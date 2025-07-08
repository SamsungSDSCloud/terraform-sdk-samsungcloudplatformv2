# VolumeReplicationPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyncCycle** | Pointer to [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | [optional] 
**SyncState** | Pointer to **string** | Replication state | [optional] 

## Methods

### NewVolumeReplicationPolicyResponse

`func NewVolumeReplicationPolicyResponse() *VolumeReplicationPolicyResponse`

NewVolumeReplicationPolicyResponse instantiates a new VolumeReplicationPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeReplicationPolicyResponseWithDefaults

`func NewVolumeReplicationPolicyResponseWithDefaults() *VolumeReplicationPolicyResponse`

NewVolumeReplicationPolicyResponseWithDefaults instantiates a new VolumeReplicationPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyncCycle

`func (o *VolumeReplicationPolicyResponse) GetSyncCycle() ReplicationCycle`

GetSyncCycle returns the SyncCycle field if non-nil, zero value otherwise.

### GetSyncCycleOk

`func (o *VolumeReplicationPolicyResponse) GetSyncCycleOk() (*ReplicationCycle, bool)`

GetSyncCycleOk returns a tuple with the SyncCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncCycle

`func (o *VolumeReplicationPolicyResponse) SetSyncCycle(v ReplicationCycle)`

SetSyncCycle sets SyncCycle field to given value.

### HasSyncCycle

`func (o *VolumeReplicationPolicyResponse) HasSyncCycle() bool`

HasSyncCycle returns a boolean if a field has been set.

### GetSyncState

`func (o *VolumeReplicationPolicyResponse) GetSyncState() string`

GetSyncState returns the SyncState field if non-nil, zero value otherwise.

### GetSyncStateOk

`func (o *VolumeReplicationPolicyResponse) GetSyncStateOk() (*string, bool)`

GetSyncStateOk returns a tuple with the SyncState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncState

`func (o *VolumeReplicationPolicyResponse) SetSyncState(v string)`

SetSyncState sets SyncState field to given value.

### HasSyncState

`func (o *VolumeReplicationPolicyResponse) HasSyncState() bool`

HasSyncState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


