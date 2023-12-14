

package commons

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	hclConfigs "github.com/hashicorp/terraform/configs"
	"github.com/khulnasoft/terrasec/pkg/iac-providers/output"
	"github.com/khulnasoft/terrasec/pkg/utils"
	"go.uber.org/zap"
)

// CreateResourceConfig creates output.ResourceConfig
func CreateResourceConfig(managedResource *hclConfigs.Resource) (resourceConfig output.ResourceConfig, err error) {

	// read source file
	fileBytes, err := os.ReadFile(managedResource.DeclRange.Filename)
	if err != nil {
		zap.S().Errorf("failed to read terraform IaC file '%s'. error: '%v'", managedResource.DeclRange.Filename, err)
		return resourceConfig, fmt.Errorf("failed to read terraform file")
	}

	// convert resource config from hcl.Body to map[string]interface{}
	c := converter{bytes: fileBytes}
	var hclBody *hclsyntax.Body
	var ok bool
	if hclBody, ok = managedResource.Config.(*hclsyntax.Body); !ok {
		return resourceConfig, fmt.Errorf("failed type assertion for hcl.Body in *hclConfigs.Resource. error: expected hcl.Body type is *hclsyntax.Body, but got %T", managedResource.Config)
	}

	goOut, lineOut, err := c.convertBody(hclBody)
	if err != nil {
		zap.S().Errorf("failed to convert hcl.Body to go struct; resource '%s', file: '%s'. error: '%v'",
			managedResource.Name, managedResource.DeclRange.Filename, err)
		return resourceConfig, fmt.Errorf("failed to convert hcl.Body to go struct")
	}

	minSeverity, maxSeverity := utils.GetMinMaxSeverity(c.rangeSource(hclBody.Range()))

	containers, initContainers := findContainers(managedResource, goOut, hclBody)
	// create a resource config
	resourceConfig = output.ResourceConfig{
		ID:                  fmt.Sprintf("%s.%s", managedResource.Type, managedResource.Name),
		Name:                managedResource.Name,
		Type:                managedResource.Type,
		Source:              managedResource.DeclRange.Filename,
		Line:                managedResource.DeclRange.Start.Line,
		Config:              goOut,
		LineConfig:          lineOut,
		SkipRules:           utils.GetSkipRules(c.rangeSource(hclBody.Range())),
		MaxSeverity:         maxSeverity,
		MinSeverity:         minSeverity,
		ContainerImages:     containers,
		InitContainerImages: initContainers,
	}

	// successful
	zap.S().Debugf("created resource config for resource '%s', file: '%s'", resourceConfig.Name, resourceConfig.Source)
	return resourceConfig, nil
}

// findContainers finds containers defined in resource
func findContainers(managedResource *hclConfigs.Resource, jsonBody jsonObj, hclBody *hclsyntax.Body) (containers []output.ContainerDetails, initContainers []output.ContainerDetails) {
	if isKubernetesResource(managedResource) {
		containers, initContainers = extractContainerImagesFromk8sResources(managedResource, hclBody)
	} else if isAzureContainerResource(managedResource) {
		containers = fetchContainersFromAzureResource(jsonBody)
	} else if isAwsContainerResource(managedResource) {
		containers = fetchContainersFromAwsResource(jsonBody, hclBody, managedResource.DeclRange.Filename)
	}
	return
}
