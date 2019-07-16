// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetwork":                 schema_pkg_apis_sriovnetwork_v1_SriovNetwork(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicy":       schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicy(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicySpec":   schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicySpec(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicyStatus": schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicyStatus(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeState":        schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeState(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateSpec":    schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateSpec(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateStatus":  schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateStatus(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkSpec":             schema_pkg_apis_sriovnetwork_v1_SriovNetworkSpec(ref),
		"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkStatus":           schema_pkg_apis_sriovnetwork_v1_SriovNetworkStatus(ref),
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetwork(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetwork is the Schema for the sriovnetworks API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkSpec", "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicy(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodePolicy is the Schema for the sriovnetworknodepolicies API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicyStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicySpec", "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodePolicyStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodePolicySpec defines the desired state of SriovNetworkNodePolicy",
				Properties: map[string]spec.Schema{
					"resourceName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"nodeSelector": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"priority": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"mtu": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"numVfs": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"nicSelector": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNicSelector"),
						},
					},
					"deviceType": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"isRdma": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"resourceName", "nodeSelector", "numVfs", "nicSelector"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNicSelector"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodePolicyStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodePolicyStatus defines the observed state of SriovNetworkNodePolicy",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeState(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodeState is the Schema for the sriovnetworknodestates API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateSpec", "github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.SriovNetworkNodeStateStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodeStateSpec defines the desired state of SriovNetworkNodeState",
				Properties: map[string]spec.Schema{
					"dpConfigVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"interfaces": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.Interface"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.Interface"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkNodeStateStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkNodeStateStatus defines the observed state of SriovNetworkNodeState",
				Properties: map[string]spec.Schema{
					"interfaces": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.InterfaceExt"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/sriov-network-operator/pkg/apis/sriovnetwork/v1.InterfaceExt"},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkSpec defines the desired state of SriovNetwork",
				Properties: map[string]spec.Schema{
					"networkNamespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"resourceName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"ipam": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"vlan": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
					"spoofChk": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"trust": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"resourceName"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_sriovnetwork_v1_SriovNetworkStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SriovNetworkStatus defines the observed state of SriovNetwork",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
