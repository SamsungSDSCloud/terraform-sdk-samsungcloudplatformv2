# SecurityGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Server Group owner&#39;s account ID | 
**Description** | **string** | Security group description | 
**Id** | **string** | Security Group ID | 
**Name** | **string** | Security Group name | 
**Rules** | [**[]SecurityGroupResponseRule**](SecurityGroupResponseRule.md) | Security group rules | 

## Methods

### NewSecurityGroupResponse

`func NewSecurityGroupResponse(accountId string, description string, id string, name string, rules []SecurityGroupResponseRule, ) *SecurityGroupResponse`

NewSecurityGroupResponse instantiates a new SecurityGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupResponseWithDefaults

`func NewSecurityGroupResponseWithDefaults() *SecurityGroupResponse`

NewSecurityGroupResponseWithDefaults instantiates a new SecurityGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SecurityGroupResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SecurityGroupResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SecurityGroupResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDescription

`func (o *SecurityGroupResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityGroupResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityGroupResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetId

`func (o *SecurityGroupResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityGroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRules

`func (o *SecurityGroupResponse) GetRules() []SecurityGroupResponseRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *SecurityGroupResponse) GetRulesOk() (*[]SecurityGroupResponseRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *SecurityGroupResponse) SetRules(v []SecurityGroupResponseRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


