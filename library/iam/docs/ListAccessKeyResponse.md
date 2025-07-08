# ListAccessKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKeys** | [**[]AccessKey**](AccessKey.md) |  | 
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewListAccessKeyResponse

`func NewListAccessKeyResponse(accessKeys []AccessKey, ) *ListAccessKeyResponse`

NewListAccessKeyResponse instantiates a new ListAccessKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccessKeyResponseWithDefaults

`func NewListAccessKeyResponseWithDefaults() *ListAccessKeyResponse`

NewListAccessKeyResponseWithDefaults instantiates a new ListAccessKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKeys

`func (o *ListAccessKeyResponse) GetAccessKeys() []AccessKey`

GetAccessKeys returns the AccessKeys field if non-nil, zero value otherwise.

### GetAccessKeysOk

`func (o *ListAccessKeyResponse) GetAccessKeysOk() (*[]AccessKey, bool)`

GetAccessKeysOk returns a tuple with the AccessKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeys

`func (o *ListAccessKeyResponse) SetAccessKeys(v []AccessKey)`

SetAccessKeys sets AccessKeys field to given value.


### GetCount

`func (o *ListAccessKeyResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListAccessKeyResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListAccessKeyResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListAccessKeyResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *ListAccessKeyResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *ListAccessKeyResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *ListAccessKeyResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListAccessKeyResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListAccessKeyResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListAccessKeyResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *ListAccessKeyResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *ListAccessKeyResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


