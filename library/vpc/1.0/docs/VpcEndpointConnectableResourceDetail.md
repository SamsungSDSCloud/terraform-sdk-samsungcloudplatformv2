# VpcEndpointConnectableResourceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceInfo** | **string** | VPC Endpoint Resource Key Info | 
**ResourceKey** | **string** | VPC Endpoint Resource Key | 
**ResourceType** | [**VpcEndpointResourceType**](VpcEndpointResourceType.md) | VPC Endpoint Resource Type | 

## Methods

### NewVpcEndpointConnectableResourceDetail

`func NewVpcEndpointConnectableResourceDetail(resourceInfo string, resourceKey string, resourceType VpcEndpointResourceType, ) *VpcEndpointConnectableResourceDetail`

NewVpcEndpointConnectableResourceDetail instantiates a new VpcEndpointConnectableResourceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcEndpointConnectableResourceDetailWithDefaults

`func NewVpcEndpointConnectableResourceDetailWithDefaults() *VpcEndpointConnectableResourceDetail`

NewVpcEndpointConnectableResourceDetailWithDefaults instantiates a new VpcEndpointConnectableResourceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceInfo

`func (o *VpcEndpointConnectableResourceDetail) GetResourceInfo() string`

GetResourceInfo returns the ResourceInfo field if non-nil, zero value otherwise.

### GetResourceInfoOk

`func (o *VpcEndpointConnectableResourceDetail) GetResourceInfoOk() (*string, bool)`

GetResourceInfoOk returns a tuple with the ResourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceInfo

`func (o *VpcEndpointConnectableResourceDetail) SetResourceInfo(v string)`

SetResourceInfo sets ResourceInfo field to given value.


### GetResourceKey

`func (o *VpcEndpointConnectableResourceDetail) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *VpcEndpointConnectableResourceDetail) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *VpcEndpointConnectableResourceDetail) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *VpcEndpointConnectableResourceDetail) GetResourceType() VpcEndpointResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *VpcEndpointConnectableResourceDetail) GetResourceTypeOk() (*VpcEndpointResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *VpcEndpointConnectableResourceDetail) SetResourceType(v VpcEndpointResourceType)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


