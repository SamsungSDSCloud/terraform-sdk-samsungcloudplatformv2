# ClusterMigrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceDatabaseName** | **string** | Database Name | 
**SourceDatabasePort** | **int32** | Database service port | 
**SourceDatabaseUserName** | **string** | Database User Name | 
**SourceDatabaseUserPassword** | **string** | Database user password | 
**SourceIpAddress** | **string** | User subnet IP address | 

## Methods

### NewClusterMigrationRequest

`func NewClusterMigrationRequest(sourceDatabaseName string, sourceDatabasePort int32, sourceDatabaseUserName string, sourceDatabaseUserPassword string, sourceIpAddress string, ) *ClusterMigrationRequest`

NewClusterMigrationRequest instantiates a new ClusterMigrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterMigrationRequestWithDefaults

`func NewClusterMigrationRequestWithDefaults() *ClusterMigrationRequest`

NewClusterMigrationRequestWithDefaults instantiates a new ClusterMigrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceDatabaseName

`func (o *ClusterMigrationRequest) GetSourceDatabaseName() string`

GetSourceDatabaseName returns the SourceDatabaseName field if non-nil, zero value otherwise.

### GetSourceDatabaseNameOk

`func (o *ClusterMigrationRequest) GetSourceDatabaseNameOk() (*string, bool)`

GetSourceDatabaseNameOk returns a tuple with the SourceDatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabaseName

`func (o *ClusterMigrationRequest) SetSourceDatabaseName(v string)`

SetSourceDatabaseName sets SourceDatabaseName field to given value.


### GetSourceDatabasePort

`func (o *ClusterMigrationRequest) GetSourceDatabasePort() int32`

GetSourceDatabasePort returns the SourceDatabasePort field if non-nil, zero value otherwise.

### GetSourceDatabasePortOk

`func (o *ClusterMigrationRequest) GetSourceDatabasePortOk() (*int32, bool)`

GetSourceDatabasePortOk returns a tuple with the SourceDatabasePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabasePort

`func (o *ClusterMigrationRequest) SetSourceDatabasePort(v int32)`

SetSourceDatabasePort sets SourceDatabasePort field to given value.


### GetSourceDatabaseUserName

`func (o *ClusterMigrationRequest) GetSourceDatabaseUserName() string`

GetSourceDatabaseUserName returns the SourceDatabaseUserName field if non-nil, zero value otherwise.

### GetSourceDatabaseUserNameOk

`func (o *ClusterMigrationRequest) GetSourceDatabaseUserNameOk() (*string, bool)`

GetSourceDatabaseUserNameOk returns a tuple with the SourceDatabaseUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabaseUserName

`func (o *ClusterMigrationRequest) SetSourceDatabaseUserName(v string)`

SetSourceDatabaseUserName sets SourceDatabaseUserName field to given value.


### GetSourceDatabaseUserPassword

`func (o *ClusterMigrationRequest) GetSourceDatabaseUserPassword() string`

GetSourceDatabaseUserPassword returns the SourceDatabaseUserPassword field if non-nil, zero value otherwise.

### GetSourceDatabaseUserPasswordOk

`func (o *ClusterMigrationRequest) GetSourceDatabaseUserPasswordOk() (*string, bool)`

GetSourceDatabaseUserPasswordOk returns a tuple with the SourceDatabaseUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDatabaseUserPassword

`func (o *ClusterMigrationRequest) SetSourceDatabaseUserPassword(v string)`

SetSourceDatabaseUserPassword sets SourceDatabaseUserPassword field to given value.


### GetSourceIpAddress

`func (o *ClusterMigrationRequest) GetSourceIpAddress() string`

GetSourceIpAddress returns the SourceIpAddress field if non-nil, zero value otherwise.

### GetSourceIpAddressOk

`func (o *ClusterMigrationRequest) GetSourceIpAddressOk() (*string, bool)`

GetSourceIpAddressOk returns a tuple with the SourceIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIpAddress

`func (o *ClusterMigrationRequest) SetSourceIpAddress(v string)`

SetSourceIpAddress sets SourceIpAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


