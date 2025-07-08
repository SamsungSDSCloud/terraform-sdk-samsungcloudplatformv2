# IngressSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**IngressClass** | **NullableString** |  | 
**IngressHost** | **[]string** | Ingress Host | 
**IngressIp** | **[]string** | Ingress Ip | 
**IngressPort** | **string** | Ingress Port | 
**Name** | **string** | Ingress Name | 
**NamespaceName** | **string** | Namespace Name | 
**Uid** | **string** | Resource ID | 

## Methods

### NewIngressSummary

`func NewIngressSummary(age string, clusterId string, createdAt time.Time, ingressClass NullableString, ingressHost []string, ingressIp []string, ingressPort string, name string, namespaceName string, uid string, ) *IngressSummary`

NewIngressSummary instantiates a new IngressSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngressSummaryWithDefaults

`func NewIngressSummaryWithDefaults() *IngressSummary`

NewIngressSummaryWithDefaults instantiates a new IngressSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *IngressSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *IngressSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *IngressSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *IngressSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *IngressSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *IngressSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *IngressSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngressSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngressSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIngressClass

`func (o *IngressSummary) GetIngressClass() string`

GetIngressClass returns the IngressClass field if non-nil, zero value otherwise.

### GetIngressClassOk

`func (o *IngressSummary) GetIngressClassOk() (*string, bool)`

GetIngressClassOk returns a tuple with the IngressClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressClass

`func (o *IngressSummary) SetIngressClass(v string)`

SetIngressClass sets IngressClass field to given value.


### SetIngressClassNil

`func (o *IngressSummary) SetIngressClassNil(b bool)`

 SetIngressClassNil sets the value for IngressClass to be an explicit nil

### UnsetIngressClass
`func (o *IngressSummary) UnsetIngressClass()`

UnsetIngressClass ensures that no value is present for IngressClass, not even an explicit nil
### GetIngressHost

`func (o *IngressSummary) GetIngressHost() []string`

GetIngressHost returns the IngressHost field if non-nil, zero value otherwise.

### GetIngressHostOk

`func (o *IngressSummary) GetIngressHostOk() (*[]string, bool)`

GetIngressHostOk returns a tuple with the IngressHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressHost

`func (o *IngressSummary) SetIngressHost(v []string)`

SetIngressHost sets IngressHost field to given value.


### GetIngressIp

`func (o *IngressSummary) GetIngressIp() []string`

GetIngressIp returns the IngressIp field if non-nil, zero value otherwise.

### GetIngressIpOk

`func (o *IngressSummary) GetIngressIpOk() (*[]string, bool)`

GetIngressIpOk returns a tuple with the IngressIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressIp

`func (o *IngressSummary) SetIngressIp(v []string)`

SetIngressIp sets IngressIp field to given value.


### GetIngressPort

`func (o *IngressSummary) GetIngressPort() string`

GetIngressPort returns the IngressPort field if non-nil, zero value otherwise.

### GetIngressPortOk

`func (o *IngressSummary) GetIngressPortOk() (*string, bool)`

GetIngressPortOk returns a tuple with the IngressPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressPort

`func (o *IngressSummary) SetIngressPort(v string)`

SetIngressPort sets IngressPort field to given value.


### GetName

`func (o *IngressSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngressSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngressSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *IngressSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *IngressSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *IngressSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetUid

`func (o *IngressSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *IngressSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *IngressSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


