# PublicDomainDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressType** | **string** | Address type | 
**AutoExtension** | **bool** | Auto extension flag | 
**CreatedAt** | **string** | Created Date | 
**CreatedBy** | **string** | Created By_ | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DomesticAddressEn** | **string** | Domestic Address (EN) | 
**DomesticAddressKo** | **string** | Domestic Address (KO) | 
**DomesticFirstAddressEn** | **string** | Domestic first address (EN) | 
**DomesticFirstAddressKo** | **string** | Domestic first address (KO) | 
**DomesticSecondAddressEn** | **string** | Domestic second address (EN) | 
**DomesticSecondAddressKo** | **string** | Domestic second address (KO) | 
**ExpiredDate** | **string** | Expired date | 
**Id** | **string** | ID | 
**ModifiedAt** | **string** | Modified Date | 
**ModifiedBy** | **string** | Modified By_ | 
**Name** | **string** | Public Domain Name | 
**OverseasAddress** | Pointer to **string** | Overseas Address | [optional] [default to ""]
**OverseasFirstAddress** | **string** | Overseas address - Address | 
**OverseasSecondAddress** | **string** | Overseas address - City | 
**OverseasThirdAddress** | **string** | Overseas address - State &amp; Country | 
**PostalCode** | **string** | Postal code | 
**RegisterEmail** | **string** | Register email | 
**RegisterNameEn** | **string** | Register Name Overseas | 
**RegisterNameKo** | **string** | Register Name Domestic | 
**RegisterTelno** | **string** | Register telephone number | 
**StartDate** | **string** | Start date | 
**Status** | **string** | The status | 

## Methods

### NewPublicDomainDetail

`func NewPublicDomainDetail(addressType string, autoExtension bool, createdAt string, createdBy string, domesticAddressEn string, domesticAddressKo string, domesticFirstAddressEn string, domesticFirstAddressKo string, domesticSecondAddressEn string, domesticSecondAddressKo string, expiredDate string, id string, modifiedAt string, modifiedBy string, name string, overseasFirstAddress string, overseasSecondAddress string, overseasThirdAddress string, postalCode string, registerEmail string, registerNameEn string, registerNameKo string, registerTelno string, startDate string, status string, ) *PublicDomainDetail`

NewPublicDomainDetail instantiates a new PublicDomainDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDomainDetailWithDefaults

`func NewPublicDomainDetailWithDefaults() *PublicDomainDetail`

NewPublicDomainDetailWithDefaults instantiates a new PublicDomainDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressType

`func (o *PublicDomainDetail) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *PublicDomainDetail) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *PublicDomainDetail) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.


### GetAutoExtension

`func (o *PublicDomainDetail) GetAutoExtension() bool`

GetAutoExtension returns the AutoExtension field if non-nil, zero value otherwise.

### GetAutoExtensionOk

`func (o *PublicDomainDetail) GetAutoExtensionOk() (*bool, bool)`

GetAutoExtensionOk returns a tuple with the AutoExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoExtension

`func (o *PublicDomainDetail) SetAutoExtension(v bool)`

SetAutoExtension sets AutoExtension field to given value.


### GetCreatedAt

`func (o *PublicDomainDetail) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicDomainDetail) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicDomainDetail) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PublicDomainDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PublicDomainDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PublicDomainDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *PublicDomainDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublicDomainDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublicDomainDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublicDomainDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PublicDomainDetail) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PublicDomainDetail) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomesticAddressEn

`func (o *PublicDomainDetail) GetDomesticAddressEn() string`

GetDomesticAddressEn returns the DomesticAddressEn field if non-nil, zero value otherwise.

### GetDomesticAddressEnOk

`func (o *PublicDomainDetail) GetDomesticAddressEnOk() (*string, bool)`

GetDomesticAddressEnOk returns a tuple with the DomesticAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticAddressEn

`func (o *PublicDomainDetail) SetDomesticAddressEn(v string)`

SetDomesticAddressEn sets DomesticAddressEn field to given value.


### GetDomesticAddressKo

`func (o *PublicDomainDetail) GetDomesticAddressKo() string`

GetDomesticAddressKo returns the DomesticAddressKo field if non-nil, zero value otherwise.

### GetDomesticAddressKoOk

`func (o *PublicDomainDetail) GetDomesticAddressKoOk() (*string, bool)`

GetDomesticAddressKoOk returns a tuple with the DomesticAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticAddressKo

`func (o *PublicDomainDetail) SetDomesticAddressKo(v string)`

SetDomesticAddressKo sets DomesticAddressKo field to given value.


### GetDomesticFirstAddressEn

`func (o *PublicDomainDetail) GetDomesticFirstAddressEn() string`

GetDomesticFirstAddressEn returns the DomesticFirstAddressEn field if non-nil, zero value otherwise.

### GetDomesticFirstAddressEnOk

`func (o *PublicDomainDetail) GetDomesticFirstAddressEnOk() (*string, bool)`

GetDomesticFirstAddressEnOk returns a tuple with the DomesticFirstAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressEn

`func (o *PublicDomainDetail) SetDomesticFirstAddressEn(v string)`

SetDomesticFirstAddressEn sets DomesticFirstAddressEn field to given value.


### GetDomesticFirstAddressKo

`func (o *PublicDomainDetail) GetDomesticFirstAddressKo() string`

GetDomesticFirstAddressKo returns the DomesticFirstAddressKo field if non-nil, zero value otherwise.

### GetDomesticFirstAddressKoOk

`func (o *PublicDomainDetail) GetDomesticFirstAddressKoOk() (*string, bool)`

GetDomesticFirstAddressKoOk returns a tuple with the DomesticFirstAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticFirstAddressKo

`func (o *PublicDomainDetail) SetDomesticFirstAddressKo(v string)`

SetDomesticFirstAddressKo sets DomesticFirstAddressKo field to given value.


### GetDomesticSecondAddressEn

`func (o *PublicDomainDetail) GetDomesticSecondAddressEn() string`

GetDomesticSecondAddressEn returns the DomesticSecondAddressEn field if non-nil, zero value otherwise.

### GetDomesticSecondAddressEnOk

`func (o *PublicDomainDetail) GetDomesticSecondAddressEnOk() (*string, bool)`

GetDomesticSecondAddressEnOk returns a tuple with the DomesticSecondAddressEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressEn

`func (o *PublicDomainDetail) SetDomesticSecondAddressEn(v string)`

SetDomesticSecondAddressEn sets DomesticSecondAddressEn field to given value.


### GetDomesticSecondAddressKo

`func (o *PublicDomainDetail) GetDomesticSecondAddressKo() string`

GetDomesticSecondAddressKo returns the DomesticSecondAddressKo field if non-nil, zero value otherwise.

### GetDomesticSecondAddressKoOk

`func (o *PublicDomainDetail) GetDomesticSecondAddressKoOk() (*string, bool)`

GetDomesticSecondAddressKoOk returns a tuple with the DomesticSecondAddressKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomesticSecondAddressKo

`func (o *PublicDomainDetail) SetDomesticSecondAddressKo(v string)`

SetDomesticSecondAddressKo sets DomesticSecondAddressKo field to given value.


### GetExpiredDate

`func (o *PublicDomainDetail) GetExpiredDate() string`

GetExpiredDate returns the ExpiredDate field if non-nil, zero value otherwise.

### GetExpiredDateOk

`func (o *PublicDomainDetail) GetExpiredDateOk() (*string, bool)`

GetExpiredDateOk returns a tuple with the ExpiredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredDate

`func (o *PublicDomainDetail) SetExpiredDate(v string)`

SetExpiredDate sets ExpiredDate field to given value.


### GetId

`func (o *PublicDomainDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublicDomainDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublicDomainDetail) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *PublicDomainDetail) GetModifiedAt() string`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PublicDomainDetail) GetModifiedAtOk() (*string, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PublicDomainDetail) SetModifiedAt(v string)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PublicDomainDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PublicDomainDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PublicDomainDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *PublicDomainDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicDomainDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicDomainDetail) SetName(v string)`

SetName sets Name field to given value.


### GetOverseasAddress

`func (o *PublicDomainDetail) GetOverseasAddress() string`

GetOverseasAddress returns the OverseasAddress field if non-nil, zero value otherwise.

### GetOverseasAddressOk

`func (o *PublicDomainDetail) GetOverseasAddressOk() (*string, bool)`

GetOverseasAddressOk returns a tuple with the OverseasAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasAddress

`func (o *PublicDomainDetail) SetOverseasAddress(v string)`

SetOverseasAddress sets OverseasAddress field to given value.

### HasOverseasAddress

`func (o *PublicDomainDetail) HasOverseasAddress() bool`

HasOverseasAddress returns a boolean if a field has been set.

### GetOverseasFirstAddress

`func (o *PublicDomainDetail) GetOverseasFirstAddress() string`

GetOverseasFirstAddress returns the OverseasFirstAddress field if non-nil, zero value otherwise.

### GetOverseasFirstAddressOk

`func (o *PublicDomainDetail) GetOverseasFirstAddressOk() (*string, bool)`

GetOverseasFirstAddressOk returns a tuple with the OverseasFirstAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasFirstAddress

`func (o *PublicDomainDetail) SetOverseasFirstAddress(v string)`

SetOverseasFirstAddress sets OverseasFirstAddress field to given value.


### GetOverseasSecondAddress

`func (o *PublicDomainDetail) GetOverseasSecondAddress() string`

GetOverseasSecondAddress returns the OverseasSecondAddress field if non-nil, zero value otherwise.

### GetOverseasSecondAddressOk

`func (o *PublicDomainDetail) GetOverseasSecondAddressOk() (*string, bool)`

GetOverseasSecondAddressOk returns a tuple with the OverseasSecondAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasSecondAddress

`func (o *PublicDomainDetail) SetOverseasSecondAddress(v string)`

SetOverseasSecondAddress sets OverseasSecondAddress field to given value.


### GetOverseasThirdAddress

`func (o *PublicDomainDetail) GetOverseasThirdAddress() string`

GetOverseasThirdAddress returns the OverseasThirdAddress field if non-nil, zero value otherwise.

### GetOverseasThirdAddressOk

`func (o *PublicDomainDetail) GetOverseasThirdAddressOk() (*string, bool)`

GetOverseasThirdAddressOk returns a tuple with the OverseasThirdAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverseasThirdAddress

`func (o *PublicDomainDetail) SetOverseasThirdAddress(v string)`

SetOverseasThirdAddress sets OverseasThirdAddress field to given value.


### GetPostalCode

`func (o *PublicDomainDetail) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PublicDomainDetail) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PublicDomainDetail) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetRegisterEmail

`func (o *PublicDomainDetail) GetRegisterEmail() string`

GetRegisterEmail returns the RegisterEmail field if non-nil, zero value otherwise.

### GetRegisterEmailOk

`func (o *PublicDomainDetail) GetRegisterEmailOk() (*string, bool)`

GetRegisterEmailOk returns a tuple with the RegisterEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterEmail

`func (o *PublicDomainDetail) SetRegisterEmail(v string)`

SetRegisterEmail sets RegisterEmail field to given value.


### GetRegisterNameEn

`func (o *PublicDomainDetail) GetRegisterNameEn() string`

GetRegisterNameEn returns the RegisterNameEn field if non-nil, zero value otherwise.

### GetRegisterNameEnOk

`func (o *PublicDomainDetail) GetRegisterNameEnOk() (*string, bool)`

GetRegisterNameEnOk returns a tuple with the RegisterNameEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNameEn

`func (o *PublicDomainDetail) SetRegisterNameEn(v string)`

SetRegisterNameEn sets RegisterNameEn field to given value.


### GetRegisterNameKo

`func (o *PublicDomainDetail) GetRegisterNameKo() string`

GetRegisterNameKo returns the RegisterNameKo field if non-nil, zero value otherwise.

### GetRegisterNameKoOk

`func (o *PublicDomainDetail) GetRegisterNameKoOk() (*string, bool)`

GetRegisterNameKoOk returns a tuple with the RegisterNameKo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterNameKo

`func (o *PublicDomainDetail) SetRegisterNameKo(v string)`

SetRegisterNameKo sets RegisterNameKo field to given value.


### GetRegisterTelno

`func (o *PublicDomainDetail) GetRegisterTelno() string`

GetRegisterTelno returns the RegisterTelno field if non-nil, zero value otherwise.

### GetRegisterTelnoOk

`func (o *PublicDomainDetail) GetRegisterTelnoOk() (*string, bool)`

GetRegisterTelnoOk returns a tuple with the RegisterTelno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisterTelno

`func (o *PublicDomainDetail) SetRegisterTelno(v string)`

SetRegisterTelno sets RegisterTelno field to given value.


### GetStartDate

`func (o *PublicDomainDetail) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PublicDomainDetail) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PublicDomainDetail) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetStatus

`func (o *PublicDomainDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PublicDomainDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PublicDomainDetail) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


