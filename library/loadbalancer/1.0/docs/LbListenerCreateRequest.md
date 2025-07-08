# LbListenerCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listener** | [**ListenerForCreate**](ListenerForCreate.md) |  | 

## Methods

### NewLbListenerCreateRequest

`func NewLbListenerCreateRequest(listener ListenerForCreate, ) *LbListenerCreateRequest`

NewLbListenerCreateRequest instantiates a new LbListenerCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbListenerCreateRequestWithDefaults

`func NewLbListenerCreateRequestWithDefaults() *LbListenerCreateRequest`

NewLbListenerCreateRequestWithDefaults instantiates a new LbListenerCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListener

`func (o *LbListenerCreateRequest) GetListener() ListenerForCreate`

GetListener returns the Listener field if non-nil, zero value otherwise.

### GetListenerOk

`func (o *LbListenerCreateRequest) GetListenerOk() (*ListenerForCreate, bool)`

GetListenerOk returns a tuple with the Listener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListener

`func (o *LbListenerCreateRequest) SetListener(v ListenerForCreate)`

SetListener sets Listener field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


