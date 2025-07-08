# PvcSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessModes** | **[]string** |  | 
**Age** | **NullableString** |  | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **NullableTime** |  | 
**Name** | **string** | PVC Name | 
**NamespaceName** | **string** | Namespace Name | 
**PvName** | **NullableString** |  | 
**PvcSize** | **NullableString** |  | 
**PvcStatus** | **NullableString** |  | 
**StorageClassName** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewPvcSummary

`func NewPvcSummary(accessModes []string, age NullableString, clusterId string, createdAt NullableTime, name string, namespaceName string, pvName NullableString, pvcSize NullableString, pvcStatus NullableString, storageClassName NullableString, uid string, ) *PvcSummary`

NewPvcSummary instantiates a new PvcSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPvcSummaryWithDefaults

`func NewPvcSummaryWithDefaults() *PvcSummary`

NewPvcSummaryWithDefaults instantiates a new PvcSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessModes

`func (o *PvcSummary) GetAccessModes() []string`

GetAccessModes returns the AccessModes field if non-nil, zero value otherwise.

### GetAccessModesOk

`func (o *PvcSummary) GetAccessModesOk() (*[]string, bool)`

GetAccessModesOk returns a tuple with the AccessModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessModes

`func (o *PvcSummary) SetAccessModes(v []string)`

SetAccessModes sets AccessModes field to given value.


### SetAccessModesNil

`func (o *PvcSummary) SetAccessModesNil(b bool)`

 SetAccessModesNil sets the value for AccessModes to be an explicit nil

### UnsetAccessModes
`func (o *PvcSummary) UnsetAccessModes()`

UnsetAccessModes ensures that no value is present for AccessModes, not even an explicit nil
### GetAge

`func (o *PvcSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *PvcSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *PvcSummary) SetAge(v string)`

SetAge sets Age field to given value.


### SetAgeNil

`func (o *PvcSummary) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *PvcSummary) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetClusterId

`func (o *PvcSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *PvcSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *PvcSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *PvcSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PvcSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PvcSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *PvcSummary) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PvcSummary) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetName

`func (o *PvcSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PvcSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PvcSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *PvcSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *PvcSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *PvcSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetPvName

`func (o *PvcSummary) GetPvName() string`

GetPvName returns the PvName field if non-nil, zero value otherwise.

### GetPvNameOk

`func (o *PvcSummary) GetPvNameOk() (*string, bool)`

GetPvNameOk returns a tuple with the PvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvName

`func (o *PvcSummary) SetPvName(v string)`

SetPvName sets PvName field to given value.


### SetPvNameNil

`func (o *PvcSummary) SetPvNameNil(b bool)`

 SetPvNameNil sets the value for PvName to be an explicit nil

### UnsetPvName
`func (o *PvcSummary) UnsetPvName()`

UnsetPvName ensures that no value is present for PvName, not even an explicit nil
### GetPvcSize

`func (o *PvcSummary) GetPvcSize() string`

GetPvcSize returns the PvcSize field if non-nil, zero value otherwise.

### GetPvcSizeOk

`func (o *PvcSummary) GetPvcSizeOk() (*string, bool)`

GetPvcSizeOk returns a tuple with the PvcSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcSize

`func (o *PvcSummary) SetPvcSize(v string)`

SetPvcSize sets PvcSize field to given value.


### SetPvcSizeNil

`func (o *PvcSummary) SetPvcSizeNil(b bool)`

 SetPvcSizeNil sets the value for PvcSize to be an explicit nil

### UnsetPvcSize
`func (o *PvcSummary) UnsetPvcSize()`

UnsetPvcSize ensures that no value is present for PvcSize, not even an explicit nil
### GetPvcStatus

`func (o *PvcSummary) GetPvcStatus() string`

GetPvcStatus returns the PvcStatus field if non-nil, zero value otherwise.

### GetPvcStatusOk

`func (o *PvcSummary) GetPvcStatusOk() (*string, bool)`

GetPvcStatusOk returns a tuple with the PvcStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvcStatus

`func (o *PvcSummary) SetPvcStatus(v string)`

SetPvcStatus sets PvcStatus field to given value.


### SetPvcStatusNil

`func (o *PvcSummary) SetPvcStatusNil(b bool)`

 SetPvcStatusNil sets the value for PvcStatus to be an explicit nil

### UnsetPvcStatus
`func (o *PvcSummary) UnsetPvcStatus()`

UnsetPvcStatus ensures that no value is present for PvcStatus, not even an explicit nil
### GetStorageClassName

`func (o *PvcSummary) GetStorageClassName() string`

GetStorageClassName returns the StorageClassName field if non-nil, zero value otherwise.

### GetStorageClassNameOk

`func (o *PvcSummary) GetStorageClassNameOk() (*string, bool)`

GetStorageClassNameOk returns a tuple with the StorageClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassName

`func (o *PvcSummary) SetStorageClassName(v string)`

SetStorageClassName sets StorageClassName field to given value.


### SetStorageClassNameNil

`func (o *PvcSummary) SetStorageClassNameNil(b bool)`

 SetStorageClassNameNil sets the value for StorageClassName to be an explicit nil

### UnsetStorageClassName
`func (o *PvcSummary) UnsetStorageClassName()`

UnsetStorageClassName ensures that no value is present for StorageClassName, not even an explicit nil
### GetUid

`func (o *PvcSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *PvcSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *PvcSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


