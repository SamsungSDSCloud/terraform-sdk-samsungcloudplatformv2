# KubernetesService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**Annotations** | **[]string** | Annotations | 
**ClusterId** | **string** | Cluster ID | 
**ClusterIp** | **NullableString** |  | 
**ClusterName** | **string** | Cluster Name | 
**CreatedAt** | **time.Time** | Created At | 
**ExternalIps** | **[]string** | External Ips | 
**ExternalTrafficPolicy** | **NullableString** |  | 
**Labels** | **[]string** | Labels | 
**Name** | **string** | Service Name | 
**NamespaceName** | **string** | Namespace Name | 
**Ports** | [**[]KubernetesServicePort**](KubernetesServicePort.md) | Ports | 
**Selectors** | **[]string** | Selectors | 
**ServiceType** | **NullableString** |  | 
**SessionAffinity** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 
**Yaml** | **string** | YAML | 

## Methods

### NewKubernetesService

`func NewKubernetesService(age string, annotations []string, clusterId string, clusterIp NullableString, clusterName string, createdAt time.Time, externalIps []string, externalTrafficPolicy NullableString, labels []string, name string, namespaceName string, ports []KubernetesServicePort, selectors []string, serviceType NullableString, sessionAffinity NullableString, uid string, yaml string, ) *KubernetesService`

NewKubernetesService instantiates a new KubernetesService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServiceWithDefaults

`func NewKubernetesServiceWithDefaults() *KubernetesService`

NewKubernetesServiceWithDefaults instantiates a new KubernetesService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *KubernetesService) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *KubernetesService) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *KubernetesService) SetAge(v string)`

SetAge sets Age field to given value.


### GetAnnotations

`func (o *KubernetesService) GetAnnotations() []string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *KubernetesService) GetAnnotationsOk() (*[]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *KubernetesService) SetAnnotations(v []string)`

SetAnnotations sets Annotations field to given value.


### GetClusterId

`func (o *KubernetesService) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *KubernetesService) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *KubernetesService) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterIp

`func (o *KubernetesService) GetClusterIp() string`

GetClusterIp returns the ClusterIp field if non-nil, zero value otherwise.

### GetClusterIpOk

`func (o *KubernetesService) GetClusterIpOk() (*string, bool)`

GetClusterIpOk returns a tuple with the ClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIp

`func (o *KubernetesService) SetClusterIp(v string)`

SetClusterIp sets ClusterIp field to given value.


### SetClusterIpNil

`func (o *KubernetesService) SetClusterIpNil(b bool)`

 SetClusterIpNil sets the value for ClusterIp to be an explicit nil

### UnsetClusterIp
`func (o *KubernetesService) UnsetClusterIp()`

UnsetClusterIp ensures that no value is present for ClusterIp, not even an explicit nil
### GetClusterName

`func (o *KubernetesService) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *KubernetesService) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *KubernetesService) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCreatedAt

`func (o *KubernetesService) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesService) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesService) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExternalIps

`func (o *KubernetesService) GetExternalIps() []string`

GetExternalIps returns the ExternalIps field if non-nil, zero value otherwise.

### GetExternalIpsOk

`func (o *KubernetesService) GetExternalIpsOk() (*[]string, bool)`

GetExternalIpsOk returns a tuple with the ExternalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIps

`func (o *KubernetesService) SetExternalIps(v []string)`

SetExternalIps sets ExternalIps field to given value.


### GetExternalTrafficPolicy

`func (o *KubernetesService) GetExternalTrafficPolicy() string`

GetExternalTrafficPolicy returns the ExternalTrafficPolicy field if non-nil, zero value otherwise.

### GetExternalTrafficPolicyOk

`func (o *KubernetesService) GetExternalTrafficPolicyOk() (*string, bool)`

GetExternalTrafficPolicyOk returns a tuple with the ExternalTrafficPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalTrafficPolicy

`func (o *KubernetesService) SetExternalTrafficPolicy(v string)`

SetExternalTrafficPolicy sets ExternalTrafficPolicy field to given value.


### SetExternalTrafficPolicyNil

`func (o *KubernetesService) SetExternalTrafficPolicyNil(b bool)`

 SetExternalTrafficPolicyNil sets the value for ExternalTrafficPolicy to be an explicit nil

### UnsetExternalTrafficPolicy
`func (o *KubernetesService) UnsetExternalTrafficPolicy()`

UnsetExternalTrafficPolicy ensures that no value is present for ExternalTrafficPolicy, not even an explicit nil
### GetLabels

`func (o *KubernetesService) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *KubernetesService) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *KubernetesService) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *KubernetesService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesService) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *KubernetesService) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *KubernetesService) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *KubernetesService) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetPorts

`func (o *KubernetesService) GetPorts() []KubernetesServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KubernetesService) GetPortsOk() (*[]KubernetesServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KubernetesService) SetPorts(v []KubernetesServicePort)`

SetPorts sets Ports field to given value.


### GetSelectors

`func (o *KubernetesService) GetSelectors() []string`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *KubernetesService) GetSelectorsOk() (*[]string, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *KubernetesService) SetSelectors(v []string)`

SetSelectors sets Selectors field to given value.


### GetServiceType

`func (o *KubernetesService) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *KubernetesService) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *KubernetesService) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### SetServiceTypeNil

`func (o *KubernetesService) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *KubernetesService) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetSessionAffinity

`func (o *KubernetesService) GetSessionAffinity() string`

GetSessionAffinity returns the SessionAffinity field if non-nil, zero value otherwise.

### GetSessionAffinityOk

`func (o *KubernetesService) GetSessionAffinityOk() (*string, bool)`

GetSessionAffinityOk returns a tuple with the SessionAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionAffinity

`func (o *KubernetesService) SetSessionAffinity(v string)`

SetSessionAffinity sets SessionAffinity field to given value.


### SetSessionAffinityNil

`func (o *KubernetesService) SetSessionAffinityNil(b bool)`

 SetSessionAffinityNil sets the value for SessionAffinity to be an explicit nil

### UnsetSessionAffinity
`func (o *KubernetesService) UnsetSessionAffinity()`

UnsetSessionAffinity ensures that no value is present for SessionAffinity, not even an explicit nil
### GetUid

`func (o *KubernetesService) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesService) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesService) SetUid(v string)`

SetUid sets Uid field to given value.


### GetYaml

`func (o *KubernetesService) GetYaml() string`

GetYaml returns the Yaml field if non-nil, zero value otherwise.

### GetYamlOk

`func (o *KubernetesService) GetYamlOk() (*string, bool)`

GetYamlOk returns a tuple with the Yaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYaml

`func (o *KubernetesService) SetYaml(v string)`

SetYaml sets Yaml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


