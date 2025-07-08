# PodSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **NullableString** |  | 
**ClusterId** | **string** | Cluster ID | 
**ContainerStatus** | **NullableString** |  | 
**CreatedAt** | **NullableTime** |  | 
**Name** | **string** | Pod Name | 
**NamespaceName** | **string** | Namespace Name | 
**NodeName** | **NullableString** |  | 
**PodIp** | **NullableString** |  | 
**PodStatus** | **NullableString** |  | 
**Restarts** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewPodSummary

`func NewPodSummary(age NullableString, clusterId string, containerStatus NullableString, createdAt NullableTime, name string, namespaceName string, nodeName NullableString, podIp NullableString, podStatus NullableString, restarts NullableString, uid string, ) *PodSummary`

NewPodSummary instantiates a new PodSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodSummaryWithDefaults

`func NewPodSummaryWithDefaults() *PodSummary`

NewPodSummaryWithDefaults instantiates a new PodSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *PodSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PodSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PodSummary) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *PodSummary) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *PodSummary) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetClusterId

`func (o *PodSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PodSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PodSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetContainerStatus

`func (o *PodSummary) GetContainerStatus() string`

GetContainerStatus returns the ContainerStatus field if non-nil, zero value otherwise.

### GetContainerStatusOk

`func (o *PodSummary) GetContainerStatusOk() (*string, bool)`

GetContainerStatusOk returns a tuple with the ContainerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStatus

`func (o *PodSummary) SetContainerStatus(v string)`

SetContainerStatus sets ContainerStatus field to given value.


### SetContainerStatusNil

`func (o *PodSummary) SetContainerStatusNil(b bool)`

 SetContainerStatusNil sets the value for ContainerStatus to be an explicit nil

### UnsetContainerStatus
`func (o *PodSummary) UnsetContainerStatus()`

UnsetContainerStatus ensures that no value is present for ContainerStatus, not even an explicit nil
### GetCreatedAt

`func (o *PodSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PodSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PodSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *PodSummary) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PodSummary) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetName

`func (o *PodSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *PodSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *PodSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *PodSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetNodeName

`func (o *PodSummary) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *PodSummary) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *PodSummary) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### SetNodeNameNil

`func (o *PodSummary) SetNodeNameNil(b bool)`

 SetNodeNameNil sets the value for NodeName to be an explicit nil

### UnsetNodeName
`func (o *PodSummary) UnsetNodeName()`

UnsetNodeName ensures that no value is present for NodeName, not even an explicit nil
### GetPodIp

`func (o *PodSummary) GetPodIp() string`

GetPodIp returns the PodIp field if non-nil, zero value otherwise.

### GetPodIpOk

`func (o *PodSummary) GetPodIpOk() (*string, bool)`

GetPodIpOk returns a tuple with the PodIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIp

`func (o *PodSummary) SetPodIp(v string)`

SetPodIp sets PodIp field to given value.


### SetPodIpNil

`func (o *PodSummary) SetPodIpNil(b bool)`

 SetPodIpNil sets the value for PodIp to be an explicit nil

### UnsetPodIp
`func (o *PodSummary) UnsetPodIp()`

UnsetPodIp ensures that no value is present for PodIp, not even an explicit nil
### GetPodStatus

`func (o *PodSummary) GetPodStatus() string`

GetPodStatus returns the PodStatus field if non-nil, zero value otherwise.

### GetPodStatusOk

`func (o *PodSummary) GetPodStatusOk() (*string, bool)`

GetPodStatusOk returns a tuple with the PodStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodStatus

`func (o *PodSummary) SetPodStatus(v string)`

SetPodStatus sets PodStatus field to given value.


### SetPodStatusNil

`func (o *PodSummary) SetPodStatusNil(b bool)`

 SetPodStatusNil sets the value for PodStatus to be an explicit nil

### UnsetPodStatus
`func (o *PodSummary) UnsetPodStatus()`

UnsetPodStatus ensures that no value is present for PodStatus, not even an explicit nil
### GetRestarts

`func (o *PodSummary) GetRestarts() string`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *PodSummary) GetRestartsOk() (*string, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *PodSummary) SetRestarts(v string)`

SetRestarts sets Restarts field to given value.


### SetRestartsNil

`func (o *PodSummary) SetRestartsNil(b bool)`

 SetRestartsNil sets the value for Restarts to be an explicit nil

### UnsetRestarts
`func (o *PodSummary) UnsetRestarts()`

UnsetRestarts ensures that no value is present for Restarts, not even an explicit nil
### GetUid

`func (o *PodSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PodSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PodSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


