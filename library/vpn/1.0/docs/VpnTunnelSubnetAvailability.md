# VpnTunnelSubnetAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Availability** | [**VpnTunnelRemoteSubnetAvailability**](VpnTunnelRemoteSubnetAvailability.md) | VPN Tunnel Remote Subnet Availability | 
**RemoteSubnet** | **string** | VPN Tunnel IPSec Remote Subnet | 
**VpnGatewayId** | **string** | VPN Gateway ID | 

## Methods

### NewVpnTunnelSubnetAvailability

`func NewVpnTunnelSubnetAvailability(availability VpnTunnelRemoteSubnetAvailability, remoteSubnet string, vpnGatewayId string, ) *VpnTunnelSubnetAvailability`

NewVpnTunnelSubnetAvailability instantiates a new VpnTunnelSubnetAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnTunnelSubnetAvailabilityWithDefaults

`func NewVpnTunnelSubnetAvailabilityWithDefaults() *VpnTunnelSubnetAvailability`

NewVpnTunnelSubnetAvailabilityWithDefaults instantiates a new VpnTunnelSubnetAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailability

`func (o *VpnTunnelSubnetAvailability) GetAvailability() VpnTunnelRemoteSubnetAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *VpnTunnelSubnetAvailability) GetAvailabilityOk() (*VpnTunnelRemoteSubnetAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *VpnTunnelSubnetAvailability) SetAvailability(v VpnTunnelRemoteSubnetAvailability)`

SetAvailability sets Availability field to given value.


### GetRemoteSubnet

`func (o *VpnTunnelSubnetAvailability) GetRemoteSubnet() string`

GetRemoteSubnet returns the RemoteSubnet field if non-nil, zero value otherwise.

### GetRemoteSubnetOk

`func (o *VpnTunnelSubnetAvailability) GetRemoteSubnetOk() (*string, bool)`

GetRemoteSubnetOk returns a tuple with the RemoteSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubnet

`func (o *VpnTunnelSubnetAvailability) SetRemoteSubnet(v string)`

SetRemoteSubnet sets RemoteSubnet field to given value.


### GetVpnGatewayId

`func (o *VpnTunnelSubnetAvailability) GetVpnGatewayId() string`

GetVpnGatewayId returns the VpnGatewayId field if non-nil, zero value otherwise.

### GetVpnGatewayIdOk

`func (o *VpnTunnelSubnetAvailability) GetVpnGatewayIdOk() (*string, bool)`

GetVpnGatewayIdOk returns a tuple with the VpnGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnGatewayId

`func (o *VpnTunnelSubnetAvailability) SetVpnGatewayId(v string)`

SetVpnGatewayId sets VpnGatewayId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


