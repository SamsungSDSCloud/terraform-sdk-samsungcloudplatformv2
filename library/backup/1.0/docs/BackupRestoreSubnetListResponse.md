# BackupRestoreSubnetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | Count | 
**Subnets** | Pointer to [**[]BackupRestoreSubnetResponse**](BackupRestoreSubnetResponse.md) |  | [optional] 

## Methods

### NewBackupRestoreSubnetListResponse

`func NewBackupRestoreSubnetListResponse(count int32, ) *BackupRestoreSubnetListResponse`

NewBackupRestoreSubnetListResponse instantiates a new BackupRestoreSubnetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRestoreSubnetListResponseWithDefaults

`func NewBackupRestoreSubnetListResponseWithDefaults() *BackupRestoreSubnetListResponse`

NewBackupRestoreSubnetListResponseWithDefaults instantiates a new BackupRestoreSubnetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *BackupRestoreSubnetListResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BackupRestoreSubnetListResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BackupRestoreSubnetListResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetSubnets

`func (o *BackupRestoreSubnetListResponse) GetSubnets() []BackupRestoreSubnetResponse`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *BackupRestoreSubnetListResponse) GetSubnetsOk() (*[]BackupRestoreSubnetResponse, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *BackupRestoreSubnetListResponse) SetSubnets(v []BackupRestoreSubnetResponse)`

SetSubnets sets Subnets field to given value.

### HasSubnets

`func (o *BackupRestoreSubnetListResponse) HasSubnets() bool`

HasSubnets returns a boolean if a field has been set.

### SetSubnetsNil

`func (o *BackupRestoreSubnetListResponse) SetSubnetsNil(b bool)`

 SetSubnetsNil sets the value for Subnets to be an explicit nil

### UnsetSubnets
`func (o *BackupRestoreSubnetListResponse) UnsetSubnets()`

UnsetSubnets ensures that no value is present for Subnets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


