# PrivateStaticNat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalIpAddress** | **string** | External IP Address | 
**Id** | **string** | NAT ID | 
**PrivateNatIpId** | Pointer to **NullableString** |  | [optional] 
**State** | **string** | NAT State | 

## Methods

### NewPrivateStaticNat

`func NewPrivateStaticNat(externalIpAddress string, id string, state string, ) *PrivateStaticNat`

NewPrivateStaticNat instantiates a new PrivateStaticNat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateStaticNatWithDefaults

`func NewPrivateStaticNatWithDefaults() *PrivateStaticNat`

NewPrivateStaticNatWithDefaults instantiates a new PrivateStaticNat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalIpAddress

`func (o *PrivateStaticNat) GetExternalIpAddress() string`

GetExternalIpAddress returns the ExternalIpAddress field if non-nil, zero value otherwise.

### GetExternalIpAddressOk

`func (o *PrivateStaticNat) GetExternalIpAddressOk() (*string, bool)`

GetExternalIpAddressOk returns a tuple with the ExternalIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpAddress

`func (o *PrivateStaticNat) SetExternalIpAddress(v string)`

SetExternalIpAddress sets ExternalIpAddress field to given value.


### GetId

`func (o *PrivateStaticNat) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateStaticNat) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateStaticNat) SetId(v string)`

SetId sets Id field to given value.


### GetPrivateNatIpId

`func (o *PrivateStaticNat) GetPrivateNatIpId() string`

GetPrivateNatIpId returns the PrivateNatIpId field if non-nil, zero value otherwise.

### GetPrivateNatIpIdOk

`func (o *PrivateStaticNat) GetPrivateNatIpIdOk() (*string, bool)`

GetPrivateNatIpIdOk returns a tuple with the PrivateNatIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNatIpId

`func (o *PrivateStaticNat) SetPrivateNatIpId(v string)`

SetPrivateNatIpId sets PrivateNatIpId field to given value.

### HasPrivateNatIpId

`func (o *PrivateStaticNat) HasPrivateNatIpId() bool`

HasPrivateNatIpId returns a boolean if a field has been set.

### SetPrivateNatIpIdNil

`func (o *PrivateStaticNat) SetPrivateNatIpIdNil(b bool)`

 SetPrivateNatIpIdNil sets the value for PrivateNatIpId to be an explicit nil

### UnsetPrivateNatIpId
`func (o *PrivateStaticNat) UnsetPrivateNatIpId()`

UnsetPrivateNatIpId ensures that no value is present for PrivateNatIpId, not even an explicit nil
### GetState

`func (o *PrivateStaticNat) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PrivateStaticNat) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PrivateStaticNat) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


