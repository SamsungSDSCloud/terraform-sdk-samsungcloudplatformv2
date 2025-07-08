# NodeSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**ContainerVersion** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**ExternalIpAddress** | **NullableString** |  | 
**Ip** | **NullableString** |  | 
**KernelVersion** | **NullableString** |  | 
**KubernetesVersion** | **NullableString** |  | 
**Name** | **string** | Node Name | 
**NodeStatus** | **string** | Node Status | 
**OsImage** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewNodeSummary

`func NewNodeSummary(age string, clusterId string, containerVersion NullableString, createdAt time.Time, externalIpAddress NullableString, ip NullableString, kernelVersion NullableString, kubernetesVersion NullableString, name string, nodeStatus string, osImage NullableString, uid string, ) *NodeSummary`

NewNodeSummary instantiates a new NodeSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeSummaryWithDefaults

`func NewNodeSummaryWithDefaults() *NodeSummary`

NewNodeSummaryWithDefaults instantiates a new NodeSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *NodeSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *NodeSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *NodeSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *NodeSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NodeSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NodeSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetContainerVersion

`func (o *NodeSummary) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *NodeSummary) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *NodeSummary) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.


### SetContainerVersionNil

`func (o *NodeSummary) SetContainerVersionNil(b bool)`

 SetContainerVersionNil sets the value for ContainerVersion to be an explicit nil

### UnsetContainerVersion
`func (o *NodeSummary) UnsetContainerVersion()`

UnsetContainerVersion ensures that no value is present for ContainerVersion, not even an explicit nil
### GetCreatedAt

`func (o *NodeSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NodeSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NodeSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExternalIpAddress

`func (o *NodeSummary) GetExternalIpAddress() string`

GetExternalIpAddress returns the ExternalIpAddress field if non-nil, zero value otherwise.

### GetExternalIpAddressOk

`func (o *NodeSummary) GetExternalIpAddressOk() (*string, bool)`

GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddress

`func (o *NodeSummary) SetExternalIpAddress(v string)`

SetExternalIpAddress sets ExternalIpAddress field to given value.


### SetExternalIpAddressNil

`func (o *NodeSummary) SetExternalIpAddressNil(b bool)`

 SetExternalIpAddressNil sets the value for ExternalIpAddress to be an explicit nil

### UnsetExternalIpAddress
`func (o *NodeSummary) UnsetExternalIpAddress()`

UnsetExternalIpAddress ensures that no value is present for ExternalIpAddress, not even an explicit nil
### GetIp

`func (o *NodeSummary) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *NodeSummary) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *NodeSummary) SetIp(v string)`

SetIp sets Ip field to given value.


### SetIpNil

`func (o *NodeSummary) SetIpNil(b bool)`

 SetIpNil sets the value for Ip to be an explicit nil

### UnsetIp
`func (o *NodeSummary) UnsetIp()`

UnsetIp ensures that no value is present for Ip, not even an explicit nil
### GetKernelVersion

`func (o *NodeSummary) GetKernelVersion() string`

GetKernelVersion returns the KernelVersion field if non-nil, zero value otherwise.

### GetKernelVersionOk

`func (o *NodeSummary) GetKernelVersionOk() (*string, bool)`

GetKernelVersionOk returns a tuple with the KernelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelVersion

`func (o *NodeSummary) SetKernelVersion(v string)`

SetKernelVersion sets KernelVersion field to given value.


### SetKernelVersionNil

`func (o *NodeSummary) SetKernelVersionNil(b bool)`

 SetKernelVersionNil sets the value for KernelVersion to be an explicit nil

### UnsetKernelVersion
`func (o *NodeSummary) UnsetKernelVersion()`

UnsetKernelVersion ensures that no value is present for KernelVersion, not even an explicit nil
### GetKubernetesVersion

`func (o *NodeSummary) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *NodeSummary) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *NodeSummary) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### SetKubernetesVersionNil

`func (o *NodeSummary) SetKubernetesVersionNil(b bool)`

 SetKubernetesVersionNil sets the value for KubernetesVersion to be an explicit nil

### UnsetKubernetesVersion
`func (o *NodeSummary) UnsetKubernetesVersion()`

UnsetKubernetesVersion ensures that no value is present for KubernetesVersion, not even an explicit nil
### GetName

`func (o *NodeSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNodeStatus

`func (o *NodeSummary) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *NodeSummary) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *NodeSummary) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.


### GetOsImage

`func (o *NodeSummary) GetOsImage() string`

GetOsImage returns the OsImage field if non-nil, zero value otherwise.

### GetOsImageOk

`func (o *NodeSummary) GetOsImageOk() (*string, bool)`

GetOsImageOk returns a tuple with the OsImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsImage

`func (o *NodeSummary) SetOsImage(v string)`

SetOsImage sets OsImage field to given value.


### SetOsImageNil

`func (o *NodeSummary) SetOsImageNil(b bool)`

 SetOsImageNil sets the value for OsImage to be an explicit nil

### UnsetOsImage
`func (o *NodeSummary) UnsetOsImage()`

UnsetOsImage ensures that no value is present for OsImage, not even an explicit nil
### GetUid

`func (o *NodeSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NodeSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NodeSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


