# SqlserverUpdateDatabasesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddDatabases** | Pointer to [**[]SqlserverDatabaseOption**](SqlserverDatabaseOption.md) |  | [optional] 
**DelDatabases** | Pointer to [**[]SqlserverDatabaseOption**](SqlserverDatabaseOption.md) |  | [optional] 

## Methods

### NewSqlserverUpdateDatabasesRequest

`func NewSqlserverUpdateDatabasesRequest() *SqlserverUpdateDatabasesRequest`

NewSqlserverUpdateDatabasesRequest instantiates a new SqlserverUpdateDatabasesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverUpdateDatabasesRequestWithDefaults

`func NewSqlserverUpdateDatabasesRequestWithDefaults() *SqlserverUpdateDatabasesRequest`

NewSqlserverUpdateDatabasesRequestWithDefaults instantiates a new SqlserverUpdateDatabasesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddDatabases

`func (o *SqlserverUpdateDatabasesRequest) GetAddDatabases() []SqlserverDatabaseOption`

GetAddDatabases returns the AddDatabases field if non-nil, zero value otherwise.

### GetAddDatabasesOk

`func (o *SqlserverUpdateDatabasesRequest) GetAddDatabasesOk() (*[]SqlserverDatabaseOption, bool)`

GetAddDatabasesOk returns a tuple with the AddDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDatabases

`func (o *SqlserverUpdateDatabasesRequest) SetAddDatabases(v []SqlserverDatabaseOption)`

SetAddDatabases sets AddDatabases field to given value.

### HasAddDatabases

`func (o *SqlserverUpdateDatabasesRequest) HasAddDatabases() bool`

HasAddDatabases returns a boolean if a field has been set.

### SetAddDatabasesNil

`func (o *SqlserverUpdateDatabasesRequest) SetAddDatabasesNil(b bool)`

 SetAddDatabasesNil sets the value for AddDatabases to be an explicit nil

### UnsetAddDatabases
`func (o *SqlserverUpdateDatabasesRequest) UnsetAddDatabases()`

UnsetAddDatabases ensures that no value is present for AddDatabases, not even an explicit nil
### GetDelDatabases

`func (o *SqlserverUpdateDatabasesRequest) GetDelDatabases() []SqlserverDatabaseOption`

GetDelDatabases returns the DelDatabases field if non-nil, zero value otherwise.

### GetDelDatabasesOk

`func (o *SqlserverUpdateDatabasesRequest) GetDelDatabasesOk() (*[]SqlserverDatabaseOption, bool)`

GetDelDatabasesOk returns a tuple with the DelDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelDatabases

`func (o *SqlserverUpdateDatabasesRequest) SetDelDatabases(v []SqlserverDatabaseOption)`

SetDelDatabases sets DelDatabases field to given value.

### HasDelDatabases

`func (o *SqlserverUpdateDatabasesRequest) HasDelDatabases() bool`

HasDelDatabases returns a boolean if a field has been set.

### SetDelDatabasesNil

`func (o *SqlserverUpdateDatabasesRequest) SetDelDatabasesNil(b bool)`

 SetDelDatabasesNil sets the value for DelDatabases to be an explicit nil

### UnsetDelDatabases
`func (o *SqlserverUpdateDatabasesRequest) UnsetDelDatabases()`

UnsetDelDatabases ensures that no value is present for DelDatabases, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


