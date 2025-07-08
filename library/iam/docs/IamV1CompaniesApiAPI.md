# \IamV1CompaniesApiAPI

All URIs are relative to *https://iam.kr-west1.dev2.samsungsdscloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCompanies**](IamV1CompaniesApiAPI.md#ListCompanies) | **Get** /v1/companies | Get list of companies registered in MDG



## ListCompanies

> ListCompanyResponse ListCompanies(ctx).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).RegistrationNumber(registrationNumber).Name(name).Id(id).Ids(ids).Execute()

Get list of companies registered in MDG



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
	registrationNumber := "registrationNumber_example" // string | 사업자 등록 번호 (optional)
	name := "name_example" // string | 회사이름 (optional)
	id := "id_example" // string | ID (optional)
	ids := "ids_example" // string | IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IamV1CompaniesApiAPI.ListCompanies(context.Background()).WithCount(withCount).Limit(limit).Marker(marker).Sort(sort).RegistrationNumber(registrationNumber).Name(name).Id(id).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IamV1CompaniesApiAPI.ListCompanies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCompanies`: ListCompanyResponse
	fmt.Fprintf(os.Stdout, "Response from `IamV1CompaniesApiAPI.ListCompanies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCompaniesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withCount** | **string** | with count | 
 **limit** | **int32** | limit | 
 **marker** | **string** | marker | 
 **sort** | **string** | sort | 
 **registrationNumber** | **string** | 사업자 등록 번호 | 
 **name** | **string** | 회사이름 | 
 **id** | **string** | ID | 
 **ids** | **string** | IDs | 

### Return type

[**ListCompanyResponse**](ListCompanyResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

