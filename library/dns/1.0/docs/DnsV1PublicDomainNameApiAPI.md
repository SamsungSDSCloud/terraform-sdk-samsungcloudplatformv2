# \DnsV1PublicDomainNameApiAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckDuplicationPublicDomain**](DnsV1PublicDomainNameApiAPI.md#CheckDuplicationPublicDomain) | **Get** /v1/public-domain-names/check-duplication | Check if public domain name is duplicated
[**CreatePublicDomain**](DnsV1PublicDomainNameApiAPI.md#CreatePublicDomain) | **Post** /v1/public-domain-names | Create Public Domain
[**GetAddressInfo**](DnsV1PublicDomainNameApiAPI.md#GetAddressInfo) | **Get** /v1/public-domain-names/address | Get Address Info by keyword
[**GetAllowedTldList**](DnsV1PublicDomainNameApiAPI.md#GetAllowedTldList) | **Get** /v1/public-domain-names/tld-list | Get Allowed TLD List
[**GetPublicDomainDetail**](DnsV1PublicDomainNameApiAPI.md#GetPublicDomainDetail) | **Get** /v1/public-domain-names/{public_domain_id} | Get Public Domain Detail
[**ListPublicDomains**](DnsV1PublicDomainNameApiAPI.md#ListPublicDomains) | **Get** /v1/public-domain-names | List Public Domains
[**PutPublicDomain**](DnsV1PublicDomainNameApiAPI.md#PutPublicDomain) | **Put** /v1/public-domain-names/{public_domain_id} | Update auto_extension or description
[**UpdateWhoisInfoPublicDomain**](DnsV1PublicDomainNameApiAPI.md#UpdateWhoisInfoPublicDomain) | **Put** /v1/public-domain-names/{public_domain_id}/information | Update WHOIS info for Public Domain



## CheckDuplicationPublicDomain

> PublicDomainCheckDuplicationResponse CheckDuplicationPublicDomain(ctx).Name(name).Execute()

Check if public domain name is duplicated



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
	name := "name_example" // string | name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PublicDomainNameApiAPI.CheckDuplicationPublicDomain(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.CheckDuplicationPublicDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckDuplicationPublicDomain`: PublicDomainCheckDuplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PublicDomainNameApiAPI.CheckDuplicationPublicDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDuplicationPublicDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | name | 

### Return type

[**PublicDomainCheckDuplicationResponse**](PublicDomainCheckDuplicationResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePublicDomain

> CreatePublicDomainResponse CreatePublicDomain(ctx).CreatePublicDomainRequest(createPublicDomainRequest).Execute()

Create Public Domain



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
	createPublicDomainRequest := *openapiclient.NewCreatePublicDomainRequest("AddressType_example", "DomesticFirstAddressEn_example", "DomesticFirstAddressKo_example", "DomesticSecondAddressEn_example", "DomesticSecondAddressKo_example", "Name_example", "OverseasFirstAddress_example", "OverseasSecondAddress_example", "OverseasThirdAddress_example", "PostalCode_example", "RegisterEmail_example", "RegisterNameEn_example", "RegisterNameKo_example", "RegisterTelno_example") // CreatePublicDomainRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PublicDomainNameApiAPI.CreatePublicDomain(context.Background()).CreatePublicDomainRequest(createPublicDomainRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.CreatePublicDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePublicDomain`: CreatePublicDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PublicDomainNameApiAPI.CreatePublicDomain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePublicDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPublicDomainRequest** | [**CreatePublicDomainRequest**](CreatePublicDomainRequest.md) |  | 

### Return type

[**CreatePublicDomainResponse**](CreatePublicDomainResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddressInfo

> GetAddressInfo(ctx).Keyword(keyword).Size(size).Page(page).Execute()

Get Address Info by keyword



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
	keyword := "keyword_example" // string | Keyword for address search
	size := int32(56) // int32 | Number of results per page (optional)
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsV1PublicDomainNameApiAPI.GetAddressInfo(context.Background()).Keyword(keyword).Size(size).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.GetAddressInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keyword** | **string** | Keyword for address search | 
 **size** | **int32** | Number of results per page | 
 **page** | **int32** | Page number | 

### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllowedTldList

> GetAllowedTldList(ctx).Execute()

Get Allowed TLD List



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DnsV1PublicDomainNameApiAPI.GetAllowedTldList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.GetAllowedTldList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowedTldListRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicDomainDetail

> PublicDomainDetailResponse GetPublicDomainDetail(ctx, publicDomainId).Execute()

Get Public Domain Detail



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
	publicDomainId := "publicDomainId_example" // string | Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PublicDomainNameApiAPI.GetPublicDomainDetail(context.Background(), publicDomainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.GetPublicDomainDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicDomainDetail`: PublicDomainDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PublicDomainNameApiAPI.GetPublicDomainDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicDomainId** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicDomainDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicDomainDetailResponse**](PublicDomainDetailResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPublicDomains

> PublicDomainListResponse ListPublicDomains(ctx).Size(size).Page(page).Sort(sort).Name(name).CreatedBy(createdBy).Execute()

List Public Domains



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
	size := int32(20) // int32 | size (optional)
	page := int32(0) // int32 | page (optional)
	sort := "created_at:desc" // string | sort (optional)
	name := "name_example" // string | name (optional)
	createdBy := "createdBy_example" // string | Created By_ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PublicDomainNameApiAPI.ListPublicDomains(context.Background()).Size(size).Page(page).Sort(sort).Name(name).CreatedBy(createdBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.ListPublicDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPublicDomains`: PublicDomainListResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PublicDomainNameApiAPI.ListPublicDomains`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPublicDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **size** | **int32** | size | 
 **page** | **int32** | page | 
 **sort** | **string** | sort | 
 **name** | **string** | name | 
 **createdBy** | **string** | Created By_ | 

### Return type

[**PublicDomainListResponse**](PublicDomainListResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPublicDomain

> PublicDomainPartialUpdateResponse PutPublicDomain(ctx, publicDomainId).PublicDomainPartialUpdateRequest(publicDomainPartialUpdateRequest).Execute()

Update auto_extension or description



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
	publicDomainId := "publicDomainId_example" // string | Id
	publicDomainPartialUpdateRequest := *openapiclient.NewPublicDomainPartialUpdateRequest() // PublicDomainPartialUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PublicDomainNameApiAPI.PutPublicDomain(context.Background(), publicDomainId).PublicDomainPartialUpdateRequest(publicDomainPartialUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.PutPublicDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPublicDomain`: PublicDomainPartialUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PublicDomainNameApiAPI.PutPublicDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicDomainId** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPublicDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicDomainPartialUpdateRequest** | [**PublicDomainPartialUpdateRequest**](PublicDomainPartialUpdateRequest.md) |  | 

### Return type

[**PublicDomainPartialUpdateResponse**](PublicDomainPartialUpdateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWhoisInfoPublicDomain

> PubblicDomainWhoisInfoUpdateResponse UpdateWhoisInfoPublicDomain(ctx, publicDomainId).PubblicDomainWhoisInfoUpdateRequest(pubblicDomainWhoisInfoUpdateRequest).Execute()

Update WHOIS info for Public Domain



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
	publicDomainId := "publicDomainId_example" // string | Id
	pubblicDomainWhoisInfoUpdateRequest := *openapiclient.NewPubblicDomainWhoisInfoUpdateRequest("AddressType_example", "DomesticFirstAddressEn_example", "DomesticFirstAddressKo_example", "DomesticSecondAddressEn_example", "DomesticSecondAddressKo_example", "OverseasFirstAddress_example", "OverseasSecondAddress_example", "OverseasThirdAddress_example", "PostalCode_example", "RegisterEmail_example", "RegisterTelno_example") // PubblicDomainWhoisInfoUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DnsV1PublicDomainNameApiAPI.UpdateWhoisInfoPublicDomain(context.Background(), publicDomainId).PubblicDomainWhoisInfoUpdateRequest(pubblicDomainWhoisInfoUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DnsV1PublicDomainNameApiAPI.UpdateWhoisInfoPublicDomain``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWhoisInfoPublicDomain`: PubblicDomainWhoisInfoUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `DnsV1PublicDomainNameApiAPI.UpdateWhoisInfoPublicDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicDomainId** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWhoisInfoPublicDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pubblicDomainWhoisInfoUpdateRequest** | [**PubblicDomainWhoisInfoUpdateRequest**](PubblicDomainWhoisInfoUpdateRequest.md) |  | 

### Return type

[**PubblicDomainWhoisInfoUpdateResponse**](PubblicDomainWhoisInfoUpdateResponse.md)

### Authorization

[X-Auth-Token](../README.md#X-Auth-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

