# TransitGatewayRuleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 
**TransitGatewayRules** | [**[]TransitGatewayRule**](TransitGatewayRule.md) |  | 

## Methods

### NewTransitGatewayRuleListResponse

`func NewTransitGatewayRuleListResponse(count int32, page int32, size int32, transitGatewayRules []TransitGatewayRule, ) *TransitGatewayRuleListResponse`

NewTransitGatewayRuleListResponse instantiates a new TransitGatewayRuleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayRuleListResponseWithDefaults

`func NewTransitGatewayRuleListResponseWithDefaults() *TransitGatewayRuleListResponse`

NewTransitGatewayRuleListResponseWithDefaults instantiates a new TransitGatewayRuleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *TransitGatewayRuleListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransitGatewayRuleListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransitGatewayRuleListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetPage

`func (o *TransitGatewayRuleListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *TransitGatewayRuleListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *TransitGatewayRuleListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *TransitGatewayRuleListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *TransitGatewayRuleListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *TransitGatewayRuleListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *TransitGatewayRuleListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *TransitGatewayRuleListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *TransitGatewayRuleListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *TransitGatewayRuleListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *TransitGatewayRuleListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *TransitGatewayRuleListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil
### GetTransitGatewayRules

`func (o *TransitGatewayRuleListResponse) GetTransitGatewayRules() []TransitGatewayRule`

GetTransitGatewayRules returns the TransitGatewayRules field if non-nil, zero value otherwise.

### GetTransitGatewayRulesOk

`func (o *TransitGatewayRuleListResponse) GetTransitGatewayRulesOk() (*[]TransitGatewayRule, bool)`

GetTransitGatewayRulesOk returns a tuple with the TransitGatewayRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayRules

`func (o *TransitGatewayRuleListResponse) SetTransitGatewayRules(v []TransitGatewayRule)`

SetTransitGatewayRules sets TransitGatewayRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


