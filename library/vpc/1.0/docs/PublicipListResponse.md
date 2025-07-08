# PublicipListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**Publicips** | [**[]Publicip**](Publicip.md) |  | 

## Methods

### NewPublicipListResponse

`func NewPublicipListResponse(publicips []Publicip, ) *PublicipListResponse`

NewPublicipListResponse instantiates a new PublicipListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicipListResponseWithDefaults

`func NewPublicipListResponseWithDefaults() *PublicipListResponse`

NewPublicipListResponseWithDefaults instantiates a new PublicipListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PublicipListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PublicipListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PublicipListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PublicipListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PublicipListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PublicipListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *PublicipListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PublicipListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PublicipListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PublicipListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *PublicipListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *PublicipListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetPublicips

`func (o *PublicipListResponse) GetPublicips() []Publicip`

GetPublicips returns the Publicips field if non-nil, zero value otherwise.

### GetPublicipsOk

`func (o *PublicipListResponse) GetPublicipsOk() (*[]Publicip, bool)`

GetPublicipsOk returns a tuple with the Publicips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicips

`func (o *PublicipListResponse) SetPublicips(v []Publicip)`

SetPublicips sets Publicips field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


