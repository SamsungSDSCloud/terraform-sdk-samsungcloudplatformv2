# SamlProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**CreatorName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**FederationType** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**ModifierName** | Pointer to **NullableString** |  | [optional] 
**SamlProviderName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSamlProviderResponse

`func NewSamlProviderResponse(accountId string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, ) *SamlProviderResponse`

NewSamlProviderResponse instantiates a new SamlProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlProviderResponseWithDefaults

`func NewSamlProviderResponseWithDefaults() *SamlProviderResponse`

NewSamlProviderResponseWithDefaults instantiates a new SamlProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SamlProviderResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SamlProviderResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SamlProviderResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *SamlProviderResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SamlProviderResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SamlProviderResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SamlProviderResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SamlProviderResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SamlProviderResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatorName

`func (o *SamlProviderResponse) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *SamlProviderResponse) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *SamlProviderResponse) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *SamlProviderResponse) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### SetCreatorNameNil

`func (o *SamlProviderResponse) SetCreatorNameNil(b bool)`

 SetCreatorNameNil sets the value for CreatorName to be an explicit nil

### UnsetCreatorName
`func (o *SamlProviderResponse) UnsetCreatorName()`

UnsetCreatorName ensures that no value is present for CreatorName, not even an explicit nil
### GetDescription

`func (o *SamlProviderResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SamlProviderResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SamlProviderResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SamlProviderResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SamlProviderResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SamlProviderResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFederationType

`func (o *SamlProviderResponse) GetFederationType() string`

GetFederationType returns the FederationType field if non-nil, zero value otherwise.

### GetFederationTypeOk

`func (o *SamlProviderResponse) GetFederationTypeOk() (*string, bool)`

GetFederationTypeOk returns a tuple with the FederationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationType

`func (o *SamlProviderResponse) SetFederationType(v string)`

SetFederationType sets FederationType field to given value.

### HasFederationType

`func (o *SamlProviderResponse) HasFederationType() bool`

HasFederationType returns a boolean if a field has been set.

### SetFederationTypeNil

`func (o *SamlProviderResponse) SetFederationTypeNil(b bool)`

 SetFederationTypeNil sets the value for FederationType to be an explicit nil

### UnsetFederationType
`func (o *SamlProviderResponse) UnsetFederationType()`

UnsetFederationType ensures that no value is present for FederationType, not even an explicit nil
### GetId

`func (o *SamlProviderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlProviderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SamlProviderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *SamlProviderResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SamlProviderResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SamlProviderResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *SamlProviderResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SamlProviderResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SamlProviderResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetModifierName

`func (o *SamlProviderResponse) GetModifierName() string`

GetModifierName returns the ModifierName field if non-nil, zero value otherwise.

### GetModifierNameOk

`func (o *SamlProviderResponse) GetModifierNameOk() (*string, bool)`

GetModifierNameOk returns a tuple with the ModifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifierName

`func (o *SamlProviderResponse) SetModifierName(v string)`

SetModifierName sets ModifierName field to given value.

### HasModifierName

`func (o *SamlProviderResponse) HasModifierName() bool`

HasModifierName returns a boolean if a field has been set.

### SetModifierNameNil

`func (o *SamlProviderResponse) SetModifierNameNil(b bool)`

 SetModifierNameNil sets the value for ModifierName to be an explicit nil

### UnsetModifierName
`func (o *SamlProviderResponse) UnsetModifierName()`

UnsetModifierName ensures that no value is present for ModifierName, not even an explicit nil
### GetSamlProviderName

`func (o *SamlProviderResponse) GetSamlProviderName() string`

GetSamlProviderName returns the SamlProviderName field if non-nil, zero value otherwise.

### GetSamlProviderNameOk

`func (o *SamlProviderResponse) GetSamlProviderNameOk() (*string, bool)`

GetSamlProviderNameOk returns a tuple with the SamlProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlProviderName

`func (o *SamlProviderResponse) SetSamlProviderName(v string)`

SetSamlProviderName sets SamlProviderName field to given value.

### HasSamlProviderName

`func (o *SamlProviderResponse) HasSamlProviderName() bool`

HasSamlProviderName returns a boolean if a field has been set.

### SetSamlProviderNameNil

`func (o *SamlProviderResponse) SetSamlProviderNameNil(b bool)`

 SetSamlProviderNameNil sets the value for SamlProviderName to be an explicit nil

### UnsetSamlProviderName
`func (o *SamlProviderResponse) UnsetSamlProviderName()`

UnsetSamlProviderName ensures that no value is present for SamlProviderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


