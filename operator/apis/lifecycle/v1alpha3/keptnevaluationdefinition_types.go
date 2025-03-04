/*
Copyright 2022.

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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeptnEvaluationDefinitionSpec defines the desired state of KeptnEvaluationDefinition
type KeptnEvaluationDefinitionSpec struct {
	// Objectives is a list of objectives that have to be met for a KeptnEvaluation referencing this
	// KeptnEvaluationDefinition to be successful.
	Objectives []Objective `json:"objectives"`
}

type Objective struct {
	// KeptnMetricRef references the KeptnMetric that should be evaluated.
	KeptnMetricRef KeptnMetricReference `json:"keptnMetricRef"`
	// EvaluationTarget specifies the target value for the references KeptnMetric.
	// Needs to start with either '<' or '>', followed by the target value (e.g. '<10').
	EvaluationTarget string `json:"evaluationTarget"`
}

type KeptnMetricReference struct {
	// Name is the name of the referenced KeptnMetric.
	Name string `json:"name"`
	// Namespace is the namespace where the referenced KeptnMetric is located.
	Namespace string `json:"namespace,omitempty"`
}

// KeptnEvaluationDefinitionStatus defines the observed state of KeptnEvaluationDefinition.
type KeptnEvaluationDefinitionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion
//+kubebuilder:resource:path=keptnevaluationdefinitions,shortName=ked

// KeptnEvaluationDefinition is the Schema for the keptnevaluationdefinitions API
type KeptnEvaluationDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec describes the desired state of the KeptnEvaluationDefinition.
	Spec KeptnEvaluationDefinitionSpec `json:"spec,omitempty"`
	// Status describes the current state of the KeptnEvaluationDefinition.
	Status KeptnEvaluationDefinitionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// KeptnEvaluationDefinitionList contains a list of KeptnEvaluationDefinition
type KeptnEvaluationDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeptnEvaluationDefinition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeptnEvaluationDefinition{}, &KeptnEvaluationDefinitionList{})
}
