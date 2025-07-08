# ListCompanyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Companies** | [**[]CompanyResponse**](CompanyResponse.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewListCompanyResponse

`func NewListCompanyResponse(companies []CompanyResponse, ) *ListCompanyResponse`

NewListCompanyResponse instantiates a new ListCompanyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompanyResponseWithDefaults

`func NewListCompanyResponseWithDefaults() *ListCompanyResponse`

NewListCompanyResponseWithDefaults instantiates a new ListCompanyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanies

`func (o *ListCompanyResponse) GetCompanies() []CompanyResponse`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *ListCompanyResponse) GetCompaniesOk() (*[]CompanyResponse, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *ListCompanyResponse) SetCompanies(v []CompanyResponse)`

SetCompanies sets Companies field to given value.


### GetCount

`func (o *ListCompanyResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListCompanyResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListCompanyResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListCompanyResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ListCompanyResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ListCompanyResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *ListCompanyResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListCompanyResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListCompanyResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListCompanyResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ListCompanyResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ListCompanyResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


