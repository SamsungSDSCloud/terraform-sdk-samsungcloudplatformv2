# EpasInitConfigOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEnabled** | Pointer to **bool** | Audit usage state | [optional] [default to false]
**BackupOption** | Pointer to [**NullableEpasBackupOption**](EpasBackupOption.md) |  | [optional] 
**DatabaseEncoding** | Pointer to **NullableString** |  | [optional] 
**DatabaseLocale** | Pointer to **NullableString** |  | [optional] 
**DatabaseName** | **string** | Database Name | 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserName** | **string** | Database User Name | 
**DatabaseUserPassword** | **string** | Database user password | 

## Methods

### NewEpasInitConfigOptionRequest

`func NewEpasInitConfigOptionRequest(databaseName string, databaseUserName string, databaseUserPassword string, ) *EpasInitConfigOptionRequest`

NewEpasInitConfigOptionRequest instantiates a new EpasInitConfigOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpasInitConfigOptionRequestWithDefaults

`func NewEpasInitConfigOptionRequestWithDefaults() *EpasInitConfigOptionRequest`

NewEpasInitConfigOptionRequestWithDefaults instantiates a new EpasInitConfigOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEnabled

`func (o *EpasInitConfigOptionRequest) GetAuditEnabled() bool`

GetAuditEnabled returns the AuditEnabled field if non-nil, zero value otherwise.

### GetAuditEnabledOk

`func (o *EpasInitConfigOptionRequest) GetAuditEnabledOk() (*bool, bool)`

GetAuditEnabledOk returns a tuple with the AuditEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEnabled

`func (o *EpasInitConfigOptionRequest) SetAuditEnabled(v bool)`

SetAuditEnabled sets AuditEnabled field to given value.

### HasAuditEnabled

`func (o *EpasInitConfigOptionRequest) HasAuditEnabled() bool`

HasAuditEnabled returns a boolean if a field has been set.

### GetBackupOption

`func (o *EpasInitConfigOptionRequest) GetBackupOption() EpasBackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *EpasInitConfigOptionRequest) GetBackupOptionOk() (*EpasBackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *EpasInitConfigOptionRequest) SetBackupOption(v EpasBackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *EpasInitConfigOptionRequest) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *EpasInitConfigOptionRequest) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *EpasInitConfigOptionRequest) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseEncoding

`func (o *EpasInitConfigOptionRequest) GetDatabaseEncoding() string`

GetDatabaseEncoding returns the DatabaseEncoding field if non-nil, zero value otherwise.

### GetDatabaseEncodingOk

`func (o *EpasInitConfigOptionRequest) GetDatabaseEncodingOk() (*string, bool)`

GetDatabaseEncodingOk returns a tuple with the DatabaseEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEncoding

`func (o *EpasInitConfigOptionRequest) SetDatabaseEncoding(v string)`

SetDatabaseEncoding sets DatabaseEncoding field to given value.

### HasDatabaseEncoding

`func (o *EpasInitConfigOptionRequest) HasDatabaseEncoding() bool`

HasDatabaseEncoding returns a boolean if a field has been set.

### SetDatabaseEncodingNil

`func (o *EpasInitConfigOptionRequest) SetDatabaseEncodingNil(b bool)`

 SetDatabaseEncodingNil sets the value for DatabaseEncoding to be an explicit nil

### UnsetDatabaseEncoding
`func (o *EpasInitConfigOptionRequest) UnsetDatabaseEncoding()`

UnsetDatabaseEncoding ensures that no value is present for DatabaseEncoding, not even an explicit nil
### GetDatabaseLocale

`func (o *EpasInitConfigOptionRequest) GetDatabaseLocale() string`

GetDatabaseLocale returns the DatabaseLocale field if non-nil, zero value otherwise.

### GetDatabaseLocaleOk

`func (o *EpasInitConfigOptionRequest) GetDatabaseLocaleOk() (*string, bool)`

GetDatabaseLocaleOk returns a tuple with the DatabaseLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseLocale

`func (o *EpasInitConfigOptionRequest) SetDatabaseLocale(v string)`

SetDatabaseLocale sets DatabaseLocale field to given value.

### HasDatabaseLocale

`func (o *EpasInitConfigOptionRequest) HasDatabaseLocale() bool`

HasDatabaseLocale returns a boolean if a field has been set.

### SetDatabaseLocaleNil

`func (o *EpasInitConfigOptionRequest) SetDatabaseLocaleNil(b bool)`

 SetDatabaseLocaleNil sets the value for DatabaseLocale to be an explicit nil

### UnsetDatabaseLocale
`func (o *EpasInitConfigOptionRequest) UnsetDatabaseLocale()`

UnsetDatabaseLocale ensures that no value is present for DatabaseLocale, not even an explicit nil
### GetDatabaseName

`func (o *EpasInitConfigOptionRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *EpasInitConfigOptionRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *EpasInitConfigOptionRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabasePort

`func (o *EpasInitConfigOptionRequest) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *EpasInitConfigOptionRequest) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *EpasInitConfigOptionRequest) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *EpasInitConfigOptionRequest) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *EpasInitConfigOptionRequest) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *EpasInitConfigOptionRequest) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *EpasInitConfigOptionRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *EpasInitConfigOptionRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *EpasInitConfigOptionRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *EpasInitConfigOptionRequest) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *EpasInitConfigOptionRequest) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *EpasInitConfigOptionRequest) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


