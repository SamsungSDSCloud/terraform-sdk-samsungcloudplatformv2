# SecretListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **NullableInt32** |  | [optional] 
**Links** | Pointer to **[]interface{}** |  | [optional] 
**Secrets** | [**[]SecretSummary**](SecretSummary.md) |  | 

## Methods

### NewSecretListResponse

`func NewSecretListResponse(secrets []SecretSummary, ) *SecretListResponse`

NewSecretListResponse instantiates a new SecretListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretListResponseWithDefaults

`func NewSecretListResponseWithDefaults() *SecretListResponse`

NewSecretListResponseWithDefaults instantiates a new SecretListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *SecretListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *SecretListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *SecretListResponse) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *SecretListResponse) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *SecretListResponse) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *SecretListResponse) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetLinks

`func (o *SecretListResponse) GetLinks() []interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecretListResponse) GetLinksOk() (*[]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecretListResponse) SetLinks(v []interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecretListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *SecretListResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *SecretListResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetSecrets

`func (o *SecretListResponse) GetSecrets() []SecretSummary`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *SecretListResponse) GetSecretsOk() (*[]SecretSummary, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *SecretListResponse) SetSecrets(v []SecretSummary)`

SetSecrets sets Secrets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


