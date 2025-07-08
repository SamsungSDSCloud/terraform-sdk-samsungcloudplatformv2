# VolumeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Filestorages** | [**[]Volume**](Volume.md) |  | 

## Methods

### NewVolumeListResponse

`func NewVolumeListResponse(count int32, filestorages []Volume, ) *VolumeListResponse`

NewVolumeListResponse instantiates a new VolumeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeListResponseWithDefaults

`func NewVolumeListResponseWithDefaults() *VolumeListResponse`

NewVolumeListResponseWithDefaults instantiates a new VolumeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VolumeListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VolumeListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VolumeListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFilestorages

`func (o *VolumeListResponse) GetFilestorages() []Volume`

GetFilestorages returns the Filestorages field if non-nil, zero value otherwise.

### GetFilestoragesOk

`func (o *VolumeListResponse) GetFilestoragesOk() (*[]Volume, bool)`

GetFilestoragesOk returns a tuple with the Filestorages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilestorages

`func (o *VolumeListResponse) SetFilestorages(v []Volume)`

SetFilestorages sets Filestorages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


