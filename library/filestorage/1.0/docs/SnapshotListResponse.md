# SnapshotListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**SnapshotSizeTotalByte** | Pointer to **NullableInt32** |  | [optional] 
**Snapshots** | [**[]Snapshot**](Snapshot.md) |  | 

## Methods

### NewSnapshotListResponse

`func NewSnapshotListResponse(count int32, snapshots []Snapshot, ) *SnapshotListResponse`

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


### GetSnapshotSizeTotalByte

`func (o *SnapshotListResponse) GetSnapshotSizeTotalByte() int32`

GetSnapshotSizeTotalByte returns the SnapshotSizeTotalByte field if non-nil, zero value otherwise.

### GetSnapshotSizeTotalByteOk

`func (o *SnapshotListResponse) GetSnapshotSizeTotalByteOk() (*int32, bool)`

GetSnapshotSizeTotalByteOk returns a tuple with the SnapshotSizeTotalByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSizeTotalByte

`func (o *SnapshotListResponse) SetSnapshotSizeTotalByte(v int32)`

SetSnapshotSizeTotalByte sets SnapshotSizeTotalByte field to given value.

### HasSnapshotSizeTotalByte

`func (o *SnapshotListResponse) HasSnapshotSizeTotalByte() bool`

HasSnapshotSizeTotalByte returns a boolean if a field has been set.

### SetSnapshotSizeTotalByteNil

`func (o *SnapshotListResponse) SetSnapshotSizeTotalByteNil(b bool)`

 SetSnapshotSizeTotalByteNil sets the value for SnapshotSizeTotalByte to be an explicit nil

### UnsetSnapshotSizeTotalByte
`func (o *SnapshotListResponse) UnsetSnapshotSizeTotalByte()`

UnsetSnapshotSizeTotalByte ensures that no value is present for SnapshotSizeTotalByte, not even an explicit nil
### GetSnapshots

`func (o *SnapshotListResponse) GetSnapshots() []Snapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *SnapshotListResponse) GetSnapshotsOk() (*[]Snapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *SnapshotListResponse) SetSnapshots(v []Snapshot)`

SetSnapshots sets Snapshots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


