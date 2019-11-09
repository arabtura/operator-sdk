// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.Memcached":         schema_pkg_apis_cache_v1alpha1_Memcached(ref),
		"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedRS":       schema_pkg_apis_cache_v1alpha1_MemcachedRS(ref),
		"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedRSSpec":   schema_pkg_apis_cache_v1alpha1_MemcachedRSSpec(ref),
		"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedRSStatus": schema_pkg_apis_cache_v1alpha1_MemcachedRSStatus(ref),
		"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedSpec":     schema_pkg_apis_cache_v1alpha1_MemcachedSpec(ref),
		"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedStatus":   schema_pkg_apis_cache_v1alpha1_MemcachedStatus(ref),
	}
}

func schema_pkg_apis_cache_v1alpha1_Memcached(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Memcached is the Schema for the memcacheds API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
							Ref: ref("github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedSpec", "github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cache_v1alpha1_MemcachedRS(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MemcachedRS is the Schema for the memcachedrs API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
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
							Ref: ref("github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedRSSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedRSStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedRSSpec", "github.com/operator-framework/operator-sdk/test/test-framework/pkg/apis/cache/v1alpha1.MemcachedRSStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_cache_v1alpha1_MemcachedRSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MemcachedRSSpec defines the desired state of MemcachedRS",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"numNodes": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int32",
						},
					},
				},
				Required: []string{"numNodes"},
			},
		},
	}
}

func schema_pkg_apis_cache_v1alpha1_MemcachedRSStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MemcachedRSStatus defines the observed state of MemcachedRS",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodeList": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"test": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
				},
				Required: []string{"nodeList", "test"},
			},
		},
	}
}

func schema_pkg_apis_cache_v1alpha1_MemcachedSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MemcachedSpec defines the desired state of Memcached",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"size": {
						SchemaProps: spec.SchemaProps{
							Description: "Size is the size of the memcached deployment",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
				},
				Required: []string{"size"},
			},
		},
	}
}

func schema_pkg_apis_cache_v1alpha1_MemcachedStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MemcachedStatus defines the observed state of Memcached",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nodes": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Nodes are the names of the memcached pods",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"nodes"},
			},
		},
	}
}