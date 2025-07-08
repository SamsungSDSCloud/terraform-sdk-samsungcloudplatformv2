# InstanceGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorageGroups** | [**[]BlockStorageGroupResponse**](BlockStorageGroupResponse.md) | Block storage groups list | 
**Id** | **string** | ID | 
**Instances** | [**[]InstanceResponse**](InstanceResponse.md) | Instances list | 
**RoleType** | [**InstanceGroupRoleType**](InstanceGroupRoleType.md) | Role type | 
**ServerTypeName** | **string** | Server type name | 

## Methods

### NewInstanceGroupResponse

`func NewInstanceGroupResponse(blockStorageGroups []BlockStorageGroupResponse, id string, instances []InstanceResponse, roleType InstanceGroupRoleType, serverTypeName string, ) *InstanceGroupResponse`

NewInstanceGroupResponse instantiates a new InstanceGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceGroupResponseWithDefaults

`func NewInstanceGroupResponseWithDefaults() *InstanceGroupResponse`

NewInstanceGroupResponseWithDefaults instantiates a new InstanceGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockStorageGroups

`func (o *InstanceGroupResponse) GetBlockStorageGroups() []BlockStorageGroupResponse`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *InstanceGroupResponse) GetBlockStorageGroupsOk() (*[]BlockStorageGroupResponse, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *InstanceGroupResponse) SetBlockStorageGroups(v []BlockStorageGroupResponse)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.


### GetId

`func (o *InstanceGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInstances

`func (o *InstanceGroupResponse) GetInstances() []InstanceResponse`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *InstanceGroupResponse) GetInstancesOk() (*[]InstanceResponse, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *InstanceGroupResponse) SetInstances(v []InstanceResponse)`

SetInstances sets Instances field to given value.


### GetRoleType

`func (o *InstanceGroupResponse) GetRoleType() InstanceGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *InstanceGroupResponse) GetRoleTypeOk() (*InstanceGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *InstanceGroupResponse) SetRoleType(v InstanceGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetServerTypeName

`func (o *InstanceGroupResponse) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *InstanceGroupResponse) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *InstanceGroupResponse) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


