# VolumeSnapshotRateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsSnapshotActivated** | Pointer to **bool** | Whether to activate snapshot | [optional] 
**SnapshotRate** | Pointer to **int32** | Snapshot rate | [optional] 
**VolumeId** | Pointer to **string** | Volume id | [optional] 

## Methods

### NewVolumeSnapshotRateResponse

`func NewVolumeSnapshotRateResponse() *VolumeSnapshotRateResponse`

NewVolumeSnapshotRateResponse instantiates a new VolumeSnapshotRateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeSnapshotRateResponseWithDefaults

`func NewVolumeSnapshotRateResponseWithDefaults() *VolumeSnapshotRateResponse`

NewVolumeSnapshotRateResponseWithDefaults instantiates a new VolumeSnapshotRateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsSnapshotActivated

`func (o *VolumeSnapshotRateResponse) GetIsSnapshotActivated() bool`

GetIsSnapshotActivated returns the IsSnapshotActivated field if non-nil, zero value otherwise.

### GetIsSnapshotActivatedOk

`func (o *VolumeSnapshotRateResponse) GetIsSnapshotActivatedOk() (*bool, bool)`

GetIsSnapshotActivatedOk returns a tuple with the IsSnapshotActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSnapshotActivated

`func (o *VolumeSnapshotRateResponse) SetIsSnapshotActivated(v bool)`

SetIsSnapshotActivated sets IsSnapshotActivated field to given value.

### HasIsSnapshotActivated

`func (o *VolumeSnapshotRateResponse) HasIsSnapshotActivated() bool`

HasIsSnapshotActivated returns a boolean if a field has been set.

### GetSnapshotRate

`func (o *VolumeSnapshotRateResponse) GetSnapshotRate() int32`

GetSnapshotRate returns the SnapshotRate field if non-nil, zero value otherwise.

### GetSnapshotRateOk

`func (o *VolumeSnapshotRateResponse) GetSnapshotRateOk() (*int32, bool)`

GetSnapshotRateOk returns a tuple with the SnapshotRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotRate

`func (o *VolumeSnapshotRateResponse) SetSnapshotRate(v int32)`

SetSnapshotRate sets SnapshotRate field to given value.

### HasSnapshotRate

`func (o *VolumeSnapshotRateResponse) HasSnapshotRate() bool`

HasSnapshotRate returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeSnapshotRateResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeSnapshotRateResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeSnapshotRateResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeSnapshotRateResponse) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


