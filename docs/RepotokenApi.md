# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeToken**](RepotokenApi.md#ChangeToken) | **Put** /api/v1/repository/{repository}/tokens/{code} | 
[**CreateToken**](RepotokenApi.md#CreateToken) | **Post** /api/v1/repository/{repository}/tokens/ | 
[**DeleteToken**](RepotokenApi.md#DeleteToken) | **Delete** /api/v1/repository/{repository}/tokens/{code} | 
[**GetTokens**](RepotokenApi.md#GetTokens) | **Get** /api/v1/repository/{repository}/tokens/{code} | 
[**ListRepoTokens**](RepotokenApi.md#ListRepoTokens) | **Get** /api/v1/repository/{repository}/tokens/ | 

# **ChangeToken**
> ChangeToken(ctx, body, code, repository)


Update the permissions for the specified repository token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TokenPermission**](TokenPermission.md)| Request body contents. | 
  **code** | **string**| The token code | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateToken**
> CreateToken(ctx, body, repository)


Create a new repository token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewToken**](NewToken.md)| Request body contents. | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteToken**
> DeleteToken(ctx, code, repository)


Delete the repository token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The token code | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTokens**
> GetTokens(ctx, code, repository)


Fetch the specified repository token information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| The token code | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRepoTokens**
> ListRepoTokens(ctx, repository)


List the tokens for the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

