# VpnGatewayCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | **string** | VPN Gateway IP Address | 
**IpId** | **string** | VPN Gateway IP ID | 
**IpType** | **string** | VPN Gateway IP Type | 
**Name** | **string** | VPN Gateway Name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**VpcId** | **string** | VPC Id | 

## Methods

### NewVpnGatewayCreateRequest

`func NewVpnGatewayCreateRequest(ipAddress string, ipId string, ipType string, name string, vpcId string, ) *VpnGatewayCreateRequest`

NewVpnGatewayCreateRequest instantiates a new VpnGatewayCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnGatewayCreateRequestWithDefaults

`func NewVpnGatewayCreateRequestWithDefaults() *VpnGatewayCreateRequest`

NewVpnGatewayCreateRequestWithDefaults instantiates a new VpnGatewayCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VpnGatewayCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpnGatewayCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpnGatewayCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpnGatewayCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VpnGatewayCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VpnGatewayCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIpAddress

`func (o *VpnGatewayCreateRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *VpnGatewayCreateRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *VpnGatewayCreateRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetIpId

`func (o *VpnGatewayCreateRequest) GetIpId() string`

GetIpId returns the IpId field if non-nil, zero value otherwise.

### GetIpIdOk

`func (o *VpnGatewayCreateRequest) GetIpIdOk() (*string, bool)`

GetIpIdOk returns a tuple with the IpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpId

`func (o *VpnGatewayCreateRequest) SetIpId(v string)`

SetIpId sets IpId field to given value.


### GetIpType

`func (o *VpnGatewayCreateRequest) GetIpType() string`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *VpnGatewayCreateRequest) GetIpTypeOk() (*string, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *VpnGatewayCreateRequest) SetIpType(v string)`

SetIpType sets IpType field to given value.


### GetName

`func (o *VpnGatewayCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnGatewayCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnGatewayCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *VpnGatewayCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpnGatewayCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpnGatewayCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpnGatewayCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVpcId

`func (o *VpnGatewayCreateRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpnGatewayCreateRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpnGatewayCreateRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


