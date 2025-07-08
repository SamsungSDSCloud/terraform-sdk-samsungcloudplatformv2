# StorageClass

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **NullableString** |  | 
**AllowVolumeExpansion** | **NullableString** |  | 
**Annotations** | **[]string** |  | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **NullableTime** |  | 
**DefaultClass** | **NullableString** |  | 
**Labels** | **[]string** |  | 
**MountOptions** | **[]string** |  | 
**Name** | **string** | Storage Class Name | 
**Provisioner** | **NullableString** |  | 
**PvReclaimPolicy** | **NullableString** |  | 
**StorageClassParameters** | **[]string** |  | 
**Uid** | **string** | Resource ID | 
**VolumeBindingMode** | **NullableString** |  | 
**Yaml** | **string** | YAML | 

## Methods

### NewStorageClass

`func NewStorageClass(age NullableString, allowVolumeExpansion NullableString, annotations []string, clusterId string, clusterName string, createdAt NullableTime, defaultClass NullableString, labels []string, mountOptions []string, name string, provisioner NullableString, pvReclaimPolicy NullableString, storageClassParameters []string, uid string, volumeBindingMode NullableString, yaml string, ) *StorageClass`

NewStorageClass instantiates a new StorageClass object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassWithDefaults

`func NewStorageClassWithDefaults() *StorageClass`

NewStorageClassWithDefaults instantiates a new StorageClass object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *StorageClass) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *StorageClass) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *StorageClass) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *StorageClass) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *StorageClass) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAllowVolumeExpansion

`func (o *StorageClass) GetAllowVolumeExpansion() string`

GetAllowVolumeExpansion returns the AllowVolumeExpansion field if non-nil, zero value otherwise.

### GetAllowVolumeExpansionOk

`func (o *StorageClass) GetAllowVolumeExpansionOk() (*string, bool)`

GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVolumeExpansion

`func (o *StorageClass) SetAllowVolumeExpansion(v string)`

SetAllowVolumeExpansion sets AllowVolumeExpansion field to given value.


### SetAllowVolumeExpansionNil

`func (o *StorageClass) SetAllowVolumeExpansionNil(b bool)`

 SetAllowVolumeExpansionNil sets the value for AllowVolumeExpansion to be an explicit nil

### UnsetAllowVolumeExpansion
`func (o *StorageClass) UnsetAllowVolumeExpansion()`

UnsetAllowVolumeExpansion ensures that no value is present for AllowVolumeExpansion, not even an explicit nil
### GetAnnotations

`func (o *StorageClass) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *StorageClass) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *StorageClass) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### SetAnnotationsNil

`func (o *StorageClass) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *StorageClass) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetClusterId

`func (o *StorageClass) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StorageClass) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StorageClass) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *StorageClass) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *StorageClass) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *StorageClass) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *StorageClass) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorageClass) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorageClass) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *StorageClass) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StorageClass) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDefaultClass

`func (o *StorageClass) GetDefaultClass() string`

GetDefaultClass returns the DefaultClass field if non-nil, zero value otherwise.

### GetDefaultClassOk

`func (o *StorageClass) GetDefaultClassOk() (*string, bool)`

GetDefaultClassOk returns a tuple with the DefaultClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClass

`func (o *StorageClass) SetDefaultClass(v string)`

SetDefaultClass sets DefaultClass field to given value.


### SetDefaultClassNil

`func (o *StorageClass) SetDefaultClassNil(b bool)`

 SetDefaultClassNil sets the value for DefaultClass to be an explicit nil

### UnsetDefaultClass
`func (o *StorageClass) UnsetDefaultClass()`

UnsetDefaultClass ensures that no value is present for DefaultClass, not even an explicit nil
### GetLabels

`func (o *StorageClass) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StorageClass) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StorageClass) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### SetLabelsNil

`func (o *StorageClass) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *StorageClass) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetMountOptions

`func (o *StorageClass) GetMountOptions() []string`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *StorageClass) GetMountOptionsOk() (*[]string, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *StorageClass) SetMountOptions(v []string)`

SetMountOptions sets MountOptions field to given value.


### SetMountOptionsNil

`func (o *StorageClass) SetMountOptionsNil(b bool)`

 SetMountOptionsNil sets the value for MountOptions to be an explicit nil

### UnsetMountOptions
`func (o *StorageClass) UnsetMountOptions()`

UnsetMountOptions ensures that no value is present for MountOptions, not even an explicit nil
### GetName

`func (o *StorageClass) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageClass) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageClass) SetName(v string)`

SetName sets Name field to given value.


### GetProvisioner

`func (o *StorageClass) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *StorageClass) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *StorageClass) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### SetProvisionerNil

`func (o *StorageClass) SetProvisionerNil(b bool)`

 SetProvisionerNil sets the value for Provisioner to be an explicit nil

### UnsetProvisioner
`func (o *StorageClass) UnsetProvisioner()`

UnsetProvisioner ensures that no value is present for Provisioner, not even an explicit nil
### GetPvReclaimPolicy

`func (o *StorageClass) GetPvReclaimPolicy() string`

GetPvReclaimPolicy returns the PvReclaimPolicy field if non-nil, zero value otherwise.

### GetPvReclaimPolicyOk

`func (o *StorageClass) GetPvReclaimPolicyOk() (*string, bool)`

GetPvReclaimPolicyOk returns a tuple with the PvReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvReclaimPolicy

`func (o *StorageClass) SetPvReclaimPolicy(v string)`

SetPvReclaimPolicy sets PvReclaimPolicy field to given value.


### SetPvReclaimPolicyNil

`func (o *StorageClass) SetPvReclaimPolicyNil(b bool)`

 SetPvReclaimPolicyNil sets the value for PvReclaimPolicy to be an explicit nil

### UnsetPvReclaimPolicy
`func (o *StorageClass) UnsetPvReclaimPolicy()`

UnsetPvReclaimPolicy ensures that no value is present for PvReclaimPolicy, not even an explicit nil
### GetStorageClassParameters

`func (o *StorageClass) GetStorageClassParameters() []string`

GetStorageClassParameters returns the StorageClassParameters field if non-nil, zero value otherwise.

### GetStorageClassParametersOk

`func (o *StorageClass) GetStorageClassParametersOk() (*[]string, bool)`

GetStorageClassParametersOk returns a tuple with the StorageClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassParameters

`func (o *StorageClass) SetStorageClassParameters(v []string)`

SetStorageClassParameters sets StorageClassParameters field to given value.


### SetStorageClassParametersNil

`func (o *StorageClass) SetStorageClassParametersNil(b bool)`

 SetStorageClassParametersNil sets the value for StorageClassParameters to be an explicit nil

### UnsetStorageClassParameters
`func (o *StorageClass) UnsetStorageClassParameters()`

UnsetStorageClassParameters ensures that no value is present for StorageClassParameters, not even an explicit nil
### GetUid

`func (o *StorageClass) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *StorageClass) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *StorageClass) SetUid(v string)`

SetUid sets Uid field to given value.


### GetVolumeBindingMode

`func (o *StorageClass) GetVolumeBindingMode() string`

GetVolumeBindingMode returns the VolumeBindingMode field if non-nil, zero value otherwise.

### GetVolumeBindingModeOk

`func (o *StorageClass) GetVolumeBindingModeOk() (*string, bool)`

GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBindingMode

`func (o *StorageClass) SetVolumeBindingMode(v string)`

SetVolumeBindingMode sets VolumeBindingMode field to given value.


### SetVolumeBindingModeNil

`func (o *StorageClass) SetVolumeBindingModeNil(b bool)`

 SetVolumeBindingModeNil sets the value for VolumeBindingMode to be an explicit nil

### UnsetVolumeBindingMode
`func (o *StorageClass) UnsetVolumeBindingMode()`

UnsetVolumeBindingMode ensures that no value is present for VolumeBindingMode, not even an explicit nil
### GetYaml

`func (o *StorageClass) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *StorageClass) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *StorageClass) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


