# VolumeSnapshotListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDt** | Pointer to **string** | Created dt | [optional] [default to ""]
**SnapshotId** | Pointer to **string** | Snapshot id | [optional] [default to ""]
**SnapshotName** | Pointer to **string** | Snapshot name | [optional] 
**SnapshotSizeMb** | Pointer to **float32** | Snapshot size(MB) | [optional] [default to 0]

## Methods

### NewVolumeSnapshotListModel

`func NewVolumeSnapshotListModel() *VolumeSnapshotListModel`

NewVolumeSnapshotListModel instantiates a new VolumeSnapshotListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotListModelWithDefaults

`func NewVolumeSnapshotListModelWithDefaults() *VolumeSnapshotListModel`

NewVolumeSnapshotListModelWithDefaults instantiates a new VolumeSnapshotListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDt

`func (o *VolumeSnapshotListModel) GetCreatedDt() string`

GetCreatedDt returns the CreatedDt field if non-nil, zero value otherwise.

### GetCreatedDtOk

`func (o *VolumeSnapshotListModel) GetCreatedDtOk() (*string, bool)`

GetCreatedDtOk returns a tuple with the CreatedDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDt

`func (o *VolumeSnapshotListModel) SetCreatedDt(v string)`

SetCreatedDt sets CreatedDt field to given value.

### HasCreatedDt

`func (o *VolumeSnapshotListModel) HasCreatedDt() bool`

HasCreatedDt returns a boolean if a field has been set.

### GetSnapshotId

`func (o *VolumeSnapshotListModel) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *VolumeSnapshotListModel) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *VolumeSnapshotListModel) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *VolumeSnapshotListModel) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetSnapshotName

`func (o *VolumeSnapshotListModel) GetSnapshotName() string`

GetSnapshotName returns the SnapshotName field if non-nil, zero value otherwise.

### GetSnapshotNameOk

`func (o *VolumeSnapshotListModel) GetSnapshotNameOk() (*string, bool)`

GetSnapshotNameOk returns a tuple with the SnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotName

`func (o *VolumeSnapshotListModel) SetSnapshotName(v string)`

SetSnapshotName sets SnapshotName field to given value.

### HasSnapshotName

`func (o *VolumeSnapshotListModel) HasSnapshotName() bool`

HasSnapshotName returns a boolean if a field has been set.

### GetSnapshotSizeMb

`func (o *VolumeSnapshotListModel) GetSnapshotSizeMb() float32`

GetSnapshotSizeMb returns the SnapshotSizeMb field if non-nil, zero value otherwise.

### GetSnapshotSizeMbOk

`func (o *VolumeSnapshotListModel) GetSnapshotSizeMbOk() (*float32, bool)`

GetSnapshotSizeMbOk returns a tuple with the SnapshotSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSizeMb

`func (o *VolumeSnapshotListModel) SetSnapshotSizeMb(v float32)`

SetSnapshotSizeMb sets SnapshotSizeMb field to given value.

### HasSnapshotSizeMb

`func (o *VolumeSnapshotListModel) HasSnapshotSizeMb() bool`

HasSnapshotSizeMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


