# ClusterRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**Groups** | **[]string** | Groups | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Cluster Role Binding Name | 
**Role** | [**RoleRef**](RoleRef.md) | Role | 
**ServiceAccounts** | **[]string** | Service Accounts | 
**Subjects** | [**[]Subject**](Subject.md) | Subjects | 
**Uid** | **string** | Resource ID | 
**Users** | **[]string** | Users | 
**Yaml** | **string** | YAML | 

## Methods

### NewClusterRoleBinding

`func NewClusterRoleBinding(age string, annotations []string, clusterId string, clusterName string, createdAt time.Time, groups []string, labels []string, name string, role RoleRef, serviceAccounts []string, subjects []Subject, uid string, users []string, yaml string, ) *ClusterRoleBinding`

NewClusterRoleBinding instantiates a new ClusterRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleBindingWithDefaults

`func NewClusterRoleBindingWithDefaults() *ClusterRoleBinding`

NewClusterRoleBindingWithDefaults instantiates a new ClusterRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *ClusterRoleBinding) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ClusterRoleBinding) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ClusterRoleBinding) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *ClusterRoleBinding) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *ClusterRoleBinding) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *ClusterRoleBinding) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *ClusterRoleBinding) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ClusterRoleBinding) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ClusterRoleBinding) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *ClusterRoleBinding) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *ClusterRoleBinding) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *ClusterRoleBinding) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *ClusterRoleBinding) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterRoleBinding) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterRoleBinding) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetGroups

`func (o *ClusterRoleBinding) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ClusterRoleBinding) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ClusterRoleBinding) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetLabels

`func (o *ClusterRoleBinding) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterRoleBinding) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterRoleBinding) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *ClusterRoleBinding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRoleBinding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRoleBinding) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *ClusterRoleBinding) GetRole() RoleRef`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ClusterRoleBinding) GetRoleOk() (*RoleRef, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ClusterRoleBinding) SetRole(v RoleRef)`

SetRole sets Role field to given value.


### GetServiceAccounts

`func (o *ClusterRoleBinding) GetServiceAccounts() []string`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ClusterRoleBinding) GetServiceAccountsOk() (*[]string, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ClusterRoleBinding) SetServiceAccounts(v []string)`

SetServiceAccounts sets ServiceAccounts field to given value.


### GetSubjects

`func (o *ClusterRoleBinding) GetSubjects() []Subject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *ClusterRoleBinding) GetSubjectsOk() (*[]Subject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *ClusterRoleBinding) SetSubjects(v []Subject)`

SetSubjects sets Subjects field to given value.


### GetUid

`func (o *ClusterRoleBinding) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ClusterRoleBinding) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ClusterRoleBinding) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUsers

`func (o *ClusterRoleBinding) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ClusterRoleBinding) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ClusterRoleBinding) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetYaml

`func (o *ClusterRoleBinding) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *ClusterRoleBinding) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *ClusterRoleBinding) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


