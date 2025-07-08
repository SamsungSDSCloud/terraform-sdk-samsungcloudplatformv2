# MemberCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberIp** | **string** | IP of Member | 
**MemberPort** | **int32** | Protocol port of Member | 
**MemberWeight** | Pointer to **int32** | Weight of Member | [optional] 
**Name** | **string** | Name of Member | 
**ObjectId** | Pointer to **string** | Object ID of Member | [optional] 
**ObjectType** | **string** | Object Type of Member | 

## Methods

### NewMemberCreateRequest

`func NewMemberCreateRequest(memberIp string, memberPort int32, name string, objectType string, ) *MemberCreateRequest`

NewMemberCreateRequest instantiates a new MemberCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberCreateRequestWithDefaults

`func NewMemberCreateRequestWithDefaults() *MemberCreateRequest`

NewMemberCreateRequestWithDefaults instantiates a new MemberCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberIp

`func (o *MemberCreateRequest) GetMemberIp() string`

GetMemberIp returns the MemberIp field if non-nil, zero value otherwise.

### GetMemberIpOk

`func (o *MemberCreateRequest) GetMemberIpOk() (*string, bool)`

GetMemberIpOk returns a tuple with the MemberIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIp

`func (o *MemberCreateRequest) SetMemberIp(v string)`

SetMemberIp sets MemberIp field to given value.


### GetMemberPort

`func (o *MemberCreateRequest) GetMemberPort() int32`

GetMemberPort returns the MemberPort field if non-nil, zero value otherwise.

### GetMemberPortOk

`func (o *MemberCreateRequest) GetMemberPortOk() (*int32, bool)`

GetMemberPortOk returns a tuple with the MemberPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPort

`func (o *MemberCreateRequest) SetMemberPort(v int32)`

SetMemberPort sets MemberPort field to given value.


### GetMemberWeight

`func (o *MemberCreateRequest) GetMemberWeight() int32`

GetMemberWeight returns the MemberWeight field if non-nil, zero value otherwise.

### GetMemberWeightOk

`func (o *MemberCreateRequest) GetMemberWeightOk() (*int32, bool)`

GetMemberWeightOk returns a tuple with the MemberWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberWeight

`func (o *MemberCreateRequest) SetMemberWeight(v int32)`

SetMemberWeight sets MemberWeight field to given value.

### HasMemberWeight

`func (o *MemberCreateRequest) HasMemberWeight() bool`

HasMemberWeight returns a boolean if a field has been set.

### GetName

`func (o *MemberCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetObjectId

`func (o *MemberCreateRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *MemberCreateRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *MemberCreateRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *MemberCreateRequest) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectType

`func (o *MemberCreateRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *MemberCreateRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *MemberCreateRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


