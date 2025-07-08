# ListenerForList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**Id** | **string** | ID | 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | The name of the listener. | 
**Protocol** | **string** | The protocol of the listener. | 
**ServerGroups** | Pointer to **[]interface{}** |  | [optional] 
**ServicePort** | **int32** | The port of the listener. | 
**State** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListenerForList

`func NewListenerForList(createdAt time.Time, createdBy string, id string, modifiedAt time.Time, modifiedBy string, name string, protocol string, servicePort int32, ) *ListenerForList`

NewListenerForList instantiates a new ListenerForList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListenerForListWithDefaults

`func NewListenerForListWithDefaults() *ListenerForList`

NewListenerForListWithDefaults instantiates a new ListenerForList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ListenerForList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListenerForList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListenerForList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ListenerForList) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListenerForList) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListenerForList) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetId

`func (o *ListenerForList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListenerForList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListenerForList) SetId(v string)`

SetId sets Id field to given value.


### GetModifiedAt

`func (o *ListenerForList) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ListenerForList) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ListenerForList) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ListenerForList) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ListenerForList) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ListenerForList) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *ListenerForList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListenerForList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListenerForList) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *ListenerForList) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ListenerForList) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ListenerForList) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetServerGroups

`func (o *ListenerForList) GetServerGroups() []interface{}`

GetServerGroups returns the ServerGroups field if non-nil, zero value otherwise.

### GetServerGroupsOk

`func (o *ListenerForList) GetServerGroupsOk() (*[]interface{}, bool)`

GetServerGroupsOk returns a tuple with the ServerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGroups

`func (o *ListenerForList) SetServerGroups(v []interface{})`

SetServerGroups sets ServerGroups field to given value.

### HasServerGroups

`func (o *ListenerForList) HasServerGroups() bool`

HasServerGroups returns a boolean if a field has been set.

### SetServerGroupsNil

`func (o *ListenerForList) SetServerGroupsNil(b bool)`

 SetServerGroupsNil sets the value for ServerGroups to be an explicit nil

### UnsetServerGroups
`func (o *ListenerForList) UnsetServerGroups()`

UnsetServerGroups ensures that no value is present for ServerGroups, not even an explicit nil
### GetServicePort

`func (o *ListenerForList) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ListenerForList) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ListenerForList) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.


### GetState

`func (o *ListenerForList) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListenerForList) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListenerForList) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ListenerForList) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *ListenerForList) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *ListenerForList) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


