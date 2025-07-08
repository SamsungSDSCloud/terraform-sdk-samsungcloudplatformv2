# MariadbClusterDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Account ID | 
**AllowableIpAddresses** | Pointer to **[]string** | Allowed IP addresses list | [optional] [default to []]
**CreatedAt** | **time.Time** | Created At | 
**CreatedBy** | **string** | Created By | 
**DbaasEngine** | **string** | DBaaS engine | 
**HaEnabled** | Pointer to **bool** | HA availability | [optional] [default to false]
**Id** | **string** | ID | 
**InitConfigOption** | [**MariadbInitConfigOptionDetail**](MariadbInitConfigOptionDetail.md) |  | 
**InstanceCount** | Pointer to **int32** | Instance Count | [optional] [default to 0]
**InstanceGroups** | [**[]InstanceGroupResponse**](InstanceGroupResponse.md) | Instance groups list | 
**MaintenanceOption** | Pointer to [**NullableMaintenanceResponseOption**](MaintenanceResponseOption.md) |  | [optional] 
**ModifiedAt** | **time.Time** | Modified At | 
**ModifiedBy** | **string** | Modified By | 
**Name** | **string** | Cluster name | 
**OriginClusterId** | Pointer to **NullableString** |  | [optional] 
**ProductType** | [**ProductType**](ProductType.md) | Product type | 
**Replicas** | Pointer to **[]string** |  | [optional] 
**RoleType** | [**NullableClusterRoleType**](ClusterRoleType.md) |  | 
**ServiceState** | [**ServiceState**](ServiceState.md) | Service state | 
**SoftwareVersion** | **string** | Software version | 
**SubnetId** | **string** | Subnet ID | 
**Timezone** | **string** | Timezone | 
**VirtualIpAddress** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMariadbClusterDetailResponse

`func NewMariadbClusterDetailResponse(accountId string, createdAt time.Time, createdBy string, dbaasEngine string, id string, initConfigOption MariadbInitConfigOptionDetail, instanceGroups []InstanceGroupResponse, modifiedAt time.Time, modifiedBy string, name string, productType ProductType, roleType NullableClusterRoleType, serviceState ServiceState, softwareVersion string, subnetId string, timezone string, ) *MariadbClusterDetailResponse`

NewMariadbClusterDetailResponse instantiates a new MariadbClusterDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbClusterDetailResponseWithDefaults

`func NewMariadbClusterDetailResponseWithDefaults() *MariadbClusterDetailResponse`

NewMariadbClusterDetailResponseWithDefaults instantiates a new MariadbClusterDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MariadbClusterDetailResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MariadbClusterDetailResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MariadbClusterDetailResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAllowableIpAddresses

`func (o *MariadbClusterDetailResponse) GetAllowableIpAddresses() []string`

GetAllowableIpAddresses returns the AllowableIpAddresses field if non-nil, zero value otherwise.

### GetAllowableIpAddressesOk

`func (o *MariadbClusterDetailResponse) GetAllowableIpAddressesOk() (*[]string, bool)`

GetAllowableIpAddressesOk returns a tuple with the AllowableIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowableIpAddresses

`func (o *MariadbClusterDetailResponse) SetAllowableIpAddresses(v []string)`

SetAllowableIpAddresses sets AllowableIpAddresses field to given value.

### HasAllowableIpAddresses

`func (o *MariadbClusterDetailResponse) HasAllowableIpAddresses() bool`

HasAllowableIpAddresses returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MariadbClusterDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MariadbClusterDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MariadbClusterDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *MariadbClusterDetailResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MariadbClusterDetailResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MariadbClusterDetailResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetDbaasEngine

`func (o *MariadbClusterDetailResponse) GetDbaasEngine() string`

GetDbaasEngine returns the DbaasEngine field if non-nil, zero value otherwise.

### GetDbaasEngineOk

`func (o *MariadbClusterDetailResponse) GetDbaasEngineOk() (*string, bool)`

GetDbaasEngineOk returns a tuple with the DbaasEngine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaasEngine

`func (o *MariadbClusterDetailResponse) SetDbaasEngine(v string)`

SetDbaasEngine sets DbaasEngine field to given value.


### GetHaEnabled

`func (o *MariadbClusterDetailResponse) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *MariadbClusterDetailResponse) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *MariadbClusterDetailResponse) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.

### HasHaEnabled

`func (o *MariadbClusterDetailResponse) HasHaEnabled() bool`

HasHaEnabled returns a boolean if a field has been set.

### GetId

`func (o *MariadbClusterDetailResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MariadbClusterDetailResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MariadbClusterDetailResponse) SetId(v string)`

SetId sets Id field to given value.


### GetInitConfigOption

`func (o *MariadbClusterDetailResponse) GetInitConfigOption() MariadbInitConfigOptionDetail`

GetInitConfigOption returns the InitConfigOption field if non-nil, zero value otherwise.

### GetInitConfigOptionOk

`func (o *MariadbClusterDetailResponse) GetInitConfigOptionOk() (*MariadbInitConfigOptionDetail, bool)`

GetInitConfigOptionOk returns a tuple with the InitConfigOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitConfigOption

`func (o *MariadbClusterDetailResponse) SetInitConfigOption(v MariadbInitConfigOptionDetail)`

SetInitConfigOption sets InitConfigOption field to given value.


### GetInstanceCount

`func (o *MariadbClusterDetailResponse) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *MariadbClusterDetailResponse) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *MariadbClusterDetailResponse) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *MariadbClusterDetailResponse) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstanceGroups

`func (o *MariadbClusterDetailResponse) GetInstanceGroups() []InstanceGroupResponse`

GetInstanceGroups returns the InstanceGroups field if non-nil, zero value otherwise.

### GetInstanceGroupsOk

`func (o *MariadbClusterDetailResponse) GetInstanceGroupsOk() (*[]InstanceGroupResponse, bool)`

GetInstanceGroupsOk returns a tuple with the InstanceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroups

`func (o *MariadbClusterDetailResponse) SetInstanceGroups(v []InstanceGroupResponse)`

SetInstanceGroups sets InstanceGroups field to given value.


### GetMaintenanceOption

`func (o *MariadbClusterDetailResponse) GetMaintenanceOption() MaintenanceResponseOption`

GetMaintenanceOption returns the MaintenanceOption field if non-nil, zero value otherwise.

### GetMaintenanceOptionOk

`func (o *MariadbClusterDetailResponse) GetMaintenanceOptionOk() (*MaintenanceResponseOption, bool)`

GetMaintenanceOptionOk returns a tuple with the MaintenanceOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceOption

`func (o *MariadbClusterDetailResponse) SetMaintenanceOption(v MaintenanceResponseOption)`

SetMaintenanceOption sets MaintenanceOption field to given value.

### HasMaintenanceOption

`func (o *MariadbClusterDetailResponse) HasMaintenanceOption() bool`

HasMaintenanceOption returns a boolean if a field has been set.

### SetMaintenanceOptionNil

`func (o *MariadbClusterDetailResponse) SetMaintenanceOptionNil(b bool)`

 SetMaintenanceOptionNil sets the value for MaintenanceOption to be an explicit nil

### UnsetMaintenanceOption
`func (o *MariadbClusterDetailResponse) UnsetMaintenanceOption()`

UnsetMaintenanceOption ensures that no value is present for MaintenanceOption, not even an explicit nil
### GetModifiedAt

`func (o *MariadbClusterDetailResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MariadbClusterDetailResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MariadbClusterDetailResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *MariadbClusterDetailResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *MariadbClusterDetailResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *MariadbClusterDetailResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetName

`func (o *MariadbClusterDetailResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MariadbClusterDetailResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MariadbClusterDetailResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOriginClusterId

`func (o *MariadbClusterDetailResponse) GetOriginClusterId() string`

GetOriginClusterId returns the OriginClusterId field if non-nil, zero value otherwise.

### GetOriginClusterIdOk

`func (o *MariadbClusterDetailResponse) GetOriginClusterIdOk() (*string, bool)`

GetOriginClusterIdOk returns a tuple with the OriginClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginClusterId

`func (o *MariadbClusterDetailResponse) SetOriginClusterId(v string)`

SetOriginClusterId sets OriginClusterId field to given value.

### HasOriginClusterId

`func (o *MariadbClusterDetailResponse) HasOriginClusterId() bool`

HasOriginClusterId returns a boolean if a field has been set.

### SetOriginClusterIdNil

`func (o *MariadbClusterDetailResponse) SetOriginClusterIdNil(b bool)`

 SetOriginClusterIdNil sets the value for OriginClusterId to be an explicit nil

### UnsetOriginClusterId
`func (o *MariadbClusterDetailResponse) UnsetOriginClusterId()`

UnsetOriginClusterId ensures that no value is present for OriginClusterId, not even an explicit nil
### GetProductType

`func (o *MariadbClusterDetailResponse) GetProductType() ProductType`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *MariadbClusterDetailResponse) GetProductTypeOk() (*ProductType, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *MariadbClusterDetailResponse) SetProductType(v ProductType)`

SetProductType sets ProductType field to given value.


### GetReplicas

`func (o *MariadbClusterDetailResponse) GetReplicas() []string`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *MariadbClusterDetailResponse) GetReplicasOk() (*[]string, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *MariadbClusterDetailResponse) SetReplicas(v []string)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *MariadbClusterDetailResponse) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### SetReplicasNil

`func (o *MariadbClusterDetailResponse) SetReplicasNil(b bool)`

 SetReplicasNil sets the value for Replicas to be an explicit nil

### UnsetReplicas
`func (o *MariadbClusterDetailResponse) UnsetReplicas()`

UnsetReplicas ensures that no value is present for Replicas, not even an explicit nil
### GetRoleType

`func (o *MariadbClusterDetailResponse) GetRoleType() ClusterRoleType`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *MariadbClusterDetailResponse) GetRoleTypeOk() (*ClusterRoleType, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *MariadbClusterDetailResponse) SetRoleType(v ClusterRoleType)`

SetRoleType sets RoleType field to given value.


### SetRoleTypeNil

`func (o *MariadbClusterDetailResponse) SetRoleTypeNil(b bool)`

 SetRoleTypeNil sets the value for RoleType to be an explicit nil

### UnsetRoleType
`func (o *MariadbClusterDetailResponse) UnsetRoleType()`

UnsetRoleType ensures that no value is present for RoleType, not even an explicit nil
### GetServiceState

`func (o *MariadbClusterDetailResponse) GetServiceState() ServiceState`

GetServiceState returns the ServiceState field if non-nil, zero value otherwise.

### GetServiceStateOk

`func (o *MariadbClusterDetailResponse) GetServiceStateOk() (*ServiceState, bool)`

GetServiceStateOk returns a tuple with the ServiceState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceState

`func (o *MariadbClusterDetailResponse) SetServiceState(v ServiceState)`

SetServiceState sets ServiceState field to given value.


### GetSoftwareVersion

`func (o *MariadbClusterDetailResponse) GetSoftwareVersion() string`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *MariadbClusterDetailResponse) GetSoftwareVersionOk() (*string, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *MariadbClusterDetailResponse) SetSoftwareVersion(v string)`

SetSoftwareVersion sets SoftwareVersion field to given value.


### GetSubnetId

`func (o *MariadbClusterDetailResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *MariadbClusterDetailResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *MariadbClusterDetailResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.


### GetTimezone

`func (o *MariadbClusterDetailResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MariadbClusterDetailResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MariadbClusterDetailResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetVirtualIpAddress

`func (o *MariadbClusterDetailResponse) GetVirtualIpAddress() string`

GetVirtualIpAddress returns the VirtualIpAddress field if non-nil, zero value otherwise.

### GetVirtualIpAddressOk

`func (o *MariadbClusterDetailResponse) GetVirtualIpAddressOk() (*string, bool)`

GetVirtualIpAddressOk returns a tuple with the VirtualIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIpAddress

`func (o *MariadbClusterDetailResponse) SetVirtualIpAddress(v string)`

SetVirtualIpAddress sets VirtualIpAddress field to given value.

### HasVirtualIpAddress

`func (o *MariadbClusterDetailResponse) HasVirtualIpAddress() bool`

HasVirtualIpAddress returns a boolean if a field has been set.

### SetVirtualIpAddressNil

`func (o *MariadbClusterDetailResponse) SetVirtualIpAddressNil(b bool)`

 SetVirtualIpAddressNil sets the value for VirtualIpAddress to be an explicit nil

### UnsetVirtualIpAddress
`func (o *MariadbClusterDetailResponse) UnsetVirtualIpAddress()`

UnsetVirtualIpAddress ensures that no value is present for VirtualIpAddress, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


