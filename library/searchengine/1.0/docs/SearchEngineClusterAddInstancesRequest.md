# SearchEngineClusterAddInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | **int32** | Instance Count | 
**PublicIpIds** | Pointer to **[]string** |  | [optional] 
**ServiceIpAddresses** | **[]string** | User subnet IP address | 

## Methods

### NewSearchEngineClusterAddInstancesRequest

`func NewSearchEngineClusterAddInstancesRequest(instanceCount int32, serviceIpAddresses []string, ) *SearchEngineClusterAddInstancesRequest`

NewSearchEngineClusterAddInstancesRequest instantiates a new SearchEngineClusterAddInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEngineClusterAddInstancesRequestWithDefaults

`func NewSearchEngineClusterAddInstancesRequestWithDefaults() *SearchEngineClusterAddInstancesRequest`

NewSearchEngineClusterAddInstancesRequestWithDefaults instantiates a new SearchEngineClusterAddInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCount

`func (o *SearchEngineClusterAddInstancesRequest) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *SearchEngineClusterAddInstancesRequest) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *SearchEngineClusterAddInstancesRequest) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.


### GetPublicIpIds

`func (o *SearchEngineClusterAddInstancesRequest) GetPublicIpIds() []string`

GetPublicIpIds returns the PublicIpIds field if non-nil, zero value otherwise.

### GetPublicIpIdsOk

`func (o *SearchEngineClusterAddInstancesRequest) GetPublicIpIdsOk() (*[]string, bool)`

GetPublicIpIdsOk returns a tuple with the PublicIpIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpIds

`func (o *SearchEngineClusterAddInstancesRequest) SetPublicIpIds(v []string)`

SetPublicIpIds sets PublicIpIds field to given value.

### HasPublicIpIds

`func (o *SearchEngineClusterAddInstancesRequest) HasPublicIpIds() bool`

HasPublicIpIds returns a boolean if a field has been set.

### SetPublicIpIdsNil

`func (o *SearchEngineClusterAddInstancesRequest) SetPublicIpIdsNil(b bool)`

 SetPublicIpIdsNil sets the value for PublicIpIds to be an explicit nil

### UnsetPublicIpIds
`func (o *SearchEngineClusterAddInstancesRequest) UnsetPublicIpIds()`

UnsetPublicIpIds ensures that no value is present for PublicIpIds, not even an explicit nil
### GetServiceIpAddresses

`func (o *SearchEngineClusterAddInstancesRequest) GetServiceIpAddresses() []string`

GetServiceIpAddresses returns the ServiceIpAddresses field if non-nil, zero value otherwise.

### GetServiceIpAddressesOk

`func (o *SearchEngineClusterAddInstancesRequest) GetServiceIpAddressesOk() (*[]string, bool)`

GetServiceIpAddressesOk returns a tuple with the ServiceIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIpAddresses

`func (o *SearchEngineClusterAddInstancesRequest) SetServiceIpAddresses(v []string)`

SetServiceIpAddresses sets ServiceIpAddresses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


