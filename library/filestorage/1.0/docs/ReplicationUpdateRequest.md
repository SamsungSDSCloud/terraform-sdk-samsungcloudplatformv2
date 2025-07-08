# ReplicationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupRetentionCount** | Pointer to **NullableInt32** |  | [optional] 
**ReplicationFrequency** | Pointer to [**NullableReplicationUpdatePolicyEnum**](ReplicationUpdatePolicyEnum.md) |  | [optional] 
**ReplicationPolicy** | Pointer to [**NullableReplicationUpdateStatusEnum**](ReplicationUpdateStatusEnum.md) |  | [optional] 
**ReplicationUpdateType** | **string** | Replication Update Type | 

## Methods

### NewReplicationUpdateRequest

`func NewReplicationUpdateRequest(replicationUpdateType string, ) *ReplicationUpdateRequest`

NewReplicationUpdateRequest instantiates a new ReplicationUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationUpdateRequestWithDefaults

`func NewReplicationUpdateRequestWithDefaults() *ReplicationUpdateRequest`

NewReplicationUpdateRequestWithDefaults instantiates a new ReplicationUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupRetentionCount

`func (o *ReplicationUpdateRequest) GetBackupRetentionCount() int32`

GetBackupRetentionCount returns the BackupRetentionCount field if non-nil, zero value otherwise.

### GetBackupRetentionCountOk

`func (o *ReplicationUpdateRequest) GetBackupRetentionCountOk() (*int32, bool)`

GetBackupRetentionCountOk returns a tuple with the BackupRetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRetentionCount

`func (o *ReplicationUpdateRequest) SetBackupRetentionCount(v int32)`

SetBackupRetentionCount sets BackupRetentionCount field to given value.

### HasBackupRetentionCount

`func (o *ReplicationUpdateRequest) HasBackupRetentionCount() bool`

HasBackupRetentionCount returns a boolean if a field has been set.

### SetBackupRetentionCountNil

`func (o *ReplicationUpdateRequest) SetBackupRetentionCountNil(b bool)`

 SetBackupRetentionCountNil sets the value for BackupRetentionCount to be an explicit nil

### UnsetBackupRetentionCount
`func (o *ReplicationUpdateRequest) UnsetBackupRetentionCount()`

UnsetBackupRetentionCount ensures that no value is present for BackupRetentionCount, not even an explicit nil
### GetReplicationFrequency

`func (o *ReplicationUpdateRequest) GetReplicationFrequency() ReplicationUpdatePolicyEnum`

GetReplicationFrequency returns the ReplicationFrequency field if non-nil, zero value otherwise.

### GetReplicationFrequencyOk

`func (o *ReplicationUpdateRequest) GetReplicationFrequencyOk() (*ReplicationUpdatePolicyEnum, bool)`

GetReplicationFrequencyOk returns a tuple with the ReplicationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFrequency

`func (o *ReplicationUpdateRequest) SetReplicationFrequency(v ReplicationUpdatePolicyEnum)`

SetReplicationFrequency sets ReplicationFrequency field to given value.

### HasReplicationFrequency

`func (o *ReplicationUpdateRequest) HasReplicationFrequency() bool`

HasReplicationFrequency returns a boolean if a field has been set.

### SetReplicationFrequencyNil

`func (o *ReplicationUpdateRequest) SetReplicationFrequencyNil(b bool)`

 SetReplicationFrequencyNil sets the value for ReplicationFrequency to be an explicit nil

### UnsetReplicationFrequency
`func (o *ReplicationUpdateRequest) UnsetReplicationFrequency()`

UnsetReplicationFrequency ensures that no value is present for ReplicationFrequency, not even an explicit nil
### GetReplicationPolicy

`func (o *ReplicationUpdateRequest) GetReplicationPolicy() ReplicationUpdateStatusEnum`

GetReplicationPolicy returns the ReplicationPolicy field if non-nil, zero value otherwise.

### GetReplicationPolicyOk

`func (o *ReplicationUpdateRequest) GetReplicationPolicyOk() (*ReplicationUpdateStatusEnum, bool)`

GetReplicationPolicyOk returns a tuple with the ReplicationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPolicy

`func (o *ReplicationUpdateRequest) SetReplicationPolicy(v ReplicationUpdateStatusEnum)`

SetReplicationPolicy sets ReplicationPolicy field to given value.

### HasReplicationPolicy

`func (o *ReplicationUpdateRequest) HasReplicationPolicy() bool`

HasReplicationPolicy returns a boolean if a field has been set.

### SetReplicationPolicyNil

`func (o *ReplicationUpdateRequest) SetReplicationPolicyNil(b bool)`

 SetReplicationPolicyNil sets the value for ReplicationPolicy to be an explicit nil

### UnsetReplicationPolicy
`func (o *ReplicationUpdateRequest) UnsetReplicationPolicy()`

UnsetReplicationPolicy ensures that no value is present for ReplicationPolicy, not even an explicit nil
### GetReplicationUpdateType

`func (o *ReplicationUpdateRequest) GetReplicationUpdateType() string`

GetReplicationUpdateType returns the ReplicationUpdateType field if non-nil, zero value otherwise.

### GetReplicationUpdateTypeOk

`func (o *ReplicationUpdateRequest) GetReplicationUpdateTypeOk() (*string, bool)`

GetReplicationUpdateTypeOk returns a tuple with the ReplicationUpdateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationUpdateType

`func (o *ReplicationUpdateRequest) SetReplicationUpdateType(v string)`

SetReplicationUpdateType sets ReplicationUpdateType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


