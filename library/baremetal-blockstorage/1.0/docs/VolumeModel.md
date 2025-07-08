# VolumeModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account id | [optional] 
**Attachments** | Pointer to [**[]AttachedServer**](AttachedServer.md) | Attached server | [optional] 
**CreatedAt** | Pointer to **string** | Created at | [optional] 
**CreatedBy** | Pointer to **string** | Created by | [optional] 
**DiskType** | Pointer to [**DiskType**](DiskType.md) | Disk type | [optional] 
**ExternalErrors** | Pointer to [**[]ExternalError**](ExternalError.md) | Error during external call | [optional] 
**HasOrigin** | Pointer to **bool** | Whether origin relation exists | [optional] 
**HasReplication** | Pointer to **bool** | Whether replication relation exists | [optional] 
**Id** | Pointer to **string** | Id | [optional] 
**IsEncryption** | Pointer to **bool** | Whether to apply encryption | [optional] 
**IscsiTargetIps** | Pointer to **[]string** | List of iscsi ips | [optional] 
**ModifiedAt** | Pointer to **string** | Modified at | [optional] 
**ModifiedBy** | Pointer to **string** | Modified by | [optional] 
**Name** | Pointer to **string** | Volume name | [optional] 
**PreviousState** | Pointer to [**BlockStorageState**](BlockStorageState.md) | Previous state | [optional] 
**Purpose** | Pointer to [**BlockStoragePurpose**](BlockStoragePurpose.md) | Volume purpose | [optional] 
**SizeGb** | Pointer to **int32** | Size gb | [optional] 
**SnapshotRate** | Pointer to **int32** | Snapshot rate | [optional] [default to 0]
**SnapshotSchedule** | Pointer to [**SnapshotSchedule**](SnapshotSchedule.md) | Snapshot schedule | [optional] 
**Srn** | Pointer to **string** | SRN | [optional] 
**State** | Pointer to [**BlockStorageState**](BlockStorageState.md) | Current state | [optional] 
**StorageVolumeName** | Pointer to **string** | Storage volume name | [optional] 
**VolumeGroup** | Pointer to [**VolumeGroupInfo**](VolumeGroupInfo.md) | Volume group | [optional] 

## Methods

### NewVolumeModel

`func NewVolumeModel() *VolumeModel`

NewVolumeModel instantiates a new VolumeModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeModelWithDefaults

`func NewVolumeModelWithDefaults() *VolumeModel`

NewVolumeModelWithDefaults instantiates a new VolumeModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VolumeModel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VolumeModel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VolumeModel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *VolumeModel) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAttachments

`func (o *VolumeModel) GetAttachments() []AttachedServer`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumeModel) GetAttachmentsOk() (*[]AttachedServer, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumeModel) SetAttachments(v []AttachedServer)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *VolumeModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumeModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VolumeModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VolumeModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VolumeModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VolumeModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDiskType

`func (o *VolumeModel) GetDiskType() DiskType`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *VolumeModel) GetDiskTypeOk() (*DiskType, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *VolumeModel) SetDiskType(v DiskType)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *VolumeModel) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetExternalErrors

`func (o *VolumeModel) GetExternalErrors() []ExternalError`

GetExternalErrors returns the ExternalErrors field if non-nil, zero value otherwise.

### GetExternalErrorsOk

`func (o *VolumeModel) GetExternalErrorsOk() (*[]ExternalError, bool)`

GetExternalErrorsOk returns a tuple with the ExternalErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalErrors

`func (o *VolumeModel) SetExternalErrors(v []ExternalError)`

SetExternalErrors sets ExternalErrors field to given value.

### HasExternalErrors

`func (o *VolumeModel) HasExternalErrors() bool`

HasExternalErrors returns a boolean if a field has been set.

### GetHasOrigin

`func (o *VolumeModel) GetHasOrigin() bool`

GetHasOrigin returns the HasOrigin field if non-nil, zero value otherwise.

### GetHasOriginOk

`func (o *VolumeModel) GetHasOriginOk() (*bool, bool)`

GetHasOriginOk returns a tuple with the HasOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOrigin

`func (o *VolumeModel) SetHasOrigin(v bool)`

SetHasOrigin sets HasOrigin field to given value.

### HasHasOrigin

`func (o *VolumeModel) HasHasOrigin() bool`

HasHasOrigin returns a boolean if a field has been set.

### GetHasReplication

`func (o *VolumeModel) GetHasReplication() bool`

GetHasReplication returns the HasReplication field if non-nil, zero value otherwise.

### GetHasReplicationOk

`func (o *VolumeModel) GetHasReplicationOk() (*bool, bool)`

GetHasReplicationOk returns a tuple with the HasReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReplication

`func (o *VolumeModel) SetHasReplication(v bool)`

SetHasReplication sets HasReplication field to given value.

### HasHasReplication

`func (o *VolumeModel) HasHasReplication() bool`

HasHasReplication returns a boolean if a field has been set.

### GetId

`func (o *VolumeModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsEncryption

`func (o *VolumeModel) GetIsEncryption() bool`

GetIsEncryption returns the IsEncryption field if non-nil, zero value otherwise.

### GetIsEncryptionOk

`func (o *VolumeModel) GetIsEncryptionOk() (*bool, bool)`

GetIsEncryptionOk returns a tuple with the IsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEncryption

`func (o *VolumeModel) SetIsEncryption(v bool)`

SetIsEncryption sets IsEncryption field to given value.

### HasIsEncryption

`func (o *VolumeModel) HasIsEncryption() bool`

HasIsEncryption returns a boolean if a field has been set.

### GetIscsiTargetIps

`func (o *VolumeModel) GetIscsiTargetIps() []string`

GetIscsiTargetIps returns the IscsiTargetIps field if non-nil, zero value otherwise.

### GetIscsiTargetIpsOk

`func (o *VolumeModel) GetIscsiTargetIpsOk() (*[]string, bool)`

GetIscsiTargetIpsOk returns a tuple with the IscsiTargetIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiTargetIps

`func (o *VolumeModel) SetIscsiTargetIps(v []string)`

SetIscsiTargetIps sets IscsiTargetIps field to given value.

### HasIscsiTargetIps

`func (o *VolumeModel) HasIscsiTargetIps() bool`

HasIscsiTargetIps returns a boolean if a field has been set.

### GetModifiedAt

`func (o *VolumeModel) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VolumeModel) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VolumeModel) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *VolumeModel) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *VolumeModel) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VolumeModel) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VolumeModel) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *VolumeModel) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *VolumeModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreviousState

`func (o *VolumeModel) GetPreviousState() BlockStorageState`

GetPreviousState returns the PreviousState field if non-nil, zero value otherwise.

### GetPreviousStateOk

`func (o *VolumeModel) GetPreviousStateOk() (*BlockStorageState, bool)`

GetPreviousStateOk returns a tuple with the PreviousState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousState

`func (o *VolumeModel) SetPreviousState(v BlockStorageState)`

SetPreviousState sets PreviousState field to given value.

### HasPreviousState

`func (o *VolumeModel) HasPreviousState() bool`

HasPreviousState returns a boolean if a field has been set.

### GetPurpose

`func (o *VolumeModel) GetPurpose() BlockStoragePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *VolumeModel) GetPurposeOk() (*BlockStoragePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *VolumeModel) SetPurpose(v BlockStoragePurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *VolumeModel) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSizeGb

`func (o *VolumeModel) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *VolumeModel) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *VolumeModel) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *VolumeModel) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetSnapshotRate

`func (o *VolumeModel) GetSnapshotRate() int32`

GetSnapshotRate returns the SnapshotRate field if non-nil, zero value otherwise.

### GetSnapshotRateOk

`func (o *VolumeModel) GetSnapshotRateOk() (*int32, bool)`

GetSnapshotRateOk returns a tuple with the SnapshotRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRate

`func (o *VolumeModel) SetSnapshotRate(v int32)`

SetSnapshotRate sets SnapshotRate field to given value.

### HasSnapshotRate

`func (o *VolumeModel) HasSnapshotRate() bool`

HasSnapshotRate returns a boolean if a field has been set.

### GetSnapshotSchedule

`func (o *VolumeModel) GetSnapshotSchedule() SnapshotSchedule`

GetSnapshotSchedule returns the SnapshotSchedule field if non-nil, zero value otherwise.

### GetSnapshotScheduleOk

`func (o *VolumeModel) GetSnapshotScheduleOk() (*SnapshotSchedule, bool)`

GetSnapshotScheduleOk returns a tuple with the SnapshotSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedule

`func (o *VolumeModel) SetSnapshotSchedule(v SnapshotSchedule)`

SetSnapshotSchedule sets SnapshotSchedule field to given value.

### HasSnapshotSchedule

`func (o *VolumeModel) HasSnapshotSchedule() bool`

HasSnapshotSchedule returns a boolean if a field has been set.

### GetSrn

`func (o *VolumeModel) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *VolumeModel) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *VolumeModel) SetSrn(v string)`

SetSrn sets Srn field to given value.

### HasSrn

`func (o *VolumeModel) HasSrn() bool`

HasSrn returns a boolean if a field has been set.

### GetState

`func (o *VolumeModel) GetState() BlockStorageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeModel) GetStateOk() (*BlockStorageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeModel) SetState(v BlockStorageState)`

SetState sets State field to given value.

### HasState

`func (o *VolumeModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageVolumeName

`func (o *VolumeModel) GetStorageVolumeName() string`

GetStorageVolumeName returns the StorageVolumeName field if non-nil, zero value otherwise.

### GetStorageVolumeNameOk

`func (o *VolumeModel) GetStorageVolumeNameOk() (*string, bool)`

GetStorageVolumeNameOk returns a tuple with the StorageVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumeName

`func (o *VolumeModel) SetStorageVolumeName(v string)`

SetStorageVolumeName sets StorageVolumeName field to given value.

### HasStorageVolumeName

`func (o *VolumeModel) HasStorageVolumeName() bool`

HasStorageVolumeName returns a boolean if a field has been set.

### GetVolumeGroup

`func (o *VolumeModel) GetVolumeGroup() VolumeGroupInfo`

GetVolumeGroup returns the VolumeGroup field if non-nil, zero value otherwise.

### GetVolumeGroupOk

`func (o *VolumeModel) GetVolumeGroupOk() (*VolumeGroupInfo, bool)`

GetVolumeGroupOk returns a tuple with the VolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroup

`func (o *VolumeModel) SetVolumeGroup(v VolumeGroupInfo)`

SetVolumeGroup sets VolumeGroup field to given value.

### HasVolumeGroup

`func (o *VolumeModel) HasVolumeGroup() bool`

HasVolumeGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


