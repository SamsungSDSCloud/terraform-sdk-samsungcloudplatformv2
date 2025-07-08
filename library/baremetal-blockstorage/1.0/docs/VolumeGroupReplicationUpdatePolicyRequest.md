# VolumeGroupReplicationUpdatePolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**ReplicationPolicyMode**](ReplicationPolicyMode.md) | Policy | 

## Methods

### NewVolumeGroupReplicationUpdatePolicyRequest

`func NewVolumeGroupReplicationUpdatePolicyRequest(policy ReplicationPolicyMode, ) *VolumeGroupReplicationUpdatePolicyRequest`

NewVolumeGroupReplicationUpdatePolicyRequest instantiates a new VolumeGroupReplicationUpdatePolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupReplicationUpdatePolicyRequestWithDefaults

`func NewVolumeGroupReplicationUpdatePolicyRequestWithDefaults() *VolumeGroupReplicationUpdatePolicyRequest`

NewVolumeGroupReplicationUpdatePolicyRequestWithDefaults instantiates a new VolumeGroupReplicationUpdatePolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VolumeGroupReplicationUpdatePolicyRequest) GetPolicy() ReplicationPolicyMode`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VolumeGroupReplicationUpdatePolicyRequest) GetPolicyOk() (*ReplicationPolicyMode, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VolumeGroupReplicationUpdatePolicyRequest) SetPolicy(v ReplicationPolicyMode)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


