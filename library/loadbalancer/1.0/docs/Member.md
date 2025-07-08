# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | ID | 
**LbServerGroupId** | **string** | ID of Lb Server Group | 
**MemberIp** | **string** | IP of Member | 
**MemberPort** | **int32** | Protocol port of Member | 
**MemberState** | **string** | Member State of Member | 
**MemberWeight** | **int32** | Weight of Member | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Name of Member | 
**ObjectId** | Pointer to **NullableString** |  | [optional] 
**ObjectType** | [**MemberType**](MemberType.md) | Object Type of Member | 
**State** | [**MemberState**](MemberState.md) | State of Member | 
**SubnetId** | **string** | ID of Subnet | 
**Uuid** | **string** | UUID of Member | 

## Methods

### NewMember

`func NewMember(createdAt time.Time, createdBy string, id string, lbServerGroupId string, memberIp string, memberPort int32, memberState string, memberWeight int32, modifiedAt time.Time, modifiedBy string, name string, objectType MemberType, state MemberState, subnetId string, uuid string, ) *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Member) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Member) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Member) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Member) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Member) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Member) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *Member) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Member) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Member) SetId(v string)`

SetId sets Id field to given value.


### GetLbServerGroupId

`func (o *Member) GetLbServerGroupId() string`

GetLbServerGroupId returns the LbServerGroupId field if non-nil, zero value otherwise.

### GetLbServerGroupIdOk

`func (o *Member) GetLbServerGroupIdOk() (*string, bool)`

GetLbServerGroupIdOk returns a tuple with the LbServerGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroupId

`func (o *Member) SetLbServerGroupId(v string)`

SetLbServerGroupId sets LbServerGroupId field to given value.


### GetMemberIp

`func (o *Member) GetMemberIp() string`

GetMemberIp returns the MemberIp field if non-nil, zero value otherwise.

### GetMemberIpOk

`func (o *Member) GetMemberIpOk() (*string, bool)`

GetMemberIpOk returns a tuple with the MemberIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIp

`func (o *Member) SetMemberIp(v string)`

SetMemberIp sets MemberIp field to given value.


### GetMemberPort

`func (o *Member) GetMemberPort() int32`

GetMemberPort returns the MemberPort field if non-nil, zero value otherwise.

### GetMemberPortOk

`func (o *Member) GetMemberPortOk() (*int32, bool)`

GetMemberPortOk returns a tuple with the MemberPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPort

`func (o *Member) SetMemberPort(v int32)`

SetMemberPort sets MemberPort field to given value.


### GetMemberState

`func (o *Member) GetMemberState() string`

GetMemberState returns the MemberState field if non-nil, zero value otherwise.

### GetMemberStateOk

`func (o *Member) GetMemberStateOk() (*string, bool)`

GetMemberStateOk returns a tuple with the MemberState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberState

`func (o *Member) SetMemberState(v string)`

SetMemberState sets MemberState field to given value.


### GetMemberWeight

`func (o *Member) GetMemberWeight() int32`

GetMemberWeight returns the MemberWeight field if non-nil, zero value otherwise.

### GetMemberWeightOk

`func (o *Member) GetMemberWeightOk() (*int32, bool)`

GetMemberWeightOk returns a tuple with the MemberWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberWeight

`func (o *Member) SetMemberWeight(v int32)`

SetMemberWeight sets MemberWeight field to given value.


### GetModifiedAt

`func (o *Member) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Member) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Member) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Member) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Member) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Member) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *Member) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Member) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Member) SetName(v string)`

SetName sets Name field to given value.


### GetObjectId

`func (o *Member) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *Member) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *Member) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *Member) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### SetObjectIdNil

`func (o *Member) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *Member) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetObjectType

`func (o *Member) GetObjectType() MemberType`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *Member) GetObjectTypeOk() (*MemberType, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *Member) SetObjectType(v MemberType)`

SetObjectType sets ObjectType field to given value.


### GetState

`func (o *Member) GetState() MemberState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Member) GetStateOk() (*MemberState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Member) SetState(v MemberState)`

SetState sets State field to given value.


### GetSubnetId

`func (o *Member) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *Member) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *Member) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetUuid

`func (o *Member) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Member) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Member) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


