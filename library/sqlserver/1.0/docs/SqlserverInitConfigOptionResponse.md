# SqlserverInitConfigOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEnabled** | Pointer to **bool** | Audit usage state | [optional] [default to false]
**BackupOption** | Pointer to [**NullableSqlserverBackupOption**](SqlserverBackupOption.md) |  | [optional] 
**DatabaseCollation** | Pointer to [**DbCollation**](DbCollation.md) |  | [optional] [default to DBCOLLATION_SQL_LATIN1_GENERAL_CP1_CI_AS]
**DatabasePort** | Pointer to **int32** | Database service port | [optional] [default to 2866]
**DatabaseServiceName** | **string** | Database Service Name | 
**DatabaseUserName** | **string** | Database User Name | 
**Databases** | [**[]SqlserverDatabaseOption**](SqlserverDatabaseOption.md) | Databases | 

## Methods

### NewSqlserverInitConfigOptionResponse

`func NewSqlserverInitConfigOptionResponse(databaseServiceName string, databaseUserName string, databases []SqlserverDatabaseOption, ) *SqlserverInitConfigOptionResponse`

NewSqlserverInitConfigOptionResponse instantiates a new SqlserverInitConfigOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverInitConfigOptionResponseWithDefaults

`func NewSqlserverInitConfigOptionResponseWithDefaults() *SqlserverInitConfigOptionResponse`

NewSqlserverInitConfigOptionResponseWithDefaults instantiates a new SqlserverInitConfigOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEnabled

`func (o *SqlserverInitConfigOptionResponse) GetAuditEnabled() bool`

GetAuditEnabled returns the AuditEnabled field if non-nil, zero value otherwise.

### GetAuditEnabledOk

`func (o *SqlserverInitConfigOptionResponse) GetAuditEnabledOk() (*bool, bool)`

GetAuditEnabledOk returns a tuple with the AuditEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEnabled

`func (o *SqlserverInitConfigOptionResponse) SetAuditEnabled(v bool)`

SetAuditEnabled sets AuditEnabled field to given value.

### HasAuditEnabled

`func (o *SqlserverInitConfigOptionResponse) HasAuditEnabled() bool`

HasAuditEnabled returns a boolean if a field has been set.

### GetBackupOption

`func (o *SqlserverInitConfigOptionResponse) GetBackupOption() SqlserverBackupOption`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *SqlserverInitConfigOptionResponse) GetBackupOptionOk() (*SqlserverBackupOption, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *SqlserverInitConfigOptionResponse) SetBackupOption(v SqlserverBackupOption)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *SqlserverInitConfigOptionResponse) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *SqlserverInitConfigOptionResponse) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *SqlserverInitConfigOptionResponse) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDatabaseCollation

`func (o *SqlserverInitConfigOptionResponse) GetDatabaseCollation() DbCollation`

GetDatabaseCollation returns the DatabaseCollation field if non-nil, zero value otherwise.

### GetDatabaseCollationOk

`func (o *SqlserverInitConfigOptionResponse) GetDatabaseCollationOk() (*DbCollation, bool)`

GetDatabaseCollationOk returns a tuple with the DatabaseCollation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseCollation

`func (o *SqlserverInitConfigOptionResponse) SetDatabaseCollation(v DbCollation)`

SetDatabaseCollation sets DatabaseCollation field to given value.

### HasDatabaseCollation

`func (o *SqlserverInitConfigOptionResponse) HasDatabaseCollation() bool`

HasDatabaseCollation returns a boolean if a field has been set.

### GetDatabasePort

`func (o *SqlserverInitConfigOptionResponse) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *SqlserverInitConfigOptionResponse) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *SqlserverInitConfigOptionResponse) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *SqlserverInitConfigOptionResponse) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### GetDatabaseServiceName

`func (o *SqlserverInitConfigOptionResponse) GetDatabaseServiceName() string`

GetDatabaseServiceName returns the DatabaseServiceName field if non-nil, zero value otherwise.

### GetDatabaseServiceNameOk

`func (o *SqlserverInitConfigOptionResponse) GetDatabaseServiceNameOk() (*string, bool)`

GetDatabaseServiceNameOk returns a tuple with the DatabaseServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseServiceName

`func (o *SqlserverInitConfigOptionResponse) SetDatabaseServiceName(v string)`

SetDatabaseServiceName sets DatabaseServiceName field to given value.


### GetDatabaseUserName

`func (o *SqlserverInitConfigOptionResponse) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *SqlserverInitConfigOptionResponse) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *SqlserverInitConfigOptionResponse) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetDatabases

`func (o *SqlserverInitConfigOptionResponse) GetDatabases() []SqlserverDatabaseOption`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *SqlserverInitConfigOptionResponse) GetDatabasesOk() (*[]SqlserverDatabaseOption, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *SqlserverInitConfigOptionResponse) SetDatabases(v []SqlserverDatabaseOption)`

SetDatabases sets Databases field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


