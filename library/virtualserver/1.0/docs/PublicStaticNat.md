# PublicStaticNat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIpAddress** | **string** | External IP Address | 
**Id** | **string** | NAT ID | 
**PublicipId** | Pointer to **NullableString** |  | [optional] 
**State** | **string** | NAT State | 

## Methods

### NewPublicStaticNat

`func NewPublicStaticNat(externalIpAddress string, id string, state string, ) *PublicStaticNat`

NewPublicStaticNat instantiates a new PublicStaticNat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicStaticNatWithDefaults

`func NewPublicStaticNatWithDefaults() *PublicStaticNat`

NewPublicStaticNatWithDefaults instantiates a new PublicStaticNat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIpAddress

`func (o *PublicStaticNat) GetExternalIpAddress() string`

GetExternalIpAddress returns the ExternalIpAddress field if non-nil, zero value otherwise.

### GetExternalIpAddressOk

`func (o *PublicStaticNat) GetExternalIpAddressOk() (*string, bool)`

GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddress

`func (o *PublicStaticNat) SetExternalIpAddress(v string)`

SetExternalIpAddress sets ExternalIpAddress field to given value.


### GetId

`func (o *PublicStaticNat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicStaticNat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicStaticNat) SetId(v string)`

SetId sets Id field to given value.


### GetPublicipId

`func (o *PublicStaticNat) GetPublicipId() string`

GetPublicipId returns the PublicipId field if non-nil, zero value otherwise.

### GetPublicipIdOk

`func (o *PublicStaticNat) GetPublicipIdOk() (*string, bool)`

GetPublicipIdOk returns a tuple with the PublicipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicipId

`func (o *PublicStaticNat) SetPublicipId(v string)`

SetPublicipId sets PublicipId field to given value.

### HasPublicipId

`func (o *PublicStaticNat) HasPublicipId() bool`

HasPublicipId returns a boolean if a field has been set.

### SetPublicipIdNil

`func (o *PublicStaticNat) SetPublicipIdNil(b bool)`

 SetPublicipIdNil sets the value for PublicipId to be an explicit nil

### UnsetPublicipId
`func (o *PublicStaticNat) UnsetPublicipId()`

UnsetPublicipId ensures that no value is present for PublicipId, not even an explicit nil
### GetState

`func (o *PublicStaticNat) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PublicStaticNat) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PublicStaticNat) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


