# TransitGatewayVpcConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | Transit Gateway VPC Connection ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**State** | [**TransitGatewayVpcConnectionState**](TransitGatewayVpcConnectionState.md) | State | 
**TransitGatewayId** | **string** | Transit Gateway ID | 
**VpcId** | **string** | VPC Id | 
**VpcName** | **string** | VPC Name | 

## Methods

### NewTransitGatewayVpcConnection

`func NewTransitGatewayVpcConnection(accountId string, createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, state TransitGatewayVpcConnectionState, transitGatewayId string, vpcId string, vpcName string, ) *TransitGatewayVpcConnection`

NewTransitGatewayVpcConnection instantiates a new TransitGatewayVpcConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitGatewayVpcConnectionWithDefaults

`func NewTransitGatewayVpcConnectionWithDefaults() *TransitGatewayVpcConnection`

NewTransitGatewayVpcConnectionWithDefaults instantiates a new TransitGatewayVpcConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *TransitGatewayVpcConnection) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransitGatewayVpcConnection) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransitGatewayVpcConnection) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCreatedAt

`func (o *TransitGatewayVpcConnection) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TransitGatewayVpcConnection) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TransitGatewayVpcConnection) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *TransitGatewayVpcConnection) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TransitGatewayVpcConnection) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TransitGatewayVpcConnection) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *TransitGatewayVpcConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransitGatewayVpcConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransitGatewayVpcConnection) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *TransitGatewayVpcConnection) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *TransitGatewayVpcConnection) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *TransitGatewayVpcConnection) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *TransitGatewayVpcConnection) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *TransitGatewayVpcConnection) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *TransitGatewayVpcConnection) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetState

`func (o *TransitGatewayVpcConnection) GetState() TransitGatewayVpcConnectionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TransitGatewayVpcConnection) GetStateOk() (*TransitGatewayVpcConnectionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TransitGatewayVpcConnection) SetState(v TransitGatewayVpcConnectionState)`

SetState sets State field to given value.


### GetTransitGatewayId

`func (o *TransitGatewayVpcConnection) GetTransitGatewayId() string`

GetTransitGatewayId returns the TransitGatewayId field if non-nil, zero value otherwise.

### GetTransitGatewayIdOk

`func (o *TransitGatewayVpcConnection) GetTransitGatewayIdOk() (*string, bool)`

GetTransitGatewayIdOk returns a tuple with the TransitGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitGatewayId

`func (o *TransitGatewayVpcConnection) SetTransitGatewayId(v string)`

SetTransitGatewayId sets TransitGatewayId field to given value.


### GetVpcId

`func (o *TransitGatewayVpcConnection) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *TransitGatewayVpcConnection) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *TransitGatewayVpcConnection) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### GetVpcName

`func (o *TransitGatewayVpcConnection) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *TransitGatewayVpcConnection) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *TransitGatewayVpcConnection) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


