# VolumeGroupCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Volume group name | 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 
**VolumeIds** | **[]string** | List of volume(Block Storage) id to add | 

## Methods

### NewVolumeGroupCreationRequest

`func NewVolumeGroupCreationRequest(name string, volumeIds []string, ) *VolumeGroupCreationRequest`

NewVolumeGroupCreationRequest instantiates a new VolumeGroupCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeGroupCreationRequestWithDefaults

`func NewVolumeGroupCreationRequestWithDefaults() *VolumeGroupCreationRequest`

NewVolumeGroupCreationRequestWithDefaults instantiates a new VolumeGroupCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeGroupCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeGroupCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeGroupCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *VolumeGroupCreationRequest) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VolumeGroupCreationRequest) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VolumeGroupCreationRequest) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VolumeGroupCreationRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VolumeGroupCreationRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VolumeGroupCreationRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetVolumeIds

`func (o *VolumeGroupCreationRequest) GetVolumeIds() []string`

GetVolumeIds returns the VolumeIds field if non-nil, zero value otherwise.

### GetVolumeIdsOk

`func (o *VolumeGroupCreationRequest) GetVolumeIdsOk() (*[]string, bool)`

GetVolumeIdsOk returns a tuple with the VolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeIds

`func (o *VolumeGroupCreationRequest) SetVolumeIds(v []string)`

SetVolumeIds sets VolumeIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


