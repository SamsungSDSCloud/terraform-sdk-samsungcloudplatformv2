# RedisInitConfigOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupOption** | Pointer to [**NullableBackupOption**](BackupOption.md) |  | [optional] 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**SentinelPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewRedisInitConfigOptionResponse

`func NewRedisInitConfigOptionResponse() *RedisInitConfigOptionResponse`

NewRedisInitConfigOptionResponse instantiates a new RedisInitConfigOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedisInitConfigOptionResponseWithDefaults

`func NewRedisInitConfigOptionResponseWithDefaults() *RedisInitConfigOptionResponse`

NewRedisInitConfigOptionResponseWithDefaults instantiates a new RedisInitConfigOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupOption

`func (o *RedisInitConfigOptionResponse) GetBackupOption() BackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *RedisInitConfigOptionResponse) GetBackupOptionOk() (*BackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *RedisInitConfigOptionResponse) SetBackupOption(v BackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *RedisInitConfigOptionResponse) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *RedisInitConfigOptionResponse) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *RedisInitConfigOptionResponse) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabasePort

`func (o *RedisInitConfigOptionResponse) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *RedisInitConfigOptionResponse) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *RedisInitConfigOptionResponse) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *RedisInitConfigOptionResponse) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *RedisInitConfigOptionResponse) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *RedisInitConfigOptionResponse) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetSentinelPort

`func (o *RedisInitConfigOptionResponse) GetSentinelPort() int32`

GetSentinelPort returns the SentinelPort field if non-nil, zero value otherwise.

### GetSentinelPortOk

`func (o *RedisInitConfigOptionResponse) GetSentinelPortOk() (*int32, bool)`

GetSentinelPortOk returns a tuple with the SentinelPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentinelPort

`func (o *RedisInitConfigOptionResponse) SetSentinelPort(v int32)`

SetSentinelPort sets SentinelPort field to given value.

### HasSentinelPort

`func (o *RedisInitConfigOptionResponse) HasSentinelPort() bool`

HasSentinelPort returns a boolean if a field has been set.

### SetSentinelPortNil

`func (o *RedisInitConfigOptionResponse) SetSentinelPortNil(b bool)`

 SetSentinelPortNil sets the value for SentinelPort to be an explicit nil

### UnsetSentinelPort
`func (o *RedisInitConfigOptionResponse) UnsetSentinelPort()`

UnsetSentinelPort ensures that no value is present for SentinelPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


