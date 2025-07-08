# RedisInitConfigOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupOption** | Pointer to [**NullableBackupOption**](BackupOption.md) |  | [optional] 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserPassword** | **NullableString** |  | 
**SentinelPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRedisInitConfigOption

`func NewRedisInitConfigOption(databaseUserPassword NullableString, ) *RedisInitConfigOption`

NewRedisInitConfigOption instantiates a new RedisInitConfigOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedisInitConfigOptionWithDefaults

`func NewRedisInitConfigOptionWithDefaults() *RedisInitConfigOption`

NewRedisInitConfigOptionWithDefaults instantiates a new RedisInitConfigOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupOption

`func (o *RedisInitConfigOption) GetBackupOption() BackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *RedisInitConfigOption) GetBackupOptionOk() (*BackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *RedisInitConfigOption) SetBackupOption(v BackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *RedisInitConfigOption) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *RedisInitConfigOption) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *RedisInitConfigOption) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabasePort

`func (o *RedisInitConfigOption) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *RedisInitConfigOption) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *RedisInitConfigOption) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *RedisInitConfigOption) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *RedisInitConfigOption) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *RedisInitConfigOption) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserPassword

`func (o *RedisInitConfigOption) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *RedisInitConfigOption) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *RedisInitConfigOption) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.


### SetDatabaseUserPasswordNil

`func (o *RedisInitConfigOption) SetDatabaseUserPasswordNil(b bool)`

 SetDatabaseUserPasswordNil sets the value for DatabaseUserPassword to be an explicit nil

### UnsetDatabaseUserPassword
`func (o *RedisInitConfigOption) UnsetDatabaseUserPassword()`

UnsetDatabaseUserPassword ensures that no value is present for DatabaseUserPassword, not even an explicit nil
### GetSentinelPort

`func (o *RedisInitConfigOption) GetSentinelPort() int32`

GetSentinelPort returns the SentinelPort field if non-nil, zero value otherwise.

### GetSentinelPortOk

`func (o *RedisInitConfigOption) GetSentinelPortOk() (*int32, bool)`

GetSentinelPortOk returns a tuple with the SentinelPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentinelPort

`func (o *RedisInitConfigOption) SetSentinelPort(v int32)`

SetSentinelPort sets SentinelPort field to given value.

### HasSentinelPort

`func (o *RedisInitConfigOption) HasSentinelPort() bool`

HasSentinelPort returns a boolean if a field has been set.

### SetSentinelPortNil

`func (o *RedisInitConfigOption) SetSentinelPortNil(b bool)`

 SetSentinelPortNil sets the value for SentinelPort to be an explicit nil

### UnsetSentinelPort
`func (o *RedisInitConfigOption) UnsetSentinelPort()`

UnsetSentinelPort ensures that no value is present for SentinelPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


