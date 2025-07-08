# SnapshotScheduleSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapshotRetentionCount** | Pointer to **NullableString** |  | [optional] 
**SnapshotSchedule** | [**SnapshotSchedule**](SnapshotSchedule.md) |  | 

## Methods

### NewSnapshotScheduleSetRequest

`func NewSnapshotScheduleSetRequest(snapshotSchedule SnapshotSchedule, ) *SnapshotScheduleSetRequest`

NewSnapshotScheduleSetRequest instantiates a new SnapshotScheduleSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotScheduleSetRequestWithDefaults

`func NewSnapshotScheduleSetRequestWithDefaults() *SnapshotScheduleSetRequest`

NewSnapshotScheduleSetRequestWithDefaults instantiates a new SnapshotScheduleSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnapshotRetentionCount

`func (o *SnapshotScheduleSetRequest) GetSnapshotRetentionCount() string`

GetSnapshotRetentionCount returns the SnapshotRetentionCount field if non-nil, zero value otherwise.

### GetSnapshotRetentionCountOk

`func (o *SnapshotScheduleSetRequest) GetSnapshotRetentionCountOk() (*string, bool)`

GetSnapshotRetentionCountOk returns a tuple with the SnapshotRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRetentionCount

`func (o *SnapshotScheduleSetRequest) SetSnapshotRetentionCount(v string)`

SetSnapshotRetentionCount sets SnapshotRetentionCount field to given value.

### HasSnapshotRetentionCount

`func (o *SnapshotScheduleSetRequest) HasSnapshotRetentionCount() bool`

HasSnapshotRetentionCount returns a boolean if a field has been set.

### SetSnapshotRetentionCountNil

`func (o *SnapshotScheduleSetRequest) SetSnapshotRetentionCountNil(b bool)`

 SetSnapshotRetentionCountNil sets the value for SnapshotRetentionCount to be an explicit nil

### UnsetSnapshotRetentionCount
`func (o *SnapshotScheduleSetRequest) UnsetSnapshotRetentionCount()`

UnsetSnapshotRetentionCount ensures that no value is present for SnapshotRetentionCount, not even an explicit nil
### GetSnapshotSchedule

`func (o *SnapshotScheduleSetRequest) GetSnapshotSchedule() SnapshotSchedule`

GetSnapshotSchedule returns the SnapshotSchedule field if non-nil, zero value otherwise.

### GetSnapshotScheduleOk

`func (o *SnapshotScheduleSetRequest) GetSnapshotScheduleOk() (*SnapshotSchedule, bool)`

GetSnapshotScheduleOk returns a tuple with the SnapshotSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedule

`func (o *SnapshotScheduleSetRequest) SetSnapshotSchedule(v SnapshotSchedule)`

SetSnapshotSchedule sets SnapshotSchedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


