# ConfigMapSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**ConfigMapDataCount** | **int32** | Config Map Data Count | 
**CreatedAt** | **time.Time** | Created At | 
**Name** | **string** | Config Map Name | 
**NamespaceName** | **string** | Namespace Name | 
**Uid** | **string** | Resource ID | 

## Methods

### NewConfigMapSummary

`func NewConfigMapSummary(age string, clusterId string, configMapDataCount int32, createdAt time.Time, name string, namespaceName string, uid string, ) *ConfigMapSummary`

NewConfigMapSummary instantiates a new ConfigMapSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigMapSummaryWithDefaults

`func NewConfigMapSummaryWithDefaults() *ConfigMapSummary`

NewConfigMapSummaryWithDefaults instantiates a new ConfigMapSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *ConfigMapSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ConfigMapSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ConfigMapSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *ConfigMapSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *ConfigMapSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *ConfigMapSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetConfigMapDataCount

`func (o *ConfigMapSummary) GetConfigMapDataCount() int32`

GetConfigMapDataCount returns the ConfigMapDataCount field if non-nil, zero value otherwise.

### GetConfigMapDataCountOk

`func (o *ConfigMapSummary) GetConfigMapDataCountOk() (*int32, bool)`

GetConfigMapDataCountOk returns a tuple with the ConfigMapDataCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapDataCount

`func (o *ConfigMapSummary) SetConfigMapDataCount(v int32)`

SetConfigMapDataCount sets ConfigMapDataCount field to given value.


### GetCreatedAt

`func (o *ConfigMapSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConfigMapSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConfigMapSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *ConfigMapSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigMapSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigMapSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *ConfigMapSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *ConfigMapSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *ConfigMapSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetUid

`func (o *ConfigMapSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *ConfigMapSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *ConfigMapSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


