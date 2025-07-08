# ImageCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerFormat** | Pointer to **string** | Container format | [optional] [default to "bare"]
**DiskFormat** | Pointer to **string** | Disk format | [optional] [default to "qcow2"]
**MinDisk** | Pointer to **int32** | Minimum disk | [optional] [default to 0]
**MinRam** | Pointer to **int32** | Minimum RAM | [optional] [default to 0]
**Name** | **string** | Name | 
**OsDistro** | [**ImageOsDistroEnum**](ImageOsDistroEnum.md) | OS distribution | 
**Protected** | Pointer to **bool** | Protected | [optional] [default to false]
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**Url** | **string** | Object Storage URL for Image file (only qcow2 format allowed) | 
**Visibility** | Pointer to **string** | Visibility | [optional] [default to "private"]

## Methods

### NewImageCreateRequest

`func NewImageCreateRequest(name string, osDistro ImageOsDistroEnum, url string, ) *ImageCreateRequest`

NewImageCreateRequest instantiates a new ImageCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageCreateRequestWithDefaults

`func NewImageCreateRequestWithDefaults() *ImageCreateRequest`

NewImageCreateRequestWithDefaults instantiates a new ImageCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerFormat

`func (o *ImageCreateRequest) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *ImageCreateRequest) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *ImageCreateRequest) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *ImageCreateRequest) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### GetDiskFormat

`func (o *ImageCreateRequest) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *ImageCreateRequest) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *ImageCreateRequest) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *ImageCreateRequest) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### GetMinDisk

`func (o *ImageCreateRequest) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ImageCreateRequest) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ImageCreateRequest) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *ImageCreateRequest) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### GetMinRam

`func (o *ImageCreateRequest) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *ImageCreateRequest) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *ImageCreateRequest) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *ImageCreateRequest) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### GetName

`func (o *ImageCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOsDistro

`func (o *ImageCreateRequest) GetOsDistro() ImageOsDistroEnum`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *ImageCreateRequest) GetOsDistroOk() (*ImageOsDistroEnum, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *ImageCreateRequest) SetOsDistro(v ImageOsDistroEnum)`

SetOsDistro sets OsDistro field to given value.


### GetProtected

`func (o *ImageCreateRequest) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *ImageCreateRequest) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *ImageCreateRequest) SetProtected(v bool)`

SetProtected sets Protected field to given value.

### HasProtected

`func (o *ImageCreateRequest) HasProtected() bool`

HasProtected returns a boolean if a field has been set.

### GetTags

`func (o *ImageCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ImageCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ImageCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ImageCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ImageCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ImageCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUrl

`func (o *ImageCreateRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageCreateRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageCreateRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVisibility

`func (o *ImageCreateRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ImageCreateRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ImageCreateRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ImageCreateRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


