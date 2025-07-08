# BackupAgentDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** |  | 
**BackupMasterIp** | **NullableString** |  | 
**BackupMasterName** | **NullableString** |  | 
**BackupPolicyCount** | **NullableInt32** |  | 
**ConnectionCheckTime** | **NullableTime** |  | 
**ConnectionState** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **NullableString** |  | 
**InstallFilePath** | **string** | Install file path | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **NullableString** |  | 
**ServerBackupAgentIp** | **NullableString** |  | 
**ServerCategory** | **NullableString** |  | 
**ServerGatewayIp** | **NullableString** |  | 
**ServerName** | **NullableString** |  | 
**ServerOsType** | **NullableString** |  | 
**ServerUuid** | **NullableString** |  | 
**State** | **NullableString** |  | 

## Methods

### NewBackupAgentDetailResponse

`func NewBackupAgentDetailResponse(accountId NullableString, backupMasterIp NullableString, backupMasterName NullableString, backupPolicyCount NullableInt32, connectionCheckTime NullableTime, connectionState NullableString, createdAt time.Time, createdBy string, id NullableString, installFilePath string, modifiedAt time.Time, modifiedBy string, name NullableString, serverBackupAgentIp NullableString, serverCategory NullableString, serverGatewayIp NullableString, serverName NullableString, serverOsType NullableString, serverUuid NullableString, state NullableString, ) *BackupAgentDetailResponse`

NewBackupAgentDetailResponse instantiates a new BackupAgentDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupAgentDetailResponseWithDefaults

`func NewBackupAgentDetailResponseWithDefaults() *BackupAgentDetailResponse`

NewBackupAgentDetailResponseWithDefaults instantiates a new BackupAgentDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BackupAgentDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BackupAgentDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BackupAgentDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *BackupAgentDetailResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *BackupAgentDetailResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetBackupMasterIp

`func (o *BackupAgentDetailResponse) GetBackupMasterIp() string`

GetBackupMasterIp returns the BackupMasterIp field if non-nil, zero value otherwise.

### GetBackupMasterIpOk

`func (o *BackupAgentDetailResponse) GetBackupMasterIpOk() (*string, bool)`

GetBackupMasterIpOk returns a tuple with the BackupMasterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMasterIp

`func (o *BackupAgentDetailResponse) SetBackupMasterIp(v string)`

SetBackupMasterIp sets BackupMasterIp field to given value.


### SetBackupMasterIpNil

`func (o *BackupAgentDetailResponse) SetBackupMasterIpNil(b bool)`

 SetBackupMasterIpNil sets the value for BackupMasterIp to be an explicit nil

### UnsetBackupMasterIp
`func (o *BackupAgentDetailResponse) UnsetBackupMasterIp()`

UnsetBackupMasterIp ensures that no value is present for BackupMasterIp, not even an explicit nil
### GetBackupMasterName

`func (o *BackupAgentDetailResponse) GetBackupMasterName() string`

GetBackupMasterName returns the BackupMasterName field if non-nil, zero value otherwise.

### GetBackupMasterNameOk

`func (o *BackupAgentDetailResponse) GetBackupMasterNameOk() (*string, bool)`

GetBackupMasterNameOk returns a tuple with the BackupMasterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMasterName

`func (o *BackupAgentDetailResponse) SetBackupMasterName(v string)`

SetBackupMasterName sets BackupMasterName field to given value.


### SetBackupMasterNameNil

`func (o *BackupAgentDetailResponse) SetBackupMasterNameNil(b bool)`

 SetBackupMasterNameNil sets the value for BackupMasterName to be an explicit nil

### UnsetBackupMasterName
`func (o *BackupAgentDetailResponse) UnsetBackupMasterName()`

UnsetBackupMasterName ensures that no value is present for BackupMasterName, not even an explicit nil
### GetBackupPolicyCount

`func (o *BackupAgentDetailResponse) GetBackupPolicyCount() int32`

GetBackupPolicyCount returns the BackupPolicyCount field if non-nil, zero value otherwise.

### GetBackupPolicyCountOk

`func (o *BackupAgentDetailResponse) GetBackupPolicyCountOk() (*int32, bool)`

GetBackupPolicyCountOk returns a tuple with the BackupPolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPolicyCount

`func (o *BackupAgentDetailResponse) SetBackupPolicyCount(v int32)`

SetBackupPolicyCount sets BackupPolicyCount field to given value.


### SetBackupPolicyCountNil

`func (o *BackupAgentDetailResponse) SetBackupPolicyCountNil(b bool)`

 SetBackupPolicyCountNil sets the value for BackupPolicyCount to be an explicit nil

### UnsetBackupPolicyCount
`func (o *BackupAgentDetailResponse) UnsetBackupPolicyCount()`

UnsetBackupPolicyCount ensures that no value is present for BackupPolicyCount, not even an explicit nil
### GetConnectionCheckTime

`func (o *BackupAgentDetailResponse) GetConnectionCheckTime() time.Time`

GetConnectionCheckTime returns the ConnectionCheckTime field if non-nil, zero value otherwise.

### GetConnectionCheckTimeOk

`func (o *BackupAgentDetailResponse) GetConnectionCheckTimeOk() (*time.Time, bool)`

GetConnectionCheckTimeOk returns a tuple with the ConnectionCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionCheckTime

`func (o *BackupAgentDetailResponse) SetConnectionCheckTime(v time.Time)`

SetConnectionCheckTime sets ConnectionCheckTime field to given value.


### SetConnectionCheckTimeNil

`func (o *BackupAgentDetailResponse) SetConnectionCheckTimeNil(b bool)`

 SetConnectionCheckTimeNil sets the value for ConnectionCheckTime to be an explicit nil

### UnsetConnectionCheckTime
`func (o *BackupAgentDetailResponse) UnsetConnectionCheckTime()`

UnsetConnectionCheckTime ensures that no value is present for ConnectionCheckTime, not even an explicit nil
### GetConnectionState

`func (o *BackupAgentDetailResponse) GetConnectionState() string`

GetConnectionState returns the ConnectionState field if non-nil, zero value otherwise.

### GetConnectionStateOk

`func (o *BackupAgentDetailResponse) GetConnectionStateOk() (*string, bool)`

GetConnectionStateOk returns a tuple with the ConnectionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionState

`func (o *BackupAgentDetailResponse) SetConnectionState(v string)`

SetConnectionState sets ConnectionState field to given value.


### SetConnectionStateNil

`func (o *BackupAgentDetailResponse) SetConnectionStateNil(b bool)`

 SetConnectionStateNil sets the value for ConnectionState to be an explicit nil

### UnsetConnectionState
`func (o *BackupAgentDetailResponse) UnsetConnectionState()`

UnsetConnectionState ensures that no value is present for ConnectionState, not even an explicit nil
### GetCreatedAt

`func (o *BackupAgentDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackupAgentDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackupAgentDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *BackupAgentDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BackupAgentDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BackupAgentDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *BackupAgentDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupAgentDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupAgentDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *BackupAgentDetailResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BackupAgentDetailResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInstallFilePath

`func (o *BackupAgentDetailResponse) GetInstallFilePath() string`

GetInstallFilePath returns the InstallFilePath field if non-nil, zero value otherwise.

### GetInstallFilePathOk

`func (o *BackupAgentDetailResponse) GetInstallFilePathOk() (*string, bool)`

GetInstallFilePathOk returns a tuple with the InstallFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallFilePath

`func (o *BackupAgentDetailResponse) SetInstallFilePath(v string)`

SetInstallFilePath sets InstallFilePath field to given value.


### GetModifiedAt

`func (o *BackupAgentDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *BackupAgentDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *BackupAgentDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *BackupAgentDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *BackupAgentDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *BackupAgentDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *BackupAgentDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupAgentDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupAgentDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BackupAgentDetailResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BackupAgentDetailResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetServerBackupAgentIp

`func (o *BackupAgentDetailResponse) GetServerBackupAgentIp() string`

GetServerBackupAgentIp returns the ServerBackupAgentIp field if non-nil, zero value otherwise.

### GetServerBackupAgentIpOk

`func (o *BackupAgentDetailResponse) GetServerBackupAgentIpOk() (*string, bool)`

GetServerBackupAgentIpOk returns a tuple with the ServerBackupAgentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBackupAgentIp

`func (o *BackupAgentDetailResponse) SetServerBackupAgentIp(v string)`

SetServerBackupAgentIp sets ServerBackupAgentIp field to given value.


### SetServerBackupAgentIpNil

`func (o *BackupAgentDetailResponse) SetServerBackupAgentIpNil(b bool)`

 SetServerBackupAgentIpNil sets the value for ServerBackupAgentIp to be an explicit nil

### UnsetServerBackupAgentIp
`func (o *BackupAgentDetailResponse) UnsetServerBackupAgentIp()`

UnsetServerBackupAgentIp ensures that no value is present for ServerBackupAgentIp, not even an explicit nil
### GetServerCategory

`func (o *BackupAgentDetailResponse) GetServerCategory() string`

GetServerCategory returns the ServerCategory field if non-nil, zero value otherwise.

### GetServerCategoryOk

`func (o *BackupAgentDetailResponse) GetServerCategoryOk() (*string, bool)`

GetServerCategoryOk returns a tuple with the ServerCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCategory

`func (o *BackupAgentDetailResponse) SetServerCategory(v string)`

SetServerCategory sets ServerCategory field to given value.


### SetServerCategoryNil

`func (o *BackupAgentDetailResponse) SetServerCategoryNil(b bool)`

 SetServerCategoryNil sets the value for ServerCategory to be an explicit nil

### UnsetServerCategory
`func (o *BackupAgentDetailResponse) UnsetServerCategory()`

UnsetServerCategory ensures that no value is present for ServerCategory, not even an explicit nil
### GetServerGatewayIp

`func (o *BackupAgentDetailResponse) GetServerGatewayIp() string`

GetServerGatewayIp returns the ServerGatewayIp field if non-nil, zero value otherwise.

### GetServerGatewayIpOk

`func (o *BackupAgentDetailResponse) GetServerGatewayIpOk() (*string, bool)`

GetServerGatewayIpOk returns a tuple with the ServerGatewayIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGatewayIp

`func (o *BackupAgentDetailResponse) SetServerGatewayIp(v string)`

SetServerGatewayIp sets ServerGatewayIp field to given value.


### SetServerGatewayIpNil

`func (o *BackupAgentDetailResponse) SetServerGatewayIpNil(b bool)`

 SetServerGatewayIpNil sets the value for ServerGatewayIp to be an explicit nil

### UnsetServerGatewayIp
`func (o *BackupAgentDetailResponse) UnsetServerGatewayIp()`

UnsetServerGatewayIp ensures that no value is present for ServerGatewayIp, not even an explicit nil
### GetServerName

`func (o *BackupAgentDetailResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *BackupAgentDetailResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *BackupAgentDetailResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### SetServerNameNil

`func (o *BackupAgentDetailResponse) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *BackupAgentDetailResponse) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetServerOsType

`func (o *BackupAgentDetailResponse) GetServerOsType() string`

GetServerOsType returns the ServerOsType field if non-nil, zero value otherwise.

### GetServerOsTypeOk

`func (o *BackupAgentDetailResponse) GetServerOsTypeOk() (*string, bool)`

GetServerOsTypeOk returns a tuple with the ServerOsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOsType

`func (o *BackupAgentDetailResponse) SetServerOsType(v string)`

SetServerOsType sets ServerOsType field to given value.


### SetServerOsTypeNil

`func (o *BackupAgentDetailResponse) SetServerOsTypeNil(b bool)`

 SetServerOsTypeNil sets the value for ServerOsType to be an explicit nil

### UnsetServerOsType
`func (o *BackupAgentDetailResponse) UnsetServerOsType()`

UnsetServerOsType ensures that no value is present for ServerOsType, not even an explicit nil
### GetServerUuid

`func (o *BackupAgentDetailResponse) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *BackupAgentDetailResponse) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *BackupAgentDetailResponse) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.


### SetServerUuidNil

`func (o *BackupAgentDetailResponse) SetServerUuidNil(b bool)`

 SetServerUuidNil sets the value for ServerUuid to be an explicit nil

### UnsetServerUuid
`func (o *BackupAgentDetailResponse) UnsetServerUuid()`

UnsetServerUuid ensures that no value is present for ServerUuid, not even an explicit nil
### GetState

`func (o *BackupAgentDetailResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BackupAgentDetailResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BackupAgentDetailResponse) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *BackupAgentDetailResponse) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *BackupAgentDetailResponse) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


