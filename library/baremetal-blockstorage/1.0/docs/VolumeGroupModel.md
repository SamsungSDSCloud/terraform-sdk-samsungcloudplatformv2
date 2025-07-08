# VolumeGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | Account id | [optional] 
**CreatedAt** | Pointer to **string** | Created at | [optional] 
**CreatedBy** | Pointer to **string** | Created by | [optional] 
**HasOrigin** | Pointer to **bool** | Whether origin relation exists | [optional] 
**HasReplication** | Pointer to **bool** | Whether replication relation exists | [optional] 
**Id** | Pointer to **string** | Volume group id | [optional] 
**MemberVolumes** | Pointer to [**[]MemberVolume**](MemberVolume.md) | List of member volumes | [optional] 
**ModifiedAt** | Pointer to **string** | Modified at | [optional] 
**ModifiedBy** | Pointer to **string** | Modified by | [optional] 
**Name** | Pointer to **string** | Volume group name | [optional] 
**Purpose** | Pointer to [**VolumeGroupPurpose**](VolumeGroupPurpose.md) | Volume group purpose | [optional] 
**SnapshotSchedule** | Pointer to [**SnapshotSchedule**](SnapshotSchedule.md) | Snapshot schedule | [optional] 
**SoStorageId** | Pointer to **int32** | SO Storage ID | [optional] 
**SoVolumePoolId** | Pointer to **string** | SO Volume Pool ID | [optional] 
**State** | Pointer to [**VolumeGroupState**](VolumeGroupState.md) | Current state | [optional] 
**StorageVolumeGroupName** | Pointer to **string** | Volume group name | [optional] 

## Methods

### NewVolumeGroupModel

`func NewVolumeGroupModel() *VolumeGroupModel`

NewVolumeGroupModel instantiates a new VolumeGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupModelWithDefaults

`func NewVolumeGroupModelWithDefaults() *VolumeGroupModel`

NewVolumeGroupModelWithDefaults instantiates a new VolumeGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VolumeGroupModel) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VolumeGroupModel) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VolumeGroupModel) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *VolumeGroupModel) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumeGroupModel) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeGroupModel) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeGroupModel) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeGroupModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VolumeGroupModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VolumeGroupModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VolumeGroupModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VolumeGroupModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetHasOrigin

`func (o *VolumeGroupModel) GetHasOrigin() bool`

GetHasOrigin returns the HasOrigin field if non-nil, zero value otherwise.

### GetHasOriginOk

`func (o *VolumeGroupModel) GetHasOriginOk() (*bool, bool)`

GetHasOriginOk returns a tuple with the HasOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOrigin

`func (o *VolumeGroupModel) SetHasOrigin(v bool)`

SetHasOrigin sets HasOrigin field to given value.

### HasHasOrigin

`func (o *VolumeGroupModel) HasHasOrigin() bool`

HasHasOrigin returns a boolean if a field has been set.

### GetHasReplication

`func (o *VolumeGroupModel) GetHasReplication() bool`

GetHasReplication returns the HasReplication field if non-nil, zero value otherwise.

### GetHasReplicationOk

`func (o *VolumeGroupModel) GetHasReplicationOk() (*bool, bool)`

GetHasReplicationOk returns a tuple with the HasReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReplication

`func (o *VolumeGroupModel) SetHasReplication(v bool)`

SetHasReplication sets HasReplication field to given value.

### HasHasReplication

`func (o *VolumeGroupModel) HasHasReplication() bool`

HasHasReplication returns a boolean if a field has been set.

### GetId

`func (o *VolumeGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeGroupModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeGroupModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemberVolumes

`func (o *VolumeGroupModel) GetMemberVolumes() []MemberVolume`

GetMemberVolumes returns the MemberVolumes field if non-nil, zero value otherwise.

### GetMemberVolumesOk

`func (o *VolumeGroupModel) GetMemberVolumesOk() (*[]MemberVolume, bool)`

GetMemberVolumesOk returns a tuple with the MemberVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberVolumes

`func (o *VolumeGroupModel) SetMemberVolumes(v []MemberVolume)`

SetMemberVolumes sets MemberVolumes field to given value.

### HasMemberVolumes

`func (o *VolumeGroupModel) HasMemberVolumes() bool`

HasMemberVolumes returns a boolean if a field has been set.

### GetModifiedAt

`func (o *VolumeGroupModel) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VolumeGroupModel) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VolumeGroupModel) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *VolumeGroupModel) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *VolumeGroupModel) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VolumeGroupModel) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VolumeGroupModel) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *VolumeGroupModel) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *VolumeGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurpose

`func (o *VolumeGroupModel) GetPurpose() VolumeGroupPurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *VolumeGroupModel) GetPurposeOk() (*VolumeGroupPurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *VolumeGroupModel) SetPurpose(v VolumeGroupPurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *VolumeGroupModel) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSnapshotSchedule

`func (o *VolumeGroupModel) GetSnapshotSchedule() SnapshotSchedule`

GetSnapshotSchedule returns the SnapshotSchedule field if non-nil, zero value otherwise.

### GetSnapshotScheduleOk

`func (o *VolumeGroupModel) GetSnapshotScheduleOk() (*SnapshotSchedule, bool)`

GetSnapshotScheduleOk returns a tuple with the SnapshotSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSchedule

`func (o *VolumeGroupModel) SetSnapshotSchedule(v SnapshotSchedule)`

SetSnapshotSchedule sets SnapshotSchedule field to given value.

### HasSnapshotSchedule

`func (o *VolumeGroupModel) HasSnapshotSchedule() bool`

HasSnapshotSchedule returns a boolean if a field has been set.

### GetSoStorageId

`func (o *VolumeGroupModel) GetSoStorageId() int32`

GetSoStorageId returns the SoStorageId field if non-nil, zero value otherwise.

### GetSoStorageIdOk

`func (o *VolumeGroupModel) GetSoStorageIdOk() (*int32, bool)`

GetSoStorageIdOk returns a tuple with the SoStorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoStorageId

`func (o *VolumeGroupModel) SetSoStorageId(v int32)`

SetSoStorageId sets SoStorageId field to given value.

### HasSoStorageId

`func (o *VolumeGroupModel) HasSoStorageId() bool`

HasSoStorageId returns a boolean if a field has been set.

### GetSoVolumePoolId

`func (o *VolumeGroupModel) GetSoVolumePoolId() string`

GetSoVolumePoolId returns the SoVolumePoolId field if non-nil, zero value otherwise.

### GetSoVolumePoolIdOk

`func (o *VolumeGroupModel) GetSoVolumePoolIdOk() (*string, bool)`

GetSoVolumePoolIdOk returns a tuple with the SoVolumePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoVolumePoolId

`func (o *VolumeGroupModel) SetSoVolumePoolId(v string)`

SetSoVolumePoolId sets SoVolumePoolId field to given value.

### HasSoVolumePoolId

`func (o *VolumeGroupModel) HasSoVolumePoolId() bool`

HasSoVolumePoolId returns a boolean if a field has been set.

### GetState

`func (o *VolumeGroupModel) GetState() VolumeGroupState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeGroupModel) GetStateOk() (*VolumeGroupState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeGroupModel) SetState(v VolumeGroupState)`

SetState sets State field to given value.

### HasState

`func (o *VolumeGroupModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageVolumeGroupName

`func (o *VolumeGroupModel) GetStorageVolumeGroupName() string`

GetStorageVolumeGroupName returns the StorageVolumeGroupName field if non-nil, zero value otherwise.

### GetStorageVolumeGroupNameOk

`func (o *VolumeGroupModel) GetStorageVolumeGroupNameOk() (*string, bool)`

GetStorageVolumeGroupNameOk returns a tuple with the StorageVolumeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumeGroupName

`func (o *VolumeGroupModel) SetStorageVolumeGroupName(v string)`

SetStorageVolumeGroupName sets StorageVolumeGroupName field to given value.

### HasStorageVolumeGroupName

`func (o *VolumeGroupModel) HasStorageVolumeGroupName() bool`

HasStorageVolumeGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


