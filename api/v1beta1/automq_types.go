/*
Copyright 2024 cuisongliu.

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

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type S3Spec struct {
	// Endpoint is the endpoint of the S3 service
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint,omitempty"`
	// Region is the region of the S3 service
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern=^[a-zA-Z0-9-]+$
	Region string `json:"region,omitempty"`
	// AccessKeyID is the access key ID of the S3 service
	// +kubebuilder:validation:Required
	AccessKeyID string `json:"accessKeyID,omitempty"`
	// SecretAccessKey is the secret access key of the S3 service
	// +kubebuilder:validation:Required
	SecretAccessKey string `json:"secretAccessKey,omitempty"`
	// Bucket is the bucket name for storing the operations data
	// +kubebuilder:validation:Required
	Bucket string `json:"bucket,omitempty"`
	// EnablePathStyle is the flag to enable the path style. Default is false.
	// Whether to enable object storage path format. Must be set to true when using MinIO as the storage service.
	// +kubebuilder:default=false
	EnablePathStyle bool `json:"enablePathStyle,omitempty"`
}

type NodeAffinity struct {
	// Type is the type of the node affinity. Supported values are "soft" and "hard"
	// +kubebuilder:validation:Enum=soft;hard
	// +kubebuilder:validation:Required
	Type string `json:"type,omitempty"`
	// NodeSelector is the node selector for the node affinity.
	// +kubebuilder:minItems=1
	NodeSelector []NodeSelector `json:"nodeSelector,omitempty"`
	// Weight is the weight of the node affinity. When the type is "soft", the weight is used to select the node. Default is 40.
	// +kubebuilder:default=40
	Weight int32 `json:"weight,omitempty"`
}

type NodeSelector struct {
	// Key is the key of the node selector
	// +kubebuilder:validation:Required
	Key string `json:"key,omitempty"`
	// Values is the value of the node selector
	// +kubebuilder:minItems=1
	Values []string `json:"values,omitempty"`
}

type PodAffinity struct {
	// Type is the type of the node affinity. Supported values are "soft" and "hard"
	// +kubebuilder:validation:Enum=soft;hard
	// +kubebuilder:validation:Required
	Type string `json:"type,omitempty"`
	// Weight is the weight of the pod affinity. When the type is "soft", the weight is used to select the pods. Default is 40.
	// +kubebuilder:default=40
	Weight int32 `json:"weight,omitempty"`
}

type PodAntiAffinity struct {
	// Type is the type of the node anti affinity. Supported values are "soft" and "hard"
	// +kubebuilder:validation:Enum=soft;hard
	// +kubebuilder:validation:Required
	Type string `json:"type,omitempty"`
	// Weight is the weight of the pod anti affinity. When the type is "soft", the weight is used to select the pods. Default is 40.
	// +kubebuilder:default=40
	Weight int32 `json:"weight,omitempty"`
}

type AffinitySpec struct {
	// NodeAffinity is the node affinity for the pod
	NodeAffinity *NodeAffinity `json:"nodeAffinity,omitempty"`
	// PodAntiAffinity is the pod anti-affinity for the pod
	PodAntiAffinity *PodAntiAffinity `json:"podAntiAffinity,omitempty"`
	// PodAffinity is the pod affinity for the pod
	PodAffinity *PodAffinity `json:"podAffinity,omitempty"`
}

func (in *AffinitySpec) ToK8sAffinity() *v1.Affinity {
	affinity := &v1.Affinity{}
	//TODO implement the conversion
	return affinity
}

type ControllerSpec struct {
	// Replicas is the number of controller replicas
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas,omitempty"`
	// JVMOptions is the JVM options for the controller
	JVMOptions []string `json:"jvmOptions,omitempty"`
	// Envs is the environment variables for the controller
	Envs []v1.EnvVar `json:"envs,omitempty"`
	// Resource is the resource requirements for the controller
	Resource v1.ResourceRequirements `json:"resource,omitempty"`
	// Affinity is the affinity for the controller
	Affinity *AffinitySpec `json:"affinity,omitempty"`
	// StorageClass is the storage class for the controller
	StorageClass string `json:"storageClass,omitempty"`
}

type BrokerSpec struct {
	// Replicas is the number of controller replicas
	// +kubebuilder:validation:Minimum=1
	Replicas int32 `json:"replicas,omitempty"`
	// JVMOptions is the JVM options for the controller
	JVMOptions []string `json:"jvmOptions,omitempty"`
	// Envs is the environment variables for the controller
	Envs []v1.EnvVar `json:"envs,omitempty"`
	// Resource is the resource requirements for the controller
	Resource v1.ResourceRequirements `json:"resource,omitempty"`
	// Affinity is the affinity for the broker
	Affinity *AffinitySpec `json:"affinity,omitempty"`
	// StorageClass is the storage class for the controller
	StorageClass string `json:"storageClass,omitempty"`
}

// MetricsSpec is the metrics configuration for the AutoMQ
type MetricsSpec struct {
	// Enable is the flag to enable the metrics
	// +kubebuilder:validation:Required
	// +kubebuilder:default=true
	Enable bool `json:"enable,omitempty"`
	// ImportDashboard is the flag to import the dashboard.
	// +kubebuilder:default=true
	ImportDashboard bool `json:"importDashboard,omitempty"`
}

// AutoMQSpec defines the desired state of AutoMQ
type AutoMQSpec struct {
	// S3 is the S3 configuration for the AutoMQ
	// +kubebuilder:validation:Required
	S3 S3Spec `json:"s3,omitempty"`
	// ClusterID is the ID of the cluster. Default is "rZdE0DjZSrqy96PXrMUZVw"
	// +kubebuilder:validation:Required
	ClusterID string `json:"clusterID,omitempty"`
	// Image is the image of the AutoMQ
	Image string `json:"image,omitempty"`
	// NodePort is the node port of the AutoMQ
	NodePort int32 `json:"nodePort,omitempty"`
	// Metrics is the metrics configuration for the AutoMQ
	Metrics MetricsSpec `json:"metrics,omitempty"`
	// Controller is the controller configuration for the AutoMQ
	// +kubebuilder:validation:Required
	Controller ControllerSpec `json:"controller,omitempty"`
	// Broker is the broker configuration for the AutoMQ
	// +kubebuilder:validation:Required
	Broker BrokerSpec `json:"broker,omitempty"`
}

type AutoMQPhase string

// These are the valid phases of node.
const (
	AutoMQPending   AutoMQPhase = "Pending"
	AutoMQError     AutoMQPhase = "Error"
	AutoMQReady     AutoMQPhase = "Ready"
	AutoMQInProcess AutoMQPhase = "InProcess"
)

// AutoMQStatus defines the observed state of AutoMQ
type AutoMQStatus struct {
	// Phase represents the current phase of AutoMQ.
	//+kubebuilder:default:=Unknown
	Phase AutoMQPhase `json:"phase,omitempty"`
	// Conditions contains the different condition statuses for this automq.
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
	// ReadyPods is the number of ready pods for the AutoMQ
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=0
	ReadyPods int32 `json:"readyPods"`
	// ControllerReplicas is the number of controller replicas for the AutoMQ
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=0
	ControllerReplicas int32 `json:"controllerReplicas"`
	// BrokerReplicas is the number of broker replicas for the AutoMQ
	// +optional
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:default=0
	BrokerReplicas int32 `json:"brokerReplicas"`
	// ControllerAddress is the address of the controller
	// +optional
	ControllerAddresses []string `json:"controllerAddresses,omitempty"`
	// BootstrapInternalAddress is the address of the bootstrap
	// +optional
	BootstrapInternalAddress string `json:"bootstrapInternalAddress,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="Ready Pods",type=string,JSONPath=`.status.readyPods`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// AutoMQ is the Schema for the automqs API
type AutoMQ struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AutoMQSpec   `json:"spec,omitempty"`
	Status AutoMQStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AutoMQList contains a list of AutoMQ
type AutoMQList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AutoMQ `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AutoMQ{}, &AutoMQList{})
}
