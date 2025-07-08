# SearchEngineInitConfigOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupOption** | Pointer to [**NullableBackupSettingExcludingArchiveRequest**](BackupSettingExcludingArchiveRequest.md) |  | [optional] 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserName** | **string** | Database User Name | 
**DatabaseUserPassword** | **string** | Database user password | 

## Methods

### NewSearchEngineInitConfigOptionRequest

`func NewSearchEngineInitConfigOptionRequest(databaseUserName string, databaseUserPassword string, ) *SearchEngineInitConfigOptionRequest`

NewSearchEngineInitConfigOptionRequest instantiates a new SearchEngineInitConfigOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEngineInitConfigOptionRequestWithDefaults

`func NewSearchEngineInitConfigOptionRequestWithDefaults() *SearchEngineInitConfigOptionRequest`

NewSearchEngineInitConfigOptionRequestWithDefaults instantiates a new SearchEngineInitConfigOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupOption

`func (o *SearchEngineInitConfigOptionRequest) GetBackupOption() BackupSettingExcludingArchiveRequest`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *SearchEngineInitConfigOptionRequest) GetBackupOptionOk() (*BackupSettingExcludingArchiveRequest, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *SearchEngineInitConfigOptionRequest) SetBackupOption(v BackupSettingExcludingArchiveRequest)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *SearchEngineInitConfigOptionRequest) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *SearchEngineInitConfigOptionRequest) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *SearchEngineInitConfigOptionRequest) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabasePort

`func (o *SearchEngineInitConfigOptionRequest) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *SearchEngineInitConfigOptionRequest) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *SearchEngineInitConfigOptionRequest) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *SearchEngineInitConfigOptionRequest) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *SearchEngineInitConfigOptionRequest) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *SearchEngineInitConfigOptionRequest) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *SearchEngineInitConfigOptionRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *SearchEngineInitConfigOptionRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *SearchEngineInitConfigOptionRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *SearchEngineInitConfigOptionRequest) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *SearchEngineInitConfigOptionRequest) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *SearchEngineInitConfigOptionRequest) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


