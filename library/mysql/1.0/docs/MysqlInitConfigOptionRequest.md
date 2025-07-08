# MysqlInitConfigOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEnabled** | Pointer to **interface{}** |  | [optional] 
**BackupOption** | Pointer to [**NullableMysqlBackupOption**](MysqlBackupOption.md) |  | [optional] 
**DatabaseCaseSensitive** | Pointer to **bool** | Database case sensitivity | [optional] [default to false]
**DatabaseCharacterSet** | Pointer to **NullableString** |  | [optional] 
**DatabaseName** | **string** | Database Name | 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserName** | **string** | Database User Name | 
**DatabaseUserPassword** | **string** | Database user password | 

## Methods

### NewMysqlInitConfigOptionRequest

`func NewMysqlInitConfigOptionRequest(databaseName string, databaseUserName string, databaseUserPassword string, ) *MysqlInitConfigOptionRequest`

NewMysqlInitConfigOptionRequest instantiates a new MysqlInitConfigOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlInitConfigOptionRequestWithDefaults

`func NewMysqlInitConfigOptionRequestWithDefaults() *MysqlInitConfigOptionRequest`

NewMysqlInitConfigOptionRequestWithDefaults instantiates a new MysqlInitConfigOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEnabled

`func (o *MysqlInitConfigOptionRequest) GetAuditEnabled() interface{}`

GetAuditEnabled returns the AuditEnabled field if non-nil, zero value otherwise.

### GetAuditEnabledOk

`func (o *MysqlInitConfigOptionRequest) GetAuditEnabledOk() (*interface{}, bool)`

GetAuditEnabledOk returns a tuple with the AuditEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEnabled

`func (o *MysqlInitConfigOptionRequest) SetAuditEnabled(v interface{})`

SetAuditEnabled sets AuditEnabled field to given value.

### HasAuditEnabled

`func (o *MysqlInitConfigOptionRequest) HasAuditEnabled() bool`

HasAuditEnabled returns a boolean if a field has been set.

### SetAuditEnabledNil

`func (o *MysqlInitConfigOptionRequest) SetAuditEnabledNil(b bool)`

 SetAuditEnabledNil sets the value for AuditEnabled to be an explicit nil

### UnsetAuditEnabled
`func (o *MysqlInitConfigOptionRequest) UnsetAuditEnabled()`

UnsetAuditEnabled ensures that no value is present for AuditEnabled, not even an explicit nil
### GetBackupOption

`func (o *MysqlInitConfigOptionRequest) GetBackupOption() MysqlBackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *MysqlInitConfigOptionRequest) GetBackupOptionOk() (*MysqlBackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *MysqlInitConfigOptionRequest) SetBackupOption(v MysqlBackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *MysqlInitConfigOptionRequest) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *MysqlInitConfigOptionRequest) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *MysqlInitConfigOptionRequest) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseCaseSensitive

`func (o *MysqlInitConfigOptionRequest) GetDatabaseCaseSensitive() bool`

GetDatabaseCaseSensitive returns the DatabaseCaseSensitive field if non-nil, zero value otherwise.

### GetDatabaseCaseSensitiveOk

`func (o *MysqlInitConfigOptionRequest) GetDatabaseCaseSensitiveOk() (*bool, bool)`

GetDatabaseCaseSensitiveOk returns a tuple with the DatabaseCaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCaseSensitive

`func (o *MysqlInitConfigOptionRequest) SetDatabaseCaseSensitive(v bool)`

SetDatabaseCaseSensitive sets DatabaseCaseSensitive field to given value.

### HasDatabaseCaseSensitive

`func (o *MysqlInitConfigOptionRequest) HasDatabaseCaseSensitive() bool`

HasDatabaseCaseSensitive returns a boolean if a field has been set.

### GetDatabaseCharacterSet

`func (o *MysqlInitConfigOptionRequest) GetDatabaseCharacterSet() string`

GetDatabaseCharacterSet returns the DatabaseCharacterSet field if non-nil, zero value otherwise.

### GetDatabaseCharacterSetOk

`func (o *MysqlInitConfigOptionRequest) GetDatabaseCharacterSetOk() (*string, bool)`

GetDatabaseCharacterSetOk returns a tuple with the DatabaseCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCharacterSet

`func (o *MysqlInitConfigOptionRequest) SetDatabaseCharacterSet(v string)`

SetDatabaseCharacterSet sets DatabaseCharacterSet field to given value.

### HasDatabaseCharacterSet

`func (o *MysqlInitConfigOptionRequest) HasDatabaseCharacterSet() bool`

HasDatabaseCharacterSet returns a boolean if a field has been set.

### SetDatabaseCharacterSetNil

`func (o *MysqlInitConfigOptionRequest) SetDatabaseCharacterSetNil(b bool)`

 SetDatabaseCharacterSetNil sets the value for DatabaseCharacterSet to be an explicit nil

### UnsetDatabaseCharacterSet
`func (o *MysqlInitConfigOptionRequest) UnsetDatabaseCharacterSet()`

UnsetDatabaseCharacterSet ensures that no value is present for DatabaseCharacterSet, not even an explicit nil
### GetDatabaseName

`func (o *MysqlInitConfigOptionRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *MysqlInitConfigOptionRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *MysqlInitConfigOptionRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabasePort

`func (o *MysqlInitConfigOptionRequest) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *MysqlInitConfigOptionRequest) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *MysqlInitConfigOptionRequest) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *MysqlInitConfigOptionRequest) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *MysqlInitConfigOptionRequest) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *MysqlInitConfigOptionRequest) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *MysqlInitConfigOptionRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *MysqlInitConfigOptionRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *MysqlInitConfigOptionRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *MysqlInitConfigOptionRequest) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *MysqlInitConfigOptionRequest) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *MysqlInitConfigOptionRequest) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


