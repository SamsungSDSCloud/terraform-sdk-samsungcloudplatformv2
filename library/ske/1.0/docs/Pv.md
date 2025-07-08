# Pv

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** |  | 
**Age** | **NullableString** |  | 
**Annotations** | **[]string** |  | 
**ClaimRefName** | **NullableString** |  | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **NullableTime** |  | 
**Labels** | **[]string** |  | 
**Name** | **string** | PV Name | 
**PvReason** | **NullableString** |  | 
**PvReclaimPolicy** | **NullableString** |  | 
**PvSize** | **NullableString** |  | 
**PvStatus** | **NullableString** |  | 
**PvVolumeMode** | **NullableString** |  | 
**StorageClassName** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewPv

`func NewPv(accessModes []string, age NullableString, annotations []string, claimRefName NullableString, clusterId string, clusterName string, createdAt NullableTime, labels []string, name string, pvReason NullableString, pvReclaimPolicy NullableString, pvSize NullableString, pvStatus NullableString, pvVolumeMode NullableString, storageClassName NullableString, uid string, yaml string, ) *Pv`

NewPv instantiates a new Pv object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvWithDefaults

`func NewPvWithDefaults() *Pv`

NewPvWithDefaults instantiates a new Pv object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *Pv) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *Pv) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *Pv) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.


### SetAccessModesNil

`func (o *Pv) SetAccessModesNil(b bool)`

 SetAccessModesNil sets the value for AccessModes to be an explicit nil

### UnsetAccessModes
`func (o *Pv) UnsetAccessModes()`

UnsetAccessModes ensures that no value is present for AccessModes, not even an explicit nil
### GetAge

`func (o *Pv) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Pv) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Pv) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *Pv) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *Pv) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAnnotations

`func (o *Pv) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Pv) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Pv) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### SetAnnotationsNil

`func (o *Pv) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *Pv) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetClaimRefName

`func (o *Pv) GetClaimRefName() string`

GetClaimRefName returns the ClaimRefName field if non-nil, zero value otherwise.

### GetClaimRefNameOk

`func (o *Pv) GetClaimRefNameOk() (*string, bool)`

GetClaimRefNameOk returns a tuple with the ClaimRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRefName

`func (o *Pv) SetClaimRefName(v string)`

SetClaimRefName sets ClaimRefName field to given value.


### SetClaimRefNameNil

`func (o *Pv) SetClaimRefNameNil(b bool)`

 SetClaimRefNameNil sets the value for ClaimRefName to be an explicit nil

### UnsetClaimRefName
`func (o *Pv) UnsetClaimRefName()`

UnsetClaimRefName ensures that no value is present for ClaimRefName, not even an explicit nil
### GetClusterId

`func (o *Pv) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Pv) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Pv) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Pv) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Pv) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Pv) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *Pv) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Pv) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Pv) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Pv) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Pv) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetLabels

`func (o *Pv) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Pv) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Pv) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### SetLabelsNil

`func (o *Pv) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *Pv) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetName

`func (o *Pv) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pv) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pv) SetName(v string)`

SetName sets Name field to given value.


### GetPvReason

`func (o *Pv) GetPvReason() string`

GetPvReason returns the PvReason field if non-nil, zero value otherwise.

### GetPvReasonOk

`func (o *Pv) GetPvReasonOk() (*string, bool)`

GetPvReasonOk returns a tuple with the PvReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvReason

`func (o *Pv) SetPvReason(v string)`

SetPvReason sets PvReason field to given value.


### SetPvReasonNil

`func (o *Pv) SetPvReasonNil(b bool)`

 SetPvReasonNil sets the value for PvReason to be an explicit nil

### UnsetPvReason
`func (o *Pv) UnsetPvReason()`

UnsetPvReason ensures that no value is present for PvReason, not even an explicit nil
### GetPvReclaimPolicy

`func (o *Pv) GetPvReclaimPolicy() string`

GetPvReclaimPolicy returns the PvReclaimPolicy field if non-nil, zero value otherwise.

### GetPvReclaimPolicyOk

`func (o *Pv) GetPvReclaimPolicyOk() (*string, bool)`

GetPvReclaimPolicyOk returns a tuple with the PvReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvReclaimPolicy

`func (o *Pv) SetPvReclaimPolicy(v string)`

SetPvReclaimPolicy sets PvReclaimPolicy field to given value.


### SetPvReclaimPolicyNil

`func (o *Pv) SetPvReclaimPolicyNil(b bool)`

 SetPvReclaimPolicyNil sets the value for PvReclaimPolicy to be an explicit nil

### UnsetPvReclaimPolicy
`func (o *Pv) UnsetPvReclaimPolicy()`

UnsetPvReclaimPolicy ensures that no value is present for PvReclaimPolicy, not even an explicit nil
### GetPvSize

`func (o *Pv) GetPvSize() string`

GetPvSize returns the PvSize field if non-nil, zero value otherwise.

### GetPvSizeOk

`func (o *Pv) GetPvSizeOk() (*string, bool)`

GetPvSizeOk returns a tuple with the PvSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvSize

`func (o *Pv) SetPvSize(v string)`

SetPvSize sets PvSize field to given value.


### SetPvSizeNil

`func (o *Pv) SetPvSizeNil(b bool)`

 SetPvSizeNil sets the value for PvSize to be an explicit nil

### UnsetPvSize
`func (o *Pv) UnsetPvSize()`

UnsetPvSize ensures that no value is present for PvSize, not even an explicit nil
### GetPvStatus

`func (o *Pv) GetPvStatus() string`

GetPvStatus returns the PvStatus field if non-nil, zero value otherwise.

### GetPvStatusOk

`func (o *Pv) GetPvStatusOk() (*string, bool)`

GetPvStatusOk returns a tuple with the PvStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvStatus

`func (o *Pv) SetPvStatus(v string)`

SetPvStatus sets PvStatus field to given value.


### SetPvStatusNil

`func (o *Pv) SetPvStatusNil(b bool)`

 SetPvStatusNil sets the value for PvStatus to be an explicit nil

### UnsetPvStatus
`func (o *Pv) UnsetPvStatus()`

UnsetPvStatus ensures that no value is present for PvStatus, not even an explicit nil
### GetPvVolumeMode

`func (o *Pv) GetPvVolumeMode() string`

GetPvVolumeMode returns the PvVolumeMode field if non-nil, zero value otherwise.

### GetPvVolumeModeOk

`func (o *Pv) GetPvVolumeModeOk() (*string, bool)`

GetPvVolumeModeOk returns a tuple with the PvVolumeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvVolumeMode

`func (o *Pv) SetPvVolumeMode(v string)`

SetPvVolumeMode sets PvVolumeMode field to given value.


### SetPvVolumeModeNil

`func (o *Pv) SetPvVolumeModeNil(b bool)`

 SetPvVolumeModeNil sets the value for PvVolumeMode to be an explicit nil

### UnsetPvVolumeMode
`func (o *Pv) UnsetPvVolumeMode()`

UnsetPvVolumeMode ensures that no value is present for PvVolumeMode, not even an explicit nil
### GetStorageClassName

`func (o *Pv) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *Pv) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *Pv) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.


### SetStorageClassNameNil

`func (o *Pv) SetStorageClassNameNil(b bool)`

 SetStorageClassNameNil sets the value for StorageClassName to be an explicit nil

### UnsetStorageClassName
`func (o *Pv) UnsetStorageClassName()`

UnsetStorageClassName ensures that no value is present for StorageClassName, not even an explicit nil
### GetUid

`func (o *Pv) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Pv) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Pv) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Pv) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Pv) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Pv) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


