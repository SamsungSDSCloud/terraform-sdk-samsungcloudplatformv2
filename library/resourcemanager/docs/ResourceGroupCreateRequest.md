# ResourceGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**Name** | **string** | 리소스 그룹명 | 
**ResourceTypes** | Pointer to **[]string** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewResourceGroupCreateRequest

`func NewResourceGroupCreateRequest(description NullableString, name string, ) *ResourceGroupCreateRequest`

NewResourceGroupCreateRequest instantiates a new ResourceGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceGroupCreateRequestWithDefaults

`func NewResourceGroupCreateRequestWithDefaults() *ResourceGroupCreateRequest`

NewResourceGroupCreateRequestWithDefaults instantiates a new ResourceGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResourceGroupCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceGroupCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceGroupCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ResourceGroupCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ResourceGroupCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *ResourceGroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceGroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceGroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetResourceTypes

`func (o *ResourceGroupCreateRequest) GetResourceTypes() []string`

GetResourceTypes returns the ResourceTypes field if non-nil, zero value otherwise.

### GetResourceTypesOk

`func (o *ResourceGroupCreateRequest) GetResourceTypesOk() (*[]string, bool)`

GetResourceTypesOk returns a tuple with the ResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypes

`func (o *ResourceGroupCreateRequest) SetResourceTypes(v []string)`

SetResourceTypes sets ResourceTypes field to given value.

### HasResourceTypes

`func (o *ResourceGroupCreateRequest) HasResourceTypes() bool`

HasResourceTypes returns a boolean if a field has been set.

### SetResourceTypesNil

`func (o *ResourceGroupCreateRequest) SetResourceTypesNil(b bool)`

 SetResourceTypesNil sets the value for ResourceTypes to be an explicit nil

### UnsetResourceTypes
`func (o *ResourceGroupCreateRequest) UnsetResourceTypes()`

UnsetResourceTypes ensures that no value is present for ResourceTypes, not even an explicit nil
### GetTags

`func (o *ResourceGroupCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResourceGroupCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResourceGroupCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResourceGroupCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ResourceGroupCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ResourceGroupCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


