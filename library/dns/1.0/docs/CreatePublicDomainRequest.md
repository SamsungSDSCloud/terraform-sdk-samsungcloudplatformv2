# CreatePublicDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | **string** | Address type | 
**AutoExtension** | Pointer to **NullableBool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DomesticFirstAddressEn** | **NullableString** |  | 
**DomesticFirstAddressKo** | **NullableString** |  | 
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
**Tags** | Pointer to [**[]Tag**](Tag.md) | Tag List | [optional] [default to []]

## Methods

### NewCreatePublicDomainRequest

`func NewCreatePublicDomainRequest(addressType string, domesticFirstAddressEn NullableString, domesticFirstAddressKo NullableString, domesticSecondAddressEn NullableString, domesticSecondAddressKo string, name string, overseasFirstAddress NullableString, overseasSecondAddress NullableString, overseasThirdAddress NullableString, postalCode string, registerEmail string, registerNameEn string, registerNameKo string, registerTelno string, ) *CreatePublicDomainRequest`

NewCreatePublicDomainRequest instantiates a new CreatePublicDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePublicDomainRequestWithDefaults

`func NewCreatePublicDomainRequestWithDefaults() *CreatePublicDomainRequest`

NewCreatePublicDomainRequestWithDefaults instantiates a new CreatePublicDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *CreatePublicDomainRequest) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *CreatePublicDomainRequest) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *CreatePublicDomainRequest) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetAutoExtension

`func (o *CreatePublicDomainRequest) GetAutoExtension() bool`

GetAutoExtension returns the AutoExtension field if non-nil, zero value otherwise.

### GetAutoExtensionOk

`func (o *CreatePublicDomainRequest) GetAutoExtensionOk() (*bool, bool)`

GetAutoExtensionOk returns a tuple with the AutoExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExtension

`func (o *CreatePublicDomainRequest) SetAutoExtension(v bool)`

SetAutoExtension sets AutoExtension field to given value.

### HasAutoExtension

`func (o *CreatePublicDomainRequest) HasAutoExtension() bool`

HasAutoExtension returns a boolean if a field has been set.

### SetAutoExtensionNil

`func (o *CreatePublicDomainRequest) SetAutoExtensionNil(b bool)`

 SetAutoExtensionNil sets the value for AutoExtension to be an explicit nil

### UnsetAutoExtension
`func (o *CreatePublicDomainRequest) UnsetAutoExtension()`

UnsetAutoExtension ensures that no value is present for AutoExtension, not even an explicit nil
### GetDescription

`func (o *CreatePublicDomainRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatePublicDomainRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatePublicDomainRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatePublicDomainRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatePublicDomainRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatePublicDomainRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomesticFirstAddressEn

`func (o *CreatePublicDomainRequest) GetDomesticFirstAddressEn() string`

GetDomesticFirstAddressEn returns the DomesticFirstAddressEn field if non-nil, zero value otherwise.

### GetDomesticFirstAddressEnOk

`func (o *CreatePublicDomainRequest) GetDomesticFirstAddressEnOk() (*string, bool)`

GetDomesticFirstAddressEnOk returns a tuple with the DomesticFirstAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressEn

`func (o *CreatePublicDomainRequest) SetDomesticFirstAddressEn(v string)`

SetDomesticFirstAddressEn sets DomesticFirstAddressEn field to given value.


### SetDomesticFirstAddressEnNil

`func (o *CreatePublicDomainRequest) SetDomesticFirstAddressEnNil(b bool)`

 SetDomesticFirstAddressEnNil sets the value for DomesticFirstAddressEn to be an explicit nil

### UnsetDomesticFirstAddressEn
`func (o *CreatePublicDomainRequest) UnsetDomesticFirstAddressEn()`

UnsetDomesticFirstAddressEn ensures that no value is present for DomesticFirstAddressEn, not even an explicit nil
### GetDomesticFirstAddressKo

`func (o *CreatePublicDomainRequest) GetDomesticFirstAddressKo() string`

GetDomesticFirstAddressKo returns the DomesticFirstAddressKo field if non-nil, zero value otherwise.

### GetDomesticFirstAddressKoOk

`func (o *CreatePublicDomainRequest) GetDomesticFirstAddressKoOk() (*string, bool)`

GetDomesticFirstAddressKoOk returns a tuple with the DomesticFirstAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressKo

`func (o *CreatePublicDomainRequest) SetDomesticFirstAddressKo(v string)`

SetDomesticFirstAddressKo sets DomesticFirstAddressKo field to given value.


### SetDomesticFirstAddressKoNil

`func (o *CreatePublicDomainRequest) SetDomesticFirstAddressKoNil(b bool)`

 SetDomesticFirstAddressKoNil sets the value for DomesticFirstAddressKo to be an explicit nil

### UnsetDomesticFirstAddressKo
`func (o *CreatePublicDomainRequest) UnsetDomesticFirstAddressKo()`

UnsetDomesticFirstAddressKo ensures that no value is present for DomesticFirstAddressKo, not even an explicit nil
### GetDomesticSecondAddressEn

`func (o *CreatePublicDomainRequest) GetDomesticSecondAddressEn() string`

GetDomesticSecondAddressEn returns the DomesticSecondAddressEn field if non-nil, zero value otherwise.

### GetDomesticSecondAddressEnOk

`func (o *CreatePublicDomainRequest) GetDomesticSecondAddressEnOk() (*string, bool)`

GetDomesticSecondAddressEnOk returns a tuple with the DomesticSecondAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressEn

`func (o *CreatePublicDomainRequest) SetDomesticSecondAddressEn(v string)`

SetDomesticSecondAddressEn sets DomesticSecondAddressEn field to given value.


### SetDomesticSecondAddressEnNil

`func (o *CreatePublicDomainRequest) SetDomesticSecondAddressEnNil(b bool)`

 SetDomesticSecondAddressEnNil sets the value for DomesticSecondAddressEn to be an explicit nil

### UnsetDomesticSecondAddressEn
`func (o *CreatePublicDomainRequest) UnsetDomesticSecondAddressEn()`

UnsetDomesticSecondAddressEn ensures that no value is present for DomesticSecondAddressEn, not even an explicit nil
### GetDomesticSecondAddressKo

`func (o *CreatePublicDomainRequest) GetDomesticSecondAddressKo() string`

GetDomesticSecondAddressKo returns the DomesticSecondAddressKo field if non-nil, zero value otherwise.

### GetDomesticSecondAddressKoOk

`func (o *CreatePublicDomainRequest) GetDomesticSecondAddressKoOk() (*string, bool)`

GetDomesticSecondAddressKoOk returns a tuple with the DomesticSecondAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressKo

`func (o *CreatePublicDomainRequest) SetDomesticSecondAddressKo(v string)`

SetDomesticSecondAddressKo sets DomesticSecondAddressKo field to given value.


### GetName

`func (o *CreatePublicDomainRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePublicDomainRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePublicDomainRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOverseasFirstAddress

`func (o *CreatePublicDomainRequest) GetOverseasFirstAddress() string`

GetOverseasFirstAddress returns the OverseasFirstAddress field if non-nil, zero value otherwise.

### GetOverseasFirstAddressOk

`func (o *CreatePublicDomainRequest) GetOverseasFirstAddressOk() (*string, bool)`

GetOverseasFirstAddressOk returns a tuple with the OverseasFirstAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasFirstAddress

`func (o *CreatePublicDomainRequest) SetOverseasFirstAddress(v string)`

SetOverseasFirstAddress sets OverseasFirstAddress field to given value.


### SetOverseasFirstAddressNil

`func (o *CreatePublicDomainRequest) SetOverseasFirstAddressNil(b bool)`

 SetOverseasFirstAddressNil sets the value for OverseasFirstAddress to be an explicit nil

### UnsetOverseasFirstAddress
`func (o *CreatePublicDomainRequest) UnsetOverseasFirstAddress()`

UnsetOverseasFirstAddress ensures that no value is present for OverseasFirstAddress, not even an explicit nil
### GetOverseasSecondAddress

`func (o *CreatePublicDomainRequest) GetOverseasSecondAddress() string`

GetOverseasSecondAddress returns the OverseasSecondAddress field if non-nil, zero value otherwise.

### GetOverseasSecondAddressOk

`func (o *CreatePublicDomainRequest) GetOverseasSecondAddressOk() (*string, bool)`

GetOverseasSecondAddressOk returns a tuple with the OverseasSecondAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasSecondAddress

`func (o *CreatePublicDomainRequest) SetOverseasSecondAddress(v string)`

SetOverseasSecondAddress sets OverseasSecondAddress field to given value.


### SetOverseasSecondAddressNil

`func (o *CreatePublicDomainRequest) SetOverseasSecondAddressNil(b bool)`

 SetOverseasSecondAddressNil sets the value for OverseasSecondAddress to be an explicit nil

### UnsetOverseasSecondAddress
`func (o *CreatePublicDomainRequest) UnsetOverseasSecondAddress()`

UnsetOverseasSecondAddress ensures that no value is present for OverseasSecondAddress, not even an explicit nil
### GetOverseasThirdAddress

`func (o *CreatePublicDomainRequest) GetOverseasThirdAddress() string`

GetOverseasThirdAddress returns the OverseasThirdAddress field if non-nil, zero value otherwise.

### GetOverseasThirdAddressOk

`func (o *CreatePublicDomainRequest) GetOverseasThirdAddressOk() (*string, bool)`

GetOverseasThirdAddressOk returns a tuple with the OverseasThirdAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasThirdAddress

`func (o *CreatePublicDomainRequest) SetOverseasThirdAddress(v string)`

SetOverseasThirdAddress sets OverseasThirdAddress field to given value.


### SetOverseasThirdAddressNil

`func (o *CreatePublicDomainRequest) SetOverseasThirdAddressNil(b bool)`

 SetOverseasThirdAddressNil sets the value for OverseasThirdAddress to be an explicit nil

### UnsetOverseasThirdAddress
`func (o *CreatePublicDomainRequest) UnsetOverseasThirdAddress()`

UnsetOverseasThirdAddress ensures that no value is present for OverseasThirdAddress, not even an explicit nil
### GetPostalCode

`func (o *CreatePublicDomainRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CreatePublicDomainRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CreatePublicDomainRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetRegisterEmail

`func (o *CreatePublicDomainRequest) GetRegisterEmail() string`

GetRegisterEmail returns the RegisterEmail field if non-nil, zero value otherwise.

### GetRegisterEmailOk

`func (o *CreatePublicDomainRequest) GetRegisterEmailOk() (*string, bool)`

GetRegisterEmailOk returns a tuple with the RegisterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterEmail

`func (o *CreatePublicDomainRequest) SetRegisterEmail(v string)`

SetRegisterEmail sets RegisterEmail field to given value.


### GetRegisterNameEn

`func (o *CreatePublicDomainRequest) GetRegisterNameEn() string`

GetRegisterNameEn returns the RegisterNameEn field if non-nil, zero value otherwise.

### GetRegisterNameEnOk

`func (o *CreatePublicDomainRequest) GetRegisterNameEnOk() (*string, bool)`

GetRegisterNameEnOk returns a tuple with the RegisterNameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNameEn

`func (o *CreatePublicDomainRequest) SetRegisterNameEn(v string)`

SetRegisterNameEn sets RegisterNameEn field to given value.


### GetRegisterNameKo

`func (o *CreatePublicDomainRequest) GetRegisterNameKo() string`

GetRegisterNameKo returns the RegisterNameKo field if non-nil, zero value otherwise.

### GetRegisterNameKoOk

`func (o *CreatePublicDomainRequest) GetRegisterNameKoOk() (*string, bool)`

GetRegisterNameKoOk returns a tuple with the RegisterNameKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNameKo

`func (o *CreatePublicDomainRequest) SetRegisterNameKo(v string)`

SetRegisterNameKo sets RegisterNameKo field to given value.


### GetRegisterTelno

`func (o *CreatePublicDomainRequest) GetRegisterTelno() string`

GetRegisterTelno returns the RegisterTelno field if non-nil, zero value otherwise.

### GetRegisterTelnoOk

`func (o *CreatePublicDomainRequest) GetRegisterTelnoOk() (*string, bool)`

GetRegisterTelnoOk returns a tuple with the RegisterTelno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterTelno

`func (o *CreatePublicDomainRequest) SetRegisterTelno(v string)`

SetRegisterTelno sets RegisterTelno field to given value.


### GetTags

`func (o *CreatePublicDomainRequest) GetTags() []Tag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePublicDomainRequest) GetTagsOk() (*[]Tag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePublicDomainRequest) SetTags(v []Tag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreatePublicDomainRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


