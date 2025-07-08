# EventStreamsClusterAddInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | **int32** | Instance Count | 
**PublicIpIds** | Pointer to **[]string** |  | [optional] 
**ServiceIpAddresses** | **[]string** | User subnet IP address | 

## Methods

### NewEventStreamsClusterAddInstancesRequest

`func NewEventStreamsClusterAddInstancesRequest(instanceCount int32, serviceIpAddresses []string, ) *EventStreamsClusterAddInstancesRequest`

NewEventStreamsClusterAddInstancesRequest instantiates a new EventStreamsClusterAddInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStreamsClusterAddInstancesRequestWithDefaults

`func NewEventStreamsClusterAddInstancesRequestWithDefaults() *EventStreamsClusterAddInstancesRequest`

NewEventStreamsClusterAddInstancesRequestWithDefaults instantiates a new EventStreamsClusterAddInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCount

`func (o *EventStreamsClusterAddInstancesRequest) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *EventStreamsClusterAddInstancesRequest) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *EventStreamsClusterAddInstancesRequest) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.


### GetPublicIpIds

`func (o *EventStreamsClusterAddInstancesRequest) GetPublicIpIds() []string`

GetPublicIpIds returns the PublicIpIds field if non-nil, zero value otherwise.

### GetPublicIpIdsOk

`func (o *EventStreamsClusterAddInstancesRequest) GetPublicIpIdsOk() (*[]string, bool)`

GetPublicIpIdsOk returns a tuple with the PublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpIds

`func (o *EventStreamsClusterAddInstancesRequest) SetPublicIpIds(v []string)`

SetPublicIpIds sets PublicIpIds field to given value.

### HasPublicIpIds

`func (o *EventStreamsClusterAddInstancesRequest) HasPublicIpIds() bool`

HasPublicIpIds returns a boolean if a field has been set.

### SetPublicIpIdsNil

`func (o *EventStreamsClusterAddInstancesRequest) SetPublicIpIdsNil(b bool)`

 SetPublicIpIdsNil sets the value for PublicIpIds to be an explicit nil

### UnsetPublicIpIds
`func (o *EventStreamsClusterAddInstancesRequest) UnsetPublicIpIds()`

UnsetPublicIpIds ensures that no value is present for PublicIpIds, not even an explicit nil
### GetServiceIpAddresses

`func (o *EventStreamsClusterAddInstancesRequest) GetServiceIpAddresses() []string`

GetServiceIpAddresses returns the ServiceIpAddresses field if non-nil, zero value otherwise.

### GetServiceIpAddressesOk

`func (o *EventStreamsClusterAddInstancesRequest) GetServiceIpAddressesOk() (*[]string, bool)`

GetServiceIpAddressesOk returns a tuple with the ServiceIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpAddresses

`func (o *EventStreamsClusterAddInstancesRequest) SetServiceIpAddresses(v []string)`

SetServiceIpAddresses sets ServiceIpAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


