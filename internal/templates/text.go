package templates

var Text = `
#####################################################################
Kubernetes As-Built Assessment - KASBA
Date: 		   {{ .CreatedAt }}
Format: 	   {{ .BOMFormat }}
KASBA Version: {{ .Version }}
#####################################################################

{{ if .Errors.HasErrors -}}
Errors:
  {{ range $index, $error := .Errors.Errors -}}
{{ $error }}
{{ end -}}
{{ end }}

{{ if .Errors.Fatal }}
Fatal errors, quiting.
{{ end }}

{{ if not .Errors.Fatal -}}

Cluster Name:         {{ (index .NodeInfo.Items 0).Metadata.Annotations.ClusterXK8SIoClusterName }}
Instance Type:        {{ (index .NodeInfo.Items 0).Metadata.Labels.NodeKubernetesIoInstanceType }}
K8s Version:          {{ (index .NodeInfo.Items 0).Status.NodeInfo.KubeletVersion }}

CNI:                  {{ .NetworkPlugin }}
Longhorn installed:   {{ .Longhorn }}

{{ range $index, $item := .NodeInfo.Items }}
Cluster Machine Name: {{ $item.Metadata.Annotations.ClusterXK8SIoMachine }}
Cluster Node Name:    {{ $item.Metadata.Name }}
Operating System:     {{ $item.Status.NodeInfo.OperatingSystem }}
OS Image:             {{ $item.Status.NodeInfo.OsImage }}
Node Arch:            {{ $item.Status.NodeInfo.Architecture }}
Kernel Version:       {{ $item.Status.NodeInfo.KernelVersion }}
System UUID:          {{ $item.Status.NodeInfo.SystemUUID }}
Container Runtime:    {{ $item.Status.NodeInfo.ContainerRuntimeVersion }}
Kube Version:         {{ $item.Status.NodeInfo.KubeletVersion }}
KubeProxy Version:    {{ $item.Status.NodeInfo.KubeProxyVersion }}
Node Args:            {{ $item.Metadata.Annotations.Rke2IoNodeArgs }}
Pod CIDR:             {{ $item.Spec.PodCIDR }}
Pod Limits:           {{ $item.Metadata.Annotations.ManagementCattleIoPodLimits }}
Pod Requests:         {{ $item.Metadata.Annotations.ManagementCattleIoPodRequests }}
--- Allocatable ---
CPU:                  {{ $item.Status.Allocatable.CPU }}
Memory:               {{ $item.Status.Allocatable.Memory }}
Ephemeral Storage:    {{ $item.Status.Allocatable.EphemeralStorage }}
Pods:                 {{ $item.Status.Allocatable.Pods }}

--- Messages {{ $item.Metadata.Name }}---
{{ range $index, $condition := $item.Status.Conditions -}}
	Condition Type: {{ $condition.Type }}
	  Last Heartbeat Time:  {{ $condition.LastHeartbeatTime }}
	  Last Transition Time: {{ $condition.LastTransitionTime }}
	  Message {{ $condition.Message }}
	  Reason: {{ $condition.Reason }}
	  Status: {{ $condition.Status }}
{{ end -}}
{{ end }}

--- Workload ---
{{ range $index, $namespace := .WorkloadInfo.Namespaces -}}
Namespace: {{ $namespace.Namespace }}
  {{- range $index, $apptype := $namespace.WorkloadTypes }}
  {{ $apptype.WorkloadType }}:
    {{- range $index, $name := $apptype.Workloads }}
	{{ $name }}
    {{- end -}}
  {{ end }}

{{ end }}

{{ end -}}

`
