package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:validation:Optional

// GlobalConfiguration defines the GlobalConfiguration resource.
type GlobalConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GlobalConfigurationSpec `json:"spec"`
}

// GlobalConfigurationSpec is the spec of the GlobalConfiguration resource.
type GlobalConfigurationSpec struct {
	Listeners []Listener `json:"listeners"`
}

// Listener defines a listener.
type Listener struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalConfigurationList is a list of the GlobalConfiguration resources.
type GlobalConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []GlobalConfiguration `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:validation:Optional

// TransportServer defines the TransportServer resource.
type TransportServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TransportServerSpec `json:"spec"`
}

// TransportServerSpec is the spec of the TransportServer resource.
type TransportServerSpec struct {
	Listener           TransportServerListener `json:"listener"`
	Upstreams          []Upstream              `json:"upstreams"`
	UpstreamParameters *UpstreamParameters     `json:"upstreamParameters"`
	Action             *Action                 `json:"action"`
}

// TransportServerListener defines a listener for a TransportServer.
type TransportServerListener struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
}

// Upstream defines an upstream.
type Upstream struct {
	Name    string `json:"name"`
	Service string `json:"service"`
	Port    int    `json:"port"`
}

// UpstreamParameters defines parameters for an upstream.
type UpstreamParameters struct {
	UDPRequests  *int `json:"udpRequests"`
	UDPResponses *int `json:"udpResponses"`
}

// Action defines an action.
type Action struct {
	Pass string `json:"pass"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TransportServerList is a list of the TransportServer resources.
type TransportServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TransportServer `json:"items"`
}
