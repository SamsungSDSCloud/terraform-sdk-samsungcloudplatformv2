# VolumeGroupReplicationUpdatePolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cycle** | Pointer to [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | [optional] 
**SyncState** | Pointer to **string** | Replication state | [optional] 

## Methods

### NewVolumeGroupReplicationUpdatePolicyResponse

`func NewVolumeGroupReplicationUpdatePolicyResponse() *VolumeGroupReplicationUpdatePolicyResponse`

NewVolumeGroupReplicationUpdatePolicyResponse instantiates a new VolumeGroupReplicationUpdatePolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupReplicationUpdatePolicyResponseWithDefaults

`func NewVolumeGroupReplicationUpdatePolicyResponseWithDefaults() *VolumeGroupReplicationUpdatePolicyResponse`

NewVolumeGroupReplicationUpdatePolicyResponseWithDefaults instantiates a new VolumeGroupReplicationUpdatePolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycle

`func (o *VolumeGroupReplicationUpdatePolicyResponse) GetCycle() ReplicationCycle`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *VolumeGroupReplicationUpdatePolicyResponse) GetCycleOk() (*ReplicationCycle, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *VolumeGroupReplicationUpdatePolicyResponse) SetCycle(v ReplicationCycle)`

SetCycle sets Cycle field to given value.

### HasCycle

`func (o *VolumeGroupReplicationUpdatePolicyResponse) HasCycle() bool`

HasCycle returns a boolean if a field has been set.

### GetSyncState

`func (o *VolumeGroupReplicationUpdatePolicyResponse) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *VolumeGroupReplicationUpdatePolicyResponse) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *VolumeGroupReplicationUpdatePolicyResponse) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *VolumeGroupReplicationUpdatePolicyResponse) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


