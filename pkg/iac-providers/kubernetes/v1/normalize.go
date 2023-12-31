package k8sv1

import (
	"encoding/json"
	"fmt"

	yamltojson "github.com/ghodss/yaml"
	"github.com/iancoleman/strcase"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/utils"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	terrasecMaxSeverity = "runterrasec.io/maxseverity"
	terrasecMinSeverity = "runterrasec.io/minseverity"
)

var (
	errUnsupportedDoc = fmt.Errorf("unsupported document type")
	// ErrNoKind is returned when the "kind" key is not available (not a valid kubernetes resource)
	ErrNoKind = fmt.Errorf("kind does not exist")

	infileInstructionNotPresentLog = "%s not present for resource: %s"
)

// k8sMetadata is used to pull the name, namespace types and annotations for a given resource
type k8sMetadata struct {
	Name         string                 `yaml:"name" json:"name"`
	GenerateName string                 `yaml:"generateName,omitempty" json:"generateName,omitempty"`
	Namespace    string                 `yaml:"namespace" json:"namespace"`
	Annotations  map[string]interface{} `yaml:"annotations" json:"annotations"`
}

// NameOrGenerateName gets the metadata's Name member, or if Name is not set then GenerateName (for CRDs, for example)
func (m k8sMetadata) NameOrGenerateName() string {
	if len(m.Name) > 0 {
		return m.Name
	}
	return m.GenerateName
}

// k8sResource is a generic struct to handle all k8s resource types
type k8sResource struct {
	APIVersion string      `yaml:"apiVersion" json:"apiVersion"`
	Kind       string      `yaml:"kind" json:"kind"`
	Metadata   k8sMetadata `yaml:"metadata" json:"metadata"`
}

// extractResource takes the incoming document and extracts the resource using a go struct
// returns the resource data and raw json byte output ready for normalization
func (k *K8sV1) extractResource(doc *utils.IacDocument) (*k8sResource, *[]byte, error) {
	var resource k8sResource
	switch doc.Type {
	case utils.YAMLDoc:
		data, err := yamltojson.YAMLToJSON(doc.Data)
		if err != nil {
			return nil, nil, err
		}
		err = yaml.Unmarshal(data, &resource)
		if err != nil {
			return nil, nil, err
		}
		return &resource, &data, nil
	case utils.JSONDoc:
		err := json.Unmarshal(doc.Data, &resource)
		if err != nil {
			return nil, nil, err
		}
		return &resource, &doc.Data, nil
	default:
		return nil, nil, errUnsupportedDoc
	}
}

// getNormalizedName returns the normalized name
// this matches the terraform-defined resource type when applicable
func (k *K8sV1) getNormalizedName(kind string) string {
	var name string
	switch kind {
	case "DaemonSet":
		name = kubernetesTypeName + "_daemonset"
	default:
		name = kubernetesTypeName + "_" + strcase.ToSnake(kind)
	}
	return name
}

// Normalize takes the input document and normalizes it
func (k *K8sV1) Normalize(doc *utils.IacDocument) (*output.ResourceConfig, error) {

	resource, jsonData, err := k.extractResource(doc)
	if err != nil {
		return nil, err
	}

	var resourceConfig output.ResourceConfig
	resourceConfig.ContainerImages = make([]output.ContainerDetails, 0)
	resourceConfig.InitContainerImages = make([]output.ContainerDetails, 0)
	var containerImages, initContainerImages []output.ContainerDetails
	resourceConfig.Type = k.getNormalizedName(resource.Kind)

	switch resource.Kind {
	case "":
		// error case
		return nil, ErrNoKind
	// non-namespaced resources
	case "ClusterRole":
		fallthrough
	// pod and all kinds of workloads
	case "Pod", "Deployment", "ReplicaSet", "ReplicationController", "Job", "CronJob", "StatefulSet", "DaemonSet":
		containerImages, initContainerImages, err = k.extractContainerImages(resource.Kind, doc)
		if err != nil {
			return nil, err
		}
		fallthrough
	default:
		// namespaced-resources
		namespace := resource.Metadata.Namespace
		if namespace == "" {
			namespace = "default"
		}

		resourceConfig.ID = resourceConfig.Type + "." + resource.Metadata.NameOrGenerateName() + "-" + namespace
	}

	resourceConfig.ContainerImages = append(resourceConfig.ContainerImages, containerImages...)
	resourceConfig.InitContainerImages = append(resourceConfig.InitContainerImages, initContainerImages...)

	// read and update skip rules, if present
	skipRules := utils.ReadSkipRulesFromMap(resource.Metadata.Annotations, resourceConfig.ID)
	if skipRules != nil {
		resourceConfig.SkipRules = append(resourceConfig.SkipRules, skipRules...)
	}

	maxSeverity, minSeverity := readMinMaxSeverityFromAnnotations(resource.Metadata.Annotations, resourceConfig.ID)

	resourceConfig.MaxSeverity = maxSeverity
	resourceConfig.MinSeverity = minSeverity

	configData := make(map[string]interface{})
	if err = json.Unmarshal(*jsonData, &configData); err != nil {
		return nil, err
	}

	resourceConfig.Name = resource.Metadata.NameOrGenerateName()
	resourceConfig.Config = configData

	return &resourceConfig, nil
}

// readMinMaxSeverityFromAnnotations finds the min max severity values set in annotations for the resource
func readMinMaxSeverityFromAnnotations(annotations map[string]interface{}, resourceID string) (maxSeverity, minSeverity string) {
	var (
		minSeverityAnnotation interface{}
		maxSeverityAnnotation interface{}
		ok                    bool
	)
	if minSeverityAnnotation, ok = annotations[terrasecMinSeverity]; !ok {
		zap.S().Debugf(infileInstructionNotPresentLog, terrasecMinSeverity, resourceID)
	} else if minSeverity, ok = minSeverityAnnotation.(string); !ok {
		zap.S().Debugf("%s must be a string containing value as (High | Low| Medium)", terrasecMinSeverity)
	}
	if maxSeverityAnnotation, ok = annotations[terrasecMaxSeverity]; !ok {
		zap.S().Debugf(infileInstructionNotPresentLog, terrasecMaxSeverity, resourceID)
	} else if maxSeverity, ok = maxSeverityAnnotation.(string); !ok {
		zap.S().Debugf("%s must be a string containing value as (High | Low| Medium)", terrasecMaxSeverity)
	}
	return
}
