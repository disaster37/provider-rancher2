package rancher

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for rancher group of resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_cluster", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"name",
		)

		/*
			r.References["cluster_template_id"] = config.Reference{
				TerraformName: "rancher2_cluster_template",
			}
		*/
	})

	p.AddResourceConfigurator("rancher2_cluster_template", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired("name")
	})

	p.AddResourceConfigurator("rancher2_bootstrap", func(r *config.Resource) {
		r.ShortGroup = "rancher"
	})

	p.AddResourceConfigurator("rancher2_cloud_credential", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"name",
		)
	})

	p.AddResourceConfigurator("rancher2_cluster_driver", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"active",
			"builtin",
			"name",
			"url",
		)
	})

	p.AddResourceConfigurator("rancher2_cluster_role_template_binding", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"name",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
		r.References["role_template_id"] = config.Reference{
			TerraformName: "rancher2_role_template",
		}
	})

	p.AddResourceConfigurator("rancher2_cluster_sync", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
		r.References["node_pool_ids"] = config.Reference{
			TerraformName: "rancher2_node_pool",
		}
	})

	p.AddResourceConfigurator("rancher2_cluster_v2", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.Kind = "ClusterV2"
		r.MarkAsRequired(
			"name",
			"kubernetes_version",
		)
	})

	p.AddResourceConfigurator("rancher2_custom_user_token", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"password",
		)
		r.References["username"] = config.Reference{
			TerraformName: "rancher2_user",
		}

	})

	p.AddResourceConfigurator("rancher2_etcd_backup", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"name",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
	})

	p.AddResourceConfigurator("rancher2_feature", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired("name")
	})

	p.AddResourceConfigurator("rancher2_global_role", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired("name")
	})

	p.AddResourceConfigurator("rancher2_global_role_binding", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.References["global_role_id"] = config.Reference{
			TerraformName: "rancher2_global_role",
		}
	})

	p.AddResourceConfigurator("rancher2_machine_config_v2", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired("generate_name")
	})

	p.AddResourceConfigurator("rancher2_node_driver", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"active",
			"builtin",
			"name",
			"url",
		)
	})

	p.AddResourceConfigurator("rancher2_node_pool", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"name",
			"hostname_prefix",
		)
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
		/*
			r.References["node_template_id"] = config.Reference{
				TerraformName: "rancher2_node_template",
			}
		*/
		r.References["cloud_credential_id"] = config.Reference{
			TerraformName: "rancher2_cloud_credential",
		}
	})

	p.AddResourceConfigurator("rancher2_node_template", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired("name")
		r.References["driver_id"] = config.Reference{
			TerraformName: "rancher2_node_driver",
		}
		r.References["cloud_credential_id"] = config.Reference{
			TerraformName: "rancher2_cloud_credential",
		}
	})

	p.AddResourceConfigurator("rancher2_role_template", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired("name")
	})

	p.AddResourceConfigurator("rancher2_setting", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"name",
			"value",
		)
	})

	p.AddResourceConfigurator("rancher2_token", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.References["cluster_id"] = config.Reference{
			TerraformName: "rancher2_cluster_v2",
		}
	})

	p.AddResourceConfigurator("rancher2_user", func(r *config.Resource) {
		r.ShortGroup = "rancher"
		r.MarkAsRequired(
			"username",
			"password",
		)
	})
}
