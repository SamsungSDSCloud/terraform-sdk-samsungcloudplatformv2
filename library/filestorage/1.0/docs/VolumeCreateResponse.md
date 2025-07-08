# VolumeCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Protocol** | **string** | Protocol | 
**SnapshotId** | Pointer to **NullableString** |  | [optional] 
**TypeName** | **string** | Volume Type Name | 
**VolumeId** | **string** | Volume ID | 
**VolumeName** | **string** | Volume Name | 

## Methods

### NewVolumeCreateResponse

`func NewVolumeCreateResponse(protocol string, typeName string, volumeId string, volumeName string, ) *VolumeCreateResponse`

NewVolumeCreateResponse instantiates a new VolumeCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeCreateResponseWithDefaults

`func NewVolumeCreateResponseWithDefaults() *VolumeCreateResponse`

NewVolumeCreateResponseWithDefaults instantiates a new VolumeCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProtocol

`func (o *VolumeCreateResponse) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VolumeCreateResponse) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VolumeCreateResponse) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetSnapshotId

`func (o *VolumeCreateResponse) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *VolumeCreateResponse) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *VolumeCreateResponse) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *VolumeCreateResponse) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *VolumeCreateResponse) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *VolumeCreateResponse) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil
### GetTypeName

`func (o *VolumeCreateResponse) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *VolumeCreateResponse) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *VolumeCreateResponse) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetVolumeId

`func (o *VolumeCreateResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeCreateResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeCreateResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.


### GetVolumeName

`func (o *VolumeCreateResponse) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *VolumeCreateResponse) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *VolumeCreateResponse) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


