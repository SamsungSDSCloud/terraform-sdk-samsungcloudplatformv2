# BackupRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Networks** | Pointer to [**[]Network**](Network.md) |  | [optional] 
**RestoreRegion** | Pointer to **NullableString** |  | [optional] 
**RestoreServerName** | **string** | Restore server name | 
**RestoreTargetId** | **string** | Restore target ID | 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**ServerTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBackupRestoreRequest

`func NewBackupRestoreRequest(restoreServerName string, restoreTargetId string, ) *BackupRestoreRequest`

NewBackupRestoreRequest instantiates a new BackupRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreRequestWithDefaults

`func NewBackupRestoreRequestWithDefaults() *BackupRestoreRequest`

NewBackupRestoreRequestWithDefaults instantiates a new BackupRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworks

`func (o *BackupRestoreRequest) GetNetworks() []Network`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *BackupRestoreRequest) GetNetworksOk() (*[]Network, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *BackupRestoreRequest) SetNetworks(v []Network)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *BackupRestoreRequest) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### SetNetworksNil

`func (o *BackupRestoreRequest) SetNetworksNil(b bool)`

 SetNetworksNil sets the value for Networks to be an explicit nil

### UnsetNetworks
`func (o *BackupRestoreRequest) UnsetNetworks()`

UnsetNetworks ensures that no value is present for Networks, not even an explicit nil
### GetRestoreRegion

`func (o *BackupRestoreRequest) GetRestoreRegion() string`

GetRestoreRegion returns the RestoreRegion field if non-nil, zero value otherwise.

### GetRestoreRegionOk

`func (o *BackupRestoreRequest) GetRestoreRegionOk() (*string, bool)`

GetRestoreRegionOk returns a tuple with the RestoreRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreRegion

`func (o *BackupRestoreRequest) SetRestoreRegion(v string)`

SetRestoreRegion sets RestoreRegion field to given value.

### HasRestoreRegion

`func (o *BackupRestoreRequest) HasRestoreRegion() bool`

HasRestoreRegion returns a boolean if a field has been set.

### SetRestoreRegionNil

`func (o *BackupRestoreRequest) SetRestoreRegionNil(b bool)`

 SetRestoreRegionNil sets the value for RestoreRegion to be an explicit nil

### UnsetRestoreRegion
`func (o *BackupRestoreRequest) UnsetRestoreRegion()`

UnsetRestoreRegion ensures that no value is present for RestoreRegion, not even an explicit nil
### GetRestoreServerName

`func (o *BackupRestoreRequest) GetRestoreServerName() string`

GetRestoreServerName returns the RestoreServerName field if non-nil, zero value otherwise.

### GetRestoreServerNameOk

`func (o *BackupRestoreRequest) GetRestoreServerNameOk() (*string, bool)`

GetRestoreServerNameOk returns a tuple with the RestoreServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreServerName

`func (o *BackupRestoreRequest) SetRestoreServerName(v string)`

SetRestoreServerName sets RestoreServerName field to given value.


### GetRestoreTargetId

`func (o *BackupRestoreRequest) GetRestoreTargetId() string`

GetRestoreTargetId returns the RestoreTargetId field if non-nil, zero value otherwise.

### GetRestoreTargetIdOk

`func (o *BackupRestoreRequest) GetRestoreTargetIdOk() (*string, bool)`

GetRestoreTargetIdOk returns a tuple with the RestoreTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTargetId

`func (o *BackupRestoreRequest) SetRestoreTargetId(v string)`

SetRestoreTargetId sets RestoreTargetId field to given value.


### GetSecurityGroups

`func (o *BackupRestoreRequest) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BackupRestoreRequest) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BackupRestoreRequest) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BackupRestoreRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BackupRestoreRequest) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BackupRestoreRequest) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetServerTypeId

`func (o *BackupRestoreRequest) GetServerTypeId() string`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *BackupRestoreRequest) GetServerTypeIdOk() (*string, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *BackupRestoreRequest) SetServerTypeId(v string)`

SetServerTypeId sets ServerTypeId field to given value.

### HasServerTypeId

`func (o *BackupRestoreRequest) HasServerTypeId() bool`

HasServerTypeId returns a boolean if a field has been set.

### SetServerTypeIdNil

`func (o *BackupRestoreRequest) SetServerTypeIdNil(b bool)`

 SetServerTypeIdNil sets the value for ServerTypeId to be an explicit nil

### UnsetServerTypeId
`func (o *BackupRestoreRequest) UnsetServerTypeId()`

UnsetServerTypeId ensures that no value is present for ServerTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


