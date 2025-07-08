# ListenerForCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**HttpsRedirection** | Pointer to **NullableBool** |  | [optional] 
**InsertClientIp** | Pointer to **NullableBool** |  | [optional] 
**LoadbalancerId** | **string** | The ID of the load balancer. | 
**Name** | **string** | The name of the listener. | 
**Persistence** | Pointer to **NullableString** |  | [optional] 
**Protocol** | **string** | The protocol of the listener. | 
**ResponseTimeout** | Pointer to **NullableInt32** |  | [optional] 
**ServerGroupId** | Pointer to **NullableString** |  | [optional] 
**ServicePort** | **int32** | The port of the listener. | 
**SessionDurationTime** | **int32** | The session duration time. | 
**SslCertificate** | Pointer to [**NullableSslCertificate**](SslCertificate.md) |  | [optional] 
**SupportHttp2** | Pointer to **NullableBool** |  | [optional] 
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]
**UrlHandler** | Pointer to **[]interface{}** |  | [optional] 
**UrlRedirection** | Pointer to **[]interface{}** |  | [optional] 
**XForwardedFor** | Pointer to **NullableBool** |  | [optional] 
**XForwardedPort** | Pointer to **NullableBool** |  | [optional] 
**XForwardedProto** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewListenerForCreate

`func NewListenerForCreate(loadbalancerId string, name string, protocol string, servicePort int32, sessionDurationTime int32, ) *ListenerForCreate`

NewListenerForCreate instantiates a new ListenerForCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerForCreateWithDefaults

`func NewListenerForCreateWithDefaults() *ListenerForCreate`

NewListenerForCreateWithDefaults instantiates a new ListenerForCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ListenerForCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListenerForCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListenerForCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListenerForCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListenerForCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListenerForCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHttpsRedirection

`func (o *ListenerForCreate) GetHttpsRedirection() bool`

GetHttpsRedirection returns the HttpsRedirection field if non-nil, zero value otherwise.

### GetHttpsRedirectionOk

`func (o *ListenerForCreate) GetHttpsRedirectionOk() (*bool, bool)`

GetHttpsRedirectionOk returns a tuple with the HttpsRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRedirection

`func (o *ListenerForCreate) SetHttpsRedirection(v bool)`

SetHttpsRedirection sets HttpsRedirection field to given value.

### HasHttpsRedirection

`func (o *ListenerForCreate) HasHttpsRedirection() bool`

HasHttpsRedirection returns a boolean if a field has been set.

### SetHttpsRedirectionNil

`func (o *ListenerForCreate) SetHttpsRedirectionNil(b bool)`

 SetHttpsRedirectionNil sets the value for HttpsRedirection to be an explicit nil

### UnsetHttpsRedirection
`func (o *ListenerForCreate) UnsetHttpsRedirection()`

UnsetHttpsRedirection ensures that no value is present for HttpsRedirection, not even an explicit nil
### GetInsertClientIp

`func (o *ListenerForCreate) GetInsertClientIp() bool`

GetInsertClientIp returns the InsertClientIp field if non-nil, zero value otherwise.

### GetInsertClientIpOk

`func (o *ListenerForCreate) GetInsertClientIpOk() (*bool, bool)`

GetInsertClientIpOk returns a tuple with the InsertClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertClientIp

`func (o *ListenerForCreate) SetInsertClientIp(v bool)`

SetInsertClientIp sets InsertClientIp field to given value.

### HasInsertClientIp

`func (o *ListenerForCreate) HasInsertClientIp() bool`

HasInsertClientIp returns a boolean if a field has been set.

### SetInsertClientIpNil

`func (o *ListenerForCreate) SetInsertClientIpNil(b bool)`

 SetInsertClientIpNil sets the value for InsertClientIp to be an explicit nil

### UnsetInsertClientIp
`func (o *ListenerForCreate) UnsetInsertClientIp()`

UnsetInsertClientIp ensures that no value is present for InsertClientIp, not even an explicit nil
### GetLoadbalancerId

`func (o *ListenerForCreate) GetLoadbalancerId() string`

GetLoadbalancerId returns the LoadbalancerId field if non-nil, zero value otherwise.

### GetLoadbalancerIdOk

`func (o *ListenerForCreate) GetLoadbalancerIdOk() (*string, bool)`

GetLoadbalancerIdOk returns a tuple with the LoadbalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancerId

`func (o *ListenerForCreate) SetLoadbalancerId(v string)`

SetLoadbalancerId sets LoadbalancerId field to given value.


### GetName

`func (o *ListenerForCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListenerForCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListenerForCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPersistence

`func (o *ListenerForCreate) GetPersistence() string`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *ListenerForCreate) GetPersistenceOk() (*string, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *ListenerForCreate) SetPersistence(v string)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *ListenerForCreate) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### SetPersistenceNil

`func (o *ListenerForCreate) SetPersistenceNil(b bool)`

 SetPersistenceNil sets the value for Persistence to be an explicit nil

### UnsetPersistence
`func (o *ListenerForCreate) UnsetPersistence()`

UnsetPersistence ensures that no value is present for Persistence, not even an explicit nil
### GetProtocol

`func (o *ListenerForCreate) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ListenerForCreate) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ListenerForCreate) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetResponseTimeout

`func (o *ListenerForCreate) GetResponseTimeout() int32`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *ListenerForCreate) GetResponseTimeoutOk() (*int32, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *ListenerForCreate) SetResponseTimeout(v int32)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *ListenerForCreate) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### SetResponseTimeoutNil

`func (o *ListenerForCreate) SetResponseTimeoutNil(b bool)`

 SetResponseTimeoutNil sets the value for ResponseTimeout to be an explicit nil

### UnsetResponseTimeout
`func (o *ListenerForCreate) UnsetResponseTimeout()`

UnsetResponseTimeout ensures that no value is present for ResponseTimeout, not even an explicit nil
### GetServerGroupId

`func (o *ListenerForCreate) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ListenerForCreate) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ListenerForCreate) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *ListenerForCreate) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### SetServerGroupIdNil

`func (o *ListenerForCreate) SetServerGroupIdNil(b bool)`

 SetServerGroupIdNil sets the value for ServerGroupId to be an explicit nil

### UnsetServerGroupId
`func (o *ListenerForCreate) UnsetServerGroupId()`

UnsetServerGroupId ensures that no value is present for ServerGroupId, not even an explicit nil
### GetServicePort

`func (o *ListenerForCreate) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ListenerForCreate) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ListenerForCreate) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.


### GetSessionDurationTime

`func (o *ListenerForCreate) GetSessionDurationTime() int32`

GetSessionDurationTime returns the SessionDurationTime field if non-nil, zero value otherwise.

### GetSessionDurationTimeOk

`func (o *ListenerForCreate) GetSessionDurationTimeOk() (*int32, bool)`

GetSessionDurationTimeOk returns a tuple with the SessionDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDurationTime

`func (o *ListenerForCreate) SetSessionDurationTime(v int32)`

SetSessionDurationTime sets SessionDurationTime field to given value.


### GetSslCertificate

`func (o *ListenerForCreate) GetSslCertificate() SslCertificate`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *ListenerForCreate) GetSslCertificateOk() (*SslCertificate, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *ListenerForCreate) SetSslCertificate(v SslCertificate)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *ListenerForCreate) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### SetSslCertificateNil

`func (o *ListenerForCreate) SetSslCertificateNil(b bool)`

 SetSslCertificateNil sets the value for SslCertificate to be an explicit nil

### UnsetSslCertificate
`func (o *ListenerForCreate) UnsetSslCertificate()`

UnsetSslCertificate ensures that no value is present for SslCertificate, not even an explicit nil
### GetSupportHttp2

`func (o *ListenerForCreate) GetSupportHttp2() bool`

GetSupportHttp2 returns the SupportHttp2 field if non-nil, zero value otherwise.

### GetSupportHttp2Ok

`func (o *ListenerForCreate) GetSupportHttp2Ok() (*bool, bool)`

GetSupportHttp2Ok returns a tuple with the SupportHttp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHttp2

`func (o *ListenerForCreate) SetSupportHttp2(v bool)`

SetSupportHttp2 sets SupportHttp2 field to given value.

### HasSupportHttp2

`func (o *ListenerForCreate) HasSupportHttp2() bool`

HasSupportHttp2 returns a boolean if a field has been set.

### SetSupportHttp2Nil

`func (o *ListenerForCreate) SetSupportHttp2Nil(b bool)`

 SetSupportHttp2Nil sets the value for SupportHttp2 to be an explicit nil

### UnsetSupportHttp2
`func (o *ListenerForCreate) UnsetSupportHttp2()`

UnsetSupportHttp2 ensures that no value is present for SupportHttp2, not even an explicit nil
### GetTags

`func (o *ListenerForCreate) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListenerForCreate) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListenerForCreate) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListenerForCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUrlHandler

`func (o *ListenerForCreate) GetUrlHandler() []interface{}`

GetUrlHandler returns the UrlHandler field if non-nil, zero value otherwise.

### GetUrlHandlerOk

`func (o *ListenerForCreate) GetUrlHandlerOk() (*[]interface{}, bool)`

GetUrlHandlerOk returns a tuple with the UrlHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlHandler

`func (o *ListenerForCreate) SetUrlHandler(v []interface{})`

SetUrlHandler sets UrlHandler field to given value.

### HasUrlHandler

`func (o *ListenerForCreate) HasUrlHandler() bool`

HasUrlHandler returns a boolean if a field has been set.

### SetUrlHandlerNil

`func (o *ListenerForCreate) SetUrlHandlerNil(b bool)`

 SetUrlHandlerNil sets the value for UrlHandler to be an explicit nil

### UnsetUrlHandler
`func (o *ListenerForCreate) UnsetUrlHandler()`

UnsetUrlHandler ensures that no value is present for UrlHandler, not even an explicit nil
### GetUrlRedirection

`func (o *ListenerForCreate) GetUrlRedirection() []interface{}`

GetUrlRedirection returns the UrlRedirection field if non-nil, zero value otherwise.

### GetUrlRedirectionOk

`func (o *ListenerForCreate) GetUrlRedirectionOk() (*[]interface{}, bool)`

GetUrlRedirectionOk returns a tuple with the UrlRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirection

`func (o *ListenerForCreate) SetUrlRedirection(v []interface{})`

SetUrlRedirection sets UrlRedirection field to given value.

### HasUrlRedirection

`func (o *ListenerForCreate) HasUrlRedirection() bool`

HasUrlRedirection returns a boolean if a field has been set.

### SetUrlRedirectionNil

`func (o *ListenerForCreate) SetUrlRedirectionNil(b bool)`

 SetUrlRedirectionNil sets the value for UrlRedirection to be an explicit nil

### UnsetUrlRedirection
`func (o *ListenerForCreate) UnsetUrlRedirection()`

UnsetUrlRedirection ensures that no value is present for UrlRedirection, not even an explicit nil
### GetXForwardedFor

`func (o *ListenerForCreate) GetXForwardedFor() bool`

GetXForwardedFor returns the XForwardedFor field if non-nil, zero value otherwise.

### GetXForwardedForOk

`func (o *ListenerForCreate) GetXForwardedForOk() (*bool, bool)`

GetXForwardedForOk returns a tuple with the XForwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedFor

`func (o *ListenerForCreate) SetXForwardedFor(v bool)`

SetXForwardedFor sets XForwardedFor field to given value.

### HasXForwardedFor

`func (o *ListenerForCreate) HasXForwardedFor() bool`

HasXForwardedFor returns a boolean if a field has been set.

### SetXForwardedForNil

`func (o *ListenerForCreate) SetXForwardedForNil(b bool)`

 SetXForwardedForNil sets the value for XForwardedFor to be an explicit nil

### UnsetXForwardedFor
`func (o *ListenerForCreate) UnsetXForwardedFor()`

UnsetXForwardedFor ensures that no value is present for XForwardedFor, not even an explicit nil
### GetXForwardedPort

`func (o *ListenerForCreate) GetXForwardedPort() bool`

GetXForwardedPort returns the XForwardedPort field if non-nil, zero value otherwise.

### GetXForwardedPortOk

`func (o *ListenerForCreate) GetXForwardedPortOk() (*bool, bool)`

GetXForwardedPortOk returns a tuple with the XForwardedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedPort

`func (o *ListenerForCreate) SetXForwardedPort(v bool)`

SetXForwardedPort sets XForwardedPort field to given value.

### HasXForwardedPort

`func (o *ListenerForCreate) HasXForwardedPort() bool`

HasXForwardedPort returns a boolean if a field has been set.

### SetXForwardedPortNil

`func (o *ListenerForCreate) SetXForwardedPortNil(b bool)`

 SetXForwardedPortNil sets the value for XForwardedPort to be an explicit nil

### UnsetXForwardedPort
`func (o *ListenerForCreate) UnsetXForwardedPort()`

UnsetXForwardedPort ensures that no value is present for XForwardedPort, not even an explicit nil
### GetXForwardedProto

`func (o *ListenerForCreate) GetXForwardedProto() bool`

GetXForwardedProto returns the XForwardedProto field if non-nil, zero value otherwise.

### GetXForwardedProtoOk

`func (o *ListenerForCreate) GetXForwardedProtoOk() (*bool, bool)`

GetXForwardedProtoOk returns a tuple with the XForwardedProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedProto

`func (o *ListenerForCreate) SetXForwardedProto(v bool)`

SetXForwardedProto sets XForwardedProto field to given value.

### HasXForwardedProto

`func (o *ListenerForCreate) HasXForwardedProto() bool`

HasXForwardedProto returns a boolean if a field has been set.

### SetXForwardedProtoNil

`func (o *ListenerForCreate) SetXForwardedProtoNil(b bool)`

 SetXForwardedProtoNil sets the value for XForwardedProto to be an explicit nil

### UnsetXForwardedProto
`func (o *ListenerForCreate) UnsetXForwardedProto()`

UnsetXForwardedProto ensures that no value is present for XForwardedProto, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


