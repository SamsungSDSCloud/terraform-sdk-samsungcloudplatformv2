# SearchEngineInitConfigOptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupOption** | Pointer to [**NullableBackupSettingExcludingArchiveRequest**](BackupSettingExcludingArchiveRequest.md) |  | [optional] 
**DashboardsPort** | Pointer to **NullableInt32** |  | [optional] 
**DatabasePort** | Pointer to **NullableInt32** |  | [optional] 
**DatabaseUserName** | **string** | Database User Name | 
**KibanaPort** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewSearchEngineInitConfigOptionResponse

`func NewSearchEngineInitConfigOptionResponse(databaseUserName string, ) *SearchEngineInitConfigOptionResponse`

NewSearchEngineInitConfigOptionResponse instantiates a new SearchEngineInitConfigOptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEngineInitConfigOptionResponseWithDefaults

`func NewSearchEngineInitConfigOptionResponseWithDefaults() *SearchEngineInitConfigOptionResponse`

NewSearchEngineInitConfigOptionResponseWithDefaults instantiates a new SearchEngineInitConfigOptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupOption

`func (o *SearchEngineInitConfigOptionResponse) GetBackupOption() BackupSettingExcludingArchiveRequest`

GetBackupOption returns the BackupOption field if non-nil, zero value otherwise.

### GetBackupOptionOk

`func (o *SearchEngineInitConfigOptionResponse) GetBackupOptionOk() (*BackupSettingExcludingArchiveRequest, bool)`

GetBackupOptionOk returns a tuple with the BackupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupOption

`func (o *SearchEngineInitConfigOptionResponse) SetBackupOption(v BackupSettingExcludingArchiveRequest)`

SetBackupOption sets BackupOption field to given value.

### HasBackupOption

`func (o *SearchEngineInitConfigOptionResponse) HasBackupOption() bool`

HasBackupOption returns a boolean if a field has been set.

### SetBackupOptionNil

`func (o *SearchEngineInitConfigOptionResponse) SetBackupOptionNil(b bool)`

 SetBackupOptionNil sets the value for BackupOption to be an explicit nil

### UnsetBackupOption
`func (o *SearchEngineInitConfigOptionResponse) UnsetBackupOption()`

UnsetBackupOption ensures that no value is present for BackupOption, not even an explicit nil
### GetDashboardsPort

`func (o *SearchEngineInitConfigOptionResponse) GetDashboardsPort() int32`

GetDashboardsPort returns the DashboardsPort field if non-nil, zero value otherwise.

### GetDashboardsPortOk

`func (o *SearchEngineInitConfigOptionResponse) GetDashboardsPortOk() (*int32, bool)`

GetDashboardsPortOk returns a tuple with the DashboardsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardsPort

`func (o *SearchEngineInitConfigOptionResponse) SetDashboardsPort(v int32)`

SetDashboardsPort sets DashboardsPort field to given value.

### HasDashboardsPort

`func (o *SearchEngineInitConfigOptionResponse) HasDashboardsPort() bool`

HasDashboardsPort returns a boolean if a field has been set.

### SetDashboardsPortNil

`func (o *SearchEngineInitConfigOptionResponse) SetDashboardsPortNil(b bool)`

 SetDashboardsPortNil sets the value for DashboardsPort to be an explicit nil

### UnsetDashboardsPort
`func (o *SearchEngineInitConfigOptionResponse) UnsetDashboardsPort()`

UnsetDashboardsPort ensures that no value is present for DashboardsPort, not even an explicit nil
### GetDatabasePort

`func (o *SearchEngineInitConfigOptionResponse) GetDatabasePort() int32`

GetDatabasePort returns the DatabasePort field if non-nil, zero value otherwise.

### GetDatabasePortOk

`func (o *SearchEngineInitConfigOptionResponse) GetDatabasePortOk() (*int32, bool)`

GetDatabasePortOk returns a tuple with the DatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabasePort

`func (o *SearchEngineInitConfigOptionResponse) SetDatabasePort(v int32)`

SetDatabasePort sets DatabasePort field to given value.

### HasDatabasePort

`func (o *SearchEngineInitConfigOptionResponse) HasDatabasePort() bool`

HasDatabasePort returns a boolean if a field has been set.

### SetDatabasePortNil

`func (o *SearchEngineInitConfigOptionResponse) SetDatabasePortNil(b bool)`

 SetDatabasePortNil sets the value for DatabasePort to be an explicit nil

### UnsetDatabasePort
`func (o *SearchEngineInitConfigOptionResponse) UnsetDatabasePort()`

UnsetDatabasePort ensures that no value is present for DatabasePort, not even an explicit nil
### GetDatabaseUserName

`func (o *SearchEngineInitConfigOptionResponse) GetDatabaseUserName() string`

GetDatabaseUserName returns the DatabaseUserName field if non-nil, zero value otherwise.

### GetDatabaseUserNameOk

`func (o *SearchEngineInitConfigOptionResponse) GetDatabaseUserNameOk() (*string, bool)`

GetDatabaseUserNameOk returns a tuple with the DatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseUserName

`func (o *SearchEngineInitConfigOptionResponse) SetDatabaseUserName(v string)`

SetDatabaseUserName sets DatabaseUserName field to given value.


### GetKibanaPort

`func (o *SearchEngineInitConfigOptionResponse) GetKibanaPort() int32`

GetKibanaPort returns the KibanaPort field if non-nil, zero value otherwise.

### GetKibanaPortOk

`func (o *SearchEngineInitConfigOptionResponse) GetKibanaPortOk() (*int32, bool)`

GetKibanaPortOk returns a tuple with the KibanaPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKibanaPort

`func (o *SearchEngineInitConfigOptionResponse) SetKibanaPort(v int32)`

SetKibanaPort sets KibanaPort field to given value.

### HasKibanaPort

`func (o *SearchEngineInitConfigOptionResponse) HasKibanaPort() bool`

HasKibanaPort returns a boolean if a field has been set.

### SetKibanaPortNil

`func (o *SearchEngineInitConfigOptionResponse) SetKibanaPortNil(b bool)`

 SetKibanaPortNil sets the value for KibanaPort to be an explicit nil

### UnsetKibanaPort
`func (o *SearchEngineInitConfigOptionResponse) UnsetKibanaPort()`

UnsetKibanaPort ensures that no value is present for KibanaPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


