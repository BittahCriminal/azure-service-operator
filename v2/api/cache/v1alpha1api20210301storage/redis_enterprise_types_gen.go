// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210301storage

import (
	"fmt"
	v20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20201201storage"
	v20210301s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1alpha1api20210301.RedisEnterprise
// Deprecated version of RedisEnterprise. Use v1beta20210301.RedisEnterprise instead
type RedisEnterprise struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisEnterprise_Spec `json:"spec,omitempty"`
	Status            Cluster_STATUS       `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisEnterprise{}

// GetConditions returns the conditions of the resource
func (enterprise *RedisEnterprise) GetConditions() conditions.Conditions {
	return enterprise.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (enterprise *RedisEnterprise) SetConditions(conditions conditions.Conditions) {
	enterprise.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisEnterprise{}

// ConvertFrom populates our RedisEnterprise from the provided hub RedisEnterprise
func (enterprise *RedisEnterprise) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210301s.RedisEnterprise)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20210301storage/RedisEnterprise but received %T instead", hub)
	}

	return enterprise.AssignProperties_From_RedisEnterprise(source)
}

// ConvertTo populates the provided hub RedisEnterprise from our RedisEnterprise
func (enterprise *RedisEnterprise) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210301s.RedisEnterprise)
	if !ok {
		return fmt.Errorf("expected cache/v1beta20210301storage/RedisEnterprise but received %T instead", hub)
	}

	return enterprise.AssignProperties_To_RedisEnterprise(destination)
}

var _ genruntime.KubernetesResource = &RedisEnterprise{}

// AzureName returns the Azure name of the resource
func (enterprise *RedisEnterprise) AzureName() string {
	return enterprise.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-03-01"
func (enterprise RedisEnterprise) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (enterprise *RedisEnterprise) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (enterprise *RedisEnterprise) GetSpec() genruntime.ConvertibleSpec {
	return &enterprise.Spec
}

// GetStatus returns the status of this resource
func (enterprise *RedisEnterprise) GetStatus() genruntime.ConvertibleStatus {
	return &enterprise.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redisEnterprise"
func (enterprise *RedisEnterprise) GetType() string {
	return "Microsoft.Cache/redisEnterprise"
}

// NewEmptyStatus returns a new empty (blank) status
func (enterprise *RedisEnterprise) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Cluster_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (enterprise *RedisEnterprise) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(enterprise.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  enterprise.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (enterprise *RedisEnterprise) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Cluster_STATUS); ok {
		enterprise.Status = *st
		return nil
	}

	// Convert status to required version
	var st Cluster_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	enterprise.Status = st
	return nil
}

// AssignProperties_From_RedisEnterprise populates our RedisEnterprise from the provided source RedisEnterprise
func (enterprise *RedisEnterprise) AssignProperties_From_RedisEnterprise(source *v20210301s.RedisEnterprise) error {

	// ObjectMeta
	enterprise.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisEnterprise_Spec
	err := spec.AssignProperties_From_RedisEnterprise_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RedisEnterprise_Spec() to populate field Spec")
	}
	enterprise.Spec = spec

	// Status
	var status Cluster_STATUS
	err = status.AssignProperties_From_Cluster_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Cluster_STATUS() to populate field Status")
	}
	enterprise.Status = status

	// No error
	return nil
}

// AssignProperties_To_RedisEnterprise populates the provided destination RedisEnterprise from our RedisEnterprise
func (enterprise *RedisEnterprise) AssignProperties_To_RedisEnterprise(destination *v20210301s.RedisEnterprise) error {

	// ObjectMeta
	destination.ObjectMeta = *enterprise.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210301s.RedisEnterprise_Spec
	err := enterprise.Spec.AssignProperties_To_RedisEnterprise_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RedisEnterprise_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210301s.Cluster_STATUS
	err = enterprise.Status.AssignProperties_To_Cluster_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Cluster_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (enterprise *RedisEnterprise) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: enterprise.Spec.OriginalVersion,
		Kind:    "RedisEnterprise",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210301.RedisEnterprise
// Deprecated version of RedisEnterprise. Use v1beta20210301.RedisEnterprise instead
type RedisEnterpriseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisEnterprise `json:"items"`
}

// Storage version of v1alpha1api20210301.APIVersion
// Deprecated version of APIVersion. Use v1beta20210301.APIVersion instead
// +kubebuilder:validation:Enum={"2021-03-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-03-01")

// Storage version of v1alpha1api20210301.Cluster_STATUS
// Deprecated version of Cluster_STATUS. Use v1beta20210301.Cluster_STATUS instead
type Cluster_STATUS struct {
	Conditions                 []conditions.Condition                                 `json:"conditions,omitempty"`
	HostName                   *string                                                `json:"hostName,omitempty"`
	Id                         *string                                                `json:"id,omitempty"`
	Location                   *string                                                `json:"location,omitempty"`
	MinimumTlsVersion          *string                                                `json:"minimumTlsVersion,omitempty"`
	Name                       *string                                                `json:"name,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_STATUS_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                `json:"provisioningState,omitempty"`
	RedisVersion               *string                                                `json:"redisVersion,omitempty"`
	ResourceState              *string                                                `json:"resourceState,omitempty"`
	Sku                        *Sku_STATUS                                            `json:"sku,omitempty"`
	Tags                       map[string]string                                      `json:"tags,omitempty"`
	Type                       *string                                                `json:"type,omitempty"`
	Zones                      []string                                               `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Cluster_STATUS{}

// ConvertStatusFrom populates our Cluster_STATUS from the provided source
func (cluster *Cluster_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210301s.Cluster_STATUS)
	if ok {
		// Populate our instance from source
		return cluster.AssignProperties_From_Cluster_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210301s.Cluster_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = cluster.AssignProperties_From_Cluster_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Cluster_STATUS
func (cluster *Cluster_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210301s.Cluster_STATUS)
	if ok {
		// Populate destination from our instance
		return cluster.AssignProperties_To_Cluster_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210301s.Cluster_STATUS{}
	err := cluster.AssignProperties_To_Cluster_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_Cluster_STATUS populates our Cluster_STATUS from the provided source Cluster_STATUS
func (cluster *Cluster_STATUS) AssignProperties_From_Cluster_STATUS(source *v20210301s.Cluster_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	cluster.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// HostName
	cluster.HostName = genruntime.ClonePointerToString(source.HostName)

	// Id
	cluster.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	cluster.Location = genruntime.ClonePointerToString(source.Location)

	// MinimumTlsVersion
	cluster.MinimumTlsVersion = genruntime.ClonePointerToString(source.MinimumTlsVersion)

	// Name
	cluster.Name = genruntime.ClonePointerToString(source.Name)

	// PrivateEndpointConnections
	if source.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]PrivateEndpointConnection_STATUS_SubResourceEmbedded, len(source.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range source.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnectionSTATUSSubResourceEmbeddedStash v20201201s.PrivateEndpointConnection_STATUS_SubResourceEmbedded
			err := privateEndpointConnectionSTATUSSubResourceEmbeddedStash.AssignProperties_From_PrivateEndpointConnection_STATUS_SubResourceEmbedded(&privateEndpointConnectionItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_PrivateEndpointConnection_STATUS_SubResourceEmbedded() to populate field PrivateEndpointConnection_STATUS_SubResourceEmbeddedStash from PrivateEndpointConnections")
			}
			var privateEndpointConnection PrivateEndpointConnection_STATUS_SubResourceEmbedded
			err = privateEndpointConnection.AssignProperties_From_PrivateEndpointConnection_STATUS_SubResourceEmbedded(&privateEndpointConnectionSTATUSSubResourceEmbeddedStash)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_From_PrivateEndpointConnection_STATUS_SubResourceEmbedded() to populate field PrivateEndpointConnections from PrivateEndpointConnection_STATUS_SubResourceEmbeddedStash")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		cluster.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		cluster.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	cluster.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// RedisVersion
	cluster.RedisVersion = genruntime.ClonePointerToString(source.RedisVersion)

	// ResourceState
	cluster.ResourceState = genruntime.ClonePointerToString(source.ResourceState)

	// Sku
	if source.Sku != nil {
		var skuSTATUSStash v20201201s.Sku_STATUS
		err := skuSTATUSStash.AssignProperties_From_Sku_STATUS(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku_STATUS() to populate field Sku_STATUSStash from Sku")
		}
		var sku Sku_STATUS
		err = sku.AssignProperties_From_Sku_STATUS(&skuSTATUSStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku_STATUS() to populate field Sku from Sku_STATUSStash")
		}
		cluster.Sku = &sku
	} else {
		cluster.Sku = nil
	}

	// Tags
	cluster.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	cluster.Type = genruntime.ClonePointerToString(source.Type)

	// Zones
	cluster.Zones = genruntime.CloneSliceOfString(source.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		cluster.PropertyBag = propertyBag
	} else {
		cluster.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Cluster_STATUS populates the provided destination Cluster_STATUS from our Cluster_STATUS
func (cluster *Cluster_STATUS) AssignProperties_To_Cluster_STATUS(destination *v20210301s.Cluster_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(cluster.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(cluster.Conditions)

	// HostName
	destination.HostName = genruntime.ClonePointerToString(cluster.HostName)

	// Id
	destination.Id = genruntime.ClonePointerToString(cluster.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(cluster.Location)

	// MinimumTlsVersion
	destination.MinimumTlsVersion = genruntime.ClonePointerToString(cluster.MinimumTlsVersion)

	// Name
	destination.Name = genruntime.ClonePointerToString(cluster.Name)

	// PrivateEndpointConnections
	if cluster.PrivateEndpointConnections != nil {
		privateEndpointConnectionList := make([]v20210301s.PrivateEndpointConnection_STATUS_SubResourceEmbedded, len(cluster.PrivateEndpointConnections))
		for privateEndpointConnectionIndex, privateEndpointConnectionItem := range cluster.PrivateEndpointConnections {
			// Shadow the loop variable to avoid aliasing
			privateEndpointConnectionItem := privateEndpointConnectionItem
			var privateEndpointConnectionSTATUSSubResourceEmbeddedStash v20201201s.PrivateEndpointConnection_STATUS_SubResourceEmbedded
			err := privateEndpointConnectionItem.AssignProperties_To_PrivateEndpointConnection_STATUS_SubResourceEmbedded(&privateEndpointConnectionSTATUSSubResourceEmbeddedStash)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_PrivateEndpointConnection_STATUS_SubResourceEmbedded() to populate field PrivateEndpointConnection_STATUS_SubResourceEmbeddedStash from PrivateEndpointConnections")
			}
			var privateEndpointConnection v20210301s.PrivateEndpointConnection_STATUS_SubResourceEmbedded
			err = privateEndpointConnectionSTATUSSubResourceEmbeddedStash.AssignProperties_To_PrivateEndpointConnection_STATUS_SubResourceEmbedded(&privateEndpointConnection)
			if err != nil {
				return errors.Wrap(err, "calling AssignProperties_To_PrivateEndpointConnection_STATUS_SubResourceEmbedded() to populate field PrivateEndpointConnections from PrivateEndpointConnection_STATUS_SubResourceEmbeddedStash")
			}
			privateEndpointConnectionList[privateEndpointConnectionIndex] = privateEndpointConnection
		}
		destination.PrivateEndpointConnections = privateEndpointConnectionList
	} else {
		destination.PrivateEndpointConnections = nil
	}

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(cluster.ProvisioningState)

	// RedisVersion
	destination.RedisVersion = genruntime.ClonePointerToString(cluster.RedisVersion)

	// ResourceState
	destination.ResourceState = genruntime.ClonePointerToString(cluster.ResourceState)

	// Sku
	if cluster.Sku != nil {
		var skuSTATUSStash v20201201s.Sku_STATUS
		err := cluster.Sku.AssignProperties_To_Sku_STATUS(&skuSTATUSStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku_STATUS() to populate field Sku_STATUSStash from Sku")
		}
		var sku v20210301s.Sku_STATUS
		err = skuSTATUSStash.AssignProperties_To_Sku_STATUS(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku_STATUS() to populate field Sku from Sku_STATUSStash")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(cluster.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(cluster.Type)

	// Zones
	destination.Zones = genruntime.CloneSliceOfString(cluster.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.RedisEnterprise_Spec
type RedisEnterprise_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName         string  `json:"azureName,omitempty"`
	Location          *string `json:"location,omitempty"`
	MinimumTlsVersion *string `json:"minimumTlsVersion,omitempty"`
	OriginalVersion   string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner       *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Sku         *Sku                               `json:"sku,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
	Zones       []string                           `json:"zones,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisEnterprise_Spec{}

// ConvertSpecFrom populates our RedisEnterprise_Spec from the provided source
func (enterprise *RedisEnterprise_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210301s.RedisEnterprise_Spec)
	if ok {
		// Populate our instance from source
		return enterprise.AssignProperties_From_RedisEnterprise_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210301s.RedisEnterprise_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = enterprise.AssignProperties_From_RedisEnterprise_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisEnterprise_Spec
func (enterprise *RedisEnterprise_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210301s.RedisEnterprise_Spec)
	if ok {
		// Populate destination from our instance
		return enterprise.AssignProperties_To_RedisEnterprise_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210301s.RedisEnterprise_Spec{}
	err := enterprise.AssignProperties_To_RedisEnterprise_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_RedisEnterprise_Spec populates our RedisEnterprise_Spec from the provided source RedisEnterprise_Spec
func (enterprise *RedisEnterprise_Spec) AssignProperties_From_RedisEnterprise_Spec(source *v20210301s.RedisEnterprise_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	enterprise.AzureName = source.AzureName

	// Location
	enterprise.Location = genruntime.ClonePointerToString(source.Location)

	// MinimumTlsVersion
	enterprise.MinimumTlsVersion = genruntime.ClonePointerToString(source.MinimumTlsVersion)

	// OriginalVersion
	enterprise.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		enterprise.Owner = &owner
	} else {
		enterprise.Owner = nil
	}

	// Sku
	if source.Sku != nil {
		var skuStash v20201201s.Sku
		err := skuStash.AssignProperties_From_Sku(source.Sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku() to populate field SkuStash from Sku")
		}
		var sku Sku
		err = sku.AssignProperties_From_Sku(&skuStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_Sku() to populate field Sku from SkuStash")
		}
		enterprise.Sku = &sku
	} else {
		enterprise.Sku = nil
	}

	// Tags
	enterprise.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Zones
	enterprise.Zones = genruntime.CloneSliceOfString(source.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		enterprise.PropertyBag = propertyBag
	} else {
		enterprise.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_RedisEnterprise_Spec populates the provided destination RedisEnterprise_Spec from our RedisEnterprise_Spec
func (enterprise *RedisEnterprise_Spec) AssignProperties_To_RedisEnterprise_Spec(destination *v20210301s.RedisEnterprise_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(enterprise.PropertyBag)

	// AzureName
	destination.AzureName = enterprise.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(enterprise.Location)

	// MinimumTlsVersion
	destination.MinimumTlsVersion = genruntime.ClonePointerToString(enterprise.MinimumTlsVersion)

	// OriginalVersion
	destination.OriginalVersion = enterprise.OriginalVersion

	// Owner
	if enterprise.Owner != nil {
		owner := enterprise.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Sku
	if enterprise.Sku != nil {
		var skuStash v20201201s.Sku
		err := enterprise.Sku.AssignProperties_To_Sku(&skuStash)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku() to populate field SkuStash from Sku")
		}
		var sku v20210301s.Sku
		err = skuStash.AssignProperties_To_Sku(&sku)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_Sku() to populate field Sku from SkuStash")
		}
		destination.Sku = &sku
	} else {
		destination.Sku = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(enterprise.Tags)

	// Zones
	destination.Zones = genruntime.CloneSliceOfString(enterprise.Zones)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.PrivateEndpointConnection_STATUS_SubResourceEmbedded
// Deprecated version of PrivateEndpointConnection_STATUS_SubResourceEmbedded. Use v1beta20210301.PrivateEndpointConnection_STATUS_SubResourceEmbedded instead
type PrivateEndpointConnection_STATUS_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_PrivateEndpointConnection_STATUS_SubResourceEmbedded populates our PrivateEndpointConnection_STATUS_SubResourceEmbedded from the provided source PrivateEndpointConnection_STATUS_SubResourceEmbedded
func (embedded *PrivateEndpointConnection_STATUS_SubResourceEmbedded) AssignProperties_From_PrivateEndpointConnection_STATUS_SubResourceEmbedded(source *v20201201s.PrivateEndpointConnection_STATUS_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Id
	embedded.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		embedded.PropertyBag = propertyBag
	} else {
		embedded.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_PrivateEndpointConnection_STATUS_SubResourceEmbedded populates the provided destination PrivateEndpointConnection_STATUS_SubResourceEmbedded from our PrivateEndpointConnection_STATUS_SubResourceEmbedded
func (embedded *PrivateEndpointConnection_STATUS_SubResourceEmbedded) AssignProperties_To_PrivateEndpointConnection_STATUS_SubResourceEmbedded(destination *v20201201s.PrivateEndpointConnection_STATUS_SubResourceEmbedded) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(embedded.PropertyBag)

	// Id
	destination.Id = genruntime.ClonePointerToString(embedded.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.Sku
// Deprecated version of Sku. Use v1beta20210301.Sku instead
type Sku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_Sku populates our Sku from the provided source Sku
func (sku *Sku) AssignProperties_From_Sku(source *v20201201s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Capacity
	sku.Capacity = genruntime.ClonePointerToInt(source.Capacity)

	// Family
	if source.Family != nil {
		propertyBag.Add("Family", *source.Family)
	} else {
		propertyBag.Remove("Family")
	}

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Sku populates the provided destination Sku from our Sku
func (sku *Sku) AssignProperties_To_Sku(destination *v20201201s.Sku) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Capacity
	destination.Capacity = genruntime.ClonePointerToInt(sku.Capacity)

	// Family
	if propertyBag.Contains("Family") {
		var family string
		err := propertyBag.Pull("Family", &family)
		if err != nil {
			return errors.Wrap(err, "pulling 'Family' from propertyBag")
		}

		destination.Family = &family
	} else {
		destination.Family = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210301.Sku_STATUS
// Deprecated version of Sku_STATUS. Use v1beta20210301.Sku_STATUS instead
type Sku_STATUS struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_Sku_STATUS populates our Sku_STATUS from the provided source Sku_STATUS
func (sku *Sku_STATUS) AssignProperties_From_Sku_STATUS(source *v20201201s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Capacity
	sku.Capacity = genruntime.ClonePointerToInt(source.Capacity)

	// Family
	if source.Family != nil {
		propertyBag.Add("Family", *source.Family)
	} else {
		propertyBag.Remove("Family")
	}

	// Name
	sku.Name = genruntime.ClonePointerToString(source.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		sku.PropertyBag = propertyBag
	} else {
		sku.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Sku_STATUS populates the provided destination Sku_STATUS from our Sku_STATUS
func (sku *Sku_STATUS) AssignProperties_To_Sku_STATUS(destination *v20201201s.Sku_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(sku.PropertyBag)

	// Capacity
	destination.Capacity = genruntime.ClonePointerToInt(sku.Capacity)

	// Family
	if propertyBag.Contains("Family") {
		var family string
		err := propertyBag.Pull("Family", &family)
		if err != nil {
			return errors.Wrap(err, "pulling 'Family' from propertyBag")
		}

		destination.Family = &family
	} else {
		destination.Family = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(sku.Name)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&RedisEnterprise{}, &RedisEnterpriseList{})
}
