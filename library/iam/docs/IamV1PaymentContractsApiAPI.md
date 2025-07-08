# \IamV1PaymentContractsApiAPI

All URIs are relative to *https://iam.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentContractList**](IamV1PaymentContractsApiAPI.md#PaymentContractList) | **Get** /v1/payment-contracts | List Payment contracts
[**PaymentContractShow**](IamV1PaymentContractsApiAPI.md#PaymentContractShow) | **Get** /v1/payment-contracts/{payment_contract_id} | Show Payment contract



## PaymentContractList

> PaymentContractListResponse PaymentContractList(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).BusinessCategoryIds(businessCategoryIds).StartCreatedAt(startCreatedAt).EndCreatedAt(endCreatedAt).IncludeAccountOfProject(includeAccountOfProject).CreatorName(creatorName).CreatorEmail(creatorEmail).CreatorCompanyName(creatorCompanyName).ModifierName(modifierName).ModifierEmail(modifierEmail).ModifierCompanyName(modifierCompanyName).RequestScope(requestScope).AccountManagerId(accountManagerId).AccountUserId(accountUserId).Execute()

List Payment contracts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	withCount := "true" // string | with count (optional)
	limit := int32(20) // int32 | limit (optional)
	marker := "607e0938521643b5b4b266f343fae693" // string | marker (optional)
	sort := "created_at:desc" // string | sort (optional)
	id := "id_example" // string | Account ID (optional)
	name := "name_example" // string | Account 이름 (optional)
	businessCategoryIds := "businessCategoryIds_example" // string | 회사 업무 구분 ID 목록 (optional)
	startCreatedAt := "startCreatedAt_example" // string | 시작 생성 일시 (optional)
	endCreatedAt := "endCreatedAt_example" // string | 끝 생성 일시 (optional)
	includeAccountOfProject := "includeAccountOfProject_example" // string | Project의 Account 포함 여부 (optional)
	creatorName := "creatorName_example" // string | Account 생성자 성, 이름 (optional)
	creatorEmail := "creatorEmail_example" // string | Account 생성자 Email (optional)
	creatorCompanyName := "creatorCompanyName_example" // string | Account 생성자 회사 이름 (optional)
	modifierName := "modifierName_example" // string | Account 수정자 성, 이름 (optional)
	modifierEmail := "modifierEmail_example" // string | Account 수정자 Email (optional)
	modifierCompanyName := "modifierCompanyName_example" // string | Account 수정자 회사 이름 (optional)
	requestScope := openapiclient.AccountRequestScopeEnum("ADMIN") // AccountRequestScopeEnum | 요청 범위 (optional)
	accountManagerId := "accountManagerId_example" // string | Account Manager ID (optional)
	accountUserId := "accountUserId_example" // string | Account User ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PaymentContractsApiAPI.PaymentContractList(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).Id(id).Name(name).BusinessCategoryIds(businessCategoryIds).StartCreatedAt(startCreatedAt).EndCreatedAt(endCreatedAt).IncludeAccountOfProject(includeAccountOfProject).CreatorName(creatorName).CreatorEmail(creatorEmail).CreatorCompanyName(creatorCompanyName).ModifierName(modifierName).ModifierEmail(modifierEmail).ModifierCompanyName(modifierCompanyName).RequestScope(requestScope).AccountManagerId(accountManagerId).AccountUserId(accountUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PaymentContractsApiAPI.PaymentContractList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentContractList`: PaymentContractListResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PaymentContractsApiAPI.PaymentContractList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentContractListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **id** | **string** | Account ID | 
 **name** | **string** | Account 이름 | 
 **businessCategoryIds** | **string** | 회사 업무 구분 ID 목록 | 
 **startCreatedAt** | **string** | 시작 생성 일시 | 
 **endCreatedAt** | **string** | 끝 생성 일시 | 
 **includeAccountOfProject** | **string** | Project의 Account 포함 여부 | 
 **creatorName** | **string** | Account 생성자 성, 이름 | 
 **creatorEmail** | **string** | Account 생성자 Email | 
 **creatorCompanyName** | **string** | Account 생성자 회사 이름 | 
 **modifierName** | **string** | Account 수정자 성, 이름 | 
 **modifierEmail** | **string** | Account 수정자 Email | 
 **modifierCompanyName** | **string** | Account 수정자 회사 이름 | 
 **requestScope** | [**AccountRequestScopeEnum**](AccountRequestScopeEnum.md) | 요청 범위 | 
 **accountManagerId** | **string** | Account Manager ID | 
 **accountUserId** | **string** | Account User ID | 

### Return type

[**PaymentContractListResponse**](PaymentContractListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentContractShow

> PaymentContractShowResponse PaymentContractShow(ctx, paymentContractId).Execute()

Show Payment contract



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	paymentContractId := "paymentContractId_example" // string | Payment contract ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1PaymentContractsApiAPI.PaymentContractShow(context.Background(), paymentContractId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1PaymentContractsApiAPI.PaymentContractShow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentContractShow`: PaymentContractShowResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1PaymentContractsApiAPI.PaymentContractShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentContractId** | **string** | Payment contract ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentContractShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaymentContractShowResponse**](PaymentContractShowResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

