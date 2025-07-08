# IngressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **NullableString** |  | 
**IngressRulePaths** | [**[]IngressRulePath**](IngressRulePath.md) | Ingress Rule Paths | 

## Methods

### NewIngressRule

`func NewIngressRule(host NullableString, ingressRulePaths []IngressRulePath, ) *IngressRule`

NewIngressRule instantiates a new IngressRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressRuleWithDefaults

`func NewIngressRuleWithDefaults() *IngressRule`

NewIngressRuleWithDefaults instantiates a new IngressRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *IngressRule) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IngressRule) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IngressRule) SetHost(v string)`

SetHost sets Host field to given value.


### SetHostNil

`func (o *IngressRule) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *IngressRule) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetIngressRulePaths

`func (o *IngressRule) GetIngressRulePaths() []IngressRulePath`

GetIngressRulePaths returns the IngressRulePaths field if non-nil, zero value otherwise.

### GetIngressRulePathsOk

`func (o *IngressRule) GetIngressRulePathsOk() (*[]IngressRulePath, bool)`

GetIngressRulePathsOk returns a tuple with the IngressRulePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRulePaths

`func (o *IngressRule) SetIngressRulePaths(v []IngressRulePath)`

SetIngressRulePaths sets IngressRulePaths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


