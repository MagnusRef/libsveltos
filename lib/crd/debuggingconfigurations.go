// Generated by *go generate* - DO NOT EDIT
/*
Copyright 2022-23. projectsveltos.io. All rights reserved.

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
package crd

var DebuggingConfigurationFile = "../../config/crd/bases/lib.projectsveltos.io_debuggingconfigurations.yaml"
var DebuggingConfigurationCRD = []byte(`---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.11.3
  creationTimestamp: null
  name: debuggingconfigurations.lib.projectsveltos.io
spec:
  group: lib.projectsveltos.io
  names:
    kind: DebuggingConfiguration
    listKind: DebuggingConfigurationList
    plural: debuggingconfigurations
    singular: debuggingconfiguration
  scope: Cluster
  versions:
  - name: v1alpha1
    schema:
      openAPIV3Schema:
        description: DebuggingConfiguration is the Schema for the debuggingconfigurations
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: DebuggingConfigurationSpec defines the desired state of DebuggingConfiguration
            properties:
              configuration:
                description: Configuration contains debugging configuration as granular
                  as per component.
                items:
                  description: ComponentConfiguration is the debugging configuration
                    to be applied to a Sveltos component.
                  properties:
                    component:
                      description: Component indicates which Sveltos component the
                        configuration applies to.
                      enum:
                      - AddonManager
                      - AddonComplianceManager
                      - Classifier
                      - ClassifierAgent
                      - SveltosClusterManager
                      - DriftDetectionManager
                      - AccessManager
                      - HealthCheckManager
                      - EventManager
                      type: string
                    logLevel:
                      description: 'LogLevel is the log severity above which logs
                        are sent to the stdout. [Default: Info]'
                      enum:
                      - LogLevelNotSet
                      - LogLevelInfo
                      - LogLevelDebug
                      - LogLevelVerbose
                      type: string
                  required:
                  - component
                  type: object
                type: array
                x-kubernetes-list-type: atomic
            type: object
        type: object
    served: true
    storage: true
`)
