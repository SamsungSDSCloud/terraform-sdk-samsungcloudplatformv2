# GroupPolicyPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Policies** | [**[]Policy**](Policy.md) |  | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGroupPolicyPageResponse

`func NewGroupPolicyPageResponse(count int32, page int32, policies []Policy, size int32, ) *GroupPolicyPageResponse`

NewGroupPolicyPageResponse instantiates a new GroupPolicyPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPolicyPageResponseWithDefaults

`func NewGroupPolicyPageResponseWithDefaults() *GroupPolicyPageResponse`

NewGroupPolicyPageResponseWithDefaults instantiates a new GroupPolicyPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *GroupPolicyPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GroupPolicyPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GroupPolicyPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *GroupPolicyPageResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GroupPolicyPageResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GroupPolicyPageResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetPolicies

`func (o *GroupPolicyPageResponse) GetPolicies() []Policy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *GroupPolicyPageResponse) GetPoliciesOk() (*[]Policy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *GroupPolicyPageResponse) SetPolicies(v []Policy)`

SetPolicies sets Policies field to given value.


### GetSize

`func (o *GroupPolicyPageResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GroupPolicyPageResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GroupPolicyPageResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *GroupPolicyPageResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GroupPolicyPageResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GroupPolicyPageResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GroupPolicyPageResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *GroupPolicyPageResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *GroupPolicyPageResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


