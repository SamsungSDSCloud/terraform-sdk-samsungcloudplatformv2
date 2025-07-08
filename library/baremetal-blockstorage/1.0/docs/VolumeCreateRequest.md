# VolumeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]AttachmentListModel**](AttachmentListModel.md) | List of server id and type | 
**DiskType** | [**DiskType**](DiskType.md) | Disk type | 
**Name** | **string** | Volume name | 
**SizeGb** | **int32** | Volume capacity(GB) | 
**Tags** | Pointer to [**[]TagModel**](TagModel.md) |  | [optional] 

## Methods

### NewVolumeCreateRequest

`func NewVolumeCreateRequest(attachments []AttachmentListModel, diskType DiskType, name string, sizeGb int32, ) *VolumeCreateRequest`

NewVolumeCreateRequest instantiates a new VolumeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCreateRequestWithDefaults

`func NewVolumeCreateRequestWithDefaults() *VolumeCreateRequest`

NewVolumeCreateRequestWithDefaults instantiates a new VolumeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *VolumeCreateRequest) GetAttachments() []AttachmentListModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumeCreateRequest) GetAttachmentsOk() (*[]AttachmentListModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumeCreateRequest) SetAttachments(v []AttachmentListModel)`

SetAttachments sets Attachments field to given value.


### GetDiskType

`func (o *VolumeCreateRequest) GetDiskType() DiskType`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *VolumeCreateRequest) GetDiskTypeOk() (*DiskType, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *VolumeCreateRequest) SetDiskType(v DiskType)`

SetDiskType sets DiskType field to given value.


### GetName

`func (o *VolumeCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSizeGb

`func (o *VolumeCreateRequest) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *VolumeCreateRequest) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *VolumeCreateRequest) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.


### GetTags

`func (o *VolumeCreateRequest) GetTags() []TagModel`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VolumeCreateRequest) GetTagsOk() (*[]TagModel, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VolumeCreateRequest) SetTags(v []TagModel)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VolumeCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *VolumeCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *VolumeCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


