# RedisClusterDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DbaasEngine** | **string** | DBaaS engine | 
**DbaasEngineVersionName** | **string** | DBaaS engine version name | 
**HaEnabled** | Pointer to **bool** | HA availability | [optional] [default to false]
**Id** | **string** | ID | 
**InitConfigOption** | [**RedisInitConfigOptionResponse**](RedisInitConfigOptionResponse.md) |  | 
**InstanceCount** | Pointer to **int32** | Instance Count | [optional] [default to 0]
**InstanceGroups** | [**[]RedisInstanceGroupResponse**](RedisInstanceGroupResponse.md) |  | 
**MaintenanceOption** | Pointer to [**NullableMaintenanceResponseOption**](MaintenanceResponseOption.md) |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **bool** | NAT availability | [optional] [default to false]
**ProductImageType** | **string** | Product image type | 
**ProductType** | [**ProductType**](ProductType.md) | Product type | 
**RoleType** | [**NullableClusterRoleType**](ClusterRoleType.md) |  | 
**ServiceState** | [**ServiceState**](ServiceState.md) | Service state | 
**SoftwareVersion** | **string** | Software version | 
**SubnetId** | **string** | Subnet ID | 
**Timezone** | **string** | Timezone | 

## Methods

### NewRedisClusterDetailResponse

`func NewRedisClusterDetailResponse(accountId string, createdAt time.Time, createdBy string, dbaasEngine string, dbaasEngineVersionName string, id string, initConfigOption RedisInitConfigOptionResponse, instanceGroups []RedisInstanceGroupResponse, modifiedAt time.Time, modifiedBy string, name string, productImageType string, productType ProductType, roleType NullableClusterRoleType, serviceState ServiceState, softwareVersion string, subnetId string, timezone string, ) *RedisClusterDetailResponse`

NewRedisClusterDetailResponse instantiates a new RedisClusterDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedisClusterDetailResponseWithDefaults

`func NewRedisClusterDetailResponseWithDefaults() *RedisClusterDetailResponse`

NewRedisClusterDetailResponseWithDefaults instantiates a new RedisClusterDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RedisClusterDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RedisClusterDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RedisClusterDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAllowableIpAddresses

`func (o *RedisClusterDetailResponse) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *RedisClusterDetailResponse) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *RedisClusterDetailResponse) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *RedisClusterDetailResponse) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RedisClusterDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RedisClusterDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RedisClusterDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *RedisClusterDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RedisClusterDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RedisClusterDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDbaasEngine

`func (o *RedisClusterDetailResponse) GetDbaasEngine() string`

GetDbaasEngine returns the DbaasEngine field if non-nil, zero value otherwise.

### GetDbaasEngineOk

`func (o *RedisClusterDetailResponse) GetDbaasEngineOk() (*string, bool)`

GetDbaasEngineOk returns a tuple with the DbaasEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngine

`func (o *RedisClusterDetailResponse) SetDbaasEngine(v string)`

SetDbaasEngine sets DbaasEngine field to given value.


### GetDbaasEngineVersionName

`func (o *RedisClusterDetailResponse) GetDbaasEngineVersionName() string`

GetDbaasEngineVersionName returns the DbaasEngineVersionName field if non-nil, zero value otherwise.

### GetDbaasEngineVersionNameOk

`func (o *RedisClusterDetailResponse) GetDbaasEngineVersionNameOk() (*string, bool)`

GetDbaasEngineVersionNameOk returns a tuple with the DbaasEngineVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionName

`func (o *RedisClusterDetailResponse) SetDbaasEngineVersionName(v string)`

SetDbaasEngineVersionName sets DbaasEngineVersionName field to given value.


### GetHaEnabled

`func (o *RedisClusterDetailResponse) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *RedisClusterDetailResponse) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *RedisClusterDetailResponse) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *RedisClusterDetailResponse) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetId

`func (o *RedisClusterDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedisClusterDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedisClusterDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInitConfigOption

`func (o *RedisClusterDetailResponse) GetInitConfigOption() RedisInitConfigOptionResponse`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *RedisClusterDetailResponse) GetInitConfigOptionOk() (*RedisInitConfigOptionResponse, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *RedisClusterDetailResponse) SetInitConfigOption(v RedisInitConfigOptionResponse)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceCount

`func (o *RedisClusterDetailResponse) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *RedisClusterDetailResponse) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *RedisClusterDetailResponse) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *RedisClusterDetailResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceGroups

`func (o *RedisClusterDetailResponse) GetInstanceGroups() []RedisInstanceGroupResponse`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *RedisClusterDetailResponse) GetInstanceGroupsOk() (*[]RedisInstanceGroupResponse, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *RedisClusterDetailResponse) SetInstanceGroups(v []RedisInstanceGroupResponse)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetMaintenanceOption

`func (o *RedisClusterDetailResponse) GetMaintenanceOption() MaintenanceResponseOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *RedisClusterDetailResponse) GetMaintenanceOptionOk() (*MaintenanceResponseOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *RedisClusterDetailResponse) SetMaintenanceOption(v MaintenanceResponseOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *RedisClusterDetailResponse) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *RedisClusterDetailResponse) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *RedisClusterDetailResponse) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetModifiedAt

`func (o *RedisClusterDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *RedisClusterDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *RedisClusterDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *RedisClusterDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *RedisClusterDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *RedisClusterDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *RedisClusterDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RedisClusterDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RedisClusterDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *RedisClusterDetailResponse) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *RedisClusterDetailResponse) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *RedisClusterDetailResponse) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *RedisClusterDetailResponse) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### GetProductImageType

`func (o *RedisClusterDetailResponse) GetProductImageType() string`

GetProductImageType returns the ProductImageType field if non-nil, zero value otherwise.

### GetProductImageTypeOk

`func (o *RedisClusterDetailResponse) GetProductImageTypeOk() (*string, bool)`

GetProductImageTypeOk returns a tuple with the ProductImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImageType

`func (o *RedisClusterDetailResponse) SetProductImageType(v string)`

SetProductImageType sets ProductImageType field to given value.


### GetProductType

`func (o *RedisClusterDetailResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *RedisClusterDetailResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *RedisClusterDetailResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetRoleType

`func (o *RedisClusterDetailResponse) GetRoleType() ClusterRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *RedisClusterDetailResponse) GetRoleTypeOk() (*ClusterRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *RedisClusterDetailResponse) SetRoleType(v ClusterRoleType)`

SetRoleType sets RoleType field to given value.


### SetRoleTypeNil

`func (o *RedisClusterDetailResponse) SetRoleTypeNil(b bool)`

 SetRoleTypeNil sets the value for RoleType to be an explicit nil

### UnsetRoleType
`func (o *RedisClusterDetailResponse) UnsetRoleType()`

UnsetRoleType ensures that no value is present for RoleType, not even an explicit nil
### GetServiceState

`func (o *RedisClusterDetailResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *RedisClusterDetailResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *RedisClusterDetailResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.


### GetSoftwareVersion

`func (o *RedisClusterDetailResponse) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *RedisClusterDetailResponse) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *RedisClusterDetailResponse) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetSubnetId

`func (o *RedisClusterDetailResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *RedisClusterDetailResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *RedisClusterDetailResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTimezone

`func (o *RedisClusterDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RedisClusterDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RedisClusterDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


