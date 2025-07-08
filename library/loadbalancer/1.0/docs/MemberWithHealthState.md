# MemberWithHealthState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**HealthState** | **string** | Health State of Member | 
**Id** | **string** | ID | 
**LbServerGroupId** | **string** | ID of Lb Server Group | 
**MemberIp** | **string** | IP of Member | 
**MemberPort** | **int32** | Protocol port of Member | 
**MemberState** | **string** | Member State of Member | 
**MemberWeight** | **int32** | Weight of Member | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Name of Member | 
**ObjectId** | **NullableString** |  | 
**ObjectType** | [**MemberType**](MemberType.md) | Object Type of Member | 
**State** | [**MemberState**](MemberState.md) | State of Member | 

## Methods

### NewMemberWithHealthState

`func NewMemberWithHealthState(createdAt time.Time, createdBy string, healthState string, id string, lbServerGroupId string, memberIp string, memberPort int32, memberState string, memberWeight int32, modifiedAt time.Time, modifiedBy string, name string, objectId NullableString, objectType MemberType, state MemberState, ) *MemberWithHealthState`

NewMemberWithHealthState instantiates a new MemberWithHealthState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithHealthStateWithDefaults

`func NewMemberWithHealthStateWithDefaults() *MemberWithHealthState`

NewMemberWithHealthStateWithDefaults instantiates a new MemberWithHealthState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MemberWithHealthState) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MemberWithHealthState) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MemberWithHealthState) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *MemberWithHealthState) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MemberWithHealthState) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MemberWithHealthState) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetHealthState

`func (o *MemberWithHealthState) GetHealthState() string`

GetHealthState returns the HealthState field if non-nil, zero value otherwise.

### GetHealthStateOk

`func (o *MemberWithHealthState) GetHealthStateOk() (*string, bool)`

GetHealthStateOk returns a tuple with the HealthState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthState

`func (o *MemberWithHealthState) SetHealthState(v string)`

SetHealthState sets HealthState field to given value.


### GetId

`func (o *MemberWithHealthState) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberWithHealthState) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberWithHealthState) SetId(v string)`

SetId sets Id field to given value.


### GetLbServerGroupId

`func (o *MemberWithHealthState) GetLbServerGroupId() string`

GetLbServerGroupId returns the LbServerGroupId field if non-nil, zero value otherwise.

### GetLbServerGroupIdOk

`func (o *MemberWithHealthState) GetLbServerGroupIdOk() (*string, bool)`

GetLbServerGroupIdOk returns a tuple with the LbServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroupId

`func (o *MemberWithHealthState) SetLbServerGroupId(v string)`

SetLbServerGroupId sets LbServerGroupId field to given value.


### GetMemberIp

`func (o *MemberWithHealthState) GetMemberIp() string`

GetMemberIp returns the MemberIp field if non-nil, zero value otherwise.

### GetMemberIpOk

`func (o *MemberWithHealthState) GetMemberIpOk() (*string, bool)`

GetMemberIpOk returns a tuple with the MemberIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIp

`func (o *MemberWithHealthState) SetMemberIp(v string)`

SetMemberIp sets MemberIp field to given value.


### GetMemberPort

`func (o *MemberWithHealthState) GetMemberPort() int32`

GetMemberPort returns the MemberPort field if non-nil, zero value otherwise.

### GetMemberPortOk

`func (o *MemberWithHealthState) GetMemberPortOk() (*int32, bool)`

GetMemberPortOk returns a tuple with the MemberPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPort

`func (o *MemberWithHealthState) SetMemberPort(v int32)`

SetMemberPort sets MemberPort field to given value.


### GetMemberState

`func (o *MemberWithHealthState) GetMemberState() string`

GetMemberState returns the MemberState field if non-nil, zero value otherwise.

### GetMemberStateOk

`func (o *MemberWithHealthState) GetMemberStateOk() (*string, bool)`

GetMemberStateOk returns a tuple with the MemberState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberState

`func (o *MemberWithHealthState) SetMemberState(v string)`

SetMemberState sets MemberState field to given value.


### GetMemberWeight

`func (o *MemberWithHealthState) GetMemberWeight() int32`

GetMemberWeight returns the MemberWeight field if non-nil, zero value otherwise.

### GetMemberWeightOk

`func (o *MemberWithHealthState) GetMemberWeightOk() (*int32, bool)`

GetMemberWeightOk returns a tuple with the MemberWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberWeight

`func (o *MemberWithHealthState) SetMemberWeight(v int32)`

SetMemberWeight sets MemberWeight field to given value.


### GetModifiedAt

`func (o *MemberWithHealthState) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MemberWithHealthState) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MemberWithHealthState) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *MemberWithHealthState) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *MemberWithHealthState) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *MemberWithHealthState) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *MemberWithHealthState) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberWithHealthState) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberWithHealthState) SetName(v string)`

SetName sets Name field to given value.


### GetObjectId

`func (o *MemberWithHealthState) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MemberWithHealthState) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MemberWithHealthState) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### SetObjectIdNil

`func (o *MemberWithHealthState) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *MemberWithHealthState) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetObjectType

`func (o *MemberWithHealthState) GetObjectType() MemberType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemberWithHealthState) GetObjectTypeOk() (*MemberType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemberWithHealthState) SetObjectType(v MemberType)`

SetObjectType sets ObjectType field to given value.


### GetState

`func (o *MemberWithHealthState) GetState() MemberState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MemberWithHealthState) GetStateOk() (*MemberState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MemberWithHealthState) SetState(v MemberState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


