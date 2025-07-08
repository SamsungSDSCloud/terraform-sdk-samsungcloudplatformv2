# SecurityGroupRulesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupId** | **string** | Security Group Rule ID | 
**SecurityGroupRules** | [**[]SecurityGroupRuleCreateRequestParam**](SecurityGroupRuleCreateRequestParam.md) |  | 

## Methods

### NewSecurityGroupRulesCreateRequest

`func NewSecurityGroupRulesCreateRequest(securityGroupId string, securityGroupRules []SecurityGroupRuleCreateRequestParam, ) *SecurityGroupRulesCreateRequest`

NewSecurityGroupRulesCreateRequest instantiates a new SecurityGroupRulesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRulesCreateRequestWithDefaults

`func NewSecurityGroupRulesCreateRequestWithDefaults() *SecurityGroupRulesCreateRequest`

NewSecurityGroupRulesCreateRequestWithDefaults instantiates a new SecurityGroupRulesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupId

`func (o *SecurityGroupRulesCreateRequest) GetSecurityGroupId() string`

GetSecurityGroupId returns the SecurityGroupId field if non-nil, zero value otherwise.

### GetSecurityGroupIdOk

`func (o *SecurityGroupRulesCreateRequest) GetSecurityGroupIdOk() (*string, bool)`

GetSecurityGroupIdOk returns a tuple with the SecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupId

`func (o *SecurityGroupRulesCreateRequest) SetSecurityGroupId(v string)`

SetSecurityGroupId sets SecurityGroupId field to given value.


### GetSecurityGroupRules

`func (o *SecurityGroupRulesCreateRequest) GetSecurityGroupRules() []SecurityGroupRuleCreateRequestParam`

GetSecurityGroupRules returns the SecurityGroupRules field if non-nil, zero value otherwise.

### GetSecurityGroupRulesOk

`func (o *SecurityGroupRulesCreateRequest) GetSecurityGroupRulesOk() (*[]SecurityGroupRuleCreateRequestParam, bool)`

GetSecurityGroupRulesOk returns a tuple with the SecurityGroupRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRules

`func (o *SecurityGroupRulesCreateRequest) SetSecurityGroupRules(v []SecurityGroupRuleCreateRequestParam)`

SetSecurityGroupRules sets SecurityGroupRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


