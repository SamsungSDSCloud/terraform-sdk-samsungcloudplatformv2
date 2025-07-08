# MemberVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskType** | Pointer to [**DiskType**](DiskType.md) | disk_type | [optional] 
**Id** | Pointer to **string** | volume id | [optional] 
**Name** | Pointer to **string** | volume name | [optional] 
**Purpose** | Pointer to [**BlockStoragePurpose**](BlockStoragePurpose.md) | volume purpose | [optional] 
**SizeGb** | Pointer to **int32** | size_gb | [optional] 
**SoVolumeId** | Pointer to **string** | so volume id | [optional] 
**SoVolumeName** | Pointer to **string** | so volume name | [optional] 

## Methods

### NewMemberVolume

`func NewMemberVolume() *MemberVolume`

NewMemberVolume instantiates a new MemberVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberVolumeWithDefaults

`func NewMemberVolumeWithDefaults() *MemberVolume`

NewMemberVolumeWithDefaults instantiates a new MemberVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskType

`func (o *MemberVolume) GetDiskType() DiskType`

GetDiskType returns the DiskType field if non-nil, zero value otherwise.

### GetDiskTypeOk

`func (o *MemberVolume) GetDiskTypeOk() (*DiskType, bool)`

GetDiskTypeOk returns a tuple with the DiskType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskType

`func (o *MemberVolume) SetDiskType(v DiskType)`

SetDiskType sets DiskType field to given value.

### HasDiskType

`func (o *MemberVolume) HasDiskType() bool`

HasDiskType returns a boolean if a field has been set.

### GetId

`func (o *MemberVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberVolume) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MemberVolume) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MemberVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPurpose

`func (o *MemberVolume) GetPurpose() BlockStoragePurpose`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *MemberVolume) GetPurposeOk() (*BlockStoragePurpose, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *MemberVolume) SetPurpose(v BlockStoragePurpose)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *MemberVolume) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSizeGb

`func (o *MemberVolume) GetSizeGb() int32`

GetSizeGb returns the SizeGb field if non-nil, zero value otherwise.

### GetSizeGbOk

`func (o *MemberVolume) GetSizeGbOk() (*int32, bool)`

GetSizeGbOk returns a tuple with the SizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeGb

`func (o *MemberVolume) SetSizeGb(v int32)`

SetSizeGb sets SizeGb field to given value.

### HasSizeGb

`func (o *MemberVolume) HasSizeGb() bool`

HasSizeGb returns a boolean if a field has been set.

### GetSoVolumeId

`func (o *MemberVolume) GetSoVolumeId() string`

GetSoVolumeId returns the SoVolumeId field if non-nil, zero value otherwise.

### GetSoVolumeIdOk

`func (o *MemberVolume) GetSoVolumeIdOk() (*string, bool)`

GetSoVolumeIdOk returns a tuple with the SoVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoVolumeId

`func (o *MemberVolume) SetSoVolumeId(v string)`

SetSoVolumeId sets SoVolumeId field to given value.

### HasSoVolumeId

`func (o *MemberVolume) HasSoVolumeId() bool`

HasSoVolumeId returns a boolean if a field has been set.

### GetSoVolumeName

`func (o *MemberVolume) GetSoVolumeName() string`

GetSoVolumeName returns the SoVolumeName field if non-nil, zero value otherwise.

### GetSoVolumeNameOk

`func (o *MemberVolume) GetSoVolumeNameOk() (*string, bool)`

GetSoVolumeNameOk returns a tuple with the SoVolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoVolumeName

`func (o *MemberVolume) SetSoVolumeName(v string)`

SetSoVolumeName sets SoVolumeName field to given value.

### HasSoVolumeName

`func (o *MemberVolume) HasSoVolumeName() bool`

HasSoVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


