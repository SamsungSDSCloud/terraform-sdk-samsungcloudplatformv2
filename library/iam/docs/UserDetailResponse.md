# UserDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessControl** | **bool** | 접근 제어 설정 유무 | 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AccountName** | Pointer to **NullableString** |  | [optional] 
**AllowedIpAddresses** | **string** | 접근 허용 IP | 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DstOffset** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** | 사용자 Email | 
**EmailAuthenticated** | **bool** | email 인증 여부 | 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | ID | 
**LastLoginAt** | Pointer to **NullableTime** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**LastPasswordUpdateAt** | Pointer to **NullableTime** |  | [optional] 
**LoginId** | **string** | Login ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**NationId** | **string** | 국가 ID | 
**PasswordReuseCount** | **int32** | 비밀 번호 재사용 가능 횟수 | 
**PhoneAuthenticated** | **bool** | 휴대폰 인증 여부 | 
**PhoneNumber** | **string** | 사용자 휴대 전화 번호 | 
**Timezone** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** | Type | 
**TzId** | Pointer to **NullableString** |  | [optional] 
**UserName** | Pointer to **NullableString** |  | [optional] 
**UtcOffset** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUserDetailResponse

`func NewUserDetailResponse(accessControl bool, allowedIpAddresses string, createdAt time.Time, createdBy string, email string, emailAuthenticated bool, id string, loginId string, modifiedAt time.Time, modifiedBy string, nationId string, passwordReuseCount int32, phoneAuthenticated bool, phoneNumber string, type_ string, ) *UserDetailResponse`

NewUserDetailResponse instantiates a new UserDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailResponseWithDefaults

`func NewUserDetailResponseWithDefaults() *UserDetailResponse`

NewUserDetailResponseWithDefaults instantiates a new UserDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessControl

`func (o *UserDetailResponse) GetAccessControl() bool`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *UserDetailResponse) GetAccessControlOk() (*bool, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *UserDetailResponse) SetAccessControl(v bool)`

SetAccessControl sets AccessControl field to given value.


### GetAccountId

`func (o *UserDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserDetailResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *UserDetailResponse) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *UserDetailResponse) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAccountName

`func (o *UserDetailResponse) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *UserDetailResponse) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *UserDetailResponse) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *UserDetailResponse) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### SetAccountNameNil

`func (o *UserDetailResponse) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *UserDetailResponse) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetAllowedIpAddresses

`func (o *UserDetailResponse) GetAllowedIpAddresses() string`

GetAllowedIpAddresses returns the AllowedIpAddresses field if non-nil, zero value otherwise.

### GetAllowedIpAddressesOk

`func (o *UserDetailResponse) GetAllowedIpAddressesOk() (*string, bool)`

GetAllowedIpAddressesOk returns a tuple with the AllowedIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIpAddresses

`func (o *UserDetailResponse) SetAllowedIpAddresses(v string)`

SetAllowedIpAddresses sets AllowedIpAddresses field to given value.


### GetCompanyName

`func (o *UserDetailResponse) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UserDetailResponse) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UserDetailResponse) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UserDetailResponse) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *UserDetailResponse) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *UserDetailResponse) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCreatedAt

`func (o *UserDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *UserDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UserDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UserDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *UserDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserDetailResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UserDetailResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UserDetailResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDstOffset

`func (o *UserDetailResponse) GetDstOffset() string`

GetDstOffset returns the DstOffset field if non-nil, zero value otherwise.

### GetDstOffsetOk

`func (o *UserDetailResponse) GetDstOffsetOk() (*string, bool)`

GetDstOffsetOk returns a tuple with the DstOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstOffset

`func (o *UserDetailResponse) SetDstOffset(v string)`

SetDstOffset sets DstOffset field to given value.

### HasDstOffset

`func (o *UserDetailResponse) HasDstOffset() bool`

HasDstOffset returns a boolean if a field has been set.

### SetDstOffsetNil

`func (o *UserDetailResponse) SetDstOffsetNil(b bool)`

 SetDstOffsetNil sets the value for DstOffset to be an explicit nil

### UnsetDstOffset
`func (o *UserDetailResponse) UnsetDstOffset()`

UnsetDstOffset ensures that no value is present for DstOffset, not even an explicit nil
### GetEmail

`func (o *UserDetailResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDetailResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDetailResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetEmailAuthenticated

`func (o *UserDetailResponse) GetEmailAuthenticated() bool`

GetEmailAuthenticated returns the EmailAuthenticated field if non-nil, zero value otherwise.

### GetEmailAuthenticatedOk

`func (o *UserDetailResponse) GetEmailAuthenticatedOk() (*bool, bool)`

GetEmailAuthenticatedOk returns a tuple with the EmailAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAuthenticated

`func (o *UserDetailResponse) SetEmailAuthenticated(v bool)`

SetEmailAuthenticated sets EmailAuthenticated field to given value.


### GetFirstName

`func (o *UserDetailResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserDetailResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserDetailResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserDetailResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserDetailResponse) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserDetailResponse) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetId

`func (o *UserDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLastLoginAt

`func (o *UserDetailResponse) GetLastLoginAt() time.Time`

GetLastLoginAt returns the LastLoginAt field if non-nil, zero value otherwise.

### GetLastLoginAtOk

`func (o *UserDetailResponse) GetLastLoginAtOk() (*time.Time, bool)`

GetLastLoginAtOk returns a tuple with the LastLoginAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginAt

`func (o *UserDetailResponse) SetLastLoginAt(v time.Time)`

SetLastLoginAt sets LastLoginAt field to given value.

### HasLastLoginAt

`func (o *UserDetailResponse) HasLastLoginAt() bool`

HasLastLoginAt returns a boolean if a field has been set.

### SetLastLoginAtNil

`func (o *UserDetailResponse) SetLastLoginAtNil(b bool)`

 SetLastLoginAtNil sets the value for LastLoginAt to be an explicit nil

### UnsetLastLoginAt
`func (o *UserDetailResponse) UnsetLastLoginAt()`

UnsetLastLoginAt ensures that no value is present for LastLoginAt, not even an explicit nil
### GetLastName

`func (o *UserDetailResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserDetailResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserDetailResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserDetailResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserDetailResponse) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserDetailResponse) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetLastPasswordUpdateAt

`func (o *UserDetailResponse) GetLastPasswordUpdateAt() time.Time`

GetLastPasswordUpdateAt returns the LastPasswordUpdateAt field if non-nil, zero value otherwise.

### GetLastPasswordUpdateAtOk

`func (o *UserDetailResponse) GetLastPasswordUpdateAtOk() (*time.Time, bool)`

GetLastPasswordUpdateAtOk returns a tuple with the LastPasswordUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPasswordUpdateAt

`func (o *UserDetailResponse) SetLastPasswordUpdateAt(v time.Time)`

SetLastPasswordUpdateAt sets LastPasswordUpdateAt field to given value.

### HasLastPasswordUpdateAt

`func (o *UserDetailResponse) HasLastPasswordUpdateAt() bool`

HasLastPasswordUpdateAt returns a boolean if a field has been set.

### SetLastPasswordUpdateAtNil

`func (o *UserDetailResponse) SetLastPasswordUpdateAtNil(b bool)`

 SetLastPasswordUpdateAtNil sets the value for LastPasswordUpdateAt to be an explicit nil

### UnsetLastPasswordUpdateAt
`func (o *UserDetailResponse) UnsetLastPasswordUpdateAt()`

UnsetLastPasswordUpdateAt ensures that no value is present for LastPasswordUpdateAt, not even an explicit nil
### GetLoginId

`func (o *UserDetailResponse) GetLoginId() string`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *UserDetailResponse) GetLoginIdOk() (*string, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *UserDetailResponse) SetLoginId(v string)`

SetLoginId sets LoginId field to given value.


### GetModifiedAt

`func (o *UserDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *UserDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *UserDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *UserDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *UserDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *UserDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetNationId

`func (o *UserDetailResponse) GetNationId() string`

GetNationId returns the NationId field if non-nil, zero value otherwise.

### GetNationIdOk

`func (o *UserDetailResponse) GetNationIdOk() (*string, bool)`

GetNationIdOk returns a tuple with the NationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationId

`func (o *UserDetailResponse) SetNationId(v string)`

SetNationId sets NationId field to given value.


### GetPasswordReuseCount

`func (o *UserDetailResponse) GetPasswordReuseCount() int32`

GetPasswordReuseCount returns the PasswordReuseCount field if non-nil, zero value otherwise.

### GetPasswordReuseCountOk

`func (o *UserDetailResponse) GetPasswordReuseCountOk() (*int32, bool)`

GetPasswordReuseCountOk returns a tuple with the PasswordReuseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordReuseCount

`func (o *UserDetailResponse) SetPasswordReuseCount(v int32)`

SetPasswordReuseCount sets PasswordReuseCount field to given value.


### GetPhoneAuthenticated

`func (o *UserDetailResponse) GetPhoneAuthenticated() bool`

GetPhoneAuthenticated returns the PhoneAuthenticated field if non-nil, zero value otherwise.

### GetPhoneAuthenticatedOk

`func (o *UserDetailResponse) GetPhoneAuthenticatedOk() (*bool, bool)`

GetPhoneAuthenticatedOk returns a tuple with the PhoneAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneAuthenticated

`func (o *UserDetailResponse) SetPhoneAuthenticated(v bool)`

SetPhoneAuthenticated sets PhoneAuthenticated field to given value.


### GetPhoneNumber

`func (o *UserDetailResponse) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserDetailResponse) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserDetailResponse) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetTimezone

`func (o *UserDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserDetailResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *UserDetailResponse) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *UserDetailResponse) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
### GetType

`func (o *UserDetailResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserDetailResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserDetailResponse) SetType(v string)`

SetType sets Type field to given value.


### GetTzId

`func (o *UserDetailResponse) GetTzId() string`

GetTzId returns the TzId field if non-nil, zero value otherwise.

### GetTzIdOk

`func (o *UserDetailResponse) GetTzIdOk() (*string, bool)`

GetTzIdOk returns a tuple with the TzId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTzId

`func (o *UserDetailResponse) SetTzId(v string)`

SetTzId sets TzId field to given value.

### HasTzId

`func (o *UserDetailResponse) HasTzId() bool`

HasTzId returns a boolean if a field has been set.

### SetTzIdNil

`func (o *UserDetailResponse) SetTzIdNil(b bool)`

 SetTzIdNil sets the value for TzId to be an explicit nil

### UnsetTzId
`func (o *UserDetailResponse) UnsetTzId()`

UnsetTzId ensures that no value is present for TzId, not even an explicit nil
### GetUserName

`func (o *UserDetailResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *UserDetailResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *UserDetailResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *UserDetailResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### SetUserNameNil

`func (o *UserDetailResponse) SetUserNameNil(b bool)`

 SetUserNameNil sets the value for UserName to be an explicit nil

### UnsetUserName
`func (o *UserDetailResponse) UnsetUserName()`

UnsetUserName ensures that no value is present for UserName, not even an explicit nil
### GetUtcOffset

`func (o *UserDetailResponse) GetUtcOffset() string`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *UserDetailResponse) GetUtcOffsetOk() (*string, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *UserDetailResponse) SetUtcOffset(v string)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *UserDetailResponse) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.

### SetUtcOffsetNil

`func (o *UserDetailResponse) SetUtcOffsetNil(b bool)`

 SetUtcOffsetNil sets the value for UtcOffset to be an explicit nil

### UnsetUtcOffset
`func (o *UserDetailResponse) UnsetUtcOffset()`

UnsetUtcOffset ensures that no value is present for UtcOffset, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


