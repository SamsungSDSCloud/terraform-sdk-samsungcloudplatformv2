# NodepoolUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**IsAutoRecovery** | Pointer to **NullableBool** |  | [optional] 
**IsAutoScale** | Pointer to **NullableBool** |  | [optional] 
**MaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**MinNodeCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewNodepoolUpdateRequest

`func NewNodepoolUpdateRequest() *NodepoolUpdateRequest`

NewNodepoolUpdateRequest instantiates a new NodepoolUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolUpdateRequestWithDefaults

`func NewNodepoolUpdateRequestWithDefaults() *NodepoolUpdateRequest`

NewNodepoolUpdateRequestWithDefaults instantiates a new NodepoolUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredNodeCount

`func (o *NodepoolUpdateRequest) GetDesiredNodeCount() int32`

GetDesiredNodeCount returns the DesiredNodeCount field if non-nil, zero value otherwise.

### GetDesiredNodeCountOk

`func (o *NodepoolUpdateRequest) GetDesiredNodeCountOk() (*int32, bool)`

GetDesiredNodeCountOk returns a tuple with the DesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodeCount

`func (o *NodepoolUpdateRequest) SetDesiredNodeCount(v int32)`

SetDesiredNodeCount sets DesiredNodeCount field to given value.

### HasDesiredNodeCount

`func (o *NodepoolUpdateRequest) HasDesiredNodeCount() bool`

HasDesiredNodeCount returns a boolean if a field has been set.

### SetDesiredNodeCountNil

`func (o *NodepoolUpdateRequest) SetDesiredNodeCountNil(b bool)`

 SetDesiredNodeCountNil sets the value for DesiredNodeCount to be an explicit nil

### UnsetDesiredNodeCount
`func (o *NodepoolUpdateRequest) UnsetDesiredNodeCount()`

UnsetDesiredNodeCount ensures that no value is present for DesiredNodeCount, not even an explicit nil
### GetIsAutoRecovery

`func (o *NodepoolUpdateRequest) GetIsAutoRecovery() bool`

GetIsAutoRecovery returns the IsAutoRecovery field if non-nil, zero value otherwise.

### GetIsAutoRecoveryOk

`func (o *NodepoolUpdateRequest) GetIsAutoRecoveryOk() (*bool, bool)`

GetIsAutoRecoveryOk returns a tuple with the IsAutoRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRecovery

`func (o *NodepoolUpdateRequest) SetIsAutoRecovery(v bool)`

SetIsAutoRecovery sets IsAutoRecovery field to given value.

### HasIsAutoRecovery

`func (o *NodepoolUpdateRequest) HasIsAutoRecovery() bool`

HasIsAutoRecovery returns a boolean if a field has been set.

### SetIsAutoRecoveryNil

`func (o *NodepoolUpdateRequest) SetIsAutoRecoveryNil(b bool)`

 SetIsAutoRecoveryNil sets the value for IsAutoRecovery to be an explicit nil

### UnsetIsAutoRecovery
`func (o *NodepoolUpdateRequest) UnsetIsAutoRecovery()`

UnsetIsAutoRecovery ensures that no value is present for IsAutoRecovery, not even an explicit nil
### GetIsAutoScale

`func (o *NodepoolUpdateRequest) GetIsAutoScale() bool`

GetIsAutoScale returns the IsAutoScale field if non-nil, zero value otherwise.

### GetIsAutoScaleOk

`func (o *NodepoolUpdateRequest) GetIsAutoScaleOk() (*bool, bool)`

GetIsAutoScaleOk returns a tuple with the IsAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoScale

`func (o *NodepoolUpdateRequest) SetIsAutoScale(v bool)`

SetIsAutoScale sets IsAutoScale field to given value.

### HasIsAutoScale

`func (o *NodepoolUpdateRequest) HasIsAutoScale() bool`

HasIsAutoScale returns a boolean if a field has been set.

### SetIsAutoScaleNil

`func (o *NodepoolUpdateRequest) SetIsAutoScaleNil(b bool)`

 SetIsAutoScaleNil sets the value for IsAutoScale to be an explicit nil

### UnsetIsAutoScale
`func (o *NodepoolUpdateRequest) UnsetIsAutoScale()`

UnsetIsAutoScale ensures that no value is present for IsAutoScale, not even an explicit nil
### GetMaxNodeCount

`func (o *NodepoolUpdateRequest) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *NodepoolUpdateRequest) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *NodepoolUpdateRequest) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *NodepoolUpdateRequest) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *NodepoolUpdateRequest) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *NodepoolUpdateRequest) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetMinNodeCount

`func (o *NodepoolUpdateRequest) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *NodepoolUpdateRequest) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *NodepoolUpdateRequest) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *NodepoolUpdateRequest) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### SetMinNodeCountNil

`func (o *NodepoolUpdateRequest) SetMinNodeCountNil(b bool)`

 SetMinNodeCountNil sets the value for MinNodeCount to be an explicit nil

### UnsetMinNodeCount
`func (o *NodepoolUpdateRequest) UnsetMinNodeCount()`

UnsetMinNodeCount ensures that no value is present for MinNodeCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


