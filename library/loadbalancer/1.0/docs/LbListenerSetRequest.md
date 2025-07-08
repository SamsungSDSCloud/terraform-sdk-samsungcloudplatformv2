# LbListenerSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listener** | [**ListenerForSet**](ListenerForSet.md) |  | 

## Methods

### NewLbListenerSetRequest

`func NewLbListenerSetRequest(listener ListenerForSet, ) *LbListenerSetRequest`

NewLbListenerSetRequest instantiates a new LbListenerSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbListenerSetRequestWithDefaults

`func NewLbListenerSetRequestWithDefaults() *LbListenerSetRequest`

NewLbListenerSetRequestWithDefaults instantiates a new LbListenerSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListener

`func (o *LbListenerSetRequest) GetListener() ListenerForSet`

GetListener returns the Listener field if non-nil, zero value otherwise.

### GetListenerOk

`func (o *LbListenerSetRequest) GetListenerOk() (*ListenerForSet, bool)`

GetListenerOk returns a tuple with the Listener field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListener

`func (o *LbListenerSetRequest) SetListener(v ListenerForSet)`

SetListener sets Listener field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


