# ClusterRoleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterRoles** | [**[]ClusterRoleSummary**](ClusterRoleSummary.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewClusterRoleListResponse

`func NewClusterRoleListResponse(clusterRoles []ClusterRoleSummary, ) *ClusterRoleListResponse`

NewClusterRoleListResponse instantiates a new ClusterRoleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleListResponseWithDefaults

`func NewClusterRoleListResponseWithDefaults() *ClusterRoleListResponse`

NewClusterRoleListResponseWithDefaults instantiates a new ClusterRoleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterRoles

`func (o *ClusterRoleListResponse) GetClusterRoles() []ClusterRoleSummary`

GetClusterRoles returns the ClusterRoles field if non-nil, zero value otherwise.

### GetClusterRolesOk

`func (o *ClusterRoleListResponse) GetClusterRolesOk() (*[]ClusterRoleSummary, bool)`

GetClusterRolesOk returns a tuple with the ClusterRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRoles

`func (o *ClusterRoleListResponse) SetClusterRoles(v []ClusterRoleSummary)`

SetClusterRoles sets ClusterRoles field to given value.


### GetCount

`func (o *ClusterRoleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ClusterRoleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ClusterRoleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ClusterRoleListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ClusterRoleListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ClusterRoleListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *ClusterRoleListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClusterRoleListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClusterRoleListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClusterRoleListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ClusterRoleListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ClusterRoleListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


