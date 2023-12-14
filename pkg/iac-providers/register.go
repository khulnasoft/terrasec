package iacprovider

import (
	"reflect"
)

// map of supported IaC providers
var supportedIacProviders = make(map[supportedIacType]map[supportedIacVersion]reflect.Type)

// map of default IaC versions for each IaC provider type
var defaultIacVersions = make(map[supportedIacType]supportedIacVersion)

// RegisterIacProvider registers an IaC provider for terrasec
// if the Iac provider does not have a version, it can be kept empty
func RegisterIacProvider(iacType supportedIacType, iacVersion, defaultIacVersion supportedIacVersion, iacProvider reflect.Type) {

	if iacVersion == "" {
		iacVersion = defaultIacVersion
	}

	if IacVersionMap, IacExists := supportedIacProviders[iacType]; IacExists {
		IacVersionMap[iacVersion] = iacProvider
	} else {
		// version support
		supportedIacVersions := make(map[supportedIacVersion]reflect.Type)
		supportedIacVersions[iacVersion] = iacProvider
		supportedIacProviders[iacType] = supportedIacVersions
	}

	// default version
	defaultIacVersions[iacType] = defaultIacVersion
}
