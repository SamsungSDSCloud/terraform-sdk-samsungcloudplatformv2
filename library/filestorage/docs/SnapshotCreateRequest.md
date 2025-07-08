# SnapshotCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VolumeId** | **string** | Volume ID | 

## Methods

### NewSnapshotCreateRequest

`func NewSnapshotCreateRequest(volumeId string, ) *SnapshotCreateRequest`

NewSnapshotCreateRequest instantiates a new SnapshotCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotCreateRequestWithDefaults

`func NewSnapshotCreateRequestWithDefaults() *SnapshotCreateRequest`

NewSnapshotCreateRequestWithDefaults instantiates a new SnapshotCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeId

`func (o *SnapshotCreateRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *SnapshotCreateRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *SnapshotCreateRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


