# VolumeGroupReplicationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cycle** | [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | 
**Name** | **string** | Replication volume group name | 
**Region** | **string** | Region | 
**ReplicationVolumeNamePrefix** | **string** | Replication volume name prefix | 

## Methods

### NewVolumeGroupReplicationCreateRequest

`func NewVolumeGroupReplicationCreateRequest(cycle ReplicationCycle, name string, region string, replicationVolumeNamePrefix string, ) *VolumeGroupReplicationCreateRequest`

NewVolumeGroupReplicationCreateRequest instantiates a new VolumeGroupReplicationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupReplicationCreateRequestWithDefaults

`func NewVolumeGroupReplicationCreateRequestWithDefaults() *VolumeGroupReplicationCreateRequest`

NewVolumeGroupReplicationCreateRequestWithDefaults instantiates a new VolumeGroupReplicationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycle

`func (o *VolumeGroupReplicationCreateRequest) GetCycle() ReplicationCycle`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *VolumeGroupReplicationCreateRequest) GetCycleOk() (*ReplicationCycle, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *VolumeGroupReplicationCreateRequest) SetCycle(v ReplicationCycle)`

SetCycle sets Cycle field to given value.


### GetName

`func (o *VolumeGroupReplicationCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroupReplicationCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroupReplicationCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *VolumeGroupReplicationCreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VolumeGroupReplicationCreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VolumeGroupReplicationCreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetReplicationVolumeNamePrefix

`func (o *VolumeGroupReplicationCreateRequest) GetReplicationVolumeNamePrefix() string`

GetReplicationVolumeNamePrefix returns the ReplicationVolumeNamePrefix field if non-nil, zero value otherwise.

### GetReplicationVolumeNamePrefixOk

`func (o *VolumeGroupReplicationCreateRequest) GetReplicationVolumeNamePrefixOk() (*string, bool)`

GetReplicationVolumeNamePrefixOk returns a tuple with the ReplicationVolumeNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeNamePrefix

`func (o *VolumeGroupReplicationCreateRequest) SetReplicationVolumeNamePrefix(v string)`

SetReplicationVolumeNamePrefix sets ReplicationVolumeNamePrefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


