# NamespaceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**Name** | **string** | Namespace Name | 
**NamespaceStatus** | **NullableString** |  | 
**Uid** | **string** | Resource ID | 

## Methods

### NewNamespaceSummary

`func NewNamespaceSummary(age string, clusterId string, createdAt time.Time, name string, namespaceStatus NullableString, uid string, ) *NamespaceSummary`

NewNamespaceSummary instantiates a new NamespaceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceSummaryWithDefaults

`func NewNamespaceSummaryWithDefaults() *NamespaceSummary`

NewNamespaceSummaryWithDefaults instantiates a new NamespaceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *NamespaceSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *NamespaceSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *NamespaceSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *NamespaceSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NamespaceSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NamespaceSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *NamespaceSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NamespaceSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NamespaceSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *NamespaceSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NamespaceSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NamespaceSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceStatus

`func (o *NamespaceSummary) GetNamespaceStatus() string`

GetNamespaceStatus returns the NamespaceStatus field if non-nil, zero value otherwise.

### GetNamespaceStatusOk

`func (o *NamespaceSummary) GetNamespaceStatusOk() (*string, bool)`

GetNamespaceStatusOk returns a tuple with the NamespaceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceStatus

`func (o *NamespaceSummary) SetNamespaceStatus(v string)`

SetNamespaceStatus sets NamespaceStatus field to given value.


### SetNamespaceStatusNil

`func (o *NamespaceSummary) SetNamespaceStatusNil(b bool)`

 SetNamespaceStatusNil sets the value for NamespaceStatus to be an explicit nil

### UnsetNamespaceStatus
`func (o *NamespaceSummary) UnsetNamespaceStatus()`

UnsetNamespaceStatus ensures that no value is present for NamespaceStatus, not even an explicit nil
### GetUid

`func (o *NamespaceSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *NamespaceSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *NamespaceSummary) SetUid(v string)`

SetUid sets Uid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


