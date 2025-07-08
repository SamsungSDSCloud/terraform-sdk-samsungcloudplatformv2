# NodeCapacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **string** | CPU | 
**EphemeralStorage** | **string** | Ephemeral Storage | 
**Gpu** | Pointer to **NullableString** |  | [optional] 
**Hugepages1Gi** | **string** | Hugepages 1Gi | 
**Hugepages2Mi** | **string** | Hugepages 2Mi | 
**Memory** | **string** | Memory | 
**NodeCapacityType** | **string** | Node Capacity Type | 
**PodCount** | **string** | Pod Count | 

## Methods

### NewNodeCapacity

`func NewNodeCapacity(cpu string, ephemeralStorage string, hugepages1Gi string, hugepages2Mi string, memory string, nodeCapacityType string, podCount string, ) *NodeCapacity`

NewNodeCapacity instantiates a new NodeCapacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeCapacityWithDefaults

`func NewNodeCapacityWithDefaults() *NodeCapacity`

NewNodeCapacityWithDefaults instantiates a new NodeCapacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *NodeCapacity) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeCapacity) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeCapacity) SetCpu(v string)`

SetCpu sets Cpu field to given value.


### GetEphemeralStorage

`func (o *NodeCapacity) GetEphemeralStorage() string`

GetEphemeralStorage returns the EphemeralStorage field if non-nil, zero value otherwise.

### GetEphemeralStorageOk

`func (o *NodeCapacity) GetEphemeralStorageOk() (*string, bool)`

GetEphemeralStorageOk returns a tuple with the EphemeralStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStorage

`func (o *NodeCapacity) SetEphemeralStorage(v string)`

SetEphemeralStorage sets EphemeralStorage field to given value.


### GetGpu

`func (o *NodeCapacity) GetGpu() string`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *NodeCapacity) GetGpuOk() (*string, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *NodeCapacity) SetGpu(v string)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *NodeCapacity) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### SetGpuNil

`func (o *NodeCapacity) SetGpuNil(b bool)`

 SetGpuNil sets the value for Gpu to be an explicit nil

### UnsetGpu
`func (o *NodeCapacity) UnsetGpu()`

UnsetGpu ensures that no value is present for Gpu, not even an explicit nil
### GetHugepages1Gi

`func (o *NodeCapacity) GetHugepages1Gi() string`

GetHugepages1Gi returns the Hugepages1Gi field if non-nil, zero value otherwise.

### GetHugepages1GiOk

`func (o *NodeCapacity) GetHugepages1GiOk() (*string, bool)`

GetHugepages1GiOk returns a tuple with the Hugepages1Gi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages1Gi

`func (o *NodeCapacity) SetHugepages1Gi(v string)`

SetHugepages1Gi sets Hugepages1Gi field to given value.


### GetHugepages2Mi

`func (o *NodeCapacity) GetHugepages2Mi() string`

GetHugepages2Mi returns the Hugepages2Mi field if non-nil, zero value otherwise.

### GetHugepages2MiOk

`func (o *NodeCapacity) GetHugepages2MiOk() (*string, bool)`

GetHugepages2MiOk returns a tuple with the Hugepages2Mi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages2Mi

`func (o *NodeCapacity) SetHugepages2Mi(v string)`

SetHugepages2Mi sets Hugepages2Mi field to given value.


### GetMemory

`func (o *NodeCapacity) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NodeCapacity) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NodeCapacity) SetMemory(v string)`

SetMemory sets Memory field to given value.


### GetNodeCapacityType

`func (o *NodeCapacity) GetNodeCapacityType() string`

GetNodeCapacityType returns the NodeCapacityType field if non-nil, zero value otherwise.

### GetNodeCapacityTypeOk

`func (o *NodeCapacity) GetNodeCapacityTypeOk() (*string, bool)`

GetNodeCapacityTypeOk returns a tuple with the NodeCapacityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCapacityType

`func (o *NodeCapacity) SetNodeCapacityType(v string)`

SetNodeCapacityType sets NodeCapacityType field to given value.


### GetPodCount

`func (o *NodeCapacity) GetPodCount() string`

GetPodCount returns the PodCount field if non-nil, zero value otherwise.

### GetPodCountOk

`func (o *NodeCapacity) GetPodCountOk() (*string, bool)`

GetPodCountOk returns a tuple with the PodCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCount

`func (o *NodeCapacity) SetPodCount(v string)`

SetPodCount sets PodCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


