# PvSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** |  | 
**Age** | **NullableString** |  | 
**ClaimRefName** | **NullableString** |  | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **NullableTime** |  | 
**Name** | **string** | PV Name | 
**PvReason** | **NullableString** |  | 
**PvReclaimPolicy** | **NullableString** |  | 
**PvSize** | **NullableString** |  | 
**PvStatus** | **NullableString** |  | 
**StorageClassName** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewPvSummary

`func NewPvSummary(accessModes []string, age NullableString, claimRefName NullableString, clusterId string, createdAt NullableTime, name string, pvReason NullableString, pvReclaimPolicy NullableString, pvSize NullableString, pvStatus NullableString, storageClassName NullableString, uid string, ) *PvSummary`

NewPvSummary instantiates a new PvSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvSummaryWithDefaults

`func NewPvSummaryWithDefaults() *PvSummary`

NewPvSummaryWithDefaults instantiates a new PvSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *PvSummary) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *PvSummary) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *PvSummary) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.


### SetAccessModesNil

`func (o *PvSummary) SetAccessModesNil(b bool)`

 SetAccessModesNil sets the value for AccessModes to be an explicit nil

### UnsetAccessModes
`func (o *PvSummary) UnsetAccessModes()`

UnsetAccessModes ensures that no value is present for AccessModes, not even an explicit nil
### GetAge

`func (o *PvSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PvSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PvSummary) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *PvSummary) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *PvSummary) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetClaimRefName

`func (o *PvSummary) GetClaimRefName() string`

GetClaimRefName returns the ClaimRefName field if non-nil, zero value otherwise.

### GetClaimRefNameOk

`func (o *PvSummary) GetClaimRefNameOk() (*string, bool)`

GetClaimRefNameOk returns a tuple with the ClaimRefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimRefName

`func (o *PvSummary) SetClaimRefName(v string)`

SetClaimRefName sets ClaimRefName field to given value.


### SetClaimRefNameNil

`func (o *PvSummary) SetClaimRefNameNil(b bool)`

 SetClaimRefNameNil sets the value for ClaimRefName to be an explicit nil

### UnsetClaimRefName
`func (o *PvSummary) UnsetClaimRefName()`

UnsetClaimRefName ensures that no value is present for ClaimRefName, not even an explicit nil
### GetClusterId

`func (o *PvSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PvSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PvSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *PvSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PvSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PvSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *PvSummary) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PvSummary) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetName

`func (o *PvSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PvSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PvSummary) SetName(v string)`

SetName sets Name field to given value.


### GetPvReason

`func (o *PvSummary) GetPvReason() string`

GetPvReason returns the PvReason field if non-nil, zero value otherwise.

### GetPvReasonOk

`func (o *PvSummary) GetPvReasonOk() (*string, bool)`

GetPvReasonOk returns a tuple with the PvReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvReason

`func (o *PvSummary) SetPvReason(v string)`

SetPvReason sets PvReason field to given value.


### SetPvReasonNil

`func (o *PvSummary) SetPvReasonNil(b bool)`

 SetPvReasonNil sets the value for PvReason to be an explicit nil

### UnsetPvReason
`func (o *PvSummary) UnsetPvReason()`

UnsetPvReason ensures that no value is present for PvReason, not even an explicit nil
### GetPvReclaimPolicy

`func (o *PvSummary) GetPvReclaimPolicy() string`

GetPvReclaimPolicy returns the PvReclaimPolicy field if non-nil, zero value otherwise.

### GetPvReclaimPolicyOk

`func (o *PvSummary) GetPvReclaimPolicyOk() (*string, bool)`

GetPvReclaimPolicyOk returns a tuple with the PvReclaimPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvReclaimPolicy

`func (o *PvSummary) SetPvReclaimPolicy(v string)`

SetPvReclaimPolicy sets PvReclaimPolicy field to given value.


### SetPvReclaimPolicyNil

`func (o *PvSummary) SetPvReclaimPolicyNil(b bool)`

 SetPvReclaimPolicyNil sets the value for PvReclaimPolicy to be an explicit nil

### UnsetPvReclaimPolicy
`func (o *PvSummary) UnsetPvReclaimPolicy()`

UnsetPvReclaimPolicy ensures that no value is present for PvReclaimPolicy, not even an explicit nil
### GetPvSize

`func (o *PvSummary) GetPvSize() string`

GetPvSize returns the PvSize field if non-nil, zero value otherwise.

### GetPvSizeOk

`func (o *PvSummary) GetPvSizeOk() (*string, bool)`

GetPvSizeOk returns a tuple with the PvSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvSize

`func (o *PvSummary) SetPvSize(v string)`

SetPvSize sets PvSize field to given value.


### SetPvSizeNil

`func (o *PvSummary) SetPvSizeNil(b bool)`

 SetPvSizeNil sets the value for PvSize to be an explicit nil

### UnsetPvSize
`func (o *PvSummary) UnsetPvSize()`

UnsetPvSize ensures that no value is present for PvSize, not even an explicit nil
### GetPvStatus

`func (o *PvSummary) GetPvStatus() string`

GetPvStatus returns the PvStatus field if non-nil, zero value otherwise.

### GetPvStatusOk

`func (o *PvSummary) GetPvStatusOk() (*string, bool)`

GetPvStatusOk returns a tuple with the PvStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvStatus

`func (o *PvSummary) SetPvStatus(v string)`

SetPvStatus sets PvStatus field to given value.


### SetPvStatusNil

`func (o *PvSummary) SetPvStatusNil(b bool)`

 SetPvStatusNil sets the value for PvStatus to be an explicit nil

### UnsetPvStatus
`func (o *PvSummary) UnsetPvStatus()`

UnsetPvStatus ensures that no value is present for PvStatus, not even an explicit nil
### GetStorageClassName

`func (o *PvSummary) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *PvSummary) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *PvSummary) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.


### SetStorageClassNameNil

`func (o *PvSummary) SetStorageClassNameNil(b bool)`

 SetStorageClassNameNil sets the value for StorageClassName to be an explicit nil

### UnsetStorageClassName
`func (o *PvSummary) UnsetStorageClassName()`

UnsetStorageClassName ensures that no value is present for StorageClassName, not even an explicit nil
### GetUid

`func (o *PvSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PvSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PvSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


