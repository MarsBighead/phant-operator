// Authors: Marsbighead <duanhmhy@126.com>
//
// # Copyright (c) 2024 Marsbighead
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package v1alpha1

import (
	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigUpdateStrategy represents the strategy to update configuration
type ConfigUpdateStrategy string

const (
	// ConfigUpdateStrategyInPlace update the configmap without changing the name
	ConfigUpdateStrategyInPlace ConfigUpdateStrategy = "InPlace"
	// ConfigUpdateStrategyRollingUpdate generate different configmap on configuration update and
	// try to rolling-update the pod controller (e.g. statefulset) to apply updates.
	ConfigUpdateStrategyRollingUpdate ConfigUpdateStrategy = "RollingUpdate"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresCluster is the control script's spec
//
// +k8s:openapi-gen=true
// +kubebuilder:resource:shortName="pgc"
type PostgresCluster struct {
	metav1.TypeMeta `json:",inline"`
	// +k8s:openapi-gen=false
	metav1.ObjectMeta `json:"metadata"`

	// Spec defines the behavior of a tidb cluster
	Spec PostgresClusterSpec `json:"spec"`

	// +k8s:openapi-gen=false
	// Most recently observed status of the tidb cluster
	Status PostgresClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PostgresClusterList is PostgresCluster list
// +k8s:openapi-gen=true
type PostgresClusterList struct {
	metav1.TypeMeta `json:",inline"`
	// +k8s:openapi-gen=false
	metav1.ListMeta `json:"metadata"`

	Items []PostgresCluster `json:"items"`
}

// NodeHealthStatus is the status for a NodeHealth resource
type PostgresClusterStatus struct {
	AvailableReplicas int32                      `json:"availableReplicas"`
	Conditions        []PostgresClusterCondition `json:"healthStatus"`
}

// PostgresClusterConditionType represents a tidb cluster condition value.
type PostgresClusterConditionType string

const (
	// TidbClusterReady indicates that the tidb cluster is ready or not.
	// This is defined as:
	// - All statefulsets are up to date (currentRevision == updateRevision).
	// - All PD members are healthy.
	// - All TiDB pods are healthy.
	// - All TiKV stores are up.
	// - All TiFlash stores are up.
	PostgresClusterReady PostgresClusterConditionType = "Ready"
)

// TidbClusterCondition describes the state of a tidb cluster at a certain point.
type PostgresClusterCondition struct {
	// Type of the condition.
	Type PostgresClusterConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	// +nullable
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	// +optional
	// +nullable
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// A human readable message indicating details about the transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// PostgresClusterSpec describes the attributes that a user creates on a tidb cluster
// +k8s:openapi-gen=true
type PostgresClusterSpec struct {

	// Specify a Service Account
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// Postgres cluster
	// +optional
	Postgres *PostgresSpec `json:"postgres,omitempty"`
}

// ComponentSpec is the base spec of each component, the fields should always accessed by the Basic<Component>Spec() method to respect the cluster-level properties
// +k8s:openapi-gen=true
type ComponentSpec struct {
	// (Deprecated) Image of the component
	// Use `baseImage` and `version` instead
	// +k8s:openapi-gen=false
	Image string `json:"image,omitempty"`

	// Version of the component. Override the cluster-level version if non-empty
	// Optional: Defaults to cluster-level setting
	// +optional
	Version *string `json:"version,omitempty"`

	// ImagePullPolicy of the component. Override the cluster-level imagePullPolicy if present
	// Optional: Defaults to cluster-level setting
	// +optional
	ImagePullPolicy *corev1.PullPolicy `json:"imagePullPolicy,omitempty"`

	// ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images.
	// +optional
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`

	// Whether Hostnetwork of the component is enabled. Override the cluster-level setting if present
	// Optional: Defaults to cluster-level setting
	// +optional
	HostNetwork *bool `json:"hostNetwork,omitempty"`

	// Affinity of the component. Override the cluster-level setting if present.
	// Optional: Defaults to cluster-level setting
	// +optional
	Affinity *corev1.Affinity `json:"affinity,omitempty"`

	// PriorityClassName of the component. Override the cluster-level one if present
	// Optional: Defaults to cluster-level setting
	// +optional
	PriorityClassName *string `json:"priorityClassName,omitempty"`

	// SchedulerName of the component. Override the cluster-level one if present
	// Optional: Defaults to cluster-level setting
	// +optional
	SchedulerName *string `json:"schedulerName,omitempty"`

	// NodeSelector of the component. Merged into the cluster-level nodeSelector if non-empty
	// Optional: Defaults to cluster-level setting
	// +optional
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// Annotations for the component. Merge into the cluster-level annotations if non-empty
	// Optional: Defaults to cluster-level setting
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Labels for the component. Merge into the cluster-level labels if non-empty
	// Optional: Defaults to cluster-level setting
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Tolerations of the component. Override the cluster-level tolerations if non-empty
	// Optional: Defaults to cluster-level setting
	// +optional
	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	// PodSecurityContext of the component
	// +optional
	PodSecurityContext *corev1.PodSecurityContext `json:"podSecurityContext,omitempty"`

	// ConfigUpostgresateStrategy of the component. Override the cluster-level upostgresateStrategy if present
	// Optional: Defaults to cluster-level setting
	// +optional
	ConfigUpostgresateStrategy *ConfigUpdateStrategy `json:"configUpostgresateStrategy,omitempty"`

	// List of environment variables to set in the container, like v1.Container.Env.
	// Note that the following env names cannot be used and will be overridden by PostgreSQL Operator builtin envs
	// - NAMESPACE
	// - TZ
	// - SERVICE_NAME
	// - PEER_SERVICE_NAME
	// - HEADLESS_SERVICE_NAME
	// - SET_NAME
	// - HOSTNAME
	// - CLUSTER_NAME
	// - POD_NAME
	// - BINLOG_ENABLED
	// - SLOW_LOG_FILE
	// +optional
	Env []corev1.EnvVar `json:"env,omitempty"`

	// Extend the use scenarios for env
	// +optional
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty"`

	// Init containers of the components
	// +optional
	InitContainers []corev1.Container `json:"initContainers,omitempty"`

	// Additional containers of the component.
	// If the container names in this field match with the ones generated by
	// PostgreSQL Operator, the container configurations will be merged into the
	// containers generated by PostgreSQL Operator via strategic merge patch.
	// If the container names in this field do not match with the ones
	// generated by PostgreSQL Operator, the container configurations will be
	// appended to the Pod container spec directly.
	// +optional
	AdditionalContainers []corev1.Container `json:"additionalContainers,omitempty"`

	// Additional volumes of component pod.
	// +optional
	AdditionalVolumes []corev1.Volume `json:"additionalVolumes,omitempty"`

	// Additional volume mounts of component pod.
	AdditionalVolumeMounts []corev1.VolumeMount `json:"additionalVolumeMounts,omitempty"`

	// DNSConfig Specifies the DNS parameters of a pod.
	// +optional
	DNSConfig *corev1.PodDNSConfig `json:"dnsConfig,omitempty"`

	// DNSPolicy Specifies the DNSPolicy parameters of a pod.
	// +optional
	DNSPolicy corev1.DNSPolicy `json:"dnsPolicy,omitempty"`

	// Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
	// Value must be non-negative integer. The value zero indicates delete immediately.
	// If this value is nil, the default grace period will be used instead.
	// The grace period is the duration in seconds after the processes running in the pod are sent
	// a termination signal and the time when the processes are forcibly halted with a kill signal.
	// Set this value longer than the expected cleanup time for your process.
	// Defaults to 30 seconds.
	// +optional
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`

	// StatefulSetUpostgresateStrategy indicates the StatefulSetUpostgresateStrategy that will be
	// employed to upostgresate Pods in the StatefulSet when a revision is made to
	// Template.
	// +optional
	StatefulSetUpdateStrategy apps.StatefulSetUpdateStrategyType `json:"statefulSetUpdateStrategy,omitempty"`

	// PodManagementPolicy of PostgreSQL cluster StatefulSets
	// +optional
	PodManagementPolicy apps.PodManagementPolicyType `json:"podManagementPolicy,omitempty"`
}

// ServiceSpec specifies the service object in k8s
// +k8s:openapi-gen=true
type ServiceSpec struct {
	// Type of the real kubernetes service
	Type corev1.ServiceType `json:"type,omitempty"`

	// Additional annotations for the service
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`

	// Additional labels for the service
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// LoadBalancerIP is the loadBalancerIP of service
	// Optional: Defaults to omitted
	// +optional
	LoadBalancerIP *string `json:"loadBalancerIP,omitempty"`

	// ClusterIP is the clusterIP of service
	// +optional
	ClusterIP *string `json:"clusterIP,omitempty"`

	// PortName is the name of service port
	// +optional
	PortName *string `json:"portName,omitempty"`

	// The port that will be exposed by this service.
	//
	// NOTE: only used for PostgreSQL
	//
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	Port *int32 `json:"port,omitempty"`

	// LoadBalancerSourceRanges is the loadBalancerSourceRanges of service
	// If specified and supported by the platform, this will restrict traffic through the cloud-provider
	// load-balancer will be restricted to the specified client IPs. This field will be ignored if the
	// cloud-provider does not support the feature."
	// More info: https://kubernetes.io/docs/concepts/services-networking/service/#aws-nlb-support
	// Optional: Defaults to omitted
	// +optional
	LoadBalancerSourceRanges []string `json:"loadBalancerSourceRanges,omitempty"`
}

// +k8s:openapi-gen=true
// postgresSpec contains details of postgres members
type PostgresSpec struct {
	ComponentSpec               `json:",inline"`
	corev1.ResourceRequirements `json:",inline"`

	// Specify a Service Account for postgres
	ServiceAccount string `json:"serviceAccount,omitempty"`

	// The desired ready replicas
	// +kubebuilder:validation:Minimum=0
	Replicas int32 `json:"replicas"`

	// Base image of the component, image tag is now allowed during validation
	// +kubebuilder:default=pingcap/postgres
	// +optional
	BaseImage string `json:"baseImage"`

	// Service defines a Kubernetes service of postgres cluster.
	// Optional: Defaults to `.spec.services` in favor of backward compatibility
	// +optional
	Service *ServiceSpec `json:"service,omitempty"`

	// MaxFailoverCount limit the max replicas could be added in failover, 0 means no failover.
	// Optional: Defaults to 3
	// +kubebuilder:validation:Minimum=0
	// +optional
	MaxFailoverCount *int32 `json:"maxFailoverCount,omitempty"`

	// The storageClassName of the persistent volume for postgres data storage.
	// Defaults to Kubernetes default storage class.
	// +optional
	StorageClassName *string `json:"storageClassName,omitempty"`

	// StorageVolumes configure additional storage for postgres pods.
	// +optional
	StorageVolumes []StorageVolume `json:"storageVolumes,omitempty"`

	// Subdirectory within the volume to store postgres Data. By default, the data
	// is stored in the root directory of volume which is mounted at
	// /var/lib/postgres.
	// Specifying this will change the data directory to a subdirectory, e.g.
	// /var/lib/postgres/data if you set the value to "data".
	// It's dangerous to change this value for a running cluster as it will
	// upgrade your cluster to use a new storage directory.
	// Defaults to "" (volume's root).
	// +optional
	DataSubDir string `json:"dataSubDir,omitempty"`

	// Config is the Configuration of postgres-servers
	// +optional
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:validation:XPreserveUnknownFields
	Config *GenericConfig `json:"config,omitempty"`

	// TLSClientSecretName is the name of secret which stores tidb server client certificate
	// which used by Dashboard.
	// +optional
	TLSClientSecretName *string `json:"tlsClientSecretName,omitempty"`

	// (Deprecated) EnableDashboardInternalProxy would directly set `internal-proxy` in the `postgresConfig`.
	// Note that this is deprecated, we should just set `dashboard.internal-proxy` in `postgres.config`.
	// +optional
	EnableDashboardInternalProxy *bool `json:"enableDashboardInternalProxy,omitempty"`

	// MountClusterClientSecret indicates whether to mount `cluster-client-secret` to the Pod
	// +optional
	MountClusterClientSecret *bool `json:"mountClusterClientSecret,omitempty"`

	// Start up script version
	// +optional
	// +kubebuilder:validation:Enum:="";"v1"
	StartUpScriptVersion string `json:"startUpScriptVersion,omitempty"`

	// Timeout threshold when postgres get started
	// +kubebuilder:default=30
	StartTimeout int `json:"startTimeout,omitempty"`
}

// StorageVolume configures additional PVC template for StatefulSets and volumeMount for pods that mount this PVC.
// Note:
// If `MountPath` is not set, volumeMount will not be generated. (You may not want to set this field when you inject volumeMount
// in somewhere else such as Mutating Admission Webhook)
// If `StorageClassName` is not set, default to the `spec.${component}.storageClassName`
type StorageVolume struct {
	Name             string  `json:"name"`
	StorageClassName *string `json:"storageClassName,omitempty"`
	StorageSize      string  `json:"storageSize"`
	MountPath        string  `json:"mountPath,omitempty"`
}

type GenericConfig struct {
	// Export this field to make "apiequality.Semantic.DeepEqual" happy now.
	// User of GenericConfig should not directly access this field.
	MP map[string]string
}
