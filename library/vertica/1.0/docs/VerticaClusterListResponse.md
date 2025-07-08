# VerticaClusterListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**ConsoleIncluded** | Pointer to **bool** | Management Console included state | [optional] [default to false]
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DatabaseName** | **NullableString** |  | 
**Id** | **string** | ID | 
**InstanceCount** | Pointer to **int32** | Instance Count | [optional] [default to 0]
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **NullableBool** |  | [optional] 
**RoleType** | [**NullableClusterRoleType**](ClusterRoleType.md) |  | 
**ServiceState** | [**ServiceState**](ServiceState.md) | Service state | 

## Methods

### NewVerticaClusterListResponse

`func NewVerticaClusterListResponse(accountId string, createdAt time.Time, createdBy string, databaseName NullableString, id string, modifiedAt time.Time, modifiedBy string, name string, roleType NullableClusterRoleType, serviceState ServiceState, ) *VerticaClusterListResponse`

NewVerticaClusterListResponse instantiates a new VerticaClusterListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerticaClusterListResponseWithDefaults

`func NewVerticaClusterListResponseWithDefaults() *VerticaClusterListResponse`

NewVerticaClusterListResponseWithDefaults instantiates a new VerticaClusterListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VerticaClusterListResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VerticaClusterListResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VerticaClusterListResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetConsoleIncluded

`func (o *VerticaClusterListResponse) GetConsoleIncluded() bool`

GetConsoleIncluded returns the ConsoleIncluded field if non-nil, zero value otherwise.

### GetConsoleIncludedOk

`func (o *VerticaClusterListResponse) GetConsoleIncludedOk() (*bool, bool)`

GetConsoleIncludedOk returns a tuple with the ConsoleIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleIncluded

`func (o *VerticaClusterListResponse) SetConsoleIncluded(v bool)`

SetConsoleIncluded sets ConsoleIncluded field to given value.

### HasConsoleIncluded

`func (o *VerticaClusterListResponse) HasConsoleIncluded() bool`

HasConsoleIncluded returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VerticaClusterListResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VerticaClusterListResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VerticaClusterListResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VerticaClusterListResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VerticaClusterListResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VerticaClusterListResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDatabaseName

`func (o *VerticaClusterListResponse) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *VerticaClusterListResponse) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *VerticaClusterListResponse) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### SetDatabaseNameNil

`func (o *VerticaClusterListResponse) SetDatabaseNameNil(b bool)`

 SetDatabaseNameNil sets the value for DatabaseName to be an explicit nil

### UnsetDatabaseName
`func (o *VerticaClusterListResponse) UnsetDatabaseName()`

UnsetDatabaseName ensures that no value is present for DatabaseName, not even an explicit nil
### GetId

`func (o *VerticaClusterListResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerticaClusterListResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerticaClusterListResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceCount

`func (o *VerticaClusterListResponse) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *VerticaClusterListResponse) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *VerticaClusterListResponse) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *VerticaClusterListResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetModifiedAt

`func (o *VerticaClusterListResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *VerticaClusterListResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *VerticaClusterListResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *VerticaClusterListResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *VerticaClusterListResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *VerticaClusterListResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *VerticaClusterListResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VerticaClusterListResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VerticaClusterListResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *VerticaClusterListResponse) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *VerticaClusterListResponse) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *VerticaClusterListResponse) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *VerticaClusterListResponse) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### SetNatEnabledNil

`func (o *VerticaClusterListResponse) SetNatEnabledNil(b bool)`

 SetNatEnabledNil sets the value for NatEnabled to be an explicit nil

### UnsetNatEnabled
`func (o *VerticaClusterListResponse) UnsetNatEnabled()`

UnsetNatEnabled ensures that no value is present for NatEnabled, not even an explicit nil
### GetRoleType

`func (o *VerticaClusterListResponse) GetRoleType() ClusterRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *VerticaClusterListResponse) GetRoleTypeOk() (*ClusterRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *VerticaClusterListResponse) SetRoleType(v ClusterRoleType)`

SetRoleType sets RoleType field to given value.


### SetRoleTypeNil

`func (o *VerticaClusterListResponse) SetRoleTypeNil(b bool)`

 SetRoleTypeNil sets the value for RoleType to be an explicit nil

### UnsetRoleType
`func (o *VerticaClusterListResponse) UnsetRoleType()`

UnsetRoleType ensures that no value is present for RoleType, not even an explicit nil
### GetServiceState

`func (o *VerticaClusterListResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *VerticaClusterListResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *VerticaClusterListResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


