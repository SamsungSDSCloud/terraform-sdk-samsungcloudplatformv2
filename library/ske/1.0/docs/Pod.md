# Pod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **NullableString** |  | 
**Annotations** | **[]string** |  | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**ContainerList** | [**[]PodContainerResponse**](PodContainerResponse.md) |  | 
**ContainerStatus** | **NullableString** |  | 
**CreatedAt** | **NullableTime** |  | 
**InitContainerList** | [**[]PodContainerResponse**](PodContainerResponse.md) |  | 
**Labels** | **[]string** |  | 
**Name** | **string** | Pod Name | 
**NamespaceName** | **string** | Namespace Name | 
**NodeIp** | **NullableString** |  | 
**NodeName** | **NullableString** |  | 
**NodeSelector** | **[]string** |  | 
**PodIp** | **NullableString** |  | 
**PodStatus** | **NullableString** |  | 
**Priority** | **NullableInt32** |  | 
**PriorityClassName** | **NullableString** |  | 
**QosClass** | **NullableString** |  | 
**Restarts** | **NullableString** |  | 
**StartTime** | **NullableTime** |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewPod

`func NewPod(age NullableString, annotations []string, clusterId string, clusterName string, containerList []PodContainerResponse, containerStatus NullableString, createdAt NullableTime, initContainerList []PodContainerResponse, labels []string, name string, namespaceName string, nodeIp NullableString, nodeName NullableString, nodeSelector []string, podIp NullableString, podStatus NullableString, priority NullableInt32, priorityClassName NullableString, qosClass NullableString, restarts NullableString, startTime NullableTime, uid string, yaml string, ) *Pod`

NewPod instantiates a new Pod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodWithDefaults

`func NewPodWithDefaults() *Pod`

NewPodWithDefaults instantiates a new Pod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Pod) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Pod) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Pod) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *Pod) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *Pod) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAnnotations

`func (o *Pod) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Pod) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Pod) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### SetAnnotationsNil

`func (o *Pod) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *Pod) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetClusterId

`func (o *Pod) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Pod) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Pod) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Pod) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Pod) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Pod) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetContainerList

`func (o *Pod) GetContainerList() []PodContainerResponse`

GetContainerList returns the ContainerList field if non-nil, zero value otherwise.

### GetContainerListOk

`func (o *Pod) GetContainerListOk() (*[]PodContainerResponse, bool)`

GetContainerListOk returns a tuple with the ContainerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerList

`func (o *Pod) SetContainerList(v []PodContainerResponse)`

SetContainerList sets ContainerList field to given value.


### SetContainerListNil

`func (o *Pod) SetContainerListNil(b bool)`

 SetContainerListNil sets the value for ContainerList to be an explicit nil

### UnsetContainerList
`func (o *Pod) UnsetContainerList()`

UnsetContainerList ensures that no value is present for ContainerList, not even an explicit nil
### GetContainerStatus

`func (o *Pod) GetContainerStatus() string`

GetContainerStatus returns the ContainerStatus field if non-nil, zero value otherwise.

### GetContainerStatusOk

`func (o *Pod) GetContainerStatusOk() (*string, bool)`

GetContainerStatusOk returns a tuple with the ContainerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerStatus

`func (o *Pod) SetContainerStatus(v string)`

SetContainerStatus sets ContainerStatus field to given value.


### SetContainerStatusNil

`func (o *Pod) SetContainerStatusNil(b bool)`

 SetContainerStatusNil sets the value for ContainerStatus to be an explicit nil

### UnsetContainerStatus
`func (o *Pod) UnsetContainerStatus()`

UnsetContainerStatus ensures that no value is present for ContainerStatus, not even an explicit nil
### GetCreatedAt

`func (o *Pod) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pod) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pod) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Pod) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Pod) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetInitContainerList

`func (o *Pod) GetInitContainerList() []PodContainerResponse`

GetInitContainerList returns the InitContainerList field if non-nil, zero value otherwise.

### GetInitContainerListOk

`func (o *Pod) GetInitContainerListOk() (*[]PodContainerResponse, bool)`

GetInitContainerListOk returns a tuple with the InitContainerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitContainerList

`func (o *Pod) SetInitContainerList(v []PodContainerResponse)`

SetInitContainerList sets InitContainerList field to given value.


### SetInitContainerListNil

`func (o *Pod) SetInitContainerListNil(b bool)`

 SetInitContainerListNil sets the value for InitContainerList to be an explicit nil

### UnsetInitContainerList
`func (o *Pod) UnsetInitContainerList()`

UnsetInitContainerList ensures that no value is present for InitContainerList, not even an explicit nil
### GetLabels

`func (o *Pod) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Pod) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Pod) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### SetLabelsNil

`func (o *Pod) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Pod) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetName

`func (o *Pod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pod) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *Pod) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Pod) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Pod) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetNodeIp

`func (o *Pod) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *Pod) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *Pod) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.


### SetNodeIpNil

`func (o *Pod) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *Pod) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetNodeName

`func (o *Pod) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *Pod) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *Pod) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.


### SetNodeNameNil

`func (o *Pod) SetNodeNameNil(b bool)`

 SetNodeNameNil sets the value for NodeName to be an explicit nil

### UnsetNodeName
`func (o *Pod) UnsetNodeName()`

UnsetNodeName ensures that no value is present for NodeName, not even an explicit nil
### GetNodeSelector

`func (o *Pod) GetNodeSelector() []string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *Pod) GetNodeSelectorOk() (*[]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *Pod) SetNodeSelector(v []string)`

SetNodeSelector sets NodeSelector field to given value.


### SetNodeSelectorNil

`func (o *Pod) SetNodeSelectorNil(b bool)`

 SetNodeSelectorNil sets the value for NodeSelector to be an explicit nil

### UnsetNodeSelector
`func (o *Pod) UnsetNodeSelector()`

UnsetNodeSelector ensures that no value is present for NodeSelector, not even an explicit nil
### GetPodIp

`func (o *Pod) GetPodIp() string`

GetPodIp returns the PodIp field if non-nil, zero value otherwise.

### GetPodIpOk

`func (o *Pod) GetPodIpOk() (*string, bool)`

GetPodIpOk returns a tuple with the PodIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodIp

`func (o *Pod) SetPodIp(v string)`

SetPodIp sets PodIp field to given value.


### SetPodIpNil

`func (o *Pod) SetPodIpNil(b bool)`

 SetPodIpNil sets the value for PodIp to be an explicit nil

### UnsetPodIp
`func (o *Pod) UnsetPodIp()`

UnsetPodIp ensures that no value is present for PodIp, not even an explicit nil
### GetPodStatus

`func (o *Pod) GetPodStatus() string`

GetPodStatus returns the PodStatus field if non-nil, zero value otherwise.

### GetPodStatusOk

`func (o *Pod) GetPodStatusOk() (*string, bool)`

GetPodStatusOk returns a tuple with the PodStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodStatus

`func (o *Pod) SetPodStatus(v string)`

SetPodStatus sets PodStatus field to given value.


### SetPodStatusNil

`func (o *Pod) SetPodStatusNil(b bool)`

 SetPodStatusNil sets the value for PodStatus to be an explicit nil

### UnsetPodStatus
`func (o *Pod) UnsetPodStatus()`

UnsetPodStatus ensures that no value is present for PodStatus, not even an explicit nil
### GetPriority

`func (o *Pod) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Pod) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Pod) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### SetPriorityNil

`func (o *Pod) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *Pod) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetPriorityClassName

`func (o *Pod) GetPriorityClassName() string`

GetPriorityClassName returns the PriorityClassName field if non-nil, zero value otherwise.

### GetPriorityClassNameOk

`func (o *Pod) GetPriorityClassNameOk() (*string, bool)`

GetPriorityClassNameOk returns a tuple with the PriorityClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClassName

`func (o *Pod) SetPriorityClassName(v string)`

SetPriorityClassName sets PriorityClassName field to given value.


### SetPriorityClassNameNil

`func (o *Pod) SetPriorityClassNameNil(b bool)`

 SetPriorityClassNameNil sets the value for PriorityClassName to be an explicit nil

### UnsetPriorityClassName
`func (o *Pod) UnsetPriorityClassName()`

UnsetPriorityClassName ensures that no value is present for PriorityClassName, not even an explicit nil
### GetQosClass

`func (o *Pod) GetQosClass() string`

GetQosClass returns the QosClass field if non-nil, zero value otherwise.

### GetQosClassOk

`func (o *Pod) GetQosClassOk() (*string, bool)`

GetQosClassOk returns a tuple with the QosClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQosClass

`func (o *Pod) SetQosClass(v string)`

SetQosClass sets QosClass field to given value.


### SetQosClassNil

`func (o *Pod) SetQosClassNil(b bool)`

 SetQosClassNil sets the value for QosClass to be an explicit nil

### UnsetQosClass
`func (o *Pod) UnsetQosClass()`

UnsetQosClass ensures that no value is present for QosClass, not even an explicit nil
### GetRestarts

`func (o *Pod) GetRestarts() string`

GetRestarts returns the Restarts field if non-nil, zero value otherwise.

### GetRestartsOk

`func (o *Pod) GetRestartsOk() (*string, bool)`

GetRestartsOk returns a tuple with the Restarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestarts

`func (o *Pod) SetRestarts(v string)`

SetRestarts sets Restarts field to given value.


### SetRestartsNil

`func (o *Pod) SetRestartsNil(b bool)`

 SetRestartsNil sets the value for Restarts to be an explicit nil

### UnsetRestarts
`func (o *Pod) UnsetRestarts()`

UnsetRestarts ensures that no value is present for Restarts, not even an explicit nil
### GetStartTime

`func (o *Pod) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Pod) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Pod) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### SetStartTimeNil

`func (o *Pod) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *Pod) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetUid

`func (o *Pod) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Pod) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Pod) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Pod) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Pod) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Pod) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


