# SqlserverDatabaseOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | Database Name | 
**DriveLetter** | Pointer to **string** | Drive letter | [optional] [default to "E"]

## Methods

### NewSqlserverDatabaseOption

`func NewSqlserverDatabaseOption(databaseName string, ) *SqlserverDatabaseOption`

NewSqlserverDatabaseOption instantiates a new SqlserverDatabaseOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverDatabaseOptionWithDefaults

`func NewSqlserverDatabaseOptionWithDefaults() *SqlserverDatabaseOption`

NewSqlserverDatabaseOptionWithDefaults instantiates a new SqlserverDatabaseOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *SqlserverDatabaseOption) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *SqlserverDatabaseOption) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *SqlserverDatabaseOption) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetDriveLetter

`func (o *SqlserverDatabaseOption) GetDriveLetter() string`

GetDriveLetter returns the DriveLetter field if non-nil, zero value otherwise.

### GetDriveLetterOk

`func (o *SqlserverDatabaseOption) GetDriveLetterOk() (*string, bool)`

GetDriveLetterOk returns a tuple with the DriveLetter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveLetter

`func (o *SqlserverDatabaseOption) SetDriveLetter(v string)`

SetDriveLetter sets DriveLetter field to given value.

### HasDriveLetter

`func (o *SqlserverDatabaseOption) HasDriveLetter() bool`

HasDriveLetter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


