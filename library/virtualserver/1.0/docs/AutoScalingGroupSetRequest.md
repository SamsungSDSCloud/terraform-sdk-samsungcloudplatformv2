# AutoScalingGroupSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredServerCountEditable** | Pointer to **NullableBool** |  | [optional] 
**DrainEnabled** | Pointer to **NullableBool** |  | [optional] 
**DrainTimeout** | Pointer to **NullableInt32** |  | [optional] 
**LaunchConfigurationId** | Pointer to **NullableString** |  | [optional] 
**SecurityGroupIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAutoScalingGroupSetRequest

`func NewAutoScalingGroupSetRequest() *AutoScalingGroupSetRequest`

NewAutoScalingGroupSetRequest instantiates a new AutoScalingGroupSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoScalingGroupSetRequestWithDefaults

`func NewAutoScalingGroupSetRequestWithDefaults() *AutoScalingGroupSetRequest`

NewAutoScalingGroupSetRequestWithDefaults instantiates a new AutoScalingGroupSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredServerCountEditable

`func (o *AutoScalingGroupSetRequest) GetDesiredServerCountEditable() bool`

GetDesiredServerCountEditable returns the DesiredServerCountEditable field if non-nil, zero value otherwise.

### GetDesiredServerCountEditableOk

`func (o *AutoScalingGroupSetRequest) GetDesiredServerCountEditableOk() (*bool, bool)`

GetDesiredServerCountEditableOk returns a tuple with the DesiredServerCountEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredServerCountEditable

`func (o *AutoScalingGroupSetRequest) SetDesiredServerCountEditable(v bool)`

SetDesiredServerCountEditable sets DesiredServerCountEditable field to given value.

### HasDesiredServerCountEditable

`func (o *AutoScalingGroupSetRequest) HasDesiredServerCountEditable() bool`

HasDesiredServerCountEditable returns a boolean if a field has been set.

### SetDesiredServerCountEditableNil

`func (o *AutoScalingGroupSetRequest) SetDesiredServerCountEditableNil(b bool)`

 SetDesiredServerCountEditableNil sets the value for DesiredServerCountEditable to be an explicit nil

### UnsetDesiredServerCountEditable
`func (o *AutoScalingGroupSetRequest) UnsetDesiredServerCountEditable()`

UnsetDesiredServerCountEditable ensures that no value is present for DesiredServerCountEditable, not even an explicit nil
### GetDrainEnabled

`func (o *AutoScalingGroupSetRequest) GetDrainEnabled() bool`

GetDrainEnabled returns the DrainEnabled field if non-nil, zero value otherwise.

### GetDrainEnabledOk

`func (o *AutoScalingGroupSetRequest) GetDrainEnabledOk() (*bool, bool)`

GetDrainEnabledOk returns a tuple with the DrainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainEnabled

`func (o *AutoScalingGroupSetRequest) SetDrainEnabled(v bool)`

SetDrainEnabled sets DrainEnabled field to given value.

### HasDrainEnabled

`func (o *AutoScalingGroupSetRequest) HasDrainEnabled() bool`

HasDrainEnabled returns a boolean if a field has been set.

### SetDrainEnabledNil

`func (o *AutoScalingGroupSetRequest) SetDrainEnabledNil(b bool)`

 SetDrainEnabledNil sets the value for DrainEnabled to be an explicit nil

### UnsetDrainEnabled
`func (o *AutoScalingGroupSetRequest) UnsetDrainEnabled()`

UnsetDrainEnabled ensures that no value is present for DrainEnabled, not even an explicit nil
### GetDrainTimeout

`func (o *AutoScalingGroupSetRequest) GetDrainTimeout() int32`

GetDrainTimeout returns the DrainTimeout field if non-nil, zero value otherwise.

### GetDrainTimeoutOk

`func (o *AutoScalingGroupSetRequest) GetDrainTimeoutOk() (*int32, bool)`

GetDrainTimeoutOk returns a tuple with the DrainTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainTimeout

`func (o *AutoScalingGroupSetRequest) SetDrainTimeout(v int32)`

SetDrainTimeout sets DrainTimeout field to given value.

### HasDrainTimeout

`func (o *AutoScalingGroupSetRequest) HasDrainTimeout() bool`

HasDrainTimeout returns a boolean if a field has been set.

### SetDrainTimeoutNil

`func (o *AutoScalingGroupSetRequest) SetDrainTimeoutNil(b bool)`

 SetDrainTimeoutNil sets the value for DrainTimeout to be an explicit nil

### UnsetDrainTimeout
`func (o *AutoScalingGroupSetRequest) UnsetDrainTimeout()`

UnsetDrainTimeout ensures that no value is present for DrainTimeout, not even an explicit nil
### GetLaunchConfigurationId

`func (o *AutoScalingGroupSetRequest) GetLaunchConfigurationId() string`

GetLaunchConfigurationId returns the LaunchConfigurationId field if non-nil, zero value otherwise.

### GetLaunchConfigurationIdOk

`func (o *AutoScalingGroupSetRequest) GetLaunchConfigurationIdOk() (*string, bool)`

GetLaunchConfigurationIdOk returns a tuple with the LaunchConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLaunchConfigurationId

`func (o *AutoScalingGroupSetRequest) SetLaunchConfigurationId(v string)`

SetLaunchConfigurationId sets LaunchConfigurationId field to given value.

### HasLaunchConfigurationId

`func (o *AutoScalingGroupSetRequest) HasLaunchConfigurationId() bool`

HasLaunchConfigurationId returns a boolean if a field has been set.

### SetLaunchConfigurationIdNil

`func (o *AutoScalingGroupSetRequest) SetLaunchConfigurationIdNil(b bool)`

 SetLaunchConfigurationIdNil sets the value for LaunchConfigurationId to be an explicit nil

### UnsetLaunchConfigurationId
`func (o *AutoScalingGroupSetRequest) UnsetLaunchConfigurationId()`

UnsetLaunchConfigurationId ensures that no value is present for LaunchConfigurationId, not even an explicit nil
### GetSecurityGroupIds

`func (o *AutoScalingGroupSetRequest) GetSecurityGroupIds() []string`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *AutoScalingGroupSetRequest) GetSecurityGroupIdsOk() (*[]string, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *AutoScalingGroupSetRequest) SetSecurityGroupIds(v []string)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *AutoScalingGroupSetRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.

### SetSecurityGroupIdsNil

`func (o *AutoScalingGroupSetRequest) SetSecurityGroupIdsNil(b bool)`

 SetSecurityGroupIdsNil sets the value for SecurityGroupIds to be an explicit nil

### UnsetSecurityGroupIds
`func (o *AutoScalingGroupSetRequest) UnsetSecurityGroupIds()`

UnsetSecurityGroupIds ensures that no value is present for SecurityGroupIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


