# PubblicDomainWhoisInfoUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | **string** | Address type | 
**DomesticFirstAddressEn** | **string** | Domestic first address (EN) | 
**DomesticFirstAddressKo** | **string** | Domestic first address (KO) | 
**DomesticSecondAddressEn** | **NullableString** |  | 
**DomesticSecondAddressKo** | **string** | Domestic second address (KO) | 
**Name** | **string** | name | 
**OverseasFirstAddress** | **NullableString** |  | 
**OverseasSecondAddress** | **NullableString** |  | 
**OverseasThirdAddress** | **NullableString** |  | 
**PostalCode** | **string** | Postal code | 
**RegisterEmail** | **string** | Register email | 
**RegisterNameEn** | **string** | Register Name Overseas | 
**RegisterNameKo** | **string** | Register Name Domestic | 
**RegisterTelno** | **string** | Register telephone number | 

## Methods

### NewPubblicDomainWhoisInfoUpdateResponse

`func NewPubblicDomainWhoisInfoUpdateResponse(addressType string, domesticFirstAddressEn string, domesticFirstAddressKo string, domesticSecondAddressEn NullableString, domesticSecondAddressKo string, name string, overseasFirstAddress NullableString, overseasSecondAddress NullableString, overseasThirdAddress NullableString, postalCode string, registerEmail string, registerNameEn string, registerNameKo string, registerTelno string, ) *PubblicDomainWhoisInfoUpdateResponse`

NewPubblicDomainWhoisInfoUpdateResponse instantiates a new PubblicDomainWhoisInfoUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubblicDomainWhoisInfoUpdateResponseWithDefaults

`func NewPubblicDomainWhoisInfoUpdateResponseWithDefaults() *PubblicDomainWhoisInfoUpdateResponse`

NewPubblicDomainWhoisInfoUpdateResponseWithDefaults instantiates a new PubblicDomainWhoisInfoUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetDomesticFirstAddressEn

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticFirstAddressEn() string`

GetDomesticFirstAddressEn returns the DomesticFirstAddressEn field if non-nil, zero value otherwise.

### GetDomesticFirstAddressEnOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticFirstAddressEnOk() (*string, bool)`

GetDomesticFirstAddressEnOk returns a tuple with the DomesticFirstAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressEn

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetDomesticFirstAddressEn(v string)`

SetDomesticFirstAddressEn sets DomesticFirstAddressEn field to given value.


### GetDomesticFirstAddressKo

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticFirstAddressKo() string`

GetDomesticFirstAddressKo returns the DomesticFirstAddressKo field if non-nil, zero value otherwise.

### GetDomesticFirstAddressKoOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticFirstAddressKoOk() (*string, bool)`

GetDomesticFirstAddressKoOk returns a tuple with the DomesticFirstAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressKo

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetDomesticFirstAddressKo(v string)`

SetDomesticFirstAddressKo sets DomesticFirstAddressKo field to given value.


### GetDomesticSecondAddressEn

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticSecondAddressEn() string`

GetDomesticSecondAddressEn returns the DomesticSecondAddressEn field if non-nil, zero value otherwise.

### GetDomesticSecondAddressEnOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticSecondAddressEnOk() (*string, bool)`

GetDomesticSecondAddressEnOk returns a tuple with the DomesticSecondAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressEn

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetDomesticSecondAddressEn(v string)`

SetDomesticSecondAddressEn sets DomesticSecondAddressEn field to given value.


### SetDomesticSecondAddressEnNil

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetDomesticSecondAddressEnNil(b bool)`

 SetDomesticSecondAddressEnNil sets the value for DomesticSecondAddressEn to be an explicit nil

### UnsetDomesticSecondAddressEn
`func (o *PubblicDomainWhoisInfoUpdateResponse) UnsetDomesticSecondAddressEn()`

UnsetDomesticSecondAddressEn ensures that no value is present for DomesticSecondAddressEn, not even an explicit nil
### GetDomesticSecondAddressKo

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticSecondAddressKo() string`

GetDomesticSecondAddressKo returns the DomesticSecondAddressKo field if non-nil, zero value otherwise.

### GetDomesticSecondAddressKoOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetDomesticSecondAddressKoOk() (*string, bool)`

GetDomesticSecondAddressKoOk returns a tuple with the DomesticSecondAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressKo

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetDomesticSecondAddressKo(v string)`

SetDomesticSecondAddressKo sets DomesticSecondAddressKo field to given value.


### GetName

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOverseasFirstAddress

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetOverseasFirstAddress() string`

GetOverseasFirstAddress returns the OverseasFirstAddress field if non-nil, zero value otherwise.

### GetOverseasFirstAddressOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetOverseasFirstAddressOk() (*string, bool)`

GetOverseasFirstAddressOk returns a tuple with the OverseasFirstAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasFirstAddress

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetOverseasFirstAddress(v string)`

SetOverseasFirstAddress sets OverseasFirstAddress field to given value.


### SetOverseasFirstAddressNil

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetOverseasFirstAddressNil(b bool)`

 SetOverseasFirstAddressNil sets the value for OverseasFirstAddress to be an explicit nil

### UnsetOverseasFirstAddress
`func (o *PubblicDomainWhoisInfoUpdateResponse) UnsetOverseasFirstAddress()`

UnsetOverseasFirstAddress ensures that no value is present for OverseasFirstAddress, not even an explicit nil
### GetOverseasSecondAddress

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetOverseasSecondAddress() string`

GetOverseasSecondAddress returns the OverseasSecondAddress field if non-nil, zero value otherwise.

### GetOverseasSecondAddressOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetOverseasSecondAddressOk() (*string, bool)`

GetOverseasSecondAddressOk returns a tuple with the OverseasSecondAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasSecondAddress

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetOverseasSecondAddress(v string)`

SetOverseasSecondAddress sets OverseasSecondAddress field to given value.


### SetOverseasSecondAddressNil

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetOverseasSecondAddressNil(b bool)`

 SetOverseasSecondAddressNil sets the value for OverseasSecondAddress to be an explicit nil

### UnsetOverseasSecondAddress
`func (o *PubblicDomainWhoisInfoUpdateResponse) UnsetOverseasSecondAddress()`

UnsetOverseasSecondAddress ensures that no value is present for OverseasSecondAddress, not even an explicit nil
### GetOverseasThirdAddress

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetOverseasThirdAddress() string`

GetOverseasThirdAddress returns the OverseasThirdAddress field if non-nil, zero value otherwise.

### GetOverseasThirdAddressOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetOverseasThirdAddressOk() (*string, bool)`

GetOverseasThirdAddressOk returns a tuple with the OverseasThirdAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasThirdAddress

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetOverseasThirdAddress(v string)`

SetOverseasThirdAddress sets OverseasThirdAddress field to given value.


### SetOverseasThirdAddressNil

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetOverseasThirdAddressNil(b bool)`

 SetOverseasThirdAddressNil sets the value for OverseasThirdAddress to be an explicit nil

### UnsetOverseasThirdAddress
`func (o *PubblicDomainWhoisInfoUpdateResponse) UnsetOverseasThirdAddress()`

UnsetOverseasThirdAddress ensures that no value is present for OverseasThirdAddress, not even an explicit nil
### GetPostalCode

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetRegisterEmail

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterEmail() string`

GetRegisterEmail returns the RegisterEmail field if non-nil, zero value otherwise.

### GetRegisterEmailOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterEmailOk() (*string, bool)`

GetRegisterEmailOk returns a tuple with the RegisterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterEmail

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetRegisterEmail(v string)`

SetRegisterEmail sets RegisterEmail field to given value.


### GetRegisterNameEn

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterNameEn() string`

GetRegisterNameEn returns the RegisterNameEn field if non-nil, zero value otherwise.

### GetRegisterNameEnOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterNameEnOk() (*string, bool)`

GetRegisterNameEnOk returns a tuple with the RegisterNameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNameEn

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetRegisterNameEn(v string)`

SetRegisterNameEn sets RegisterNameEn field to given value.


### GetRegisterNameKo

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterNameKo() string`

GetRegisterNameKo returns the RegisterNameKo field if non-nil, zero value otherwise.

### GetRegisterNameKoOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterNameKoOk() (*string, bool)`

GetRegisterNameKoOk returns a tuple with the RegisterNameKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNameKo

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetRegisterNameKo(v string)`

SetRegisterNameKo sets RegisterNameKo field to given value.


### GetRegisterTelno

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterTelno() string`

GetRegisterTelno returns the RegisterTelno field if non-nil, zero value otherwise.

### GetRegisterTelnoOk

`func (o *PubblicDomainWhoisInfoUpdateResponse) GetRegisterTelnoOk() (*string, bool)`

GetRegisterTelnoOk returns a tuple with the RegisterTelno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterTelno

`func (o *PubblicDomainWhoisInfoUpdateResponse) SetRegisterTelno(v string)`

SetRegisterTelno sets RegisterTelno field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


