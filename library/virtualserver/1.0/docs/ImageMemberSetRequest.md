# ImageMemberSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**ImageMemberStatusEnum**](ImageMemberStatusEnum.md) | Member status | 

## Methods

### NewImageMemberSetRequest

`func NewImageMemberSetRequest(status ImageMemberStatusEnum, ) *ImageMemberSetRequest`

NewImageMemberSetRequest instantiates a new ImageMemberSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageMemberSetRequestWithDefaults

`func NewImageMemberSetRequestWithDefaults() *ImageMemberSetRequest`

NewImageMemberSetRequestWithDefaults instantiates a new ImageMemberSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ImageMemberSetRequest) GetStatus() ImageMemberStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageMemberSetRequest) GetStatusOk() (*ImageMemberStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageMemberSetRequest) SetStatus(v ImageMemberStatusEnum)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


