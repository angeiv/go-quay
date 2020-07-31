# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAppToken**](AppspecifictokensApi.md#CreateAppToken) | **Post** /api/v1/user/apptoken | 
[**GetAppToken**](AppspecifictokensApi.md#GetAppToken) | **Get** /api/v1/user/apptoken/{token_uuid} | 
[**ListAppTokens**](AppspecifictokensApi.md#ListAppTokens) | **Get** /api/v1/user/apptoken | 
[**RevokeAppToken**](AppspecifictokensApi.md#RevokeAppToken) | **Delete** /api/v1/user/apptoken/{token_uuid} | 

# **CreateAppToken**
> CreateAppToken(ctx, body)


Create a new app specific token for user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewToken**](NewToken.md)| Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAppToken**
> GetAppToken(ctx, tokenUuid)


Returns a specific app token for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenUuid** | **string**| The uuid of the app specific token | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAppTokens**
> ListAppTokens(ctx, optional)


Lists the app specific tokens for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppspecifictokensApiListAppTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppspecifictokensApiListAppTokensOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expiring** | **optional.Bool**| If true, only returns those tokens expiring soon | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeAppToken**
> RevokeAppToken(ctx, tokenUuid)


Revokes a specific app token for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenUuid** | **string**| The uuid of the app specific token | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

