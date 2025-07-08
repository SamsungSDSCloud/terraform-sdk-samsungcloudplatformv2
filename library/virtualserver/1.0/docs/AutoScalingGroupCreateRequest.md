# AutoScalingGroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredServerCount** | **int32** | Desired server count | 
**DesiredServerCountEditable** | **bool** | Desired server count editable | 
**DrainEnabled** | Pointer to **NullableBool** |  | [optional] 
**DrainTimeout** | Pointer to **NullableInt32** |  | [optional] 
**LaunchConfigurationId** | **string** | Launch Configuration ID | 
**LbServerGroups** | Pointer to [**[]AutoScalingGroupLbServerGroup**](AutoScalingGroupLbServerGroup.md) |  | [optional] 
**MaxServerCount** | **int32** | Max server count | 
**MinServerCount** | **int32** | Min server count | 
**Name** | **string** | Auto-Scaling Group name | 
**Notifications** | Pointer to [**[]AutoScalingGroupNotificationCreateRequest**](AutoScalingGroupNotificationCreateRequest.md) |  | [optional] 
**ScalingPolicies** | Pointer to [**[]AutoScalingGroupPolicyCreateRequest**](AutoScalingGroupPolicyCreateRequest.md) |  | [optional] 
**SecurityGroupIds** | Pointer to **[]string** |  | [optional] 
**ServerNamePrefix** | **string** | Server name prefix | 
**SubnetIds** | **[]string** | Subnet ID list | 
**Tags** | Pointer to [**[]Tag**](Tag.md) |  | [optional] 

## Methods

### NewAutoScalingGroupCreateRequest

`func NewAutoScalingGroupCreateRequest(desiredServerCount int32, desiredServerCountEditable bool, launchConfigurationId string, maxServerCount int32, minServerCount int32, name string, serverNamePrefix string, subnetIds []string, ) *AutoScalingGroupCreateRequest`

NewAutoScalingGroupCreateRequest instantiates a new AutoScalingGroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupCreateRequestWithDefaults

`func NewAutoScalingGroupCreateRequestWithDefaults() *AutoScalingGroupCreateRequest`

NewAutoScalingGroupCreateRequestWithDefaults instantiates a new AutoScalingGroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredServerCount

`func (o *AutoScalingGroupCreateRequest) GetDesiredServerCount() int32`

GetDesiredServerCount returns the DesiredServerCount field if non-nil, zero value otherwise.

### GetDesiredServerCountOk

`func (o *AutoScalingGroupCreateRequest) GetDesiredServerCountOk() (*int32, bool)`

GetDesiredServerCountOk returns a tuple with the DesiredServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCount

`func (o *AutoScalingGroupCreateRequest) SetDesiredServerCount(v int32)`

SetDesiredServerCount sets DesiredServerCount field to given value.


### GetDesiredServerCountEditable

`func (o *AutoScalingGroupCreateRequest) GetDesiredServerCountEditable() bool`

GetDesiredServerCountEditable returns the DesiredServerCountEditable field if non-nil, zero value otherwise.

### GetDesiredServerCountEditableOk

`func (o *AutoScalingGroupCreateRequest) GetDesiredServerCountEditableOk() (*bool, bool)`

GetDesiredServerCountEditableOk returns a tuple with the DesiredServerCountEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCountEditable

`func (o *AutoScalingGroupCreateRequest) SetDesiredServerCountEditable(v bool)`

SetDesiredServerCountEditable sets DesiredServerCountEditable field to given value.


### GetDrainEnabled

`func (o *AutoScalingGroupCreateRequest) GetDrainEnabled() bool`

GetDrainEnabled returns the DrainEnabled field if non-nil, zero value otherwise.

### GetDrainEnabledOk

`func (o *AutoScalingGroupCreateRequest) GetDrainEnabledOk() (*bool, bool)`

GetDrainEnabledOk returns a tuple with the DrainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainEnabled

`func (o *AutoScalingGroupCreateRequest) SetDrainEnabled(v bool)`

SetDrainEnabled sets DrainEnabled field to given value.

### HasDrainEnabled

`func (o *AutoScalingGroupCreateRequest) HasDrainEnabled() bool`

HasDrainEnabled returns a boolean if a field has been set.

### SetDrainEnabledNil

`func (o *AutoScalingGroupCreateRequest) SetDrainEnabledNil(b bool)`

 SetDrainEnabledNil sets the value for DrainEnabled to be an explicit nil

### UnsetDrainEnabled
`func (o *AutoScalingGroupCreateRequest) UnsetDrainEnabled()`

UnsetDrainEnabled ensures that no value is present for DrainEnabled, not even an explicit nil
### GetDrainTimeout

`func (o *AutoScalingGroupCreateRequest) GetDrainTimeout() int32`

GetDrainTimeout returns the DrainTimeout field if non-nil, zero value otherwise.

### GetDrainTimeoutOk

`func (o *AutoScalingGroupCreateRequest) GetDrainTimeoutOk() (*int32, bool)`

GetDrainTimeoutOk returns a tuple with the DrainTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainTimeout

`func (o *AutoScalingGroupCreateRequest) SetDrainTimeout(v int32)`

SetDrainTimeout sets DrainTimeout field to given value.

### HasDrainTimeout

`func (o *AutoScalingGroupCreateRequest) HasDrainTimeout() bool`

HasDrainTimeout returns a boolean if a field has been set.

### SetDrainTimeoutNil

`func (o *AutoScalingGroupCreateRequest) SetDrainTimeoutNil(b bool)`

 SetDrainTimeoutNil sets the value for DrainTimeout to be an explicit nil

### UnsetDrainTimeout
`func (o *AutoScalingGroupCreateRequest) UnsetDrainTimeout()`

UnsetDrainTimeout ensures that no value is present for DrainTimeout, not even an explicit nil
### GetLaunchConfigurationId

`func (o *AutoScalingGroupCreateRequest) GetLaunchConfigurationId() string`

GetLaunchConfigurationId returns the LaunchConfigurationId field if non-nil, zero value otherwise.

### GetLaunchConfigurationIdOk

`func (o *AutoScalingGroupCreateRequest) GetLaunchConfigurationIdOk() (*string, bool)`

GetLaunchConfigurationIdOk returns a tuple with the LaunchConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfigurationId

`func (o *AutoScalingGroupCreateRequest) SetLaunchConfigurationId(v string)`

SetLaunchConfigurationId sets LaunchConfigurationId field to given value.


### GetLbServerGroups

`func (o *AutoScalingGroupCreateRequest) GetLbServerGroups() []AutoScalingGroupLbServerGroup`

GetLbServerGroups returns the LbServerGroups field if non-nil, zero value otherwise.

### GetLbServerGroupsOk

`func (o *AutoScalingGroupCreateRequest) GetLbServerGroupsOk() (*[]AutoScalingGroupLbServerGroup, bool)`

GetLbServerGroupsOk returns a tuple with the LbServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroups

`func (o *AutoScalingGroupCreateRequest) SetLbServerGroups(v []AutoScalingGroupLbServerGroup)`

SetLbServerGroups sets LbServerGroups field to given value.

### HasLbServerGroups

`func (o *AutoScalingGroupCreateRequest) HasLbServerGroups() bool`

HasLbServerGroups returns a boolean if a field has been set.

### SetLbServerGroupsNil

`func (o *AutoScalingGroupCreateRequest) SetLbServerGroupsNil(b bool)`

 SetLbServerGroupsNil sets the value for LbServerGroups to be an explicit nil

### UnsetLbServerGroups
`func (o *AutoScalingGroupCreateRequest) UnsetLbServerGroups()`

UnsetLbServerGroups ensures that no value is present for LbServerGroups, not even an explicit nil
### GetMaxServerCount

`func (o *AutoScalingGroupCreateRequest) GetMaxServerCount() int32`

GetMaxServerCount returns the MaxServerCount field if non-nil, zero value otherwise.

### GetMaxServerCountOk

`func (o *AutoScalingGroupCreateRequest) GetMaxServerCountOk() (*int32, bool)`

GetMaxServerCountOk returns a tuple with the MaxServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxServerCount

`func (o *AutoScalingGroupCreateRequest) SetMaxServerCount(v int32)`

SetMaxServerCount sets MaxServerCount field to given value.


### GetMinServerCount

`func (o *AutoScalingGroupCreateRequest) GetMinServerCount() int32`

GetMinServerCount returns the MinServerCount field if non-nil, zero value otherwise.

### GetMinServerCountOk

`func (o *AutoScalingGroupCreateRequest) GetMinServerCountOk() (*int32, bool)`

GetMinServerCountOk returns a tuple with the MinServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinServerCount

`func (o *AutoScalingGroupCreateRequest) SetMinServerCount(v int32)`

SetMinServerCount sets MinServerCount field to given value.


### GetName

`func (o *AutoScalingGroupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoScalingGroupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoScalingGroupCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNotifications

`func (o *AutoScalingGroupCreateRequest) GetNotifications() []AutoScalingGroupNotificationCreateRequest`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *AutoScalingGroupCreateRequest) GetNotificationsOk() (*[]AutoScalingGroupNotificationCreateRequest, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *AutoScalingGroupCreateRequest) SetNotifications(v []AutoScalingGroupNotificationCreateRequest)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *AutoScalingGroupCreateRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### SetNotificationsNil

`func (o *AutoScalingGroupCreateRequest) SetNotificationsNil(b bool)`

 SetNotificationsNil sets the value for Notifications to be an explicit nil

### UnsetNotifications
`func (o *AutoScalingGroupCreateRequest) UnsetNotifications()`

UnsetNotifications ensures that no value is present for Notifications, not even an explicit nil
### GetScalingPolicies

`func (o *AutoScalingGroupCreateRequest) GetScalingPolicies() []AutoScalingGroupPolicyCreateRequest`

GetScalingPolicies returns the ScalingPolicies field if non-nil, zero value otherwise.

### GetScalingPoliciesOk

`func (o *AutoScalingGroupCreateRequest) GetScalingPoliciesOk() (*[]AutoScalingGroupPolicyCreateRequest, bool)`

GetScalingPoliciesOk returns a tuple with the ScalingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingPolicies

`func (o *AutoScalingGroupCreateRequest) SetScalingPolicies(v []AutoScalingGroupPolicyCreateRequest)`

SetScalingPolicies sets ScalingPolicies field to given value.

### HasScalingPolicies

`func (o *AutoScalingGroupCreateRequest) HasScalingPolicies() bool`

HasScalingPolicies returns a boolean if a field has been set.

### SetScalingPoliciesNil

`func (o *AutoScalingGroupCreateRequest) SetScalingPoliciesNil(b bool)`

 SetScalingPoliciesNil sets the value for ScalingPolicies to be an explicit nil

### UnsetScalingPolicies
`func (o *AutoScalingGroupCreateRequest) UnsetScalingPolicies()`

UnsetScalingPolicies ensures that no value is present for ScalingPolicies, not even an explicit nil
### GetSecurityGroupIds

`func (o *AutoScalingGroupCreateRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *AutoScalingGroupCreateRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *AutoScalingGroupCreateRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *AutoScalingGroupCreateRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIdsNil

`func (o *AutoScalingGroupCreateRequest) SetSecurityGroupIdsNil(b bool)`

 SetSecurityGroupIdsNil sets the value for SecurityGroupIds to be an explicit nil

### UnsetSecurityGroupIds
`func (o *AutoScalingGroupCreateRequest) UnsetSecurityGroupIds()`

UnsetSecurityGroupIds ensures that no value is present for SecurityGroupIds, not even an explicit nil
### GetServerNamePrefix

`func (o *AutoScalingGroupCreateRequest) GetServerNamePrefix() string`

GetServerNamePrefix returns the ServerNamePrefix field if non-nil, zero value otherwise.

### GetServerNamePrefixOk

`func (o *AutoScalingGroupCreateRequest) GetServerNamePrefixOk() (*string, bool)`

GetServerNamePrefixOk returns a tuple with the ServerNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamePrefix

`func (o *AutoScalingGroupCreateRequest) SetServerNamePrefix(v string)`

SetServerNamePrefix sets ServerNamePrefix field to given value.


### GetSubnetIds

`func (o *AutoScalingGroupCreateRequest) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *AutoScalingGroupCreateRequest) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *AutoScalingGroupCreateRequest) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.


### GetTags

`func (o *AutoScalingGroupCreateRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AutoScalingGroupCreateRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AutoScalingGroupCreateRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AutoScalingGroupCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *AutoScalingGroupCreateRequest) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *AutoScalingGroupCreateRequest) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


