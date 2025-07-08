# DaemonSetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Age** | **string** | Age | 
**ClusterId** | **string** | Cluster ID | 
**CreatedAt** | **time.Time** | Created At | 
**CurrentNumberScheduled** | **int32** | Current Number Scheduled | 
**DesiredNumberScheduled** | **int32** | Desired Number Scheduled | 
**Name** | **string** | Daemon Set Name | 
**NamespaceName** | **string** | Namespace Name | 
**NodeSelector** | **[]string** | Node Selector | 
**NumberAvailable** | **int32** | Number Available | 
**NumberReady** | **int32** | Number Ready | 
**Uid** | **string** | Resource ID | 
**UpdatedNumberScheduled** | **int32** | Updated Number Scheduled | 

## Methods

### NewDaemonSetSummary

`func NewDaemonSetSummary(age string, clusterId string, createdAt time.Time, currentNumberScheduled int32, desiredNumberScheduled int32, name string, namespaceName string, nodeSelector []string, numberAvailable int32, numberReady int32, uid string, updatedNumberScheduled int32, ) *DaemonSetSummary`

NewDaemonSetSummary instantiates a new DaemonSetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaemonSetSummaryWithDefaults

`func NewDaemonSetSummaryWithDefaults() *DaemonSetSummary`

NewDaemonSetSummaryWithDefaults instantiates a new DaemonSetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAge

`func (o *DaemonSetSummary) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *DaemonSetSummary) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *DaemonSetSummary) SetAge(v string)`

SetAge sets Age field to given value.


### GetClusterId

`func (o *DaemonSetSummary) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *DaemonSetSummary) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *DaemonSetSummary) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetCreatedAt

`func (o *DaemonSetSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaemonSetSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaemonSetSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrentNumberScheduled

`func (o *DaemonSetSummary) GetCurrentNumberScheduled() int32`

GetCurrentNumberScheduled returns the CurrentNumberScheduled field if non-nil, zero value otherwise.

### GetCurrentNumberScheduledOk

`func (o *DaemonSetSummary) GetCurrentNumberScheduledOk() (*int32, bool)`

GetCurrentNumberScheduledOk returns a tuple with the CurrentNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNumberScheduled

`func (o *DaemonSetSummary) SetCurrentNumberScheduled(v int32)`

SetCurrentNumberScheduled sets CurrentNumberScheduled field to given value.


### GetDesiredNumberScheduled

`func (o *DaemonSetSummary) GetDesiredNumberScheduled() int32`

GetDesiredNumberScheduled returns the DesiredNumberScheduled field if non-nil, zero value otherwise.

### GetDesiredNumberScheduledOk

`func (o *DaemonSetSummary) GetDesiredNumberScheduledOk() (*int32, bool)`

GetDesiredNumberScheduledOk returns a tuple with the DesiredNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredNumberScheduled

`func (o *DaemonSetSummary) SetDesiredNumberScheduled(v int32)`

SetDesiredNumberScheduled sets DesiredNumberScheduled field to given value.


### GetName

`func (o *DaemonSetSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaemonSetSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaemonSetSummary) SetName(v string)`

SetName sets Name field to given value.


### GetNamespaceName

`func (o *DaemonSetSummary) GetNamespaceName() string`

GetNamespaceName returns the NamespaceName field if non-nil, zero value otherwise.

### GetNamespaceNameOk

`func (o *DaemonSetSummary) GetNamespaceNameOk() (*string, bool)`

GetNamespaceNameOk returns a tuple with the NamespaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceName

`func (o *DaemonSetSummary) SetNamespaceName(v string)`

SetNamespaceName sets NamespaceName field to given value.


### GetNodeSelector

`func (o *DaemonSetSummary) GetNodeSelector() []string`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *DaemonSetSummary) GetNodeSelectorOk() (*[]string, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *DaemonSetSummary) SetNodeSelector(v []string)`

SetNodeSelector sets NodeSelector field to given value.


### GetNumberAvailable

`func (o *DaemonSetSummary) GetNumberAvailable() int32`

GetNumberAvailable returns the NumberAvailable field if non-nil, zero value otherwise.

### GetNumberAvailableOk

`func (o *DaemonSetSummary) GetNumberAvailableOk() (*int32, bool)`

GetNumberAvailableOk returns a tuple with the NumberAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberAvailable

`func (o *DaemonSetSummary) SetNumberAvailable(v int32)`

SetNumberAvailable sets NumberAvailable field to given value.


### GetNumberReady

`func (o *DaemonSetSummary) GetNumberReady() int32`

GetNumberReady returns the NumberReady field if non-nil, zero value otherwise.

### GetNumberReadyOk

`func (o *DaemonSetSummary) GetNumberReadyOk() (*int32, bool)`

GetNumberReadyOk returns a tuple with the NumberReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberReady

`func (o *DaemonSetSummary) SetNumberReady(v int32)`

SetNumberReady sets NumberReady field to given value.


### GetUid

`func (o *DaemonSetSummary) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *DaemonSetSummary) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *DaemonSetSummary) SetUid(v string)`

SetUid sets Uid field to given value.


### GetUpdatedNumberScheduled

`func (o *DaemonSetSummary) GetUpdatedNumberScheduled() int32`

GetUpdatedNumberScheduled returns the UpdatedNumberScheduled field if non-nil, zero value otherwise.

### GetUpdatedNumberScheduledOk

`func (o *DaemonSetSummary) GetUpdatedNumberScheduledOk() (*int32, bool)`

GetUpdatedNumberScheduledOk returns a tuple with the UpdatedNumberScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedNumberScheduled

`func (o *DaemonSetSummary) SetUpdatedNumberScheduled(v int32)`

SetUpdatedNumberScheduled sets UpdatedNumberScheduled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


