# ReplicasCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**MaintenanceOption** | Pointer to [**NullableMaintenanceOption**](MaintenanceOption.md) |  | [optional] 
**ReplicaCount** | **int32** | Replica count | 
**Replicas** | [**[]Replicas**](Replicas.md) | Replicas list | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewReplicasCreateRequest

`func NewReplicasCreateRequest(replicaCount int32, replicas []Replicas, ) *ReplicasCreateRequest`

NewReplicasCreateRequest instantiates a new ReplicasCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicasCreateRequestWithDefaults

`func NewReplicasCreateRequestWithDefaults() *ReplicasCreateRequest`

NewReplicasCreateRequestWithDefaults instantiates a new ReplicasCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowableIpAddresses

`func (o *ReplicasCreateRequest) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *ReplicasCreateRequest) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *ReplicasCreateRequest) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *ReplicasCreateRequest) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetMaintenanceOption

`func (o *ReplicasCreateRequest) GetMaintenanceOption() MaintenanceOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *ReplicasCreateRequest) GetMaintenanceOptionOk() (*MaintenanceOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *ReplicasCreateRequest) SetMaintenanceOption(v MaintenanceOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *ReplicasCreateRequest) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *ReplicasCreateRequest) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *ReplicasCreateRequest) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetReplicaCount

`func (o *ReplicasCreateRequest) GetReplicaCount() int32`

GetReplicaCount returns the ReplicaCount field if non-nil, zero value otherwise.

### GetReplicaCountOk

`func (o *ReplicasCreateRequest) GetReplicaCountOk() (*int32, bool)`

GetReplicaCountOk returns a tuple with the ReplicaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaCount

`func (o *ReplicasCreateRequest) SetReplicaCount(v int32)`

SetReplicaCount sets ReplicaCount field to given value.


### GetReplicas

`func (o *ReplicasCreateRequest) GetReplicas() []Replicas`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *ReplicasCreateRequest) GetReplicasOk() (*[]Replicas, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *ReplicasCreateRequest) SetReplicas(v []Replicas)`

SetReplicas sets Replicas field to given value.


### GetTags

`func (o *ReplicasCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReplicasCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReplicasCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReplicasCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ReplicasCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ReplicasCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


