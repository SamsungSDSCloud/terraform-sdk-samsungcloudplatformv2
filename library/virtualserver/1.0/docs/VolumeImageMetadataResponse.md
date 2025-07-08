# VolumeImageMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to **NullableString** |  | [optional] 
**ContainerFormat** | Pointer to **NullableString** |  | [optional] 
**DiskFormat** | Pointer to **NullableString** |  | [optional] 
**ImageId** | Pointer to **NullableString** |  | [optional] 
**ImageName** | Pointer to **NullableString** |  | [optional] 
**MinDisk** | Pointer to **NullableInt32** |  | [optional] 
**MinRam** | Pointer to **NullableInt32** |  | [optional] 
**OsDistro** | Pointer to **NullableString** |  | [optional] 
**ScpImageType** | Pointer to **NullableString** |  | [optional] 
**ScpK8sVersion** | Pointer to **NullableString** |  | [optional] 
**ScpOriginalImageType** | Pointer to **NullableString** |  | [optional] 
**SignatureVerified** | Pointer to **NullableBool** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 

## Methods

### NewVolumeImageMetadataResponse

`func NewVolumeImageMetadataResponse() *VolumeImageMetadataResponse`

NewVolumeImageMetadataResponse instantiates a new VolumeImageMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeImageMetadataResponseWithDefaults

`func NewVolumeImageMetadataResponseWithDefaults() *VolumeImageMetadataResponse`

NewVolumeImageMetadataResponseWithDefaults instantiates a new VolumeImageMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *VolumeImageMetadataResponse) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *VolumeImageMetadataResponse) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *VolumeImageMetadataResponse) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *VolumeImageMetadataResponse) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *VolumeImageMetadataResponse) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *VolumeImageMetadataResponse) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetContainerFormat

`func (o *VolumeImageMetadataResponse) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *VolumeImageMetadataResponse) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *VolumeImageMetadataResponse) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.

### HasContainerFormat

`func (o *VolumeImageMetadataResponse) HasContainerFormat() bool`

HasContainerFormat returns a boolean if a field has been set.

### SetContainerFormatNil

`func (o *VolumeImageMetadataResponse) SetContainerFormatNil(b bool)`

 SetContainerFormatNil sets the value for ContainerFormat to be an explicit nil

### UnsetContainerFormat
`func (o *VolumeImageMetadataResponse) UnsetContainerFormat()`

UnsetContainerFormat ensures that no value is present for ContainerFormat, not even an explicit nil
### GetDiskFormat

`func (o *VolumeImageMetadataResponse) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *VolumeImageMetadataResponse) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *VolumeImageMetadataResponse) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.

### HasDiskFormat

`func (o *VolumeImageMetadataResponse) HasDiskFormat() bool`

HasDiskFormat returns a boolean if a field has been set.

### SetDiskFormatNil

`func (o *VolumeImageMetadataResponse) SetDiskFormatNil(b bool)`

 SetDiskFormatNil sets the value for DiskFormat to be an explicit nil

### UnsetDiskFormat
`func (o *VolumeImageMetadataResponse) UnsetDiskFormat()`

UnsetDiskFormat ensures that no value is present for DiskFormat, not even an explicit nil
### GetImageId

`func (o *VolumeImageMetadataResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *VolumeImageMetadataResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *VolumeImageMetadataResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *VolumeImageMetadataResponse) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### SetImageIdNil

`func (o *VolumeImageMetadataResponse) SetImageIdNil(b bool)`

 SetImageIdNil sets the value for ImageId to be an explicit nil

### UnsetImageId
`func (o *VolumeImageMetadataResponse) UnsetImageId()`

UnsetImageId ensures that no value is present for ImageId, not even an explicit nil
### GetImageName

`func (o *VolumeImageMetadataResponse) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *VolumeImageMetadataResponse) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *VolumeImageMetadataResponse) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *VolumeImageMetadataResponse) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### SetImageNameNil

`func (o *VolumeImageMetadataResponse) SetImageNameNil(b bool)`

 SetImageNameNil sets the value for ImageName to be an explicit nil

### UnsetImageName
`func (o *VolumeImageMetadataResponse) UnsetImageName()`

UnsetImageName ensures that no value is present for ImageName, not even an explicit nil
### GetMinDisk

`func (o *VolumeImageMetadataResponse) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *VolumeImageMetadataResponse) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *VolumeImageMetadataResponse) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.

### HasMinDisk

`func (o *VolumeImageMetadataResponse) HasMinDisk() bool`

HasMinDisk returns a boolean if a field has been set.

### SetMinDiskNil

`func (o *VolumeImageMetadataResponse) SetMinDiskNil(b bool)`

 SetMinDiskNil sets the value for MinDisk to be an explicit nil

### UnsetMinDisk
`func (o *VolumeImageMetadataResponse) UnsetMinDisk()`

UnsetMinDisk ensures that no value is present for MinDisk, not even an explicit nil
### GetMinRam

`func (o *VolumeImageMetadataResponse) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *VolumeImageMetadataResponse) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *VolumeImageMetadataResponse) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.

### HasMinRam

`func (o *VolumeImageMetadataResponse) HasMinRam() bool`

HasMinRam returns a boolean if a field has been set.

### SetMinRamNil

`func (o *VolumeImageMetadataResponse) SetMinRamNil(b bool)`

 SetMinRamNil sets the value for MinRam to be an explicit nil

### UnsetMinRam
`func (o *VolumeImageMetadataResponse) UnsetMinRam()`

UnsetMinRam ensures that no value is present for MinRam, not even an explicit nil
### GetOsDistro

`func (o *VolumeImageMetadataResponse) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *VolumeImageMetadataResponse) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *VolumeImageMetadataResponse) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *VolumeImageMetadataResponse) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *VolumeImageMetadataResponse) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *VolumeImageMetadataResponse) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetScpImageType

`func (o *VolumeImageMetadataResponse) GetScpImageType() string`

GetScpImageType returns the ScpImageType field if non-nil, zero value otherwise.

### GetScpImageTypeOk

`func (o *VolumeImageMetadataResponse) GetScpImageTypeOk() (*string, bool)`

GetScpImageTypeOk returns a tuple with the ScpImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpImageType

`func (o *VolumeImageMetadataResponse) SetScpImageType(v string)`

SetScpImageType sets ScpImageType field to given value.

### HasScpImageType

`func (o *VolumeImageMetadataResponse) HasScpImageType() bool`

HasScpImageType returns a boolean if a field has been set.

### SetScpImageTypeNil

`func (o *VolumeImageMetadataResponse) SetScpImageTypeNil(b bool)`

 SetScpImageTypeNil sets the value for ScpImageType to be an explicit nil

### UnsetScpImageType
`func (o *VolumeImageMetadataResponse) UnsetScpImageType()`

UnsetScpImageType ensures that no value is present for ScpImageType, not even an explicit nil
### GetScpK8sVersion

`func (o *VolumeImageMetadataResponse) GetScpK8sVersion() string`

GetScpK8sVersion returns the ScpK8sVersion field if non-nil, zero value otherwise.

### GetScpK8sVersionOk

`func (o *VolumeImageMetadataResponse) GetScpK8sVersionOk() (*string, bool)`

GetScpK8sVersionOk returns a tuple with the ScpK8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpK8sVersion

`func (o *VolumeImageMetadataResponse) SetScpK8sVersion(v string)`

SetScpK8sVersion sets ScpK8sVersion field to given value.

### HasScpK8sVersion

`func (o *VolumeImageMetadataResponse) HasScpK8sVersion() bool`

HasScpK8sVersion returns a boolean if a field has been set.

### SetScpK8sVersionNil

`func (o *VolumeImageMetadataResponse) SetScpK8sVersionNil(b bool)`

 SetScpK8sVersionNil sets the value for ScpK8sVersion to be an explicit nil

### UnsetScpK8sVersion
`func (o *VolumeImageMetadataResponse) UnsetScpK8sVersion()`

UnsetScpK8sVersion ensures that no value is present for ScpK8sVersion, not even an explicit nil
### GetScpOriginalImageType

`func (o *VolumeImageMetadataResponse) GetScpOriginalImageType() string`

GetScpOriginalImageType returns the ScpOriginalImageType field if non-nil, zero value otherwise.

### GetScpOriginalImageTypeOk

`func (o *VolumeImageMetadataResponse) GetScpOriginalImageTypeOk() (*string, bool)`

GetScpOriginalImageTypeOk returns a tuple with the ScpOriginalImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpOriginalImageType

`func (o *VolumeImageMetadataResponse) SetScpOriginalImageType(v string)`

SetScpOriginalImageType sets ScpOriginalImageType field to given value.

### HasScpOriginalImageType

`func (o *VolumeImageMetadataResponse) HasScpOriginalImageType() bool`

HasScpOriginalImageType returns a boolean if a field has been set.

### SetScpOriginalImageTypeNil

`func (o *VolumeImageMetadataResponse) SetScpOriginalImageTypeNil(b bool)`

 SetScpOriginalImageTypeNil sets the value for ScpOriginalImageType to be an explicit nil

### UnsetScpOriginalImageType
`func (o *VolumeImageMetadataResponse) UnsetScpOriginalImageType()`

UnsetScpOriginalImageType ensures that no value is present for ScpOriginalImageType, not even an explicit nil
### GetSignatureVerified

`func (o *VolumeImageMetadataResponse) GetSignatureVerified() bool`

GetSignatureVerified returns the SignatureVerified field if non-nil, zero value otherwise.

### GetSignatureVerifiedOk

`func (o *VolumeImageMetadataResponse) GetSignatureVerifiedOk() (*bool, bool)`

GetSignatureVerifiedOk returns a tuple with the SignatureVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerified

`func (o *VolumeImageMetadataResponse) SetSignatureVerified(v bool)`

SetSignatureVerified sets SignatureVerified field to given value.

### HasSignatureVerified

`func (o *VolumeImageMetadataResponse) HasSignatureVerified() bool`

HasSignatureVerified returns a boolean if a field has been set.

### SetSignatureVerifiedNil

`func (o *VolumeImageMetadataResponse) SetSignatureVerifiedNil(b bool)`

 SetSignatureVerifiedNil sets the value for SignatureVerified to be an explicit nil

### UnsetSignatureVerified
`func (o *VolumeImageMetadataResponse) UnsetSignatureVerified()`

UnsetSignatureVerified ensures that no value is present for SignatureVerified, not even an explicit nil
### GetSize

`func (o *VolumeImageMetadataResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VolumeImageMetadataResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VolumeImageMetadataResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *VolumeImageMetadataResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VolumeImageMetadataResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VolumeImageMetadataResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


