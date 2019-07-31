//Package v1alpha1 Defines the API to support Multicluster Endpoints (klusterlets).
//IBM Confidential
//OCO Source Materials
//5737-E67
//(C) Copyright IBM Corporation 2019 All Rights Reserved
//The source code for this program is not published or otherwise divested of its trade secrets, irrespective of what has been deposited with the U.S. Copyright Office.
package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IMPORTANT: Run "operator-sdk generate k8s" to regenerate code after modifying this file

// KlusterletServiceSpec defines the desired state of KlusterletService
// +k8s:openapi-gen=true
type KlusterletServiceSpec struct {
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// +kubebuilder:validation:MinLength=1
	Version string `json:"version"`

	// +kubebuilder:validation:MinLength=1
	ClusterName string `json:"clusterName"`

	// +kubebuilder:validation:MinLength=1
	ClusterNamespace string `json:"clusterNamespace"`

	ClusterLabels map[string]string `json:"clusterLabels"`

	BootStrapConfig         map[string]string                   `json:"bootstrapConfig"`
	TillerIntegration       KlusterletTillerIntegrationSpec     `json:"tillerIntegration"`
	PrometheusIntegration   KlusterletPrometheusIntegrationSpec `json:"prometheusIntegration"`
	TopologyCollectorConfig KlusterletTopologyCollectorSpec     `json:"topologyCollector"`
	SearchCollectorConfig   KlusterletSearchCollectorSpec       `json:"searchCollector"`

	// +kubebuilder:validation:MinLength=1
	ImageRegistry string `json:"imageRegistry"`

	// +kubebuilder:validation:MinLength=1
	ImagePullSecret string `json:"imagePullSecret"`

	// +kubebuilder:validation:Enum=Always,Never,IfNotPresent
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// Override KlusterletOverride `json:"override,omitempty"`
}

// KlusterletOverride defines configuration override for all components
// +k8s:openapi-gen=true
type KlusterletOverride struct {
	CertManager       CertManager       `json:"certManager"`
	ConnectionManager ConnectionManager `json:"connectionManager"`
	// PolicyController  PolicyController  `json:"policyController,omitempty"`
	SearchCollector SearchCollector `json:"searchCollector,omitempty"`
	// ServiceRegistry   ServiceRegistry   `json:"serviceRegistry,omitempty"`
	WorkManager       WorkManager       `json:"workManager,omitempty"`
	Tiller            Tiller            `json:"tiller"`
	TopologyCollector TopologyCollector `json:"topologyCollector,omitempty"`
}

// KlusterletSearchCollectorSpec defines configuration for the SearchCollector component
// +k8s:openapi-gen=true
type KlusterletSearchCollectorSpec struct {
	Enabled bool `json:"enabled"`
}

// KlusterletWorkManagerSpec defines configuration for the WorkManager component
// +k8s:openapi-gen=true
type KlusterletWorkManagerSpec struct {
	ClusterLabels map[string]string `json:"clusterLabels"`
}

// KlusterletTopologyCollectorSpec defines configuration for the WorkManager Promtheus Integration
// +k8s:openapi-gen=true
type KlusterletTopologyCollectorSpec struct {
	Enabled                 bool  `json:"enabled"`
	CollectorUpdateInterval int32 `json:"updateInterval"`
}

// KlusterletPrometheusIntegrationSpec defines configuration for the WorkManager Promtheus Integration
// +k8s:openapi-gen=true
type KlusterletPrometheusIntegrationSpec struct {
	Enabled bool `json:"enabled"`
}

// KlusterletTillerIntegrationSpec defines configuration for the WorkManager Tiller Integration
// +k8s:openapi-gen=true
type KlusterletTillerIntegrationSpec struct {
	Enabled bool `json:"enabled"`
}

// KlusterletConnectionManagerSpec defines configuration for the ConnectionManager component
// +k8s:openapi-gen=true
type KlusterletConnectionManagerSpec struct {
	BootStrapConfig map[string]string `json:"bootstrapConfig"`
}

// KlusterletServiceStatus defines the observed state of KlusterletService
// +k8s:openapi-gen=true
type KlusterletServiceStatus struct {
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	Test string `json:"test"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KlusterletService is the Schema for the klusterletservices API
// +k8s:openapi-gen=true
type KlusterletService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KlusterletServiceSpec   `json:"spec,omitempty"`
	Status KlusterletServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KlusterletServiceList contains a list of KlusterletService
type KlusterletServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KlusterletService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KlusterletService{}, &KlusterletServiceList{})
}