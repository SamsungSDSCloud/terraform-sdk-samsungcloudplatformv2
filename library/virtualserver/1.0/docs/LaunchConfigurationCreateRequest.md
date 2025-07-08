# LaunchConfigurationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | Image ID | 
**KeypairName** | **string** | Keypair name | 
**Name** | **string** | Launch Configuration name | 
**ServerTypeId** | **string** | Server type ID | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 
**UserData** | Pointer to **NullableString** |  | [optional] 
**Volumes** | Pointer to [**[]LaunchConfigurationVolume**](LaunchConfigurationVolume.md) |  | [optional] 

## Methods

### NewLaunchConfigurationCreateRequest

`func NewLaunchConfigurationCreateRequest(imageId string, keypairName string, name string, serverTypeId string, ) *LaunchConfigurationCreateRequest`

NewLaunchConfigurationCreateRequest instantiates a new LaunchConfigurationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLaunchConfigurationCreateRequestWithDefaults

`func NewLaunchConfigurationCreateRequestWithDefaults() *LaunchConfigurationCreateRequest`

NewLaunchConfigurationCreateRequestWithDefaults instantiates a new LaunchConfigurationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *LaunchConfigurationCreateRequest) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *LaunchConfigurationCreateRequest) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *LaunchConfigurationCreateRequest) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetKeypairName

`func (o *LaunchConfigurationCreateRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *LaunchConfigurationCreateRequest) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *LaunchConfigurationCreateRequest) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.


### GetName

`func (o *LaunchConfigurationCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LaunchConfigurationCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LaunchConfigurationCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServerTypeId

`func (o *LaunchConfigurationCreateRequest) GetServerTypeId() string`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *LaunchConfigurationCreateRequest) GetServerTypeIdOk() (*string, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *LaunchConfigurationCreateRequest) SetServerTypeId(v string)`

SetServerTypeId sets ServerTypeId field to given value.


### GetTags

`func (o *LaunchConfigurationCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LaunchConfigurationCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LaunchConfigurationCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LaunchConfigurationCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *LaunchConfigurationCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *LaunchConfigurationCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetUserData

`func (o *LaunchConfigurationCreateRequest) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *LaunchConfigurationCreateRequest) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *LaunchConfigurationCreateRequest) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *LaunchConfigurationCreateRequest) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *LaunchConfigurationCreateRequest) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *LaunchConfigurationCreateRequest) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetVolumes

`func (o *LaunchConfigurationCreateRequest) GetVolumes() []LaunchConfigurationVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *LaunchConfigurationCreateRequest) GetVolumesOk() (*[]LaunchConfigurationVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *LaunchConfigurationCreateRequest) SetVolumes(v []LaunchConfigurationVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *LaunchConfigurationCreateRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *LaunchConfigurationCreateRequest) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *LaunchConfigurationCreateRequest) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


