# VerticaInitConfigOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupOption** | Pointer to [**NullableBackupSettingExcludingArchiveRequest**](BackupSettingExcludingArchiveRequest.md) |  | [optional] 
**DatabaseLocale** | Pointer to **NullableString** |  | [optional] 
**DatabaseName** | **string** | Database Name | 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserName** | **string** | Database User Name | 
**McPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewVerticaInitConfigOptionResponse

`func NewVerticaInitConfigOptionResponse(databaseName string, databaseUserName string, ) *VerticaInitConfigOptionResponse`

NewVerticaInitConfigOptionResponse instantiates a new VerticaInitConfigOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticaInitConfigOptionResponseWithDefaults

`func NewVerticaInitConfigOptionResponseWithDefaults() *VerticaInitConfigOptionResponse`

NewVerticaInitConfigOptionResponseWithDefaults instantiates a new VerticaInitConfigOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupOption

`func (o *VerticaInitConfigOptionResponse) GetBackupOption() BackupSettingExcludingArchiveRequest`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *VerticaInitConfigOptionResponse) GetBackupOptionOk() (*BackupSettingExcludingArchiveRequest, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *VerticaInitConfigOptionResponse) SetBackupOption(v BackupSettingExcludingArchiveRequest)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *VerticaInitConfigOptionResponse) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *VerticaInitConfigOptionResponse) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *VerticaInitConfigOptionResponse) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseLocale

`func (o *VerticaInitConfigOptionResponse) GetDatabaseLocale() string`

GetDatabaseLocale returns the DatabaseLocale field if non-nil, zero value otherwise.

### GetDatabaseLocaleOk

`func (o *VerticaInitConfigOptionResponse) GetDatabaseLocaleOk() (*string, bool)`

GetDatabaseLocaleOk returns a tuple with the DatabaseLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseLocale

`func (o *VerticaInitConfigOptionResponse) SetDatabaseLocale(v string)`

SetDatabaseLocale sets DatabaseLocale field to given value.

### HasDatabaseLocale

`func (o *VerticaInitConfigOptionResponse) HasDatabaseLocale() bool`

HasDatabaseLocale returns a boolean if a field has been set.

### SetDatabaseLocaleNil

`func (o *VerticaInitConfigOptionResponse) SetDatabaseLocaleNil(b bool)`

 SetDatabaseLocaleNil sets the value for DatabaseLocale to be an explicit nil

### UnsetDatabaseLocale
`func (o *VerticaInitConfigOptionResponse) UnsetDatabaseLocale()`

UnsetDatabaseLocale ensures that no value is present for DatabaseLocale, not even an explicit nil
### GetDatabaseName

`func (o *VerticaInitConfigOptionResponse) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *VerticaInitConfigOptionResponse) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *VerticaInitConfigOptionResponse) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabasePort

`func (o *VerticaInitConfigOptionResponse) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *VerticaInitConfigOptionResponse) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *VerticaInitConfigOptionResponse) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *VerticaInitConfigOptionResponse) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *VerticaInitConfigOptionResponse) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *VerticaInitConfigOptionResponse) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *VerticaInitConfigOptionResponse) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *VerticaInitConfigOptionResponse) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *VerticaInitConfigOptionResponse) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetMcPort

`func (o *VerticaInitConfigOptionResponse) GetMcPort() int32`

GetMcPort returns the McPort field if non-nil, zero value otherwise.

### GetMcPortOk

`func (o *VerticaInitConfigOptionResponse) GetMcPortOk() (*int32, bool)`

GetMcPortOk returns a tuple with the McPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcPort

`func (o *VerticaInitConfigOptionResponse) SetMcPort(v int32)`

SetMcPort sets McPort field to given value.

### HasMcPort

`func (o *VerticaInitConfigOptionResponse) HasMcPort() bool`

HasMcPort returns a boolean if a field has been set.

### SetMcPortNil

`func (o *VerticaInitConfigOptionResponse) SetMcPortNil(b bool)`

 SetMcPortNil sets the value for McPort to be an explicit nil

### UnsetMcPort
`func (o *VerticaInitConfigOptionResponse) UnsetMcPort()`

UnsetMcPort ensures that no value is present for McPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


