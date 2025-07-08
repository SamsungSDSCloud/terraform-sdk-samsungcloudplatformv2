# EpasInitConfigOptionDetail

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

## Methods

### NewEpasInitConfigOptionDetail

`func NewEpasInitConfigOptionDetail(databaseName string, databaseUserName string, ) *EpasInitConfigOptionDetail`

NewEpasInitConfigOptionDetail instantiates a new EpasInitConfigOptionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpasInitConfigOptionDetailWithDefaults

`func NewEpasInitConfigOptionDetailWithDefaults() *EpasInitConfigOptionDetail`

NewEpasInitConfigOptionDetailWithDefaults instantiates a new EpasInitConfigOptionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEnabled

`func (o *EpasInitConfigOptionDetail) GetAuditEnabled() bool`

GetAuditEnabled returns the AuditEnabled field if non-nil, zero value otherwise.

### GetAuditEnabledOk

`func (o *EpasInitConfigOptionDetail) GetAuditEnabledOk() (*bool, bool)`

GetAuditEnabledOk returns a tuple with the AuditEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEnabled

`func (o *EpasInitConfigOptionDetail) SetAuditEnabled(v bool)`

SetAuditEnabled sets AuditEnabled field to given value.

### HasAuditEnabled

`func (o *EpasInitConfigOptionDetail) HasAuditEnabled() bool`

HasAuditEnabled returns a boolean if a field has been set.

### GetBackupOption

`func (o *EpasInitConfigOptionDetail) GetBackupOption() EpasBackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *EpasInitConfigOptionDetail) GetBackupOptionOk() (*EpasBackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *EpasInitConfigOptionDetail) SetBackupOption(v EpasBackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *EpasInitConfigOptionDetail) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *EpasInitConfigOptionDetail) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *EpasInitConfigOptionDetail) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseEncoding

`func (o *EpasInitConfigOptionDetail) GetDatabaseEncoding() string`

GetDatabaseEncoding returns the DatabaseEncoding field if non-nil, zero value otherwise.

### GetDatabaseEncodingOk

`func (o *EpasInitConfigOptionDetail) GetDatabaseEncodingOk() (*string, bool)`

GetDatabaseEncodingOk returns a tuple with the DatabaseEncoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseEncoding

`func (o *EpasInitConfigOptionDetail) SetDatabaseEncoding(v string)`

SetDatabaseEncoding sets DatabaseEncoding field to given value.

### HasDatabaseEncoding

`func (o *EpasInitConfigOptionDetail) HasDatabaseEncoding() bool`

HasDatabaseEncoding returns a boolean if a field has been set.

### SetDatabaseEncodingNil

`func (o *EpasInitConfigOptionDetail) SetDatabaseEncodingNil(b bool)`

 SetDatabaseEncodingNil sets the value for DatabaseEncoding to be an explicit nil

### UnsetDatabaseEncoding
`func (o *EpasInitConfigOptionDetail) UnsetDatabaseEncoding()`

UnsetDatabaseEncoding ensures that no value is present for DatabaseEncoding, not even an explicit nil
### GetDatabaseLocale

`func (o *EpasInitConfigOptionDetail) GetDatabaseLocale() string`

GetDatabaseLocale returns the DatabaseLocale field if non-nil, zero value otherwise.

### GetDatabaseLocaleOk

`func (o *EpasInitConfigOptionDetail) GetDatabaseLocaleOk() (*string, bool)`

GetDatabaseLocaleOk returns a tuple with the DatabaseLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseLocale

`func (o *EpasInitConfigOptionDetail) SetDatabaseLocale(v string)`

SetDatabaseLocale sets DatabaseLocale field to given value.

### HasDatabaseLocale

`func (o *EpasInitConfigOptionDetail) HasDatabaseLocale() bool`

HasDatabaseLocale returns a boolean if a field has been set.

### SetDatabaseLocaleNil

`func (o *EpasInitConfigOptionDetail) SetDatabaseLocaleNil(b bool)`

 SetDatabaseLocaleNil sets the value for DatabaseLocale to be an explicit nil

### UnsetDatabaseLocale
`func (o *EpasInitConfigOptionDetail) UnsetDatabaseLocale()`

UnsetDatabaseLocale ensures that no value is present for DatabaseLocale, not even an explicit nil
### GetDatabaseName

`func (o *EpasInitConfigOptionDetail) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *EpasInitConfigOptionDetail) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *EpasInitConfigOptionDetail) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDatabasePort

`func (o *EpasInitConfigOptionDetail) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *EpasInitConfigOptionDetail) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *EpasInitConfigOptionDetail) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *EpasInitConfigOptionDetail) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *EpasInitConfigOptionDetail) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *EpasInitConfigOptionDetail) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *EpasInitConfigOptionDetail) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *EpasInitConfigOptionDetail) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *EpasInitConfigOptionDetail) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


