# VolumeGroupReplicationUpdateCycleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cycle** | Pointer to [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | [optional] 
**VolumeGroupId** | Pointer to **string** | Volume group id | [optional] 

## Methods

### NewVolumeGroupReplicationUpdateCycleResponse

`func NewVolumeGroupReplicationUpdateCycleResponse() *VolumeGroupReplicationUpdateCycleResponse`

NewVolumeGroupReplicationUpdateCycleResponse instantiates a new VolumeGroupReplicationUpdateCycleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupReplicationUpdateCycleResponseWithDefaults

`func NewVolumeGroupReplicationUpdateCycleResponseWithDefaults() *VolumeGroupReplicationUpdateCycleResponse`

NewVolumeGroupReplicationUpdateCycleResponseWithDefaults instantiates a new VolumeGroupReplicationUpdateCycleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycle

`func (o *VolumeGroupReplicationUpdateCycleResponse) GetCycle() ReplicationCycle`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *VolumeGroupReplicationUpdateCycleResponse) GetCycleOk() (*ReplicationCycle, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *VolumeGroupReplicationUpdateCycleResponse) SetCycle(v ReplicationCycle)`

SetCycle sets Cycle field to given value.

### HasCycle

`func (o *VolumeGroupReplicationUpdateCycleResponse) HasCycle() bool`

HasCycle returns a boolean if a field has been set.

### GetVolumeGroupId

`func (o *VolumeGroupReplicationUpdateCycleResponse) GetVolumeGroupId() string`

GetVolumeGroupId returns the VolumeGroupId field if non-nil, zero value otherwise.

### GetVolumeGroupIdOk

`func (o *VolumeGroupReplicationUpdateCycleResponse) GetVolumeGroupIdOk() (*string, bool)`

GetVolumeGroupIdOk returns a tuple with the VolumeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeGroupId

`func (o *VolumeGroupReplicationUpdateCycleResponse) SetVolumeGroupId(v string)`

SetVolumeGroupId sets VolumeGroupId field to given value.

### HasVolumeGroupId

`func (o *VolumeGroupReplicationUpdateCycleResponse) HasVolumeGroupId() bool`

HasVolumeGroupId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


