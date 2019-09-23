package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChaosTestSpec defines the desired state of ChaosTest
// +k8s:openapi-gen=true
type ChaosTestSpec struct {
	Targets  []string `json:"targets"`
	Actions  []string `json:"actions"`
	Interval int      `json:"interval"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// ChaosTestStatus defines the observed state of ChaosTest
// +k8s:openapi-gen=true
type ChaosTestStatus struct {
  Actions []string  `json:"actions"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChaosTest is the Schema for the chaostests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ChaosTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChaosTestSpec   `json:"spec,omitempty"`
	Status ChaosTestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChaosTestList contains a list of ChaosTest
type ChaosTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChaosTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChaosTest{}, &ChaosTestList{})
}
