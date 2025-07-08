# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**Architecture** | **NullableString** |  | 
**BootId** | **NullableString** |  | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**ContainerVersion** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**ExternalIpAddress** | **NullableString** |  | 
**HostName** | **NullableString** |  | 
**Ip** | **NullableString** |  | 
**KernelVersion** | **NullableString** |  | 
**KubeProxyVersion** | **NullableString** |  | 
**KubeletVersion** | **NullableString** |  | 
**KubernetesVersion** | **NullableString** |  | 
**Labels** | **[]string** | Labels | 
**MachineId** | **NullableString** |  | 
**Name** | **string** | Node Name | 
**NodeCapacities** | [**[]NodeCapacity**](NodeCapacity.md) | Node Capacities | 
**NodeConditions** | [**[]NodeCondition**](NodeCondition.md) | Node Conditions | 
**NodeStatus** | **string** | Node Status | 
**OperatingSystem** | **NullableString** |  | 
**OsImage** | **NullableString** |  | 
**PodCidr** | **NullableString** |  | 
**SystemUuid** | **NullableString** |  | 
**Taints** | **[]string** | Taints | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewNode

`func NewNode(age string, annotations []string, architecture NullableString, bootId NullableString, clusterId string, clusterName string, containerVersion NullableString, createdAt time.Time, createdBy string, externalIpAddress NullableString, hostName NullableString, ip NullableString, kernelVersion NullableString, kubeProxyVersion NullableString, kubeletVersion NullableString, kubernetesVersion NullableString, labels []string, machineId NullableString, name string, nodeCapacities []NodeCapacity, nodeConditions []NodeCondition, nodeStatus string, operatingSystem NullableString, osImage NullableString, podCidr NullableString, systemUuid NullableString, taints []string, uid string, yaml string, ) *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Node) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Node) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Node) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *Node) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Node) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Node) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetArchitecture

`func (o *Node) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *Node) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *Node) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### SetArchitectureNil

`func (o *Node) SetArchitectureNil(b bool)`

 SetArchitectureNil sets the value for Architecture to be an explicit nil

### UnsetArchitecture
`func (o *Node) UnsetArchitecture()`

UnsetArchitecture ensures that no value is present for Architecture, not even an explicit nil
### GetBootId

`func (o *Node) GetBootId() string`

GetBootId returns the BootId field if non-nil, zero value otherwise.

### GetBootIdOk

`func (o *Node) GetBootIdOk() (*string, bool)`

GetBootIdOk returns a tuple with the BootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootId

`func (o *Node) SetBootId(v string)`

SetBootId sets BootId field to given value.


### SetBootIdNil

`func (o *Node) SetBootIdNil(b bool)`

 SetBootIdNil sets the value for BootId to be an explicit nil

### UnsetBootId
`func (o *Node) UnsetBootId()`

UnsetBootId ensures that no value is present for BootId, not even an explicit nil
### GetClusterId

`func (o *Node) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Node) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Node) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Node) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Node) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Node) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetContainerVersion

`func (o *Node) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *Node) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *Node) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.


### SetContainerVersionNil

`func (o *Node) SetContainerVersionNil(b bool)`

 SetContainerVersionNil sets the value for ContainerVersion to be an explicit nil

### UnsetContainerVersion
`func (o *Node) UnsetContainerVersion()`

UnsetContainerVersion ensures that no value is present for ContainerVersion, not even an explicit nil
### GetCreatedAt

`func (o *Node) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Node) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Node) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Node) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Node) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Node) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetExternalIpAddress

`func (o *Node) GetExternalIpAddress() string`

GetExternalIpAddress returns the ExternalIpAddress field if non-nil, zero value otherwise.

### GetExternalIpAddressOk

`func (o *Node) GetExternalIpAddressOk() (*string, bool)`

GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddress

`func (o *Node) SetExternalIpAddress(v string)`

SetExternalIpAddress sets ExternalIpAddress field to given value.


### SetExternalIpAddressNil

`func (o *Node) SetExternalIpAddressNil(b bool)`

 SetExternalIpAddressNil sets the value for ExternalIpAddress to be an explicit nil

### UnsetExternalIpAddress
`func (o *Node) UnsetExternalIpAddress()`

UnsetExternalIpAddress ensures that no value is present for ExternalIpAddress, not even an explicit nil
### GetHostName

`func (o *Node) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Node) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *Node) SetHostName(v string)`

SetHostName sets HostName field to given value.


### SetHostNameNil

`func (o *Node) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *Node) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetIp

`func (o *Node) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Node) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Node) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *Node) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *Node) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetKernelVersion

`func (o *Node) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *Node) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *Node) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### SetKernelVersionNil

`func (o *Node) SetKernelVersionNil(b bool)`

 SetKernelVersionNil sets the value for KernelVersion to be an explicit nil

### UnsetKernelVersion
`func (o *Node) UnsetKernelVersion()`

UnsetKernelVersion ensures that no value is present for KernelVersion, not even an explicit nil
### GetKubeProxyVersion

`func (o *Node) GetKubeProxyVersion() string`

GetKubeProxyVersion returns the KubeProxyVersion field if non-nil, zero value otherwise.

### GetKubeProxyVersionOk

`func (o *Node) GetKubeProxyVersionOk() (*string, bool)`

GetKubeProxyVersionOk returns a tuple with the KubeProxyVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeProxyVersion

`func (o *Node) SetKubeProxyVersion(v string)`

SetKubeProxyVersion sets KubeProxyVersion field to given value.


### SetKubeProxyVersionNil

`func (o *Node) SetKubeProxyVersionNil(b bool)`

 SetKubeProxyVersionNil sets the value for KubeProxyVersion to be an explicit nil

### UnsetKubeProxyVersion
`func (o *Node) UnsetKubeProxyVersion()`

UnsetKubeProxyVersion ensures that no value is present for KubeProxyVersion, not even an explicit nil
### GetKubeletVersion

`func (o *Node) GetKubeletVersion() string`

GetKubeletVersion returns the KubeletVersion field if non-nil, zero value otherwise.

### GetKubeletVersionOk

`func (o *Node) GetKubeletVersionOk() (*string, bool)`

GetKubeletVersionOk returns a tuple with the KubeletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeletVersion

`func (o *Node) SetKubeletVersion(v string)`

SetKubeletVersion sets KubeletVersion field to given value.


### SetKubeletVersionNil

`func (o *Node) SetKubeletVersionNil(b bool)`

 SetKubeletVersionNil sets the value for KubeletVersion to be an explicit nil

### UnsetKubeletVersion
`func (o *Node) UnsetKubeletVersion()`

UnsetKubeletVersion ensures that no value is present for KubeletVersion, not even an explicit nil
### GetKubernetesVersion

`func (o *Node) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *Node) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *Node) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### SetKubernetesVersionNil

`func (o *Node) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *Node) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetLabels

`func (o *Node) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Node) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Node) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetMachineId

`func (o *Node) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *Node) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *Node) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.


### SetMachineIdNil

`func (o *Node) SetMachineIdNil(b bool)`

 SetMachineIdNil sets the value for MachineId to be an explicit nil

### UnsetMachineId
`func (o *Node) UnsetMachineId()`

UnsetMachineId ensures that no value is present for MachineId, not even an explicit nil
### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.


### GetNodeCapacities

`func (o *Node) GetNodeCapacities() []NodeCapacity`

GetNodeCapacities returns the NodeCapacities field if non-nil, zero value otherwise.

### GetNodeCapacitiesOk

`func (o *Node) GetNodeCapacitiesOk() (*[]NodeCapacity, bool)`

GetNodeCapacitiesOk returns a tuple with the NodeCapacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCapacities

`func (o *Node) SetNodeCapacities(v []NodeCapacity)`

SetNodeCapacities sets NodeCapacities field to given value.


### GetNodeConditions

`func (o *Node) GetNodeConditions() []NodeCondition`

GetNodeConditions returns the NodeConditions field if non-nil, zero value otherwise.

### GetNodeConditionsOk

`func (o *Node) GetNodeConditionsOk() (*[]NodeCondition, bool)`

GetNodeConditionsOk returns a tuple with the NodeConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConditions

`func (o *Node) SetNodeConditions(v []NodeCondition)`

SetNodeConditions sets NodeConditions field to given value.


### GetNodeStatus

`func (o *Node) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *Node) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *Node) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.


### GetOperatingSystem

`func (o *Node) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Node) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Node) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### SetOperatingSystemNil

`func (o *Node) SetOperatingSystemNil(b bool)`

 SetOperatingSystemNil sets the value for OperatingSystem to be an explicit nil

### UnsetOperatingSystem
`func (o *Node) UnsetOperatingSystem()`

UnsetOperatingSystem ensures that no value is present for OperatingSystem, not even an explicit nil
### GetOsImage

`func (o *Node) GetOsImage() string`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *Node) GetOsImageOk() (*string, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *Node) SetOsImage(v string)`

SetOsImage sets OsImage field to given value.


### SetOsImageNil

`func (o *Node) SetOsImageNil(b bool)`

 SetOsImageNil sets the value for OsImage to be an explicit nil

### UnsetOsImage
`func (o *Node) UnsetOsImage()`

UnsetOsImage ensures that no value is present for OsImage, not even an explicit nil
### GetPodCidr

`func (o *Node) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *Node) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *Node) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.


### SetPodCidrNil

`func (o *Node) SetPodCidrNil(b bool)`

 SetPodCidrNil sets the value for PodCidr to be an explicit nil

### UnsetPodCidr
`func (o *Node) UnsetPodCidr()`

UnsetPodCidr ensures that no value is present for PodCidr, not even an explicit nil
### GetSystemUuid

`func (o *Node) GetSystemUuid() string`

GetSystemUuid returns the SystemUuid field if non-nil, zero value otherwise.

### GetSystemUuidOk

`func (o *Node) GetSystemUuidOk() (*string, bool)`

GetSystemUuidOk returns a tuple with the SystemUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemUuid

`func (o *Node) SetSystemUuid(v string)`

SetSystemUuid sets SystemUuid field to given value.


### SetSystemUuidNil

`func (o *Node) SetSystemUuidNil(b bool)`

 SetSystemUuidNil sets the value for SystemUuid to be an explicit nil

### UnsetSystemUuid
`func (o *Node) UnsetSystemUuid()`

UnsetSystemUuid ensures that no value is present for SystemUuid, not even an explicit nil
### GetTaints

`func (o *Node) GetTaints() []string`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *Node) GetTaintsOk() (*[]string, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *Node) SetTaints(v []string)`

SetTaints sets Taints field to given value.


### GetUid

`func (o *Node) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Node) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Node) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Node) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Node) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Node) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


