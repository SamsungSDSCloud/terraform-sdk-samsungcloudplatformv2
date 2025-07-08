# ResourceGroupPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**ResourceGroups** | [**[]ResourceGroup**](ResourceGroup.md) |  | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewResourceGroupPageResponse

`func NewResourceGroupPageResponse(count int32, page int32, resourceGroups []ResourceGroup, size int32, ) *ResourceGroupPageResponse`

NewResourceGroupPageResponse instantiates a new ResourceGroupPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupPageResponseWithDefaults

`func NewResourceGroupPageResponseWithDefaults() *ResourceGroupPageResponse`

NewResourceGroupPageResponseWithDefaults instantiates a new ResourceGroupPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ResourceGroupPageResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ResourceGroupPageResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ResourceGroupPageResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *ResourceGroupPageResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ResourceGroupPageResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ResourceGroupPageResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetResourceGroups

`func (o *ResourceGroupPageResponse) GetResourceGroups() []ResourceGroup`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *ResourceGroupPageResponse) GetResourceGroupsOk() (*[]ResourceGroup, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *ResourceGroupPageResponse) SetResourceGroups(v []ResourceGroup)`

SetResourceGroups sets ResourceGroups field to given value.


### GetSize

`func (o *ResourceGroupPageResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ResourceGroupPageResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ResourceGroupPageResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *ResourceGroupPageResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ResourceGroupPageResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ResourceGroupPageResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ResourceGroupPageResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *ResourceGroupPageResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *ResourceGroupPageResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


