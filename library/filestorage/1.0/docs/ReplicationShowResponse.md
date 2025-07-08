# ReplicationShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRetentionCount** | **NullableInt32** |  | 
**ReplicationFrequency** | **string** | Frequency | 
**ReplicationId** | **string** | Replication ID | 
**ReplicationPolicy** | **string** | Policy | 
**ReplicationStatus** | **string** | Replication Status | 
**ReplicationType** | **string** | Replication Type | 
**ReplicationVolumeAccessLevel** | **string** | Target Access Level | 
**ReplicationVolumeId** | **string** | Target Volume ID | 
**ReplicationVolumeName** | **string** | Target Volume Name | 
**ReplicationVolumeRegion** | **string** | Target Region | 
**SourceVolumeAccessLevel** | **string** | Source Access Level | 
**SourceVolumeId** | **string** | Source Volume ID | 
**SourceVolumeName** | **string** | Source Volume Name | 
**SourceVolumeRegion** | **string** | Source Region | 

## Methods

### NewReplicationShowResponse

`func NewReplicationShowResponse(backupRetentionCount NullableInt32, replicationFrequency string, replicationId string, replicationPolicy string, replicationStatus string, replicationType string, replicationVolumeAccessLevel string, replicationVolumeId string, replicationVolumeName string, replicationVolumeRegion string, sourceVolumeAccessLevel string, sourceVolumeId string, sourceVolumeName string, sourceVolumeRegion string, ) *ReplicationShowResponse`

NewReplicationShowResponse instantiates a new ReplicationShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationShowResponseWithDefaults

`func NewReplicationShowResponseWithDefaults() *ReplicationShowResponse`

NewReplicationShowResponseWithDefaults instantiates a new ReplicationShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRetentionCount

`func (o *ReplicationShowResponse) GetBackupRetentionCount() int32`

GetBackupRetentionCount returns the BackupRetentionCount field if non-nil, zero value otherwise.

### GetBackupRetentionCountOk

`func (o *ReplicationShowResponse) GetBackupRetentionCountOk() (*int32, bool)`

GetBackupRetentionCountOk returns a tuple with the BackupRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionCount

`func (o *ReplicationShowResponse) SetBackupRetentionCount(v int32)`

SetBackupRetentionCount sets BackupRetentionCount field to given value.


### SetBackupRetentionCountNil

`func (o *ReplicationShowResponse) SetBackupRetentionCountNil(b bool)`

 SetBackupRetentionCountNil sets the value for BackupRetentionCount to be an explicit nil

### UnsetBackupRetentionCount
`func (o *ReplicationShowResponse) UnsetBackupRetentionCount()`

UnsetBackupRetentionCount ensures that no value is present for BackupRetentionCount, not even an explicit nil
### GetReplicationFrequency

`func (o *ReplicationShowResponse) GetReplicationFrequency() string`

GetReplicationFrequency returns the ReplicationFrequency field if non-nil, zero value otherwise.

### GetReplicationFrequencyOk

`func (o *ReplicationShowResponse) GetReplicationFrequencyOk() (*string, bool)`

GetReplicationFrequencyOk returns a tuple with the ReplicationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFrequency

`func (o *ReplicationShowResponse) SetReplicationFrequency(v string)`

SetReplicationFrequency sets ReplicationFrequency field to given value.


### GetReplicationId

`func (o *ReplicationShowResponse) GetReplicationId() string`

GetReplicationId returns the ReplicationId field if non-nil, zero value otherwise.

### GetReplicationIdOk

`func (o *ReplicationShowResponse) GetReplicationIdOk() (*string, bool)`

GetReplicationIdOk returns a tuple with the ReplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationId

`func (o *ReplicationShowResponse) SetReplicationId(v string)`

SetReplicationId sets ReplicationId field to given value.


### GetReplicationPolicy

`func (o *ReplicationShowResponse) GetReplicationPolicy() string`

GetReplicationPolicy returns the ReplicationPolicy field if non-nil, zero value otherwise.

### GetReplicationPolicyOk

`func (o *ReplicationShowResponse) GetReplicationPolicyOk() (*string, bool)`

GetReplicationPolicyOk returns a tuple with the ReplicationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPolicy

`func (o *ReplicationShowResponse) SetReplicationPolicy(v string)`

SetReplicationPolicy sets ReplicationPolicy field to given value.


### GetReplicationStatus

`func (o *ReplicationShowResponse) GetReplicationStatus() string`

GetReplicationStatus returns the ReplicationStatus field if non-nil, zero value otherwise.

### GetReplicationStatusOk

`func (o *ReplicationShowResponse) GetReplicationStatusOk() (*string, bool)`

GetReplicationStatusOk returns a tuple with the ReplicationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStatus

`func (o *ReplicationShowResponse) SetReplicationStatus(v string)`

SetReplicationStatus sets ReplicationStatus field to given value.


### GetReplicationType

`func (o *ReplicationShowResponse) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *ReplicationShowResponse) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *ReplicationShowResponse) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.


### GetReplicationVolumeAccessLevel

`func (o *ReplicationShowResponse) GetReplicationVolumeAccessLevel() string`

GetReplicationVolumeAccessLevel returns the ReplicationVolumeAccessLevel field if non-nil, zero value otherwise.

### GetReplicationVolumeAccessLevelOk

`func (o *ReplicationShowResponse) GetReplicationVolumeAccessLevelOk() (*string, bool)`

GetReplicationVolumeAccessLevelOk returns a tuple with the ReplicationVolumeAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeAccessLevel

`func (o *ReplicationShowResponse) SetReplicationVolumeAccessLevel(v string)`

SetReplicationVolumeAccessLevel sets ReplicationVolumeAccessLevel field to given value.


### GetReplicationVolumeId

`func (o *ReplicationShowResponse) GetReplicationVolumeId() string`

GetReplicationVolumeId returns the ReplicationVolumeId field if non-nil, zero value otherwise.

### GetReplicationVolumeIdOk

`func (o *ReplicationShowResponse) GetReplicationVolumeIdOk() (*string, bool)`

GetReplicationVolumeIdOk returns a tuple with the ReplicationVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeId

`func (o *ReplicationShowResponse) SetReplicationVolumeId(v string)`

SetReplicationVolumeId sets ReplicationVolumeId field to given value.


### GetReplicationVolumeName

`func (o *ReplicationShowResponse) GetReplicationVolumeName() string`

GetReplicationVolumeName returns the ReplicationVolumeName field if non-nil, zero value otherwise.

### GetReplicationVolumeNameOk

`func (o *ReplicationShowResponse) GetReplicationVolumeNameOk() (*string, bool)`

GetReplicationVolumeNameOk returns a tuple with the ReplicationVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeName

`func (o *ReplicationShowResponse) SetReplicationVolumeName(v string)`

SetReplicationVolumeName sets ReplicationVolumeName field to given value.


### GetReplicationVolumeRegion

`func (o *ReplicationShowResponse) GetReplicationVolumeRegion() string`

GetReplicationVolumeRegion returns the ReplicationVolumeRegion field if non-nil, zero value otherwise.

### GetReplicationVolumeRegionOk

`func (o *ReplicationShowResponse) GetReplicationVolumeRegionOk() (*string, bool)`

GetReplicationVolumeRegionOk returns a tuple with the ReplicationVolumeRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationVolumeRegion

`func (o *ReplicationShowResponse) SetReplicationVolumeRegion(v string)`

SetReplicationVolumeRegion sets ReplicationVolumeRegion field to given value.


### GetSourceVolumeAccessLevel

`func (o *ReplicationShowResponse) GetSourceVolumeAccessLevel() string`

GetSourceVolumeAccessLevel returns the SourceVolumeAccessLevel field if non-nil, zero value otherwise.

### GetSourceVolumeAccessLevelOk

`func (o *ReplicationShowResponse) GetSourceVolumeAccessLevelOk() (*string, bool)`

GetSourceVolumeAccessLevelOk returns a tuple with the SourceVolumeAccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeAccessLevel

`func (o *ReplicationShowResponse) SetSourceVolumeAccessLevel(v string)`

SetSourceVolumeAccessLevel sets SourceVolumeAccessLevel field to given value.


### GetSourceVolumeId

`func (o *ReplicationShowResponse) GetSourceVolumeId() string`

GetSourceVolumeId returns the SourceVolumeId field if non-nil, zero value otherwise.

### GetSourceVolumeIdOk

`func (o *ReplicationShowResponse) GetSourceVolumeIdOk() (*string, bool)`

GetSourceVolumeIdOk returns a tuple with the SourceVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeId

`func (o *ReplicationShowResponse) SetSourceVolumeId(v string)`

SetSourceVolumeId sets SourceVolumeId field to given value.


### GetSourceVolumeName

`func (o *ReplicationShowResponse) GetSourceVolumeName() string`

GetSourceVolumeName returns the SourceVolumeName field if non-nil, zero value otherwise.

### GetSourceVolumeNameOk

`func (o *ReplicationShowResponse) GetSourceVolumeNameOk() (*string, bool)`

GetSourceVolumeNameOk returns a tuple with the SourceVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeName

`func (o *ReplicationShowResponse) SetSourceVolumeName(v string)`

SetSourceVolumeName sets SourceVolumeName field to given value.


### GetSourceVolumeRegion

`func (o *ReplicationShowResponse) GetSourceVolumeRegion() string`

GetSourceVolumeRegion returns the SourceVolumeRegion field if non-nil, zero value otherwise.

### GetSourceVolumeRegionOk

`func (o *ReplicationShowResponse) GetSourceVolumeRegionOk() (*string, bool)`

GetSourceVolumeRegionOk returns a tuple with the SourceVolumeRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeRegion

`func (o *ReplicationShowResponse) SetSourceVolumeRegion(v string)`

SetSourceVolumeRegion sets SourceVolumeRegion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


