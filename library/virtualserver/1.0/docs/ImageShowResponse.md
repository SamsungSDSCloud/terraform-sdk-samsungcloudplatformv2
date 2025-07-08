# ImageShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | Pointer to **NullableString** |  | [optional] 
**ContainerFormat** | **string** | Container format | 
**CreatedAt** | **string** | Created at | 
**DiskFormat** | **string** | Disk format | 
**File** | **string** | File | 
**Id** | **string** | Image ID | 
**MinDisk** | **int32** | Minimum disk | 
**MinRam** | **int32** | Minimum RAM | 
**Name** | **string** | Name | 
**OsDistro** | Pointer to **NullableString** |  | [optional] 
**OsHashAlgo** | Pointer to **NullableString** |  | [optional] 
**OsHashValue** | Pointer to **NullableString** |  | [optional] 
**OsHidden** | **bool** | OS hidden | 
**Owner** | **string** | Owner project ID | 
**OwnerAccountName** | Pointer to **NullableString** |  | [optional] 
**OwnerUserName** | Pointer to **NullableString** |  | [optional] 
**Protected** | **bool** | Protected | 
**RootDeviceName** | Pointer to **NullableString** |  | [optional] 
**ScpImageType** | Pointer to **NullableString** |  | [optional] 
**ScpK8sVersion** | Pointer to **NullableString** |  | [optional] 
**ScpOriginalImageType** | Pointer to **NullableString** |  | [optional] 
**ScpOsBuildVersion** | Pointer to **NullableString** |  | [optional] 
**ScpOsVersion** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableInt64** |  | [optional] 
**Status** | **string** | Status | 
**UpdatedAt** | **string** | Updated at | 
**Url** | Pointer to **NullableString** |  | [optional] 
**VirtualSize** | Pointer to **NullableInt64** |  | [optional] 
**Visibility** | **string** | Visibility | 
**Volumes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewImageShowResponse

`func NewImageShowResponse(containerFormat string, createdAt string, diskFormat string, file string, id string, minDisk int32, minRam int32, name string, osHidden bool, owner string, protected bool, status string, updatedAt string, visibility string, ) *ImageShowResponse`

NewImageShowResponse instantiates a new ImageShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageShowResponseWithDefaults

`func NewImageShowResponseWithDefaults() *ImageShowResponse`

NewImageShowResponseWithDefaults instantiates a new ImageShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *ImageShowResponse) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ImageShowResponse) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ImageShowResponse) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ImageShowResponse) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *ImageShowResponse) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *ImageShowResponse) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetContainerFormat

`func (o *ImageShowResponse) GetContainerFormat() string`

GetContainerFormat returns the ContainerFormat field if non-nil, zero value otherwise.

### GetContainerFormatOk

`func (o *ImageShowResponse) GetContainerFormatOk() (*string, bool)`

GetContainerFormatOk returns a tuple with the ContainerFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerFormat

`func (o *ImageShowResponse) SetContainerFormat(v string)`

SetContainerFormat sets ContainerFormat field to given value.


### GetCreatedAt

`func (o *ImageShowResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ImageShowResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ImageShowResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDiskFormat

`func (o *ImageShowResponse) GetDiskFormat() string`

GetDiskFormat returns the DiskFormat field if non-nil, zero value otherwise.

### GetDiskFormatOk

`func (o *ImageShowResponse) GetDiskFormatOk() (*string, bool)`

GetDiskFormatOk returns a tuple with the DiskFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskFormat

`func (o *ImageShowResponse) SetDiskFormat(v string)`

SetDiskFormat sets DiskFormat field to given value.


### GetFile

`func (o *ImageShowResponse) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ImageShowResponse) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ImageShowResponse) SetFile(v string)`

SetFile sets File field to given value.


### GetId

`func (o *ImageShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMinDisk

`func (o *ImageShowResponse) GetMinDisk() int32`

GetMinDisk returns the MinDisk field if non-nil, zero value otherwise.

### GetMinDiskOk

`func (o *ImageShowResponse) GetMinDiskOk() (*int32, bool)`

GetMinDiskOk returns a tuple with the MinDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDisk

`func (o *ImageShowResponse) SetMinDisk(v int32)`

SetMinDisk sets MinDisk field to given value.


### GetMinRam

`func (o *ImageShowResponse) GetMinRam() int32`

GetMinRam returns the MinRam field if non-nil, zero value otherwise.

### GetMinRamOk

`func (o *ImageShowResponse) GetMinRamOk() (*int32, bool)`

GetMinRamOk returns a tuple with the MinRam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRam

`func (o *ImageShowResponse) SetMinRam(v int32)`

SetMinRam sets MinRam field to given value.


### GetName

`func (o *ImageShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOsDistro

`func (o *ImageShowResponse) GetOsDistro() string`

GetOsDistro returns the OsDistro field if non-nil, zero value otherwise.

### GetOsDistroOk

`func (o *ImageShowResponse) GetOsDistroOk() (*string, bool)`

GetOsDistroOk returns a tuple with the OsDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsDistro

`func (o *ImageShowResponse) SetOsDistro(v string)`

SetOsDistro sets OsDistro field to given value.

### HasOsDistro

`func (o *ImageShowResponse) HasOsDistro() bool`

HasOsDistro returns a boolean if a field has been set.

### SetOsDistroNil

`func (o *ImageShowResponse) SetOsDistroNil(b bool)`

 SetOsDistroNil sets the value for OsDistro to be an explicit nil

### UnsetOsDistro
`func (o *ImageShowResponse) UnsetOsDistro()`

UnsetOsDistro ensures that no value is present for OsDistro, not even an explicit nil
### GetOsHashAlgo

`func (o *ImageShowResponse) GetOsHashAlgo() string`

GetOsHashAlgo returns the OsHashAlgo field if non-nil, zero value otherwise.

### GetOsHashAlgoOk

`func (o *ImageShowResponse) GetOsHashAlgoOk() (*string, bool)`

GetOsHashAlgoOk returns a tuple with the OsHashAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsHashAlgo

`func (o *ImageShowResponse) SetOsHashAlgo(v string)`

SetOsHashAlgo sets OsHashAlgo field to given value.

### HasOsHashAlgo

`func (o *ImageShowResponse) HasOsHashAlgo() bool`

HasOsHashAlgo returns a boolean if a field has been set.

### SetOsHashAlgoNil

`func (o *ImageShowResponse) SetOsHashAlgoNil(b bool)`

 SetOsHashAlgoNil sets the value for OsHashAlgo to be an explicit nil

### UnsetOsHashAlgo
`func (o *ImageShowResponse) UnsetOsHashAlgo()`

UnsetOsHashAlgo ensures that no value is present for OsHashAlgo, not even an explicit nil
### GetOsHashValue

`func (o *ImageShowResponse) GetOsHashValue() string`

GetOsHashValue returns the OsHashValue field if non-nil, zero value otherwise.

### GetOsHashValueOk

`func (o *ImageShowResponse) GetOsHashValueOk() (*string, bool)`

GetOsHashValueOk returns a tuple with the OsHashValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsHashValue

`func (o *ImageShowResponse) SetOsHashValue(v string)`

SetOsHashValue sets OsHashValue field to given value.

### HasOsHashValue

`func (o *ImageShowResponse) HasOsHashValue() bool`

HasOsHashValue returns a boolean if a field has been set.

### SetOsHashValueNil

`func (o *ImageShowResponse) SetOsHashValueNil(b bool)`

 SetOsHashValueNil sets the value for OsHashValue to be an explicit nil

### UnsetOsHashValue
`func (o *ImageShowResponse) UnsetOsHashValue()`

UnsetOsHashValue ensures that no value is present for OsHashValue, not even an explicit nil
### GetOsHidden

`func (o *ImageShowResponse) GetOsHidden() bool`

GetOsHidden returns the OsHidden field if non-nil, zero value otherwise.

### GetOsHiddenOk

`func (o *ImageShowResponse) GetOsHiddenOk() (*bool, bool)`

GetOsHiddenOk returns a tuple with the OsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsHidden

`func (o *ImageShowResponse) SetOsHidden(v bool)`

SetOsHidden sets OsHidden field to given value.


### GetOwner

`func (o *ImageShowResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ImageShowResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ImageShowResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetOwnerAccountName

`func (o *ImageShowResponse) GetOwnerAccountName() string`

GetOwnerAccountName returns the OwnerAccountName field if non-nil, zero value otherwise.

### GetOwnerAccountNameOk

`func (o *ImageShowResponse) GetOwnerAccountNameOk() (*string, bool)`

GetOwnerAccountNameOk returns a tuple with the OwnerAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAccountName

`func (o *ImageShowResponse) SetOwnerAccountName(v string)`

SetOwnerAccountName sets OwnerAccountName field to given value.

### HasOwnerAccountName

`func (o *ImageShowResponse) HasOwnerAccountName() bool`

HasOwnerAccountName returns a boolean if a field has been set.

### SetOwnerAccountNameNil

`func (o *ImageShowResponse) SetOwnerAccountNameNil(b bool)`

 SetOwnerAccountNameNil sets the value for OwnerAccountName to be an explicit nil

### UnsetOwnerAccountName
`func (o *ImageShowResponse) UnsetOwnerAccountName()`

UnsetOwnerAccountName ensures that no value is present for OwnerAccountName, not even an explicit nil
### GetOwnerUserName

`func (o *ImageShowResponse) GetOwnerUserName() string`

GetOwnerUserName returns the OwnerUserName field if non-nil, zero value otherwise.

### GetOwnerUserNameOk

`func (o *ImageShowResponse) GetOwnerUserNameOk() (*string, bool)`

GetOwnerUserNameOk returns a tuple with the OwnerUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUserName

`func (o *ImageShowResponse) SetOwnerUserName(v string)`

SetOwnerUserName sets OwnerUserName field to given value.

### HasOwnerUserName

`func (o *ImageShowResponse) HasOwnerUserName() bool`

HasOwnerUserName returns a boolean if a field has been set.

### SetOwnerUserNameNil

`func (o *ImageShowResponse) SetOwnerUserNameNil(b bool)`

 SetOwnerUserNameNil sets the value for OwnerUserName to be an explicit nil

### UnsetOwnerUserName
`func (o *ImageShowResponse) UnsetOwnerUserName()`

UnsetOwnerUserName ensures that no value is present for OwnerUserName, not even an explicit nil
### GetProtected

`func (o *ImageShowResponse) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *ImageShowResponse) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *ImageShowResponse) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetRootDeviceName

`func (o *ImageShowResponse) GetRootDeviceName() string`

GetRootDeviceName returns the RootDeviceName field if non-nil, zero value otherwise.

### GetRootDeviceNameOk

`func (o *ImageShowResponse) GetRootDeviceNameOk() (*string, bool)`

GetRootDeviceNameOk returns a tuple with the RootDeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDeviceName

`func (o *ImageShowResponse) SetRootDeviceName(v string)`

SetRootDeviceName sets RootDeviceName field to given value.

### HasRootDeviceName

`func (o *ImageShowResponse) HasRootDeviceName() bool`

HasRootDeviceName returns a boolean if a field has been set.

### SetRootDeviceNameNil

`func (o *ImageShowResponse) SetRootDeviceNameNil(b bool)`

 SetRootDeviceNameNil sets the value for RootDeviceName to be an explicit nil

### UnsetRootDeviceName
`func (o *ImageShowResponse) UnsetRootDeviceName()`

UnsetRootDeviceName ensures that no value is present for RootDeviceName, not even an explicit nil
### GetScpImageType

`func (o *ImageShowResponse) GetScpImageType() string`

GetScpImageType returns the ScpImageType field if non-nil, zero value otherwise.

### GetScpImageTypeOk

`func (o *ImageShowResponse) GetScpImageTypeOk() (*string, bool)`

GetScpImageTypeOk returns a tuple with the ScpImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpImageType

`func (o *ImageShowResponse) SetScpImageType(v string)`

SetScpImageType sets ScpImageType field to given value.

### HasScpImageType

`func (o *ImageShowResponse) HasScpImageType() bool`

HasScpImageType returns a boolean if a field has been set.

### SetScpImageTypeNil

`func (o *ImageShowResponse) SetScpImageTypeNil(b bool)`

 SetScpImageTypeNil sets the value for ScpImageType to be an explicit nil

### UnsetScpImageType
`func (o *ImageShowResponse) UnsetScpImageType()`

UnsetScpImageType ensures that no value is present for ScpImageType, not even an explicit nil
### GetScpK8sVersion

`func (o *ImageShowResponse) GetScpK8sVersion() string`

GetScpK8sVersion returns the ScpK8sVersion field if non-nil, zero value otherwise.

### GetScpK8sVersionOk

`func (o *ImageShowResponse) GetScpK8sVersionOk() (*string, bool)`

GetScpK8sVersionOk returns a tuple with the ScpK8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpK8sVersion

`func (o *ImageShowResponse) SetScpK8sVersion(v string)`

SetScpK8sVersion sets ScpK8sVersion field to given value.

### HasScpK8sVersion

`func (o *ImageShowResponse) HasScpK8sVersion() bool`

HasScpK8sVersion returns a boolean if a field has been set.

### SetScpK8sVersionNil

`func (o *ImageShowResponse) SetScpK8sVersionNil(b bool)`

 SetScpK8sVersionNil sets the value for ScpK8sVersion to be an explicit nil

### UnsetScpK8sVersion
`func (o *ImageShowResponse) UnsetScpK8sVersion()`

UnsetScpK8sVersion ensures that no value is present for ScpK8sVersion, not even an explicit nil
### GetScpOriginalImageType

`func (o *ImageShowResponse) GetScpOriginalImageType() string`

GetScpOriginalImageType returns the ScpOriginalImageType field if non-nil, zero value otherwise.

### GetScpOriginalImageTypeOk

`func (o *ImageShowResponse) GetScpOriginalImageTypeOk() (*string, bool)`

GetScpOriginalImageTypeOk returns a tuple with the ScpOriginalImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpOriginalImageType

`func (o *ImageShowResponse) SetScpOriginalImageType(v string)`

SetScpOriginalImageType sets ScpOriginalImageType field to given value.

### HasScpOriginalImageType

`func (o *ImageShowResponse) HasScpOriginalImageType() bool`

HasScpOriginalImageType returns a boolean if a field has been set.

### SetScpOriginalImageTypeNil

`func (o *ImageShowResponse) SetScpOriginalImageTypeNil(b bool)`

 SetScpOriginalImageTypeNil sets the value for ScpOriginalImageType to be an explicit nil

### UnsetScpOriginalImageType
`func (o *ImageShowResponse) UnsetScpOriginalImageType()`

UnsetScpOriginalImageType ensures that no value is present for ScpOriginalImageType, not even an explicit nil
### GetScpOsBuildVersion

`func (o *ImageShowResponse) GetScpOsBuildVersion() string`

GetScpOsBuildVersion returns the ScpOsBuildVersion field if non-nil, zero value otherwise.

### GetScpOsBuildVersionOk

`func (o *ImageShowResponse) GetScpOsBuildVersionOk() (*string, bool)`

GetScpOsBuildVersionOk returns a tuple with the ScpOsBuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpOsBuildVersion

`func (o *ImageShowResponse) SetScpOsBuildVersion(v string)`

SetScpOsBuildVersion sets ScpOsBuildVersion field to given value.

### HasScpOsBuildVersion

`func (o *ImageShowResponse) HasScpOsBuildVersion() bool`

HasScpOsBuildVersion returns a boolean if a field has been set.

### SetScpOsBuildVersionNil

`func (o *ImageShowResponse) SetScpOsBuildVersionNil(b bool)`

 SetScpOsBuildVersionNil sets the value for ScpOsBuildVersion to be an explicit nil

### UnsetScpOsBuildVersion
`func (o *ImageShowResponse) UnsetScpOsBuildVersion()`

UnsetScpOsBuildVersion ensures that no value is present for ScpOsBuildVersion, not even an explicit nil
### GetScpOsVersion

`func (o *ImageShowResponse) GetScpOsVersion() string`

GetScpOsVersion returns the ScpOsVersion field if non-nil, zero value otherwise.

### GetScpOsVersionOk

`func (o *ImageShowResponse) GetScpOsVersionOk() (*string, bool)`

GetScpOsVersionOk returns a tuple with the ScpOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScpOsVersion

`func (o *ImageShowResponse) SetScpOsVersion(v string)`

SetScpOsVersion sets ScpOsVersion field to given value.

### HasScpOsVersion

`func (o *ImageShowResponse) HasScpOsVersion() bool`

HasScpOsVersion returns a boolean if a field has been set.

### SetScpOsVersionNil

`func (o *ImageShowResponse) SetScpOsVersionNil(b bool)`

 SetScpOsVersionNil sets the value for ScpOsVersion to be an explicit nil

### UnsetScpOsVersion
`func (o *ImageShowResponse) UnsetScpOsVersion()`

UnsetScpOsVersion ensures that no value is present for ScpOsVersion, not even an explicit nil
### GetSize

`func (o *ImageShowResponse) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ImageShowResponse) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ImageShowResponse) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ImageShowResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ImageShowResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ImageShowResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetStatus

`func (o *ImageShowResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageShowResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageShowResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *ImageShowResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ImageShowResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ImageShowResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *ImageShowResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ImageShowResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ImageShowResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ImageShowResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ImageShowResponse) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ImageShowResponse) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetVirtualSize

`func (o *ImageShowResponse) GetVirtualSize() int64`

GetVirtualSize returns the VirtualSize field if non-nil, zero value otherwise.

### GetVirtualSizeOk

`func (o *ImageShowResponse) GetVirtualSizeOk() (*int64, bool)`

GetVirtualSizeOk returns a tuple with the VirtualSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualSize

`func (o *ImageShowResponse) SetVirtualSize(v int64)`

SetVirtualSize sets VirtualSize field to given value.

### HasVirtualSize

`func (o *ImageShowResponse) HasVirtualSize() bool`

HasVirtualSize returns a boolean if a field has been set.

### SetVirtualSizeNil

`func (o *ImageShowResponse) SetVirtualSizeNil(b bool)`

 SetVirtualSizeNil sets the value for VirtualSize to be an explicit nil

### UnsetVirtualSize
`func (o *ImageShowResponse) UnsetVirtualSize()`

UnsetVirtualSize ensures that no value is present for VirtualSize, not even an explicit nil
### GetVisibility

`func (o *ImageShowResponse) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ImageShowResponse) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ImageShowResponse) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetVolumes

`func (o *ImageShowResponse) GetVolumes() string`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ImageShowResponse) GetVolumesOk() (*string, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ImageShowResponse) SetVolumes(v string)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ImageShowResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *ImageShowResponse) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *ImageShowResponse) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


