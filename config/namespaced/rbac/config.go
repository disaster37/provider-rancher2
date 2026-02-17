package rbac

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for rbac group of resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_project_role_template_binding", func(r *config.Resource) {
		r.ShortGroup = "rbac"
		r.MarkAsRequired(
			"name",
		)
		r.References["project_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
		r.References["role_template_id"] = config.Reference{
			TerraformName: "rancher2_role_template",
		}
	})
}
