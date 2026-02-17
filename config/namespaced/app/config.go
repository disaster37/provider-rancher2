package app

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for app group of resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_app", func(r *config.Resource) {
		r.ShortGroup = "app"
		r.MarkAsRequired(
			"catalog_nam",
			"name",
			"template_name",
		)
		r.References["catalog_name"] = config.Reference{
			TerraformName: "rancher2_catalog",
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
		r.References["target_namespace"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
	})

	p.AddResourceConfigurator("rancher2_app_v2", func(r *config.Resource) {
		r.ShortGroup = "app"
		r.Kind = "AppV2"
		r.MarkAsRequired(
			"name",
			"chart_name",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
		r.References["namespace"] = config.Reference{
			TerraformName: "rancher2_namespace",
		}
		r.References["repo_name"] = config.Reference{
			TerraformName: "rancher2_catalog_v2",
		}
	})

	p.AddResourceConfigurator("rancher2_catalog", func(r *config.Resource) {
		r.ShortGroup = "app"
		r.MarkAsRequired(
			"name",
			"url",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster",
		}
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
	})

	p.AddResourceConfigurator("rancher2_catalog_v2", func(r *config.Resource) {
		r.ShortGroup = "app"
		r.Kind = "CatalogV2"
		r.MarkAsRequired(
			"name",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
	})

	p.AddResourceConfigurator("rancher2_multi_cluster_app", func(r *config.Resource) {
		r.ShortGroup = "app"
		r.MarkAsRequired(
			"name",
			"targets",
			"template_name",
		)
		r.References["catalog_name"] = config.Reference{
			TerraformName: "rancher2_catalog_v2",
		}
		r.References["targets.project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
	})

	p.AddResourceConfigurator("rancher2_namespace", func(r *config.Resource) {
		r.ShortGroup = "k8s"
		r.MarkAsRequired(
			"name",
		)
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_project",
		}
	})

}
