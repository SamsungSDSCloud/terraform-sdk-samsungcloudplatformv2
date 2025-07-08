# ServersVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnTermination** | **bool** | Delete on termination | 
**Device** | **string** | Device Name | 
**Id** | **string** | Volume Attachment ID | 
**ServerId** | **string** | Server ID | 
**VolumeId** | **string** | Volume ID | 

## Methods

### NewServersVolumeResponse

`func NewServersVolumeResponse(deleteOnTermination bool, device string, id string, serverId string, volumeId string, ) *ServersVolumeResponse`

NewServersVolumeResponse instantiates a new ServersVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServersVolumeResponseWithDefaults

`func NewServersVolumeResponseWithDefaults() *ServersVolumeResponse`

NewServersVolumeResponseWithDefaults instantiates a new ServersVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnTermination

`func (o *ServersVolumeResponse) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *ServersVolumeResponse) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *ServersVolumeResponse) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.


### GetDevice

`func (o *ServersVolumeResponse) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ServersVolumeResponse) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ServersVolumeResponse) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetId

`func (o *ServersVolumeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServersVolumeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServersVolumeResponse) SetId(v string)`

SetId sets Id field to given value.


### GetServerId

`func (o *ServersVolumeResponse) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ServersVolumeResponse) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ServersVolumeResponse) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### GetVolumeId

`func (o *ServersVolumeResponse) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ServersVolumeResponse) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ServersVolumeResponse) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


