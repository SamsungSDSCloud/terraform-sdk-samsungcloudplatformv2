# Pvc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** |  | 
**Age** | **NullableString** |  | 
**Annotations** | **[]string** |  | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **NullableTime** |  | 
**Labels** | **[]string** |  | 
**Name** | **string** | PVC Name | 
**NamespaceName** | **string** | Namespace Name | 
**PvName** | **NullableString** |  | 
**PvVolumeMode** | **NullableString** |  | 
**PvcSize** | **NullableString** |  | 
**PvcStatus** | **NullableString** |  | 
**StorageClassName** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewPvc

`func NewPvc(accessModes []string, age NullableString, annotations []string, clusterId string, clusterName string, createdAt NullableTime, labels []string, name string, namespaceName string, pvName NullableString, pvVolumeMode NullableString, pvcSize NullableString, pvcStatus NullableString, storageClassName NullableString, uid string, yaml string, ) *Pvc`

NewPvc instantiates a new Pvc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvcWithDefaults

`func NewPvcWithDefaults() *Pvc`

NewPvcWithDefaults instantiates a new Pvc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *Pvc) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *Pvc) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *Pvc) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.


### SetAccessModesNil

`func (o *Pvc) SetAccessModesNil(b bool)`

 SetAccessModesNil sets the value for AccessModes to be an explicit nil

### UnsetAccessModes
`func (o *Pvc) UnsetAccessModes()`

UnsetAccessModes ensures that no value is present for AccessModes, not even an explicit nil
### GetAge

`func (o *Pvc) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Pvc) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Pvc) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *Pvc) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *Pvc) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAnnotations

`func (o *Pvc) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Pvc) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Pvc) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### SetAnnotationsNil

`func (o *Pvc) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *Pvc) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetClusterId

`func (o *Pvc) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Pvc) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Pvc) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Pvc) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Pvc) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Pvc) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *Pvc) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pvc) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pvc) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Pvc) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Pvc) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLabels

`func (o *Pvc) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Pvc) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Pvc) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### SetLabelsNil

`func (o *Pvc) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Pvc) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetName

`func (o *Pvc) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pvc) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pvc) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *Pvc) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Pvc) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Pvc) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetPvName

`func (o *Pvc) GetPvName() string`

GetPvName returns the PvName field if non-nil, zero value otherwise.

### GetPvNameOk

`func (o *Pvc) GetPvNameOk() (*string, bool)`

GetPvNameOk returns a tuple with the PvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvName

`func (o *Pvc) SetPvName(v string)`

SetPvName sets PvName field to given value.


### SetPvNameNil

`func (o *Pvc) SetPvNameNil(b bool)`

 SetPvNameNil sets the value for PvName to be an explicit nil

### UnsetPvName
`func (o *Pvc) UnsetPvName()`

UnsetPvName ensures that no value is present for PvName, not even an explicit nil
### GetPvVolumeMode

`func (o *Pvc) GetPvVolumeMode() string`

GetPvVolumeMode returns the PvVolumeMode field if non-nil, zero value otherwise.

### GetPvVolumeModeOk

`func (o *Pvc) GetPvVolumeModeOk() (*string, bool)`

GetPvVolumeModeOk returns a tuple with the PvVolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvVolumeMode

`func (o *Pvc) SetPvVolumeMode(v string)`

SetPvVolumeMode sets PvVolumeMode field to given value.


### SetPvVolumeModeNil

`func (o *Pvc) SetPvVolumeModeNil(b bool)`

 SetPvVolumeModeNil sets the value for PvVolumeMode to be an explicit nil

### UnsetPvVolumeMode
`func (o *Pvc) UnsetPvVolumeMode()`

UnsetPvVolumeMode ensures that no value is present for PvVolumeMode, not even an explicit nil
### GetPvcSize

`func (o *Pvc) GetPvcSize() string`

GetPvcSize returns the PvcSize field if non-nil, zero value otherwise.

### GetPvcSizeOk

`func (o *Pvc) GetPvcSizeOk() (*string, bool)`

GetPvcSizeOk returns a tuple with the PvcSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcSize

`func (o *Pvc) SetPvcSize(v string)`

SetPvcSize sets PvcSize field to given value.


### SetPvcSizeNil

`func (o *Pvc) SetPvcSizeNil(b bool)`

 SetPvcSizeNil sets the value for PvcSize to be an explicit nil

### UnsetPvcSize
`func (o *Pvc) UnsetPvcSize()`

UnsetPvcSize ensures that no value is present for PvcSize, not even an explicit nil
### GetPvcStatus

`func (o *Pvc) GetPvcStatus() string`

GetPvcStatus returns the PvcStatus field if non-nil, zero value otherwise.

### GetPvcStatusOk

`func (o *Pvc) GetPvcStatusOk() (*string, bool)`

GetPvcStatusOk returns a tuple with the PvcStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcStatus

`func (o *Pvc) SetPvcStatus(v string)`

SetPvcStatus sets PvcStatus field to given value.


### SetPvcStatusNil

`func (o *Pvc) SetPvcStatusNil(b bool)`

 SetPvcStatusNil sets the value for PvcStatus to be an explicit nil

### UnsetPvcStatus
`func (o *Pvc) UnsetPvcStatus()`

UnsetPvcStatus ensures that no value is present for PvcStatus, not even an explicit nil
### GetStorageClassName

`func (o *Pvc) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *Pvc) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *Pvc) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.


### SetStorageClassNameNil

`func (o *Pvc) SetStorageClassNameNil(b bool)`

 SetStorageClassNameNil sets the value for StorageClassName to be an explicit nil

### UnsetStorageClassName
`func (o *Pvc) UnsetStorageClassName()`

UnsetStorageClassName ensures that no value is present for StorageClassName, not even an explicit nil
### GetUid

`func (o *Pvc) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Pvc) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Pvc) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Pvc) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Pvc) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Pvc) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


