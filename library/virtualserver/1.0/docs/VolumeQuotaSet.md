# VolumeQuotaSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Quota set id | 
**PerVolumeSize** | [**VolumeQuotaSetObject**](VolumeQuotaSetObject.md) | Per volume size quota | 
**Snapshots** | [**VolumeQuotaSetObject**](VolumeQuotaSetObject.md) | Snapshot quota | 
**Usages** | [**VolumeQuotaSetObject**](VolumeQuotaSetObject.md) | Usage quota | 
**Volumes** | [**VolumeQuotaSetObject**](VolumeQuotaSetObject.md) | Volume quota | 

## Methods

### NewVolumeQuotaSet

`func NewVolumeQuotaSet(id string, perVolumeSize VolumeQuotaSetObject, snapshots VolumeQuotaSetObject, usages VolumeQuotaSetObject, volumes VolumeQuotaSetObject, ) *VolumeQuotaSet`

NewVolumeQuotaSet instantiates a new VolumeQuotaSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeQuotaSetWithDefaults

`func NewVolumeQuotaSetWithDefaults() *VolumeQuotaSet`

NewVolumeQuotaSetWithDefaults instantiates a new VolumeQuotaSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumeQuotaSet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeQuotaSet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeQuotaSet) SetId(v string)`

SetId sets Id field to given value.


### GetPerVolumeSize

`func (o *VolumeQuotaSet) GetPerVolumeSize() VolumeQuotaSetObject`

GetPerVolumeSize returns the PerVolumeSize field if non-nil, zero value otherwise.

### GetPerVolumeSizeOk

`func (o *VolumeQuotaSet) GetPerVolumeSizeOk() (*VolumeQuotaSetObject, bool)`

GetPerVolumeSizeOk returns a tuple with the PerVolumeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerVolumeSize

`func (o *VolumeQuotaSet) SetPerVolumeSize(v VolumeQuotaSetObject)`

SetPerVolumeSize sets PerVolumeSize field to given value.


### GetSnapshots

`func (o *VolumeQuotaSet) GetSnapshots() VolumeQuotaSetObject`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *VolumeQuotaSet) GetSnapshotsOk() (*VolumeQuotaSetObject, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *VolumeQuotaSet) SetSnapshots(v VolumeQuotaSetObject)`

SetSnapshots sets Snapshots field to given value.


### GetUsages

`func (o *VolumeQuotaSet) GetUsages() VolumeQuotaSetObject`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *VolumeQuotaSet) GetUsagesOk() (*VolumeQuotaSetObject, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *VolumeQuotaSet) SetUsages(v VolumeQuotaSetObject)`

SetUsages sets Usages field to given value.


### GetVolumes

`func (o *VolumeQuotaSet) GetVolumes() VolumeQuotaSetObject`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *VolumeQuotaSet) GetVolumesOk() (*VolumeQuotaSetObject, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *VolumeQuotaSet) SetVolumes(v VolumeQuotaSetObject)`

SetVolumes sets Volumes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


