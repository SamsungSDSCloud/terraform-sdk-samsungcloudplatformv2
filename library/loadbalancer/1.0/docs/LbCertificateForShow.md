# LbCertificateForShow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The project ID of the certificate | 
**CertBody** | **string** | The body of the certificate | 
**CertChain** | Pointer to **NullableString** |  | [optional] 
**CertKind** | Pointer to **string** |  | [optional] 
**Cn** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | Pointer to **string** |  | [optional] 
**NotAfterDt** | Pointer to **time.Time** |  | [optional] 
**NotBeforeDt** | Pointer to **time.Time** |  | [optional] 
**Organization** | **string** | The organization of the certificate | 
**PrivateKey** | **string** | The private key of the certificate | 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewLbCertificateForShow

`func NewLbCertificateForShow(accountId string, certBody string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, organization string, privateKey string, ) *LbCertificateForShow`

NewLbCertificateForShow instantiates a new LbCertificateForShow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbCertificateForShowWithDefaults

`func NewLbCertificateForShowWithDefaults() *LbCertificateForShow`

NewLbCertificateForShowWithDefaults instantiates a new LbCertificateForShow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LbCertificateForShow) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LbCertificateForShow) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LbCertificateForShow) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCertBody

`func (o *LbCertificateForShow) GetCertBody() string`

GetCertBody returns the CertBody field if non-nil, zero value otherwise.

### GetCertBodyOk

`func (o *LbCertificateForShow) GetCertBodyOk() (*string, bool)`

GetCertBodyOk returns a tuple with the CertBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertBody

`func (o *LbCertificateForShow) SetCertBody(v string)`

SetCertBody sets CertBody field to given value.


### GetCertChain

`func (o *LbCertificateForShow) GetCertChain() string`

GetCertChain returns the CertChain field if non-nil, zero value otherwise.

### GetCertChainOk

`func (o *LbCertificateForShow) GetCertChainOk() (*string, bool)`

GetCertChainOk returns a tuple with the CertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertChain

`func (o *LbCertificateForShow) SetCertChain(v string)`

SetCertChain sets CertChain field to given value.

### HasCertChain

`func (o *LbCertificateForShow) HasCertChain() bool`

HasCertChain returns a boolean if a field has been set.

### SetCertChainNil

`func (o *LbCertificateForShow) SetCertChainNil(b bool)`

 SetCertChainNil sets the value for CertChain to be an explicit nil

### UnsetCertChain
`func (o *LbCertificateForShow) UnsetCertChain()`

UnsetCertChain ensures that no value is present for CertChain, not even an explicit nil
### GetCertKind

`func (o *LbCertificateForShow) GetCertKind() string`

GetCertKind returns the CertKind field if non-nil, zero value otherwise.

### GetCertKindOk

`func (o *LbCertificateForShow) GetCertKindOk() (*string, bool)`

GetCertKindOk returns a tuple with the CertKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertKind

`func (o *LbCertificateForShow) SetCertKind(v string)`

SetCertKind sets CertKind field to given value.

### HasCertKind

`func (o *LbCertificateForShow) HasCertKind() bool`

HasCertKind returns a boolean if a field has been set.

### GetCn

`func (o *LbCertificateForShow) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *LbCertificateForShow) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *LbCertificateForShow) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *LbCertificateForShow) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LbCertificateForShow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LbCertificateForShow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LbCertificateForShow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LbCertificateForShow) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LbCertificateForShow) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LbCertificateForShow) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *LbCertificateForShow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LbCertificateForShow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LbCertificateForShow) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *LbCertificateForShow) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LbCertificateForShow) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LbCertificateForShow) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LbCertificateForShow) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LbCertificateForShow) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LbCertificateForShow) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LbCertificateForShow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbCertificateForShow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbCertificateForShow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LbCertificateForShow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotAfterDt

`func (o *LbCertificateForShow) GetNotAfterDt() time.Time`

GetNotAfterDt returns the NotAfterDt field if non-nil, zero value otherwise.

### GetNotAfterDtOk

`func (o *LbCertificateForShow) GetNotAfterDtOk() (*time.Time, bool)`

GetNotAfterDtOk returns a tuple with the NotAfterDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfterDt

`func (o *LbCertificateForShow) SetNotAfterDt(v time.Time)`

SetNotAfterDt sets NotAfterDt field to given value.

### HasNotAfterDt

`func (o *LbCertificateForShow) HasNotAfterDt() bool`

HasNotAfterDt returns a boolean if a field has been set.

### GetNotBeforeDt

`func (o *LbCertificateForShow) GetNotBeforeDt() time.Time`

GetNotBeforeDt returns the NotBeforeDt field if non-nil, zero value otherwise.

### GetNotBeforeDtOk

`func (o *LbCertificateForShow) GetNotBeforeDtOk() (*time.Time, bool)`

GetNotBeforeDtOk returns a tuple with the NotBeforeDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDt

`func (o *LbCertificateForShow) SetNotBeforeDt(v time.Time)`

SetNotBeforeDt sets NotBeforeDt field to given value.

### HasNotBeforeDt

`func (o *LbCertificateForShow) HasNotBeforeDt() bool`

HasNotBeforeDt returns a boolean if a field has been set.

### GetOrganization

`func (o *LbCertificateForShow) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *LbCertificateForShow) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *LbCertificateForShow) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetPrivateKey

`func (o *LbCertificateForShow) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *LbCertificateForShow) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *LbCertificateForShow) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetState

`func (o *LbCertificateForShow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LbCertificateForShow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LbCertificateForShow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LbCertificateForShow) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


