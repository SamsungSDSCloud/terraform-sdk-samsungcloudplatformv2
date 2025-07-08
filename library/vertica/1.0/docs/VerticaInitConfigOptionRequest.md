# VerticaInitConfigOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupOption** | Pointer to [**NullableBackupSettingExcludingArchiveRequest**](BackupSettingExcludingArchiveRequest.md) |  | [optional] 
**DatabaseLocale** | Pointer to **NullableString** |  | [optional] 
**DatabaseName** | **string** | Database Name | 
**DatabaseUserName** | **string** | Database User Name | 
**DatabaseUserPassword** | **string** | Database user password | 

## Methods

### NewVerticaInitConfigOptionRequest

`func NewVerticaInitConfigOptionRequest(databaseName string, databaseUserName string, databaseUserPassword string, ) *VerticaInitConfigOptionRequest`

NewVerticaInitConfigOptionRequest instantiates a new VerticaInitConfigOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticaInitConfigOptionRequestWithDefaults

`func NewVerticaInitConfigOptionRequestWithDefaults() *VerticaInitConfigOptionRequest`

NewVerticaInitConfigOptionRequestWithDefaults instantiates a new VerticaInitConfigOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupOption

`func (o *VerticaInitConfigOptionRequest) GetBackupOption() BackupSettingExcludingArchiveRequest`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *VerticaInitConfigOptionRequest) GetBackupOptionOk() (*BackupSettingExcludingArchiveRequest, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *VerticaInitConfigOptionRequest) SetBackupOption(v BackupSettingExcludingArchiveRequest)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *VerticaInitConfigOptionRequest) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *VerticaInitConfigOptionRequest) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *VerticaInitConfigOptionRequest) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseLocale

`func (o *VerticaInitConfigOptionRequest) GetDatabaseLocale() string`

GetDatabaseLocale returns the DatabaseLocale field if non-nil, zero value otherwise.

### GetDatabaseLocaleOk

`func (o *VerticaInitConfigOptionRequest) GetDatabaseLocaleOk() (*string, bool)`

GetDatabaseLocaleOk returns a tuple with the DatabaseLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseLocale

`func (o *VerticaInitConfigOptionRequest) SetDatabaseLocale(v string)`

SetDatabaseLocale sets DatabaseLocale field to given value.

### HasDatabaseLocale

`func (o *VerticaInitConfigOptionRequest) HasDatabaseLocale() bool`

HasDatabaseLocale returns a boolean if a field has been set.

### SetDatabaseLocaleNil

`func (o *VerticaInitConfigOptionRequest) SetDatabaseLocaleNil(b bool)`

 SetDatabaseLocaleNil sets the value for DatabaseLocale to be an explicit nil

### UnsetDatabaseLocale
`func (o *VerticaInitConfigOptionRequest) UnsetDatabaseLocale()`

UnsetDatabaseLocale ensures that no value is present for DatabaseLocale, not even an explicit nil
### GetDatabaseName

`func (o *VerticaInitConfigOptionRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *VerticaInitConfigOptionRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *VerticaInitConfigOptionRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabaseUserName

`func (o *VerticaInitConfigOptionRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *VerticaInitConfigOptionRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *VerticaInitConfigOptionRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *VerticaInitConfigOptionRequest) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *VerticaInitConfigOptionRequest) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *VerticaInitConfigOptionRequest) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


