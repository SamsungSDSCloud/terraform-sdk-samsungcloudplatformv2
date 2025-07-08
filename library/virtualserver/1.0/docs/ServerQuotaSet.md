# ServerQuotaSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cores** | [**ServerQuotaSetObject**](ServerQuotaSetObject.md) | Core quota | 
**Id** | **string** | Quota set ID | 
**Instances** | [**ServerQuotaSetObject**](ServerQuotaSetObject.md) | Instance quota | 
**Keypairs** | [**ServerQuotaSetObject**](ServerQuotaSetObject.md) | Keypairs quota | 
**MetadataItems** | [**ServerQuotaSetObject**](ServerQuotaSetObject.md) | Metadata items quota | 
**Ram** | [**ServerQuotaSetObject**](ServerQuotaSetObject.md) | RAM quota | 
**ServerGroupMembers** | [**ServerQuotaSetObject**](ServerQuotaSetObject.md) | Server group members quota | 
**ServerGroups** | [**ServerQuotaSetObject**](ServerQuotaSetObject.md) | Server groups quota | 

## Methods

### NewServerQuotaSet

`func NewServerQuotaSet(cores ServerQuotaSetObject, id string, instances ServerQuotaSetObject, keypairs ServerQuotaSetObject, metadataItems ServerQuotaSetObject, ram ServerQuotaSetObject, serverGroupMembers ServerQuotaSetObject, serverGroups ServerQuotaSetObject, ) *ServerQuotaSet`

NewServerQuotaSet instantiates a new ServerQuotaSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerQuotaSetWithDefaults

`func NewServerQuotaSetWithDefaults() *ServerQuotaSet`

NewServerQuotaSetWithDefaults instantiates a new ServerQuotaSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCores

`func (o *ServerQuotaSet) GetCores() ServerQuotaSetObject`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *ServerQuotaSet) GetCoresOk() (*ServerQuotaSetObject, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *ServerQuotaSet) SetCores(v ServerQuotaSetObject)`

SetCores sets Cores field to given value.


### GetId

`func (o *ServerQuotaSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerQuotaSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerQuotaSet) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *ServerQuotaSet) GetInstances() ServerQuotaSetObject`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ServerQuotaSet) GetInstancesOk() (*ServerQuotaSetObject, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ServerQuotaSet) SetInstances(v ServerQuotaSetObject)`

SetInstances sets Instances field to given value.


### GetKeypairs

`func (o *ServerQuotaSet) GetKeypairs() ServerQuotaSetObject`

GetKeypairs returns the Keypairs field if non-nil, zero value otherwise.

### GetKeypairsOk

`func (o *ServerQuotaSet) GetKeypairsOk() (*ServerQuotaSetObject, bool)`

GetKeypairsOk returns a tuple with the Keypairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairs

`func (o *ServerQuotaSet) SetKeypairs(v ServerQuotaSetObject)`

SetKeypairs sets Keypairs field to given value.


### GetMetadataItems

`func (o *ServerQuotaSet) GetMetadataItems() ServerQuotaSetObject`

GetMetadataItems returns the MetadataItems field if non-nil, zero value otherwise.

### GetMetadataItemsOk

`func (o *ServerQuotaSet) GetMetadataItemsOk() (*ServerQuotaSetObject, bool)`

GetMetadataItemsOk returns a tuple with the MetadataItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataItems

`func (o *ServerQuotaSet) SetMetadataItems(v ServerQuotaSetObject)`

SetMetadataItems sets MetadataItems field to given value.


### GetRam

`func (o *ServerQuotaSet) GetRam() ServerQuotaSetObject`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *ServerQuotaSet) GetRamOk() (*ServerQuotaSetObject, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *ServerQuotaSet) SetRam(v ServerQuotaSetObject)`

SetRam sets Ram field to given value.


### GetServerGroupMembers

`func (o *ServerQuotaSet) GetServerGroupMembers() ServerQuotaSetObject`

GetServerGroupMembers returns the ServerGroupMembers field if non-nil, zero value otherwise.

### GetServerGroupMembersOk

`func (o *ServerQuotaSet) GetServerGroupMembersOk() (*ServerQuotaSetObject, bool)`

GetServerGroupMembersOk returns a tuple with the ServerGroupMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupMembers

`func (o *ServerQuotaSet) SetServerGroupMembers(v ServerQuotaSetObject)`

SetServerGroupMembers sets ServerGroupMembers field to given value.


### GetServerGroups

`func (o *ServerQuotaSet) GetServerGroups() ServerQuotaSetObject`

GetServerGroups returns the ServerGroups field if non-nil, zero value otherwise.

### GetServerGroupsOk

`func (o *ServerQuotaSet) GetServerGroupsOk() (*ServerQuotaSetObject, bool)`

GetServerGroupsOk returns a tuple with the ServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroups

`func (o *ServerQuotaSet) SetServerGroups(v ServerQuotaSetObject)`

SetServerGroups sets ServerGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


