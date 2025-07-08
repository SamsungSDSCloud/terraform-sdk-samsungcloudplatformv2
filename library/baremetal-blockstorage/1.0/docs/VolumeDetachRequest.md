# VolumeDetachRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | **[]string** | List of object id to detach | 

## Methods

### NewVolumeDetachRequest

`func NewVolumeDetachRequest(attachments []string, ) *VolumeDetachRequest`

NewVolumeDetachRequest instantiates a new VolumeDetachRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeDetachRequestWithDefaults

`func NewVolumeDetachRequestWithDefaults() *VolumeDetachRequest`

NewVolumeDetachRequestWithDefaults instantiates a new VolumeDetachRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *VolumeDetachRequest) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *VolumeDetachRequest) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *VolumeDetachRequest) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


