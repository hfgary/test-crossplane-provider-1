/*
Copyright 2025 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// MyTypeAParameters are the configurable fields of a MyTypeA.
type MyTypeAParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// MyTypeAObservation are the observable fields of a MyTypeA.
type MyTypeAObservation struct {
	ConfigurableField string `json:"configurableField"`
	ObservableField   string `json:"observableField,omitempty"`
}

// A MyTypeASpec defines the desired state of a MyTypeA.
type MyTypeASpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MyTypeAParameters `json:"forProvider"`
}

// A MyTypeAStatus represents the observed state of a MyTypeA.
type MyTypeAStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MyTypeAObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A MyTypeA is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,myprovidera}
type MyTypeA struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MyTypeASpec   `json:"spec"`
	Status MyTypeAStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MyTypeAList contains a list of MyTypeA
type MyTypeAList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MyTypeA `json:"items"`
}

// MyTypeA type metadata.
var (
	MyTypeAKind             = reflect.TypeOf(MyTypeA{}).Name()
	MyTypeAGroupKind        = schema.GroupKind{Group: Group, Kind: MyTypeAKind}.String()
	MyTypeAKindAPIVersion   = MyTypeAKind + "." + SchemeGroupVersion.String()
	MyTypeAGroupVersionKind = SchemeGroupVersion.WithKind(MyTypeAKind)
)

func init() {
	SchemeBuilder.Register(&MyTypeA{}, &MyTypeAList{})
}
