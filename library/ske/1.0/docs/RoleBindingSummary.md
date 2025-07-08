# RoleBindingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**Groups** | **[]string** | Groups | 
**Name** | **string** | Role Binding Name | 
**NamespaceName** | **string** | Namespace Name | 
**Role** | [**RoleRef**](RoleRef.md) | Role | 
**ServiceAccounts** | **[]string** | Service Accounts | 
**Uid** | **string** | Resource ID | 
**Users** | **[]string** | Users | 

## Methods

### NewRoleBindingSummary

`func NewRoleBindingSummary(age string, clusterId string, createdAt time.Time, groups []string, name string, namespaceName string, role RoleRef, serviceAccounts []string, uid string, users []string, ) *RoleBindingSummary`

NewRoleBindingSummary instantiates a new RoleBindingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBindingSummaryWithDefaults

`func NewRoleBindingSummaryWithDefaults() *RoleBindingSummary`

NewRoleBindingSummaryWithDefaults instantiates a new RoleBindingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *RoleBindingSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *RoleBindingSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *RoleBindingSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *RoleBindingSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RoleBindingSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RoleBindingSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *RoleBindingSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RoleBindingSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RoleBindingSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGroups

`func (o *RoleBindingSummary) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RoleBindingSummary) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RoleBindingSummary) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetName

`func (o *RoleBindingSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleBindingSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleBindingSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *RoleBindingSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *RoleBindingSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *RoleBindingSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetRole

`func (o *RoleBindingSummary) GetRole() RoleRef`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleBindingSummary) GetRoleOk() (*RoleRef, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleBindingSummary) SetRole(v RoleRef)`

SetRole sets Role field to given value.


### GetServiceAccounts

`func (o *RoleBindingSummary) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *RoleBindingSummary) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *RoleBindingSummary) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.


### GetUid

`func (o *RoleBindingSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *RoleBindingSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *RoleBindingSummary) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUsers

`func (o *RoleBindingSummary) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *RoleBindingSummary) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *RoleBindingSummary) SetUsers(v []string)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


