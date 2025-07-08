# NetworkLoggingConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**BucketName** | **string** | Storage Bucket Name | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DownInterface** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Network Logging Configuration ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ResourceId** | **string** | Resource ID | 
**ResourceName** | **string** | Resource Name | 
**ResourceType** | [**NetworkLoggingResourceType**](NetworkLoggingResourceType.md) | Resource Type | 
**SecurityGroupLogId** | Pointer to **NullableString** |  | [optional] 
**UpInterface** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewNetworkLoggingConfiguration

`func NewNetworkLoggingConfiguration(accountId string, bucketName string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, resourceId string, resourceName string, resourceType NetworkLoggingResourceType, ) *NetworkLoggingConfiguration`

NewNetworkLoggingConfiguration instantiates a new NetworkLoggingConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkLoggingConfigurationWithDefaults

`func NewNetworkLoggingConfigurationWithDefaults() *NetworkLoggingConfiguration`

NewNetworkLoggingConfigurationWithDefaults instantiates a new NetworkLoggingConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *NetworkLoggingConfiguration) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *NetworkLoggingConfiguration) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *NetworkLoggingConfiguration) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetBucketName

`func (o *NetworkLoggingConfiguration) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *NetworkLoggingConfiguration) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *NetworkLoggingConfiguration) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetCreatedAt

`func (o *NetworkLoggingConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NetworkLoggingConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NetworkLoggingConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *NetworkLoggingConfiguration) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *NetworkLoggingConfiguration) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *NetworkLoggingConfiguration) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDownInterface

`func (o *NetworkLoggingConfiguration) GetDownInterface() string`

GetDownInterface returns the DownInterface field if non-nil, zero value otherwise.

### GetDownInterfaceOk

`func (o *NetworkLoggingConfiguration) GetDownInterfaceOk() (*string, bool)`

GetDownInterfaceOk returns a tuple with the DownInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownInterface

`func (o *NetworkLoggingConfiguration) SetDownInterface(v string)`

SetDownInterface sets DownInterface field to given value.

### HasDownInterface

`func (o *NetworkLoggingConfiguration) HasDownInterface() bool`

HasDownInterface returns a boolean if a field has been set.

### SetDownInterfaceNil

`func (o *NetworkLoggingConfiguration) SetDownInterfaceNil(b bool)`

 SetDownInterfaceNil sets the value for DownInterface to be an explicit nil

### UnsetDownInterface
`func (o *NetworkLoggingConfiguration) UnsetDownInterface()`

UnsetDownInterface ensures that no value is present for DownInterface, not even an explicit nil
### GetId

`func (o *NetworkLoggingConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkLoggingConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkLoggingConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *NetworkLoggingConfiguration) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *NetworkLoggingConfiguration) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *NetworkLoggingConfiguration) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *NetworkLoggingConfiguration) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *NetworkLoggingConfiguration) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *NetworkLoggingConfiguration) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetResourceId

`func (o *NetworkLoggingConfiguration) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *NetworkLoggingConfiguration) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *NetworkLoggingConfiguration) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceName

`func (o *NetworkLoggingConfiguration) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *NetworkLoggingConfiguration) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *NetworkLoggingConfiguration) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### GetResourceType

`func (o *NetworkLoggingConfiguration) GetResourceType() NetworkLoggingResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *NetworkLoggingConfiguration) GetResourceTypeOk() (*NetworkLoggingResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *NetworkLoggingConfiguration) SetResourceType(v NetworkLoggingResourceType)`

SetResourceType sets ResourceType field to given value.


### GetSecurityGroupLogId

`func (o *NetworkLoggingConfiguration) GetSecurityGroupLogId() string`

GetSecurityGroupLogId returns the SecurityGroupLogId field if non-nil, zero value otherwise.

### GetSecurityGroupLogIdOk

`func (o *NetworkLoggingConfiguration) GetSecurityGroupLogIdOk() (*string, bool)`

GetSecurityGroupLogIdOk returns a tuple with the SecurityGroupLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupLogId

`func (o *NetworkLoggingConfiguration) SetSecurityGroupLogId(v string)`

SetSecurityGroupLogId sets SecurityGroupLogId field to given value.

### HasSecurityGroupLogId

`func (o *NetworkLoggingConfiguration) HasSecurityGroupLogId() bool`

HasSecurityGroupLogId returns a boolean if a field has been set.

### SetSecurityGroupLogIdNil

`func (o *NetworkLoggingConfiguration) SetSecurityGroupLogIdNil(b bool)`

 SetSecurityGroupLogIdNil sets the value for SecurityGroupLogId to be an explicit nil

### UnsetSecurityGroupLogId
`func (o *NetworkLoggingConfiguration) UnsetSecurityGroupLogId()`

UnsetSecurityGroupLogId ensures that no value is present for SecurityGroupLogId, not even an explicit nil
### GetUpInterface

`func (o *NetworkLoggingConfiguration) GetUpInterface() string`

GetUpInterface returns the UpInterface field if non-nil, zero value otherwise.

### GetUpInterfaceOk

`func (o *NetworkLoggingConfiguration) GetUpInterfaceOk() (*string, bool)`

GetUpInterfaceOk returns a tuple with the UpInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpInterface

`func (o *NetworkLoggingConfiguration) SetUpInterface(v string)`

SetUpInterface sets UpInterface field to given value.

### HasUpInterface

`func (o *NetworkLoggingConfiguration) HasUpInterface() bool`

HasUpInterface returns a boolean if a field has been set.

### SetUpInterfaceNil

`func (o *NetworkLoggingConfiguration) SetUpInterfaceNil(b bool)`

 SetUpInterfaceNil sets the value for UpInterface to be an explicit nil

### UnsetUpInterface
`func (o *NetworkLoggingConfiguration) UnsetUpInterface()`

UnsetUpInterface ensures that no value is present for UpInterface, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


