# NetworkLoggingStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**BucketName** | **string** | Storage Bucket Name | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | Network Logging Storage ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ResourceType** | [**NetworkLoggingResourceType**](NetworkLoggingResourceType.md) | Resource Type | 

## Methods

### NewNetworkLoggingStorage

`func NewNetworkLoggingStorage(accountId string, bucketName string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, resourceType NetworkLoggingResourceType, ) *NetworkLoggingStorage`

NewNetworkLoggingStorage instantiates a new NetworkLoggingStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoggingStorageWithDefaults

`func NewNetworkLoggingStorageWithDefaults() *NetworkLoggingStorage`

NewNetworkLoggingStorageWithDefaults instantiates a new NetworkLoggingStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NetworkLoggingStorage) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NetworkLoggingStorage) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NetworkLoggingStorage) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBucketName

`func (o *NetworkLoggingStorage) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *NetworkLoggingStorage) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *NetworkLoggingStorage) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetCreatedAt

`func (o *NetworkLoggingStorage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkLoggingStorage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkLoggingStorage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *NetworkLoggingStorage) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NetworkLoggingStorage) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NetworkLoggingStorage) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *NetworkLoggingStorage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLoggingStorage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLoggingStorage) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *NetworkLoggingStorage) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NetworkLoggingStorage) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NetworkLoggingStorage) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *NetworkLoggingStorage) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *NetworkLoggingStorage) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *NetworkLoggingStorage) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetResourceType

`func (o *NetworkLoggingStorage) GetResourceType() NetworkLoggingResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NetworkLoggingStorage) GetResourceTypeOk() (*NetworkLoggingResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NetworkLoggingStorage) SetResourceType(v NetworkLoggingResourceType)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


