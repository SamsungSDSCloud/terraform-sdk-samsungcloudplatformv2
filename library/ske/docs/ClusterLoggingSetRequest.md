# ClusterLoggingSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudLoggingEnabled** | **bool** | Cloud Logging Enabled | 

## Methods

### NewClusterLoggingSetRequest

`func NewClusterLoggingSetRequest(cloudLoggingEnabled bool, ) *ClusterLoggingSetRequest`

NewClusterLoggingSetRequest instantiates a new ClusterLoggingSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLoggingSetRequestWithDefaults

`func NewClusterLoggingSetRequestWithDefaults() *ClusterLoggingSetRequest`

NewClusterLoggingSetRequestWithDefaults instantiates a new ClusterLoggingSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudLoggingEnabled

`func (o *ClusterLoggingSetRequest) GetCloudLoggingEnabled() bool`

GetCloudLoggingEnabled returns the CloudLoggingEnabled field if non-nil, zero value otherwise.

### GetCloudLoggingEnabledOk

`func (o *ClusterLoggingSetRequest) GetCloudLoggingEnabledOk() (*bool, bool)`

GetCloudLoggingEnabledOk returns a tuple with the CloudLoggingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudLoggingEnabled

`func (o *ClusterLoggingSetRequest) SetCloudLoggingEnabled(v bool)`

SetCloudLoggingEnabled sets CloudLoggingEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


