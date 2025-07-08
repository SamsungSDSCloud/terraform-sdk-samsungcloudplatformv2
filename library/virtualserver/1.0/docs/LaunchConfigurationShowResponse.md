# LaunchConfigurationShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AutoScalingGroupCount** | Pointer to **NullableInt32** |  | [optional] 
**BootDiskSize** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**ExtraDiskSize** | Pointer to **NullableInt32** |  | [optional] 
**Id** | **string** | ID | 
**ImageId** | **string** | Image ID | 
**ImageName** | **string** | Image name | 
**KeypairName** | **string** | Keypair name | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Launch Configuration name | 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**ServerTypeId** | **string** | Server type ID | 
**State** | **string** | Launch Configuration state | 
**UserData** | Pointer to **NullableString** |  | [optional] 
**Volumes** | Pointer to [**[]LaunchConfigurationVolume**](LaunchConfigurationVolume.md) |  | [optional] 

## Methods

### NewLaunchConfigurationShowResponse

`func NewLaunchConfigurationShowResponse(accountId string, createdAt time.Time, createdBy string, id string, imageId string, imageName string, keypairName string, modifiedAt time.Time, modifiedBy string, name string, serverTypeId string, state string, ) *LaunchConfigurationShowResponse`

NewLaunchConfigurationShowResponse instantiates a new LaunchConfigurationShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchConfigurationShowResponseWithDefaults

`func NewLaunchConfigurationShowResponseWithDefaults() *LaunchConfigurationShowResponse`

NewLaunchConfigurationShowResponseWithDefaults instantiates a new LaunchConfigurationShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LaunchConfigurationShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LaunchConfigurationShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LaunchConfigurationShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAutoScalingGroupCount

`func (o *LaunchConfigurationShowResponse) GetAutoScalingGroupCount() int32`

GetAutoScalingGroupCount returns the AutoScalingGroupCount field if non-nil, zero value otherwise.

### GetAutoScalingGroupCountOk

`func (o *LaunchConfigurationShowResponse) GetAutoScalingGroupCountOk() (*int32, bool)`

GetAutoScalingGroupCountOk returns a tuple with the AutoScalingGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroupCount

`func (o *LaunchConfigurationShowResponse) SetAutoScalingGroupCount(v int32)`

SetAutoScalingGroupCount sets AutoScalingGroupCount field to given value.

### HasAutoScalingGroupCount

`func (o *LaunchConfigurationShowResponse) HasAutoScalingGroupCount() bool`

HasAutoScalingGroupCount returns a boolean if a field has been set.

### SetAutoScalingGroupCountNil

`func (o *LaunchConfigurationShowResponse) SetAutoScalingGroupCountNil(b bool)`

 SetAutoScalingGroupCountNil sets the value for AutoScalingGroupCount to be an explicit nil

### UnsetAutoScalingGroupCount
`func (o *LaunchConfigurationShowResponse) UnsetAutoScalingGroupCount()`

UnsetAutoScalingGroupCount ensures that no value is present for AutoScalingGroupCount, not even an explicit nil
### GetBootDiskSize

`func (o *LaunchConfigurationShowResponse) GetBootDiskSize() int32`

GetBootDiskSize returns the BootDiskSize field if non-nil, zero value otherwise.

### GetBootDiskSizeOk

`func (o *LaunchConfigurationShowResponse) GetBootDiskSizeOk() (*int32, bool)`

GetBootDiskSizeOk returns a tuple with the BootDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiskSize

`func (o *LaunchConfigurationShowResponse) SetBootDiskSize(v int32)`

SetBootDiskSize sets BootDiskSize field to given value.

### HasBootDiskSize

`func (o *LaunchConfigurationShowResponse) HasBootDiskSize() bool`

HasBootDiskSize returns a boolean if a field has been set.

### SetBootDiskSizeNil

`func (o *LaunchConfigurationShowResponse) SetBootDiskSizeNil(b bool)`

 SetBootDiskSizeNil sets the value for BootDiskSize to be an explicit nil

### UnsetBootDiskSize
`func (o *LaunchConfigurationShowResponse) UnsetBootDiskSize()`

UnsetBootDiskSize ensures that no value is present for BootDiskSize, not even an explicit nil
### GetCreatedAt

`func (o *LaunchConfigurationShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LaunchConfigurationShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LaunchConfigurationShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LaunchConfigurationShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LaunchConfigurationShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LaunchConfigurationShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetExtraDiskSize

`func (o *LaunchConfigurationShowResponse) GetExtraDiskSize() int32`

GetExtraDiskSize returns the ExtraDiskSize field if non-nil, zero value otherwise.

### GetExtraDiskSizeOk

`func (o *LaunchConfigurationShowResponse) GetExtraDiskSizeOk() (*int32, bool)`

GetExtraDiskSizeOk returns a tuple with the ExtraDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDiskSize

`func (o *LaunchConfigurationShowResponse) SetExtraDiskSize(v int32)`

SetExtraDiskSize sets ExtraDiskSize field to given value.

### HasExtraDiskSize

`func (o *LaunchConfigurationShowResponse) HasExtraDiskSize() bool`

HasExtraDiskSize returns a boolean if a field has been set.

### SetExtraDiskSizeNil

`func (o *LaunchConfigurationShowResponse) SetExtraDiskSizeNil(b bool)`

 SetExtraDiskSizeNil sets the value for ExtraDiskSize to be an explicit nil

### UnsetExtraDiskSize
`func (o *LaunchConfigurationShowResponse) UnsetExtraDiskSize()`

UnsetExtraDiskSize ensures that no value is present for ExtraDiskSize, not even an explicit nil
### GetId

`func (o *LaunchConfigurationShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LaunchConfigurationShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LaunchConfigurationShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *LaunchConfigurationShowResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *LaunchConfigurationShowResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *LaunchConfigurationShowResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetImageName

`func (o *LaunchConfigurationShowResponse) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *LaunchConfigurationShowResponse) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *LaunchConfigurationShowResponse) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetKeypairName

`func (o *LaunchConfigurationShowResponse) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *LaunchConfigurationShowResponse) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *LaunchConfigurationShowResponse) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.


### GetModifiedAt

`func (o *LaunchConfigurationShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LaunchConfigurationShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LaunchConfigurationShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LaunchConfigurationShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LaunchConfigurationShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LaunchConfigurationShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LaunchConfigurationShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LaunchConfigurationShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LaunchConfigurationShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *LaunchConfigurationShowResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LaunchConfigurationShowResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LaunchConfigurationShowResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LaunchConfigurationShowResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LaunchConfigurationShowResponse) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LaunchConfigurationShowResponse) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetServerTypeId

`func (o *LaunchConfigurationShowResponse) GetServerTypeId() string`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *LaunchConfigurationShowResponse) GetServerTypeIdOk() (*string, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *LaunchConfigurationShowResponse) SetServerTypeId(v string)`

SetServerTypeId sets ServerTypeId field to given value.


### GetState

`func (o *LaunchConfigurationShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LaunchConfigurationShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LaunchConfigurationShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetUserData

`func (o *LaunchConfigurationShowResponse) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *LaunchConfigurationShowResponse) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *LaunchConfigurationShowResponse) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *LaunchConfigurationShowResponse) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *LaunchConfigurationShowResponse) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *LaunchConfigurationShowResponse) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVolumes

`func (o *LaunchConfigurationShowResponse) GetVolumes() []LaunchConfigurationVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *LaunchConfigurationShowResponse) GetVolumesOk() (*[]LaunchConfigurationVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *LaunchConfigurationShowResponse) SetVolumes(v []LaunchConfigurationVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *LaunchConfigurationShowResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *LaunchConfigurationShowResponse) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *LaunchConfigurationShowResponse) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


