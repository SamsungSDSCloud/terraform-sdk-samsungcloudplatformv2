# MysqlInitConfigOptionResponseDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupOption** | Pointer to [**NullableMysqlBackupOption**](MysqlBackupOption.md) |  | [optional] 
**DatabaseCharacterSet** | Pointer to **NullableString** |  | [optional] 
**DatabaseName** | **string** | Database Name | 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserName** | **string** | Database User Name | 

## Methods

### NewMysqlInitConfigOptionResponseDetail

`func NewMysqlInitConfigOptionResponseDetail(databaseName string, databaseUserName string, ) *MysqlInitConfigOptionResponseDetail`

NewMysqlInitConfigOptionResponseDetail instantiates a new MysqlInitConfigOptionResponseDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlInitConfigOptionResponseDetailWithDefaults

`func NewMysqlInitConfigOptionResponseDetailWithDefaults() *MysqlInitConfigOptionResponseDetail`

NewMysqlInitConfigOptionResponseDetailWithDefaults instantiates a new MysqlInitConfigOptionResponseDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupOption

`func (o *MysqlInitConfigOptionResponseDetail) GetBackupOption() MysqlBackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *MysqlInitConfigOptionResponseDetail) GetBackupOptionOk() (*MysqlBackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *MysqlInitConfigOptionResponseDetail) SetBackupOption(v MysqlBackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *MysqlInitConfigOptionResponseDetail) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *MysqlInitConfigOptionResponseDetail) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *MysqlInitConfigOptionResponseDetail) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseCharacterSet

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabaseCharacterSet() string`

GetDatabaseCharacterSet returns the DatabaseCharacterSet field if non-nil, zero value otherwise.

### GetDatabaseCharacterSetOk

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabaseCharacterSetOk() (*string, bool)`

GetDatabaseCharacterSetOk returns a tuple with the DatabaseCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCharacterSet

`func (o *MysqlInitConfigOptionResponseDetail) SetDatabaseCharacterSet(v string)`

SetDatabaseCharacterSet sets DatabaseCharacterSet field to given value.

### HasDatabaseCharacterSet

`func (o *MysqlInitConfigOptionResponseDetail) HasDatabaseCharacterSet() bool`

HasDatabaseCharacterSet returns a boolean if a field has been set.

### SetDatabaseCharacterSetNil

`func (o *MysqlInitConfigOptionResponseDetail) SetDatabaseCharacterSetNil(b bool)`

 SetDatabaseCharacterSetNil sets the value for DatabaseCharacterSet to be an explicit nil

### UnsetDatabaseCharacterSet
`func (o *MysqlInitConfigOptionResponseDetail) UnsetDatabaseCharacterSet()`

UnsetDatabaseCharacterSet ensures that no value is present for DatabaseCharacterSet, not even an explicit nil
### GetDatabaseName

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *MysqlInitConfigOptionResponseDetail) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabasePort

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *MysqlInitConfigOptionResponseDetail) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *MysqlInitConfigOptionResponseDetail) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *MysqlInitConfigOptionResponseDetail) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *MysqlInitConfigOptionResponseDetail) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *MysqlInitConfigOptionResponseDetail) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *MysqlInitConfigOptionResponseDetail) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


