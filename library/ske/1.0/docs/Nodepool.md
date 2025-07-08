# Nodepool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AutoRecoveryEnabled** | **bool** | Is Auto Recovery | 
**AutoScaleEnabled** | **bool** | Is Auto Scale | 
**Cluster** | [**ClusterOfNodepool**](ClusterOfNodepool.md) | Cluster | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**CurrentNodeCount** | **int32** | Current Node Count | 
**DesiredNodeCount** | **int32** | Desired Node Count | 
**Id** | **string** | Nodepool ID | 
**Image** | [**Image**](Image.md) | Image | 
**Keypair** | [**Keypair**](Keypair.md) | Keypair Name | 
**KubernetesVersion** | **string** | Kubernetes Version | 
**Labels** | Pointer to [**[]NodepoolLabel**](NodepoolLabel.md) |  | [optional] 
**MaxNodeCount** | **int32** | Max Node Count | 
**MinNodeCount** | **int32** | Min Node Count | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Nodepool Name | 
**ServerType** | [**ServerType**](ServerType.md) | Server Type | 
**Status** | **string** | Node Pool Status | 
**Taints** | Pointer to [**[]NodepoolTaint**](NodepoolTaint.md) |  | [optional] 
**VolumeSize** | **int32** | Volume Size | 
**VolumeType** | [**VolumeType**](VolumeType.md) | Volume Type | 

## Methods

### NewNodepool

`func NewNodepool(accountId string, autoRecoveryEnabled bool, autoScaleEnabled bool, cluster ClusterOfNodepool, createdAt time.Time, createdBy string, currentNodeCount int32, desiredNodeCount int32, id string, image Image, keypair Keypair, kubernetesVersion string, maxNodeCount int32, minNodeCount int32, modifiedAt time.Time, modifiedBy string, name string, serverType ServerType, status string, volumeSize int32, volumeType VolumeType, ) *Nodepool`

NewNodepool instantiates a new Nodepool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolWithDefaults

`func NewNodepoolWithDefaults() *Nodepool`

NewNodepoolWithDefaults instantiates a new Nodepool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Nodepool) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Nodepool) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Nodepool) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAutoRecoveryEnabled

`func (o *Nodepool) GetAutoRecoveryEnabled() bool`

GetAutoRecoveryEnabled returns the AutoRecoveryEnabled field if non-nil, zero value otherwise.

### GetAutoRecoveryEnabledOk

`func (o *Nodepool) GetAutoRecoveryEnabledOk() (*bool, bool)`

GetAutoRecoveryEnabledOk returns a tuple with the AutoRecoveryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoveryEnabled

`func (o *Nodepool) SetAutoRecoveryEnabled(v bool)`

SetAutoRecoveryEnabled sets AutoRecoveryEnabled field to given value.


### GetAutoScaleEnabled

`func (o *Nodepool) GetAutoScaleEnabled() bool`

GetAutoScaleEnabled returns the AutoScaleEnabled field if non-nil, zero value otherwise.

### GetAutoScaleEnabledOk

`func (o *Nodepool) GetAutoScaleEnabledOk() (*bool, bool)`

GetAutoScaleEnabledOk returns a tuple with the AutoScaleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoScaleEnabled

`func (o *Nodepool) SetAutoScaleEnabled(v bool)`

SetAutoScaleEnabled sets AutoScaleEnabled field to given value.


### GetCluster

`func (o *Nodepool) GetCluster() ClusterOfNodepool`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *Nodepool) GetClusterOk() (*ClusterOfNodepool, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *Nodepool) SetCluster(v ClusterOfNodepool)`

SetCluster sets Cluster field to given value.


### GetCreatedAt

`func (o *Nodepool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Nodepool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Nodepool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Nodepool) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Nodepool) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Nodepool) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCurrentNodeCount

`func (o *Nodepool) GetCurrentNodeCount() int32`

GetCurrentNodeCount returns the CurrentNodeCount field if non-nil, zero value otherwise.

### GetCurrentNodeCountOk

`func (o *Nodepool) GetCurrentNodeCountOk() (*int32, bool)`

GetCurrentNodeCountOk returns a tuple with the CurrentNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodeCount

`func (o *Nodepool) SetCurrentNodeCount(v int32)`

SetCurrentNodeCount sets CurrentNodeCount field to given value.


### GetDesiredNodeCount

`func (o *Nodepool) GetDesiredNodeCount() int32`

GetDesiredNodeCount returns the DesiredNodeCount field if non-nil, zero value otherwise.

### GetDesiredNodeCountOk

`func (o *Nodepool) GetDesiredNodeCountOk() (*int32, bool)`

GetDesiredNodeCountOk returns a tuple with the DesiredNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNodeCount

`func (o *Nodepool) SetDesiredNodeCount(v int32)`

SetDesiredNodeCount sets DesiredNodeCount field to given value.


### GetId

`func (o *Nodepool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Nodepool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Nodepool) SetId(v string)`

SetId sets Id field to given value.


### GetImage

`func (o *Nodepool) GetImage() Image`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Nodepool) GetImageOk() (*Image, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Nodepool) SetImage(v Image)`

SetImage sets Image field to given value.


### GetKeypair

`func (o *Nodepool) GetKeypair() Keypair`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *Nodepool) GetKeypairOk() (*Keypair, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *Nodepool) SetKeypair(v Keypair)`

SetKeypair sets Keypair field to given value.


### GetKubernetesVersion

`func (o *Nodepool) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *Nodepool) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *Nodepool) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.


### GetLabels

`func (o *Nodepool) GetLabels() []NodepoolLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Nodepool) GetLabelsOk() (*[]NodepoolLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Nodepool) SetLabels(v []NodepoolLabel)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Nodepool) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *Nodepool) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Nodepool) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMaxNodeCount

`func (o *Nodepool) GetMaxNodeCount() int32`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *Nodepool) GetMaxNodeCountOk() (*int32, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *Nodepool) SetMaxNodeCount(v int32)`

SetMaxNodeCount sets MaxNodeCount field to given value.


### GetMinNodeCount

`func (o *Nodepool) GetMinNodeCount() int32`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *Nodepool) GetMinNodeCountOk() (*int32, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *Nodepool) SetMinNodeCount(v int32)`

SetMinNodeCount sets MinNodeCount field to given value.


### GetModifiedAt

`func (o *Nodepool) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Nodepool) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Nodepool) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Nodepool) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Nodepool) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Nodepool) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *Nodepool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Nodepool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Nodepool) SetName(v string)`

SetName sets Name field to given value.


### GetServerType

`func (o *Nodepool) GetServerType() ServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *Nodepool) GetServerTypeOk() (*ServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *Nodepool) SetServerType(v ServerType)`

SetServerType sets ServerType field to given value.


### GetStatus

`func (o *Nodepool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Nodepool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Nodepool) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTaints

`func (o *Nodepool) GetTaints() []NodepoolTaint`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *Nodepool) GetTaintsOk() (*[]NodepoolTaint, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *Nodepool) SetTaints(v []NodepoolTaint)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *Nodepool) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### SetTaintsNil

`func (o *Nodepool) SetTaintsNil(b bool)`

 SetTaintsNil sets the value for Taints to be an explicit nil

### UnsetTaints
`func (o *Nodepool) UnsetTaints()`

UnsetTaints ensures that no value is present for Taints, not even an explicit nil
### GetVolumeSize

`func (o *Nodepool) GetVolumeSize() int32`

GetVolumeSize returns the VolumeSize field if non-nil, zero value otherwise.

### GetVolumeSizeOk

`func (o *Nodepool) GetVolumeSizeOk() (*int32, bool)`

GetVolumeSizeOk returns a tuple with the VolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeSize

`func (o *Nodepool) SetVolumeSize(v int32)`

SetVolumeSize sets VolumeSize field to given value.


### GetVolumeType

`func (o *Nodepool) GetVolumeType() VolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *Nodepool) GetVolumeTypeOk() (*VolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *Nodepool) SetVolumeType(v VolumeType)`

SetVolumeType sets VolumeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


