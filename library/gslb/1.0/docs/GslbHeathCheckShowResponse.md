# GslbHeathCheckShowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**HealthCheckInterval** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckProbeTimeout** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckUserId** | Pointer to **NullableString** |  | [optional] 
**HealthCheckUserPassword** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Protocol** | **string** | The GSLB Health Check Protocol. | 
**ReceiveString** | Pointer to **NullableString** |  | [optional] 
**SendString** | Pointer to **NullableString** |  | [optional] 
**ServicePort** | Pointer to **NullableInt32** |  | [optional] 
**Timeout** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewGslbHeathCheckShowResponse

`func NewGslbHeathCheckShowResponse(createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, protocol string, ) *GslbHeathCheckShowResponse`

NewGslbHeathCheckShowResponse instantiates a new GslbHeathCheckShowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGslbHeathCheckShowResponseWithDefaults

`func NewGslbHeathCheckShowResponseWithDefaults() *GslbHeathCheckShowResponse`

NewGslbHeathCheckShowResponseWithDefaults instantiates a new GslbHeathCheckShowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *GslbHeathCheckShowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GslbHeathCheckShowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GslbHeathCheckShowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *GslbHeathCheckShowResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GslbHeathCheckShowResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GslbHeathCheckShowResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetHealthCheckInterval

`func (o *GslbHeathCheckShowResponse) GetHealthCheckInterval() int32`

GetHealthCheckInterval returns the HealthCheckInterval field if non-nil, zero value otherwise.

### GetHealthCheckIntervalOk

`func (o *GslbHeathCheckShowResponse) GetHealthCheckIntervalOk() (*int32, bool)`

GetHealthCheckIntervalOk returns a tuple with the HealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckInterval

`func (o *GslbHeathCheckShowResponse) SetHealthCheckInterval(v int32)`

SetHealthCheckInterval sets HealthCheckInterval field to given value.

### HasHealthCheckInterval

`func (o *GslbHeathCheckShowResponse) HasHealthCheckInterval() bool`

HasHealthCheckInterval returns a boolean if a field has been set.

### SetHealthCheckIntervalNil

`func (o *GslbHeathCheckShowResponse) SetHealthCheckIntervalNil(b bool)`

 SetHealthCheckIntervalNil sets the value for HealthCheckInterval to be an explicit nil

### UnsetHealthCheckInterval
`func (o *GslbHeathCheckShowResponse) UnsetHealthCheckInterval()`

UnsetHealthCheckInterval ensures that no value is present for HealthCheckInterval, not even an explicit nil
### GetHealthCheckProbeTimeout

`func (o *GslbHeathCheckShowResponse) GetHealthCheckProbeTimeout() int32`

GetHealthCheckProbeTimeout returns the HealthCheckProbeTimeout field if non-nil, zero value otherwise.

### GetHealthCheckProbeTimeoutOk

`func (o *GslbHeathCheckShowResponse) GetHealthCheckProbeTimeoutOk() (*int32, bool)`

GetHealthCheckProbeTimeoutOk returns a tuple with the HealthCheckProbeTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckProbeTimeout

`func (o *GslbHeathCheckShowResponse) SetHealthCheckProbeTimeout(v int32)`

SetHealthCheckProbeTimeout sets HealthCheckProbeTimeout field to given value.

### HasHealthCheckProbeTimeout

`func (o *GslbHeathCheckShowResponse) HasHealthCheckProbeTimeout() bool`

HasHealthCheckProbeTimeout returns a boolean if a field has been set.

### SetHealthCheckProbeTimeoutNil

`func (o *GslbHeathCheckShowResponse) SetHealthCheckProbeTimeoutNil(b bool)`

 SetHealthCheckProbeTimeoutNil sets the value for HealthCheckProbeTimeout to be an explicit nil

### UnsetHealthCheckProbeTimeout
`func (o *GslbHeathCheckShowResponse) UnsetHealthCheckProbeTimeout()`

UnsetHealthCheckProbeTimeout ensures that no value is present for HealthCheckProbeTimeout, not even an explicit nil
### GetHealthCheckUserId

`func (o *GslbHeathCheckShowResponse) GetHealthCheckUserId() string`

GetHealthCheckUserId returns the HealthCheckUserId field if non-nil, zero value otherwise.

### GetHealthCheckUserIdOk

`func (o *GslbHeathCheckShowResponse) GetHealthCheckUserIdOk() (*string, bool)`

GetHealthCheckUserIdOk returns a tuple with the HealthCheckUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUserId

`func (o *GslbHeathCheckShowResponse) SetHealthCheckUserId(v string)`

SetHealthCheckUserId sets HealthCheckUserId field to given value.

### HasHealthCheckUserId

`func (o *GslbHeathCheckShowResponse) HasHealthCheckUserId() bool`

HasHealthCheckUserId returns a boolean if a field has been set.

### SetHealthCheckUserIdNil

`func (o *GslbHeathCheckShowResponse) SetHealthCheckUserIdNil(b bool)`

 SetHealthCheckUserIdNil sets the value for HealthCheckUserId to be an explicit nil

### UnsetHealthCheckUserId
`func (o *GslbHeathCheckShowResponse) UnsetHealthCheckUserId()`

UnsetHealthCheckUserId ensures that no value is present for HealthCheckUserId, not even an explicit nil
### GetHealthCheckUserPassword

`func (o *GslbHeathCheckShowResponse) GetHealthCheckUserPassword() string`

GetHealthCheckUserPassword returns the HealthCheckUserPassword field if non-nil, zero value otherwise.

### GetHealthCheckUserPasswordOk

`func (o *GslbHeathCheckShowResponse) GetHealthCheckUserPasswordOk() (*string, bool)`

GetHealthCheckUserPasswordOk returns a tuple with the HealthCheckUserPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUserPassword

`func (o *GslbHeathCheckShowResponse) SetHealthCheckUserPassword(v string)`

SetHealthCheckUserPassword sets HealthCheckUserPassword field to given value.

### HasHealthCheckUserPassword

`func (o *GslbHeathCheckShowResponse) HasHealthCheckUserPassword() bool`

HasHealthCheckUserPassword returns a boolean if a field has been set.

### SetHealthCheckUserPasswordNil

`func (o *GslbHeathCheckShowResponse) SetHealthCheckUserPasswordNil(b bool)`

 SetHealthCheckUserPasswordNil sets the value for HealthCheckUserPassword to be an explicit nil

### UnsetHealthCheckUserPassword
`func (o *GslbHeathCheckShowResponse) UnsetHealthCheckUserPassword()`

UnsetHealthCheckUserPassword ensures that no value is present for HealthCheckUserPassword, not even an explicit nil
### GetId

`func (o *GslbHeathCheckShowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GslbHeathCheckShowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GslbHeathCheckShowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *GslbHeathCheckShowResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *GslbHeathCheckShowResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *GslbHeathCheckShowResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *GslbHeathCheckShowResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *GslbHeathCheckShowResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *GslbHeathCheckShowResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetProtocol

`func (o *GslbHeathCheckShowResponse) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GslbHeathCheckShowResponse) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GslbHeathCheckShowResponse) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetReceiveString

`func (o *GslbHeathCheckShowResponse) GetReceiveString() string`

GetReceiveString returns the ReceiveString field if non-nil, zero value otherwise.

### GetReceiveStringOk

`func (o *GslbHeathCheckShowResponse) GetReceiveStringOk() (*string, bool)`

GetReceiveStringOk returns a tuple with the ReceiveString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveString

`func (o *GslbHeathCheckShowResponse) SetReceiveString(v string)`

SetReceiveString sets ReceiveString field to given value.

### HasReceiveString

`func (o *GslbHeathCheckShowResponse) HasReceiveString() bool`

HasReceiveString returns a boolean if a field has been set.

### SetReceiveStringNil

`func (o *GslbHeathCheckShowResponse) SetReceiveStringNil(b bool)`

 SetReceiveStringNil sets the value for ReceiveString to be an explicit nil

### UnsetReceiveString
`func (o *GslbHeathCheckShowResponse) UnsetReceiveString()`

UnsetReceiveString ensures that no value is present for ReceiveString, not even an explicit nil
### GetSendString

`func (o *GslbHeathCheckShowResponse) GetSendString() string`

GetSendString returns the SendString field if non-nil, zero value otherwise.

### GetSendStringOk

`func (o *GslbHeathCheckShowResponse) GetSendStringOk() (*string, bool)`

GetSendStringOk returns a tuple with the SendString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendString

`func (o *GslbHeathCheckShowResponse) SetSendString(v string)`

SetSendString sets SendString field to given value.

### HasSendString

`func (o *GslbHeathCheckShowResponse) HasSendString() bool`

HasSendString returns a boolean if a field has been set.

### SetSendStringNil

`func (o *GslbHeathCheckShowResponse) SetSendStringNil(b bool)`

 SetSendStringNil sets the value for SendString to be an explicit nil

### UnsetSendString
`func (o *GslbHeathCheckShowResponse) UnsetSendString()`

UnsetSendString ensures that no value is present for SendString, not even an explicit nil
### GetServicePort

`func (o *GslbHeathCheckShowResponse) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *GslbHeathCheckShowResponse) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *GslbHeathCheckShowResponse) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *GslbHeathCheckShowResponse) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *GslbHeathCheckShowResponse) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *GslbHeathCheckShowResponse) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetTimeout

`func (o *GslbHeathCheckShowResponse) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GslbHeathCheckShowResponse) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GslbHeathCheckShowResponse) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GslbHeathCheckShowResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *GslbHeathCheckShowResponse) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *GslbHeathCheckShowResponse) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


