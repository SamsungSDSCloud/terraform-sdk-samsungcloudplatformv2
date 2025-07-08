# SecurityGroupQuotaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupRules** | [**Quota**](Quota.md) |  | 
**SecurityGroupRulesPerSg** | [**NullableQuota**](Quota.md) |  | 
**SecurityGroups** | [**Quota**](Quota.md) |  | 

## Methods

### NewSecurityGroupQuotaResponse

`func NewSecurityGroupQuotaResponse(securityGroupRules Quota, securityGroupRulesPerSg NullableQuota, securityGroups Quota, ) *SecurityGroupQuotaResponse`

NewSecurityGroupQuotaResponse instantiates a new SecurityGroupQuotaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupQuotaResponseWithDefaults

`func NewSecurityGroupQuotaResponseWithDefaults() *SecurityGroupQuotaResponse`

NewSecurityGroupQuotaResponseWithDefaults instantiates a new SecurityGroupQuotaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupRules

`func (o *SecurityGroupQuotaResponse) GetSecurityGroupRules() Quota`

GetSecurityGroupRules returns the SecurityGroupRules field if non-nil, zero value otherwise.

### GetSecurityGroupRulesOk

`func (o *SecurityGroupQuotaResponse) GetSecurityGroupRulesOk() (*Quota, bool)`

GetSecurityGroupRulesOk returns a tuple with the SecurityGroupRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRules

`func (o *SecurityGroupQuotaResponse) SetSecurityGroupRules(v Quota)`

SetSecurityGroupRules sets SecurityGroupRules field to given value.


### GetSecurityGroupRulesPerSg

`func (o *SecurityGroupQuotaResponse) GetSecurityGroupRulesPerSg() Quota`

GetSecurityGroupRulesPerSg returns the SecurityGroupRulesPerSg field if non-nil, zero value otherwise.

### GetSecurityGroupRulesPerSgOk

`func (o *SecurityGroupQuotaResponse) GetSecurityGroupRulesPerSgOk() (*Quota, bool)`

GetSecurityGroupRulesPerSgOk returns a tuple with the SecurityGroupRulesPerSg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupRulesPerSg

`func (o *SecurityGroupQuotaResponse) SetSecurityGroupRulesPerSg(v Quota)`

SetSecurityGroupRulesPerSg sets SecurityGroupRulesPerSg field to given value.


### SetSecurityGroupRulesPerSgNil

`func (o *SecurityGroupQuotaResponse) SetSecurityGroupRulesPerSgNil(b bool)`

 SetSecurityGroupRulesPerSgNil sets the value for SecurityGroupRulesPerSg to be an explicit nil

### UnsetSecurityGroupRulesPerSg
`func (o *SecurityGroupQuotaResponse) UnsetSecurityGroupRulesPerSg()`

UnsetSecurityGroupRulesPerSg ensures that no value is present for SecurityGroupRulesPerSg, not even an explicit nil
### GetSecurityGroups

`func (o *SecurityGroupQuotaResponse) GetSecurityGroups() Quota`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *SecurityGroupQuotaResponse) GetSecurityGroupsOk() (*Quota, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *SecurityGroupQuotaResponse) SetSecurityGroups(v Quota)`

SetSecurityGroups sets SecurityGroups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


