# LbHealthCheckShow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **NullableString** |  | 
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Description** | Pointer to **NullableString** |  | [optional] 
**HealthCheckCount** | Pointer to **int32** | LB Health Check count | [optional] [default to 3]
**HealthCheckInterval** | Pointer to **int32** | LB Health Check interval | [optional] [default to 5]
**HealthCheckPort** | **int32** | LB Health Check port | 
**HealthCheckTimeout** | Pointer to **int32** | LB Health Check timeout | [optional] [default to 5]
**HealthCheckType** | [**LbMonitorType**](LbMonitorType.md) | LB Health Check type | 
**HealthCheckUrl** | **NullableString** |  | 
**HttpMethod** | **NullableString** |  | 
**Id** | **string** | LB Health Check ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**LbMonitorProtocol**](LbMonitorProtocol.md) | LB Health Check Protocol | [optional] [default to LBMONITORPROTOCOL_TCP]
**RequestData** | **NullableString** |  | 
**ResponseCode** | **NullableString** |  | 
**State** | **string** | LB Health Check state | 
**SubnetId** | **NullableString** |  | 
**VpcId** | **NullableString** |  | 

## Methods

### NewLbHealthCheckShow

`func NewLbHealthCheckShow(accountId NullableString, createdAt time.Time, createdBy string, healthCheckPort int32, healthCheckType LbMonitorType, healthCheckUrl NullableString, httpMethod NullableString, id string, modifiedAt time.Time, modifiedBy string, requestData NullableString, responseCode NullableString, state string, subnetId NullableString, vpcId NullableString, ) *LbHealthCheckShow`

NewLbHealthCheckShow instantiates a new LbHealthCheckShow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLbHealthCheckShowWithDefaults

`func NewLbHealthCheckShowWithDefaults() *LbHealthCheckShow`

NewLbHealthCheckShowWithDefaults instantiates a new LbHealthCheckShow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LbHealthCheckShow) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LbHealthCheckShow) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LbHealthCheckShow) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *LbHealthCheckShow) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *LbHealthCheckShow) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetCreatedAt

`func (o *LbHealthCheckShow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LbHealthCheckShow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LbHealthCheckShow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LbHealthCheckShow) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LbHealthCheckShow) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LbHealthCheckShow) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDescription

`func (o *LbHealthCheckShow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LbHealthCheckShow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LbHealthCheckShow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LbHealthCheckShow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LbHealthCheckShow) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LbHealthCheckShow) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetHealthCheckCount

`func (o *LbHealthCheckShow) GetHealthCheckCount() int32`

GetHealthCheckCount returns the HealthCheckCount field if non-nil, zero value otherwise.

### GetHealthCheckCountOk

`func (o *LbHealthCheckShow) GetHealthCheckCountOk() (*int32, bool)`

GetHealthCheckCountOk returns a tuple with the HealthCheckCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckCount

`func (o *LbHealthCheckShow) SetHealthCheckCount(v int32)`

SetHealthCheckCount sets HealthCheckCount field to given value.

### HasHealthCheckCount

`func (o *LbHealthCheckShow) HasHealthCheckCount() bool`

HasHealthCheckCount returns a boolean if a field has been set.

### GetHealthCheckInterval

`func (o *LbHealthCheckShow) GetHealthCheckInterval() int32`

GetHealthCheckInterval returns the HealthCheckInterval field if non-nil, zero value otherwise.

### GetHealthCheckIntervalOk

`func (o *LbHealthCheckShow) GetHealthCheckIntervalOk() (*int32, bool)`

GetHealthCheckIntervalOk returns a tuple with the HealthCheckInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckInterval

`func (o *LbHealthCheckShow) SetHealthCheckInterval(v int32)`

SetHealthCheckInterval sets HealthCheckInterval field to given value.

### HasHealthCheckInterval

`func (o *LbHealthCheckShow) HasHealthCheckInterval() bool`

HasHealthCheckInterval returns a boolean if a field has been set.

### GetHealthCheckPort

`func (o *LbHealthCheckShow) GetHealthCheckPort() int32`

GetHealthCheckPort returns the HealthCheckPort field if non-nil, zero value otherwise.

### GetHealthCheckPortOk

`func (o *LbHealthCheckShow) GetHealthCheckPortOk() (*int32, bool)`

GetHealthCheckPortOk returns a tuple with the HealthCheckPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckPort

`func (o *LbHealthCheckShow) SetHealthCheckPort(v int32)`

SetHealthCheckPort sets HealthCheckPort field to given value.


### GetHealthCheckTimeout

`func (o *LbHealthCheckShow) GetHealthCheckTimeout() int32`

GetHealthCheckTimeout returns the HealthCheckTimeout field if non-nil, zero value otherwise.

### GetHealthCheckTimeoutOk

`func (o *LbHealthCheckShow) GetHealthCheckTimeoutOk() (*int32, bool)`

GetHealthCheckTimeoutOk returns a tuple with the HealthCheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckTimeout

`func (o *LbHealthCheckShow) SetHealthCheckTimeout(v int32)`

SetHealthCheckTimeout sets HealthCheckTimeout field to given value.

### HasHealthCheckTimeout

`func (o *LbHealthCheckShow) HasHealthCheckTimeout() bool`

HasHealthCheckTimeout returns a boolean if a field has been set.

### GetHealthCheckType

`func (o *LbHealthCheckShow) GetHealthCheckType() LbMonitorType`

GetHealthCheckType returns the HealthCheckType field if non-nil, zero value otherwise.

### GetHealthCheckTypeOk

`func (o *LbHealthCheckShow) GetHealthCheckTypeOk() (*LbMonitorType, bool)`

GetHealthCheckTypeOk returns a tuple with the HealthCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckType

`func (o *LbHealthCheckShow) SetHealthCheckType(v LbMonitorType)`

SetHealthCheckType sets HealthCheckType field to given value.


### GetHealthCheckUrl

`func (o *LbHealthCheckShow) GetHealthCheckUrl() string`

GetHealthCheckUrl returns the HealthCheckUrl field if non-nil, zero value otherwise.

### GetHealthCheckUrlOk

`func (o *LbHealthCheckShow) GetHealthCheckUrlOk() (*string, bool)`

GetHealthCheckUrlOk returns a tuple with the HealthCheckUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthCheckUrl

`func (o *LbHealthCheckShow) SetHealthCheckUrl(v string)`

SetHealthCheckUrl sets HealthCheckUrl field to given value.


### SetHealthCheckUrlNil

`func (o *LbHealthCheckShow) SetHealthCheckUrlNil(b bool)`

 SetHealthCheckUrlNil sets the value for HealthCheckUrl to be an explicit nil

### UnsetHealthCheckUrl
`func (o *LbHealthCheckShow) UnsetHealthCheckUrl()`

UnsetHealthCheckUrl ensures that no value is present for HealthCheckUrl, not even an explicit nil
### GetHttpMethod

`func (o *LbHealthCheckShow) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *LbHealthCheckShow) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *LbHealthCheckShow) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.


### SetHttpMethodNil

`func (o *LbHealthCheckShow) SetHttpMethodNil(b bool)`

 SetHttpMethodNil sets the value for HttpMethod to be an explicit nil

### UnsetHttpMethod
`func (o *LbHealthCheckShow) UnsetHttpMethod()`

UnsetHttpMethod ensures that no value is present for HttpMethod, not even an explicit nil
### GetId

`func (o *LbHealthCheckShow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LbHealthCheckShow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LbHealthCheckShow) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *LbHealthCheckShow) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LbHealthCheckShow) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LbHealthCheckShow) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LbHealthCheckShow) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LbHealthCheckShow) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LbHealthCheckShow) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *LbHealthCheckShow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LbHealthCheckShow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LbHealthCheckShow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LbHealthCheckShow) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LbHealthCheckShow) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LbHealthCheckShow) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProtocol

`func (o *LbHealthCheckShow) GetProtocol() LbMonitorProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LbHealthCheckShow) GetProtocolOk() (*LbMonitorProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LbHealthCheckShow) SetProtocol(v LbMonitorProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LbHealthCheckShow) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRequestData

`func (o *LbHealthCheckShow) GetRequestData() string`

GetRequestData returns the RequestData field if non-nil, zero value otherwise.

### GetRequestDataOk

`func (o *LbHealthCheckShow) GetRequestDataOk() (*string, bool)`

GetRequestDataOk returns a tuple with the RequestData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestData

`func (o *LbHealthCheckShow) SetRequestData(v string)`

SetRequestData sets RequestData field to given value.


### SetRequestDataNil

`func (o *LbHealthCheckShow) SetRequestDataNil(b bool)`

 SetRequestDataNil sets the value for RequestData to be an explicit nil

### UnsetRequestData
`func (o *LbHealthCheckShow) UnsetRequestData()`

UnsetRequestData ensures that no value is present for RequestData, not even an explicit nil
### GetResponseCode

`func (o *LbHealthCheckShow) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *LbHealthCheckShow) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *LbHealthCheckShow) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### SetResponseCodeNil

`func (o *LbHealthCheckShow) SetResponseCodeNil(b bool)`

 SetResponseCodeNil sets the value for ResponseCode to be an explicit nil

### UnsetResponseCode
`func (o *LbHealthCheckShow) UnsetResponseCode()`

UnsetResponseCode ensures that no value is present for ResponseCode, not even an explicit nil
### GetState

`func (o *LbHealthCheckShow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *LbHealthCheckShow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *LbHealthCheckShow) SetState(v string)`

SetState sets State field to given value.


### GetSubnetId

`func (o *LbHealthCheckShow) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LbHealthCheckShow) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LbHealthCheckShow) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### SetSubnetIdNil

`func (o *LbHealthCheckShow) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *LbHealthCheckShow) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetVpcId

`func (o *LbHealthCheckShow) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LbHealthCheckShow) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LbHealthCheckShow) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.


### SetVpcIdNil

`func (o *LbHealthCheckShow) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *LbHealthCheckShow) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


