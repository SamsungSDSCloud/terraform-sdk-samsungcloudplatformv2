# PrivateDnsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedVpcIds** | Pointer to **[]string** | Connected VPC ID | [optional] [default to []]
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Private DNS Name | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewPrivateDnsCreateRequest

`func NewPrivateDnsCreateRequest(name string, ) *PrivateDnsCreateRequest`

NewPrivateDnsCreateRequest instantiates a new PrivateDnsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateDnsCreateRequestWithDefaults

`func NewPrivateDnsCreateRequestWithDefaults() *PrivateDnsCreateRequest`

NewPrivateDnsCreateRequestWithDefaults instantiates a new PrivateDnsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedVpcIds

`func (o *PrivateDnsCreateRequest) GetConnectedVpcIds() []string`

GetConnectedVpcIds returns the ConnectedVpcIds field if non-nil, zero value otherwise.

### GetConnectedVpcIdsOk

`func (o *PrivateDnsCreateRequest) GetConnectedVpcIdsOk() (*[]string, bool)`

GetConnectedVpcIdsOk returns a tuple with the ConnectedVpcIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedVpcIds

`func (o *PrivateDnsCreateRequest) SetConnectedVpcIds(v []string)`

SetConnectedVpcIds sets ConnectedVpcIds field to given value.

### HasConnectedVpcIds

`func (o *PrivateDnsCreateRequest) HasConnectedVpcIds() bool`

HasConnectedVpcIds returns a boolean if a field has been set.

### GetDescription

`func (o *PrivateDnsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateDnsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateDnsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateDnsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PrivateDnsCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PrivateDnsCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *PrivateDnsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateDnsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateDnsCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *PrivateDnsCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PrivateDnsCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PrivateDnsCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PrivateDnsCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *PrivateDnsCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *PrivateDnsCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


