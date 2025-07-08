# SqlserverInitConfigOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEnabled** | Pointer to **bool** | Audit usage state | [optional] [default to false]
**BackupOption** | Pointer to [**NullableSqlserverBackupOption**](SqlserverBackupOption.md) |  | [optional] 
**DatabaseCollation** | Pointer to [**DbCollation**](DbCollation.md) |  | [optional] [default to DBCOLLATION_SQL_LATIN1_GENERAL_CP1_CI_AS]
**DatabaseName** | Pointer to **interface{}** |  | [optional] 
**DatabasePort** | Pointer to **int32** | Database service port | [optional] [default to 2866]
**DatabaseServiceName** | **string** | Database Service Name | 
**DatabaseUserName** | **string** | Database User Name | 
**DatabaseUserPassword** | **string** | Database user password | 
**Databases** | [**[]SqlserverDatabaseOption**](SqlserverDatabaseOption.md) | Databases | 
**License** | **string** | license | 

## Methods

### NewSqlserverInitConfigOptionRequest

`func NewSqlserverInitConfigOptionRequest(databaseServiceName string, databaseUserName string, databaseUserPassword string, databases []SqlserverDatabaseOption, license string, ) *SqlserverInitConfigOptionRequest`

NewSqlserverInitConfigOptionRequest instantiates a new SqlserverInitConfigOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverInitConfigOptionRequestWithDefaults

`func NewSqlserverInitConfigOptionRequestWithDefaults() *SqlserverInitConfigOptionRequest`

NewSqlserverInitConfigOptionRequestWithDefaults instantiates a new SqlserverInitConfigOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEnabled

`func (o *SqlserverInitConfigOptionRequest) GetAuditEnabled() bool`

GetAuditEnabled returns the AuditEnabled field if non-nil, zero value otherwise.

### GetAuditEnabledOk

`func (o *SqlserverInitConfigOptionRequest) GetAuditEnabledOk() (*bool, bool)`

GetAuditEnabledOk returns a tuple with the AuditEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEnabled

`func (o *SqlserverInitConfigOptionRequest) SetAuditEnabled(v bool)`

SetAuditEnabled sets AuditEnabled field to given value.

### HasAuditEnabled

`func (o *SqlserverInitConfigOptionRequest) HasAuditEnabled() bool`

HasAuditEnabled returns a boolean if a field has been set.

### GetBackupOption

`func (o *SqlserverInitConfigOptionRequest) GetBackupOption() SqlserverBackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *SqlserverInitConfigOptionRequest) GetBackupOptionOk() (*SqlserverBackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *SqlserverInitConfigOptionRequest) SetBackupOption(v SqlserverBackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *SqlserverInitConfigOptionRequest) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *SqlserverInitConfigOptionRequest) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *SqlserverInitConfigOptionRequest) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseCollation

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseCollation() DbCollation`

GetDatabaseCollation returns the DatabaseCollation field if non-nil, zero value otherwise.

### GetDatabaseCollationOk

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseCollationOk() (*DbCollation, bool)`

GetDatabaseCollationOk returns a tuple with the DatabaseCollation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCollation

`func (o *SqlserverInitConfigOptionRequest) SetDatabaseCollation(v DbCollation)`

SetDatabaseCollation sets DatabaseCollation field to given value.

### HasDatabaseCollation

`func (o *SqlserverInitConfigOptionRequest) HasDatabaseCollation() bool`

HasDatabaseCollation returns a boolean if a field has been set.

### GetDatabaseName

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseName() interface{}`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseNameOk() (*interface{}, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *SqlserverInitConfigOptionRequest) SetDatabaseName(v interface{})`

SetDatabaseName sets DatabaseName field to given value.

### HasDatabaseName

`func (o *SqlserverInitConfigOptionRequest) HasDatabaseName() bool`

HasDatabaseName returns a boolean if a field has been set.

### SetDatabaseNameNil

`func (o *SqlserverInitConfigOptionRequest) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *SqlserverInitConfigOptionRequest) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetDatabasePort

`func (o *SqlserverInitConfigOptionRequest) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *SqlserverInitConfigOptionRequest) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *SqlserverInitConfigOptionRequest) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *SqlserverInitConfigOptionRequest) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### GetDatabaseServiceName

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseServiceName() string`

GetDatabaseServiceName returns the DatabaseServiceName field if non-nil, zero value otherwise.

### GetDatabaseServiceNameOk

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseServiceNameOk() (*string, bool)`

GetDatabaseServiceNameOk returns a tuple with the DatabaseServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseServiceName

`func (o *SqlserverInitConfigOptionRequest) SetDatabaseServiceName(v string)`

SetDatabaseServiceName sets DatabaseServiceName field to given value.


### GetDatabaseUserName

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *SqlserverInitConfigOptionRequest) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabaseUserPassword

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseUserPassword() string`

GetDatabaseUserPassword returns the DatabaseUserPassword field if non-nil, zero value otherwise.

### GetDatabaseUserPasswordOk

`func (o *SqlserverInitConfigOptionRequest) GetDatabaseUserPasswordOk() (*string, bool)`

GetDatabaseUserPasswordOk returns a tuple with the DatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserPassword

`func (o *SqlserverInitConfigOptionRequest) SetDatabaseUserPassword(v string)`

SetDatabaseUserPassword sets DatabaseUserPassword field to given value.


### GetDatabases

`func (o *SqlserverInitConfigOptionRequest) GetDatabases() []SqlserverDatabaseOption`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *SqlserverInitConfigOptionRequest) GetDatabasesOk() (*[]SqlserverDatabaseOption, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *SqlserverInitConfigOptionRequest) SetDatabases(v []SqlserverDatabaseOption)`

SetDatabases sets Databases field to given value.


### GetLicense

`func (o *SqlserverInitConfigOptionRequest) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SqlserverInitConfigOptionRequest) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SqlserverInitConfigOptionRequest) SetLicense(v string)`

SetLicense sets License field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


