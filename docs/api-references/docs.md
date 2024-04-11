---
title: TiDB Operator API Document
summary: Reference of TiDB Operator API
category: how-to
---
<h1>API Document</h1>
<h2 id="phant.io/v1alpha1">phant.io/v1alpha1</h2>
<p>
<p>Package v1alpha1 is the v1alpha1 version of the API.</p>
</p>
Resource Types:
<ul><li>
<a href="#postgrescluster">PostgresCluster</a>
</li></ul>
<h3 id="postgrescluster">PostgresCluster</h3>
<p>
<p>PostgresCluster is the control script&rsquo;s spec</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code></br>
string</td>
<td>
<code>
phant.io/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code></br>
string
</td>
<td><code>PostgresCluster</code></td>
</tr>
<tr>
<td>
<code>metadata</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code></br>
<em>
<a href="#postgresclusterspec">
PostgresClusterSpec
</a>
</em>
</td>
<td>
<p>Spec defines the behavior of a tidb cluster</p>
<br/>
<br/>
<table>
<tr>
<td>
<code>serviceAccount</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify a Service Account</p>
</td>
</tr>
<tr>
<td>
<code>postgres</code></br>
<em>
<a href="#postgresspec">
PostgresSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Postgres cluster</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="#postgresclusterstatus">
PostgresClusterStatus
</a>
</em>
</td>
<td>
<p>Most recently observed status of the tidb cluster</p>
</td>
</tr>
</tbody>
</table>
<h3 id="componentspec">ComponentSpec</h3>
<p>
(<em>Appears on:</em>
<a href="#postgresspec">PostgresSpec</a>)
</p>
<p>
<p>ComponentSpec is the base spec of each component, the fields should always accessed by the Basic<Component>Spec() method to respect the cluster-level properties</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>image</code></br>
<em>
string
</em>
</td>
<td>
<p>(Deprecated) Image of the component
Use <code>baseImage</code> and <code>version</code> instead</p>
</td>
</tr>
<tr>
<td>
<code>version</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Version of the component. Override the cluster-level version if non-empty
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>imagePullPolicy</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#pullpolicy-v1-core">
Kubernetes core/v1.PullPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ImagePullPolicy of the component. Override the cluster-level imagePullPolicy if present
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecrets</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#localobjectreference-v1-core">
[]Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images.</p>
</td>
</tr>
<tr>
<td>
<code>hostNetwork</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>Whether Hostnetwork of the component is enabled. Override the cluster-level setting if present
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>affinity</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#affinity-v1-core">
Kubernetes core/v1.Affinity
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Affinity of the component. Override the cluster-level setting if present.
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>priorityClassName</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>PriorityClassName of the component. Override the cluster-level one if present
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>schedulerName</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>SchedulerName of the component. Override the cluster-level one if present
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>nodeSelector</code></br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>NodeSelector of the component. Merged into the cluster-level nodeSelector if non-empty
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code></br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Annotations for the component. Merge into the cluster-level annotations if non-empty
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>labels</code></br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Labels for the component. Merge into the cluster-level labels if non-empty
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#toleration-v1-core">
[]Kubernetes core/v1.Toleration
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Tolerations of the component. Override the cluster-level tolerations if non-empty
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>podSecurityContext</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#podsecuritycontext-v1-core">
Kubernetes core/v1.PodSecurityContext
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodSecurityContext of the component</p>
</td>
</tr>
<tr>
<td>
<code>configUpostgresateStrategy</code></br>
<em>
<a href="#configupdatestrategy">
ConfigUpdateStrategy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>ConfigUpostgresateStrategy of the component. Override the cluster-level upostgresateStrategy if present
Optional: Defaults to cluster-level setting</p>
</td>
</tr>
<tr>
<td>
<code>env</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#envvar-v1-core">
[]Kubernetes core/v1.EnvVar
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>List of environment variables to set in the container, like v1.Container.Env.
Note that the following env names cannot be used and will be overridden by PostgreSQL Operator builtin envs
- NAMESPACE
- TZ
- SERVICE_NAME
- PEER_SERVICE_NAME
- HEADLESS_SERVICE_NAME
- SET_NAME
- HOSTNAME
- CLUSTER_NAME
- POD_NAME
- BINLOG_ENABLED
- SLOW_LOG_FILE</p>
</td>
</tr>
<tr>
<td>
<code>envFrom</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#envfromsource-v1-core">
[]Kubernetes core/v1.EnvFromSource
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Extend the use scenarios for env</p>
</td>
</tr>
<tr>
<td>
<code>initContainers</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Init containers of the components</p>
</td>
</tr>
<tr>
<td>
<code>additionalContainers</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#container-v1-core">
[]Kubernetes core/v1.Container
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional containers of the component.
If the container names in this field match with the ones generated by
PostgreSQL Operator, the container configurations will be merged into the
containers generated by PostgreSQL Operator via strategic merge patch.
If the container names in this field do not match with the ones
generated by PostgreSQL Operator, the container configurations will be
appended to the Pod container spec directly.</p>
</td>
</tr>
<tr>
<td>
<code>additionalVolumes</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#volume-v1-core">
[]Kubernetes core/v1.Volume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional volumes of component pod.</p>
</td>
</tr>
<tr>
<td>
<code>additionalVolumeMounts</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#volumemount-v1-core">
[]Kubernetes core/v1.VolumeMount
</a>
</em>
</td>
<td>
<p>Additional volume mounts of component pod.</p>
</td>
</tr>
<tr>
<td>
<code>dnsConfig</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#poddnsconfig-v1-core">
Kubernetes core/v1.PodDNSConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DNSConfig Specifies the DNS parameters of a pod.</p>
</td>
</tr>
<tr>
<td>
<code>dnsPolicy</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#dnspolicy-v1-core">
Kubernetes core/v1.DNSPolicy
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>DNSPolicy Specifies the DNSPolicy parameters of a pod.</p>
</td>
</tr>
<tr>
<td>
<code>terminationGracePeriodSeconds</code></br>
<em>
int64
</em>
</td>
<td>
<em>(Optional)</em>
<p>Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request.
Value must be non-negative integer. The value zero indicates delete immediately.
If this value is nil, the default grace period will be used instead.
The grace period is the duration in seconds after the processes running in the pod are sent
a termination signal and the time when the processes are forcibly halted with a kill signal.
Set this value longer than the expected cleanup time for your process.
Defaults to 30 seconds.</p>
</td>
</tr>
<tr>
<td>
<code>statefulSetUpdateStrategy</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#statefulsetupdatestrategytype-v1-apps">
Kubernetes apps/v1.StatefulSetUpdateStrategyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>StatefulSetUpostgresateStrategy indicates the StatefulSetUpostgresateStrategy that will be
employed to upostgresate Pods in the StatefulSet when a revision is made to
Template.</p>
</td>
</tr>
<tr>
<td>
<code>podManagementPolicy</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#podmanagementpolicytype-v1-apps">
Kubernetes apps/v1.PodManagementPolicyType
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>PodManagementPolicy of PostgreSQL cluster StatefulSets</p>
</td>
</tr>
</tbody>
</table>
<h3 id="configupdatestrategy">ConfigUpdateStrategy</h3>
<p>
(<em>Appears on:</em>
<a href="#componentspec">ComponentSpec</a>)
</p>
<p>
<p>ConfigUpdateStrategy represents the strategy to update configuration</p>
</p>
<h3 id="genericconfig">GenericConfig</h3>
<p>
(<em>Appears on:</em>
<a href="#postgresspec">PostgresSpec</a>)
</p>
<p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>MP</code></br>
<em>
map[string]string
</em>
</td>
<td>
<p>Export this field to make &ldquo;apiequality.Semantic.DeepEqual&rdquo; happy now.
User of GenericConfig should not directly access this field.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="postgresclustercondition">PostgresClusterCondition</h3>
<p>
(<em>Appears on:</em>
<a href="#postgresclusterstatus">PostgresClusterStatus</a>)
</p>
<p>
<p>TidbClusterCondition describes the state of a tidb cluster at a certain point.</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code></br>
<em>
<a href="#postgresclusterconditiontype">
PostgresClusterConditionType
</a>
</em>
</td>
<td>
<p>Type of the condition.</p>
</td>
</tr>
<tr>
<td>
<code>status</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#conditionstatus-v1-core">
Kubernetes core/v1.ConditionStatus
</a>
</em>
</td>
<td>
<p>Status of the condition, one of True, False, Unknown.</p>
</td>
</tr>
<tr>
<td>
<code>lastUpdateTime</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<p>The last time this condition was updated.</p>
</td>
</tr>
<tr>
<td>
<code>lastTransitionTime</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Last time the condition transitioned from one status to another.</p>
</td>
</tr>
<tr>
<td>
<code>reason</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The reason for the condition&rsquo;s last transition.</p>
</td>
</tr>
<tr>
<td>
<code>message</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>A human readable message indicating details about the transition.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="postgresclusterconditiontype">PostgresClusterConditionType</h3>
<p>
(<em>Appears on:</em>
<a href="#postgresclustercondition">PostgresClusterCondition</a>)
</p>
<p>
<p>PostgresClusterConditionType represents a tidb cluster condition value.</p>
</p>
<h3 id="postgresclusterspec">PostgresClusterSpec</h3>
<p>
(<em>Appears on:</em>
<a href="#postgrescluster">PostgresCluster</a>)
</p>
<p>
<p>PostgresClusterSpec describes the attributes that a user creates on a tidb cluster</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>serviceAccount</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify a Service Account</p>
</td>
</tr>
<tr>
<td>
<code>postgres</code></br>
<em>
<a href="#postgresspec">
PostgresSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Postgres cluster</p>
</td>
</tr>
</tbody>
</table>
<h3 id="postgresclusterstatus">PostgresClusterStatus</h3>
<p>
(<em>Appears on:</em>
<a href="#postgrescluster">PostgresCluster</a>)
</p>
<p>
<p>NodeHealthStatus is the status for a NodeHealth resource</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>availableReplicas</code></br>
<em>
int32
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>healthStatus</code></br>
<em>
<a href="#postgresclustercondition">
[]PostgresClusterCondition
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="postgresspec">PostgresSpec</h3>
<p>
(<em>Appears on:</em>
<a href="#postgresclusterspec">PostgresClusterSpec</a>)
</p>
<p>
<p>postgresSpec contains details of postgres members</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ComponentSpec</code></br>
<em>
<a href="#componentspec">
ComponentSpec
</a>
</em>
</td>
<td>
<p>
(Members of <code>ComponentSpec</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>ResourceRequirements</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#resourcerequirements-v1-core">
Kubernetes core/v1.ResourceRequirements
</a>
</em>
</td>
<td>
<p>
(Members of <code>ResourceRequirements</code> are embedded into this type.)
</p>
</td>
</tr>
<tr>
<td>
<code>serviceAccount</code></br>
<em>
string
</em>
</td>
<td>
<p>Specify a Service Account for postgres</p>
</td>
</tr>
<tr>
<td>
<code>replicas</code></br>
<em>
int32
</em>
</td>
<td>
<p>The desired ready replicas</p>
</td>
</tr>
<tr>
<td>
<code>baseImage</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Base image of the component, image tag is now allowed during validation</p>
</td>
</tr>
<tr>
<td>
<code>service</code></br>
<em>
<a href="#servicespec">
ServiceSpec
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Service defines a Kubernetes service of postgres cluster.
Optional: Defaults to <code>.spec.services</code> in favor of backward compatibility</p>
</td>
</tr>
<tr>
<td>
<code>maxFailoverCount</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>MaxFailoverCount limit the max replicas could be added in failover, 0 means no failover.
Optional: Defaults to 3</p>
</td>
</tr>
<tr>
<td>
<code>storageClassName</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The storageClassName of the persistent volume for postgres data storage.
Defaults to Kubernetes default storage class.</p>
</td>
</tr>
<tr>
<td>
<code>storageVolumes</code></br>
<em>
<a href="#storagevolume">
[]StorageVolume
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>StorageVolumes configure additional storage for postgres pods.</p>
</td>
</tr>
<tr>
<td>
<code>dataSubDir</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Subdirectory within the volume to store postgres Data. By default, the data
is stored in the root directory of volume which is mounted at
/var/lib/postgres.
Specifying this will change the data directory to a subdirectory, e.g.
/var/lib/postgres/data if you set the value to &ldquo;data&rdquo;.
It&rsquo;s dangerous to change this value for a running cluster as it will
upgrade your cluster to use a new storage directory.
Defaults to &ldquo;&rdquo; (volume&rsquo;s root).</p>
</td>
</tr>
<tr>
<td>
<code>config</code></br>
<em>
<a href="#genericconfig">
GenericConfig
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Config is the Configuration of postgres-servers</p>
</td>
</tr>
<tr>
<td>
<code>tlsClientSecretName</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>TLSClientSecretName is the name of secret which stores tidb server client certificate
which used by Dashboard.</p>
</td>
</tr>
<tr>
<td>
<code>enableDashboardInternalProxy</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>(Deprecated) EnableDashboardInternalProxy would directly set <code>internal-proxy</code> in the <code>postgresConfig</code>.
Note that this is deprecated, we should just set <code>dashboard.internal-proxy</code> in <code>postgres.config</code>.</p>
</td>
</tr>
<tr>
<td>
<code>mountClusterClientSecret</code></br>
<em>
bool
</em>
</td>
<td>
<em>(Optional)</em>
<p>MountClusterClientSecret indicates whether to mount <code>cluster-client-secret</code> to the Pod</p>
</td>
</tr>
<tr>
<td>
<code>startUpScriptVersion</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Start up script version</p>
</td>
</tr>
<tr>
<td>
<code>startTimeout</code></br>
<em>
int
</em>
</td>
<td>
<p>Timeout threshold when postgres get started</p>
</td>
</tr>
</tbody>
</table>
<h3 id="servicespec">ServiceSpec</h3>
<p>
(<em>Appears on:</em>
<a href="#postgresspec">PostgresSpec</a>)
</p>
<p>
<p>ServiceSpec specifies the service object in k8s</p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code></br>
<em>
<a href="https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.19/#servicetype-v1-core">
Kubernetes core/v1.ServiceType
</a>
</em>
</td>
<td>
<p>Type of the real kubernetes service</p>
</td>
</tr>
<tr>
<td>
<code>annotations</code></br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional annotations for the service</p>
</td>
</tr>
<tr>
<td>
<code>labels</code></br>
<em>
map[string]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>Additional labels for the service</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerIP</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>LoadBalancerIP is the loadBalancerIP of service
Optional: Defaults to omitted</p>
</td>
</tr>
<tr>
<td>
<code>clusterIP</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>ClusterIP is the clusterIP of service</p>
</td>
</tr>
<tr>
<td>
<code>portName</code></br>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>PortName is the name of service port</p>
</td>
</tr>
<tr>
<td>
<code>port</code></br>
<em>
int32
</em>
</td>
<td>
<em>(Optional)</em>
<p>The port that will be exposed by this service.</p>
<p>NOTE: only used for PostgreSQL</p>
</td>
</tr>
<tr>
<td>
<code>loadBalancerSourceRanges</code></br>
<em>
[]string
</em>
</td>
<td>
<em>(Optional)</em>
<p>LoadBalancerSourceRanges is the loadBalancerSourceRanges of service
If specified and supported by the platform, this will restrict traffic through the cloud-provider
load-balancer will be restricted to the specified client IPs. This field will be ignored if the
cloud-provider does not support the feature.&rdquo;
More info: <a href="https://kubernetes.io/docs/concepts/services-networking/service/#aws-nlb-support">https://kubernetes.io/docs/concepts/services-networking/service/#aws-nlb-support</a>
Optional: Defaults to omitted</p>
</td>
</tr>
</tbody>
</table>
<h3 id="storagevolume">StorageVolume</h3>
<p>
(<em>Appears on:</em>
<a href="#postgresspec">PostgresSpec</a>)
</p>
<p>
<p>StorageVolume configures additional PVC template for StatefulSets and volumeMount for pods that mount this PVC.
Note:
If <code>MountPath</code> is not set, volumeMount will not be generated. (You may not want to set this field when you inject volumeMount
in somewhere else such as Mutating Admission Webhook)
If <code>StorageClassName</code> is not set, default to the <code>spec.${component}.storageClassName</code></p>
</p>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageClassName</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>storageSize</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>mountPath</code></br>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
</em></p>
