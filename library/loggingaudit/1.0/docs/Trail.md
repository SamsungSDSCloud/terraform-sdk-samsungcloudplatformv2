# Trail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AccountName** | Pointer to **NullableString** |  | [optional] 
**BucketName** | **string** | Bucket name | 
**BucketRegion** | **string** | Bucket region | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**CreatedUserId** | Pointer to **NullableString** |  | [optional] 
**DelYn** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | ID | 
**LogTypeTotalYn** | Pointer to **NullableString** |  | [optional] 
**LogVerificationYn** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**RegionNames** | Pointer to **[]interface{}** | Region name list | [optional] [default to []]
**RegionTotalYn** | Pointer to **NullableString** |  | [optional] 
**ResourceTypeTotalYn** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 
**TargetLogTypes** | Pointer to **[]interface{}** | Target log type list | [optional] [default to []]
**TargetResourceTypes** | Pointer to **[]interface{}** | Target resource type list | [optional] [default to []]
**TargetUsers** | Pointer to **[]interface{}** | Target user list | [optional] [default to []]
**TrailBatchEndAt** | Pointer to **NullableTime** |  | [optional] 
**TrailBatchFirstStartAt** | Pointer to **NullableTime** |  | [optional] 
**TrailBatchLastState** | Pointer to **NullableString** |  | [optional] 
**TrailBatchStartAt** | Pointer to **NullableTime** |  | [optional] 
**TrailBatchSuccessAt** | Pointer to **NullableTime** |  | [optional] 
**TrailDescription** | Pointer to **NullableString** |  | [optional] 
**TrailName** | **string** | Trail name | 
**TrailSaveType** | **string** | Trail save type | 
**UserTotalYn** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrail

`func NewTrail(bucketName string, bucketRegion string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, trailName string, trailSaveType string, ) *Trail`

NewTrail instantiates a new Trail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailWithDefaults

`func NewTrailWithDefaults() *Trail`

NewTrailWithDefaults instantiates a new Trail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *Trail) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Trail) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Trail) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Trail) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *Trail) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Trail) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAccountName

`func (o *Trail) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Trail) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Trail) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *Trail) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *Trail) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *Trail) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetBucketName

`func (o *Trail) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *Trail) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *Trail) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetBucketRegion

`func (o *Trail) GetBucketRegion() string`

GetBucketRegion returns the BucketRegion field if non-nil, zero value otherwise.

### GetBucketRegionOk

`func (o *Trail) GetBucketRegionOk() (*string, bool)`

GetBucketRegionOk returns a tuple with the BucketRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketRegion

`func (o *Trail) SetBucketRegion(v string)`

SetBucketRegion sets BucketRegion field to given value.


### GetCreatedAt

`func (o *Trail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Trail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Trail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Trail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Trail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Trail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedUserId

`func (o *Trail) GetCreatedUserId() string`

GetCreatedUserId returns the CreatedUserId field if non-nil, zero value otherwise.

### GetCreatedUserIdOk

`func (o *Trail) GetCreatedUserIdOk() (*string, bool)`

GetCreatedUserIdOk returns a tuple with the CreatedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUserId

`func (o *Trail) SetCreatedUserId(v string)`

SetCreatedUserId sets CreatedUserId field to given value.

### HasCreatedUserId

`func (o *Trail) HasCreatedUserId() bool`

HasCreatedUserId returns a boolean if a field has been set.

### SetCreatedUserIdNil

`func (o *Trail) SetCreatedUserIdNil(b bool)`

 SetCreatedUserIdNil sets the value for CreatedUserId to be an explicit nil

### UnsetCreatedUserId
`func (o *Trail) UnsetCreatedUserId()`

UnsetCreatedUserId ensures that no value is present for CreatedUserId, not even an explicit nil
### GetDelYn

`func (o *Trail) GetDelYn() string`

GetDelYn returns the DelYn field if non-nil, zero value otherwise.

### GetDelYnOk

`func (o *Trail) GetDelYnOk() (*string, bool)`

GetDelYnOk returns a tuple with the DelYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelYn

`func (o *Trail) SetDelYn(v string)`

SetDelYn sets DelYn field to given value.

### HasDelYn

`func (o *Trail) HasDelYn() bool`

HasDelYn returns a boolean if a field has been set.

### SetDelYnNil

`func (o *Trail) SetDelYnNil(b bool)`

 SetDelYnNil sets the value for DelYn to be an explicit nil

### UnsetDelYn
`func (o *Trail) UnsetDelYn()`

UnsetDelYn ensures that no value is present for DelYn, not even an explicit nil
### GetId

`func (o *Trail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trail) SetId(v string)`

SetId sets Id field to given value.


### GetLogTypeTotalYn

`func (o *Trail) GetLogTypeTotalYn() string`

GetLogTypeTotalYn returns the LogTypeTotalYn field if non-nil, zero value otherwise.

### GetLogTypeTotalYnOk

`func (o *Trail) GetLogTypeTotalYnOk() (*string, bool)`

GetLogTypeTotalYnOk returns a tuple with the LogTypeTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypeTotalYn

`func (o *Trail) SetLogTypeTotalYn(v string)`

SetLogTypeTotalYn sets LogTypeTotalYn field to given value.

### HasLogTypeTotalYn

`func (o *Trail) HasLogTypeTotalYn() bool`

HasLogTypeTotalYn returns a boolean if a field has been set.

### SetLogTypeTotalYnNil

`func (o *Trail) SetLogTypeTotalYnNil(b bool)`

 SetLogTypeTotalYnNil sets the value for LogTypeTotalYn to be an explicit nil

### UnsetLogTypeTotalYn
`func (o *Trail) UnsetLogTypeTotalYn()`

UnsetLogTypeTotalYn ensures that no value is present for LogTypeTotalYn, not even an explicit nil
### GetLogVerificationYn

`func (o *Trail) GetLogVerificationYn() string`

GetLogVerificationYn returns the LogVerificationYn field if non-nil, zero value otherwise.

### GetLogVerificationYnOk

`func (o *Trail) GetLogVerificationYnOk() (*string, bool)`

GetLogVerificationYnOk returns a tuple with the LogVerificationYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogVerificationYn

`func (o *Trail) SetLogVerificationYn(v string)`

SetLogVerificationYn sets LogVerificationYn field to given value.

### HasLogVerificationYn

`func (o *Trail) HasLogVerificationYn() bool`

HasLogVerificationYn returns a boolean if a field has been set.

### SetLogVerificationYnNil

`func (o *Trail) SetLogVerificationYnNil(b bool)`

 SetLogVerificationYnNil sets the value for LogVerificationYn to be an explicit nil

### UnsetLogVerificationYn
`func (o *Trail) UnsetLogVerificationYn()`

UnsetLogVerificationYn ensures that no value is present for LogVerificationYn, not even an explicit nil
### GetModifiedAt

`func (o *Trail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Trail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Trail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Trail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Trail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Trail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetRegionNames

`func (o *Trail) GetRegionNames() []interface{}`

GetRegionNames returns the RegionNames field if non-nil, zero value otherwise.

### GetRegionNamesOk

`func (o *Trail) GetRegionNamesOk() (*[]interface{}, bool)`

GetRegionNamesOk returns a tuple with the RegionNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNames

`func (o *Trail) SetRegionNames(v []interface{})`

SetRegionNames sets RegionNames field to given value.

### HasRegionNames

`func (o *Trail) HasRegionNames() bool`

HasRegionNames returns a boolean if a field has been set.

### GetRegionTotalYn

`func (o *Trail) GetRegionTotalYn() string`

GetRegionTotalYn returns the RegionTotalYn field if non-nil, zero value otherwise.

### GetRegionTotalYnOk

`func (o *Trail) GetRegionTotalYnOk() (*string, bool)`

GetRegionTotalYnOk returns a tuple with the RegionTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionTotalYn

`func (o *Trail) SetRegionTotalYn(v string)`

SetRegionTotalYn sets RegionTotalYn field to given value.

### HasRegionTotalYn

`func (o *Trail) HasRegionTotalYn() bool`

HasRegionTotalYn returns a boolean if a field has been set.

### SetRegionTotalYnNil

`func (o *Trail) SetRegionTotalYnNil(b bool)`

 SetRegionTotalYnNil sets the value for RegionTotalYn to be an explicit nil

### UnsetRegionTotalYn
`func (o *Trail) UnsetRegionTotalYn()`

UnsetRegionTotalYn ensures that no value is present for RegionTotalYn, not even an explicit nil
### GetResourceTypeTotalYn

`func (o *Trail) GetResourceTypeTotalYn() string`

GetResourceTypeTotalYn returns the ResourceTypeTotalYn field if non-nil, zero value otherwise.

### GetResourceTypeTotalYnOk

`func (o *Trail) GetResourceTypeTotalYnOk() (*string, bool)`

GetResourceTypeTotalYnOk returns a tuple with the ResourceTypeTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTypeTotalYn

`func (o *Trail) SetResourceTypeTotalYn(v string)`

SetResourceTypeTotalYn sets ResourceTypeTotalYn field to given value.

### HasResourceTypeTotalYn

`func (o *Trail) HasResourceTypeTotalYn() bool`

HasResourceTypeTotalYn returns a boolean if a field has been set.

### SetResourceTypeTotalYnNil

`func (o *Trail) SetResourceTypeTotalYnNil(b bool)`

 SetResourceTypeTotalYnNil sets the value for ResourceTypeTotalYn to be an explicit nil

### UnsetResourceTypeTotalYn
`func (o *Trail) UnsetResourceTypeTotalYn()`

UnsetResourceTypeTotalYn ensures that no value is present for ResourceTypeTotalYn, not even an explicit nil
### GetState

`func (o *Trail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Trail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Trail) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Trail) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *Trail) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *Trail) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetTargetLogTypes

`func (o *Trail) GetTargetLogTypes() []interface{}`

GetTargetLogTypes returns the TargetLogTypes field if non-nil, zero value otherwise.

### GetTargetLogTypesOk

`func (o *Trail) GetTargetLogTypesOk() (*[]interface{}, bool)`

GetTargetLogTypesOk returns a tuple with the TargetLogTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLogTypes

`func (o *Trail) SetTargetLogTypes(v []interface{})`

SetTargetLogTypes sets TargetLogTypes field to given value.

### HasTargetLogTypes

`func (o *Trail) HasTargetLogTypes() bool`

HasTargetLogTypes returns a boolean if a field has been set.

### GetTargetResourceTypes

`func (o *Trail) GetTargetResourceTypes() []interface{}`

GetTargetResourceTypes returns the TargetResourceTypes field if non-nil, zero value otherwise.

### GetTargetResourceTypesOk

`func (o *Trail) GetTargetResourceTypesOk() (*[]interface{}, bool)`

GetTargetResourceTypesOk returns a tuple with the TargetResourceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceTypes

`func (o *Trail) SetTargetResourceTypes(v []interface{})`

SetTargetResourceTypes sets TargetResourceTypes field to given value.

### HasTargetResourceTypes

`func (o *Trail) HasTargetResourceTypes() bool`

HasTargetResourceTypes returns a boolean if a field has been set.

### GetTargetUsers

`func (o *Trail) GetTargetUsers() []interface{}`

GetTargetUsers returns the TargetUsers field if non-nil, zero value otherwise.

### GetTargetUsersOk

`func (o *Trail) GetTargetUsersOk() (*[]interface{}, bool)`

GetTargetUsersOk returns a tuple with the TargetUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUsers

`func (o *Trail) SetTargetUsers(v []interface{})`

SetTargetUsers sets TargetUsers field to given value.

### HasTargetUsers

`func (o *Trail) HasTargetUsers() bool`

HasTargetUsers returns a boolean if a field has been set.

### GetTrailBatchEndAt

`func (o *Trail) GetTrailBatchEndAt() time.Time`

GetTrailBatchEndAt returns the TrailBatchEndAt field if non-nil, zero value otherwise.

### GetTrailBatchEndAtOk

`func (o *Trail) GetTrailBatchEndAtOk() (*time.Time, bool)`

GetTrailBatchEndAtOk returns a tuple with the TrailBatchEndAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailBatchEndAt

`func (o *Trail) SetTrailBatchEndAt(v time.Time)`

SetTrailBatchEndAt sets TrailBatchEndAt field to given value.

### HasTrailBatchEndAt

`func (o *Trail) HasTrailBatchEndAt() bool`

HasTrailBatchEndAt returns a boolean if a field has been set.

### SetTrailBatchEndAtNil

`func (o *Trail) SetTrailBatchEndAtNil(b bool)`

 SetTrailBatchEndAtNil sets the value for TrailBatchEndAt to be an explicit nil

### UnsetTrailBatchEndAt
`func (o *Trail) UnsetTrailBatchEndAt()`

UnsetTrailBatchEndAt ensures that no value is present for TrailBatchEndAt, not even an explicit nil
### GetTrailBatchFirstStartAt

`func (o *Trail) GetTrailBatchFirstStartAt() time.Time`

GetTrailBatchFirstStartAt returns the TrailBatchFirstStartAt field if non-nil, zero value otherwise.

### GetTrailBatchFirstStartAtOk

`func (o *Trail) GetTrailBatchFirstStartAtOk() (*time.Time, bool)`

GetTrailBatchFirstStartAtOk returns a tuple with the TrailBatchFirstStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailBatchFirstStartAt

`func (o *Trail) SetTrailBatchFirstStartAt(v time.Time)`

SetTrailBatchFirstStartAt sets TrailBatchFirstStartAt field to given value.

### HasTrailBatchFirstStartAt

`func (o *Trail) HasTrailBatchFirstStartAt() bool`

HasTrailBatchFirstStartAt returns a boolean if a field has been set.

### SetTrailBatchFirstStartAtNil

`func (o *Trail) SetTrailBatchFirstStartAtNil(b bool)`

 SetTrailBatchFirstStartAtNil sets the value for TrailBatchFirstStartAt to be an explicit nil

### UnsetTrailBatchFirstStartAt
`func (o *Trail) UnsetTrailBatchFirstStartAt()`

UnsetTrailBatchFirstStartAt ensures that no value is present for TrailBatchFirstStartAt, not even an explicit nil
### GetTrailBatchLastState

`func (o *Trail) GetTrailBatchLastState() string`

GetTrailBatchLastState returns the TrailBatchLastState field if non-nil, zero value otherwise.

### GetTrailBatchLastStateOk

`func (o *Trail) GetTrailBatchLastStateOk() (*string, bool)`

GetTrailBatchLastStateOk returns a tuple with the TrailBatchLastState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailBatchLastState

`func (o *Trail) SetTrailBatchLastState(v string)`

SetTrailBatchLastState sets TrailBatchLastState field to given value.

### HasTrailBatchLastState

`func (o *Trail) HasTrailBatchLastState() bool`

HasTrailBatchLastState returns a boolean if a field has been set.

### SetTrailBatchLastStateNil

`func (o *Trail) SetTrailBatchLastStateNil(b bool)`

 SetTrailBatchLastStateNil sets the value for TrailBatchLastState to be an explicit nil

### UnsetTrailBatchLastState
`func (o *Trail) UnsetTrailBatchLastState()`

UnsetTrailBatchLastState ensures that no value is present for TrailBatchLastState, not even an explicit nil
### GetTrailBatchStartAt

`func (o *Trail) GetTrailBatchStartAt() time.Time`

GetTrailBatchStartAt returns the TrailBatchStartAt field if non-nil, zero value otherwise.

### GetTrailBatchStartAtOk

`func (o *Trail) GetTrailBatchStartAtOk() (*time.Time, bool)`

GetTrailBatchStartAtOk returns a tuple with the TrailBatchStartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailBatchStartAt

`func (o *Trail) SetTrailBatchStartAt(v time.Time)`

SetTrailBatchStartAt sets TrailBatchStartAt field to given value.

### HasTrailBatchStartAt

`func (o *Trail) HasTrailBatchStartAt() bool`

HasTrailBatchStartAt returns a boolean if a field has been set.

### SetTrailBatchStartAtNil

`func (o *Trail) SetTrailBatchStartAtNil(b bool)`

 SetTrailBatchStartAtNil sets the value for TrailBatchStartAt to be an explicit nil

### UnsetTrailBatchStartAt
`func (o *Trail) UnsetTrailBatchStartAt()`

UnsetTrailBatchStartAt ensures that no value is present for TrailBatchStartAt, not even an explicit nil
### GetTrailBatchSuccessAt

`func (o *Trail) GetTrailBatchSuccessAt() time.Time`

GetTrailBatchSuccessAt returns the TrailBatchSuccessAt field if non-nil, zero value otherwise.

### GetTrailBatchSuccessAtOk

`func (o *Trail) GetTrailBatchSuccessAtOk() (*time.Time, bool)`

GetTrailBatchSuccessAtOk returns a tuple with the TrailBatchSuccessAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailBatchSuccessAt

`func (o *Trail) SetTrailBatchSuccessAt(v time.Time)`

SetTrailBatchSuccessAt sets TrailBatchSuccessAt field to given value.

### HasTrailBatchSuccessAt

`func (o *Trail) HasTrailBatchSuccessAt() bool`

HasTrailBatchSuccessAt returns a boolean if a field has been set.

### SetTrailBatchSuccessAtNil

`func (o *Trail) SetTrailBatchSuccessAtNil(b bool)`

 SetTrailBatchSuccessAtNil sets the value for TrailBatchSuccessAt to be an explicit nil

### UnsetTrailBatchSuccessAt
`func (o *Trail) UnsetTrailBatchSuccessAt()`

UnsetTrailBatchSuccessAt ensures that no value is present for TrailBatchSuccessAt, not even an explicit nil
### GetTrailDescription

`func (o *Trail) GetTrailDescription() string`

GetTrailDescription returns the TrailDescription field if non-nil, zero value otherwise.

### GetTrailDescriptionOk

`func (o *Trail) GetTrailDescriptionOk() (*string, bool)`

GetTrailDescriptionOk returns a tuple with the TrailDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailDescription

`func (o *Trail) SetTrailDescription(v string)`

SetTrailDescription sets TrailDescription field to given value.

### HasTrailDescription

`func (o *Trail) HasTrailDescription() bool`

HasTrailDescription returns a boolean if a field has been set.

### SetTrailDescriptionNil

`func (o *Trail) SetTrailDescriptionNil(b bool)`

 SetTrailDescriptionNil sets the value for TrailDescription to be an explicit nil

### UnsetTrailDescription
`func (o *Trail) UnsetTrailDescription()`

UnsetTrailDescription ensures that no value is present for TrailDescription, not even an explicit nil
### GetTrailName

`func (o *Trail) GetTrailName() string`

GetTrailName returns the TrailName field if non-nil, zero value otherwise.

### GetTrailNameOk

`func (o *Trail) GetTrailNameOk() (*string, bool)`

GetTrailNameOk returns a tuple with the TrailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailName

`func (o *Trail) SetTrailName(v string)`

SetTrailName sets TrailName field to given value.


### GetTrailSaveType

`func (o *Trail) GetTrailSaveType() string`

GetTrailSaveType returns the TrailSaveType field if non-nil, zero value otherwise.

### GetTrailSaveTypeOk

`func (o *Trail) GetTrailSaveTypeOk() (*string, bool)`

GetTrailSaveTypeOk returns a tuple with the TrailSaveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailSaveType

`func (o *Trail) SetTrailSaveType(v string)`

SetTrailSaveType sets TrailSaveType field to given value.


### GetUserTotalYn

`func (o *Trail) GetUserTotalYn() string`

GetUserTotalYn returns the UserTotalYn field if non-nil, zero value otherwise.

### GetUserTotalYnOk

`func (o *Trail) GetUserTotalYnOk() (*string, bool)`

GetUserTotalYnOk returns a tuple with the UserTotalYn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTotalYn

`func (o *Trail) SetUserTotalYn(v string)`

SetUserTotalYn sets UserTotalYn field to given value.

### HasUserTotalYn

`func (o *Trail) HasUserTotalYn() bool`

HasUserTotalYn returns a boolean if a field has been set.

### SetUserTotalYnNil

`func (o *Trail) SetUserTotalYnNil(b bool)`

 SetUserTotalYnNil sets the value for UserTotalYn to be an explicit nil

### UnsetUserTotalYn
`func (o *Trail) UnsetUserTotalYn()`

UnsetUserTotalYn ensures that no value is present for UserTotalYn, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


