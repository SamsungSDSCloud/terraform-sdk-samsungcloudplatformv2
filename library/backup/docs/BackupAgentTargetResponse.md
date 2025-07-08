# BackupAgentTargetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerBackupAgentIp** | **string** | Server&#39;s Backup Agent IP | 
**ServerName** | **string** | Backup Agent target server name | 
**ServerState** | **string** | Backup server state | 
**ServerUuid** | **string** | Backup Agent target server UUID | 

## Methods

### NewBackupAgentTargetResponse

`func NewBackupAgentTargetResponse(serverBackupAgentIp string, serverName string, serverState string, serverUuid string, ) *BackupAgentTargetResponse`

NewBackupAgentTargetResponse instantiates a new BackupAgentTargetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupAgentTargetResponseWithDefaults

`func NewBackupAgentTargetResponseWithDefaults() *BackupAgentTargetResponse`

NewBackupAgentTargetResponseWithDefaults instantiates a new BackupAgentTargetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerBackupAgentIp

`func (o *BackupAgentTargetResponse) GetServerBackupAgentIp() string`

GetServerBackupAgentIp returns the ServerBackupAgentIp field if non-nil, zero value otherwise.

### GetServerBackupAgentIpOk

`func (o *BackupAgentTargetResponse) GetServerBackupAgentIpOk() (*string, bool)`

GetServerBackupAgentIpOk returns a tuple with the ServerBackupAgentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBackupAgentIp

`func (o *BackupAgentTargetResponse) SetServerBackupAgentIp(v string)`

SetServerBackupAgentIp sets ServerBackupAgentIp field to given value.


### GetServerName

`func (o *BackupAgentTargetResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *BackupAgentTargetResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *BackupAgentTargetResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerState

`func (o *BackupAgentTargetResponse) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *BackupAgentTargetResponse) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *BackupAgentTargetResponse) SetServerState(v string)`

SetServerState sets ServerState field to given value.


### GetServerUuid

`func (o *BackupAgentTargetResponse) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *BackupAgentTargetResponse) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *BackupAgentTargetResponse) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


