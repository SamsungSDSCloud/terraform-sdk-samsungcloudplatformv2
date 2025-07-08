# LbServerGroupList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | LB Server Group ID | 
**LbHealthCheckId** | **NullableString** |  | 
**LbName** | Pointer to **NullableString** |  | [optional] 
**LbServerGroupMemberCount** | **int32** | LB Server Group Member count | 
**LoadbalancerId** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | LB Server Group name | 
**Protocol** | [**LbServerGroupProtocol**](LbServerGroupProtocol.md) | Protocol | 
**State** | **string** | LB Server Group state | 
**VpcId** | **string** | VPC ID | 

## Methods

### NewLbServerGroupList

`func NewLbServerGroupList(createdAt time.Time, createdBy string, id string, lbHealthCheckId NullableString, lbServerGroupMemberCount int32, modifiedAt time.Time, modifiedBy string, name string, protocol LbServerGroupProtocol, state string, vpcId string, ) *LbServerGroupList`

NewLbServerGroupList instantiates a new LbServerGroupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbServerGroupListWithDefaults

`func NewLbServerGroupListWithDefaults() *LbServerGroupList`

NewLbServerGroupListWithDefaults instantiates a new LbServerGroupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LbServerGroupList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LbServerGroupList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LbServerGroupList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LbServerGroupList) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LbServerGroupList) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LbServerGroupList) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *LbServerGroupList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LbServerGroupList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LbServerGroupList) SetId(v string)`

SetId sets Id field to given value.


### GetLbHealthCheckId

`func (o *LbServerGroupList) GetLbHealthCheckId() string`

GetLbHealthCheckId returns the LbHealthCheckId field if non-nil, zero value otherwise.

### GetLbHealthCheckIdOk

`func (o *LbServerGroupList) GetLbHealthCheckIdOk() (*string, bool)`

GetLbHealthCheckIdOk returns a tuple with the LbHealthCheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbHealthCheckId

`func (o *LbServerGroupList) SetLbHealthCheckId(v string)`

SetLbHealthCheckId sets LbHealthCheckId field to given value.


### SetLbHealthCheckIdNil

`func (o *LbServerGroupList) SetLbHealthCheckIdNil(b bool)`

 SetLbHealthCheckIdNil sets the value for LbHealthCheckId to be an explicit nil

### UnsetLbHealthCheckId
`func (o *LbServerGroupList) UnsetLbHealthCheckId()`

UnsetLbHealthCheckId ensures that no value is present for LbHealthCheckId, not even an explicit nil
### GetLbName

`func (o *LbServerGroupList) GetLbName() string`

GetLbName returns the LbName field if non-nil, zero value otherwise.

### GetLbNameOk

`func (o *LbServerGroupList) GetLbNameOk() (*string, bool)`

GetLbNameOk returns a tuple with the LbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbName

`func (o *LbServerGroupList) SetLbName(v string)`

SetLbName sets LbName field to given value.

### HasLbName

`func (o *LbServerGroupList) HasLbName() bool`

HasLbName returns a boolean if a field has been set.

### SetLbNameNil

`func (o *LbServerGroupList) SetLbNameNil(b bool)`

 SetLbNameNil sets the value for LbName to be an explicit nil

### UnsetLbName
`func (o *LbServerGroupList) UnsetLbName()`

UnsetLbName ensures that no value is present for LbName, not even an explicit nil
### GetLbServerGroupMemberCount

`func (o *LbServerGroupList) GetLbServerGroupMemberCount() int32`

GetLbServerGroupMemberCount returns the LbServerGroupMemberCount field if non-nil, zero value otherwise.

### GetLbServerGroupMemberCountOk

`func (o *LbServerGroupList) GetLbServerGroupMemberCountOk() (*int32, bool)`

GetLbServerGroupMemberCountOk returns a tuple with the LbServerGroupMemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroupMemberCount

`func (o *LbServerGroupList) SetLbServerGroupMemberCount(v int32)`

SetLbServerGroupMemberCount sets LbServerGroupMemberCount field to given value.


### GetLoadbalancerId

`func (o *LbServerGroupList) GetLoadbalancerId() string`

GetLoadbalancerId returns the LoadbalancerId field if non-nil, zero value otherwise.

### GetLoadbalancerIdOk

`func (o *LbServerGroupList) GetLoadbalancerIdOk() (*string, bool)`

GetLoadbalancerIdOk returns a tuple with the LoadbalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadbalancerId

`func (o *LbServerGroupList) SetLoadbalancerId(v string)`

SetLoadbalancerId sets LoadbalancerId field to given value.

### HasLoadbalancerId

`func (o *LbServerGroupList) HasLoadbalancerId() bool`

HasLoadbalancerId returns a boolean if a field has been set.

### SetLoadbalancerIdNil

`func (o *LbServerGroupList) SetLoadbalancerIdNil(b bool)`

 SetLoadbalancerIdNil sets the value for LoadbalancerId to be an explicit nil

### UnsetLoadbalancerId
`func (o *LbServerGroupList) UnsetLoadbalancerId()`

UnsetLoadbalancerId ensures that no value is present for LoadbalancerId, not even an explicit nil
### GetModifiedAt

`func (o *LbServerGroupList) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LbServerGroupList) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LbServerGroupList) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LbServerGroupList) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LbServerGroupList) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LbServerGroupList) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LbServerGroupList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbServerGroupList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbServerGroupList) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *LbServerGroupList) GetProtocol() LbServerGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LbServerGroupList) GetProtocolOk() (*LbServerGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LbServerGroupList) SetProtocol(v LbServerGroupProtocol)`

SetProtocol sets Protocol field to given value.


### GetState

`func (o *LbServerGroupList) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LbServerGroupList) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LbServerGroupList) SetState(v string)`

SetState sets State field to given value.


### GetVpcId

`func (o *LbServerGroupList) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LbServerGroupList) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LbServerGroupList) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


