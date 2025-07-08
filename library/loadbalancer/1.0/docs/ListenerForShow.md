# ListenerForShow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**HttpsRedirection** | Pointer to **NullableBool** |  | [optional] 
**Id** | **string** | ID | 
**InsertClientIp** | Pointer to **NullableBool** |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | The name of the listener. | 
**Persistence** | Pointer to **NullableString** |  | [optional] 
**Protocol** | **string** | The protocol of the listener. | 
**ResponseTimeout** | Pointer to **NullableInt32** |  | [optional] 
**ServerGroupId** | Pointer to **NullableString** |  | [optional] 
**ServerGroupName** | Pointer to **NullableString** |  | [optional] 
**ServicePort** | **int32** | The port of the listener. | 
**SessionDurationTime** | Pointer to **NullableInt32** |  | [optional] 
**SslCertificate** | Pointer to [**NullableSslCertificate**](SslCertificate.md) |  | [optional] 
**State** | **string** | The state of the listener. | 
**SupportHttp2** | Pointer to **NullableBool** |  | [optional] 
**UrlHandler** | Pointer to **[]interface{}** |  | [optional] 
**UrlRedirection** | Pointer to **[]interface{}** |  | [optional] 
**XForwardedFor** | Pointer to **NullableBool** |  | [optional] 
**XForwardedPort** | Pointer to **NullableBool** |  | [optional] 
**XForwardedProto** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewListenerForShow

`func NewListenerForShow(createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, protocol string, servicePort int32, state string, ) *ListenerForShow`

NewListenerForShow instantiates a new ListenerForShow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerForShowWithDefaults

`func NewListenerForShowWithDefaults() *ListenerForShow`

NewListenerForShowWithDefaults instantiates a new ListenerForShow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ListenerForShow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListenerForShow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListenerForShow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ListenerForShow) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListenerForShow) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListenerForShow) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *ListenerForShow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListenerForShow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListenerForShow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListenerForShow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListenerForShow) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListenerForShow) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHttpsRedirection

`func (o *ListenerForShow) GetHttpsRedirection() bool`

GetHttpsRedirection returns the HttpsRedirection field if non-nil, zero value otherwise.

### GetHttpsRedirectionOk

`func (o *ListenerForShow) GetHttpsRedirectionOk() (*bool, bool)`

GetHttpsRedirectionOk returns a tuple with the HttpsRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRedirection

`func (o *ListenerForShow) SetHttpsRedirection(v bool)`

SetHttpsRedirection sets HttpsRedirection field to given value.

### HasHttpsRedirection

`func (o *ListenerForShow) HasHttpsRedirection() bool`

HasHttpsRedirection returns a boolean if a field has been set.

### SetHttpsRedirectionNil

`func (o *ListenerForShow) SetHttpsRedirectionNil(b bool)`

 SetHttpsRedirectionNil sets the value for HttpsRedirection to be an explicit nil

### UnsetHttpsRedirection
`func (o *ListenerForShow) UnsetHttpsRedirection()`

UnsetHttpsRedirection ensures that no value is present for HttpsRedirection, not even an explicit nil
### GetId

`func (o *ListenerForShow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListenerForShow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListenerForShow) SetId(v string)`

SetId sets Id field to given value.


### GetInsertClientIp

`func (o *ListenerForShow) GetInsertClientIp() bool`

GetInsertClientIp returns the InsertClientIp field if non-nil, zero value otherwise.

### GetInsertClientIpOk

`func (o *ListenerForShow) GetInsertClientIpOk() (*bool, bool)`

GetInsertClientIpOk returns a tuple with the InsertClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertClientIp

`func (o *ListenerForShow) SetInsertClientIp(v bool)`

SetInsertClientIp sets InsertClientIp field to given value.

### HasInsertClientIp

`func (o *ListenerForShow) HasInsertClientIp() bool`

HasInsertClientIp returns a boolean if a field has been set.

### SetInsertClientIpNil

`func (o *ListenerForShow) SetInsertClientIpNil(b bool)`

 SetInsertClientIpNil sets the value for InsertClientIp to be an explicit nil

### UnsetInsertClientIp
`func (o *ListenerForShow) UnsetInsertClientIp()`

UnsetInsertClientIp ensures that no value is present for InsertClientIp, not even an explicit nil
### GetModifiedAt

`func (o *ListenerForShow) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ListenerForShow) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ListenerForShow) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ListenerForShow) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ListenerForShow) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ListenerForShow) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *ListenerForShow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListenerForShow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListenerForShow) SetName(v string)`

SetName sets Name field to given value.


### GetPersistence

`func (o *ListenerForShow) GetPersistence() string`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *ListenerForShow) GetPersistenceOk() (*string, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *ListenerForShow) SetPersistence(v string)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *ListenerForShow) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### SetPersistenceNil

`func (o *ListenerForShow) SetPersistenceNil(b bool)`

 SetPersistenceNil sets the value for Persistence to be an explicit nil

### UnsetPersistence
`func (o *ListenerForShow) UnsetPersistence()`

UnsetPersistence ensures that no value is present for Persistence, not even an explicit nil
### GetProtocol

`func (o *ListenerForShow) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ListenerForShow) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ListenerForShow) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetResponseTimeout

`func (o *ListenerForShow) GetResponseTimeout() int32`

GetResponseTimeout returns the ResponseTimeout field if non-nil, zero value otherwise.

### GetResponseTimeoutOk

`func (o *ListenerForShow) GetResponseTimeoutOk() (*int32, bool)`

GetResponseTimeoutOk returns a tuple with the ResponseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeout

`func (o *ListenerForShow) SetResponseTimeout(v int32)`

SetResponseTimeout sets ResponseTimeout field to given value.

### HasResponseTimeout

`func (o *ListenerForShow) HasResponseTimeout() bool`

HasResponseTimeout returns a boolean if a field has been set.

### SetResponseTimeoutNil

`func (o *ListenerForShow) SetResponseTimeoutNil(b bool)`

 SetResponseTimeoutNil sets the value for ResponseTimeout to be an explicit nil

### UnsetResponseTimeout
`func (o *ListenerForShow) UnsetResponseTimeout()`

UnsetResponseTimeout ensures that no value is present for ResponseTimeout, not even an explicit nil
### GetServerGroupId

`func (o *ListenerForShow) GetServerGroupId() string`

GetServerGroupId returns the ServerGroupId field if non-nil, zero value otherwise.

### GetServerGroupIdOk

`func (o *ListenerForShow) GetServerGroupIdOk() (*string, bool)`

GetServerGroupIdOk returns a tuple with the ServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupId

`func (o *ListenerForShow) SetServerGroupId(v string)`

SetServerGroupId sets ServerGroupId field to given value.

### HasServerGroupId

`func (o *ListenerForShow) HasServerGroupId() bool`

HasServerGroupId returns a boolean if a field has been set.

### SetServerGroupIdNil

`func (o *ListenerForShow) SetServerGroupIdNil(b bool)`

 SetServerGroupIdNil sets the value for ServerGroupId to be an explicit nil

### UnsetServerGroupId
`func (o *ListenerForShow) UnsetServerGroupId()`

UnsetServerGroupId ensures that no value is present for ServerGroupId, not even an explicit nil
### GetServerGroupName

`func (o *ListenerForShow) GetServerGroupName() string`

GetServerGroupName returns the ServerGroupName field if non-nil, zero value otherwise.

### GetServerGroupNameOk

`func (o *ListenerForShow) GetServerGroupNameOk() (*string, bool)`

GetServerGroupNameOk returns a tuple with the ServerGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroupName

`func (o *ListenerForShow) SetServerGroupName(v string)`

SetServerGroupName sets ServerGroupName field to given value.

### HasServerGroupName

`func (o *ListenerForShow) HasServerGroupName() bool`

HasServerGroupName returns a boolean if a field has been set.

### SetServerGroupNameNil

`func (o *ListenerForShow) SetServerGroupNameNil(b bool)`

 SetServerGroupNameNil sets the value for ServerGroupName to be an explicit nil

### UnsetServerGroupName
`func (o *ListenerForShow) UnsetServerGroupName()`

UnsetServerGroupName ensures that no value is present for ServerGroupName, not even an explicit nil
### GetServicePort

`func (o *ListenerForShow) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ListenerForShow) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ListenerForShow) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.


### GetSessionDurationTime

`func (o *ListenerForShow) GetSessionDurationTime() int32`

GetSessionDurationTime returns the SessionDurationTime field if non-nil, zero value otherwise.

### GetSessionDurationTimeOk

`func (o *ListenerForShow) GetSessionDurationTimeOk() (*int32, bool)`

GetSessionDurationTimeOk returns a tuple with the SessionDurationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionDurationTime

`func (o *ListenerForShow) SetSessionDurationTime(v int32)`

SetSessionDurationTime sets SessionDurationTime field to given value.

### HasSessionDurationTime

`func (o *ListenerForShow) HasSessionDurationTime() bool`

HasSessionDurationTime returns a boolean if a field has been set.

### SetSessionDurationTimeNil

`func (o *ListenerForShow) SetSessionDurationTimeNil(b bool)`

 SetSessionDurationTimeNil sets the value for SessionDurationTime to be an explicit nil

### UnsetSessionDurationTime
`func (o *ListenerForShow) UnsetSessionDurationTime()`

UnsetSessionDurationTime ensures that no value is present for SessionDurationTime, not even an explicit nil
### GetSslCertificate

`func (o *ListenerForShow) GetSslCertificate() SslCertificate`

GetSslCertificate returns the SslCertificate field if non-nil, zero value otherwise.

### GetSslCertificateOk

`func (o *ListenerForShow) GetSslCertificateOk() (*SslCertificate, bool)`

GetSslCertificateOk returns a tuple with the SslCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertificate

`func (o *ListenerForShow) SetSslCertificate(v SslCertificate)`

SetSslCertificate sets SslCertificate field to given value.

### HasSslCertificate

`func (o *ListenerForShow) HasSslCertificate() bool`

HasSslCertificate returns a boolean if a field has been set.

### SetSslCertificateNil

`func (o *ListenerForShow) SetSslCertificateNil(b bool)`

 SetSslCertificateNil sets the value for SslCertificate to be an explicit nil

### UnsetSslCertificate
`func (o *ListenerForShow) UnsetSslCertificate()`

UnsetSslCertificate ensures that no value is present for SslCertificate, not even an explicit nil
### GetState

`func (o *ListenerForShow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListenerForShow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListenerForShow) SetState(v string)`

SetState sets State field to given value.


### GetSupportHttp2

`func (o *ListenerForShow) GetSupportHttp2() bool`

GetSupportHttp2 returns the SupportHttp2 field if non-nil, zero value otherwise.

### GetSupportHttp2Ok

`func (o *ListenerForShow) GetSupportHttp2Ok() (*bool, bool)`

GetSupportHttp2Ok returns a tuple with the SupportHttp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportHttp2

`func (o *ListenerForShow) SetSupportHttp2(v bool)`

SetSupportHttp2 sets SupportHttp2 field to given value.

### HasSupportHttp2

`func (o *ListenerForShow) HasSupportHttp2() bool`

HasSupportHttp2 returns a boolean if a field has been set.

### SetSupportHttp2Nil

`func (o *ListenerForShow) SetSupportHttp2Nil(b bool)`

 SetSupportHttp2Nil sets the value for SupportHttp2 to be an explicit nil

### UnsetSupportHttp2
`func (o *ListenerForShow) UnsetSupportHttp2()`

UnsetSupportHttp2 ensures that no value is present for SupportHttp2, not even an explicit nil
### GetUrlHandler

`func (o *ListenerForShow) GetUrlHandler() []interface{}`

GetUrlHandler returns the UrlHandler field if non-nil, zero value otherwise.

### GetUrlHandlerOk

`func (o *ListenerForShow) GetUrlHandlerOk() (*[]interface{}, bool)`

GetUrlHandlerOk returns a tuple with the UrlHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlHandler

`func (o *ListenerForShow) SetUrlHandler(v []interface{})`

SetUrlHandler sets UrlHandler field to given value.

### HasUrlHandler

`func (o *ListenerForShow) HasUrlHandler() bool`

HasUrlHandler returns a boolean if a field has been set.

### SetUrlHandlerNil

`func (o *ListenerForShow) SetUrlHandlerNil(b bool)`

 SetUrlHandlerNil sets the value for UrlHandler to be an explicit nil

### UnsetUrlHandler
`func (o *ListenerForShow) UnsetUrlHandler()`

UnsetUrlHandler ensures that no value is present for UrlHandler, not even an explicit nil
### GetUrlRedirection

`func (o *ListenerForShow) GetUrlRedirection() []interface{}`

GetUrlRedirection returns the UrlRedirection field if non-nil, zero value otherwise.

### GetUrlRedirectionOk

`func (o *ListenerForShow) GetUrlRedirectionOk() (*[]interface{}, bool)`

GetUrlRedirectionOk returns a tuple with the UrlRedirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlRedirection

`func (o *ListenerForShow) SetUrlRedirection(v []interface{})`

SetUrlRedirection sets UrlRedirection field to given value.

### HasUrlRedirection

`func (o *ListenerForShow) HasUrlRedirection() bool`

HasUrlRedirection returns a boolean if a field has been set.

### SetUrlRedirectionNil

`func (o *ListenerForShow) SetUrlRedirectionNil(b bool)`

 SetUrlRedirectionNil sets the value for UrlRedirection to be an explicit nil

### UnsetUrlRedirection
`func (o *ListenerForShow) UnsetUrlRedirection()`

UnsetUrlRedirection ensures that no value is present for UrlRedirection, not even an explicit nil
### GetXForwardedFor

`func (o *ListenerForShow) GetXForwardedFor() bool`

GetXForwardedFor returns the XForwardedFor field if non-nil, zero value otherwise.

### GetXForwardedForOk

`func (o *ListenerForShow) GetXForwardedForOk() (*bool, bool)`

GetXForwardedForOk returns a tuple with the XForwardedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedFor

`func (o *ListenerForShow) SetXForwardedFor(v bool)`

SetXForwardedFor sets XForwardedFor field to given value.

### HasXForwardedFor

`func (o *ListenerForShow) HasXForwardedFor() bool`

HasXForwardedFor returns a boolean if a field has been set.

### SetXForwardedForNil

`func (o *ListenerForShow) SetXForwardedForNil(b bool)`

 SetXForwardedForNil sets the value for XForwardedFor to be an explicit nil

### UnsetXForwardedFor
`func (o *ListenerForShow) UnsetXForwardedFor()`

UnsetXForwardedFor ensures that no value is present for XForwardedFor, not even an explicit nil
### GetXForwardedPort

`func (o *ListenerForShow) GetXForwardedPort() bool`

GetXForwardedPort returns the XForwardedPort field if non-nil, zero value otherwise.

### GetXForwardedPortOk

`func (o *ListenerForShow) GetXForwardedPortOk() (*bool, bool)`

GetXForwardedPortOk returns a tuple with the XForwardedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedPort

`func (o *ListenerForShow) SetXForwardedPort(v bool)`

SetXForwardedPort sets XForwardedPort field to given value.

### HasXForwardedPort

`func (o *ListenerForShow) HasXForwardedPort() bool`

HasXForwardedPort returns a boolean if a field has been set.

### SetXForwardedPortNil

`func (o *ListenerForShow) SetXForwardedPortNil(b bool)`

 SetXForwardedPortNil sets the value for XForwardedPort to be an explicit nil

### UnsetXForwardedPort
`func (o *ListenerForShow) UnsetXForwardedPort()`

UnsetXForwardedPort ensures that no value is present for XForwardedPort, not even an explicit nil
### GetXForwardedProto

`func (o *ListenerForShow) GetXForwardedProto() bool`

GetXForwardedProto returns the XForwardedProto field if non-nil, zero value otherwise.

### GetXForwardedProtoOk

`func (o *ListenerForShow) GetXForwardedProtoOk() (*bool, bool)`

GetXForwardedProtoOk returns a tuple with the XForwardedProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXForwardedProto

`func (o *ListenerForShow) SetXForwardedProto(v bool)`

SetXForwardedProto sets XForwardedProto field to given value.

### HasXForwardedProto

`func (o *ListenerForShow) HasXForwardedProto() bool`

HasXForwardedProto returns a boolean if a field has been set.

### SetXForwardedProtoNil

`func (o *ListenerForShow) SetXForwardedProtoNil(b bool)`

 SetXForwardedProtoNil sets the value for XForwardedProto to be an explicit nil

### UnsetXForwardedProto
`func (o *ListenerForShow) UnsetXForwardedProto()`

UnsetXForwardedProto ensures that no value is present for XForwardedProto, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


