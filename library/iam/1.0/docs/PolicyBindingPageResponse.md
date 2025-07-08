# PolicyBindingPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Groups** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Page** | **int32** | page | 
**PolicyId** | Pointer to **string** | Policy ID | [optional] 
**Roles** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 
**Users** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewPolicyBindingPageResponse

`func NewPolicyBindingPageResponse(count int32, page int32, size int32, ) *PolicyBindingPageResponse`

NewPolicyBindingPageResponse instantiates a new PolicyBindingPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyBindingPageResponseWithDefaults

`func NewPolicyBindingPageResponseWithDefaults() *PolicyBindingPageResponse`

NewPolicyBindingPageResponseWithDefaults instantiates a new PolicyBindingPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PolicyBindingPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PolicyBindingPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PolicyBindingPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetGroups

`func (o *PolicyBindingPageResponse) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *PolicyBindingPageResponse) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *PolicyBindingPageResponse) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *PolicyBindingPageResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### SetGroupsNil

`func (o *PolicyBindingPageResponse) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *PolicyBindingPageResponse) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetPage

`func (o *PolicyBindingPageResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PolicyBindingPageResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PolicyBindingPageResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPolicyId

`func (o *PolicyBindingPageResponse) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyBindingPageResponse) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyBindingPageResponse) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyBindingPageResponse) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetRoles

`func (o *PolicyBindingPageResponse) GetRoles() []map[string]interface{}`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *PolicyBindingPageResponse) GetRolesOk() (*[]map[string]interface{}, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *PolicyBindingPageResponse) SetRoles(v []map[string]interface{})`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *PolicyBindingPageResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *PolicyBindingPageResponse) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *PolicyBindingPageResponse) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetSize

`func (o *PolicyBindingPageResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PolicyBindingPageResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PolicyBindingPageResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *PolicyBindingPageResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PolicyBindingPageResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PolicyBindingPageResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PolicyBindingPageResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *PolicyBindingPageResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *PolicyBindingPageResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetUsers

`func (o *PolicyBindingPageResponse) GetUsers() []map[string]interface{}`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PolicyBindingPageResponse) GetUsersOk() (*[]map[string]interface{}, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PolicyBindingPageResponse) SetUsers(v []map[string]interface{})`

SetUsers sets Users field to given value.

### HasUsers

`func (o *PolicyBindingPageResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *PolicyBindingPageResponse) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *PolicyBindingPageResponse) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


