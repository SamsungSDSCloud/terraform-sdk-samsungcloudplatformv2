# IngressTls

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | **[]string** | hosts | 
**SecretName** | **NullableString** |  | 

## Methods

### NewIngressTls

`func NewIngressTls(hosts []string, secretName NullableString, ) *IngressTls`

NewIngressTls instantiates a new IngressTls object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressTlsWithDefaults

`func NewIngressTlsWithDefaults() *IngressTls`

NewIngressTlsWithDefaults instantiates a new IngressTls object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *IngressTls) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *IngressTls) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *IngressTls) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetSecretName

`func (o *IngressTls) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *IngressTls) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *IngressTls) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### SetSecretNameNil

`func (o *IngressTls) SetSecretNameNil(b bool)`

 SetSecretNameNil sets the value for SecretName to be an explicit nil

### UnsetSecretName
`func (o *IngressTls) UnsetSecretName()`

UnsetSecretName ensures that no value is present for SecretName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


