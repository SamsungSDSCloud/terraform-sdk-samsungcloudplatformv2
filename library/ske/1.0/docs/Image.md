# Image

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomImageName** | Pointer to **NullableString** |  | [optional] 
**Os** | **string** | Image OS | 
**OsVersion** | **string** | Image OS Version | 

## Methods

### NewImage

`func NewImage(os string, osVersion string, ) *Image`

NewImage instantiates a new Image object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageWithDefaults

`func NewImageWithDefaults() *Image`

NewImageWithDefaults instantiates a new Image object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomImageName

`func (o *Image) GetCustomImageName() string`

GetCustomImageName returns the CustomImageName field if non-nil, zero value otherwise.

### GetCustomImageNameOk

`func (o *Image) GetCustomImageNameOk() (*string, bool)`

GetCustomImageNameOk returns a tuple with the CustomImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageName

`func (o *Image) SetCustomImageName(v string)`

SetCustomImageName sets CustomImageName field to given value.

### HasCustomImageName

`func (o *Image) HasCustomImageName() bool`

HasCustomImageName returns a boolean if a field has been set.

### SetCustomImageNameNil

`func (o *Image) SetCustomImageNameNil(b bool)`

 SetCustomImageNameNil sets the value for CustomImageName to be an explicit nil

### UnsetCustomImageName
`func (o *Image) UnsetCustomImageName()`

UnsetCustomImageName ensures that no value is present for CustomImageName, not even an explicit nil
### GetOs

`func (o *Image) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Image) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Image) SetOs(v string)`

SetOs sets Os field to given value.


### GetOsVersion

`func (o *Image) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *Image) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *Image) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


