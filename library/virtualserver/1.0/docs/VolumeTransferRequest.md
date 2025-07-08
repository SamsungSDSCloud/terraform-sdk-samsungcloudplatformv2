# VolumeTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewVolumeTransferRequest

`func NewVolumeTransferRequest(volumeId string, ) *VolumeTransferRequest`

NewVolumeTransferRequest instantiates a new VolumeTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeTransferRequestWithDefaults

`func NewVolumeTransferRequestWithDefaults() *VolumeTransferRequest`

NewVolumeTransferRequestWithDefaults instantiates a new VolumeTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VolumeTransferRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeTransferRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeTransferRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeTransferRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VolumeTransferRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VolumeTransferRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVolumeId

`func (o *VolumeTransferRequest) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeTransferRequest) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeTransferRequest) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


