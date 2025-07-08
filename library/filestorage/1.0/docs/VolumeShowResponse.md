# VolumeShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CifsId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**EncryptionEnabled** | **bool** | Volume Encryption Enabled | 
**EndpointPath** | Pointer to **NullableString** |  | [optional] 
**FileUnitRecoveryEnabled** | Pointer to **NullableBool** |  | [optional] 
**Id** | **string** | Volume ID | 
**Name** | **string** | Volume Name | 
**Path** | Pointer to **NullableString** |  | [optional] 
**Protocol** | **string** | Protocol | 
**Purpose** | **string** | Purpose | 
**State** | **string** | Volume State | 
**TypeId** | **string** | Volume Type ID | 
**TypeName** | **string** | Volume Type Name | 
**Usage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewVolumeShowResponse

`func NewVolumeShowResponse(accountId string, createdAt time.Time, encryptionEnabled bool, id string, name string, protocol string, purpose string, state string, typeId string, typeName string, ) *VolumeShowResponse`

NewVolumeShowResponse instantiates a new VolumeShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeShowResponseWithDefaults

`func NewVolumeShowResponseWithDefaults() *VolumeShowResponse`

NewVolumeShowResponseWithDefaults instantiates a new VolumeShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VolumeShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VolumeShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VolumeShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCifsId

`func (o *VolumeShowResponse) GetCifsId() string`

GetCifsId returns the CifsId field if non-nil, zero value otherwise.

### GetCifsIdOk

`func (o *VolumeShowResponse) GetCifsIdOk() (*string, bool)`

GetCifsIdOk returns a tuple with the CifsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCifsId

`func (o *VolumeShowResponse) SetCifsId(v string)`

SetCifsId sets CifsId field to given value.

### HasCifsId

`func (o *VolumeShowResponse) HasCifsId() bool`

HasCifsId returns a boolean if a field has been set.

### SetCifsIdNil

`func (o *VolumeShowResponse) SetCifsIdNil(b bool)`

 SetCifsIdNil sets the value for CifsId to be an explicit nil

### UnsetCifsId
`func (o *VolumeShowResponse) UnsetCifsId()`

UnsetCifsId ensures that no value is present for CifsId, not even an explicit nil
### GetCreatedAt

`func (o *VolumeShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEncryptionEnabled

`func (o *VolumeShowResponse) GetEncryptionEnabled() bool`

GetEncryptionEnabled returns the EncryptionEnabled field if non-nil, zero value otherwise.

### GetEncryptionEnabledOk

`func (o *VolumeShowResponse) GetEncryptionEnabledOk() (*bool, bool)`

GetEncryptionEnabledOk returns a tuple with the EncryptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionEnabled

`func (o *VolumeShowResponse) SetEncryptionEnabled(v bool)`

SetEncryptionEnabled sets EncryptionEnabled field to given value.


### GetEndpointPath

`func (o *VolumeShowResponse) GetEndpointPath() string`

GetEndpointPath returns the EndpointPath field if non-nil, zero value otherwise.

### GetEndpointPathOk

`func (o *VolumeShowResponse) GetEndpointPathOk() (*string, bool)`

GetEndpointPathOk returns a tuple with the EndpointPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointPath

`func (o *VolumeShowResponse) SetEndpointPath(v string)`

SetEndpointPath sets EndpointPath field to given value.

### HasEndpointPath

`func (o *VolumeShowResponse) HasEndpointPath() bool`

HasEndpointPath returns a boolean if a field has been set.

### SetEndpointPathNil

`func (o *VolumeShowResponse) SetEndpointPathNil(b bool)`

 SetEndpointPathNil sets the value for EndpointPath to be an explicit nil

### UnsetEndpointPath
`func (o *VolumeShowResponse) UnsetEndpointPath()`

UnsetEndpointPath ensures that no value is present for EndpointPath, not even an explicit nil
### GetFileUnitRecoveryEnabled

`func (o *VolumeShowResponse) GetFileUnitRecoveryEnabled() bool`

GetFileUnitRecoveryEnabled returns the FileUnitRecoveryEnabled field if non-nil, zero value otherwise.

### GetFileUnitRecoveryEnabledOk

`func (o *VolumeShowResponse) GetFileUnitRecoveryEnabledOk() (*bool, bool)`

GetFileUnitRecoveryEnabledOk returns a tuple with the FileUnitRecoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUnitRecoveryEnabled

`func (o *VolumeShowResponse) SetFileUnitRecoveryEnabled(v bool)`

SetFileUnitRecoveryEnabled sets FileUnitRecoveryEnabled field to given value.

### HasFileUnitRecoveryEnabled

`func (o *VolumeShowResponse) HasFileUnitRecoveryEnabled() bool`

HasFileUnitRecoveryEnabled returns a boolean if a field has been set.

### SetFileUnitRecoveryEnabledNil

`func (o *VolumeShowResponse) SetFileUnitRecoveryEnabledNil(b bool)`

 SetFileUnitRecoveryEnabledNil sets the value for FileUnitRecoveryEnabled to be an explicit nil

### UnsetFileUnitRecoveryEnabled
`func (o *VolumeShowResponse) UnsetFileUnitRecoveryEnabled()`

UnsetFileUnitRecoveryEnabled ensures that no value is present for FileUnitRecoveryEnabled, not even an explicit nil
### GetId

`func (o *VolumeShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VolumeShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VolumeShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VolumeShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *VolumeShowResponse) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VolumeShowResponse) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VolumeShowResponse) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VolumeShowResponse) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *VolumeShowResponse) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *VolumeShowResponse) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetProtocol

`func (o *VolumeShowResponse) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *VolumeShowResponse) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *VolumeShowResponse) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetPurpose

`func (o *VolumeShowResponse) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *VolumeShowResponse) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *VolumeShowResponse) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetState

`func (o *VolumeShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VolumeShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VolumeShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetTypeId

`func (o *VolumeShowResponse) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *VolumeShowResponse) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *VolumeShowResponse) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.


### GetTypeName

`func (o *VolumeShowResponse) GetTypeName() string`

GetTypeName returns the TypeName field if non-nil, zero value otherwise.

### GetTypeNameOk

`func (o *VolumeShowResponse) GetTypeNameOk() (*string, bool)`

GetTypeNameOk returns a tuple with the TypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeName

`func (o *VolumeShowResponse) SetTypeName(v string)`

SetTypeName sets TypeName field to given value.


### GetUsage

`func (o *VolumeShowResponse) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *VolumeShowResponse) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *VolumeShowResponse) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *VolumeShowResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *VolumeShowResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *VolumeShowResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


