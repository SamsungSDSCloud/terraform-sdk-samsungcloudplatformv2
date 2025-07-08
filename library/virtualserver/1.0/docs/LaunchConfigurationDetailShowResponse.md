# LaunchConfigurationDetailShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AutoScalingGroupCount** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
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

### NewLaunchConfigurationDetailShowResponse

`func NewLaunchConfigurationDetailShowResponse(accountId string, createdAt time.Time, createdBy string, id string, imageId string, imageName string, keypairName string, modifiedAt time.Time, modifiedBy string, name string, serverTypeId string, state string, ) *LaunchConfigurationDetailShowResponse`

NewLaunchConfigurationDetailShowResponse instantiates a new LaunchConfigurationDetailShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchConfigurationDetailShowResponseWithDefaults

`func NewLaunchConfigurationDetailShowResponseWithDefaults() *LaunchConfigurationDetailShowResponse`

NewLaunchConfigurationDetailShowResponseWithDefaults instantiates a new LaunchConfigurationDetailShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LaunchConfigurationDetailShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LaunchConfigurationDetailShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LaunchConfigurationDetailShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAutoScalingGroupCount

`func (o *LaunchConfigurationDetailShowResponse) GetAutoScalingGroupCount() int32`

GetAutoScalingGroupCount returns the AutoScalingGroupCount field if non-nil, zero value otherwise.

### GetAutoScalingGroupCountOk

`func (o *LaunchConfigurationDetailShowResponse) GetAutoScalingGroupCountOk() (*int32, bool)`

GetAutoScalingGroupCountOk returns a tuple with the AutoScalingGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScalingGroupCount

`func (o *LaunchConfigurationDetailShowResponse) SetAutoScalingGroupCount(v int32)`

SetAutoScalingGroupCount sets AutoScalingGroupCount field to given value.

### HasAutoScalingGroupCount

`func (o *LaunchConfigurationDetailShowResponse) HasAutoScalingGroupCount() bool`

HasAutoScalingGroupCount returns a boolean if a field has been set.

### SetAutoScalingGroupCountNil

`func (o *LaunchConfigurationDetailShowResponse) SetAutoScalingGroupCountNil(b bool)`

 SetAutoScalingGroupCountNil sets the value for AutoScalingGroupCount to be an explicit nil

### UnsetAutoScalingGroupCount
`func (o *LaunchConfigurationDetailShowResponse) UnsetAutoScalingGroupCount()`

UnsetAutoScalingGroupCount ensures that no value is present for AutoScalingGroupCount, not even an explicit nil
### GetCreatedAt

`func (o *LaunchConfigurationDetailShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LaunchConfigurationDetailShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LaunchConfigurationDetailShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LaunchConfigurationDetailShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LaunchConfigurationDetailShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LaunchConfigurationDetailShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *LaunchConfigurationDetailShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LaunchConfigurationDetailShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LaunchConfigurationDetailShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetImageId

`func (o *LaunchConfigurationDetailShowResponse) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *LaunchConfigurationDetailShowResponse) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *LaunchConfigurationDetailShowResponse) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetImageName

`func (o *LaunchConfigurationDetailShowResponse) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *LaunchConfigurationDetailShowResponse) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *LaunchConfigurationDetailShowResponse) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetKeypairName

`func (o *LaunchConfigurationDetailShowResponse) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *LaunchConfigurationDetailShowResponse) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *LaunchConfigurationDetailShowResponse) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.


### GetModifiedAt

`func (o *LaunchConfigurationDetailShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LaunchConfigurationDetailShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LaunchConfigurationDetailShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LaunchConfigurationDetailShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LaunchConfigurationDetailShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LaunchConfigurationDetailShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LaunchConfigurationDetailShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LaunchConfigurationDetailShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LaunchConfigurationDetailShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *LaunchConfigurationDetailShowResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *LaunchConfigurationDetailShowResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *LaunchConfigurationDetailShowResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *LaunchConfigurationDetailShowResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *LaunchConfigurationDetailShowResponse) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *LaunchConfigurationDetailShowResponse) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetServerTypeId

`func (o *LaunchConfigurationDetailShowResponse) GetServerTypeId() string`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *LaunchConfigurationDetailShowResponse) GetServerTypeIdOk() (*string, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *LaunchConfigurationDetailShowResponse) SetServerTypeId(v string)`

SetServerTypeId sets ServerTypeId field to given value.


### GetState

`func (o *LaunchConfigurationDetailShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LaunchConfigurationDetailShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LaunchConfigurationDetailShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetUserData

`func (o *LaunchConfigurationDetailShowResponse) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *LaunchConfigurationDetailShowResponse) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *LaunchConfigurationDetailShowResponse) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *LaunchConfigurationDetailShowResponse) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *LaunchConfigurationDetailShowResponse) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *LaunchConfigurationDetailShowResponse) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVolumes

`func (o *LaunchConfigurationDetailShowResponse) GetVolumes() []LaunchConfigurationVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *LaunchConfigurationDetailShowResponse) GetVolumesOk() (*[]LaunchConfigurationVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *LaunchConfigurationDetailShowResponse) SetVolumes(v []LaunchConfigurationVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *LaunchConfigurationDetailShowResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *LaunchConfigurationDetailShowResponse) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *LaunchConfigurationDetailShowResponse) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


