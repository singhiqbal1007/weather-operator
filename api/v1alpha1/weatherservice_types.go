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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WeatherServiceSpec defines the desired state of WeatherService
type WeatherServiceSpec struct {
	City string `json:"city,omitempty"`
	Days int    `json:"days,omitempty"`
}

// WeatherServiceStatus defines the observed state of WeatherService
type WeatherServiceStatus struct {
	Executed bool `json:"executed,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WeatherService is the Schema for the weatherservices API
type WeatherService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WeatherServiceSpec   `json:"spec,omitempty"`
	Status WeatherServiceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WeatherServiceList contains a list of WeatherService
type WeatherServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WeatherService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WeatherService{}, &WeatherServiceList{})
}
