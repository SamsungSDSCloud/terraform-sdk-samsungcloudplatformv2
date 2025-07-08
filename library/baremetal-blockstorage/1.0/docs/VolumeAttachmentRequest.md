# VolumeAttachmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | [**[]AttachmentListModel**](AttachmentListModel.md) | List of object info to attach | 

## Methods

### NewVolumeAttachmentRequest

`func NewVolumeAttachmentRequest(attachments []AttachmentListModel, ) *VolumeAttachmentRequest`

NewVolumeAttachmentRequest instantiates a new VolumeAttachmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeAttachmentRequestWithDefaults

`func NewVolumeAttachmentRequestWithDefaults() *VolumeAttachmentRequest`

NewVolumeAttachmentRequestWithDefaults instantiates a new VolumeAttachmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *VolumeAttachmentRequest) GetAttachments() []AttachmentListModel`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumeAttachmentRequest) GetAttachmentsOk() (*[]AttachmentListModel, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumeAttachmentRequest) SetAttachments(v []AttachmentListModel)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


