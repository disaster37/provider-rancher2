// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	appv2 "github.com/disaster37/provider-rancher2/internal/controller/namespaced/app/appv2"
	catalogv2 "github.com/disaster37/provider-rancher2/internal/controller/namespaced/app/catalogv2"
	authconfigactivedirectory "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigactivedirectory"
	authconfigadfs "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigadfs"
	authconfigazuread "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigazuread"
	authconfigfreeipa "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigfreeipa"
	authconfiggithub "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfiggithub"
	authconfigkeycloak "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigkeycloak"
	authconfigokta "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigokta"
	authconfigopenldap "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigopenldap"
	authconfigping "github.com/disaster37/provider-rancher2/internal/controller/namespaced/auth/authconfigping"
	certificate "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/certificate"
	configmapv2 "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/configmapv2"
	namespace "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/namespace"
	project "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/project"
	registry "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/registry"
	secret "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/secret"
	secretv2 "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/secretv2"
	storageclassv2 "github.com/disaster37/provider-rancher2/internal/controller/namespaced/k8s/storageclassv2"
	providerconfig "github.com/disaster37/provider-rancher2/internal/controller/namespaced/providerconfig"
	bootstrap "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/bootstrap"
	cloudcredential "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/cloudcredential"
	cluster "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/cluster"
	clusterdriver "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/clusterdriver"
	clusterroletemplatebinding "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/clusterroletemplatebinding"
	clustersync "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/clustersync"
	clusterv2 "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/clusterv2"
	customusertoken "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/customusertoken"
	feature "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/feature"
	globalrole "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/globalrole"
	globalrolebinding "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/globalrolebinding"
	machineconfigv2 "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/machineconfigv2"
	nodedriver "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/nodedriver"
	nodepool "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/nodepool"
	roletemplate "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/roletemplate"
	setting "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/setting"
	token "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/token"
	user "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rancher/user"
	projectroletemplatebinding "github.com/disaster37/provider-rancher2/internal/controller/namespaced/rbac/projectroletemplatebinding"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appv2.Setup,
		catalogv2.Setup,
		authconfigactivedirectory.Setup,
		authconfigadfs.Setup,
		authconfigazuread.Setup,
		authconfigfreeipa.Setup,
		authconfiggithub.Setup,
		authconfigkeycloak.Setup,
		authconfigokta.Setup,
		authconfigopenldap.Setup,
		authconfigping.Setup,
		certificate.Setup,
		configmapv2.Setup,
		namespace.Setup,
		project.Setup,
		registry.Setup,
		secret.Setup,
		secretv2.Setup,
		storageclassv2.Setup,
		providerconfig.Setup,
		bootstrap.Setup,
		cloudcredential.Setup,
		cluster.Setup,
		clusterdriver.Setup,
		clusterroletemplatebinding.Setup,
		clustersync.Setup,
		clusterv2.Setup,
		customusertoken.Setup,
		feature.Setup,
		globalrole.Setup,
		globalrolebinding.Setup,
		machineconfigv2.Setup,
		nodedriver.Setup,
		nodepool.Setup,
		roletemplate.Setup,
		setting.Setup,
		token.Setup,
		user.Setup,
		projectroletemplatebinding.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		appv2.SetupGated,
		catalogv2.SetupGated,
		authconfigactivedirectory.SetupGated,
		authconfigadfs.SetupGated,
		authconfigazuread.SetupGated,
		authconfigfreeipa.SetupGated,
		authconfiggithub.SetupGated,
		authconfigkeycloak.SetupGated,
		authconfigokta.SetupGated,
		authconfigopenldap.SetupGated,
		authconfigping.SetupGated,
		certificate.SetupGated,
		configmapv2.SetupGated,
		namespace.SetupGated,
		project.SetupGated,
		registry.SetupGated,
		secret.SetupGated,
		secretv2.SetupGated,
		storageclassv2.SetupGated,
		providerconfig.SetupGated,
		bootstrap.SetupGated,
		cloudcredential.SetupGated,
		cluster.SetupGated,
		clusterdriver.SetupGated,
		clusterroletemplatebinding.SetupGated,
		clustersync.SetupGated,
		clusterv2.SetupGated,
		customusertoken.SetupGated,
		feature.SetupGated,
		globalrole.SetupGated,
		globalrolebinding.SetupGated,
		machineconfigv2.SetupGated,
		nodedriver.SetupGated,
		nodepool.SetupGated,
		roletemplate.SetupGated,
		setting.SetupGated,
		token.SetupGated,
		user.SetupGated,
		projectroletemplatebinding.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
