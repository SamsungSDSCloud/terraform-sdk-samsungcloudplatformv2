# VolumeListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **[]map[string]interface{}** | Attached server | [optional] 
**CreatedAt** | Pointer to **time.Time** | Created at | [optional] 
**CreatedBy** | Pointer to **string** | Created by | [optional] 
**DiskType** | Pointer to [**DiskType**](DiskType.md) | Disk type | [optional] 
**HasRelation** | Pointer to **bool** | Whether relation exists | [optional] 
**Id** | Pointer to **string** | Id | [optional] 
**IsInVolumeGroup** | Pointer to **bool** | Whether to belong to the volume group | [optional] 
**IsSnapshotActivated** | Pointer to **bool** | Whether to activate snapshot | [optional] 
**IscsiTargetIps** | Pointer to **[]string** | List of iscsi ips | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Modified at | [optional] 
**ModifiedBy** | Pointer to **string** | Modified by | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Purpose** | Pointer to [**BlockStoragePurpose**](BlockStoragePurpose.md) | Purpose | [optional] 
**SizeGb** | Pointer to **int32** | Size gb | [optional] 
**State** | Pointer to [**BlockStorageState**](BlockStorageState.md) | State | [optional] 
**VolumeGroup** | Pointer to **map[string]interface{}** | Volume group | [optional] 

## Methods

### NewVolumeListModel

`func NewVolumeListModel() *VolumeListModel`

NewVolumeListModel instantiates a new VolumeListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeListModelWithDefaults

`func NewVolumeListModelWithDefaults() *VolumeListModel`

NewVolumeListModelWithDefaults instantiates a new VolumeListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *VolumeListModel) GetAttachments() []map[string]interface{}`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumeListModel) GetAttachmentsOk() (*[]map[string]interface{}, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumeListModel) SetAttachments(v []map[string]interface{})`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *VolumeListModel) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VolumeListModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeListModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeListModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeListModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VolumeListModel) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VolumeListModel) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VolumeListModel) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VolumeListModel) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDiskType

`func (o *VolumeListModel) GetDiskType() DiskType`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *VolumeListModel) GetDiskTypeOk() (*DiskType, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *VolumeListModel) SetDiskType(v DiskType)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *VolumeListModel) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetHasRelation

`func (o *VolumeListModel) GetHasRelation() bool`

GetHasRelation returns the HasRelation field if non-nil, zero value otherwise.

### GetHasRelationOk

`func (o *VolumeListModel) GetHasRelationOk() (*bool, bool)`

GetHasRelationOk returns a tuple with the HasRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRelation

`func (o *VolumeListModel) SetHasRelation(v bool)`

SetHasRelation sets HasRelation field to given value.

### HasHasRelation

`func (o *VolumeListModel) HasHasRelation() bool`

HasHasRelation returns a boolean if a field has been set.

### GetId

`func (o *VolumeListModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeListModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeListModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeListModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsInVolumeGroup

`func (o *VolumeListModel) GetIsInVolumeGroup() bool`

GetIsInVolumeGroup returns the IsInVolumeGroup field if non-nil, zero value otherwise.

### GetIsInVolumeGroupOk

`func (o *VolumeListModel) GetIsInVolumeGroupOk() (*bool, bool)`

GetIsInVolumeGroupOk returns a tuple with the IsInVolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInVolumeGroup

`func (o *VolumeListModel) SetIsInVolumeGroup(v bool)`

SetIsInVolumeGroup sets IsInVolumeGroup field to given value.

### HasIsInVolumeGroup

`func (o *VolumeListModel) HasIsInVolumeGroup() bool`

HasIsInVolumeGroup returns a boolean if a field has been set.

### GetIsSnapshotActivated

`func (o *VolumeListModel) GetIsSnapshotActivated() bool`

GetIsSnapshotActivated returns the IsSnapshotActivated field if non-nil, zero value otherwise.

### GetIsSnapshotActivatedOk

`func (o *VolumeListModel) GetIsSnapshotActivatedOk() (*bool, bool)`

GetIsSnapshotActivatedOk returns a tuple with the IsSnapshotActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotActivated

`func (o *VolumeListModel) SetIsSnapshotActivated(v bool)`

SetIsSnapshotActivated sets IsSnapshotActivated field to given value.

### HasIsSnapshotActivated

`func (o *VolumeListModel) HasIsSnapshotActivated() bool`

HasIsSnapshotActivated returns a boolean if a field has been set.

### GetIscsiTargetIps

`func (o *VolumeListModel) GetIscsiTargetIps() []string`

GetIscsiTargetIps returns the IscsiTargetIps field if non-nil, zero value otherwise.

### GetIscsiTargetIpsOk

`func (o *VolumeListModel) GetIscsiTargetIpsOk() (*[]string, bool)`

GetIscsiTargetIpsOk returns a tuple with the IscsiTargetIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIscsiTargetIps

`func (o *VolumeListModel) SetIscsiTargetIps(v []string)`

SetIscsiTargetIps sets IscsiTargetIps field to given value.

### HasIscsiTargetIps

`func (o *VolumeListModel) HasIscsiTargetIps() bool`

HasIscsiTargetIps returns a boolean if a field has been set.

### GetModifiedAt

`func (o *VolumeListModel) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VolumeListModel) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VolumeListModel) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *VolumeListModel) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetModifiedBy

`func (o *VolumeListModel) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VolumeListModel) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VolumeListModel) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *VolumeListModel) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetName

`func (o *VolumeListModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeListModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeListModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeListModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurpose

`func (o *VolumeListModel) GetPurpose() BlockStoragePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *VolumeListModel) GetPurposeOk() (*BlockStoragePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *VolumeListModel) SetPurpose(v BlockStoragePurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *VolumeListModel) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSizeGb

`func (o *VolumeListModel) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *VolumeListModel) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *VolumeListModel) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *VolumeListModel) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetState

`func (o *VolumeListModel) GetState() BlockStorageState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeListModel) GetStateOk() (*BlockStorageState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeListModel) SetState(v BlockStorageState)`

SetState sets State field to given value.

### HasState

`func (o *VolumeListModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVolumeGroup

`func (o *VolumeListModel) GetVolumeGroup() map[string]interface{}`

GetVolumeGroup returns the VolumeGroup field if non-nil, zero value otherwise.

### GetVolumeGroupOk

`func (o *VolumeListModel) GetVolumeGroupOk() (*map[string]interface{}, bool)`

GetVolumeGroupOk returns a tuple with the VolumeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroup

`func (o *VolumeListModel) SetVolumeGroup(v map[string]interface{})`

SetVolumeGroup sets VolumeGroup field to given value.

### HasVolumeGroup

`func (o *VolumeListModel) HasVolumeGroup() bool`

HasVolumeGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


