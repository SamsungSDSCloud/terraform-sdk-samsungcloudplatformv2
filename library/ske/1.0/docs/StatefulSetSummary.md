# StatefulSetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**Name** | **string** | Stateful Set Name | 
**NamespaceName** | **string** | Namespace Name | 
**ReadyReplicas** | **int32** | Ready Replicas | 
**SpecReplicas** | **int32** | Spec Replicas | 
**StatusReplicas** | **int32** | Status Replicas | 
**Uid** | **string** | Resource ID | 

## Methods

### NewStatefulSetSummary

`func NewStatefulSetSummary(age string, clusterId string, createdAt time.Time, name string, namespaceName string, readyReplicas int32, specReplicas int32, statusReplicas int32, uid string, ) *StatefulSetSummary`

NewStatefulSetSummary instantiates a new StatefulSetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatefulSetSummaryWithDefaults

`func NewStatefulSetSummaryWithDefaults() *StatefulSetSummary`

NewStatefulSetSummaryWithDefaults instantiates a new StatefulSetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *StatefulSetSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *StatefulSetSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *StatefulSetSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *StatefulSetSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *StatefulSetSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *StatefulSetSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *StatefulSetSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StatefulSetSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StatefulSetSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *StatefulSetSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatefulSetSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatefulSetSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *StatefulSetSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *StatefulSetSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *StatefulSetSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetReadyReplicas

`func (o *StatefulSetSummary) GetReadyReplicas() int32`

GetReadyReplicas returns the ReadyReplicas field if non-nil, zero value otherwise.

### GetReadyReplicasOk

`func (o *StatefulSetSummary) GetReadyReplicasOk() (*int32, bool)`

GetReadyReplicasOk returns a tuple with the ReadyReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyReplicas

`func (o *StatefulSetSummary) SetReadyReplicas(v int32)`

SetReadyReplicas sets ReadyReplicas field to given value.


### GetSpecReplicas

`func (o *StatefulSetSummary) GetSpecReplicas() int32`

GetSpecReplicas returns the SpecReplicas field if non-nil, zero value otherwise.

### GetSpecReplicasOk

`func (o *StatefulSetSummary) GetSpecReplicasOk() (*int32, bool)`

GetSpecReplicasOk returns a tuple with the SpecReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecReplicas

`func (o *StatefulSetSummary) SetSpecReplicas(v int32)`

SetSpecReplicas sets SpecReplicas field to given value.


### GetStatusReplicas

`func (o *StatefulSetSummary) GetStatusReplicas() int32`

GetStatusReplicas returns the StatusReplicas field if non-nil, zero value otherwise.

### GetStatusReplicasOk

`func (o *StatefulSetSummary) GetStatusReplicasOk() (*int32, bool)`

GetStatusReplicasOk returns a tuple with the StatusReplicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReplicas

`func (o *StatefulSetSummary) SetStatusReplicas(v int32)`

SetStatusReplicas sets StatusReplicas field to given value.


### GetUid

`func (o *StatefulSetSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *StatefulSetSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *StatefulSetSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


