# TagsSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | [**map[string][]Tag**](array.md) |  | 

## Methods

### NewTagsSetRequest

`func NewTagsSetRequest(tags map[string][]Tag, ) *TagsSetRequest`

NewTagsSetRequest instantiates a new TagsSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsSetRequestWithDefaults

`func NewTagsSetRequestWithDefaults() *TagsSetRequest`

NewTagsSetRequestWithDefaults instantiates a new TagsSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *TagsSetRequest) GetTags() map[string][]Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TagsSetRequest) GetTagsOk() (*map[string][]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TagsSetRequest) SetTags(v map[string][]Tag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


