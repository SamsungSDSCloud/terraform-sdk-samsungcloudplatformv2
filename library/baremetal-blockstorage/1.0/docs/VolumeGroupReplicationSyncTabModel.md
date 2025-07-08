# VolumeGroupReplicationSyncTabModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelatedVolumeGroups** | Pointer to [**[]RelationVolumeGroupModel**](RelationVolumeGroupModel.md) | List of related volume groups | [optional] [default to []]
**SyncCycle** | Pointer to [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | [optional] 
**SyncState** | Pointer to **string** | Replication state | [optional] 

## Methods

### NewVolumeGroupReplicationSyncTabModel

`func NewVolumeGroupReplicationSyncTabModel() *VolumeGroupReplicationSyncTabModel`

NewVolumeGroupReplicationSyncTabModel instantiates a new VolumeGroupReplicationSyncTabModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupReplicationSyncTabModelWithDefaults

`func NewVolumeGroupReplicationSyncTabModelWithDefaults() *VolumeGroupReplicationSyncTabModel`

NewVolumeGroupReplicationSyncTabModelWithDefaults instantiates a new VolumeGroupReplicationSyncTabModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelatedVolumeGroups

`func (o *VolumeGroupReplicationSyncTabModel) GetRelatedVolumeGroups() []RelationVolumeGroupModel`

GetRelatedVolumeGroups returns the RelatedVolumeGroups field if non-nil, zero value otherwise.

### GetRelatedVolumeGroupsOk

`func (o *VolumeGroupReplicationSyncTabModel) GetRelatedVolumeGroupsOk() (*[]RelationVolumeGroupModel, bool)`

GetRelatedVolumeGroupsOk returns a tuple with the RelatedVolumeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedVolumeGroups

`func (o *VolumeGroupReplicationSyncTabModel) SetRelatedVolumeGroups(v []RelationVolumeGroupModel)`

SetRelatedVolumeGroups sets RelatedVolumeGroups field to given value.

### HasRelatedVolumeGroups

`func (o *VolumeGroupReplicationSyncTabModel) HasRelatedVolumeGroups() bool`

HasRelatedVolumeGroups returns a boolean if a field has been set.

### GetSyncCycle

`func (o *VolumeGroupReplicationSyncTabModel) GetSyncCycle() ReplicationCycle`

GetSyncCycle returns the SyncCycle field if non-nil, zero value otherwise.

### GetSyncCycleOk

`func (o *VolumeGroupReplicationSyncTabModel) GetSyncCycleOk() (*ReplicationCycle, bool)`

GetSyncCycleOk returns a tuple with the SyncCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCycle

`func (o *VolumeGroupReplicationSyncTabModel) SetSyncCycle(v ReplicationCycle)`

SetSyncCycle sets SyncCycle field to given value.

### HasSyncCycle

`func (o *VolumeGroupReplicationSyncTabModel) HasSyncCycle() bool`

HasSyncCycle returns a boolean if a field has been set.

### GetSyncState

`func (o *VolumeGroupReplicationSyncTabModel) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *VolumeGroupReplicationSyncTabModel) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *VolumeGroupReplicationSyncTabModel) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *VolumeGroupReplicationSyncTabModel) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


