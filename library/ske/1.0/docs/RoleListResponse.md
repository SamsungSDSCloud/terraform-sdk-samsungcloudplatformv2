# RoleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**Roles** | [**[]RoleSummary**](RoleSummary.md) |  | 

## Methods

### NewRoleListResponse

`func NewRoleListResponse(roles []RoleSummary, ) *RoleListResponse`

NewRoleListResponse instantiates a new RoleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleListResponseWithDefaults

`func NewRoleListResponseWithDefaults() *RoleListResponse`

NewRoleListResponseWithDefaults instantiates a new RoleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RoleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RoleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RoleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RoleListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *RoleListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *RoleListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *RoleListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *RoleListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *RoleListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRoles

`func (o *RoleListResponse) GetRoles() []RoleSummary`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RoleListResponse) GetRolesOk() (*[]RoleSummary, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RoleListResponse) SetRoles(v []RoleSummary)`

SetRoles sets Roles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


