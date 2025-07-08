# PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiGroupForSort** | Pointer to **string** | Api Group For Sort | [optional] [default to ""]
**NonResourceUrls** | Pointer to **[]string** | Non Resource Urls | [optional] [default to []]
**Resource** | Pointer to **string** | Resource | [optional] [default to ""]
**ResourceNames** | Pointer to **[]string** | Resource Names | [optional] [default to []]
**Verbs** | Pointer to **[]string** | Verbs | [optional] [default to []]

## Methods

### NewPolicyRule

`func NewPolicyRule() *PolicyRule`

NewPolicyRule instantiates a new PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleWithDefaults

`func NewPolicyRuleWithDefaults() *PolicyRule`

NewPolicyRuleWithDefaults instantiates a new PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiGroupForSort

`func (o *PolicyRule) GetApiGroupForSort() string`

GetApiGroupForSort returns the ApiGroupForSort field if non-nil, zero value otherwise.

### GetApiGroupForSortOk

`func (o *PolicyRule) GetApiGroupForSortOk() (*string, bool)`

GetApiGroupForSortOk returns a tuple with the ApiGroupForSort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiGroupForSort

`func (o *PolicyRule) SetApiGroupForSort(v string)`

SetApiGroupForSort sets ApiGroupForSort field to given value.

### HasApiGroupForSort

`func (o *PolicyRule) HasApiGroupForSort() bool`

HasApiGroupForSort returns a boolean if a field has been set.

### GetNonResourceUrls

`func (o *PolicyRule) GetNonResourceUrls() []string`

GetNonResourceUrls returns the NonResourceUrls field if non-nil, zero value otherwise.

### GetNonResourceUrlsOk

`func (o *PolicyRule) GetNonResourceUrlsOk() (*[]string, bool)`

GetNonResourceUrlsOk returns a tuple with the NonResourceUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceUrls

`func (o *PolicyRule) SetNonResourceUrls(v []string)`

SetNonResourceUrls sets NonResourceUrls field to given value.

### HasNonResourceUrls

`func (o *PolicyRule) HasNonResourceUrls() bool`

HasNonResourceUrls returns a boolean if a field has been set.

### GetResource

`func (o *PolicyRule) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *PolicyRule) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *PolicyRule) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *PolicyRule) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceNames

`func (o *PolicyRule) GetResourceNames() []string`

GetResourceNames returns the ResourceNames field if non-nil, zero value otherwise.

### GetResourceNamesOk

`func (o *PolicyRule) GetResourceNamesOk() (*[]string, bool)`

GetResourceNamesOk returns a tuple with the ResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNames

`func (o *PolicyRule) SetResourceNames(v []string)`

SetResourceNames sets ResourceNames field to given value.

### HasResourceNames

`func (o *PolicyRule) HasResourceNames() bool`

HasResourceNames returns a boolean if a field has been set.

### GetVerbs

`func (o *PolicyRule) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *PolicyRule) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *PolicyRule) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.

### HasVerbs

`func (o *PolicyRule) HasVerbs() bool`

HasVerbs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


