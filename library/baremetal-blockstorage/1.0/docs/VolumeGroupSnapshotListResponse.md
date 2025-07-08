# VolumeGroupSnapshotListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSnapshotPolicy** | Pointer to **bool** | Whether to activate snapshot | [optional] [default to false]
**Snapshots** | Pointer to [**[]VolumeGroupSnapshotListModel**](VolumeGroupSnapshotListModel.md) | List of snapshots | [optional] [default to []]
**VolumeGroupId** | Pointer to **string** | Volume group id | [optional] [default to ""]

## Methods

### NewVolumeGroupSnapshotListResponse

`func NewVolumeGroupSnapshotListResponse() *VolumeGroupSnapshotListResponse`

NewVolumeGroupSnapshotListResponse instantiates a new VolumeGroupSnapshotListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupSnapshotListResponseWithDefaults

`func NewVolumeGroupSnapshotListResponseWithDefaults() *VolumeGroupSnapshotListResponse`

NewVolumeGroupSnapshotListResponseWithDefaults instantiates a new VolumeGroupSnapshotListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSnapshotPolicy

`func (o *VolumeGroupSnapshotListResponse) GetIsSnapshotPolicy() bool`

GetIsSnapshotPolicy returns the IsSnapshotPolicy field if non-nil, zero value otherwise.

### GetIsSnapshotPolicyOk

`func (o *VolumeGroupSnapshotListResponse) GetIsSnapshotPolicyOk() (*bool, bool)`

GetIsSnapshotPolicyOk returns a tuple with the IsSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotPolicy

`func (o *VolumeGroupSnapshotListResponse) SetIsSnapshotPolicy(v bool)`

SetIsSnapshotPolicy sets IsSnapshotPolicy field to given value.

### HasIsSnapshotPolicy

`func (o *VolumeGroupSnapshotListResponse) HasIsSnapshotPolicy() bool`

HasIsSnapshotPolicy returns a boolean if a field has been set.

### GetSnapshots

`func (o *VolumeGroupSnapshotListResponse) GetSnapshots() []VolumeGroupSnapshotListModel`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *VolumeGroupSnapshotListResponse) GetSnapshotsOk() (*[]VolumeGroupSnapshotListModel, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *VolumeGroupSnapshotListResponse) SetSnapshots(v []VolumeGroupSnapshotListModel)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *VolumeGroupSnapshotListResponse) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetVolumeGroupId

`func (o *VolumeGroupSnapshotListResponse) GetVolumeGroupId() string`

GetVolumeGroupId returns the VolumeGroupId field if non-nil, zero value otherwise.

### GetVolumeGroupIdOk

`func (o *VolumeGroupSnapshotListResponse) GetVolumeGroupIdOk() (*string, bool)`

GetVolumeGroupIdOk returns a tuple with the VolumeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupId

`func (o *VolumeGroupSnapshotListResponse) SetVolumeGroupId(v string)`

SetVolumeGroupId sets VolumeGroupId field to given value.

### HasVolumeGroupId

`func (o *VolumeGroupSnapshotListResponse) HasVolumeGroupId() bool`

HasVolumeGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


