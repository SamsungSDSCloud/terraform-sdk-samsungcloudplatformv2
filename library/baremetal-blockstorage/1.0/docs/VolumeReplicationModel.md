# VolumeReplicationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelatedVolumes** | Pointer to [**[]RelationVolumeModel**](RelationVolumeModel.md) | List of related volumes | [optional] [default to []]
**SyncCycle** | Pointer to [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | [optional] 
**SyncState** | Pointer to **string** | Replication state | [optional] 

## Methods

### NewVolumeReplicationModel

`func NewVolumeReplicationModel() *VolumeReplicationModel`

NewVolumeReplicationModel instantiates a new VolumeReplicationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeReplicationModelWithDefaults

`func NewVolumeReplicationModelWithDefaults() *VolumeReplicationModel`

NewVolumeReplicationModelWithDefaults instantiates a new VolumeReplicationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelatedVolumes

`func (o *VolumeReplicationModel) GetRelatedVolumes() []RelationVolumeModel`

GetRelatedVolumes returns the RelatedVolumes field if non-nil, zero value otherwise.

### GetRelatedVolumesOk

`func (o *VolumeReplicationModel) GetRelatedVolumesOk() (*[]RelationVolumeModel, bool)`

GetRelatedVolumesOk returns a tuple with the RelatedVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedVolumes

`func (o *VolumeReplicationModel) SetRelatedVolumes(v []RelationVolumeModel)`

SetRelatedVolumes sets RelatedVolumes field to given value.

### HasRelatedVolumes

`func (o *VolumeReplicationModel) HasRelatedVolumes() bool`

HasRelatedVolumes returns a boolean if a field has been set.

### GetSyncCycle

`func (o *VolumeReplicationModel) GetSyncCycle() ReplicationCycle`

GetSyncCycle returns the SyncCycle field if non-nil, zero value otherwise.

### GetSyncCycleOk

`func (o *VolumeReplicationModel) GetSyncCycleOk() (*ReplicationCycle, bool)`

GetSyncCycleOk returns a tuple with the SyncCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCycle

`func (o *VolumeReplicationModel) SetSyncCycle(v ReplicationCycle)`

SetSyncCycle sets SyncCycle field to given value.

### HasSyncCycle

`func (o *VolumeReplicationModel) HasSyncCycle() bool`

HasSyncCycle returns a boolean if a field has been set.

### GetSyncState

`func (o *VolumeReplicationModel) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *VolumeReplicationModel) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *VolumeReplicationModel) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *VolumeReplicationModel) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


