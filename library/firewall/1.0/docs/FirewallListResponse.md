# FirewallListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | count | 
**Firewalls** | [**[]Firewall**](Firewall.md) |  | 
**Page** | **int32** | page | 
**Size** | **int32** | size | 
**Sort** | Pointer to **[]string** |  | [optional] 

## Methods

### NewFirewallListResponse

`func NewFirewallListResponse(count int32, firewalls []Firewall, page int32, size int32, ) *FirewallListResponse`

NewFirewallListResponse instantiates a new FirewallListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallListResponseWithDefaults

`func NewFirewallListResponseWithDefaults() *FirewallListResponse`

NewFirewallListResponseWithDefaults instantiates a new FirewallListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FirewallListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FirewallListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FirewallListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetFirewalls

`func (o *FirewallListResponse) GetFirewalls() []Firewall`

GetFirewalls returns the Firewalls field if non-nil, zero value otherwise.

### GetFirewallsOk

`func (o *FirewallListResponse) GetFirewallsOk() (*[]Firewall, bool)`

GetFirewallsOk returns a tuple with the Firewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewalls

`func (o *FirewallListResponse) SetFirewalls(v []Firewall)`

SetFirewalls sets Firewalls field to given value.


### GetPage

`func (o *FirewallListResponse) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FirewallListResponse) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FirewallListResponse) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSize

`func (o *FirewallListResponse) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FirewallListResponse) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FirewallListResponse) SetSize(v int32)`

SetSize sets Size field to given value.


### GetSort

`func (o *FirewallListResponse) GetSort() []string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *FirewallListResponse) GetSortOk() (*[]string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *FirewallListResponse) SetSort(v []string)`

SetSort sets Sort field to given value.

### HasSort

`func (o *FirewallListResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### SetSortNil

`func (o *FirewallListResponse) SetSortNil(b bool)`

 SetSortNil sets the value for Sort to be an explicit nil

### UnsetSort
`func (o *FirewallListResponse) UnsetSort()`

UnsetSort ensures that no value is present for Sort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


