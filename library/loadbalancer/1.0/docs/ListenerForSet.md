# ListenerForSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**HttpsRedirection** | Pointer to **NullableBool** |  | [optional] 
**InsertClientIp** | Pointer to **NullableBool** |  | [optional] 
**Persistence** | Pointer to **NullableString** |  | [optional] 
**ResponseTimeout** | Pointer to **NullableInt32** |  | [optional] 
**ServerGroupId** | Pointer to **NullableString** |  | [optional] 
**SessionDurationTime** | Pointer to **NullableInt32** |  | [optional] 
**SslCertificate** | Pointer to [**NullableSslCertificate**](SslCertificate.md) |  | [optional] 
**SupportHttp2** | Pointer to **bool** | Whether the HTTP 2.0 is supported | [optional] 
**UrlHandler** | Pointer to **[]interface{}** |  | [optional] 
**UrlRedirection** | Pointer to **[]interface{}** |  | [optional] 
**XForwardedFor** | Pointer to **NullableBool** |  | [optional] 
**XForwardedPort** | Pointer to **NullableBool** |  | [optional] 
**XForwardedProto** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewListenerForSet

`func NewListenerForSet() *ListenerForSet`

NewListenerForSet instantiates a new ListenerForSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerForSetWithDefaults

`func NewListenerForSetWithDefaults() *ListenerForSet`

NewListenerForSetWithDefaults instantiates a new ListenerForSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ListenerForSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListenerForSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListenerForSet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListenerForSet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListenerForSet) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListenerForSet) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHttpsRedirection

`func (o *ListenerForSet) GetHttpsRedirection() bool`

GetHttpsRedirection returns the HttpsRedirection field if non-nil, zero value otherwise.

### GetHttpsRedirectionOk

`func (o *ListenerForSet) GetHttpsRedirectionOk() (*bool, bool)`

GetHttpsRedirectionOk returns a tuple with the HttpsRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRedirection

`func (o *ListenerForSet) SetHttpsRedirection(v bool)`

SetHttpsRedirection sets HttpsRedirection field to given value.

### HasHttpsRedirection

`func (o *ListenerForSet) HasHttpsRedirection() bool`

HasHttpsRedirection returns a boolean if a field has been set.

### SetHttpsRedirectionNil

`func (o *ListenerForSet) SetHttpsRedirectionNil(b bool)`

 SetHttpsRedirectionNil sets the value for HttpsRedirection to be an explicit nil

### UnsetHttpsRedirection
`func (o *ListenerForSet) UnsetHttpsRedirection()`

UnsetHttpsRedirection ensures that no value is present for HttpsRedirection, not even an explicit nil
### GetInsertClientIp

`func (o *ListenerForSet) GetInsertClientIp() bool`

GetInsertClientIp returns the InsertClientIp field if non-nil, zero value otherwise.

### GetInsertClientIpOk

`func (o *ListenerForSet) GetInsertClientIpOk() (*bool, bool)`

GetInsertClientIpOk returns a tuple with the InsertClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertClientIp

`func (o *ListenerForSet) SetInsertClientIp(v bool)`

SetInsertClientIp sets InsertClientIp field to given value.

### HasInsertClientIp

`func (o *ListenerForSet) HasInsertClientIp() bool`

HasInsertClientIp returns a boolean if a field has been set.

### SetInsertClientIpNil

`func (o *ListenerForSet) SetInsertClientIpNil(b bool)`

 SetInsertClientIpNil sets the value for InsertClientIp to be an explicit nil

### UnsetInsertClientIp
`func (o *ListenerForSet) UnsetInsertClientIp()`

UnsetInsertClientIp ensures that no value is present for InsertClientIp, not even an explicit nil
### GetPersistence

`func (o *ListenerForSet) GetPersistence() string`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *ListenerForSet) GetPersistenceOk() (*string, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *ListenerForSet) SetPersistence(v string)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *ListenerForSet) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### SetPersistenceNil

`func (o *ListenerForSet) SetPersistenceNil(b bool)`

 SetPersistenceNil sets the value for Persistence to be an explicit nil

### UnsetPersistence
`func (o *ListenerForSet) UnsetPersistence()`

UnsetPersistence ensures that no value is present for Persistence, not even an explicit nil
### GetResponseTimeout

`func (o *ListenerForSet) GetResponseTimeout() int32`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *ListenerForSet) GetResponseTimeoutOk() (*int32, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *ListenerForSet) SetResponseTimeout(v int32)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *ListenerForSet) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### SetResponseTimeoutNil

`func (o *ListenerForSet) SetResponseTimeoutNil(b bool)`

 SetResponseTimeoutNil sets the value for ResponseTimeout to be an explicit nil

### UnsetResponseTimeout
`func (o *ListenerForSet) UnsetResponseTimeout()`

UnsetResponseTimeout ensures that no value is present for ResponseTimeout, not even an explicit nil
### GetServerGroupId

`func (o *ListenerForSet) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ListenerForSet) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ListenerForSet) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *ListenerForSet) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### SetServerGroupIdNil

`func (o *ListenerForSet) SetServerGroupIdNil(b bool)`

 SetServerGroupIdNil sets the value for ServerGroupId to be an explicit nil

### UnsetServerGroupId
`func (o *ListenerForSet) UnsetServerGroupId()`

UnsetServerGroupId ensures that no value is present for ServerGroupId, not even an explicit nil
### GetSessionDurationTime

`func (o *ListenerForSet) GetSessionDurationTime() int32`

GetSessionDurationTime returns the SessionDurationTime field if non-nil, zero value otherwise.

### GetSessionDurationTimeOk

`func (o *ListenerForSet) GetSessionDurationTimeOk() (*int32, bool)`

GetSessionDurationTimeOk returns a tuple with the SessionDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDurationTime

`func (o *ListenerForSet) SetSessionDurationTime(v int32)`

SetSessionDurationTime sets SessionDurationTime field to given value.

### HasSessionDurationTime

`func (o *ListenerForSet) HasSessionDurationTime() bool`

HasSessionDurationTime returns a boolean if a field has been set.

### SetSessionDurationTimeNil

`func (o *ListenerForSet) SetSessionDurationTimeNil(b bool)`

 SetSessionDurationTimeNil sets the value for SessionDurationTime to be an explicit nil

### UnsetSessionDurationTime
`func (o *ListenerForSet) UnsetSessionDurationTime()`

UnsetSessionDurationTime ensures that no value is present for SessionDurationTime, not even an explicit nil
### GetSslCertificate

`func (o *ListenerForSet) GetSslCertificate() SslCertificate`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *ListenerForSet) GetSslCertificateOk() (*SslCertificate, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *ListenerForSet) SetSslCertificate(v SslCertificate)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *ListenerForSet) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### SetSslCertificateNil

`func (o *ListenerForSet) SetSslCertificateNil(b bool)`

 SetSslCertificateNil sets the value for SslCertificate to be an explicit nil

### UnsetSslCertificate
`func (o *ListenerForSet) UnsetSslCertificate()`

UnsetSslCertificate ensures that no value is present for SslCertificate, not even an explicit nil
### GetSupportHttp2

`func (o *ListenerForSet) GetSupportHttp2() bool`

GetSupportHttp2 returns the SupportHttp2 field if non-nil, zero value otherwise.

### GetSupportHttp2Ok

`func (o *ListenerForSet) GetSupportHttp2Ok() (*bool, bool)`

GetSupportHttp2Ok returns a tuple with the SupportHttp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHttp2

`func (o *ListenerForSet) SetSupportHttp2(v bool)`

SetSupportHttp2 sets SupportHttp2 field to given value.

### HasSupportHttp2

`func (o *ListenerForSet) HasSupportHttp2() bool`

HasSupportHttp2 returns a boolean if a field has been set.

### GetUrlHandler

`func (o *ListenerForSet) GetUrlHandler() []interface{}`

GetUrlHandler returns the UrlHandler field if non-nil, zero value otherwise.

### GetUrlHandlerOk

`func (o *ListenerForSet) GetUrlHandlerOk() (*[]interface{}, bool)`

GetUrlHandlerOk returns a tuple with the UrlHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlHandler

`func (o *ListenerForSet) SetUrlHandler(v []interface{})`

SetUrlHandler sets UrlHandler field to given value.

### HasUrlHandler

`func (o *ListenerForSet) HasUrlHandler() bool`

HasUrlHandler returns a boolean if a field has been set.

### SetUrlHandlerNil

`func (o *ListenerForSet) SetUrlHandlerNil(b bool)`

 SetUrlHandlerNil sets the value for UrlHandler to be an explicit nil

### UnsetUrlHandler
`func (o *ListenerForSet) UnsetUrlHandler()`

UnsetUrlHandler ensures that no value is present for UrlHandler, not even an explicit nil
### GetUrlRedirection

`func (o *ListenerForSet) GetUrlRedirection() []interface{}`

GetUrlRedirection returns the UrlRedirection field if non-nil, zero value otherwise.

### GetUrlRedirectionOk

`func (o *ListenerForSet) GetUrlRedirectionOk() (*[]interface{}, bool)`

GetUrlRedirectionOk returns a tuple with the UrlRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirection

`func (o *ListenerForSet) SetUrlRedirection(v []interface{})`

SetUrlRedirection sets UrlRedirection field to given value.

### HasUrlRedirection

`func (o *ListenerForSet) HasUrlRedirection() bool`

HasUrlRedirection returns a boolean if a field has been set.

### SetUrlRedirectionNil

`func (o *ListenerForSet) SetUrlRedirectionNil(b bool)`

 SetUrlRedirectionNil sets the value for UrlRedirection to be an explicit nil

### UnsetUrlRedirection
`func (o *ListenerForSet) UnsetUrlRedirection()`

UnsetUrlRedirection ensures that no value is present for UrlRedirection, not even an explicit nil
### GetXForwardedFor

`func (o *ListenerForSet) GetXForwardedFor() bool`

GetXForwardedFor returns the XForwardedFor field if non-nil, zero value otherwise.

### GetXForwardedForOk

`func (o *ListenerForSet) GetXForwardedForOk() (*bool, bool)`

GetXForwardedForOk returns a tuple with the XForwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedFor

`func (o *ListenerForSet) SetXForwardedFor(v bool)`

SetXForwardedFor sets XForwardedFor field to given value.

### HasXForwardedFor

`func (o *ListenerForSet) HasXForwardedFor() bool`

HasXForwardedFor returns a boolean if a field has been set.

### SetXForwardedForNil

`func (o *ListenerForSet) SetXForwardedForNil(b bool)`

 SetXForwardedForNil sets the value for XForwardedFor to be an explicit nil

### UnsetXForwardedFor
`func (o *ListenerForSet) UnsetXForwardedFor()`

UnsetXForwardedFor ensures that no value is present for XForwardedFor, not even an explicit nil
### GetXForwardedPort

`func (o *ListenerForSet) GetXForwardedPort() bool`

GetXForwardedPort returns the XForwardedPort field if non-nil, zero value otherwise.

### GetXForwardedPortOk

`func (o *ListenerForSet) GetXForwardedPortOk() (*bool, bool)`

GetXForwardedPortOk returns a tuple with the XForwardedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedPort

`func (o *ListenerForSet) SetXForwardedPort(v bool)`

SetXForwardedPort sets XForwardedPort field to given value.

### HasXForwardedPort

`func (o *ListenerForSet) HasXForwardedPort() bool`

HasXForwardedPort returns a boolean if a field has been set.

### SetXForwardedPortNil

`func (o *ListenerForSet) SetXForwardedPortNil(b bool)`

 SetXForwardedPortNil sets the value for XForwardedPort to be an explicit nil

### UnsetXForwardedPort
`func (o *ListenerForSet) UnsetXForwardedPort()`

UnsetXForwardedPort ensures that no value is present for XForwardedPort, not even an explicit nil
### GetXForwardedProto

`func (o *ListenerForSet) GetXForwardedProto() bool`

GetXForwardedProto returns the XForwardedProto field if non-nil, zero value otherwise.

### GetXForwardedProtoOk

`func (o *ListenerForSet) GetXForwardedProtoOk() (*bool, bool)`

GetXForwardedProtoOk returns a tuple with the XForwardedProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedProto

`func (o *ListenerForSet) SetXForwardedProto(v bool)`

SetXForwardedProto sets XForwardedProto field to given value.

### HasXForwardedProto

`func (o *ListenerForSet) HasXForwardedProto() bool`

HasXForwardedProto returns a boolean if a field has been set.

### SetXForwardedProtoNil

`func (o *ListenerForSet) SetXForwardedProtoNil(b bool)`

 SetXForwardedProtoNil sets the value for XForwardedProto to be an explicit nil

### UnsetXForwardedProto
`func (o *ListenerForSet) UnsetXForwardedProto()`

UnsetXForwardedProto ensures that no value is present for XForwardedProto, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


