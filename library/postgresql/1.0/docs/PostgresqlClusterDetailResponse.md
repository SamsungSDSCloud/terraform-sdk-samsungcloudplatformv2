# PostgresqlClusterDetailResponse

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
**InitConfigOption** | [**PostgresqlInitConfigOptionDetail**](PostgresqlInitConfigOptionDetail.md) | DB initial configuration option | 
**InstanceCount** | Pointer to **int32** | Instance Count | [optional] [default to 0]
**InstanceGroups** | [**[]InstanceGroupResponse**](InstanceGroupResponse.md) | Instance groups list | 
**MaintenanceOption** | Pointer to [**NullableMaintenanceResponseOption**](MaintenanceResponseOption.md) |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Cluster name | 
**NatEnabled** | Pointer to **bool** | NAT availability | [optional] [default to false]
**OriginClusterId** | Pointer to **NullableString** |  | [optional] 
**ProductImageType** | **string** | Product image type | 
**ProductType** | [**ProductType**](ProductType.md) | Product type | 
**Replicas** | **[]string** | Replicas list | 
**RoleType** | [**NullableClusterRoleType**](ClusterRoleType.md) |  | 
**ServiceState** | [**ServiceState**](ServiceState.md) | Service state | 
**SoftwareVersion** | **string** | Software version | 
**SubnetId** | **string** | Subnet ID | 
**Timezone** | **string** | Timezone | 
**VipPublicIpAddress** | Pointer to **NullableString** |  | [optional] 
**VipPublicIpId** | Pointer to **NullableString** |  | [optional] 
**VirtualIpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPostgresqlClusterDetailResponse

`func NewPostgresqlClusterDetailResponse(accountId string, createdAt time.Time, createdBy string, dbaasEngine string, dbaasEngineVersionName string, id string, initConfigOption PostgresqlInitConfigOptionDetail, instanceGroups []InstanceGroupResponse, modifiedAt time.Time, modifiedBy string, name string, productImageType string, productType ProductType, replicas []string, roleType NullableClusterRoleType, serviceState ServiceState, softwareVersion string, subnetId string, timezone string, ) *PostgresqlClusterDetailResponse`

NewPostgresqlClusterDetailResponse instantiates a new PostgresqlClusterDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlClusterDetailResponseWithDefaults

`func NewPostgresqlClusterDetailResponseWithDefaults() *PostgresqlClusterDetailResponse`

NewPostgresqlClusterDetailResponseWithDefaults instantiates a new PostgresqlClusterDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *PostgresqlClusterDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PostgresqlClusterDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PostgresqlClusterDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAllowableIpAddresses

`func (o *PostgresqlClusterDetailResponse) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *PostgresqlClusterDetailResponse) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *PostgresqlClusterDetailResponse) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *PostgresqlClusterDetailResponse) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PostgresqlClusterDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostgresqlClusterDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostgresqlClusterDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *PostgresqlClusterDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PostgresqlClusterDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PostgresqlClusterDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDbaasEngine

`func (o *PostgresqlClusterDetailResponse) GetDbaasEngine() string`

GetDbaasEngine returns the DbaasEngine field if non-nil, zero value otherwise.

### GetDbaasEngineOk

`func (o *PostgresqlClusterDetailResponse) GetDbaasEngineOk() (*string, bool)`

GetDbaasEngineOk returns a tuple with the DbaasEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngine

`func (o *PostgresqlClusterDetailResponse) SetDbaasEngine(v string)`

SetDbaasEngine sets DbaasEngine field to given value.


### GetDbaasEngineVersionName

`func (o *PostgresqlClusterDetailResponse) GetDbaasEngineVersionName() string`

GetDbaasEngineVersionName returns the DbaasEngineVersionName field if non-nil, zero value otherwise.

### GetDbaasEngineVersionNameOk

`func (o *PostgresqlClusterDetailResponse) GetDbaasEngineVersionNameOk() (*string, bool)`

GetDbaasEngineVersionNameOk returns a tuple with the DbaasEngineVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionName

`func (o *PostgresqlClusterDetailResponse) SetDbaasEngineVersionName(v string)`

SetDbaasEngineVersionName sets DbaasEngineVersionName field to given value.


### GetHaEnabled

`func (o *PostgresqlClusterDetailResponse) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *PostgresqlClusterDetailResponse) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *PostgresqlClusterDetailResponse) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *PostgresqlClusterDetailResponse) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetId

`func (o *PostgresqlClusterDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostgresqlClusterDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostgresqlClusterDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInitConfigOption

`func (o *PostgresqlClusterDetailResponse) GetInitConfigOption() PostgresqlInitConfigOptionDetail`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *PostgresqlClusterDetailResponse) GetInitConfigOptionOk() (*PostgresqlInitConfigOptionDetail, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *PostgresqlClusterDetailResponse) SetInitConfigOption(v PostgresqlInitConfigOptionDetail)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceCount

`func (o *PostgresqlClusterDetailResponse) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *PostgresqlClusterDetailResponse) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *PostgresqlClusterDetailResponse) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *PostgresqlClusterDetailResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceGroups

`func (o *PostgresqlClusterDetailResponse) GetInstanceGroups() []InstanceGroupResponse`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *PostgresqlClusterDetailResponse) GetInstanceGroupsOk() (*[]InstanceGroupResponse, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *PostgresqlClusterDetailResponse) SetInstanceGroups(v []InstanceGroupResponse)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetMaintenanceOption

`func (o *PostgresqlClusterDetailResponse) GetMaintenanceOption() MaintenanceResponseOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *PostgresqlClusterDetailResponse) GetMaintenanceOptionOk() (*MaintenanceResponseOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *PostgresqlClusterDetailResponse) SetMaintenanceOption(v MaintenanceResponseOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *PostgresqlClusterDetailResponse) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *PostgresqlClusterDetailResponse) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *PostgresqlClusterDetailResponse) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetModifiedAt

`func (o *PostgresqlClusterDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *PostgresqlClusterDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *PostgresqlClusterDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *PostgresqlClusterDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *PostgresqlClusterDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *PostgresqlClusterDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *PostgresqlClusterDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostgresqlClusterDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostgresqlClusterDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *PostgresqlClusterDetailResponse) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *PostgresqlClusterDetailResponse) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *PostgresqlClusterDetailResponse) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *PostgresqlClusterDetailResponse) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### GetOriginClusterId

`func (o *PostgresqlClusterDetailResponse) GetOriginClusterId() string`

GetOriginClusterId returns the OriginClusterId field if non-nil, zero value otherwise.

### GetOriginClusterIdOk

`func (o *PostgresqlClusterDetailResponse) GetOriginClusterIdOk() (*string, bool)`

GetOriginClusterIdOk returns a tuple with the OriginClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginClusterId

`func (o *PostgresqlClusterDetailResponse) SetOriginClusterId(v string)`

SetOriginClusterId sets OriginClusterId field to given value.

### HasOriginClusterId

`func (o *PostgresqlClusterDetailResponse) HasOriginClusterId() bool`

HasOriginClusterId returns a boolean if a field has been set.

### SetOriginClusterIdNil

`func (o *PostgresqlClusterDetailResponse) SetOriginClusterIdNil(b bool)`

 SetOriginClusterIdNil sets the value for OriginClusterId to be an explicit nil

### UnsetOriginClusterId
`func (o *PostgresqlClusterDetailResponse) UnsetOriginClusterId()`

UnsetOriginClusterId ensures that no value is present for OriginClusterId, not even an explicit nil
### GetProductImageType

`func (o *PostgresqlClusterDetailResponse) GetProductImageType() string`

GetProductImageType returns the ProductImageType field if non-nil, zero value otherwise.

### GetProductImageTypeOk

`func (o *PostgresqlClusterDetailResponse) GetProductImageTypeOk() (*string, bool)`

GetProductImageTypeOk returns a tuple with the ProductImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImageType

`func (o *PostgresqlClusterDetailResponse) SetProductImageType(v string)`

SetProductImageType sets ProductImageType field to given value.


### GetProductType

`func (o *PostgresqlClusterDetailResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *PostgresqlClusterDetailResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *PostgresqlClusterDetailResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetReplicas

`func (o *PostgresqlClusterDetailResponse) GetReplicas() []string`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PostgresqlClusterDetailResponse) GetReplicasOk() (*[]string, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PostgresqlClusterDetailResponse) SetReplicas(v []string)`

SetReplicas sets Replicas field to given value.


### GetRoleType

`func (o *PostgresqlClusterDetailResponse) GetRoleType() ClusterRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *PostgresqlClusterDetailResponse) GetRoleTypeOk() (*ClusterRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *PostgresqlClusterDetailResponse) SetRoleType(v ClusterRoleType)`

SetRoleType sets RoleType field to given value.


### SetRoleTypeNil

`func (o *PostgresqlClusterDetailResponse) SetRoleTypeNil(b bool)`

 SetRoleTypeNil sets the value for RoleType to be an explicit nil

### UnsetRoleType
`func (o *PostgresqlClusterDetailResponse) UnsetRoleType()`

UnsetRoleType ensures that no value is present for RoleType, not even an explicit nil
### GetServiceState

`func (o *PostgresqlClusterDetailResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *PostgresqlClusterDetailResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *PostgresqlClusterDetailResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.


### GetSoftwareVersion

`func (o *PostgresqlClusterDetailResponse) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *PostgresqlClusterDetailResponse) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *PostgresqlClusterDetailResponse) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetSubnetId

`func (o *PostgresqlClusterDetailResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *PostgresqlClusterDetailResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *PostgresqlClusterDetailResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTimezone

`func (o *PostgresqlClusterDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *PostgresqlClusterDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *PostgresqlClusterDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetVipPublicIpAddress

`func (o *PostgresqlClusterDetailResponse) GetVipPublicIpAddress() string`

GetVipPublicIpAddress returns the VipPublicIpAddress field if non-nil, zero value otherwise.

### GetVipPublicIpAddressOk

`func (o *PostgresqlClusterDetailResponse) GetVipPublicIpAddressOk() (*string, bool)`

GetVipPublicIpAddressOk returns a tuple with the VipPublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPublicIpAddress

`func (o *PostgresqlClusterDetailResponse) SetVipPublicIpAddress(v string)`

SetVipPublicIpAddress sets VipPublicIpAddress field to given value.

### HasVipPublicIpAddress

`func (o *PostgresqlClusterDetailResponse) HasVipPublicIpAddress() bool`

HasVipPublicIpAddress returns a boolean if a field has been set.

### SetVipPublicIpAddressNil

`func (o *PostgresqlClusterDetailResponse) SetVipPublicIpAddressNil(b bool)`

 SetVipPublicIpAddressNil sets the value for VipPublicIpAddress to be an explicit nil

### UnsetVipPublicIpAddress
`func (o *PostgresqlClusterDetailResponse) UnsetVipPublicIpAddress()`

UnsetVipPublicIpAddress ensures that no value is present for VipPublicIpAddress, not even an explicit nil
### GetVipPublicIpId

`func (o *PostgresqlClusterDetailResponse) GetVipPublicIpId() string`

GetVipPublicIpId returns the VipPublicIpId field if non-nil, zero value otherwise.

### GetVipPublicIpIdOk

`func (o *PostgresqlClusterDetailResponse) GetVipPublicIpIdOk() (*string, bool)`

GetVipPublicIpIdOk returns a tuple with the VipPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPublicIpId

`func (o *PostgresqlClusterDetailResponse) SetVipPublicIpId(v string)`

SetVipPublicIpId sets VipPublicIpId field to given value.

### HasVipPublicIpId

`func (o *PostgresqlClusterDetailResponse) HasVipPublicIpId() bool`

HasVipPublicIpId returns a boolean if a field has been set.

### SetVipPublicIpIdNil

`func (o *PostgresqlClusterDetailResponse) SetVipPublicIpIdNil(b bool)`

 SetVipPublicIpIdNil sets the value for VipPublicIpId to be an explicit nil

### UnsetVipPublicIpId
`func (o *PostgresqlClusterDetailResponse) UnsetVipPublicIpId()`

UnsetVipPublicIpId ensures that no value is present for VipPublicIpId, not even an explicit nil
### GetVirtualIpAddress

`func (o *PostgresqlClusterDetailResponse) GetVirtualIpAddress() string`

GetVirtualIpAddress returns the VirtualIpAddress field if non-nil, zero value otherwise.

### GetVirtualIpAddressOk

`func (o *PostgresqlClusterDetailResponse) GetVirtualIpAddressOk() (*string, bool)`

GetVirtualIpAddressOk returns a tuple with the VirtualIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIpAddress

`func (o *PostgresqlClusterDetailResponse) SetVirtualIpAddress(v string)`

SetVirtualIpAddress sets VirtualIpAddress field to given value.

### HasVirtualIpAddress

`func (o *PostgresqlClusterDetailResponse) HasVirtualIpAddress() bool`

HasVirtualIpAddress returns a boolean if a field has been set.

### SetVirtualIpAddressNil

`func (o *PostgresqlClusterDetailResponse) SetVirtualIpAddressNil(b bool)`

 SetVirtualIpAddressNil sets the value for VirtualIpAddress to be an explicit nil

### UnsetVirtualIpAddress
`func (o *PostgresqlClusterDetailResponse) UnsetVirtualIpAddress()`

UnsetVirtualIpAddress ensures that no value is present for VirtualIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


