# EventStreamsInitConfigOptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AkhqId** | Pointer to **NullableString** |  | [optional] 
**AkhqPassword** | Pointer to **NullableString** |  | [optional] 
**BrokerPort** | Pointer to **int32** | broker port | [optional] [default to 9091]
**BrokerSaslId** | **string** | broker sasl id | 
**BrokerSaslPassword** | **string** | broker sasl password | 
**ZookeeperPort** | Pointer to **int32** | zookeeper port | [optional] [default to 2180]
**ZookeeperSaslId** | **string** | zookeeper sasl id | 
**ZookeeperSaslPassword** | **string** | zookeeper sasl password | 

## Methods

### NewEventStreamsInitConfigOptionRequest

`func NewEventStreamsInitConfigOptionRequest(brokerSaslId string, brokerSaslPassword string, zookeeperSaslId string, zookeeperSaslPassword string, ) *EventStreamsInitConfigOptionRequest`

NewEventStreamsInitConfigOptionRequest instantiates a new EventStreamsInitConfigOptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventStreamsInitConfigOptionRequestWithDefaults

`func NewEventStreamsInitConfigOptionRequestWithDefaults() *EventStreamsInitConfigOptionRequest`

NewEventStreamsInitConfigOptionRequestWithDefaults instantiates a new EventStreamsInitConfigOptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAkhqId

`func (o *EventStreamsInitConfigOptionRequest) GetAkhqId() string`

GetAkhqId returns the AkhqId field if non-nil, zero value otherwise.

### GetAkhqIdOk

`func (o *EventStreamsInitConfigOptionRequest) GetAkhqIdOk() (*string, bool)`

GetAkhqIdOk returns a tuple with the AkhqId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkhqId

`func (o *EventStreamsInitConfigOptionRequest) SetAkhqId(v string)`

SetAkhqId sets AkhqId field to given value.

### HasAkhqId

`func (o *EventStreamsInitConfigOptionRequest) HasAkhqId() bool`

HasAkhqId returns a boolean if a field has been set.

### SetAkhqIdNil

`func (o *EventStreamsInitConfigOptionRequest) SetAkhqIdNil(b bool)`

 SetAkhqIdNil sets the value for AkhqId to be an explicit nil

### UnsetAkhqId
`func (o *EventStreamsInitConfigOptionRequest) UnsetAkhqId()`

UnsetAkhqId ensures that no value is present for AkhqId, not even an explicit nil
### GetAkhqPassword

`func (o *EventStreamsInitConfigOptionRequest) GetAkhqPassword() string`

GetAkhqPassword returns the AkhqPassword field if non-nil, zero value otherwise.

### GetAkhqPasswordOk

`func (o *EventStreamsInitConfigOptionRequest) GetAkhqPasswordOk() (*string, bool)`

GetAkhqPasswordOk returns a tuple with the AkhqPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAkhqPassword

`func (o *EventStreamsInitConfigOptionRequest) SetAkhqPassword(v string)`

SetAkhqPassword sets AkhqPassword field to given value.

### HasAkhqPassword

`func (o *EventStreamsInitConfigOptionRequest) HasAkhqPassword() bool`

HasAkhqPassword returns a boolean if a field has been set.

### SetAkhqPasswordNil

`func (o *EventStreamsInitConfigOptionRequest) SetAkhqPasswordNil(b bool)`

 SetAkhqPasswordNil sets the value for AkhqPassword to be an explicit nil

### UnsetAkhqPassword
`func (o *EventStreamsInitConfigOptionRequest) UnsetAkhqPassword()`

UnsetAkhqPassword ensures that no value is present for AkhqPassword, not even an explicit nil
### GetBrokerPort

`func (o *EventStreamsInitConfigOptionRequest) GetBrokerPort() int32`

GetBrokerPort returns the BrokerPort field if non-nil, zero value otherwise.

### GetBrokerPortOk

`func (o *EventStreamsInitConfigOptionRequest) GetBrokerPortOk() (*int32, bool)`

GetBrokerPortOk returns a tuple with the BrokerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerPort

`func (o *EventStreamsInitConfigOptionRequest) SetBrokerPort(v int32)`

SetBrokerPort sets BrokerPort field to given value.

### HasBrokerPort

`func (o *EventStreamsInitConfigOptionRequest) HasBrokerPort() bool`

HasBrokerPort returns a boolean if a field has been set.

### GetBrokerSaslId

`func (o *EventStreamsInitConfigOptionRequest) GetBrokerSaslId() string`

GetBrokerSaslId returns the BrokerSaslId field if non-nil, zero value otherwise.

### GetBrokerSaslIdOk

`func (o *EventStreamsInitConfigOptionRequest) GetBrokerSaslIdOk() (*string, bool)`

GetBrokerSaslIdOk returns a tuple with the BrokerSaslId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerSaslId

`func (o *EventStreamsInitConfigOptionRequest) SetBrokerSaslId(v string)`

SetBrokerSaslId sets BrokerSaslId field to given value.


### GetBrokerSaslPassword

`func (o *EventStreamsInitConfigOptionRequest) GetBrokerSaslPassword() string`

GetBrokerSaslPassword returns the BrokerSaslPassword field if non-nil, zero value otherwise.

### GetBrokerSaslPasswordOk

`func (o *EventStreamsInitConfigOptionRequest) GetBrokerSaslPasswordOk() (*string, bool)`

GetBrokerSaslPasswordOk returns a tuple with the BrokerSaslPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerSaslPassword

`func (o *EventStreamsInitConfigOptionRequest) SetBrokerSaslPassword(v string)`

SetBrokerSaslPassword sets BrokerSaslPassword field to given value.


### GetZookeeperPort

`func (o *EventStreamsInitConfigOptionRequest) GetZookeeperPort() int32`

GetZookeeperPort returns the ZookeeperPort field if non-nil, zero value otherwise.

### GetZookeeperPortOk

`func (o *EventStreamsInitConfigOptionRequest) GetZookeeperPortOk() (*int32, bool)`

GetZookeeperPortOk returns a tuple with the ZookeeperPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZookeeperPort

`func (o *EventStreamsInitConfigOptionRequest) SetZookeeperPort(v int32)`

SetZookeeperPort sets ZookeeperPort field to given value.

### HasZookeeperPort

`func (o *EventStreamsInitConfigOptionRequest) HasZookeeperPort() bool`

HasZookeeperPort returns a boolean if a field has been set.

### GetZookeeperSaslId

`func (o *EventStreamsInitConfigOptionRequest) GetZookeeperSaslId() string`

GetZookeeperSaslId returns the ZookeeperSaslId field if non-nil, zero value otherwise.

### GetZookeeperSaslIdOk

`func (o *EventStreamsInitConfigOptionRequest) GetZookeeperSaslIdOk() (*string, bool)`

GetZookeeperSaslIdOk returns a tuple with the ZookeeperSaslId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZookeeperSaslId

`func (o *EventStreamsInitConfigOptionRequest) SetZookeeperSaslId(v string)`

SetZookeeperSaslId sets ZookeeperSaslId field to given value.


### GetZookeeperSaslPassword

`func (o *EventStreamsInitConfigOptionRequest) GetZookeeperSaslPassword() string`

GetZookeeperSaslPassword returns the ZookeeperSaslPassword field if non-nil, zero value otherwise.

### GetZookeeperSaslPasswordOk

`func (o *EventStreamsInitConfigOptionRequest) GetZookeeperSaslPasswordOk() (*string, bool)`

GetZookeeperSaslPasswordOk returns a tuple with the ZookeeperSaslPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZookeeperSaslPassword

`func (o *EventStreamsInitConfigOptionRequest) SetZookeeperSaslPassword(v string)`

SetZookeeperSaslPassword sets ZookeeperSaslPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


