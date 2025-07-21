{{/*
Expand the name of the chart.
*/}}
{{- define "jelastic-golang-hello.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "jelastic-golang-hello.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "jelastic-golang-hello.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "jelastic-golang-hello.labels" -}}
helm.sh/chart: {{ include "jelastic-golang-hello.chart" . }}
{{ include "jelastic-golang-hello.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "jelastic-golang-hello.selectorLabels" -}}
app.kubernetes.io/name: {{ include "jelastic-golang-hello.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "jelastic-golang-hello.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "jelastic-golang-hello.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Create the image name
*/}}
{{- define "jelastic-golang-hello.image" -}}
{{- $tag := .Values.image.tag | default .Chart.AppVersion }}
{{- printf "%s:%s" .Values.image.repository $tag }}
{{- end }}

{{/*
Create environment-specific values
*/}}
{{- define "jelastic-golang-hello.environmentConfig" -}}
{{- $env := .Values.global.environment | default "production" }}
{{- if hasKey .Values.environments $env }}
{{- $envConfig := index .Values.environments $env }}
{{- toYaml $envConfig }}
{{- end }}
{{- end }}