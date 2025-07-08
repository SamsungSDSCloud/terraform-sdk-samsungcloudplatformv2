# AgentBackupRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilesystemPaths** | [**[]RestoreFilesystemPath**](RestoreFilesystemPath.md) | List of filesystem backup path | 
**Overwrite** | Pointer to **NullableBool** |  | [optional] 
**RestoreServerUuid** | **string** | Restored server ID | 
**RestoreTargetId** | **string** | Restore target ID | 

## Methods

### NewAgentBackupRestoreRequest

`func NewAgentBackupRestoreRequest(filesystemPaths []RestoreFilesystemPath, restoreServerUuid string, restoreTargetId string, ) *AgentBackupRestoreRequest`

NewAgentBackupRestoreRequest instantiates a new AgentBackupRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentBackupRestoreRequestWithDefaults

`func NewAgentBackupRestoreRequestWithDefaults() *AgentBackupRestoreRequest`

NewAgentBackupRestoreRequestWithDefaults instantiates a new AgentBackupRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilesystemPaths

`func (o *AgentBackupRestoreRequest) GetFilesystemPaths() []RestoreFilesystemPath`

GetFilesystemPaths returns the FilesystemPaths field if non-nil, zero value otherwise.

### GetFilesystemPathsOk

`func (o *AgentBackupRestoreRequest) GetFilesystemPathsOk() (*[]RestoreFilesystemPath, bool)`

GetFilesystemPathsOk returns a tuple with the FilesystemPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemPaths

`func (o *AgentBackupRestoreRequest) SetFilesystemPaths(v []RestoreFilesystemPath)`

SetFilesystemPaths sets FilesystemPaths field to given value.


### GetOverwrite

`func (o *AgentBackupRestoreRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *AgentBackupRestoreRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *AgentBackupRestoreRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *AgentBackupRestoreRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *AgentBackupRestoreRequest) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *AgentBackupRestoreRequest) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetRestoreServerUuid

`func (o *AgentBackupRestoreRequest) GetRestoreServerUuid() string`

GetRestoreServerUuid returns the RestoreServerUuid field if non-nil, zero value otherwise.

### GetRestoreServerUuidOk

`func (o *AgentBackupRestoreRequest) GetRestoreServerUuidOk() (*string, bool)`

GetRestoreServerUuidOk returns a tuple with the RestoreServerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServerUuid

`func (o *AgentBackupRestoreRequest) SetRestoreServerUuid(v string)`

SetRestoreServerUuid sets RestoreServerUuid field to given value.


### GetRestoreTargetId

`func (o *AgentBackupRestoreRequest) GetRestoreTargetId() string`

GetRestoreTargetId returns the RestoreTargetId field if non-nil, zero value otherwise.

### GetRestoreTargetIdOk

`func (o *AgentBackupRestoreRequest) GetRestoreTargetIdOk() (*string, bool)`

GetRestoreTargetIdOk returns a tuple with the RestoreTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTargetId

`func (o *AgentBackupRestoreRequest) SetRestoreTargetId(v string)`

SetRestoreTargetId sets RestoreTargetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


