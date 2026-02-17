package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	appCluster "github.com/disaster37/provider-rancher2/config/cluster/app"
	authCluster "github.com/disaster37/provider-rancher2/config/cluster/auth"
	k8sCluster "github.com/disaster37/provider-rancher2/config/cluster/k8s"
	rancherCluster "github.com/disaster37/provider-rancher2/config/cluster/rancher"
	rbacCluster "github.com/disaster37/provider-rancher2/config/cluster/rbac"
	appNamespaced "github.com/disaster37/provider-rancher2/config/namespaced/app"
	authNamespaced "github.com/disaster37/provider-rancher2/config/namespaced/auth"
	k8sNamespaced "github.com/disaster37/provider-rancher2/config/namespaced/k8s"
	rancherNamespaced "github.com/disaster37/provider-rancher2/config/namespaced/rancher"
	rbacNamespaced "github.com/disaster37/provider-rancher2/config/namespaced/rbac"
)

const (
	resourcePrefix = "rancher2"
	modulePath     = "github.com/disaster37/provider-rancher2"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("rancher2.contrib.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			KindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		appCluster.Configure,
		authCluster.Configure,
		k8sCluster.Configure,
		rancherCluster.Configure,
		rbacCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("rancher2.m.contrib.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			KindOverrides(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		appNamespaced.Configure,
		authNamespaced.Configure,
		k8sNamespaced.Configure,
		rancherNamespaced.Configure,
		rbacNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// KindOverrides overrides the kind of the resources given in KindMap.
func KindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		subElement := strings.Split(strings.TrimPrefix(r.Name, "rancher2_"), "_")
		for i, elem := range subElement {
			subElement[i] = cases.Title(language.English).String(elem)
		}
		r.Kind = strings.Join(subElement, "")
	}
}
