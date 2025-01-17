package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GrafanaSpec defines the desired state of Grafana
type GrafanaSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Hostname                string                  `json:"hostname,omitempty"`
	Containers              []v1.Container          `json:"containers,omitempty"`
	Secrets                 []string                `json:"secrets,omitempty"`
	DashboardLabelSelectors []*metav1.LabelSelector `json:"dashboardLabelSelector,omitempty"`
	LogLevel                string                  `json:"logLevel"`
	AdminUser               string                  `json:"adminUser"`
	AdminPassword           string                  `json:"adminPassword"`
	BasicAuth               bool                    `json:"basicAuth"`
	DisableLoginForm        bool                    `json:"disableLoginForm"`
	DisableSignoutMenu      bool                    `json:"disableSignoutMenu"`
	Anonymous               bool                    `json:"anonymous"`
}

// GrafanaStatus defines the observed state of Grafana
type GrafanaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Phase            int        `json:"phase"`
	InstalledPlugins PluginList `json:"installedPlugins"`
	FailedPlugins    PluginList `json:"failedPlugins"`
}

// GrafanaPlugin contains information about a single plugin
type GrafanaPlugin struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Origin  *GrafanaDashboard `json:"-"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Grafana is the Schema for the grafanas API
// +k8s:openapi-gen=true
type Grafana struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrafanaSpec   `json:"spec,omitempty"`
	Status GrafanaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GrafanaList contains a list of Grafana
type GrafanaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Grafana `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Grafana{}, &GrafanaList{})
}
