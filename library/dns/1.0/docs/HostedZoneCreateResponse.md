# HostedZoneCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | current action in progress | 
**Attributes** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Email** | **string** | email | 
**HostedZoneType** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Id | 
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Masters** | Pointer to **[]string** |  | [optional] 
**Name** | **string** | name | 
**PoolId** | **string** | Designate Pool ID | 
**ProjectId** | **string** | ProjectID | 
**Serial** | **int32** | serial number | 
**Shared** | Pointer to **NullableBool** |  | [optional] 
**Status** | **string** | The status | 
**TransferredAt** | Pointer to **NullableString** |  | [optional] 
**Ttl** | Pointer to **NullableInt32** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewHostedZoneCreateResponse

`func NewHostedZoneCreateResponse(action string, email string, id string, name string, poolId string, projectId string, serial int32, status string, ) *HostedZoneCreateResponse`

NewHostedZoneCreateResponse instantiates a new HostedZoneCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedZoneCreateResponseWithDefaults

`func NewHostedZoneCreateResponseWithDefaults() *HostedZoneCreateResponse`

NewHostedZoneCreateResponseWithDefaults instantiates a new HostedZoneCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *HostedZoneCreateResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HostedZoneCreateResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HostedZoneCreateResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetAttributes

`func (o *HostedZoneCreateResponse) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *HostedZoneCreateResponse) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *HostedZoneCreateResponse) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *HostedZoneCreateResponse) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *HostedZoneCreateResponse) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *HostedZoneCreateResponse) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetCreatedAt

`func (o *HostedZoneCreateResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HostedZoneCreateResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HostedZoneCreateResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HostedZoneCreateResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *HostedZoneCreateResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *HostedZoneCreateResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *HostedZoneCreateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HostedZoneCreateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HostedZoneCreateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HostedZoneCreateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *HostedZoneCreateResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *HostedZoneCreateResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmail

`func (o *HostedZoneCreateResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *HostedZoneCreateResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *HostedZoneCreateResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetHostedZoneType

`func (o *HostedZoneCreateResponse) GetHostedZoneType() string`

GetHostedZoneType returns the HostedZoneType field if non-nil, zero value otherwise.

### GetHostedZoneTypeOk

`func (o *HostedZoneCreateResponse) GetHostedZoneTypeOk() (*string, bool)`

GetHostedZoneTypeOk returns a tuple with the HostedZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedZoneType

`func (o *HostedZoneCreateResponse) SetHostedZoneType(v string)`

SetHostedZoneType sets HostedZoneType field to given value.

### HasHostedZoneType

`func (o *HostedZoneCreateResponse) HasHostedZoneType() bool`

HasHostedZoneType returns a boolean if a field has been set.

### SetHostedZoneTypeNil

`func (o *HostedZoneCreateResponse) SetHostedZoneTypeNil(b bool)`

 SetHostedZoneTypeNil sets the value for HostedZoneType to be an explicit nil

### UnsetHostedZoneType
`func (o *HostedZoneCreateResponse) UnsetHostedZoneType()`

UnsetHostedZoneType ensures that no value is present for HostedZoneType, not even an explicit nil
### GetId

`func (o *HostedZoneCreateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostedZoneCreateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostedZoneCreateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLinks

`func (o *HostedZoneCreateResponse) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *HostedZoneCreateResponse) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *HostedZoneCreateResponse) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *HostedZoneCreateResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *HostedZoneCreateResponse) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *HostedZoneCreateResponse) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetMasters

`func (o *HostedZoneCreateResponse) GetMasters() []string`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *HostedZoneCreateResponse) GetMastersOk() (*[]string, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *HostedZoneCreateResponse) SetMasters(v []string)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *HostedZoneCreateResponse) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### SetMastersNil

`func (o *HostedZoneCreateResponse) SetMastersNil(b bool)`

 SetMastersNil sets the value for Masters to be an explicit nil

### UnsetMasters
`func (o *HostedZoneCreateResponse) UnsetMasters()`

UnsetMasters ensures that no value is present for Masters, not even an explicit nil
### GetName

`func (o *HostedZoneCreateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostedZoneCreateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostedZoneCreateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPoolId

`func (o *HostedZoneCreateResponse) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *HostedZoneCreateResponse) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *HostedZoneCreateResponse) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.


### GetProjectId

`func (o *HostedZoneCreateResponse) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *HostedZoneCreateResponse) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *HostedZoneCreateResponse) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSerial

`func (o *HostedZoneCreateResponse) GetSerial() int32`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *HostedZoneCreateResponse) GetSerialOk() (*int32, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *HostedZoneCreateResponse) SetSerial(v int32)`

SetSerial sets Serial field to given value.


### GetShared

`func (o *HostedZoneCreateResponse) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *HostedZoneCreateResponse) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *HostedZoneCreateResponse) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *HostedZoneCreateResponse) HasShared() bool`

HasShared returns a boolean if a field has been set.

### SetSharedNil

`func (o *HostedZoneCreateResponse) SetSharedNil(b bool)`

 SetSharedNil sets the value for Shared to be an explicit nil

### UnsetShared
`func (o *HostedZoneCreateResponse) UnsetShared()`

UnsetShared ensures that no value is present for Shared, not even an explicit nil
### GetStatus

`func (o *HostedZoneCreateResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostedZoneCreateResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostedZoneCreateResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransferredAt

`func (o *HostedZoneCreateResponse) GetTransferredAt() string`

GetTransferredAt returns the TransferredAt field if non-nil, zero value otherwise.

### GetTransferredAtOk

`func (o *HostedZoneCreateResponse) GetTransferredAtOk() (*string, bool)`

GetTransferredAtOk returns a tuple with the TransferredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferredAt

`func (o *HostedZoneCreateResponse) SetTransferredAt(v string)`

SetTransferredAt sets TransferredAt field to given value.

### HasTransferredAt

`func (o *HostedZoneCreateResponse) HasTransferredAt() bool`

HasTransferredAt returns a boolean if a field has been set.

### SetTransferredAtNil

`func (o *HostedZoneCreateResponse) SetTransferredAtNil(b bool)`

 SetTransferredAtNil sets the value for TransferredAt to be an explicit nil

### UnsetTransferredAt
`func (o *HostedZoneCreateResponse) UnsetTransferredAt()`

UnsetTransferredAt ensures that no value is present for TransferredAt, not even an explicit nil
### GetTtl

`func (o *HostedZoneCreateResponse) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *HostedZoneCreateResponse) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *HostedZoneCreateResponse) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *HostedZoneCreateResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### SetTtlNil

`func (o *HostedZoneCreateResponse) SetTtlNil(b bool)`

 SetTtlNil sets the value for Ttl to be an explicit nil

### UnsetTtl
`func (o *HostedZoneCreateResponse) UnsetTtl()`

UnsetTtl ensures that no value is present for Ttl, not even an explicit nil
### GetType

`func (o *HostedZoneCreateResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostedZoneCreateResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostedZoneCreateResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HostedZoneCreateResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *HostedZoneCreateResponse) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *HostedZoneCreateResponse) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUpdatedAt

`func (o *HostedZoneCreateResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HostedZoneCreateResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HostedZoneCreateResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HostedZoneCreateResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *HostedZoneCreateResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *HostedZoneCreateResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetVersion

`func (o *HostedZoneCreateResponse) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HostedZoneCreateResponse) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HostedZoneCreateResponse) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HostedZoneCreateResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HostedZoneCreateResponse) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HostedZoneCreateResponse) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


