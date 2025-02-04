// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkInterface_Spec_ARM struct {
	// ExtendedLocation: The extended location of the network interface.
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the network interface.
	Properties *NetworkInterface_Properties_Spec_ARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkInterface_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (networkInterface NetworkInterface_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (networkInterface *NetworkInterface_Spec_ARM) GetName() string {
	return networkInterface.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkInterfaces"
func (networkInterface *NetworkInterface_Spec_ARM) GetType() string {
	return "Microsoft.Network/networkInterfaces"
}

type NetworkInterface_Properties_Spec_ARM struct {
	// DnsSettings: The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettings_ARM `json:"dnsSettings,omitempty"`

	// EnableAcceleratedNetworking: If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `json:"enableAcceleratedNetworking,omitempty"`

	// EnableIPForwarding: Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `json:"enableIPForwarding,omitempty"`

	// IpConfigurations: A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterface_Properties_IpConfigurations_Spec_ARM `json:"ipConfigurations,omitempty"`

	// NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *SubResource_ARM `json:"networkSecurityGroup,omitempty"`
}

type NetworkInterface_Properties_IpConfigurations_Spec_ARM struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Network interface IP configuration properties.
	Properties *NetworkInterfaceIPConfigurationPropertiesFormat_ARM `json:"properties,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceDnsSettings
type NetworkInterfaceDnsSettings_ARM struct {
	// DnsServers: List of DNS servers IP addresses. Use 'AzureProvidedDNS' to switch to azure provided DNS resolution.
	// 'AzureProvidedDNS' value cannot be combined with other IPs, it must be the only value in dnsServers collection.
	DnsServers []string `json:"dnsServers,omitempty"`

	// InternalDnsNameLabel: Relative DNS name for this NIC used for internal communications between VMs in the same virtual
	// network.
	InternalDnsNameLabel *string `json:"internalDnsNameLabel,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource
type SubResource_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceIPConfigurationPropertiesFormat
type NetworkInterfaceIPConfigurationPropertiesFormat_ARM struct {
	// ApplicationGatewayBackendAddressPools: The reference to ApplicationGatewayBackendAddressPool resource.
	ApplicationGatewayBackendAddressPools []SubResource_ARM `json:"applicationGatewayBackendAddressPools,omitempty"`

	// ApplicationSecurityGroups: Application security groups in which the IP configuration is included.
	ApplicationSecurityGroups []SubResource_ARM `json:"applicationSecurityGroups,omitempty"`

	// LoadBalancerBackendAddressPools: The reference to LoadBalancerBackendAddressPool resource.
	LoadBalancerBackendAddressPools []SubResource_ARM `json:"loadBalancerBackendAddressPools,omitempty"`

	// LoadBalancerInboundNatRules: A list of references of LoadBalancerInboundNatRules.
	LoadBalancerInboundNatRules []SubResource_ARM `json:"loadBalancerInboundNatRules,omitempty"`

	// Primary: Whether this is a primary customer address on the network interface.
	Primary *bool `json:"primary,omitempty"`

	// PrivateIPAddress: Private IP address of the IP configuration.
	PrivateIPAddress *string `json:"privateIPAddress,omitempty"`

	// PrivateIPAddressVersion: Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4.
	PrivateIPAddressVersion *NetworkInterfaceIPConfigurationPropertiesFormat_PrivateIPAddressVersion `json:"privateIPAddressVersion,omitempty"`

	// PrivateIPAllocationMethod: The private IP address allocation method.
	PrivateIPAllocationMethod *NetworkInterfaceIPConfigurationPropertiesFormat_PrivateIPAllocationMethod `json:"privateIPAllocationMethod,omitempty"`

	// PublicIPAddress: Public IP address bound to the IP configuration.
	PublicIPAddress *SubResource_ARM `json:"publicIPAddress,omitempty"`

	// Subnet: Subnet bound to the IP configuration.
	Subnet *SubResource_ARM `json:"subnet,omitempty"`

	// VirtualNetworkTaps: The reference to Virtual Network Taps.
	VirtualNetworkTaps []SubResource_ARM `json:"virtualNetworkTaps,omitempty"`
}
