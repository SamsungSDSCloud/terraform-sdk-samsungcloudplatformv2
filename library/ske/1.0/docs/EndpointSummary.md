# EndpointSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**Name** | **string** | Endpoint Name | 
**NamespaceName** | **string** | Namespace Name | 
**Subsets** | [**[]Subset**](Subset.md) | Subsets | 
**Uid** | **string** | Resource ID | 

## Methods

### NewEndpointSummary

`func NewEndpointSummary(age string, clusterId string, createdAt time.Time, name string, namespaceName string, subsets []Subset, uid string, ) *EndpointSummary`

NewEndpointSummary instantiates a new EndpointSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointSummaryWithDefaults

`func NewEndpointSummaryWithDefaults() *EndpointSummary`

NewEndpointSummaryWithDefaults instantiates a new EndpointSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *EndpointSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *EndpointSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *EndpointSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *EndpointSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EndpointSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EndpointSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *EndpointSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EndpointSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EndpointSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *EndpointSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *EndpointSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *EndpointSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *EndpointSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetSubsets

`func (o *EndpointSummary) GetSubsets() []Subset`

GetSubsets returns the Subsets field if non-nil, zero value otherwise.

### GetSubsetsOk

`func (o *EndpointSummary) GetSubsetsOk() (*[]Subset, bool)`

GetSubsetsOk returns a tuple with the Subsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsets

`func (o *EndpointSummary) SetSubsets(v []Subset)`

SetSubsets sets Subsets field to given value.


### GetUid

`func (o *EndpointSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *EndpointSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *EndpointSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


