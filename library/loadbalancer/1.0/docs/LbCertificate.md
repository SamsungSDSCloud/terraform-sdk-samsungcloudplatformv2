# LbCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewLbCertificate

`func NewLbCertificate(createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, ) *LbCertificate`

NewLbCertificate instantiates a new LbCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbCertificateWithDefaults

`func NewLbCertificateWithDefaults() *LbCertificate`

NewLbCertificateWithDefaults instantiates a new LbCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertKind

`func (o *LbCertificate) GetCertKind() string`

GetCertKind returns the CertKind field if non-nil, zero value otherwise.

### GetCertKindOk

`func (o *LbCertificate) GetCertKindOk() (*string, bool)`

GetCertKindOk returns a tuple with the CertKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertKind

`func (o *LbCertificate) SetCertKind(v string)`

SetCertKind sets CertKind field to given value.

### HasCertKind

`func (o *LbCertificate) HasCertKind() bool`

HasCertKind returns a boolean if a field has been set.

### GetCn

`func (o *LbCertificate) GetCn() string`

GetCn returns the Cn field if non-nil, zero value otherwise.

### GetCnOk

`func (o *LbCertificate) GetCnOk() (*string, bool)`

GetCnOk returns a tuple with the Cn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCn

`func (o *LbCertificate) SetCn(v string)`

SetCn sets Cn field to given value.

### HasCn

`func (o *LbCertificate) HasCn() bool`

HasCn returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LbCertificate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LbCertificate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LbCertificate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LbCertificate) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LbCertificate) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LbCertificate) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *LbCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LbCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LbCertificate) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *LbCertificate) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LbCertificate) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LbCertificate) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LbCertificate) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LbCertificate) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LbCertificate) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LbCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LbCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotAfterDt

`func (o *LbCertificate) GetNotAfterDt() time.Time`

GetNotAfterDt returns the NotAfterDt field if non-nil, zero value otherwise.

### GetNotAfterDtOk

`func (o *LbCertificate) GetNotAfterDtOk() (*time.Time, bool)`

GetNotAfterDtOk returns a tuple with the NotAfterDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfterDt

`func (o *LbCertificate) SetNotAfterDt(v time.Time)`

SetNotAfterDt sets NotAfterDt field to given value.

### HasNotAfterDt

`func (o *LbCertificate) HasNotAfterDt() bool`

HasNotAfterDt returns a boolean if a field has been set.

### GetNotBeforeDt

`func (o *LbCertificate) GetNotBeforeDt() time.Time`

GetNotBeforeDt returns the NotBeforeDt field if non-nil, zero value otherwise.

### GetNotBeforeDtOk

`func (o *LbCertificate) GetNotBeforeDtOk() (*time.Time, bool)`

GetNotBeforeDtOk returns a tuple with the NotBeforeDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDt

`func (o *LbCertificate) SetNotBeforeDt(v time.Time)`

SetNotBeforeDt sets NotBeforeDt field to given value.

### HasNotBeforeDt

`func (o *LbCertificate) HasNotBeforeDt() bool`

HasNotBeforeDt returns a boolean if a field has been set.

### GetState

`func (o *LbCertificate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LbCertificate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LbCertificate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LbCertificate) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


