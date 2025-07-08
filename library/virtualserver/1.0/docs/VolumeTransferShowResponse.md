# VolumeTransferShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | 생성일시 | 
**Id** | **string** | Transfer ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewVolumeTransferShowResponse

`func NewVolumeTransferShowResponse(createdAt string, id string, volumeId string, ) *VolumeTransferShowResponse`

NewVolumeTransferShowResponse instantiates a new VolumeTransferShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeTransferShowResponseWithDefaults

`func NewVolumeTransferShowResponseWithDefaults() *VolumeTransferShowResponse`

NewVolumeTransferShowResponseWithDefaults instantiates a new VolumeTransferShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VolumeTransferShowResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeTransferShowResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeTransferShowResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *VolumeTransferShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeTransferShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeTransferShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VolumeTransferShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeTransferShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeTransferShowResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeTransferShowResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VolumeTransferShowResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VolumeTransferShowResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVolumeId

`func (o *VolumeTransferShowResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeTransferShowResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeTransferShowResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


