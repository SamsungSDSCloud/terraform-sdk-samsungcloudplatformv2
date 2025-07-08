# VolumeSnapshotListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSnapshotPolicy** | Pointer to **bool** | Whether to activate snapshot | [optional] [default to false]
**SnapshotCapacityMb** | Pointer to **float32** | Snapshot capacity(MB) | [optional] [default to 0]
**SnapshotCapacityRate** | Pointer to **float32** | Snapshot capacity rate | [optional] [default to 0]
**SnapshotTotalUsage** | Pointer to **float32** | Snapshot total usage | [optional] [default to 0]
**Snapshots** | Pointer to [**[]VolumeSnapshotListModel**](VolumeSnapshotListModel.md) | List of snapshots | [optional] [default to []]
**VolumeId** | Pointer to **string** | Volume id | [optional] [default to ""]

## Methods

### NewVolumeSnapshotListResponse

`func NewVolumeSnapshotListResponse() *VolumeSnapshotListResponse`

NewVolumeSnapshotListResponse instantiates a new VolumeSnapshotListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotListResponseWithDefaults

`func NewVolumeSnapshotListResponseWithDefaults() *VolumeSnapshotListResponse`

NewVolumeSnapshotListResponseWithDefaults instantiates a new VolumeSnapshotListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSnapshotPolicy

`func (o *VolumeSnapshotListResponse) GetIsSnapshotPolicy() bool`

GetIsSnapshotPolicy returns the IsSnapshotPolicy field if non-nil, zero value otherwise.

### GetIsSnapshotPolicyOk

`func (o *VolumeSnapshotListResponse) GetIsSnapshotPolicyOk() (*bool, bool)`

GetIsSnapshotPolicyOk returns a tuple with the IsSnapshotPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotPolicy

`func (o *VolumeSnapshotListResponse) SetIsSnapshotPolicy(v bool)`

SetIsSnapshotPolicy sets IsSnapshotPolicy field to given value.

### HasIsSnapshotPolicy

`func (o *VolumeSnapshotListResponse) HasIsSnapshotPolicy() bool`

HasIsSnapshotPolicy returns a boolean if a field has been set.

### GetSnapshotCapacityMb

`func (o *VolumeSnapshotListResponse) GetSnapshotCapacityMb() float32`

GetSnapshotCapacityMb returns the SnapshotCapacityMb field if non-nil, zero value otherwise.

### GetSnapshotCapacityMbOk

`func (o *VolumeSnapshotListResponse) GetSnapshotCapacityMbOk() (*float32, bool)`

GetSnapshotCapacityMbOk returns a tuple with the SnapshotCapacityMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCapacityMb

`func (o *VolumeSnapshotListResponse) SetSnapshotCapacityMb(v float32)`

SetSnapshotCapacityMb sets SnapshotCapacityMb field to given value.

### HasSnapshotCapacityMb

`func (o *VolumeSnapshotListResponse) HasSnapshotCapacityMb() bool`

HasSnapshotCapacityMb returns a boolean if a field has been set.

### GetSnapshotCapacityRate

`func (o *VolumeSnapshotListResponse) GetSnapshotCapacityRate() float32`

GetSnapshotCapacityRate returns the SnapshotCapacityRate field if non-nil, zero value otherwise.

### GetSnapshotCapacityRateOk

`func (o *VolumeSnapshotListResponse) GetSnapshotCapacityRateOk() (*float32, bool)`

GetSnapshotCapacityRateOk returns a tuple with the SnapshotCapacityRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotCapacityRate

`func (o *VolumeSnapshotListResponse) SetSnapshotCapacityRate(v float32)`

SetSnapshotCapacityRate sets SnapshotCapacityRate field to given value.

### HasSnapshotCapacityRate

`func (o *VolumeSnapshotListResponse) HasSnapshotCapacityRate() bool`

HasSnapshotCapacityRate returns a boolean if a field has been set.

### GetSnapshotTotalUsage

`func (o *VolumeSnapshotListResponse) GetSnapshotTotalUsage() float32`

GetSnapshotTotalUsage returns the SnapshotTotalUsage field if non-nil, zero value otherwise.

### GetSnapshotTotalUsageOk

`func (o *VolumeSnapshotListResponse) GetSnapshotTotalUsageOk() (*float32, bool)`

GetSnapshotTotalUsageOk returns a tuple with the SnapshotTotalUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTotalUsage

`func (o *VolumeSnapshotListResponse) SetSnapshotTotalUsage(v float32)`

SetSnapshotTotalUsage sets SnapshotTotalUsage field to given value.

### HasSnapshotTotalUsage

`func (o *VolumeSnapshotListResponse) HasSnapshotTotalUsage() bool`

HasSnapshotTotalUsage returns a boolean if a field has been set.

### GetSnapshots

`func (o *VolumeSnapshotListResponse) GetSnapshots() []VolumeSnapshotListModel`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *VolumeSnapshotListResponse) GetSnapshotsOk() (*[]VolumeSnapshotListModel, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *VolumeSnapshotListResponse) SetSnapshots(v []VolumeSnapshotListModel)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *VolumeSnapshotListResponse) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeSnapshotListResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeSnapshotListResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeSnapshotListResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeSnapshotListResponse) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


