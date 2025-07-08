# VolumeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CifsPassword** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Volume Name | 
**Protocol** | **string** | Protocol | 
**SnapshotId** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**TypeName** | **string** | Volume Type Name | 

## Methods

### NewVolumeCreateRequest

`func NewVolumeCreateRequest(name string, protocol string, typeName string, ) *VolumeCreateRequest`

NewVolumeCreateRequest instantiates a new VolumeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCreateRequestWithDefaults

`func NewVolumeCreateRequestWithDefaults() *VolumeCreateRequest`

NewVolumeCreateRequestWithDefaults instantiates a new VolumeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCifsPassword

`func (o *VolumeCreateRequest) GetCifsPassword() string`

GetCifsPassword returns the CifsPassword field if non-nil, zero value otherwise.

### GetCifsPasswordOk

`func (o *VolumeCreateRequest) GetCifsPasswordOk() (*string, bool)`

GetCifsPasswordOk returns a tuple with the CifsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsPassword

`func (o *VolumeCreateRequest) SetCifsPassword(v string)`

SetCifsPassword sets CifsPassword field to given value.

### HasCifsPassword

`func (o *VolumeCreateRequest) HasCifsPassword() bool`

HasCifsPassword returns a boolean if a field has been set.

### SetCifsPasswordNil

`func (o *VolumeCreateRequest) SetCifsPasswordNil(b bool)`

 SetCifsPasswordNil sets the value for CifsPassword to be an explicit nil

### UnsetCifsPassword
`func (o *VolumeCreateRequest) UnsetCifsPassword()`

UnsetCifsPassword ensures that no value is present for CifsPassword, not even an explicit nil
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


### GetProtocol

`func (o *VolumeCreateRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VolumeCreateRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VolumeCreateRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSnapshotId

`func (o *VolumeCreateRequest) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *VolumeCreateRequest) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *VolumeCreateRequest) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *VolumeCreateRequest) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *VolumeCreateRequest) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *VolumeCreateRequest) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetTags

`func (o *VolumeCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VolumeCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VolumeCreateRequest) SetTags(v []Tag)`

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
### GetTypeName

`func (o *VolumeCreateRequest) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *VolumeCreateRequest) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *VolumeCreateRequest) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


