# TrailCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** |  | [optional] 
**BucketName** | **string** | Bucket name | 
**BucketRegion** | **string** | Bucket region | 
**LogTypeTotalYn** | Pointer to **NullableString** |  | [optional] 
**LogVerificationYn** | Pointer to **NullableString** |  | [optional] 
**RegionNames** | Pointer to **[]interface{}** | Region name list | [optional] [default to []]
**RegionTotalYn** | Pointer to **NullableString** |  | [optional] 
**ResourceTypeTotalYn** | Pointer to **NullableString** |  | [optional] 
**TagCreateRequests** | Pointer to **[]map[string]string** |  | [optional] 
**TargetLogTypes** | Pointer to **[]interface{}** | Target log type list | [optional] [default to []]
**TargetResourceTypes** | Pointer to **[]interface{}** | Target resource type list | [optional] [default to []]
**TargetUsers** | Pointer to **[]interface{}** | Target user list | [optional] [default to []]
**TrailDescription** | Pointer to **NullableString** |  | [optional] 
**TrailName** | **string** | Trail name | 
**TrailSaveType** | **string** | Trail save type | 
**UserTotalYn** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrailCreateRequest

`func NewTrailCreateRequest(bucketName string, bucketRegion string, trailName string, trailSaveType string, ) *TrailCreateRequest`

NewTrailCreateRequest instantiates a new TrailCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailCreateRequestWithDefaults

`func NewTrailCreateRequestWithDefaults() *TrailCreateRequest`

NewTrailCreateRequestWithDefaults instantiates a new TrailCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TrailCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TrailCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TrailCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TrailCreateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TrailCreateRequest) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TrailCreateRequest) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetBucketName

`func (o *TrailCreateRequest) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *TrailCreateRequest) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *TrailCreateRequest) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetBucketRegion

`func (o *TrailCreateRequest) GetBucketRegion() string`

GetBucketRegion returns the BucketRegion field if non-nil, zero value otherwise.

### GetBucketRegionOk

`func (o *TrailCreateRequest) GetBucketRegionOk() (*string, bool)`

GetBucketRegionOk returns a tuple with the BucketRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRegion

`func (o *TrailCreateRequest) SetBucketRegion(v string)`

SetBucketRegion sets BucketRegion field to given value.


### GetLogTypeTotalYn

`func (o *TrailCreateRequest) GetLogTypeTotalYn() string`

GetLogTypeTotalYn returns the LogTypeTotalYn field if non-nil, zero value otherwise.

### GetLogTypeTotalYnOk

`func (o *TrailCreateRequest) GetLogTypeTotalYnOk() (*string, bool)`

GetLogTypeTotalYnOk returns a tuple with the LogTypeTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypeTotalYn

`func (o *TrailCreateRequest) SetLogTypeTotalYn(v string)`

SetLogTypeTotalYn sets LogTypeTotalYn field to given value.

### HasLogTypeTotalYn

`func (o *TrailCreateRequest) HasLogTypeTotalYn() bool`

HasLogTypeTotalYn returns a boolean if a field has been set.

### SetLogTypeTotalYnNil

`func (o *TrailCreateRequest) SetLogTypeTotalYnNil(b bool)`

 SetLogTypeTotalYnNil sets the value for LogTypeTotalYn to be an explicit nil

### UnsetLogTypeTotalYn
`func (o *TrailCreateRequest) UnsetLogTypeTotalYn()`

UnsetLogTypeTotalYn ensures that no value is present for LogTypeTotalYn, not even an explicit nil
### GetLogVerificationYn

`func (o *TrailCreateRequest) GetLogVerificationYn() string`

GetLogVerificationYn returns the LogVerificationYn field if non-nil, zero value otherwise.

### GetLogVerificationYnOk

`func (o *TrailCreateRequest) GetLogVerificationYnOk() (*string, bool)`

GetLogVerificationYnOk returns a tuple with the LogVerificationYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogVerificationYn

`func (o *TrailCreateRequest) SetLogVerificationYn(v string)`

SetLogVerificationYn sets LogVerificationYn field to given value.

### HasLogVerificationYn

`func (o *TrailCreateRequest) HasLogVerificationYn() bool`

HasLogVerificationYn returns a boolean if a field has been set.

### SetLogVerificationYnNil

`func (o *TrailCreateRequest) SetLogVerificationYnNil(b bool)`

 SetLogVerificationYnNil sets the value for LogVerificationYn to be an explicit nil

### UnsetLogVerificationYn
`func (o *TrailCreateRequest) UnsetLogVerificationYn()`

UnsetLogVerificationYn ensures that no value is present for LogVerificationYn, not even an explicit nil
### GetRegionNames

`func (o *TrailCreateRequest) GetRegionNames() []interface{}`

GetRegionNames returns the RegionNames field if non-nil, zero value otherwise.

### GetRegionNamesOk

`func (o *TrailCreateRequest) GetRegionNamesOk() (*[]interface{}, bool)`

GetRegionNamesOk returns a tuple with the RegionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNames

`func (o *TrailCreateRequest) SetRegionNames(v []interface{})`

SetRegionNames sets RegionNames field to given value.

### HasRegionNames

`func (o *TrailCreateRequest) HasRegionNames() bool`

HasRegionNames returns a boolean if a field has been set.

### GetRegionTotalYn

`func (o *TrailCreateRequest) GetRegionTotalYn() string`

GetRegionTotalYn returns the RegionTotalYn field if non-nil, zero value otherwise.

### GetRegionTotalYnOk

`func (o *TrailCreateRequest) GetRegionTotalYnOk() (*string, bool)`

GetRegionTotalYnOk returns a tuple with the RegionTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionTotalYn

`func (o *TrailCreateRequest) SetRegionTotalYn(v string)`

SetRegionTotalYn sets RegionTotalYn field to given value.

### HasRegionTotalYn

`func (o *TrailCreateRequest) HasRegionTotalYn() bool`

HasRegionTotalYn returns a boolean if a field has been set.

### SetRegionTotalYnNil

`func (o *TrailCreateRequest) SetRegionTotalYnNil(b bool)`

 SetRegionTotalYnNil sets the value for RegionTotalYn to be an explicit nil

### UnsetRegionTotalYn
`func (o *TrailCreateRequest) UnsetRegionTotalYn()`

UnsetRegionTotalYn ensures that no value is present for RegionTotalYn, not even an explicit nil
### GetResourceTypeTotalYn

`func (o *TrailCreateRequest) GetResourceTypeTotalYn() string`

GetResourceTypeTotalYn returns the ResourceTypeTotalYn field if non-nil, zero value otherwise.

### GetResourceTypeTotalYnOk

`func (o *TrailCreateRequest) GetResourceTypeTotalYnOk() (*string, bool)`

GetResourceTypeTotalYnOk returns a tuple with the ResourceTypeTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeTotalYn

`func (o *TrailCreateRequest) SetResourceTypeTotalYn(v string)`

SetResourceTypeTotalYn sets ResourceTypeTotalYn field to given value.

### HasResourceTypeTotalYn

`func (o *TrailCreateRequest) HasResourceTypeTotalYn() bool`

HasResourceTypeTotalYn returns a boolean if a field has been set.

### SetResourceTypeTotalYnNil

`func (o *TrailCreateRequest) SetResourceTypeTotalYnNil(b bool)`

 SetResourceTypeTotalYnNil sets the value for ResourceTypeTotalYn to be an explicit nil

### UnsetResourceTypeTotalYn
`func (o *TrailCreateRequest) UnsetResourceTypeTotalYn()`

UnsetResourceTypeTotalYn ensures that no value is present for ResourceTypeTotalYn, not even an explicit nil
### GetTagCreateRequests

`func (o *TrailCreateRequest) GetTagCreateRequests() []map[string]string`

GetTagCreateRequests returns the TagCreateRequests field if non-nil, zero value otherwise.

### GetTagCreateRequestsOk

`func (o *TrailCreateRequest) GetTagCreateRequestsOk() (*[]map[string]string, bool)`

GetTagCreateRequestsOk returns a tuple with the TagCreateRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCreateRequests

`func (o *TrailCreateRequest) SetTagCreateRequests(v []map[string]string)`

SetTagCreateRequests sets TagCreateRequests field to given value.

### HasTagCreateRequests

`func (o *TrailCreateRequest) HasTagCreateRequests() bool`

HasTagCreateRequests returns a boolean if a field has been set.

### SetTagCreateRequestsNil

`func (o *TrailCreateRequest) SetTagCreateRequestsNil(b bool)`

 SetTagCreateRequestsNil sets the value for TagCreateRequests to be an explicit nil

### UnsetTagCreateRequests
`func (o *TrailCreateRequest) UnsetTagCreateRequests()`

UnsetTagCreateRequests ensures that no value is present for TagCreateRequests, not even an explicit nil
### GetTargetLogTypes

`func (o *TrailCreateRequest) GetTargetLogTypes() []interface{}`

GetTargetLogTypes returns the TargetLogTypes field if non-nil, zero value otherwise.

### GetTargetLogTypesOk

`func (o *TrailCreateRequest) GetTargetLogTypesOk() (*[]interface{}, bool)`

GetTargetLogTypesOk returns a tuple with the TargetLogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLogTypes

`func (o *TrailCreateRequest) SetTargetLogTypes(v []interface{})`

SetTargetLogTypes sets TargetLogTypes field to given value.

### HasTargetLogTypes

`func (o *TrailCreateRequest) HasTargetLogTypes() bool`

HasTargetLogTypes returns a boolean if a field has been set.

### GetTargetResourceTypes

`func (o *TrailCreateRequest) GetTargetResourceTypes() []interface{}`

GetTargetResourceTypes returns the TargetResourceTypes field if non-nil, zero value otherwise.

### GetTargetResourceTypesOk

`func (o *TrailCreateRequest) GetTargetResourceTypesOk() (*[]interface{}, bool)`

GetTargetResourceTypesOk returns a tuple with the TargetResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceTypes

`func (o *TrailCreateRequest) SetTargetResourceTypes(v []interface{})`

SetTargetResourceTypes sets TargetResourceTypes field to given value.

### HasTargetResourceTypes

`func (o *TrailCreateRequest) HasTargetResourceTypes() bool`

HasTargetResourceTypes returns a boolean if a field has been set.

### GetTargetUsers

`func (o *TrailCreateRequest) GetTargetUsers() []interface{}`

GetTargetUsers returns the TargetUsers field if non-nil, zero value otherwise.

### GetTargetUsersOk

`func (o *TrailCreateRequest) GetTargetUsersOk() (*[]interface{}, bool)`

GetTargetUsersOk returns a tuple with the TargetUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUsers

`func (o *TrailCreateRequest) SetTargetUsers(v []interface{})`

SetTargetUsers sets TargetUsers field to given value.

### HasTargetUsers

`func (o *TrailCreateRequest) HasTargetUsers() bool`

HasTargetUsers returns a boolean if a field has been set.

### GetTrailDescription

`func (o *TrailCreateRequest) GetTrailDescription() string`

GetTrailDescription returns the TrailDescription field if non-nil, zero value otherwise.

### GetTrailDescriptionOk

`func (o *TrailCreateRequest) GetTrailDescriptionOk() (*string, bool)`

GetTrailDescriptionOk returns a tuple with the TrailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailDescription

`func (o *TrailCreateRequest) SetTrailDescription(v string)`

SetTrailDescription sets TrailDescription field to given value.

### HasTrailDescription

`func (o *TrailCreateRequest) HasTrailDescription() bool`

HasTrailDescription returns a boolean if a field has been set.

### SetTrailDescriptionNil

`func (o *TrailCreateRequest) SetTrailDescriptionNil(b bool)`

 SetTrailDescriptionNil sets the value for TrailDescription to be an explicit nil

### UnsetTrailDescription
`func (o *TrailCreateRequest) UnsetTrailDescription()`

UnsetTrailDescription ensures that no value is present for TrailDescription, not even an explicit nil
### GetTrailName

`func (o *TrailCreateRequest) GetTrailName() string`

GetTrailName returns the TrailName field if non-nil, zero value otherwise.

### GetTrailNameOk

`func (o *TrailCreateRequest) GetTrailNameOk() (*string, bool)`

GetTrailNameOk returns a tuple with the TrailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailName

`func (o *TrailCreateRequest) SetTrailName(v string)`

SetTrailName sets TrailName field to given value.


### GetTrailSaveType

`func (o *TrailCreateRequest) GetTrailSaveType() string`

GetTrailSaveType returns the TrailSaveType field if non-nil, zero value otherwise.

### GetTrailSaveTypeOk

`func (o *TrailCreateRequest) GetTrailSaveTypeOk() (*string, bool)`

GetTrailSaveTypeOk returns a tuple with the TrailSaveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailSaveType

`func (o *TrailCreateRequest) SetTrailSaveType(v string)`

SetTrailSaveType sets TrailSaveType field to given value.


### GetUserTotalYn

`func (o *TrailCreateRequest) GetUserTotalYn() string`

GetUserTotalYn returns the UserTotalYn field if non-nil, zero value otherwise.

### GetUserTotalYnOk

`func (o *TrailCreateRequest) GetUserTotalYnOk() (*string, bool)`

GetUserTotalYnOk returns a tuple with the UserTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTotalYn

`func (o *TrailCreateRequest) SetUserTotalYn(v string)`

SetUserTotalYn sets UserTotalYn field to given value.

### HasUserTotalYn

`func (o *TrailCreateRequest) HasUserTotalYn() bool`

HasUserTotalYn returns a boolean if a field has been set.

### SetUserTotalYnNil

`func (o *TrailCreateRequest) SetUserTotalYnNil(b bool)`

 SetUserTotalYnNil sets the value for UserTotalYn to be an explicit nil

### UnsetUserTotalYn
`func (o *TrailCreateRequest) UnsetUserTotalYn()`

UnsetUserTotalYn ensures that no value is present for UserTotalYn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


