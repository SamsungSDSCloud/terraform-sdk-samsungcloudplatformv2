# VolumeReplicationCycleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cycle** | Pointer to [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | [optional] 
**VolumeId** | Pointer to **string** | Volume id | [optional] 

## Methods

### NewVolumeReplicationCycleResponse

`func NewVolumeReplicationCycleResponse() *VolumeReplicationCycleResponse`

NewVolumeReplicationCycleResponse instantiates a new VolumeReplicationCycleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeReplicationCycleResponseWithDefaults

`func NewVolumeReplicationCycleResponseWithDefaults() *VolumeReplicationCycleResponse`

NewVolumeReplicationCycleResponseWithDefaults instantiates a new VolumeReplicationCycleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycle

`func (o *VolumeReplicationCycleResponse) GetCycle() ReplicationCycle`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *VolumeReplicationCycleResponse) GetCycleOk() (*ReplicationCycle, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *VolumeReplicationCycleResponse) SetCycle(v ReplicationCycle)`

SetCycle sets Cycle field to given value.

### HasCycle

`func (o *VolumeReplicationCycleResponse) HasCycle() bool`

HasCycle returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeReplicationCycleResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeReplicationCycleResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeReplicationCycleResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeReplicationCycleResponse) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


