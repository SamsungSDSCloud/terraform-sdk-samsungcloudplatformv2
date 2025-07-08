# FirewallRuleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**FirewallRules** | [**[]FirewallRule**](FirewallRule.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirewallRuleListResponse

`func NewFirewallRuleListResponse(count int32, firewallRules []FirewallRule, page int32, size int32, ) *FirewallRuleListResponse`

NewFirewallRuleListResponse instantiates a new FirewallRuleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleListResponseWithDefaults

`func NewFirewallRuleListResponseWithDefaults() *FirewallRuleListResponse`

NewFirewallRuleListResponseWithDefaults instantiates a new FirewallRuleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FirewallRuleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FirewallRuleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FirewallRuleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFirewallRules

`func (o *FirewallRuleListResponse) GetFirewallRules() []FirewallRule`

GetFirewallRules returns the FirewallRules field if non-nil, zero value otherwise.

### GetFirewallRulesOk

`func (o *FirewallRuleListResponse) GetFirewallRulesOk() (*[]FirewallRule, bool)`

GetFirewallRulesOk returns a tuple with the FirewallRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallRules

`func (o *FirewallRuleListResponse) SetFirewallRules(v []FirewallRule)`

SetFirewallRules sets FirewallRules field to given value.


### GetPage

`func (o *FirewallRuleListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FirewallRuleListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FirewallRuleListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *FirewallRuleListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirewallRuleListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirewallRuleListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *FirewallRuleListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *FirewallRuleListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *FirewallRuleListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *FirewallRuleListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *FirewallRuleListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *FirewallRuleListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


