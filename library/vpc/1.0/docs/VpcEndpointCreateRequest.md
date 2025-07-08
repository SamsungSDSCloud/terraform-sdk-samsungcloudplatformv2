# VpcEndpointCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description | [optional] [default to ""]
**EndpointIpAddress** | **string** | VPC Endpoint IP Address | 
**Name** | **string** | VPC Endpoint Name | 
**ResourceInfo** | **string** | VPC Endpoint Resource Key Info | 
**ResourceKey** | **string** | VPC Endpoint Resource Key | 
**ResourceType** | [**VpcEndpointResourceType**](VpcEndpointResourceType.md) | VPC Endpoint Resource Type | 
**SubnetId** | **string** | Subnet Id | 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**VpcId** | **string** | VPC Id | 

## Methods

### NewVpcEndpointCreateRequest

`func NewVpcEndpointCreateRequest(endpointIpAddress string, name string, resourceInfo string, resourceKey string, resourceType VpcEndpointResourceType, subnetId string, vpcId string, ) *VpcEndpointCreateRequest`

NewVpcEndpointCreateRequest instantiates a new VpcEndpointCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcEndpointCreateRequestWithDefaults

`func NewVpcEndpointCreateRequestWithDefaults() *VpcEndpointCreateRequest`

NewVpcEndpointCreateRequestWithDefaults instantiates a new VpcEndpointCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *VpcEndpointCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VpcEndpointCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VpcEndpointCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VpcEndpointCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpointIpAddress

`func (o *VpcEndpointCreateRequest) GetEndpointIpAddress() string`

GetEndpointIpAddress returns the EndpointIpAddress field if non-nil, zero value otherwise.

### GetEndpointIpAddressOk

`func (o *VpcEndpointCreateRequest) GetEndpointIpAddressOk() (*string, bool)`

GetEndpointIpAddressOk returns a tuple with the EndpointIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointIpAddress

`func (o *VpcEndpointCreateRequest) SetEndpointIpAddress(v string)`

SetEndpointIpAddress sets EndpointIpAddress field to given value.


### GetName

`func (o *VpcEndpointCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpcEndpointCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpcEndpointCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetResourceInfo

`func (o *VpcEndpointCreateRequest) GetResourceInfo() string`

GetResourceInfo returns the ResourceInfo field if non-nil, zero value otherwise.

### GetResourceInfoOk

`func (o *VpcEndpointCreateRequest) GetResourceInfoOk() (*string, bool)`

GetResourceInfoOk returns a tuple with the ResourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInfo

`func (o *VpcEndpointCreateRequest) SetResourceInfo(v string)`

SetResourceInfo sets ResourceInfo field to given value.


### GetResourceKey

`func (o *VpcEndpointCreateRequest) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *VpcEndpointCreateRequest) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *VpcEndpointCreateRequest) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *VpcEndpointCreateRequest) GetResourceType() VpcEndpointResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *VpcEndpointCreateRequest) GetResourceTypeOk() (*VpcEndpointResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *VpcEndpointCreateRequest) SetResourceType(v VpcEndpointResourceType)`

SetResourceType sets ResourceType field to given value.


### GetSubnetId

`func (o *VpcEndpointCreateRequest) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *VpcEndpointCreateRequest) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *VpcEndpointCreateRequest) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTags

`func (o *VpcEndpointCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VpcEndpointCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VpcEndpointCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VpcEndpointCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVpcId

`func (o *VpcEndpointCreateRequest) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpcEndpointCreateRequest) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpcEndpointCreateRequest) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


