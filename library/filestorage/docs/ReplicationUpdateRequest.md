# ReplicationUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplicationFrequency** | Pointer to **string** | Frequency | [optional] 
**ReplicationPolicy** | Pointer to **string** | Policy | [optional] 
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

### GetReplicationFrequency

`func (o *ReplicationUpdateRequest) GetReplicationFrequency() string`

GetReplicationFrequency returns the ReplicationFrequency field if non-nil, zero value otherwise.

### GetReplicationFrequencyOk

`func (o *ReplicationUpdateRequest) GetReplicationFrequencyOk() (*string, bool)`

GetReplicationFrequencyOk returns a tuple with the ReplicationFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFrequency

`func (o *ReplicationUpdateRequest) SetReplicationFrequency(v string)`

SetReplicationFrequency sets ReplicationFrequency field to given value.

### HasReplicationFrequency

`func (o *ReplicationUpdateRequest) HasReplicationFrequency() bool`

HasReplicationFrequency returns a boolean if a field has been set.

### GetReplicationPolicy

`func (o *ReplicationUpdateRequest) GetReplicationPolicy() string`

GetReplicationPolicy returns the ReplicationPolicy field if non-nil, zero value otherwise.

### GetReplicationPolicyOk

`func (o *ReplicationUpdateRequest) GetReplicationPolicyOk() (*string, bool)`

GetReplicationPolicyOk returns a tuple with the ReplicationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationPolicy

`func (o *ReplicationUpdateRequest) SetReplicationPolicy(v string)`

SetReplicationPolicy sets ReplicationPolicy field to given value.

### HasReplicationPolicy

`func (o *ReplicationUpdateRequest) HasReplicationPolicy() bool`

HasReplicationPolicy returns a boolean if a field has been set.

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


