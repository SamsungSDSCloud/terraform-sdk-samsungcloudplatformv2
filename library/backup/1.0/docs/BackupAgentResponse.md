# BackupAgentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** |  | 
**BackupPolicyCount** | **NullableInt32** |  | 
**ConnectionState** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **NullableString** |  | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **NullableString** |  | 
**ServerBackupAgentIp** | **NullableString** |  | 
**ServerName** | **NullableString** |  | 
**State** | **NullableString** |  | 

## Methods

### NewBackupAgentResponse

`func NewBackupAgentResponse(accountId NullableString, backupPolicyCount NullableInt32, connectionState NullableString, createdAt time.Time, createdBy string, id NullableString, modifiedAt time.Time, modifiedBy string, name NullableString, serverBackupAgentIp NullableString, serverName NullableString, state NullableString, ) *BackupAgentResponse`

NewBackupAgentResponse instantiates a new BackupAgentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupAgentResponseWithDefaults

`func NewBackupAgentResponseWithDefaults() *BackupAgentResponse`

NewBackupAgentResponseWithDefaults instantiates a new BackupAgentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BackupAgentResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BackupAgentResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BackupAgentResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *BackupAgentResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *BackupAgentResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetBackupPolicyCount

`func (o *BackupAgentResponse) GetBackupPolicyCount() int32`

GetBackupPolicyCount returns the BackupPolicyCount field if non-nil, zero value otherwise.

### GetBackupPolicyCountOk

`func (o *BackupAgentResponse) GetBackupPolicyCountOk() (*int32, bool)`

GetBackupPolicyCountOk returns a tuple with the BackupPolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicyCount

`func (o *BackupAgentResponse) SetBackupPolicyCount(v int32)`

SetBackupPolicyCount sets BackupPolicyCount field to given value.


### SetBackupPolicyCountNil

`func (o *BackupAgentResponse) SetBackupPolicyCountNil(b bool)`

 SetBackupPolicyCountNil sets the value for BackupPolicyCount to be an explicit nil

### UnsetBackupPolicyCount
`func (o *BackupAgentResponse) UnsetBackupPolicyCount()`

UnsetBackupPolicyCount ensures that no value is present for BackupPolicyCount, not even an explicit nil
### GetConnectionState

`func (o *BackupAgentResponse) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *BackupAgentResponse) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *BackupAgentResponse) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.


### SetConnectionStateNil

`func (o *BackupAgentResponse) SetConnectionStateNil(b bool)`

 SetConnectionStateNil sets the value for ConnectionState to be an explicit nil

### UnsetConnectionState
`func (o *BackupAgentResponse) UnsetConnectionState()`

UnsetConnectionState ensures that no value is present for ConnectionState, not even an explicit nil
### GetCreatedAt

`func (o *BackupAgentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupAgentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupAgentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *BackupAgentResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupAgentResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupAgentResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *BackupAgentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupAgentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupAgentResponse) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *BackupAgentResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BackupAgentResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetModifiedAt

`func (o *BackupAgentResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BackupAgentResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BackupAgentResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *BackupAgentResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BackupAgentResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BackupAgentResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *BackupAgentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupAgentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupAgentResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BackupAgentResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupAgentResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServerBackupAgentIp

`func (o *BackupAgentResponse) GetServerBackupAgentIp() string`

GetServerBackupAgentIp returns the ServerBackupAgentIp field if non-nil, zero value otherwise.

### GetServerBackupAgentIpOk

`func (o *BackupAgentResponse) GetServerBackupAgentIpOk() (*string, bool)`

GetServerBackupAgentIpOk returns a tuple with the ServerBackupAgentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBackupAgentIp

`func (o *BackupAgentResponse) SetServerBackupAgentIp(v string)`

SetServerBackupAgentIp sets ServerBackupAgentIp field to given value.


### SetServerBackupAgentIpNil

`func (o *BackupAgentResponse) SetServerBackupAgentIpNil(b bool)`

 SetServerBackupAgentIpNil sets the value for ServerBackupAgentIp to be an explicit nil

### UnsetServerBackupAgentIp
`func (o *BackupAgentResponse) UnsetServerBackupAgentIp()`

UnsetServerBackupAgentIp ensures that no value is present for ServerBackupAgentIp, not even an explicit nil
### GetServerName

`func (o *BackupAgentResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *BackupAgentResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *BackupAgentResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### SetServerNameNil

`func (o *BackupAgentResponse) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *BackupAgentResponse) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetState

`func (o *BackupAgentResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupAgentResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupAgentResponse) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *BackupAgentResponse) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BackupAgentResponse) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


