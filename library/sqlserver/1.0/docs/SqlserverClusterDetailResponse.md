# SqlserverClusterDetailResponse

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
**InitConfigOption** | [**SqlserverInitConfigOptionResponse**](SqlserverInitConfigOptionResponse.md) |  | 
**InstanceCount** | Pointer to **int32** | Instance Count | [optional] [default to 0]
**InstanceGroups** | [**[]SqlserverInstanceGroupResponse**](SqlserverInstanceGroupResponse.md) |  | 
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
**VipPublicIpAddress** | Pointer to **NullableString** |  | [optional] 
**VipPublicIpId** | Pointer to **NullableString** |  | [optional] 
**VirtualIpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSqlserverClusterDetailResponse

`func NewSqlserverClusterDetailResponse(accountId string, createdAt time.Time, createdBy string, dbaasEngine string, dbaasEngineVersionName string, id string, initConfigOption SqlserverInitConfigOptionResponse, instanceGroups []SqlserverInstanceGroupResponse, modifiedAt time.Time, modifiedBy string, name string, productImageType string, productType ProductType, roleType NullableClusterRoleType, serviceState ServiceState, softwareVersion string, subnetId string, timezone string, ) *SqlserverClusterDetailResponse`

NewSqlserverClusterDetailResponse instantiates a new SqlserverClusterDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSqlserverClusterDetailResponseWithDefaults

`func NewSqlserverClusterDetailResponseWithDefaults() *SqlserverClusterDetailResponse`

NewSqlserverClusterDetailResponseWithDefaults instantiates a new SqlserverClusterDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SqlserverClusterDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SqlserverClusterDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SqlserverClusterDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAllowableIpAddresses

`func (o *SqlserverClusterDetailResponse) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *SqlserverClusterDetailResponse) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *SqlserverClusterDetailResponse) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *SqlserverClusterDetailResponse) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SqlserverClusterDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SqlserverClusterDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SqlserverClusterDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SqlserverClusterDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SqlserverClusterDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SqlserverClusterDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDbaasEngine

`func (o *SqlserverClusterDetailResponse) GetDbaasEngine() string`

GetDbaasEngine returns the DbaasEngine field if non-nil, zero value otherwise.

### GetDbaasEngineOk

`func (o *SqlserverClusterDetailResponse) GetDbaasEngineOk() (*string, bool)`

GetDbaasEngineOk returns a tuple with the DbaasEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngine

`func (o *SqlserverClusterDetailResponse) SetDbaasEngine(v string)`

SetDbaasEngine sets DbaasEngine field to given value.


### GetDbaasEngineVersionName

`func (o *SqlserverClusterDetailResponse) GetDbaasEngineVersionName() string`

GetDbaasEngineVersionName returns the DbaasEngineVersionName field if non-nil, zero value otherwise.

### GetDbaasEngineVersionNameOk

`func (o *SqlserverClusterDetailResponse) GetDbaasEngineVersionNameOk() (*string, bool)`

GetDbaasEngineVersionNameOk returns a tuple with the DbaasEngineVersionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngineVersionName

`func (o *SqlserverClusterDetailResponse) SetDbaasEngineVersionName(v string)`

SetDbaasEngineVersionName sets DbaasEngineVersionName field to given value.


### GetHaEnabled

`func (o *SqlserverClusterDetailResponse) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *SqlserverClusterDetailResponse) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *SqlserverClusterDetailResponse) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *SqlserverClusterDetailResponse) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetId

`func (o *SqlserverClusterDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SqlserverClusterDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SqlserverClusterDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInitConfigOption

`func (o *SqlserverClusterDetailResponse) GetInitConfigOption() SqlserverInitConfigOptionResponse`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *SqlserverClusterDetailResponse) GetInitConfigOptionOk() (*SqlserverInitConfigOptionResponse, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *SqlserverClusterDetailResponse) SetInitConfigOption(v SqlserverInitConfigOptionResponse)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceCount

`func (o *SqlserverClusterDetailResponse) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *SqlserverClusterDetailResponse) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *SqlserverClusterDetailResponse) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *SqlserverClusterDetailResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceGroups

`func (o *SqlserverClusterDetailResponse) GetInstanceGroups() []SqlserverInstanceGroupResponse`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *SqlserverClusterDetailResponse) GetInstanceGroupsOk() (*[]SqlserverInstanceGroupResponse, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *SqlserverClusterDetailResponse) SetInstanceGroups(v []SqlserverInstanceGroupResponse)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetMaintenanceOption

`func (o *SqlserverClusterDetailResponse) GetMaintenanceOption() MaintenanceResponseOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *SqlserverClusterDetailResponse) GetMaintenanceOptionOk() (*MaintenanceResponseOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *SqlserverClusterDetailResponse) SetMaintenanceOption(v MaintenanceResponseOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *SqlserverClusterDetailResponse) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *SqlserverClusterDetailResponse) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *SqlserverClusterDetailResponse) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetModifiedAt

`func (o *SqlserverClusterDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SqlserverClusterDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SqlserverClusterDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *SqlserverClusterDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SqlserverClusterDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SqlserverClusterDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *SqlserverClusterDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SqlserverClusterDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SqlserverClusterDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNatEnabled

`func (o *SqlserverClusterDetailResponse) GetNatEnabled() bool`

GetNatEnabled returns the NatEnabled field if non-nil, zero value otherwise.

### GetNatEnabledOk

`func (o *SqlserverClusterDetailResponse) GetNatEnabledOk() (*bool, bool)`

GetNatEnabledOk returns a tuple with the NatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatEnabled

`func (o *SqlserverClusterDetailResponse) SetNatEnabled(v bool)`

SetNatEnabled sets NatEnabled field to given value.

### HasNatEnabled

`func (o *SqlserverClusterDetailResponse) HasNatEnabled() bool`

HasNatEnabled returns a boolean if a field has been set.

### GetProductImageType

`func (o *SqlserverClusterDetailResponse) GetProductImageType() string`

GetProductImageType returns the ProductImageType field if non-nil, zero value otherwise.

### GetProductImageTypeOk

`func (o *SqlserverClusterDetailResponse) GetProductImageTypeOk() (*string, bool)`

GetProductImageTypeOk returns a tuple with the ProductImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImageType

`func (o *SqlserverClusterDetailResponse) SetProductImageType(v string)`

SetProductImageType sets ProductImageType field to given value.


### GetProductType

`func (o *SqlserverClusterDetailResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *SqlserverClusterDetailResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *SqlserverClusterDetailResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetRoleType

`func (o *SqlserverClusterDetailResponse) GetRoleType() ClusterRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *SqlserverClusterDetailResponse) GetRoleTypeOk() (*ClusterRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *SqlserverClusterDetailResponse) SetRoleType(v ClusterRoleType)`

SetRoleType sets RoleType field to given value.


### SetRoleTypeNil

`func (o *SqlserverClusterDetailResponse) SetRoleTypeNil(b bool)`

 SetRoleTypeNil sets the value for RoleType to be an explicit nil

### UnsetRoleType
`func (o *SqlserverClusterDetailResponse) UnsetRoleType()`

UnsetRoleType ensures that no value is present for RoleType, not even an explicit nil
### GetServiceState

`func (o *SqlserverClusterDetailResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *SqlserverClusterDetailResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *SqlserverClusterDetailResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.


### GetSoftwareVersion

`func (o *SqlserverClusterDetailResponse) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *SqlserverClusterDetailResponse) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *SqlserverClusterDetailResponse) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetSubnetId

`func (o *SqlserverClusterDetailResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *SqlserverClusterDetailResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *SqlserverClusterDetailResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTimezone

`func (o *SqlserverClusterDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *SqlserverClusterDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *SqlserverClusterDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetVipPublicIpAddress

`func (o *SqlserverClusterDetailResponse) GetVipPublicIpAddress() string`

GetVipPublicIpAddress returns the VipPublicIpAddress field if non-nil, zero value otherwise.

### GetVipPublicIpAddressOk

`func (o *SqlserverClusterDetailResponse) GetVipPublicIpAddressOk() (*string, bool)`

GetVipPublicIpAddressOk returns a tuple with the VipPublicIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPublicIpAddress

`func (o *SqlserverClusterDetailResponse) SetVipPublicIpAddress(v string)`

SetVipPublicIpAddress sets VipPublicIpAddress field to given value.

### HasVipPublicIpAddress

`func (o *SqlserverClusterDetailResponse) HasVipPublicIpAddress() bool`

HasVipPublicIpAddress returns a boolean if a field has been set.

### SetVipPublicIpAddressNil

`func (o *SqlserverClusterDetailResponse) SetVipPublicIpAddressNil(b bool)`

 SetVipPublicIpAddressNil sets the value for VipPublicIpAddress to be an explicit nil

### UnsetVipPublicIpAddress
`func (o *SqlserverClusterDetailResponse) UnsetVipPublicIpAddress()`

UnsetVipPublicIpAddress ensures that no value is present for VipPublicIpAddress, not even an explicit nil
### GetVipPublicIpId

`func (o *SqlserverClusterDetailResponse) GetVipPublicIpId() string`

GetVipPublicIpId returns the VipPublicIpId field if non-nil, zero value otherwise.

### GetVipPublicIpIdOk

`func (o *SqlserverClusterDetailResponse) GetVipPublicIpIdOk() (*string, bool)`

GetVipPublicIpIdOk returns a tuple with the VipPublicIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPublicIpId

`func (o *SqlserverClusterDetailResponse) SetVipPublicIpId(v string)`

SetVipPublicIpId sets VipPublicIpId field to given value.

### HasVipPublicIpId

`func (o *SqlserverClusterDetailResponse) HasVipPublicIpId() bool`

HasVipPublicIpId returns a boolean if a field has been set.

### SetVipPublicIpIdNil

`func (o *SqlserverClusterDetailResponse) SetVipPublicIpIdNil(b bool)`

 SetVipPublicIpIdNil sets the value for VipPublicIpId to be an explicit nil

### UnsetVipPublicIpId
`func (o *SqlserverClusterDetailResponse) UnsetVipPublicIpId()`

UnsetVipPublicIpId ensures that no value is present for VipPublicIpId, not even an explicit nil
### GetVirtualIpAddress

`func (o *SqlserverClusterDetailResponse) GetVirtualIpAddress() string`

GetVirtualIpAddress returns the VirtualIpAddress field if non-nil, zero value otherwise.

### GetVirtualIpAddressOk

`func (o *SqlserverClusterDetailResponse) GetVirtualIpAddressOk() (*string, bool)`

GetVirtualIpAddressOk returns a tuple with the VirtualIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIpAddress

`func (o *SqlserverClusterDetailResponse) SetVirtualIpAddress(v string)`

SetVirtualIpAddress sets VirtualIpAddress field to given value.

### HasVirtualIpAddress

`func (o *SqlserverClusterDetailResponse) HasVirtualIpAddress() bool`

HasVirtualIpAddress returns a boolean if a field has been set.

### SetVirtualIpAddressNil

`func (o *SqlserverClusterDetailResponse) SetVirtualIpAddressNil(b bool)`

 SetVirtualIpAddressNil sets the value for VirtualIpAddress to be an explicit nil

### UnsetVirtualIpAddress
`func (o *SqlserverClusterDetailResponse) UnsetVirtualIpAddress()`

UnsetVirtualIpAddress ensures that no value is present for VirtualIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


