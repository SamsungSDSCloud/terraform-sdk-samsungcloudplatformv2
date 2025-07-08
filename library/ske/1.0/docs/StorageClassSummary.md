# StorageClassSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **NullableString** |  | 
**AllowVolumeExpansion** | **NullableString** |  | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **NullableTime** |  | 
**DefaultClass** | **NullableString** |  | 
**Name** | **string** | Storage Class Name | 
**Provisioner** | **NullableString** |  | 
**PvReclaimPolicy** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 
**VolumeBindingMode** | **NullableString** |  | 

## Methods

### NewStorageClassSummary

`func NewStorageClassSummary(age NullableString, allowVolumeExpansion NullableString, clusterId string, createdAt NullableTime, defaultClass NullableString, name string, provisioner NullableString, pvReclaimPolicy NullableString, uid string, volumeBindingMode NullableString, ) *StorageClassSummary`

NewStorageClassSummary instantiates a new StorageClassSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassSummaryWithDefaults

`func NewStorageClassSummaryWithDefaults() *StorageClassSummary`

NewStorageClassSummaryWithDefaults instantiates a new StorageClassSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *StorageClassSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *StorageClassSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *StorageClassSummary) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *StorageClassSummary) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *StorageClassSummary) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAllowVolumeExpansion

`func (o *StorageClassSummary) GetAllowVolumeExpansion() string`

GetAllowVolumeExpansion returns the AllowVolumeExpansion field if non-nil, zero value otherwise.

### GetAllowVolumeExpansionOk

`func (o *StorageClassSummary) GetAllowVolumeExpansionOk() (*string, bool)`

GetAllowVolumeExpansionOk returns a tuple with the AllowVolumeExpansion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVolumeExpansion

`func (o *StorageClassSummary) SetAllowVolumeExpansion(v string)`

SetAllowVolumeExpansion sets AllowVolumeExpansion field to given value.


### SetAllowVolumeExpansionNil

`func (o *StorageClassSummary) SetAllowVolumeExpansionNil(b bool)`

 SetAllowVolumeExpansionNil sets the value for AllowVolumeExpansion to be an explicit nil

### UnsetAllowVolumeExpansion
`func (o *StorageClassSummary) UnsetAllowVolumeExpansion()`

UnsetAllowVolumeExpansion ensures that no value is present for AllowVolumeExpansion, not even an explicit nil
### GetClusterId

`func (o *StorageClassSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StorageClassSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StorageClassSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *StorageClassSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StorageClassSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StorageClassSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *StorageClassSummary) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *StorageClassSummary) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDefaultClass

`func (o *StorageClassSummary) GetDefaultClass() string`

GetDefaultClass returns the DefaultClass field if non-nil, zero value otherwise.

### GetDefaultClassOk

`func (o *StorageClassSummary) GetDefaultClassOk() (*string, bool)`

GetDefaultClassOk returns a tuple with the DefaultClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClass

`func (o *StorageClassSummary) SetDefaultClass(v string)`

SetDefaultClass sets DefaultClass field to given value.


### SetDefaultClassNil

`func (o *StorageClassSummary) SetDefaultClassNil(b bool)`

 SetDefaultClassNil sets the value for DefaultClass to be an explicit nil

### UnsetDefaultClass
`func (o *StorageClassSummary) UnsetDefaultClass()`

UnsetDefaultClass ensures that no value is present for DefaultClass, not even an explicit nil
### GetName

`func (o *StorageClassSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageClassSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageClassSummary) SetName(v string)`

SetName sets Name field to given value.


### GetProvisioner

`func (o *StorageClassSummary) GetProvisioner() string`

GetProvisioner returns the Provisioner field if non-nil, zero value otherwise.

### GetProvisionerOk

`func (o *StorageClassSummary) GetProvisionerOk() (*string, bool)`

GetProvisionerOk returns a tuple with the Provisioner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioner

`func (o *StorageClassSummary) SetProvisioner(v string)`

SetProvisioner sets Provisioner field to given value.


### SetProvisionerNil

`func (o *StorageClassSummary) SetProvisionerNil(b bool)`

 SetProvisionerNil sets the value for Provisioner to be an explicit nil

### UnsetProvisioner
`func (o *StorageClassSummary) UnsetProvisioner()`

UnsetProvisioner ensures that no value is present for Provisioner, not even an explicit nil
### GetPvReclaimPolicy

`func (o *StorageClassSummary) GetPvReclaimPolicy() string`

GetPvReclaimPolicy returns the PvReclaimPolicy field if non-nil, zero value otherwise.

### GetPvReclaimPolicyOk

`func (o *StorageClassSummary) GetPvReclaimPolicyOk() (*string, bool)`

GetPvReclaimPolicyOk returns a tuple with the PvReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvReclaimPolicy

`func (o *StorageClassSummary) SetPvReclaimPolicy(v string)`

SetPvReclaimPolicy sets PvReclaimPolicy field to given value.


### SetPvReclaimPolicyNil

`func (o *StorageClassSummary) SetPvReclaimPolicyNil(b bool)`

 SetPvReclaimPolicyNil sets the value for PvReclaimPolicy to be an explicit nil

### UnsetPvReclaimPolicy
`func (o *StorageClassSummary) UnsetPvReclaimPolicy()`

UnsetPvReclaimPolicy ensures that no value is present for PvReclaimPolicy, not even an explicit nil
### GetUid

`func (o *StorageClassSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *StorageClassSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *StorageClassSummary) SetUid(v string)`

SetUid sets Uid field to given value.


### GetVolumeBindingMode

`func (o *StorageClassSummary) GetVolumeBindingMode() string`

GetVolumeBindingMode returns the VolumeBindingMode field if non-nil, zero value otherwise.

### GetVolumeBindingModeOk

`func (o *StorageClassSummary) GetVolumeBindingModeOk() (*string, bool)`

GetVolumeBindingModeOk returns a tuple with the VolumeBindingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeBindingMode

`func (o *StorageClassSummary) SetVolumeBindingMode(v string)`

SetVolumeBindingMode sets VolumeBindingMode field to given value.


### SetVolumeBindingModeNil

`func (o *StorageClassSummary) SetVolumeBindingModeNil(b bool)`

 SetVolumeBindingModeNil sets the value for VolumeBindingMode to be an explicit nil

### UnsetVolumeBindingMode
`func (o *StorageClassSummary) UnsetVolumeBindingMode()`

UnsetVolumeBindingMode ensures that no value is present for VolumeBindingMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


