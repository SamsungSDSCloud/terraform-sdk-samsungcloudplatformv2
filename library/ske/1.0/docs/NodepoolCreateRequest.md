# NodepoolCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | Cluster ID | 
**CustomImageId** | Pointer to **NullableString** |  | [optional] 
**DesiredNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**ImageOs** | **string** | Image OS | 
**ImageOsVersion** | **string** | Image OS Version | 
**IsAutoRecovery** | **bool** | Is Auto Recovery | 
**IsAutoScale** | **bool** | Is Auto Scale | 
**KeypairName** | **string** | Keypair Name | 
**KubernetesVersion** | **string** | Kubernetes Version | 
**Labels** | Pointer to [**[]NodepoolLabel**](NodepoolLabel.md) |  | [optional] 
**MaxNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**MinNodeCount** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** | Nodepool Name | 
**ServerTypeId** | **string** | Server Type ID | 
**Taints** | Pointer to [**[]NodepoolTaint**](NodepoolTaint.md) |  | [optional] 
**VolumeSize** | **int32** | Volume Size | 
**VolumeTypeName** | **string** | Volume Type Name | 

## Methods

### NewNodepoolCreateRequest

`func NewNodepoolCreateRequest(clusterId string, imageOs string, imageOsVersion string, isAutoRecovery bool, isAutoScale bool, keypairName string, kubernetesVersion string, name string, serverTypeId string, volumeSize int32, volumeTypeName string, ) *NodepoolCreateRequest`

NewNodepoolCreateRequest instantiates a new NodepoolCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolCreateRequestWithDefaults

`func NewNodepoolCreateRequestWithDefaults() *NodepoolCreateRequest`

NewNodepoolCreateRequestWithDefaults instantiates a new NodepoolCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *NodepoolCreateRequest) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NodepoolCreateRequest) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NodepoolCreateRequest) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCustomImageId

`func (o *NodepoolCreateRequest) GetCustomImageId() string`

GetCustomImageId returns the CustomImageId field if non-nil, zero value otherwise.

### GetCustomImageIdOk

`func (o *NodepoolCreateRequest) GetCustomImageIdOk() (*string, bool)`

GetCustomImageIdOk returns a tuple with the CustomImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomImageId

`func (o *NodepoolCreateRequest) SetCustomImageId(v string)`

SetCustomImageId sets CustomImageId field to given value.

### HasCustomImageId

`func (o *NodepoolCreateRequest) HasCustomImageId() bool`

HasCustomImageId returns a boolean if a field has been set.

### SetCustomImageIdNil

`func (o *NodepoolCreateRequest) SetCustomImageIdNil(b bool)`

 SetCustomImageIdNil sets the value for CustomImageId to be an explicit nil

### UnsetCustomImageId
`func (o *NodepoolCreateRequest) UnsetCustomImageId()`

UnsetCustomImageId ensures that no value is present for CustomImageId, not even an explicit nil
### GetDesiredNodeCount

`func (o *NodepoolCreateRequest) GetDesiredNodeCount() int32`

GetDesiredNodeCount returns the DesiredNodeCount field if non-nil, zero value otherwise.

### GetDesiredNodeCountOk

`func (o *NodepoolCreateRequest) GetDesiredNodeCountOk() (*int32, bool)`

GetDesiredNodeCountOk returns a tuple with the DesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodeCount

`func (o *NodepoolCreateRequest) SetDesiredNodeCount(v int32)`

SetDesiredNodeCount sets DesiredNodeCount field to given value.

### HasDesiredNodeCount

`func (o *NodepoolCreateRequest) HasDesiredNodeCount() bool`

HasDesiredNodeCount returns a boolean if a field has been set.

### SetDesiredNodeCountNil

`func (o *NodepoolCreateRequest) SetDesiredNodeCountNil(b bool)`

 SetDesiredNodeCountNil sets the value for DesiredNodeCount to be an explicit nil

### UnsetDesiredNodeCount
`func (o *NodepoolCreateRequest) UnsetDesiredNodeCount()`

UnsetDesiredNodeCount ensures that no value is present for DesiredNodeCount, not even an explicit nil
### GetImageOs

`func (o *NodepoolCreateRequest) GetImageOs() string`

GetImageOs returns the ImageOs field if non-nil, zero value otherwise.

### GetImageOsOk

`func (o *NodepoolCreateRequest) GetImageOsOk() (*string, bool)`

GetImageOsOk returns a tuple with the ImageOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOs

`func (o *NodepoolCreateRequest) SetImageOs(v string)`

SetImageOs sets ImageOs field to given value.


### GetImageOsVersion

`func (o *NodepoolCreateRequest) GetImageOsVersion() string`

GetImageOsVersion returns the ImageOsVersion field if non-nil, zero value otherwise.

### GetImageOsVersionOk

`func (o *NodepoolCreateRequest) GetImageOsVersionOk() (*string, bool)`

GetImageOsVersionOk returns a tuple with the ImageOsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageOsVersion

`func (o *NodepoolCreateRequest) SetImageOsVersion(v string)`

SetImageOsVersion sets ImageOsVersion field to given value.


### GetIsAutoRecovery

`func (o *NodepoolCreateRequest) GetIsAutoRecovery() bool`

GetIsAutoRecovery returns the IsAutoRecovery field if non-nil, zero value otherwise.

### GetIsAutoRecoveryOk

`func (o *NodepoolCreateRequest) GetIsAutoRecoveryOk() (*bool, bool)`

GetIsAutoRecoveryOk returns a tuple with the IsAutoRecovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRecovery

`func (o *NodepoolCreateRequest) SetIsAutoRecovery(v bool)`

SetIsAutoRecovery sets IsAutoRecovery field to given value.


### GetIsAutoScale

`func (o *NodepoolCreateRequest) GetIsAutoScale() bool`

GetIsAutoScale returns the IsAutoScale field if non-nil, zero value otherwise.

### GetIsAutoScaleOk

`func (o *NodepoolCreateRequest) GetIsAutoScaleOk() (*bool, bool)`

GetIsAutoScaleOk returns a tuple with the IsAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoScale

`func (o *NodepoolCreateRequest) SetIsAutoScale(v bool)`

SetIsAutoScale sets IsAutoScale field to given value.


### GetKeypairName

`func (o *NodepoolCreateRequest) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *NodepoolCreateRequest) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *NodepoolCreateRequest) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.


### GetKubernetesVersion

`func (o *NodepoolCreateRequest) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *NodepoolCreateRequest) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *NodepoolCreateRequest) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetLabels

`func (o *NodepoolCreateRequest) GetLabels() []NodepoolLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NodepoolCreateRequest) GetLabelsOk() (*[]NodepoolLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NodepoolCreateRequest) SetLabels(v []NodepoolLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NodepoolCreateRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *NodepoolCreateRequest) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *NodepoolCreateRequest) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMaxNodeCount

`func (o *NodepoolCreateRequest) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *NodepoolCreateRequest) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *NodepoolCreateRequest) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *NodepoolCreateRequest) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### SetMaxNodeCountNil

`func (o *NodepoolCreateRequest) SetMaxNodeCountNil(b bool)`

 SetMaxNodeCountNil sets the value for MaxNodeCount to be an explicit nil

### UnsetMaxNodeCount
`func (o *NodepoolCreateRequest) UnsetMaxNodeCount()`

UnsetMaxNodeCount ensures that no value is present for MaxNodeCount, not even an explicit nil
### GetMinNodeCount

`func (o *NodepoolCreateRequest) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *NodepoolCreateRequest) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *NodepoolCreateRequest) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *NodepoolCreateRequest) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### SetMinNodeCountNil

`func (o *NodepoolCreateRequest) SetMinNodeCountNil(b bool)`

 SetMinNodeCountNil sets the value for MinNodeCount to be an explicit nil

### UnsetMinNodeCount
`func (o *NodepoolCreateRequest) UnsetMinNodeCount()`

UnsetMinNodeCount ensures that no value is present for MinNodeCount, not even an explicit nil
### GetName

`func (o *NodepoolCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodepoolCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodepoolCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetServerTypeId

`func (o *NodepoolCreateRequest) GetServerTypeId() string`

GetServerTypeId returns the ServerTypeId field if non-nil, zero value otherwise.

### GetServerTypeIdOk

`func (o *NodepoolCreateRequest) GetServerTypeIdOk() (*string, bool)`

GetServerTypeIdOk returns a tuple with the ServerTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeId

`func (o *NodepoolCreateRequest) SetServerTypeId(v string)`

SetServerTypeId sets ServerTypeId field to given value.


### GetTaints

`func (o *NodepoolCreateRequest) GetTaints() []NodepoolTaint`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *NodepoolCreateRequest) GetTaintsOk() (*[]NodepoolTaint, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *NodepoolCreateRequest) SetTaints(v []NodepoolTaint)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *NodepoolCreateRequest) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### SetTaintsNil

`func (o *NodepoolCreateRequest) SetTaintsNil(b bool)`

 SetTaintsNil sets the value for Taints to be an explicit nil

### UnsetTaints
`func (o *NodepoolCreateRequest) UnsetTaints()`

UnsetTaints ensures that no value is present for Taints, not even an explicit nil
### GetVolumeSize

`func (o *NodepoolCreateRequest) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *NodepoolCreateRequest) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *NodepoolCreateRequest) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetVolumeTypeName

`func (o *NodepoolCreateRequest) GetVolumeTypeName() string`

GetVolumeTypeName returns the VolumeTypeName field if non-nil, zero value otherwise.

### GetVolumeTypeNameOk

`func (o *NodepoolCreateRequest) GetVolumeTypeNameOk() (*string, bool)`

GetVolumeTypeNameOk returns a tuple with the VolumeTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeName

`func (o *NodepoolCreateRequest) SetVolumeTypeName(v string)`

SetVolumeTypeName sets VolumeTypeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


