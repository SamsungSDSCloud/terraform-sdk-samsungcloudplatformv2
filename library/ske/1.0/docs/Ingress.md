# Ingress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**IngressClass** | **NullableString** |  | 
**IngressDefaultBackend** | **NullableString** |  | 
**IngressHost** | **[]string** | Ingress Host | 
**IngressIp** | **[]string** | Ingress Ip | 
**IngressPort** | **string** | Ingress Port | 
**IngressRules** | [**[]IngressRule**](IngressRule.md) | Ingress Rules | 
**IngressTlses** | [**[]IngressTls**](IngressTls.md) | Ingress Tlses | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Ingress Name | 
**NamespaceName** | **string** | Namespace Name | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewIngress

`func NewIngress(age string, annotations []string, clusterId string, clusterName string, createdAt time.Time, ingressClass NullableString, ingressDefaultBackend NullableString, ingressHost []string, ingressIp []string, ingressPort string, ingressRules []IngressRule, ingressTlses []IngressTls, labels []string, name string, namespaceName string, uid string, yaml string, ) *Ingress`

NewIngress instantiates a new Ingress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressWithDefaults

`func NewIngressWithDefaults() *Ingress`

NewIngressWithDefaults instantiates a new Ingress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *Ingress) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *Ingress) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *Ingress) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *Ingress) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *Ingress) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *Ingress) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *Ingress) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Ingress) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Ingress) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *Ingress) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Ingress) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Ingress) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *Ingress) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Ingress) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Ingress) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIngressClass

`func (o *Ingress) GetIngressClass() string`

GetIngressClass returns the IngressClass field if non-nil, zero value otherwise.

### GetIngressClassOk

`func (o *Ingress) GetIngressClassOk() (*string, bool)`

GetIngressClassOk returns a tuple with the IngressClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressClass

`func (o *Ingress) SetIngressClass(v string)`

SetIngressClass sets IngressClass field to given value.


### SetIngressClassNil

`func (o *Ingress) SetIngressClassNil(b bool)`

 SetIngressClassNil sets the value for IngressClass to be an explicit nil

### UnsetIngressClass
`func (o *Ingress) UnsetIngressClass()`

UnsetIngressClass ensures that no value is present for IngressClass, not even an explicit nil
### GetIngressDefaultBackend

`func (o *Ingress) GetIngressDefaultBackend() string`

GetIngressDefaultBackend returns the IngressDefaultBackend field if non-nil, zero value otherwise.

### GetIngressDefaultBackendOk

`func (o *Ingress) GetIngressDefaultBackendOk() (*string, bool)`

GetIngressDefaultBackendOk returns a tuple with the IngressDefaultBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressDefaultBackend

`func (o *Ingress) SetIngressDefaultBackend(v string)`

SetIngressDefaultBackend sets IngressDefaultBackend field to given value.


### SetIngressDefaultBackendNil

`func (o *Ingress) SetIngressDefaultBackendNil(b bool)`

 SetIngressDefaultBackendNil sets the value for IngressDefaultBackend to be an explicit nil

### UnsetIngressDefaultBackend
`func (o *Ingress) UnsetIngressDefaultBackend()`

UnsetIngressDefaultBackend ensures that no value is present for IngressDefaultBackend, not even an explicit nil
### GetIngressHost

`func (o *Ingress) GetIngressHost() []string`

GetIngressHost returns the IngressHost field if non-nil, zero value otherwise.

### GetIngressHostOk

`func (o *Ingress) GetIngressHostOk() (*[]string, bool)`

GetIngressHostOk returns a tuple with the IngressHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressHost

`func (o *Ingress) SetIngressHost(v []string)`

SetIngressHost sets IngressHost field to given value.


### GetIngressIp

`func (o *Ingress) GetIngressIp() []string`

GetIngressIp returns the IngressIp field if non-nil, zero value otherwise.

### GetIngressIpOk

`func (o *Ingress) GetIngressIpOk() (*[]string, bool)`

GetIngressIpOk returns a tuple with the IngressIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressIp

`func (o *Ingress) SetIngressIp(v []string)`

SetIngressIp sets IngressIp field to given value.


### GetIngressPort

`func (o *Ingress) GetIngressPort() string`

GetIngressPort returns the IngressPort field if non-nil, zero value otherwise.

### GetIngressPortOk

`func (o *Ingress) GetIngressPortOk() (*string, bool)`

GetIngressPortOk returns a tuple with the IngressPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPort

`func (o *Ingress) SetIngressPort(v string)`

SetIngressPort sets IngressPort field to given value.


### GetIngressRules

`func (o *Ingress) GetIngressRules() []IngressRule`

GetIngressRules returns the IngressRules field if non-nil, zero value otherwise.

### GetIngressRulesOk

`func (o *Ingress) GetIngressRulesOk() (*[]IngressRule, bool)`

GetIngressRulesOk returns a tuple with the IngressRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressRules

`func (o *Ingress) SetIngressRules(v []IngressRule)`

SetIngressRules sets IngressRules field to given value.


### GetIngressTlses

`func (o *Ingress) GetIngressTlses() []IngressTls`

GetIngressTlses returns the IngressTlses field if non-nil, zero value otherwise.

### GetIngressTlsesOk

`func (o *Ingress) GetIngressTlsesOk() (*[]IngressTls, bool)`

GetIngressTlsesOk returns a tuple with the IngressTlses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressTlses

`func (o *Ingress) SetIngressTlses(v []IngressTls)`

SetIngressTlses sets IngressTlses field to given value.


### GetLabels

`func (o *Ingress) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Ingress) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Ingress) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *Ingress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ingress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ingress) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *Ingress) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *Ingress) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *Ingress) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetUid

`func (o *Ingress) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Ingress) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Ingress) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *Ingress) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *Ingress) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *Ingress) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


