# GslbHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthCheckInterval** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckProbeTimeout** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckUserId** | Pointer to **NullableString** |  | [optional] 
**HealthCheckUserPassword** | Pointer to **NullableString** |  | [optional] 
**Protocol** | **string** | The GSLB Health Check Protocol. | 
**ReceiveString** | Pointer to **NullableString** |  | [optional] 
**SendString** | Pointer to **NullableString** |  | [optional] 
**ServicePort** | Pointer to **NullableInt32** |  | [optional] 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGslbHealthCheck

`func NewGslbHealthCheck(protocol string, ) *GslbHealthCheck`

NewGslbHealthCheck instantiates a new GslbHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGslbHealthCheckWithDefaults

`func NewGslbHealthCheckWithDefaults() *GslbHealthCheck`

NewGslbHealthCheckWithDefaults instantiates a new GslbHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthCheckInterval

`func (o *GslbHealthCheck) GetHealthCheckInterval() int32`

GetHealthCheckInterval returns the HealthCheckInterval field if non-nil, zero value otherwise.

### GetHealthCheckIntervalOk

`func (o *GslbHealthCheck) GetHealthCheckIntervalOk() (*int32, bool)`

GetHealthCheckIntervalOk returns a tuple with the HealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckInterval

`func (o *GslbHealthCheck) SetHealthCheckInterval(v int32)`

SetHealthCheckInterval sets HealthCheckInterval field to given value.

### HasHealthCheckInterval

`func (o *GslbHealthCheck) HasHealthCheckInterval() bool`

HasHealthCheckInterval returns a boolean if a field has been set.

### SetHealthCheckIntervalNil

`func (o *GslbHealthCheck) SetHealthCheckIntervalNil(b bool)`

 SetHealthCheckIntervalNil sets the value for HealthCheckInterval to be an explicit nil

### UnsetHealthCheckInterval
`func (o *GslbHealthCheck) UnsetHealthCheckInterval()`

UnsetHealthCheckInterval ensures that no value is present for HealthCheckInterval, not even an explicit nil
### GetHealthCheckProbeTimeout

`func (o *GslbHealthCheck) GetHealthCheckProbeTimeout() int32`

GetHealthCheckProbeTimeout returns the HealthCheckProbeTimeout field if non-nil, zero value otherwise.

### GetHealthCheckProbeTimeoutOk

`func (o *GslbHealthCheck) GetHealthCheckProbeTimeoutOk() (*int32, bool)`

GetHealthCheckProbeTimeoutOk returns a tuple with the HealthCheckProbeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckProbeTimeout

`func (o *GslbHealthCheck) SetHealthCheckProbeTimeout(v int32)`

SetHealthCheckProbeTimeout sets HealthCheckProbeTimeout field to given value.

### HasHealthCheckProbeTimeout

`func (o *GslbHealthCheck) HasHealthCheckProbeTimeout() bool`

HasHealthCheckProbeTimeout returns a boolean if a field has been set.

### SetHealthCheckProbeTimeoutNil

`func (o *GslbHealthCheck) SetHealthCheckProbeTimeoutNil(b bool)`

 SetHealthCheckProbeTimeoutNil sets the value for HealthCheckProbeTimeout to be an explicit nil

### UnsetHealthCheckProbeTimeout
`func (o *GslbHealthCheck) UnsetHealthCheckProbeTimeout()`

UnsetHealthCheckProbeTimeout ensures that no value is present for HealthCheckProbeTimeout, not even an explicit nil
### GetHealthCheckUserId

`func (o *GslbHealthCheck) GetHealthCheckUserId() string`

GetHealthCheckUserId returns the HealthCheckUserId field if non-nil, zero value otherwise.

### GetHealthCheckUserIdOk

`func (o *GslbHealthCheck) GetHealthCheckUserIdOk() (*string, bool)`

GetHealthCheckUserIdOk returns a tuple with the HealthCheckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUserId

`func (o *GslbHealthCheck) SetHealthCheckUserId(v string)`

SetHealthCheckUserId sets HealthCheckUserId field to given value.

### HasHealthCheckUserId

`func (o *GslbHealthCheck) HasHealthCheckUserId() bool`

HasHealthCheckUserId returns a boolean if a field has been set.

### SetHealthCheckUserIdNil

`func (o *GslbHealthCheck) SetHealthCheckUserIdNil(b bool)`

 SetHealthCheckUserIdNil sets the value for HealthCheckUserId to be an explicit nil

### UnsetHealthCheckUserId
`func (o *GslbHealthCheck) UnsetHealthCheckUserId()`

UnsetHealthCheckUserId ensures that no value is present for HealthCheckUserId, not even an explicit nil
### GetHealthCheckUserPassword

`func (o *GslbHealthCheck) GetHealthCheckUserPassword() string`

GetHealthCheckUserPassword returns the HealthCheckUserPassword field if non-nil, zero value otherwise.

### GetHealthCheckUserPasswordOk

`func (o *GslbHealthCheck) GetHealthCheckUserPasswordOk() (*string, bool)`

GetHealthCheckUserPasswordOk returns a tuple with the HealthCheckUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUserPassword

`func (o *GslbHealthCheck) SetHealthCheckUserPassword(v string)`

SetHealthCheckUserPassword sets HealthCheckUserPassword field to given value.

### HasHealthCheckUserPassword

`func (o *GslbHealthCheck) HasHealthCheckUserPassword() bool`

HasHealthCheckUserPassword returns a boolean if a field has been set.

### SetHealthCheckUserPasswordNil

`func (o *GslbHealthCheck) SetHealthCheckUserPasswordNil(b bool)`

 SetHealthCheckUserPasswordNil sets the value for HealthCheckUserPassword to be an explicit nil

### UnsetHealthCheckUserPassword
`func (o *GslbHealthCheck) UnsetHealthCheckUserPassword()`

UnsetHealthCheckUserPassword ensures that no value is present for HealthCheckUserPassword, not even an explicit nil
### GetProtocol

`func (o *GslbHealthCheck) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GslbHealthCheck) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GslbHealthCheck) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetReceiveString

`func (o *GslbHealthCheck) GetReceiveString() string`

GetReceiveString returns the ReceiveString field if non-nil, zero value otherwise.

### GetReceiveStringOk

`func (o *GslbHealthCheck) GetReceiveStringOk() (*string, bool)`

GetReceiveStringOk returns a tuple with the ReceiveString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveString

`func (o *GslbHealthCheck) SetReceiveString(v string)`

SetReceiveString sets ReceiveString field to given value.

### HasReceiveString

`func (o *GslbHealthCheck) HasReceiveString() bool`

HasReceiveString returns a boolean if a field has been set.

### SetReceiveStringNil

`func (o *GslbHealthCheck) SetReceiveStringNil(b bool)`

 SetReceiveStringNil sets the value for ReceiveString to be an explicit nil

### UnsetReceiveString
`func (o *GslbHealthCheck) UnsetReceiveString()`

UnsetReceiveString ensures that no value is present for ReceiveString, not even an explicit nil
### GetSendString

`func (o *GslbHealthCheck) GetSendString() string`

GetSendString returns the SendString field if non-nil, zero value otherwise.

### GetSendStringOk

`func (o *GslbHealthCheck) GetSendStringOk() (*string, bool)`

GetSendStringOk returns a tuple with the SendString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendString

`func (o *GslbHealthCheck) SetSendString(v string)`

SetSendString sets SendString field to given value.

### HasSendString

`func (o *GslbHealthCheck) HasSendString() bool`

HasSendString returns a boolean if a field has been set.

### SetSendStringNil

`func (o *GslbHealthCheck) SetSendStringNil(b bool)`

 SetSendStringNil sets the value for SendString to be an explicit nil

### UnsetSendString
`func (o *GslbHealthCheck) UnsetSendString()`

UnsetSendString ensures that no value is present for SendString, not even an explicit nil
### GetServicePort

`func (o *GslbHealthCheck) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *GslbHealthCheck) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *GslbHealthCheck) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *GslbHealthCheck) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *GslbHealthCheck) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *GslbHealthCheck) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetTimeout

`func (o *GslbHealthCheck) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GslbHealthCheck) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GslbHealthCheck) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GslbHealthCheck) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *GslbHealthCheck) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *GslbHealthCheck) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


