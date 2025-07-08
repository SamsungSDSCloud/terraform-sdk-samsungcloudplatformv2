# SamlProviderDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AcsUrl** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**FederationType** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | ID | 
**IdpCertEnd** | Pointer to **NullableTime** |  | [optional] 
**IdpFileData** | Pointer to **NullableString** |  | [optional] 
**IdpFileName** | Pointer to **NullableString** |  | [optional] 
**IdpFileSize** | Pointer to **NullableInt32** |  | [optional] 
**IdpIssuerUrl** | Pointer to **NullableString** |  | [optional] 
**IdpSsoUrl** | Pointer to **NullableString** |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ModifierName** | Pointer to **NullableString** |  | [optional] 
**SamlProviderName** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]map[string]string** |  | [optional] 

## Methods

### NewSamlProviderDetailResponse

`func NewSamlProviderDetailResponse(accountId string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, ) *SamlProviderDetailResponse`

NewSamlProviderDetailResponse instantiates a new SamlProviderDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProviderDetailResponseWithDefaults

`func NewSamlProviderDetailResponseWithDefaults() *SamlProviderDetailResponse`

NewSamlProviderDetailResponseWithDefaults instantiates a new SamlProviderDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SamlProviderDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SamlProviderDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SamlProviderDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAcsUrl

`func (o *SamlProviderDetailResponse) GetAcsUrl() string`

GetAcsUrl returns the AcsUrl field if non-nil, zero value otherwise.

### GetAcsUrlOk

`func (o *SamlProviderDetailResponse) GetAcsUrlOk() (*string, bool)`

GetAcsUrlOk returns a tuple with the AcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcsUrl

`func (o *SamlProviderDetailResponse) SetAcsUrl(v string)`

SetAcsUrl sets AcsUrl field to given value.

### HasAcsUrl

`func (o *SamlProviderDetailResponse) HasAcsUrl() bool`

HasAcsUrl returns a boolean if a field has been set.

### SetAcsUrlNil

`func (o *SamlProviderDetailResponse) SetAcsUrlNil(b bool)`

 SetAcsUrlNil sets the value for AcsUrl to be an explicit nil

### UnsetAcsUrl
`func (o *SamlProviderDetailResponse) UnsetAcsUrl()`

UnsetAcsUrl ensures that no value is present for AcsUrl, not even an explicit nil
### GetCreatedAt

`func (o *SamlProviderDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SamlProviderDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SamlProviderDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SamlProviderDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SamlProviderDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SamlProviderDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorName

`func (o *SamlProviderDetailResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *SamlProviderDetailResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *SamlProviderDetailResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *SamlProviderDetailResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *SamlProviderDetailResponse) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *SamlProviderDetailResponse) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetDescription

`func (o *SamlProviderDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SamlProviderDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SamlProviderDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SamlProviderDetailResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SamlProviderDetailResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SamlProviderDetailResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFederationType

`func (o *SamlProviderDetailResponse) GetFederationType() string`

GetFederationType returns the FederationType field if non-nil, zero value otherwise.

### GetFederationTypeOk

`func (o *SamlProviderDetailResponse) GetFederationTypeOk() (*string, bool)`

GetFederationTypeOk returns a tuple with the FederationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationType

`func (o *SamlProviderDetailResponse) SetFederationType(v string)`

SetFederationType sets FederationType field to given value.

### HasFederationType

`func (o *SamlProviderDetailResponse) HasFederationType() bool`

HasFederationType returns a boolean if a field has been set.

### SetFederationTypeNil

`func (o *SamlProviderDetailResponse) SetFederationTypeNil(b bool)`

 SetFederationTypeNil sets the value for FederationType to be an explicit nil

### UnsetFederationType
`func (o *SamlProviderDetailResponse) UnsetFederationType()`

UnsetFederationType ensures that no value is present for FederationType, not even an explicit nil
### GetId

`func (o *SamlProviderDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlProviderDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SamlProviderDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetIdpCertEnd

`func (o *SamlProviderDetailResponse) GetIdpCertEnd() time.Time`

GetIdpCertEnd returns the IdpCertEnd field if non-nil, zero value otherwise.

### GetIdpCertEndOk

`func (o *SamlProviderDetailResponse) GetIdpCertEndOk() (*time.Time, bool)`

GetIdpCertEndOk returns a tuple with the IdpCertEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpCertEnd

`func (o *SamlProviderDetailResponse) SetIdpCertEnd(v time.Time)`

SetIdpCertEnd sets IdpCertEnd field to given value.

### HasIdpCertEnd

`func (o *SamlProviderDetailResponse) HasIdpCertEnd() bool`

HasIdpCertEnd returns a boolean if a field has been set.

### SetIdpCertEndNil

`func (o *SamlProviderDetailResponse) SetIdpCertEndNil(b bool)`

 SetIdpCertEndNil sets the value for IdpCertEnd to be an explicit nil

### UnsetIdpCertEnd
`func (o *SamlProviderDetailResponse) UnsetIdpCertEnd()`

UnsetIdpCertEnd ensures that no value is present for IdpCertEnd, not even an explicit nil
### GetIdpFileData

`func (o *SamlProviderDetailResponse) GetIdpFileData() string`

GetIdpFileData returns the IdpFileData field if non-nil, zero value otherwise.

### GetIdpFileDataOk

`func (o *SamlProviderDetailResponse) GetIdpFileDataOk() (*string, bool)`

GetIdpFileDataOk returns a tuple with the IdpFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpFileData

`func (o *SamlProviderDetailResponse) SetIdpFileData(v string)`

SetIdpFileData sets IdpFileData field to given value.

### HasIdpFileData

`func (o *SamlProviderDetailResponse) HasIdpFileData() bool`

HasIdpFileData returns a boolean if a field has been set.

### SetIdpFileDataNil

`func (o *SamlProviderDetailResponse) SetIdpFileDataNil(b bool)`

 SetIdpFileDataNil sets the value for IdpFileData to be an explicit nil

### UnsetIdpFileData
`func (o *SamlProviderDetailResponse) UnsetIdpFileData()`

UnsetIdpFileData ensures that no value is present for IdpFileData, not even an explicit nil
### GetIdpFileName

`func (o *SamlProviderDetailResponse) GetIdpFileName() string`

GetIdpFileName returns the IdpFileName field if non-nil, zero value otherwise.

### GetIdpFileNameOk

`func (o *SamlProviderDetailResponse) GetIdpFileNameOk() (*string, bool)`

GetIdpFileNameOk returns a tuple with the IdpFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpFileName

`func (o *SamlProviderDetailResponse) SetIdpFileName(v string)`

SetIdpFileName sets IdpFileName field to given value.

### HasIdpFileName

`func (o *SamlProviderDetailResponse) HasIdpFileName() bool`

HasIdpFileName returns a boolean if a field has been set.

### SetIdpFileNameNil

`func (o *SamlProviderDetailResponse) SetIdpFileNameNil(b bool)`

 SetIdpFileNameNil sets the value for IdpFileName to be an explicit nil

### UnsetIdpFileName
`func (o *SamlProviderDetailResponse) UnsetIdpFileName()`

UnsetIdpFileName ensures that no value is present for IdpFileName, not even an explicit nil
### GetIdpFileSize

`func (o *SamlProviderDetailResponse) GetIdpFileSize() int32`

GetIdpFileSize returns the IdpFileSize field if non-nil, zero value otherwise.

### GetIdpFileSizeOk

`func (o *SamlProviderDetailResponse) GetIdpFileSizeOk() (*int32, bool)`

GetIdpFileSizeOk returns a tuple with the IdpFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpFileSize

`func (o *SamlProviderDetailResponse) SetIdpFileSize(v int32)`

SetIdpFileSize sets IdpFileSize field to given value.

### HasIdpFileSize

`func (o *SamlProviderDetailResponse) HasIdpFileSize() bool`

HasIdpFileSize returns a boolean if a field has been set.

### SetIdpFileSizeNil

`func (o *SamlProviderDetailResponse) SetIdpFileSizeNil(b bool)`

 SetIdpFileSizeNil sets the value for IdpFileSize to be an explicit nil

### UnsetIdpFileSize
`func (o *SamlProviderDetailResponse) UnsetIdpFileSize()`

UnsetIdpFileSize ensures that no value is present for IdpFileSize, not even an explicit nil
### GetIdpIssuerUrl

`func (o *SamlProviderDetailResponse) GetIdpIssuerUrl() string`

GetIdpIssuerUrl returns the IdpIssuerUrl field if non-nil, zero value otherwise.

### GetIdpIssuerUrlOk

`func (o *SamlProviderDetailResponse) GetIdpIssuerUrlOk() (*string, bool)`

GetIdpIssuerUrlOk returns a tuple with the IdpIssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpIssuerUrl

`func (o *SamlProviderDetailResponse) SetIdpIssuerUrl(v string)`

SetIdpIssuerUrl sets IdpIssuerUrl field to given value.

### HasIdpIssuerUrl

`func (o *SamlProviderDetailResponse) HasIdpIssuerUrl() bool`

HasIdpIssuerUrl returns a boolean if a field has been set.

### SetIdpIssuerUrlNil

`func (o *SamlProviderDetailResponse) SetIdpIssuerUrlNil(b bool)`

 SetIdpIssuerUrlNil sets the value for IdpIssuerUrl to be an explicit nil

### UnsetIdpIssuerUrl
`func (o *SamlProviderDetailResponse) UnsetIdpIssuerUrl()`

UnsetIdpIssuerUrl ensures that no value is present for IdpIssuerUrl, not even an explicit nil
### GetIdpSsoUrl

`func (o *SamlProviderDetailResponse) GetIdpSsoUrl() string`

GetIdpSsoUrl returns the IdpSsoUrl field if non-nil, zero value otherwise.

### GetIdpSsoUrlOk

`func (o *SamlProviderDetailResponse) GetIdpSsoUrlOk() (*string, bool)`

GetIdpSsoUrlOk returns a tuple with the IdpSsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpSsoUrl

`func (o *SamlProviderDetailResponse) SetIdpSsoUrl(v string)`

SetIdpSsoUrl sets IdpSsoUrl field to given value.

### HasIdpSsoUrl

`func (o *SamlProviderDetailResponse) HasIdpSsoUrl() bool`

HasIdpSsoUrl returns a boolean if a field has been set.

### SetIdpSsoUrlNil

`func (o *SamlProviderDetailResponse) SetIdpSsoUrlNil(b bool)`

 SetIdpSsoUrlNil sets the value for IdpSsoUrl to be an explicit nil

### UnsetIdpSsoUrl
`func (o *SamlProviderDetailResponse) UnsetIdpSsoUrl()`

UnsetIdpSsoUrl ensures that no value is present for IdpSsoUrl, not even an explicit nil
### GetModifiedAt

`func (o *SamlProviderDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SamlProviderDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SamlProviderDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *SamlProviderDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SamlProviderDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SamlProviderDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierName

`func (o *SamlProviderDetailResponse) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *SamlProviderDetailResponse) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *SamlProviderDetailResponse) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.

### HasModifierName

`func (o *SamlProviderDetailResponse) HasModifierName() bool`

HasModifierName returns a boolean if a field has been set.

### SetModifierNameNil

`func (o *SamlProviderDetailResponse) SetModifierNameNil(b bool)`

 SetModifierNameNil sets the value for ModifierName to be an explicit nil

### UnsetModifierName
`func (o *SamlProviderDetailResponse) UnsetModifierName()`

UnsetModifierName ensures that no value is present for ModifierName, not even an explicit nil
### GetSamlProviderName

`func (o *SamlProviderDetailResponse) GetSamlProviderName() string`

GetSamlProviderName returns the SamlProviderName field if non-nil, zero value otherwise.

### GetSamlProviderNameOk

`func (o *SamlProviderDetailResponse) GetSamlProviderNameOk() (*string, bool)`

GetSamlProviderNameOk returns a tuple with the SamlProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlProviderName

`func (o *SamlProviderDetailResponse) SetSamlProviderName(v string)`

SetSamlProviderName sets SamlProviderName field to given value.

### HasSamlProviderName

`func (o *SamlProviderDetailResponse) HasSamlProviderName() bool`

HasSamlProviderName returns a boolean if a field has been set.

### SetSamlProviderNameNil

`func (o *SamlProviderDetailResponse) SetSamlProviderNameNil(b bool)`

 SetSamlProviderNameNil sets the value for SamlProviderName to be an explicit nil

### UnsetSamlProviderName
`func (o *SamlProviderDetailResponse) UnsetSamlProviderName()`

UnsetSamlProviderName ensures that no value is present for SamlProviderName, not even an explicit nil
### GetTags

`func (o *SamlProviderDetailResponse) GetTags() []map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SamlProviderDetailResponse) GetTagsOk() (*[]map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SamlProviderDetailResponse) SetTags(v []map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SamlProviderDetailResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *SamlProviderDetailResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *SamlProviderDetailResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


