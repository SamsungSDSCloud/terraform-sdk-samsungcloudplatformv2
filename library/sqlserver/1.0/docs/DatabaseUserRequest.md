# DatabaseUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** | Database User Name | 
**State** | Pointer to [**NullableDatabaseUserState**](DatabaseUserState.md) |  | [optional] 

## Methods

### NewDatabaseUserRequest

`func NewDatabaseUserRequest(name string, ) *DatabaseUserRequest`

NewDatabaseUserRequest instantiates a new DatabaseUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseUserRequestWithDefaults

`func NewDatabaseUserRequestWithDefaults() *DatabaseUserRequest`

NewDatabaseUserRequestWithDefaults instantiates a new DatabaseUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DatabaseUserRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatabaseUserRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatabaseUserRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatabaseUserRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DatabaseUserRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DatabaseUserRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *DatabaseUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetState

`func (o *DatabaseUserRequest) GetState() DatabaseUserState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DatabaseUserRequest) GetStateOk() (*DatabaseUserState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DatabaseUserRequest) SetState(v DatabaseUserState)`

SetState sets State field to given value.

### HasState

`func (o *DatabaseUserRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *DatabaseUserRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *DatabaseUserRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


