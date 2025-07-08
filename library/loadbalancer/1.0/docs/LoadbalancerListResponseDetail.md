# LoadbalancerListResponseDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | ID | 
**ListenerCount** | Pointer to **NullableInt32** |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | Pointer to **NullableString** |  | [optional] 
**PublicNatIp** | Pointer to **NullableString** |  | [optional] 
**ServiceIp** | Pointer to **NullableString** |  | [optional] 
**SourceNatIp** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewLoadbalancerListResponseDetail

`func NewLoadbalancerListResponseDetail(createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, ) *LoadbalancerListResponseDetail`

NewLoadbalancerListResponseDetail instantiates a new LoadbalancerListResponseDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadbalancerListResponseDetailWithDefaults

`func NewLoadbalancerListResponseDetailWithDefaults() *LoadbalancerListResponseDetail`

NewLoadbalancerListResponseDetailWithDefaults instantiates a new LoadbalancerListResponseDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LoadbalancerListResponseDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadbalancerListResponseDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadbalancerListResponseDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LoadbalancerListResponseDetail) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LoadbalancerListResponseDetail) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LoadbalancerListResponseDetail) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *LoadbalancerListResponseDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadbalancerListResponseDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadbalancerListResponseDetail) SetId(v string)`

SetId sets Id field to given value.


### GetListenerCount

`func (o *LoadbalancerListResponseDetail) GetListenerCount() int32`

GetListenerCount returns the ListenerCount field if non-nil, zero value otherwise.

### GetListenerCountOk

`func (o *LoadbalancerListResponseDetail) GetListenerCountOk() (*int32, bool)`

GetListenerCountOk returns a tuple with the ListenerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenerCount

`func (o *LoadbalancerListResponseDetail) SetListenerCount(v int32)`

SetListenerCount sets ListenerCount field to given value.

### HasListenerCount

`func (o *LoadbalancerListResponseDetail) HasListenerCount() bool`

HasListenerCount returns a boolean if a field has been set.

### SetListenerCountNil

`func (o *LoadbalancerListResponseDetail) SetListenerCountNil(b bool)`

 SetListenerCountNil sets the value for ListenerCount to be an explicit nil

### UnsetListenerCount
`func (o *LoadbalancerListResponseDetail) UnsetListenerCount()`

UnsetListenerCount ensures that no value is present for ListenerCount, not even an explicit nil
### GetModifiedAt

`func (o *LoadbalancerListResponseDetail) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LoadbalancerListResponseDetail) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LoadbalancerListResponseDetail) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LoadbalancerListResponseDetail) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LoadbalancerListResponseDetail) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LoadbalancerListResponseDetail) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LoadbalancerListResponseDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadbalancerListResponseDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadbalancerListResponseDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadbalancerListResponseDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LoadbalancerListResponseDetail) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LoadbalancerListResponseDetail) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPublicNatIp

`func (o *LoadbalancerListResponseDetail) GetPublicNatIp() string`

GetPublicNatIp returns the PublicNatIp field if non-nil, zero value otherwise.

### GetPublicNatIpOk

`func (o *LoadbalancerListResponseDetail) GetPublicNatIpOk() (*string, bool)`

GetPublicNatIpOk returns a tuple with the PublicNatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNatIp

`func (o *LoadbalancerListResponseDetail) SetPublicNatIp(v string)`

SetPublicNatIp sets PublicNatIp field to given value.

### HasPublicNatIp

`func (o *LoadbalancerListResponseDetail) HasPublicNatIp() bool`

HasPublicNatIp returns a boolean if a field has been set.

### SetPublicNatIpNil

`func (o *LoadbalancerListResponseDetail) SetPublicNatIpNil(b bool)`

 SetPublicNatIpNil sets the value for PublicNatIp to be an explicit nil

### UnsetPublicNatIp
`func (o *LoadbalancerListResponseDetail) UnsetPublicNatIp()`

UnsetPublicNatIp ensures that no value is present for PublicNatIp, not even an explicit nil
### GetServiceIp

`func (o *LoadbalancerListResponseDetail) GetServiceIp() string`

GetServiceIp returns the ServiceIp field if non-nil, zero value otherwise.

### GetServiceIpOk

`func (o *LoadbalancerListResponseDetail) GetServiceIpOk() (*string, bool)`

GetServiceIpOk returns a tuple with the ServiceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIp

`func (o *LoadbalancerListResponseDetail) SetServiceIp(v string)`

SetServiceIp sets ServiceIp field to given value.

### HasServiceIp

`func (o *LoadbalancerListResponseDetail) HasServiceIp() bool`

HasServiceIp returns a boolean if a field has been set.

### SetServiceIpNil

`func (o *LoadbalancerListResponseDetail) SetServiceIpNil(b bool)`

 SetServiceIpNil sets the value for ServiceIp to be an explicit nil

### UnsetServiceIp
`func (o *LoadbalancerListResponseDetail) UnsetServiceIp()`

UnsetServiceIp ensures that no value is present for ServiceIp, not even an explicit nil
### GetSourceNatIp

`func (o *LoadbalancerListResponseDetail) GetSourceNatIp() string`

GetSourceNatIp returns the SourceNatIp field if non-nil, zero value otherwise.

### GetSourceNatIpOk

`func (o *LoadbalancerListResponseDetail) GetSourceNatIpOk() (*string, bool)`

GetSourceNatIpOk returns a tuple with the SourceNatIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceNatIp

`func (o *LoadbalancerListResponseDetail) SetSourceNatIp(v string)`

SetSourceNatIp sets SourceNatIp field to given value.

### HasSourceNatIp

`func (o *LoadbalancerListResponseDetail) HasSourceNatIp() bool`

HasSourceNatIp returns a boolean if a field has been set.

### SetSourceNatIpNil

`func (o *LoadbalancerListResponseDetail) SetSourceNatIpNil(b bool)`

 SetSourceNatIpNil sets the value for SourceNatIp to be an explicit nil

### UnsetSourceNatIp
`func (o *LoadbalancerListResponseDetail) UnsetSourceNatIp()`

UnsetSourceNatIp ensures that no value is present for SourceNatIp, not even an explicit nil
### GetState

`func (o *LoadbalancerListResponseDetail) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LoadbalancerListResponseDetail) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LoadbalancerListResponseDetail) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *LoadbalancerListResponseDetail) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *LoadbalancerListResponseDetail) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *LoadbalancerListResponseDetail) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


