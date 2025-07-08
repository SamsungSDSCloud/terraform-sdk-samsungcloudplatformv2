# LbHealthCheckSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**HealthCheckCount** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckInterval** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckPort** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckTimeout** | Pointer to **NullableInt32** |  | [optional] 
**HealthCheckUrl** | Pointer to **NullableString** |  | [optional] 
**HttpMethod** | Pointer to [**NullableLbMonitorHttpMethod**](LbMonitorHttpMethod.md) |  | [optional] 
**Protocol** | Pointer to [**NullableLbMonitorProtocol**](LbMonitorProtocol.md) |  | [optional] 
**RequestData** | Pointer to **NullableString** |  | [optional] 
**ResponseCode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLbHealthCheckSet

`func NewLbHealthCheckSet() *LbHealthCheckSet`

NewLbHealthCheckSet instantiates a new LbHealthCheckSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbHealthCheckSetWithDefaults

`func NewLbHealthCheckSetWithDefaults() *LbHealthCheckSet`

NewLbHealthCheckSetWithDefaults instantiates a new LbHealthCheckSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LbHealthCheckSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LbHealthCheckSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LbHealthCheckSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LbHealthCheckSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LbHealthCheckSet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LbHealthCheckSet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHealthCheckCount

`func (o *LbHealthCheckSet) GetHealthCheckCount() int32`

GetHealthCheckCount returns the HealthCheckCount field if non-nil, zero value otherwise.

### GetHealthCheckCountOk

`func (o *LbHealthCheckSet) GetHealthCheckCountOk() (*int32, bool)`

GetHealthCheckCountOk returns a tuple with the HealthCheckCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckCount

`func (o *LbHealthCheckSet) SetHealthCheckCount(v int32)`

SetHealthCheckCount sets HealthCheckCount field to given value.

### HasHealthCheckCount

`func (o *LbHealthCheckSet) HasHealthCheckCount() bool`

HasHealthCheckCount returns a boolean if a field has been set.

### SetHealthCheckCountNil

`func (o *LbHealthCheckSet) SetHealthCheckCountNil(b bool)`

 SetHealthCheckCountNil sets the value for HealthCheckCount to be an explicit nil

### UnsetHealthCheckCount
`func (o *LbHealthCheckSet) UnsetHealthCheckCount()`

UnsetHealthCheckCount ensures that no value is present for HealthCheckCount, not even an explicit nil
### GetHealthCheckInterval

`func (o *LbHealthCheckSet) GetHealthCheckInterval() int32`

GetHealthCheckInterval returns the HealthCheckInterval field if non-nil, zero value otherwise.

### GetHealthCheckIntervalOk

`func (o *LbHealthCheckSet) GetHealthCheckIntervalOk() (*int32, bool)`

GetHealthCheckIntervalOk returns a tuple with the HealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckInterval

`func (o *LbHealthCheckSet) SetHealthCheckInterval(v int32)`

SetHealthCheckInterval sets HealthCheckInterval field to given value.

### HasHealthCheckInterval

`func (o *LbHealthCheckSet) HasHealthCheckInterval() bool`

HasHealthCheckInterval returns a boolean if a field has been set.

### SetHealthCheckIntervalNil

`func (o *LbHealthCheckSet) SetHealthCheckIntervalNil(b bool)`

 SetHealthCheckIntervalNil sets the value for HealthCheckInterval to be an explicit nil

### UnsetHealthCheckInterval
`func (o *LbHealthCheckSet) UnsetHealthCheckInterval()`

UnsetHealthCheckInterval ensures that no value is present for HealthCheckInterval, not even an explicit nil
### GetHealthCheckPort

`func (o *LbHealthCheckSet) GetHealthCheckPort() int32`

GetHealthCheckPort returns the HealthCheckPort field if non-nil, zero value otherwise.

### GetHealthCheckPortOk

`func (o *LbHealthCheckSet) GetHealthCheckPortOk() (*int32, bool)`

GetHealthCheckPortOk returns a tuple with the HealthCheckPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPort

`func (o *LbHealthCheckSet) SetHealthCheckPort(v int32)`

SetHealthCheckPort sets HealthCheckPort field to given value.

### HasHealthCheckPort

`func (o *LbHealthCheckSet) HasHealthCheckPort() bool`

HasHealthCheckPort returns a boolean if a field has been set.

### SetHealthCheckPortNil

`func (o *LbHealthCheckSet) SetHealthCheckPortNil(b bool)`

 SetHealthCheckPortNil sets the value for HealthCheckPort to be an explicit nil

### UnsetHealthCheckPort
`func (o *LbHealthCheckSet) UnsetHealthCheckPort()`

UnsetHealthCheckPort ensures that no value is present for HealthCheckPort, not even an explicit nil
### GetHealthCheckTimeout

`func (o *LbHealthCheckSet) GetHealthCheckTimeout() int32`

GetHealthCheckTimeout returns the HealthCheckTimeout field if non-nil, zero value otherwise.

### GetHealthCheckTimeoutOk

`func (o *LbHealthCheckSet) GetHealthCheckTimeoutOk() (*int32, bool)`

GetHealthCheckTimeoutOk returns a tuple with the HealthCheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckTimeout

`func (o *LbHealthCheckSet) SetHealthCheckTimeout(v int32)`

SetHealthCheckTimeout sets HealthCheckTimeout field to given value.

### HasHealthCheckTimeout

`func (o *LbHealthCheckSet) HasHealthCheckTimeout() bool`

HasHealthCheckTimeout returns a boolean if a field has been set.

### SetHealthCheckTimeoutNil

`func (o *LbHealthCheckSet) SetHealthCheckTimeoutNil(b bool)`

 SetHealthCheckTimeoutNil sets the value for HealthCheckTimeout to be an explicit nil

### UnsetHealthCheckTimeout
`func (o *LbHealthCheckSet) UnsetHealthCheckTimeout()`

UnsetHealthCheckTimeout ensures that no value is present for HealthCheckTimeout, not even an explicit nil
### GetHealthCheckUrl

`func (o *LbHealthCheckSet) GetHealthCheckUrl() string`

GetHealthCheckUrl returns the HealthCheckUrl field if non-nil, zero value otherwise.

### GetHealthCheckUrlOk

`func (o *LbHealthCheckSet) GetHealthCheckUrlOk() (*string, bool)`

GetHealthCheckUrlOk returns a tuple with the HealthCheckUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUrl

`func (o *LbHealthCheckSet) SetHealthCheckUrl(v string)`

SetHealthCheckUrl sets HealthCheckUrl field to given value.

### HasHealthCheckUrl

`func (o *LbHealthCheckSet) HasHealthCheckUrl() bool`

HasHealthCheckUrl returns a boolean if a field has been set.

### SetHealthCheckUrlNil

`func (o *LbHealthCheckSet) SetHealthCheckUrlNil(b bool)`

 SetHealthCheckUrlNil sets the value for HealthCheckUrl to be an explicit nil

### UnsetHealthCheckUrl
`func (o *LbHealthCheckSet) UnsetHealthCheckUrl()`

UnsetHealthCheckUrl ensures that no value is present for HealthCheckUrl, not even an explicit nil
### GetHttpMethod

`func (o *LbHealthCheckSet) GetHttpMethod() LbMonitorHttpMethod`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *LbHealthCheckSet) GetHttpMethodOk() (*LbMonitorHttpMethod, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *LbHealthCheckSet) SetHttpMethod(v LbMonitorHttpMethod)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *LbHealthCheckSet) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### SetHttpMethodNil

`func (o *LbHealthCheckSet) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *LbHealthCheckSet) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetProtocol

`func (o *LbHealthCheckSet) GetProtocol() LbMonitorProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LbHealthCheckSet) GetProtocolOk() (*LbMonitorProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LbHealthCheckSet) SetProtocol(v LbMonitorProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LbHealthCheckSet) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *LbHealthCheckSet) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *LbHealthCheckSet) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetRequestData

`func (o *LbHealthCheckSet) GetRequestData() string`

GetRequestData returns the RequestData field if non-nil, zero value otherwise.

### GetRequestDataOk

`func (o *LbHealthCheckSet) GetRequestDataOk() (*string, bool)`

GetRequestDataOk returns a tuple with the RequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestData

`func (o *LbHealthCheckSet) SetRequestData(v string)`

SetRequestData sets RequestData field to given value.

### HasRequestData

`func (o *LbHealthCheckSet) HasRequestData() bool`

HasRequestData returns a boolean if a field has been set.

### SetRequestDataNil

`func (o *LbHealthCheckSet) SetRequestDataNil(b bool)`

 SetRequestDataNil sets the value for RequestData to be an explicit nil

### UnsetRequestData
`func (o *LbHealthCheckSet) UnsetRequestData()`

UnsetRequestData ensures that no value is present for RequestData, not even an explicit nil
### GetResponseCode

`func (o *LbHealthCheckSet) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *LbHealthCheckSet) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *LbHealthCheckSet) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *LbHealthCheckSet) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### SetResponseCodeNil

`func (o *LbHealthCheckSet) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *LbHealthCheckSet) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


