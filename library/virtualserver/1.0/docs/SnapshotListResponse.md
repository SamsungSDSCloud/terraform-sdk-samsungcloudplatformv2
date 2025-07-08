# SnapshotListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Snapshots** | [**[]SnapshotDetailResponse**](SnapshotDetailResponse.md) |  | 

## Methods

### NewSnapshotListResponse

`func NewSnapshotListResponse(snapshots []SnapshotDetailResponse, ) *SnapshotListResponse`

NewSnapshotListResponse instantiates a new SnapshotListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotListResponseWithDefaults

`func NewSnapshotListResponseWithDefaults() *SnapshotListResponse`

NewSnapshotListResponseWithDefaults instantiates a new SnapshotListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SnapshotListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SnapshotListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SnapshotListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SnapshotListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *SnapshotListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *SnapshotListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetSnapshots

`func (o *SnapshotListResponse) GetSnapshots() []SnapshotDetailResponse`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *SnapshotListResponse) GetSnapshotsOk() (*[]SnapshotDetailResponse, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *SnapshotListResponse) SetSnapshots(v []SnapshotDetailResponse)`

SetSnapshots sets Snapshots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


