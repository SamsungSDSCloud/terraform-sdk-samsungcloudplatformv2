# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** |  | 
**CreatedAt** | **NullableTime** |  | 
**CreatedBy** | **NullableString** |  | 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ProductName** | **NullableString** |  | 
**Region** | **NullableString** |  | 
**ResourceName** | **NullableString** |  | 
**ResourceType** | **NullableString** |  | 
**ResourceTypeDisplayName** | **NullableString** |  | 
**Srn** | **NullableString** |  | 
**State** | **NullableString** |  | 
**Tags** | **[]map[string]interface{}** |  | 

## Methods

### NewResource

`func NewResource(accountId NullableString, createdAt NullableTime, createdBy NullableString, id string, modifiedAt time.Time, modifiedBy string, productName NullableString, region NullableString, resourceName NullableString, resourceType NullableString, resourceTypeDisplayName NullableString, srn NullableString, state NullableString, tags []map[string]interface{}, ) *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Resource) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Resource) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Resource) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *Resource) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Resource) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetCreatedAt

`func (o *Resource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Resource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Resource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Resource) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Resource) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedBy

`func (o *Resource) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Resource) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Resource) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### SetCreatedByNil

`func (o *Resource) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *Resource) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *Resource) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Resource) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Resource) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Resource) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Resource) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Resource) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetProductName

`func (o *Resource) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *Resource) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *Resource) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### SetProductNameNil

`func (o *Resource) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *Resource) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetRegion

`func (o *Resource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Resource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Resource) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *Resource) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Resource) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetResourceName

`func (o *Resource) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *Resource) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *Resource) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.


### SetResourceNameNil

`func (o *Resource) SetResourceNameNil(b bool)`

 SetResourceNameNil sets the value for ResourceName to be an explicit nil

### UnsetResourceName
`func (o *Resource) UnsetResourceName()`

UnsetResourceName ensures that no value is present for ResourceName, not even an explicit nil
### GetResourceType

`func (o *Resource) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *Resource) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *Resource) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### SetResourceTypeNil

`func (o *Resource) SetResourceTypeNil(b bool)`

 SetResourceTypeNil sets the value for ResourceType to be an explicit nil

### UnsetResourceType
`func (o *Resource) UnsetResourceType()`

UnsetResourceType ensures that no value is present for ResourceType, not even an explicit nil
### GetResourceTypeDisplayName

`func (o *Resource) GetResourceTypeDisplayName() string`

GetResourceTypeDisplayName returns the ResourceTypeDisplayName field if non-nil, zero value otherwise.

### GetResourceTypeDisplayNameOk

`func (o *Resource) GetResourceTypeDisplayNameOk() (*string, bool)`

GetResourceTypeDisplayNameOk returns a tuple with the ResourceTypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeDisplayName

`func (o *Resource) SetResourceTypeDisplayName(v string)`

SetResourceTypeDisplayName sets ResourceTypeDisplayName field to given value.


### SetResourceTypeDisplayNameNil

`func (o *Resource) SetResourceTypeDisplayNameNil(b bool)`

 SetResourceTypeDisplayNameNil sets the value for ResourceTypeDisplayName to be an explicit nil

### UnsetResourceTypeDisplayName
`func (o *Resource) UnsetResourceTypeDisplayName()`

UnsetResourceTypeDisplayName ensures that no value is present for ResourceTypeDisplayName, not even an explicit nil
### GetSrn

`func (o *Resource) GetSrn() string`

GetSrn returns the Srn field if non-nil, zero value otherwise.

### GetSrnOk

`func (o *Resource) GetSrnOk() (*string, bool)`

GetSrnOk returns a tuple with the Srn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrn

`func (o *Resource) SetSrn(v string)`

SetSrn sets Srn field to given value.


### SetSrnNil

`func (o *Resource) SetSrnNil(b bool)`

 SetSrnNil sets the value for Srn to be an explicit nil

### UnsetSrn
`func (o *Resource) UnsetSrn()`

UnsetSrn ensures that no value is present for Srn, not even an explicit nil
### GetState

`func (o *Resource) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Resource) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Resource) SetState(v string)`

SetState sets State field to given value.


### SetStateNil

`func (o *Resource) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Resource) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTags

`func (o *Resource) GetTags() []map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Resource) GetTagsOk() (*[]map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Resource) SetTags(v []map[string]interface{})`

SetTags sets Tags field to given value.


### SetTagsNil

`func (o *Resource) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Resource) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


