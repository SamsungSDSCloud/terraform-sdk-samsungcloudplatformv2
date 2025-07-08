# KubernetesServiceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**ClusterIp** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**ExternalIps** | **[]string** | External Ips | 
**Name** | **string** | Service Name | 
**NamespaceName** | **string** | Namespace Name | 
**Ports** | [**[]KubernetesServicePort**](KubernetesServicePort.md) | Ports | 
**ServiceType** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewKubernetesServiceSummary

`func NewKubernetesServiceSummary(age string, clusterId string, clusterIp NullableString, createdAt time.Time, externalIps []string, name string, namespaceName string, ports []KubernetesServicePort, serviceType NullableString, uid string, ) *KubernetesServiceSummary`

NewKubernetesServiceSummary instantiates a new KubernetesServiceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesServiceSummaryWithDefaults

`func NewKubernetesServiceSummaryWithDefaults() *KubernetesServiceSummary`

NewKubernetesServiceSummaryWithDefaults instantiates a new KubernetesServiceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *KubernetesServiceSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *KubernetesServiceSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *KubernetesServiceSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *KubernetesServiceSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *KubernetesServiceSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *KubernetesServiceSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetClusterIp

`func (o *KubernetesServiceSummary) GetClusterIp() string`

GetClusterIp returns the ClusterIp field if non-nil, zero value otherwise.

### GetClusterIpOk

`func (o *KubernetesServiceSummary) GetClusterIpOk() (*string, bool)`

GetClusterIpOk returns a tuple with the ClusterIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIp

`func (o *KubernetesServiceSummary) SetClusterIp(v string)`

SetClusterIp sets ClusterIp field to given value.


### SetClusterIpNil

`func (o *KubernetesServiceSummary) SetClusterIpNil(b bool)`

 SetClusterIpNil sets the value for ClusterIp to be an explicit nil

### UnsetClusterIp
`func (o *KubernetesServiceSummary) UnsetClusterIp()`

UnsetClusterIp ensures that no value is present for ClusterIp, not even an explicit nil
### GetCreatedAt

`func (o *KubernetesServiceSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KubernetesServiceSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KubernetesServiceSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExternalIps

`func (o *KubernetesServiceSummary) GetExternalIps() []string`

GetExternalIps returns the ExternalIps field if non-nil, zero value otherwise.

### GetExternalIpsOk

`func (o *KubernetesServiceSummary) GetExternalIpsOk() (*[]string, bool)`

GetExternalIpsOk returns a tuple with the ExternalIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIps

`func (o *KubernetesServiceSummary) SetExternalIps(v []string)`

SetExternalIps sets ExternalIps field to given value.


### GetName

`func (o *KubernetesServiceSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesServiceSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesServiceSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *KubernetesServiceSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *KubernetesServiceSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *KubernetesServiceSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetPorts

`func (o *KubernetesServiceSummary) GetPorts() []KubernetesServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *KubernetesServiceSummary) GetPortsOk() (*[]KubernetesServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *KubernetesServiceSummary) SetPorts(v []KubernetesServicePort)`

SetPorts sets Ports field to given value.


### GetServiceType

`func (o *KubernetesServiceSummary) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *KubernetesServiceSummary) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *KubernetesServiceSummary) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### SetServiceTypeNil

`func (o *KubernetesServiceSummary) SetServiceTypeNil(b bool)`

 SetServiceTypeNil sets the value for ServiceType to be an explicit nil

### UnsetServiceType
`func (o *KubernetesServiceSummary) UnsetServiceType()`

UnsetServiceType ensures that no value is present for ServiceType, not even an explicit nil
### GetUid

`func (o *KubernetesServiceSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *KubernetesServiceSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *KubernetesServiceSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


