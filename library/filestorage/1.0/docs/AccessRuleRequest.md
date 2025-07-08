# AccessRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Access Rule Action | 
**ObjectId** | **string** | Object ID | 
**ObjectType** | **string** | Object Type | 

## Methods

### NewAccessRuleRequest

`func NewAccessRuleRequest(action string, objectId string, objectType string, ) *AccessRuleRequest`

NewAccessRuleRequest instantiates a new AccessRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleRequestWithDefaults

`func NewAccessRuleRequestWithDefaults() *AccessRuleRequest`

NewAccessRuleRequestWithDefaults instantiates a new AccessRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AccessRuleRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AccessRuleRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AccessRuleRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetObjectId

`func (o *AccessRuleRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AccessRuleRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AccessRuleRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectType

`func (o *AccessRuleRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AccessRuleRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AccessRuleRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


