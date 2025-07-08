# ServerShowResponseVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteOnTermination** | **bool** | Delete on termination | 
**Id** | **string** | Volume ID | 

## Methods

### NewServerShowResponseVolume

`func NewServerShowResponseVolume(deleteOnTermination bool, id string, ) *ServerShowResponseVolume`

NewServerShowResponseVolume instantiates a new ServerShowResponseVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerShowResponseVolumeWithDefaults

`func NewServerShowResponseVolumeWithDefaults() *ServerShowResponseVolume`

NewServerShowResponseVolumeWithDefaults instantiates a new ServerShowResponseVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteOnTermination

`func (o *ServerShowResponseVolume) GetDeleteOnTermination() bool`

GetDeleteOnTermination returns the DeleteOnTermination field if non-nil, zero value otherwise.

### GetDeleteOnTerminationOk

`func (o *ServerShowResponseVolume) GetDeleteOnTerminationOk() (*bool, bool)`

GetDeleteOnTerminationOk returns a tuple with the DeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOnTermination

`func (o *ServerShowResponseVolume) SetDeleteOnTermination(v bool)`

SetDeleteOnTermination sets DeleteOnTermination field to given value.


### GetId

`func (o *ServerShowResponseVolume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerShowResponseVolume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerShowResponseVolume) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


