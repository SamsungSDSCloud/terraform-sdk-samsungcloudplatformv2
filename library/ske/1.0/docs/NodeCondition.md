# NodeCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastHeartbeatTime** | **NullableTime** |  | 
**LastTransitionTime** | **NullableTime** |  | 
**NodeConditionMessage** | **NullableString** |  | 
**NodeConditionReason** | **NullableString** |  | 
**NodeConditionStatus** | **NullableString** |  | 
**NodeConditionType** | **NullableString** |  | 

## Methods

### NewNodeCondition

`func NewNodeCondition(lastHeartbeatTime NullableTime, lastTransitionTime NullableTime, nodeConditionMessage NullableString, nodeConditionReason NullableString, nodeConditionStatus NullableString, nodeConditionType NullableString, ) *NodeCondition`

NewNodeCondition instantiates a new NodeCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeConditionWithDefaults

`func NewNodeConditionWithDefaults() *NodeCondition`

NewNodeConditionWithDefaults instantiates a new NodeCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastHeartbeatTime

`func (o *NodeCondition) GetLastHeartbeatTime() time.Time`

GetLastHeartbeatTime returns the LastHeartbeatTime field if non-nil, zero value otherwise.

### GetLastHeartbeatTimeOk

`func (o *NodeCondition) GetLastHeartbeatTimeOk() (*time.Time, bool)`

GetLastHeartbeatTimeOk returns a tuple with the LastHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastHeartbeatTime

`func (o *NodeCondition) SetLastHeartbeatTime(v time.Time)`

SetLastHeartbeatTime sets LastHeartbeatTime field to given value.


### SetLastHeartbeatTimeNil

`func (o *NodeCondition) SetLastHeartbeatTimeNil(b bool)`

 SetLastHeartbeatTimeNil sets the value for LastHeartbeatTime to be an explicit nil

### UnsetLastHeartbeatTime
`func (o *NodeCondition) UnsetLastHeartbeatTime()`

UnsetLastHeartbeatTime ensures that no value is present for LastHeartbeatTime, not even an explicit nil
### GetLastTransitionTime

`func (o *NodeCondition) GetLastTransitionTime() time.Time`

GetLastTransitionTime returns the LastTransitionTime field if non-nil, zero value otherwise.

### GetLastTransitionTimeOk

`func (o *NodeCondition) GetLastTransitionTimeOk() (*time.Time, bool)`

GetLastTransitionTimeOk returns a tuple with the LastTransitionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransitionTime

`func (o *NodeCondition) SetLastTransitionTime(v time.Time)`

SetLastTransitionTime sets LastTransitionTime field to given value.


### SetLastTransitionTimeNil

`func (o *NodeCondition) SetLastTransitionTimeNil(b bool)`

 SetLastTransitionTimeNil sets the value for LastTransitionTime to be an explicit nil

### UnsetLastTransitionTime
`func (o *NodeCondition) UnsetLastTransitionTime()`

UnsetLastTransitionTime ensures that no value is present for LastTransitionTime, not even an explicit nil
### GetNodeConditionMessage

`func (o *NodeCondition) GetNodeConditionMessage() string`

GetNodeConditionMessage returns the NodeConditionMessage field if non-nil, zero value otherwise.

### GetNodeConditionMessageOk

`func (o *NodeCondition) GetNodeConditionMessageOk() (*string, bool)`

GetNodeConditionMessageOk returns a tuple with the NodeConditionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConditionMessage

`func (o *NodeCondition) SetNodeConditionMessage(v string)`

SetNodeConditionMessage sets NodeConditionMessage field to given value.


### SetNodeConditionMessageNil

`func (o *NodeCondition) SetNodeConditionMessageNil(b bool)`

 SetNodeConditionMessageNil sets the value for NodeConditionMessage to be an explicit nil

### UnsetNodeConditionMessage
`func (o *NodeCondition) UnsetNodeConditionMessage()`

UnsetNodeConditionMessage ensures that no value is present for NodeConditionMessage, not even an explicit nil
### GetNodeConditionReason

`func (o *NodeCondition) GetNodeConditionReason() string`

GetNodeConditionReason returns the NodeConditionReason field if non-nil, zero value otherwise.

### GetNodeConditionReasonOk

`func (o *NodeCondition) GetNodeConditionReasonOk() (*string, bool)`

GetNodeConditionReasonOk returns a tuple with the NodeConditionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConditionReason

`func (o *NodeCondition) SetNodeConditionReason(v string)`

SetNodeConditionReason sets NodeConditionReason field to given value.


### SetNodeConditionReasonNil

`func (o *NodeCondition) SetNodeConditionReasonNil(b bool)`

 SetNodeConditionReasonNil sets the value for NodeConditionReason to be an explicit nil

### UnsetNodeConditionReason
`func (o *NodeCondition) UnsetNodeConditionReason()`

UnsetNodeConditionReason ensures that no value is present for NodeConditionReason, not even an explicit nil
### GetNodeConditionStatus

`func (o *NodeCondition) GetNodeConditionStatus() string`

GetNodeConditionStatus returns the NodeConditionStatus field if non-nil, zero value otherwise.

### GetNodeConditionStatusOk

`func (o *NodeCondition) GetNodeConditionStatusOk() (*string, bool)`

GetNodeConditionStatusOk returns a tuple with the NodeConditionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConditionStatus

`func (o *NodeCondition) SetNodeConditionStatus(v string)`

SetNodeConditionStatus sets NodeConditionStatus field to given value.


### SetNodeConditionStatusNil

`func (o *NodeCondition) SetNodeConditionStatusNil(b bool)`

 SetNodeConditionStatusNil sets the value for NodeConditionStatus to be an explicit nil

### UnsetNodeConditionStatus
`func (o *NodeCondition) UnsetNodeConditionStatus()`

UnsetNodeConditionStatus ensures that no value is present for NodeConditionStatus, not even an explicit nil
### GetNodeConditionType

`func (o *NodeCondition) GetNodeConditionType() string`

GetNodeConditionType returns the NodeConditionType field if non-nil, zero value otherwise.

### GetNodeConditionTypeOk

`func (o *NodeCondition) GetNodeConditionTypeOk() (*string, bool)`

GetNodeConditionTypeOk returns a tuple with the NodeConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConditionType

`func (o *NodeCondition) SetNodeConditionType(v string)`

SetNodeConditionType sets NodeConditionType field to given value.


### SetNodeConditionTypeNil

`func (o *NodeCondition) SetNodeConditionTypeNil(b bool)`

 SetNodeConditionTypeNil sets the value for NodeConditionType to be an explicit nil

### UnsetNodeConditionType
`func (o *NodeCondition) UnsetNodeConditionType()`

UnsetNodeConditionType ensures that no value is present for NodeConditionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


