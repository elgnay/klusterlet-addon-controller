// (c) Copyright IBM Corporation 2019, 2020. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// U.S. Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.
//
// Copyright (c) Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package v1

import (
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	agentv1 "github.com/open-cluster-management/klusterlet-addon-controller/pkg/apis/agent/v1"
)

// constants for work manager
const (
	WorkManager             = "klusterlet-addon-workmgr"
	WorkMgr                 = "workmgr"
	RequiresHubKubeConfig   = true
	managedClusterAddOnName = "work-manager"
	addonNameEnv            = "WORKMGR_NAME"
)

var log = logf.Log.WithName("workmgr")

type AddonWorkMgr struct{}

func (addon AddonWorkMgr) IsEnabled(instance *agentv1.KlusterletAddonConfig) bool {
	return true
}

func (addon AddonWorkMgr) CheckHubKubeconfigRequired() bool {
	return RequiresHubKubeConfig
}

func (addon AddonWorkMgr) GetAddonName() string {
	return WorkMgr
}

func (addon AddonWorkMgr) GetManagedClusterAddOnName() string {
	if n := os.Getenv(addonNameEnv); len(n) != 0 {
		return n
	}
	log.Info("failed to get addon name from env var " + addonNameEnv)
	return managedClusterAddOnName
}

func (addon AddonWorkMgr) NewAddonCR(
	instance *agentv1.KlusterletAddonConfig,
	namespace string,
) (runtime.Object, error) {
	return newWorkManagerCR(instance, namespace)
}

// newWorkManagerCR - create CR for component work manager
func newWorkManagerCR(
	instance *agentv1.KlusterletAddonConfig,
	namespace string,
) (*agentv1.WorkManager, error) {
	labels := map[string]string{
		"app": instance.Name,
	}

	gv := agentv1.GlobalValues{
		ImagePullPolicy: instance.Spec.ImagePullPolicy,
		ImagePullSecret: instance.Spec.ImagePullSecret,
		ImageOverrides:  make(map[string]string, 1),
	}

	imageRepository, err := instance.GetImage("multicloud_manager")
	if err != nil {
		log.Error(err, "Fail to get Image", "Component.Name", "work-manager")
		return nil, err
	}
	gv.ImageOverrides["multicloud_manager"] = imageRepository

	if imageRepositoryLease, err := instance.GetImage("klusterlet_addon_lease_controller"); err != nil {
		log.Error(err, "Fail to get Image", "Image.Key", "klusterlet_addon_lease_controller")
	} else {
		gv.ImageOverrides["klusterlet_addon_lease_controller"] = imageRepositoryLease
	}

	clusterLabels := instance.Spec.ClusterLabels

	return &agentv1.WorkManager{
		TypeMeta: metav1.TypeMeta{
			APIVersion: agentv1.SchemeGroupVersion.String(),
			Kind:       "WorkManager",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      WorkManager,
			Namespace: namespace,
			Labels:    labels,
		},
		Spec: agentv1.WorkManagerSpec{
			FullNameOverride: WorkManager,

			ClusterName:      instance.Spec.ClusterName,
			ClusterNamespace: instance.Spec.ClusterNamespace,
			ClusterLabels:    clusterLabels,

			HubKubeconfigSecret: WorkMgr + "-hub-kubeconfig",

			GlobalValues: gv,
		},
	}, nil
}
