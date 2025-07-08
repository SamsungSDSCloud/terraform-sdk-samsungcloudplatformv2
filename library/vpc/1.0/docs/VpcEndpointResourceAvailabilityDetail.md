# VpcEndpointResourceAvailabilityDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Availability of Endpoint Resource | 
**Message** | **string** | Vpc Endpoint Message | 
**ResourceKey** | **string** | VPC Endpoint Resource Key | 
**ResourceType** | [**VpcEndpointResourceType**](VpcEndpointResourceType.md) | VPC Endpoint Resource Type | 
**VpcId** | **string** | VPC Id | 

## Methods

### NewVpcEndpointResourceAvailabilityDetail

`func NewVpcEndpointResourceAvailabilityDetail(available bool, message string, resourceKey string, resourceType VpcEndpointResourceType, vpcId string, ) *VpcEndpointResourceAvailabilityDetail`

NewVpcEndpointResourceAvailabilityDetail instantiates a new VpcEndpointResourceAvailabilityDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpcEndpointResourceAvailabilityDetailWithDefaults

`func NewVpcEndpointResourceAvailabilityDetailWithDefaults() *VpcEndpointResourceAvailabilityDetail`

NewVpcEndpointResourceAvailabilityDetailWithDefaults instantiates a new VpcEndpointResourceAvailabilityDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *VpcEndpointResourceAvailabilityDetail) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *VpcEndpointResourceAvailabilityDetail) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *VpcEndpointResourceAvailabilityDetail) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetMessage

`func (o *VpcEndpointResourceAvailabilityDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VpcEndpointResourceAvailabilityDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VpcEndpointResourceAvailabilityDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetResourceKey

`func (o *VpcEndpointResourceAvailabilityDetail) GetResourceKey() string`

GetResourceKey returns the ResourceKey field if non-nil, zero value otherwise.

### GetResourceKeyOk

`func (o *VpcEndpointResourceAvailabilityDetail) GetResourceKeyOk() (*string, bool)`

GetResourceKeyOk returns a tuple with the ResourceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceKey

`func (o *VpcEndpointResourceAvailabilityDetail) SetResourceKey(v string)`

SetResourceKey sets ResourceKey field to given value.


### GetResourceType

`func (o *VpcEndpointResourceAvailabilityDetail) GetResourceType() VpcEndpointResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *VpcEndpointResourceAvailabilityDetail) GetResourceTypeOk() (*VpcEndpointResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *VpcEndpointResourceAvailabilityDetail) SetResourceType(v VpcEndpointResourceType)`

SetResourceType sets ResourceType field to given value.


### GetVpcId

`func (o *VpcEndpointResourceAvailabilityDetail) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *VpcEndpointResourceAvailabilityDetail) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *VpcEndpointResourceAvailabilityDetail) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


