# VolumeServerAttachRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnTermination** | Pointer to **NullableBool** |  | [optional] 
**ServerId** | **string** | Server ID | 

## Methods

### NewVolumeServerAttachRequest

`func NewVolumeServerAttachRequest(serverId string, ) *VolumeServerAttachRequest`

NewVolumeServerAttachRequest instantiates a new VolumeServerAttachRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeServerAttachRequestWithDefaults

`func NewVolumeServerAttachRequestWithDefaults() *VolumeServerAttachRequest`

NewVolumeServerAttachRequestWithDefaults instantiates a new VolumeServerAttachRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnTermination

`func (o *VolumeServerAttachRequest) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *VolumeServerAttachRequest) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *VolumeServerAttachRequest) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.

### HasDeleteOnTermination

`func (o *VolumeServerAttachRequest) HasDeleteOnTermination() bool`

HasDeleteOnTermination returns a boolean if a field has been set.

### SetDeleteOnTerminationNil

`func (o *VolumeServerAttachRequest) SetDeleteOnTerminationNil(b bool)`

 SetDeleteOnTerminationNil sets the value for DeleteOnTermination to be an explicit nil

### UnsetDeleteOnTermination
`func (o *VolumeServerAttachRequest) UnsetDeleteOnTermination()`

UnsetDeleteOnTermination ensures that no value is present for DeleteOnTermination, not even an explicit nil
### GetServerId

`func (o *VolumeServerAttachRequest) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *VolumeServerAttachRequest) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *VolumeServerAttachRequest) SetServerId(v string)`

SetServerId sets ServerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


