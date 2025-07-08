# ReplicationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRetentionCount** | Pointer to **int32** | Backup Retention Count | [optional] 
**CifsPassword** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Target Volume Name | 
**Region** | **string** | Target Region | 
**ReplicationFrequency** | **string** | Frequency | 
**ReplicationType** | **string** | replication | 
**VolumeId** | **string** | Source Volume ID | 

## Methods

### NewReplicationCreateRequest

`func NewReplicationCreateRequest(name string, region string, replicationFrequency string, replicationType string, volumeId string, ) *ReplicationCreateRequest`

NewReplicationCreateRequest instantiates a new ReplicationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationCreateRequestWithDefaults

`func NewReplicationCreateRequestWithDefaults() *ReplicationCreateRequest`

NewReplicationCreateRequestWithDefaults instantiates a new ReplicationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRetentionCount

`func (o *ReplicationCreateRequest) GetBackupRetentionCount() int32`

GetBackupRetentionCount returns the BackupRetentionCount field if non-nil, zero value otherwise.

### GetBackupRetentionCountOk

`func (o *ReplicationCreateRequest) GetBackupRetentionCountOk() (*int32, bool)`

GetBackupRetentionCountOk returns a tuple with the BackupRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionCount

`func (o *ReplicationCreateRequest) SetBackupRetentionCount(v int32)`

SetBackupRetentionCount sets BackupRetentionCount field to given value.

### HasBackupRetentionCount

`func (o *ReplicationCreateRequest) HasBackupRetentionCount() bool`

HasBackupRetentionCount returns a boolean if a field has been set.

### GetCifsPassword

`func (o *ReplicationCreateRequest) GetCifsPassword() string`

GetCifsPassword returns the CifsPassword field if non-nil, zero value otherwise.

### GetCifsPasswordOk

`func (o *ReplicationCreateRequest) GetCifsPasswordOk() (*string, bool)`

GetCifsPasswordOk returns a tuple with the CifsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsPassword

`func (o *ReplicationCreateRequest) SetCifsPassword(v string)`

SetCifsPassword sets CifsPassword field to given value.

### HasCifsPassword

`func (o *ReplicationCreateRequest) HasCifsPassword() bool`

HasCifsPassword returns a boolean if a field has been set.

### SetCifsPasswordNil

`func (o *ReplicationCreateRequest) SetCifsPasswordNil(b bool)`

 SetCifsPasswordNil sets the value for CifsPassword to be an explicit nil

### UnsetCifsPassword
`func (o *ReplicationCreateRequest) UnsetCifsPassword()`

UnsetCifsPassword ensures that no value is present for CifsPassword, not even an explicit nil
### GetName

`func (o *ReplicationCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReplicationCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReplicationCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *ReplicationCreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ReplicationCreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ReplicationCreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetReplicationFrequency

`func (o *ReplicationCreateRequest) GetReplicationFrequency() string`

GetReplicationFrequency returns the ReplicationFrequency field if non-nil, zero value otherwise.

### GetReplicationFrequencyOk

`func (o *ReplicationCreateRequest) GetReplicationFrequencyOk() (*string, bool)`

GetReplicationFrequencyOk returns a tuple with the ReplicationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFrequency

`func (o *ReplicationCreateRequest) SetReplicationFrequency(v string)`

SetReplicationFrequency sets ReplicationFrequency field to given value.


### GetReplicationType

`func (o *ReplicationCreateRequest) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *ReplicationCreateRequest) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *ReplicationCreateRequest) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.


### GetVolumeId

`func (o *ReplicationCreateRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ReplicationCreateRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ReplicationCreateRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


