# SrnTagList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Srn** | **string** | SRN | 
**Tags** | [**[]Tag**](Tag.md) |  | 

## Methods

### NewSrnTagList

`func NewSrnTagList(srn string, tags []Tag, ) *SrnTagList`

NewSrnTagList instantiates a new SrnTagList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrnTagListWithDefaults

`func NewSrnTagListWithDefaults() *SrnTagList`

NewSrnTagListWithDefaults instantiates a new SrnTagList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSrn

`func (o *SrnTagList) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *SrnTagList) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *SrnTagList) SetSrn(v string)`

SetSrn sets Srn field to given value.


### GetTags

`func (o *SrnTagList) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SrnTagList) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SrnTagList) SetTags(v []Tag)`

SetTags sets Tags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


