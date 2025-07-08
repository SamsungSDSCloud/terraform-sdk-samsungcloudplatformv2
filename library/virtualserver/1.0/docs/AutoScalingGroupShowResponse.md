# AutoScalingGroupShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DesiredServerCount** | **int32** | Desired server count | 
**DesiredServerCountEditable** | **bool** | Desired server count editable | 
**DrainEnabled** | **bool** | Drain enabled | 
**DrainTimeout** | Pointer to **NullableInt32** |  | [optional] 
**Id** | **string** | ID | 
**LaunchConfigurationId** | **string** | Launch Configuration ID | 
**LaunchConfigurationName** | **string** | Launch Configuration name | 
**LbServerGroups** | Pointer to [**[]AutoScalingGroupLbServerGroup**](AutoScalingGroupLbServerGroup.md) |  | [optional] 
**MaxServerCount** | **int32** | Max server count | 
**MinServerCount** | **int32** | Min server count | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Auto-Scaling Group name | 
**SecurityGroupIds** | Pointer to **[]string** |  | [optional] 
**ServerNamePrefix** | **string** | Server name prefix | 
**State** | **string** | Auto-Scaling Group state | 
**SubnetIds** | Pointer to **[]string** | Subnet ID list | [optional] [default to []]
**TotalServerCount** | **int32** | Total server count | 
**VpcId** | **string** | VPC ID | 

## Methods

### NewAutoScalingGroupShowResponse

`func NewAutoScalingGroupShowResponse(accountId string, createdAt time.Time, createdBy string, desiredServerCount int32, desiredServerCountEditable bool, drainEnabled bool, id string, launchConfigurationId string, launchConfigurationName string, maxServerCount int32, minServerCount int32, modifiedAt time.Time, modifiedBy string, name string, serverNamePrefix string, state string, totalServerCount int32, vpcId string, ) *AutoScalingGroupShowResponse`

NewAutoScalingGroupShowResponse instantiates a new AutoScalingGroupShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupShowResponseWithDefaults

`func NewAutoScalingGroupShowResponseWithDefaults() *AutoScalingGroupShowResponse`

NewAutoScalingGroupShowResponseWithDefaults instantiates a new AutoScalingGroupShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AutoScalingGroupShowResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AutoScalingGroupShowResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AutoScalingGroupShowResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *AutoScalingGroupShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AutoScalingGroupShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AutoScalingGroupShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AutoScalingGroupShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AutoScalingGroupShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AutoScalingGroupShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDesiredServerCount

`func (o *AutoScalingGroupShowResponse) GetDesiredServerCount() int32`

GetDesiredServerCount returns the DesiredServerCount field if non-nil, zero value otherwise.

### GetDesiredServerCountOk

`func (o *AutoScalingGroupShowResponse) GetDesiredServerCountOk() (*int32, bool)`

GetDesiredServerCountOk returns a tuple with the DesiredServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCount

`func (o *AutoScalingGroupShowResponse) SetDesiredServerCount(v int32)`

SetDesiredServerCount sets DesiredServerCount field to given value.


### GetDesiredServerCountEditable

`func (o *AutoScalingGroupShowResponse) GetDesiredServerCountEditable() bool`

GetDesiredServerCountEditable returns the DesiredServerCountEditable field if non-nil, zero value otherwise.

### GetDesiredServerCountEditableOk

`func (o *AutoScalingGroupShowResponse) GetDesiredServerCountEditableOk() (*bool, bool)`

GetDesiredServerCountEditableOk returns a tuple with the DesiredServerCountEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCountEditable

`func (o *AutoScalingGroupShowResponse) SetDesiredServerCountEditable(v bool)`

SetDesiredServerCountEditable sets DesiredServerCountEditable field to given value.


### GetDrainEnabled

`func (o *AutoScalingGroupShowResponse) GetDrainEnabled() bool`

GetDrainEnabled returns the DrainEnabled field if non-nil, zero value otherwise.

### GetDrainEnabledOk

`func (o *AutoScalingGroupShowResponse) GetDrainEnabledOk() (*bool, bool)`

GetDrainEnabledOk returns a tuple with the DrainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainEnabled

`func (o *AutoScalingGroupShowResponse) SetDrainEnabled(v bool)`

SetDrainEnabled sets DrainEnabled field to given value.


### GetDrainTimeout

`func (o *AutoScalingGroupShowResponse) GetDrainTimeout() int32`

GetDrainTimeout returns the DrainTimeout field if non-nil, zero value otherwise.

### GetDrainTimeoutOk

`func (o *AutoScalingGroupShowResponse) GetDrainTimeoutOk() (*int32, bool)`

GetDrainTimeoutOk returns a tuple with the DrainTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainTimeout

`func (o *AutoScalingGroupShowResponse) SetDrainTimeout(v int32)`

SetDrainTimeout sets DrainTimeout field to given value.

### HasDrainTimeout

`func (o *AutoScalingGroupShowResponse) HasDrainTimeout() bool`

HasDrainTimeout returns a boolean if a field has been set.

### SetDrainTimeoutNil

`func (o *AutoScalingGroupShowResponse) SetDrainTimeoutNil(b bool)`

 SetDrainTimeoutNil sets the value for DrainTimeout to be an explicit nil

### UnsetDrainTimeout
`func (o *AutoScalingGroupShowResponse) UnsetDrainTimeout()`

UnsetDrainTimeout ensures that no value is present for DrainTimeout, not even an explicit nil
### GetId

`func (o *AutoScalingGroupShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoScalingGroupShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoScalingGroupShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLaunchConfigurationId

`func (o *AutoScalingGroupShowResponse) GetLaunchConfigurationId() string`

GetLaunchConfigurationId returns the LaunchConfigurationId field if non-nil, zero value otherwise.

### GetLaunchConfigurationIdOk

`func (o *AutoScalingGroupShowResponse) GetLaunchConfigurationIdOk() (*string, bool)`

GetLaunchConfigurationIdOk returns a tuple with the LaunchConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfigurationId

`func (o *AutoScalingGroupShowResponse) SetLaunchConfigurationId(v string)`

SetLaunchConfigurationId sets LaunchConfigurationId field to given value.


### GetLaunchConfigurationName

`func (o *AutoScalingGroupShowResponse) GetLaunchConfigurationName() string`

GetLaunchConfigurationName returns the LaunchConfigurationName field if non-nil, zero value otherwise.

### GetLaunchConfigurationNameOk

`func (o *AutoScalingGroupShowResponse) GetLaunchConfigurationNameOk() (*string, bool)`

GetLaunchConfigurationNameOk returns a tuple with the LaunchConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfigurationName

`func (o *AutoScalingGroupShowResponse) SetLaunchConfigurationName(v string)`

SetLaunchConfigurationName sets LaunchConfigurationName field to given value.


### GetLbServerGroups

`func (o *AutoScalingGroupShowResponse) GetLbServerGroups() []AutoScalingGroupLbServerGroup`

GetLbServerGroups returns the LbServerGroups field if non-nil, zero value otherwise.

### GetLbServerGroupsOk

`func (o *AutoScalingGroupShowResponse) GetLbServerGroupsOk() (*[]AutoScalingGroupLbServerGroup, bool)`

GetLbServerGroupsOk returns a tuple with the LbServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroups

`func (o *AutoScalingGroupShowResponse) SetLbServerGroups(v []AutoScalingGroupLbServerGroup)`

SetLbServerGroups sets LbServerGroups field to given value.

### HasLbServerGroups

`func (o *AutoScalingGroupShowResponse) HasLbServerGroups() bool`

HasLbServerGroups returns a boolean if a field has been set.

### SetLbServerGroupsNil

`func (o *AutoScalingGroupShowResponse) SetLbServerGroupsNil(b bool)`

 SetLbServerGroupsNil sets the value for LbServerGroups to be an explicit nil

### UnsetLbServerGroups
`func (o *AutoScalingGroupShowResponse) UnsetLbServerGroups()`

UnsetLbServerGroups ensures that no value is present for LbServerGroups, not even an explicit nil
### GetMaxServerCount

`func (o *AutoScalingGroupShowResponse) GetMaxServerCount() int32`

GetMaxServerCount returns the MaxServerCount field if non-nil, zero value otherwise.

### GetMaxServerCountOk

`func (o *AutoScalingGroupShowResponse) GetMaxServerCountOk() (*int32, bool)`

GetMaxServerCountOk returns a tuple with the MaxServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerCount

`func (o *AutoScalingGroupShowResponse) SetMaxServerCount(v int32)`

SetMaxServerCount sets MaxServerCount field to given value.


### GetMinServerCount

`func (o *AutoScalingGroupShowResponse) GetMinServerCount() int32`

GetMinServerCount returns the MinServerCount field if non-nil, zero value otherwise.

### GetMinServerCountOk

`func (o *AutoScalingGroupShowResponse) GetMinServerCountOk() (*int32, bool)`

GetMinServerCountOk returns a tuple with the MinServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinServerCount

`func (o *AutoScalingGroupShowResponse) SetMinServerCount(v int32)`

SetMinServerCount sets MinServerCount field to given value.


### GetModifiedAt

`func (o *AutoScalingGroupShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AutoScalingGroupShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AutoScalingGroupShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *AutoScalingGroupShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *AutoScalingGroupShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *AutoScalingGroupShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *AutoScalingGroupShowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupShowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupShowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSecurityGroupIds

`func (o *AutoScalingGroupShowResponse) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *AutoScalingGroupShowResponse) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *AutoScalingGroupShowResponse) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *AutoScalingGroupShowResponse) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIdsNil

`func (o *AutoScalingGroupShowResponse) SetSecurityGroupIdsNil(b bool)`

 SetSecurityGroupIdsNil sets the value for SecurityGroupIds to be an explicit nil

### UnsetSecurityGroupIds
`func (o *AutoScalingGroupShowResponse) UnsetSecurityGroupIds()`

UnsetSecurityGroupIds ensures that no value is present for SecurityGroupIds, not even an explicit nil
### GetServerNamePrefix

`func (o *AutoScalingGroupShowResponse) GetServerNamePrefix() string`

GetServerNamePrefix returns the ServerNamePrefix field if non-nil, zero value otherwise.

### GetServerNamePrefixOk

`func (o *AutoScalingGroupShowResponse) GetServerNamePrefixOk() (*string, bool)`

GetServerNamePrefixOk returns a tuple with the ServerNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamePrefix

`func (o *AutoScalingGroupShowResponse) SetServerNamePrefix(v string)`

SetServerNamePrefix sets ServerNamePrefix field to given value.


### GetState

`func (o *AutoScalingGroupShowResponse) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AutoScalingGroupShowResponse) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AutoScalingGroupShowResponse) SetState(v string)`

SetState sets State field to given value.


### GetSubnetIds

`func (o *AutoScalingGroupShowResponse) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *AutoScalingGroupShowResponse) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *AutoScalingGroupShowResponse) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.

### HasSubnetIds

`func (o *AutoScalingGroupShowResponse) HasSubnetIds() bool`

HasSubnetIds returns a boolean if a field has been set.

### GetTotalServerCount

`func (o *AutoScalingGroupShowResponse) GetTotalServerCount() int32`

GetTotalServerCount returns the TotalServerCount field if non-nil, zero value otherwise.

### GetTotalServerCountOk

`func (o *AutoScalingGroupShowResponse) GetTotalServerCountOk() (*int32, bool)`

GetTotalServerCountOk returns a tuple with the TotalServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServerCount

`func (o *AutoScalingGroupShowResponse) SetTotalServerCount(v int32)`

SetTotalServerCount sets TotalServerCount field to given value.


### GetVpcId

`func (o *AutoScalingGroupShowResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *AutoScalingGroupShowResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *AutoScalingGroupShowResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


