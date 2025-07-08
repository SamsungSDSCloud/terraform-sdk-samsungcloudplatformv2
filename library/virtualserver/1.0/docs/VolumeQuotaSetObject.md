# VolumeQuotaSetObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InUse** | **int32** | Quota set in use | 
**Limit** | **int32** | Quota set limit | 
**Reserved** | **int32** | Quota set reserved | 

## Methods

### NewVolumeQuotaSetObject

`func NewVolumeQuotaSetObject(inUse int32, limit int32, reserved int32, ) *VolumeQuotaSetObject`

NewVolumeQuotaSetObject instantiates a new VolumeQuotaSetObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeQuotaSetObjectWithDefaults

`func NewVolumeQuotaSetObjectWithDefaults() *VolumeQuotaSetObject`

NewVolumeQuotaSetObjectWithDefaults instantiates a new VolumeQuotaSetObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInUse

`func (o *VolumeQuotaSetObject) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *VolumeQuotaSetObject) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *VolumeQuotaSetObject) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetLimit

`func (o *VolumeQuotaSetObject) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VolumeQuotaSetObject) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VolumeQuotaSetObject) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetReserved

`func (o *VolumeQuotaSetObject) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *VolumeQuotaSetObject) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *VolumeQuotaSetObject) SetReserved(v int32)`

SetReserved sets Reserved field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


