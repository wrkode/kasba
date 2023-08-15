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
Monitoring Installed: {{ .Monitoring }}
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

--- Storage ---
  Storage Classes:
  {{- range $index, $sc := .StorageClass }}
    - Name: {{ $sc.Name }}
      Provisioner: {{ $sc.Provisioner }}
      Parameters:
      {{- range $key, $value := $sc.Parameters }}
        {{ $key }}: {{ $value }}
      {{- end }}
  {{- end }}

  Persistent Volumes:
{{- range $index, $pv := .PersistentVolumes }}
  - Name: {{ $pv.Name }}
    Type: {{ $pv.Type }}
    Capacity: {{ index $pv.Size }}
    Access Modes:
    {{- range $index, $mode := $pv.AccessModes }}
      - {{ $mode }}
    {{- end }}
    Reclamation Policy: {{ $pv.ReclamationPolicy }}
{{- end }}

  Persistent Volume Claims:
  {{- range $index, $pvc := .PersistentVolumeClaims }}
    - Namespace: {{ $pvc.Namespace }}
      Name: {{ $pvc.Name}}
      Status: {{ $pvc.Status  }}
      Volume: {{ $pvc.Volume  }}
      Capacity: {{ $pvc.Capacity }}
      AccessModes: {{ $pvc.AccessModes }}
      StorageClass: {{ $pvc.StorageClass }}
      Age: {{ $pvc.Age }}
  {{- end }}

  Config Maps:
{{- $currentNamespace := "" -}}
{{- range $index, $cm := .ConfigMaps -}}
{{- if ne $cm.Namespace $currentNamespace }}
Namespace: {{ $cm.Namespace }}
{{- $currentNamespace = $cm.Namespace -}}
{{- end }}
  Name:      {{ $cm.Name }}
  Age:       {{ $cm.Age }}
{{- end }}

--- Service Discovery ---
  Services:
{{- $currentNamespace := "" -}}
{{- range $index, $serviceItem := .Services -}}
{{- if ne $serviceItem.Namespace $currentNamespace }}
Namespace: {{ $serviceItem.Namespace }}
{{- $currentNamespace = $serviceItem.Namespace -}}
{{- end }}
  Name:      {{ $serviceItem.Name }}
    Type:      {{ $serviceItem.Type }}
    ClusterIP: {{ $serviceItem.ClusterIP }}
    ExternalIP: {{ if $serviceItem.ExternalIP }}{{ $serviceItem.ExternalIP }}{{ else }}<none>{{ end }}
    Ports:     
    {{- range $pIndex, $port := $serviceItem.Ports }}
      {{ $port.Name }}: {{ $port.Port }}/{{ $port.Protocol }} {{ if $port.NodePort }}(NodePort: {{ $port.NodePort }}){{ end }}
    {{- end }}
    Age:       {{ $serviceItem.Age }}d
{{- end }}

  Ingresses:
{{- $currentNamespace := "" -}}
{{- range $index, $ingressItem := .Ingresses -}}
{{- if ne $ingressItem.Namespace $currentNamespace }}
Namespace: {{ $ingressItem.Namespace }}
{{ $currentNamespace = $ingressItem.Namespace }}{{/* This is the main adjustment to the variable assignment */}}
{{- end }}
    Name: {{ $ingressItem.Name }}
      Hosts: {{ $ingressItem.Hosts }}
      DefaultBackend: {{ $ingressItem.DefaultBackend }}
      Age: {{ $ingressItem.Age }}d
{{- end }}
{{- end }}
`
