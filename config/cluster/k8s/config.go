package k8s

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for k8s group of resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_certificate", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.MarkAsRequired(
			"certs",
			"key",
		)
		r.References["namespace_id"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
	})

	p.AddResourceConfigurator("rancher2_config_map_v2", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.Kind = "ConfigMapV2"
		r.MarkAsRequired(
			"name",
			"data",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
		r.References["namespace"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
	})

	p.AddResourceConfigurator("rancher2_project", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.MarkAsRequired(
			"name",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
	})

	p.AddResourceConfigurator("rancher2_registry", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.MarkAsRequired(
			"name",
			"registries",
		)
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
		r.References["namespace_id"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
	})

	p.AddResourceConfigurator("rancher2_secret", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.MarkAsRequired(
			"data",
		)
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
		r.References["namespace_id"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
	})

	p.AddResourceConfigurator("rancher2_secret_v2", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.Kind = "SecretV2"
		r.MarkAsRequired(
			"name",
			"data",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
		r.References["namespace"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
	})

	p.AddResourceConfigurator("rancher2_storage_class_v2", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.Kind = "StorageClassV2"
		r.MarkAsRequired(
			"name",
			"k8s_provisioner",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
	})

}
