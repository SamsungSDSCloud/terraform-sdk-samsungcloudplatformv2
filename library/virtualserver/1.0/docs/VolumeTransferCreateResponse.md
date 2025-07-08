# VolumeTransferCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthKey** | **string** | Auth Key | 
**CreatedAt** | **string** | 생성일시 | 
**Id** | **string** | Transfer ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewVolumeTransferCreateResponse

`func NewVolumeTransferCreateResponse(authKey string, createdAt string, id string, volumeId string, ) *VolumeTransferCreateResponse`

NewVolumeTransferCreateResponse instantiates a new VolumeTransferCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeTransferCreateResponseWithDefaults

`func NewVolumeTransferCreateResponseWithDefaults() *VolumeTransferCreateResponse`

NewVolumeTransferCreateResponseWithDefaults instantiates a new VolumeTransferCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthKey

`func (o *VolumeTransferCreateResponse) GetAuthKey() string`

GetAuthKey returns the AuthKey field if non-nil, zero value otherwise.

### GetAuthKeyOk

`func (o *VolumeTransferCreateResponse) GetAuthKeyOk() (*string, bool)`

GetAuthKeyOk returns a tuple with the AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthKey

`func (o *VolumeTransferCreateResponse) SetAuthKey(v string)`

SetAuthKey sets AuthKey field to given value.


### GetCreatedAt

`func (o *VolumeTransferCreateResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeTransferCreateResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeTransferCreateResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *VolumeTransferCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeTransferCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeTransferCreateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VolumeTransferCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeTransferCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeTransferCreateResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VolumeTransferCreateResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VolumeTransferCreateResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VolumeTransferCreateResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetVolumeId

`func (o *VolumeTransferCreateResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeTransferCreateResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeTransferCreateResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


