// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStore":       schema_pkg_apis_noobaa_v1alpha1_BackingStore(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreSpec":   schema_pkg_apis_noobaa_v1alpha1_BackingStoreSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreStatus": schema_pkg_apis_noobaa_v1alpha1_BackingStoreStatus(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClass":        schema_pkg_apis_noobaa_v1alpha1_BucketClass(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassSpec":    schema_pkg_apis_noobaa_v1alpha1_BucketClassSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassStatus":  schema_pkg_apis_noobaa_v1alpha1_BucketClassStatus(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaa":             schema_pkg_apis_noobaa_v1alpha1_NooBaa(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec":         schema_pkg_apis_noobaa_v1alpha1_NooBaaSpec(ref),
		"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus":       schema_pkg_apis_noobaa_v1alpha1_NooBaaStatus(ref),
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BackingStore(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BackingStore is the Schema for the backingstores API",
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
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BackingStoreStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BackingStoreSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BackingStoreSpec defines the desired state of BackingStore",
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"bucketName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"secret": {
						SchemaProps: spec.SchemaProps{
							Description: "Secret refers to a secret that provides the credentials",
							Ref:         ref("k8s.io/api/core/v1.SecretReference"),
						},
					},
					"s3Options": {
						SchemaProps: spec.SchemaProps{
							Description: "S3Options specifies client options for the backing store",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.S3Options"),
						},
					},
				},
				Required: []string{"type", "bucketName", "secret"},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.S3Options", "k8s.io/api/core/v1.SecretReference"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BackingStoreStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BackingStoreStatus defines the observed state of BackingStore",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the System is in its lifecycle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a list of conditions related to operator reconciliation",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects that are \"interesting\" or related to this operator.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
				},
				Required: []string{"phase"},
			},
		},
		Dependencies: []string{
			"github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BucketClass(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BucketClass is the Schema for the bucketclasses API",
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
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.BucketClassStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BucketClassSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BucketClassSpec defines the desired state of BucketClass",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_BucketClassStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "BucketClassStatus defines the observed state of BucketClass",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaa(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaa is the Schema for the NooBaas API",
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
							Description: "Standard object metadata.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Description: "Specification of the desired behavior of the noobaa system.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Most recently observed status of the noobaa system.",
							Ref:         ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaSpec", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.NooBaaStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaaSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaaSpec defines the desired state of System",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image (optional) overrides the default image for the server container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mongoImage": {
						SchemaProps: spec.SchemaProps{
							Description: "MongoImage (optional) overrides the default image for the mongodb container",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"coreResources": {
						SchemaProps: spec.SchemaProps{
							Description: "CoreResources (optional) overrides the default resource requirements for the server container",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"mongoResources": {
						SchemaProps: spec.SchemaProps{
							Description: "MongoResources (optional) overrides the default resource requirements for the mongodb container",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"imagePullSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "ImagePullSecret (optional) sets a pull secret for the system image",
							Ref:         ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
					"storageClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageClassName (optional) overrides the default StorageClass for the PVC that the operator creates, this affects where the system stores its database which contains system config, buckets, objects meta-data and mapping file parts to storage locations.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dbVolumeResources": {
						SchemaProps: spec.SchemaProps{
							Description: "DBVolumeResources (optional) overrides the default PVC resource requirements for the database volume (mongo). Updates to this value are supported only for increasing the size, and only if the storage class specifies `allowVolumeExpansion: true`.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.LocalObjectReference", "k8s.io/api/core/v1.ResourceRequirements"},
	}
}

func schema_pkg_apis_noobaa_v1alpha1_NooBaaStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NooBaaStatus defines the observed state of System",
				Properties: map[string]spec.Schema{
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "ObservedGeneration is the most recent generation observed for this noobaa system. It corresponds to the CR generation, which is updated on mutation by the API Server.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the System is in its lifecycle",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-patch-merge-key": "type",
								"x-kubernetes-patch-strategy":  "merge",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions is a list of conditions related to operator reconciliation",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/openshift/custom-resource-status/conditions/v1.Condition"),
									},
								},
							},
						},
					},
					"relatedObjects": {
						SchemaProps: spec.SchemaProps{
							Description: "RelatedObjects is a list of objects that are \"interesting\" or related to this operator.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"actualImage": {
						SchemaProps: spec.SchemaProps{
							Description: "ActualImage is set to report which image the operator is using",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accounts": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AccountsStatus"),
						},
					},
					"services": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.ServicesStatus"),
						},
					},
					"readme": {
						SchemaProps: spec.SchemaProps{
							Description: "Readme is a user readable string with explanations on the system",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"observedGeneration", "phase", "actualImage", "accounts", "services", "readme"},
			},
		},
		Dependencies: []string{
			"github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.AccountsStatus", "github.com/noobaa/noobaa-operator/pkg/apis/noobaa/v1alpha1.ServicesStatus", "github.com/openshift/custom-resource-status/conditions/v1.Condition", "k8s.io/api/core/v1.ObjectReference"},
	}
}
