# NetworkLoggingStorageCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | **string** | Storage Bucket Name | 
**ResourceType** | [**NetworkLoggingResourceType**](NetworkLoggingResourceType.md) | Resource Type | 

## Methods

### NewNetworkLoggingStorageCreateRequest

`func NewNetworkLoggingStorageCreateRequest(bucketName string, resourceType NetworkLoggingResourceType, ) *NetworkLoggingStorageCreateRequest`

NewNetworkLoggingStorageCreateRequest instantiates a new NetworkLoggingStorageCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoggingStorageCreateRequestWithDefaults

`func NewNetworkLoggingStorageCreateRequestWithDefaults() *NetworkLoggingStorageCreateRequest`

NewNetworkLoggingStorageCreateRequestWithDefaults instantiates a new NetworkLoggingStorageCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *NetworkLoggingStorageCreateRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *NetworkLoggingStorageCreateRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *NetworkLoggingStorageCreateRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetResourceType

`func (o *NetworkLoggingStorageCreateRequest) GetResourceType() NetworkLoggingResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NetworkLoggingStorageCreateRequest) GetResourceTypeOk() (*NetworkLoggingResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NetworkLoggingStorageCreateRequest) SetResourceType(v NetworkLoggingResourceType)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


