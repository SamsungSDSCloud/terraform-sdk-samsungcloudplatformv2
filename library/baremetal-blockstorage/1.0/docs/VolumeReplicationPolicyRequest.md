# VolumeReplicationPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Policy** | [**ReplicationPolicyMode**](ReplicationPolicyMode.md) | Policy | 

## Methods

### NewVolumeReplicationPolicyRequest

`func NewVolumeReplicationPolicyRequest(policy ReplicationPolicyMode, ) *VolumeReplicationPolicyRequest`

NewVolumeReplicationPolicyRequest instantiates a new VolumeReplicationPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeReplicationPolicyRequestWithDefaults

`func NewVolumeReplicationPolicyRequestWithDefaults() *VolumeReplicationPolicyRequest`

NewVolumeReplicationPolicyRequestWithDefaults instantiates a new VolumeReplicationPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicy

`func (o *VolumeReplicationPolicyRequest) GetPolicy() ReplicationPolicyMode`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *VolumeReplicationPolicyRequest) GetPolicyOk() (*ReplicationPolicyMode, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *VolumeReplicationPolicyRequest) SetPolicy(v ReplicationPolicyMode)`

SetPolicy sets Policy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


