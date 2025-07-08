# RoleBindingListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**RoleBindings** | [**[]RoleBindingSummary**](RoleBindingSummary.md) |  | 

## Methods

### NewRoleBindingListResponse

`func NewRoleBindingListResponse(roleBindings []RoleBindingSummary, ) *RoleBindingListResponse`

NewRoleBindingListResponse instantiates a new RoleBindingListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleBindingListResponseWithDefaults

`func NewRoleBindingListResponseWithDefaults() *RoleBindingListResponse`

NewRoleBindingListResponseWithDefaults instantiates a new RoleBindingListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *RoleBindingListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *RoleBindingListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *RoleBindingListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *RoleBindingListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *RoleBindingListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *RoleBindingListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *RoleBindingListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleBindingListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleBindingListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleBindingListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *RoleBindingListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *RoleBindingListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetRoleBindings

`func (o *RoleBindingListResponse) GetRoleBindings() []RoleBindingSummary`

GetRoleBindings returns the RoleBindings field if non-nil, zero value otherwise.

### GetRoleBindingsOk

`func (o *RoleBindingListResponse) GetRoleBindingsOk() (*[]RoleBindingSummary, bool)`

GetRoleBindingsOk returns a tuple with the RoleBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleBindings

`func (o *RoleBindingListResponse) SetRoleBindings(v []RoleBindingSummary)`

SetRoleBindings sets RoleBindings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


