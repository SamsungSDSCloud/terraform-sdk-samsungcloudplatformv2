# ClusterRoleBindingListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterRoleBindings** | [**[]ClusterRoleBindingSummary**](ClusterRoleBindingSummary.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewClusterRoleBindingListResponse

`func NewClusterRoleBindingListResponse(clusterRoleBindings []ClusterRoleBindingSummary, ) *ClusterRoleBindingListResponse`

NewClusterRoleBindingListResponse instantiates a new ClusterRoleBindingListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleBindingListResponseWithDefaults

`func NewClusterRoleBindingListResponseWithDefaults() *ClusterRoleBindingListResponse`

NewClusterRoleBindingListResponseWithDefaults instantiates a new ClusterRoleBindingListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterRoleBindings

`func (o *ClusterRoleBindingListResponse) GetClusterRoleBindings() []ClusterRoleBindingSummary`

GetClusterRoleBindings returns the ClusterRoleBindings field if non-nil, zero value otherwise.

### GetClusterRoleBindingsOk

`func (o *ClusterRoleBindingListResponse) GetClusterRoleBindingsOk() (*[]ClusterRoleBindingSummary, bool)`

GetClusterRoleBindingsOk returns a tuple with the ClusterRoleBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRoleBindings

`func (o *ClusterRoleBindingListResponse) SetClusterRoleBindings(v []ClusterRoleBindingSummary)`

SetClusterRoleBindings sets ClusterRoleBindings field to given value.


### GetCount

`func (o *ClusterRoleBindingListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ClusterRoleBindingListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ClusterRoleBindingListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ClusterRoleBindingListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ClusterRoleBindingListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ClusterRoleBindingListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *ClusterRoleBindingListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClusterRoleBindingListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClusterRoleBindingListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClusterRoleBindingListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ClusterRoleBindingListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ClusterRoleBindingListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


