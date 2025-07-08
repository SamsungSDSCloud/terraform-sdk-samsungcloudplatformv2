# LbHealthCheckList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**HealthCheckType** | [**LbMonitorType**](LbMonitorType.md) | LB Health Check type | 
**Id** | **string** | LB Health Check ID | 
**LbServerGroupCount** | **int32** | LB Server Group count | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | LB Health Check name | 
**Protocol** | [**LbMonitorProtocol**](LbMonitorProtocol.md) | LB Health Check Protocol | 
**State** | **string** | LB Health Check state | 
**SubnetId** | **string** | Service Subnet ID | 

## Methods

### NewLbHealthCheckList

`func NewLbHealthCheckList(createdAt time.Time, createdBy string, healthCheckType LbMonitorType, id string, lbServerGroupCount int32, modifiedAt time.Time, modifiedBy string, name string, protocol LbMonitorProtocol, state string, subnetId string, ) *LbHealthCheckList`

NewLbHealthCheckList instantiates a new LbHealthCheckList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbHealthCheckListWithDefaults

`func NewLbHealthCheckListWithDefaults() *LbHealthCheckList`

NewLbHealthCheckListWithDefaults instantiates a new LbHealthCheckList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LbHealthCheckList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LbHealthCheckList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LbHealthCheckList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LbHealthCheckList) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LbHealthCheckList) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LbHealthCheckList) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetHealthCheckType

`func (o *LbHealthCheckList) GetHealthCheckType() LbMonitorType`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *LbHealthCheckList) GetHealthCheckTypeOk() (*LbMonitorType, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *LbHealthCheckList) SetHealthCheckType(v LbMonitorType)`

SetHealthCheckType sets HealthCheckType field to given value.


### GetId

`func (o *LbHealthCheckList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LbHealthCheckList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LbHealthCheckList) SetId(v string)`

SetId sets Id field to given value.


### GetLbServerGroupCount

`func (o *LbHealthCheckList) GetLbServerGroupCount() int32`

GetLbServerGroupCount returns the LbServerGroupCount field if non-nil, zero value otherwise.

### GetLbServerGroupCountOk

`func (o *LbHealthCheckList) GetLbServerGroupCountOk() (*int32, bool)`

GetLbServerGroupCountOk returns a tuple with the LbServerGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLbServerGroupCount

`func (o *LbHealthCheckList) SetLbServerGroupCount(v int32)`

SetLbServerGroupCount sets LbServerGroupCount field to given value.


### GetModifiedAt

`func (o *LbHealthCheckList) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LbHealthCheckList) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LbHealthCheckList) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LbHealthCheckList) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LbHealthCheckList) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LbHealthCheckList) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LbHealthCheckList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbHealthCheckList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbHealthCheckList) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *LbHealthCheckList) GetProtocol() LbMonitorProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LbHealthCheckList) GetProtocolOk() (*LbMonitorProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LbHealthCheckList) SetProtocol(v LbMonitorProtocol)`

SetProtocol sets Protocol field to given value.


### GetState

`func (o *LbHealthCheckList) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LbHealthCheckList) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LbHealthCheckList) SetState(v string)`

SetState sets State field to given value.


### GetSubnetId

`func (o *LbHealthCheckList) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LbHealthCheckList) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LbHealthCheckList) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


