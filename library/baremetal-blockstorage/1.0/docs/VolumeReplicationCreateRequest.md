# VolumeReplicationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cycle** | [**ReplicationCycle**](ReplicationCycle.md) | Replication cycle | 
**Name** | **string** | Replication volume name | 
**Region** | **string** | Region | 

## Methods

### NewVolumeReplicationCreateRequest

`func NewVolumeReplicationCreateRequest(cycle ReplicationCycle, name string, region string, ) *VolumeReplicationCreateRequest`

NewVolumeReplicationCreateRequest instantiates a new VolumeReplicationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeReplicationCreateRequestWithDefaults

`func NewVolumeReplicationCreateRequestWithDefaults() *VolumeReplicationCreateRequest`

NewVolumeReplicationCreateRequestWithDefaults instantiates a new VolumeReplicationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycle

`func (o *VolumeReplicationCreateRequest) GetCycle() ReplicationCycle`

GetCycle returns the Cycle field if non-nil, zero value otherwise.

### GetCycleOk

`func (o *VolumeReplicationCreateRequest) GetCycleOk() (*ReplicationCycle, bool)`

GetCycleOk returns a tuple with the Cycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycle

`func (o *VolumeReplicationCreateRequest) SetCycle(v ReplicationCycle)`

SetCycle sets Cycle field to given value.


### GetName

`func (o *VolumeReplicationCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeReplicationCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeReplicationCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *VolumeReplicationCreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VolumeReplicationCreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VolumeReplicationCreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


