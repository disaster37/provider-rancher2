package auth

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure adds configurations for auth group of resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("rancher2_auth_config_activedirectory", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"servers",
			"service_account_username",
			"service_account_password",
			"user_search_base",
			"test_username",
			"test_password",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_adfs", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"display_name_field",
			"groups_field",
			"idp_metadata_content",
			"rancher_api_host",
			"sp_cert",
			"sp_key",
			"uid_field",
			"user_name_field",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_azuread", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"application_id",
			"application_secret",
			"auth_endpoint",
			"graph_endpoint",
			"rancher_url",
			"tenant_id",
			"token_endpoint",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_freeipa", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"servers",
			"service_account_distinguished_name",
			"service_account_password",
			"user_search_base",
			"test_username",
			"test_password",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_github", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"client_id",
			"client_secret",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_keycloak", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"display_name_field",
			"groups_field",
			"idp_metadata_content",
			"rancher_api_host",
			"sp_cert",
			"sp_key",
			"uid_field",
			"user_name_field",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_okta", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"display_name_field",
			"groups_field",
			"idp_metadata_content",
			"rancher_api_host",
			"sp_cert",
			"sp_key",
			"uid_field",
			"user_name_field",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_openldap", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"servers",
			"service_account_distinguished_name",
			"service_account_password",
			"user_search_base",
			"test_username",
			"test_password",
		)
	})

	p.AddResourceConfigurator("rancher2_auth_config_ping", func(r *config.Resource) {
		r.ShortGroup = "auth"
		r.MarkAsRequired(
			"display_name_field",
			"groups_field",
			"idp_metadata_content",
			"rancher_api_host",
			"sp_cert",
			"sp_key",
			"uid_field",
			"user_name_field",
		)
	})
}
