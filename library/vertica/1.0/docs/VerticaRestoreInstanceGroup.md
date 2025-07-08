# VerticaRestoreInstanceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockStorageGroups** | [**[]BlockStorageGroupRequest**](BlockStorageGroupRequest.md) | Block storage groups list | 
**RoleType** | [**InstanceGroupRoleType**](InstanceGroupRoleType.md) | Role type | 
**ServerTypeName** | **string** | Server type name | 

## Methods

### NewVerticaRestoreInstanceGroup

`func NewVerticaRestoreInstanceGroup(blockStorageGroups []BlockStorageGroupRequest, roleType InstanceGroupRoleType, serverTypeName string, ) *VerticaRestoreInstanceGroup`

NewVerticaRestoreInstanceGroup instantiates a new VerticaRestoreInstanceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticaRestoreInstanceGroupWithDefaults

`func NewVerticaRestoreInstanceGroupWithDefaults() *VerticaRestoreInstanceGroup`

NewVerticaRestoreInstanceGroupWithDefaults instantiates a new VerticaRestoreInstanceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockStorageGroups

`func (o *VerticaRestoreInstanceGroup) GetBlockStorageGroups() []BlockStorageGroupRequest`

GetBlockStorageGroups returns the BlockStorageGroups field if non-nil, zero value otherwise.

### GetBlockStorageGroupsOk

`func (o *VerticaRestoreInstanceGroup) GetBlockStorageGroupsOk() (*[]BlockStorageGroupRequest, bool)`

GetBlockStorageGroupsOk returns a tuple with the BlockStorageGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockStorageGroups

`func (o *VerticaRestoreInstanceGroup) SetBlockStorageGroups(v []BlockStorageGroupRequest)`

SetBlockStorageGroups sets BlockStorageGroups field to given value.


### GetRoleType

`func (o *VerticaRestoreInstanceGroup) GetRoleType() InstanceGroupRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *VerticaRestoreInstanceGroup) GetRoleTypeOk() (*InstanceGroupRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *VerticaRestoreInstanceGroup) SetRoleType(v InstanceGroupRoleType)`

SetRoleType sets RoleType field to given value.


### GetServerTypeName

`func (o *VerticaRestoreInstanceGroup) GetServerTypeName() string`

GetServerTypeName returns the ServerTypeName field if non-nil, zero value otherwise.

### GetServerTypeNameOk

`func (o *VerticaRestoreInstanceGroup) GetServerTypeNameOk() (*string, bool)`

GetServerTypeNameOk returns a tuple with the ServerTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeName

`func (o *VerticaRestoreInstanceGroup) SetServerTypeName(v string)`

SetServerTypeName sets ServerTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


