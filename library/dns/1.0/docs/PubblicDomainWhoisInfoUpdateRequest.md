# PubblicDomainWhoisInfoUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | **string** | Address type | 
**DomesticFirstAddressEn** | **string** | Domestic first address (EN) | 
**DomesticFirstAddressKo** | **string** | Domestic first address (KO) | 
**DomesticSecondAddressEn** | **NullableString** |  | 
**DomesticSecondAddressKo** | **string** | Domestic second address (KO) | 
**OverseasFirstAddress** | **NullableString** |  | 
**OverseasSecondAddress** | **NullableString** |  | 
**OverseasThirdAddress** | **NullableString** |  | 
**PostalCode** | **string** | Postal code | 
**RegisterEmail** | **string** | Register email | 
**RegisterTelno** | **string** | Register telephone number | 

## Methods

### NewPubblicDomainWhoisInfoUpdateRequest

`func NewPubblicDomainWhoisInfoUpdateRequest(addressType string, domesticFirstAddressEn string, domesticFirstAddressKo string, domesticSecondAddressEn NullableString, domesticSecondAddressKo string, overseasFirstAddress NullableString, overseasSecondAddress NullableString, overseasThirdAddress NullableString, postalCode string, registerEmail string, registerTelno string, ) *PubblicDomainWhoisInfoUpdateRequest`

NewPubblicDomainWhoisInfoUpdateRequest instantiates a new PubblicDomainWhoisInfoUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubblicDomainWhoisInfoUpdateRequestWithDefaults

`func NewPubblicDomainWhoisInfoUpdateRequestWithDefaults() *PubblicDomainWhoisInfoUpdateRequest`

NewPubblicDomainWhoisInfoUpdateRequestWithDefaults instantiates a new PubblicDomainWhoisInfoUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetDomesticFirstAddressEn

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticFirstAddressEn() string`

GetDomesticFirstAddressEn returns the DomesticFirstAddressEn field if non-nil, zero value otherwise.

### GetDomesticFirstAddressEnOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticFirstAddressEnOk() (*string, bool)`

GetDomesticFirstAddressEnOk returns a tuple with the DomesticFirstAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressEn

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetDomesticFirstAddressEn(v string)`

SetDomesticFirstAddressEn sets DomesticFirstAddressEn field to given value.


### GetDomesticFirstAddressKo

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticFirstAddressKo() string`

GetDomesticFirstAddressKo returns the DomesticFirstAddressKo field if non-nil, zero value otherwise.

### GetDomesticFirstAddressKoOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticFirstAddressKoOk() (*string, bool)`

GetDomesticFirstAddressKoOk returns a tuple with the DomesticFirstAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressKo

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetDomesticFirstAddressKo(v string)`

SetDomesticFirstAddressKo sets DomesticFirstAddressKo field to given value.


### GetDomesticSecondAddressEn

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticSecondAddressEn() string`

GetDomesticSecondAddressEn returns the DomesticSecondAddressEn field if non-nil, zero value otherwise.

### GetDomesticSecondAddressEnOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticSecondAddressEnOk() (*string, bool)`

GetDomesticSecondAddressEnOk returns a tuple with the DomesticSecondAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressEn

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetDomesticSecondAddressEn(v string)`

SetDomesticSecondAddressEn sets DomesticSecondAddressEn field to given value.


### SetDomesticSecondAddressEnNil

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetDomesticSecondAddressEnNil(b bool)`

 SetDomesticSecondAddressEnNil sets the value for DomesticSecondAddressEn to be an explicit nil

### UnsetDomesticSecondAddressEn
`func (o *PubblicDomainWhoisInfoUpdateRequest) UnsetDomesticSecondAddressEn()`

UnsetDomesticSecondAddressEn ensures that no value is present for DomesticSecondAddressEn, not even an explicit nil
### GetDomesticSecondAddressKo

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticSecondAddressKo() string`

GetDomesticSecondAddressKo returns the DomesticSecondAddressKo field if non-nil, zero value otherwise.

### GetDomesticSecondAddressKoOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetDomesticSecondAddressKoOk() (*string, bool)`

GetDomesticSecondAddressKoOk returns a tuple with the DomesticSecondAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressKo

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetDomesticSecondAddressKo(v string)`

SetDomesticSecondAddressKo sets DomesticSecondAddressKo field to given value.


### GetOverseasFirstAddress

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetOverseasFirstAddress() string`

GetOverseasFirstAddress returns the OverseasFirstAddress field if non-nil, zero value otherwise.

### GetOverseasFirstAddressOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetOverseasFirstAddressOk() (*string, bool)`

GetOverseasFirstAddressOk returns a tuple with the OverseasFirstAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasFirstAddress

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetOverseasFirstAddress(v string)`

SetOverseasFirstAddress sets OverseasFirstAddress field to given value.


### SetOverseasFirstAddressNil

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetOverseasFirstAddressNil(b bool)`

 SetOverseasFirstAddressNil sets the value for OverseasFirstAddress to be an explicit nil

### UnsetOverseasFirstAddress
`func (o *PubblicDomainWhoisInfoUpdateRequest) UnsetOverseasFirstAddress()`

UnsetOverseasFirstAddress ensures that no value is present for OverseasFirstAddress, not even an explicit nil
### GetOverseasSecondAddress

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetOverseasSecondAddress() string`

GetOverseasSecondAddress returns the OverseasSecondAddress field if non-nil, zero value otherwise.

### GetOverseasSecondAddressOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetOverseasSecondAddressOk() (*string, bool)`

GetOverseasSecondAddressOk returns a tuple with the OverseasSecondAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasSecondAddress

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetOverseasSecondAddress(v string)`

SetOverseasSecondAddress sets OverseasSecondAddress field to given value.


### SetOverseasSecondAddressNil

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetOverseasSecondAddressNil(b bool)`

 SetOverseasSecondAddressNil sets the value for OverseasSecondAddress to be an explicit nil

### UnsetOverseasSecondAddress
`func (o *PubblicDomainWhoisInfoUpdateRequest) UnsetOverseasSecondAddress()`

UnsetOverseasSecondAddress ensures that no value is present for OverseasSecondAddress, not even an explicit nil
### GetOverseasThirdAddress

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetOverseasThirdAddress() string`

GetOverseasThirdAddress returns the OverseasThirdAddress field if non-nil, zero value otherwise.

### GetOverseasThirdAddressOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetOverseasThirdAddressOk() (*string, bool)`

GetOverseasThirdAddressOk returns a tuple with the OverseasThirdAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasThirdAddress

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetOverseasThirdAddress(v string)`

SetOverseasThirdAddress sets OverseasThirdAddress field to given value.


### SetOverseasThirdAddressNil

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetOverseasThirdAddressNil(b bool)`

 SetOverseasThirdAddressNil sets the value for OverseasThirdAddress to be an explicit nil

### UnsetOverseasThirdAddress
`func (o *PubblicDomainWhoisInfoUpdateRequest) UnsetOverseasThirdAddress()`

UnsetOverseasThirdAddress ensures that no value is present for OverseasThirdAddress, not even an explicit nil
### GetPostalCode

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetRegisterEmail

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetRegisterEmail() string`

GetRegisterEmail returns the RegisterEmail field if non-nil, zero value otherwise.

### GetRegisterEmailOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetRegisterEmailOk() (*string, bool)`

GetRegisterEmailOk returns a tuple with the RegisterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterEmail

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetRegisterEmail(v string)`

SetRegisterEmail sets RegisterEmail field to given value.


### GetRegisterTelno

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetRegisterTelno() string`

GetRegisterTelno returns the RegisterTelno field if non-nil, zero value otherwise.

### GetRegisterTelnoOk

`func (o *PubblicDomainWhoisInfoUpdateRequest) GetRegisterTelnoOk() (*string, bool)`

GetRegisterTelnoOk returns a tuple with the RegisterTelno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterTelno

`func (o *PubblicDomainWhoisInfoUpdateRequest) SetRegisterTelno(v string)`

SetRegisterTelno sets RegisterTelno field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


