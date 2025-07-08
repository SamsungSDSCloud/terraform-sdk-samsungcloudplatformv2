# SecurityGroupRuleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**SecurityGroupRules** | [**[]SecurityGroupRule**](SecurityGroupRule.md) |  | 

## Methods

### NewSecurityGroupRuleListResponse

`func NewSecurityGroupRuleListResponse(securityGroupRules []SecurityGroupRule, ) *SecurityGroupRuleListResponse`

NewSecurityGroupRuleListResponse instantiates a new SecurityGroupRuleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupRuleListResponseWithDefaults

`func NewSecurityGroupRuleListResponseWithDefaults() *SecurityGroupRuleListResponse`

NewSecurityGroupRuleListResponseWithDefaults instantiates a new SecurityGroupRuleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SecurityGroupRuleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SecurityGroupRuleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SecurityGroupRuleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SecurityGroupRuleListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *SecurityGroupRuleListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *SecurityGroupRuleListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *SecurityGroupRuleListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityGroupRuleListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityGroupRuleListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityGroupRuleListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *SecurityGroupRuleListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *SecurityGroupRuleListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSecurityGroupRules

`func (o *SecurityGroupRuleListResponse) GetSecurityGroupRules() []SecurityGroupRule`

GetSecurityGroupRules returns the SecurityGroupRules field if non-nil, zero value otherwise.

### GetSecurityGroupRulesOk

`func (o *SecurityGroupRuleListResponse) GetSecurityGroupRulesOk() (*[]SecurityGroupRule, bool)`

GetSecurityGroupRulesOk returns a tuple with the SecurityGroupRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRules

`func (o *SecurityGroupRuleListResponse) SetSecurityGroupRules(v []SecurityGroupRule)`

SetSecurityGroupRules sets SecurityGroupRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


