# VolumeSnapshotRestoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account id | [optional] [default to ""]
**SnapshotId** | Pointer to **string** | Snapshot id | [optional] [default to ""]
**VolumeId** | Pointer to **string** | Volume id | [optional] [default to ""]

## Methods

### NewVolumeSnapshotRestoreResponse

`func NewVolumeSnapshotRestoreResponse() *VolumeSnapshotRestoreResponse`

NewVolumeSnapshotRestoreResponse instantiates a new VolumeSnapshotRestoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotRestoreResponseWithDefaults

`func NewVolumeSnapshotRestoreResponseWithDefaults() *VolumeSnapshotRestoreResponse`

NewVolumeSnapshotRestoreResponseWithDefaults instantiates a new VolumeSnapshotRestoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VolumeSnapshotRestoreResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VolumeSnapshotRestoreResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VolumeSnapshotRestoreResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *VolumeSnapshotRestoreResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *VolumeSnapshotRestoreResponse) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *VolumeSnapshotRestoreResponse) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *VolumeSnapshotRestoreResponse) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *VolumeSnapshotRestoreResponse) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeSnapshotRestoreResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeSnapshotRestoreResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeSnapshotRestoreResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeSnapshotRestoreResponse) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


