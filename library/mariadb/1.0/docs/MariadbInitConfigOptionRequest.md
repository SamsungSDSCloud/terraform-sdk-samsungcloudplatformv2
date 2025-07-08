# MariadbInitConfigOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEnabled** | Pointer to **bool** | Audit usage state | [optional] [default to false]
**BackupOption** | Pointer to [**NullableMariadbBackupOption**](MariadbBackupOption.md) |  | [optional] 
**DatabaseCharacterSet** | Pointer to **NullableString** |  | [optional] 
**DatabaseName** | **string** | Database Name | 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserName** | **string** | Database User Name | 
**DatabaseUserPassword** | **string** | Database user password | 

## Methods

### NewMariadbInitConfigOptionRequest

`func NewMariadbInitConfigOptionRequest(databaseName string, databaseUserName string, databaseUserPassword string, ) *MariadbInitConfigOptionRequest`

NewMariadbInitConfigOptionRequest instantiates a new MariadbInitConfigOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbInitConfigOptionRequestWithDefaults

`func NewMariadbInitConfigOptionRequestWithDefaults() *MariadbInitConfigOptionRequest`

NewMariadbInitConfigOptionRequestWithDefaults instantiates a new MariadbInitConfigOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEnabled

`func (o *MariadbInitConfigOptionRequest) GetAuditEnabled() bool`

GetAuditEnabled returns the AuditEnabled field if non-nil, zero value otherwise.

### GetAuditEnabledOk

`func (o *MariadbInitConfigOptionRequest) GetAuditEnabledOk() (*bool, bool)`

GetAuditEnabledOk returns a tuple with the AuditEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEnabled

`func (o *MariadbInitConfigOptionRequest) SetAuditEnabled(v bool)`

SetAuditEnabled sets AuditEnabled field to given value.

### HasAuditEnabled

`func (o *MariadbInitConfigOptionRequest) HasAuditEnabled() bool`

HasAuditEnabled returns a boolean if a field has been set.

### GetBackupOption

`func (o *MariadbInitConfigOptionRequest) GetBackupOption() MariadbBackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *MariadbInitConfigOptionRequest) GetBackupOptionOk() (*MariadbBackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *MariadbInitConfigOptionRequest) SetBackupOption(v MariadbBackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *MariadbInitConfigOptionRequest) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *MariadbInitConfigOptionRequest) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *MariadbInitConfigOptionRequest) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseCharacterSet

`func (o *MariadbInitConfigOptionRequest) GetDatabaseCharacterSet() string`

GetDatabaseCharacterSet returns the DatabaseCharacterSet field if non-nil, zero value otherwise.

### GetDatabaseCharacterSetOk

`func (o *MariadbInitConfigOptionRequest) GetDatabaseCharacterSetOk() (*string, bool)`

GetDatabaseCharacterSetOk returns a tuple with the DatabaseCharacterSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCharacterSet

`func (o *MariadbInitConfigOptionRequest) SetDatabaseCharacterSet(v string)`

SetDatabaseCharacterSet sets DatabaseCharacterSet field to given value.

### HasDatabaseCharacterSet

`func (o *MariadbInitConfigOptionRequest) HasDatabaseCharacterSet() bool`

HasDatabaseCharacterSet returns a boolean if a field has been set.

### SetDatabaseCharacterSetNil

`func (o *MariadbInitConfigOptionRequest) SetDatabaseCharacterSetNil(b bool)`

 SetDatabaseCharacterSetNil sets the value for DatabaseCharacterSet to be an explicit nil

### UnsetDatabaseCharacterSet
`func (o *MariadbInitConfigOptionRequest) UnsetDatabaseCharacterSet()`

UnsetDatabaseCharacterSet ensures that no value is present for DatabaseCharacterSet, not even an explicit nil
### GetDatabaseName

`func (o *MariadbInitConfigOptionRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *MariadbInitConfigOptionRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *MariadbInitConfigOptionRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabasePort

`func (o *MariadbInitConfigOptionRequest) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *MariadbInitConfigOptionRequest) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *MariadbInitConfigOptionRequest) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *MariadbInitConfigOptionRequest) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *MariadbInitConfigOptionRequest) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *MariadbInitConfigOptionRequest) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *MariadbInitConfigOptionRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *MariadbInitConfigOptionRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *MariadbInitConfigOptionRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *MariadbInitConfigOptionRequest) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *MariadbInitConfigOptionRequest) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *MariadbInitConfigOptionRequest) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


