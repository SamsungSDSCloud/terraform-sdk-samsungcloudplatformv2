# AgentBackupRestoreTargetServerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerName** | **string** | Backup server name | 
**ServerOsVersion** | **string** | Server&#39;s OS Type and Version | 
**ServerUuid** | **string** | Backup server UUID | 

## Methods

### NewAgentBackupRestoreTargetServerResponse

`func NewAgentBackupRestoreTargetServerResponse(serverName string, serverOsVersion string, serverUuid string, ) *AgentBackupRestoreTargetServerResponse`

NewAgentBackupRestoreTargetServerResponse instantiates a new AgentBackupRestoreTargetServerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentBackupRestoreTargetServerResponseWithDefaults

`func NewAgentBackupRestoreTargetServerResponseWithDefaults() *AgentBackupRestoreTargetServerResponse`

NewAgentBackupRestoreTargetServerResponseWithDefaults instantiates a new AgentBackupRestoreTargetServerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerName

`func (o *AgentBackupRestoreTargetServerResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *AgentBackupRestoreTargetServerResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *AgentBackupRestoreTargetServerResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetServerOsVersion

`func (o *AgentBackupRestoreTargetServerResponse) GetServerOsVersion() string`

GetServerOsVersion returns the ServerOsVersion field if non-nil, zero value otherwise.

### GetServerOsVersionOk

`func (o *AgentBackupRestoreTargetServerResponse) GetServerOsVersionOk() (*string, bool)`

GetServerOsVersionOk returns a tuple with the ServerOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOsVersion

`func (o *AgentBackupRestoreTargetServerResponse) SetServerOsVersion(v string)`

SetServerOsVersion sets ServerOsVersion field to given value.


### GetServerUuid

`func (o *AgentBackupRestoreTargetServerResponse) GetServerUuid() string`

GetServerUuid returns the ServerUuid field if non-nil, zero value otherwise.

### GetServerUuidOk

`func (o *AgentBackupRestoreTargetServerResponse) GetServerUuidOk() (*string, bool)`

GetServerUuidOk returns a tuple with the ServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUuid

`func (o *AgentBackupRestoreTargetServerResponse) SetServerUuid(v string)`

SetServerUuid sets ServerUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


