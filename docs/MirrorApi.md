# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeRepoMirrorConfig**](MirrorApi.md#ChangeRepoMirrorConfig) | **Put** /api/v1/repository/{repository}/mirror | 
[**CreateRepoMirrorConfig**](MirrorApi.md#CreateRepoMirrorConfig) | **Post** /api/v1/repository/{repository}/mirror | 
[**GetRepoMirrorConfig**](MirrorApi.md#GetRepoMirrorConfig) | **Get** /api/v1/repository/{repository}/mirror | 
[**SyncCancel**](MirrorApi.md#SyncCancel) | **Post** /api/v1/repository/{repository}/mirror/sync-cancel | 
[**SyncNow**](MirrorApi.md#SyncNow) | **Post** /api/v1/repository/{repository}/mirror/sync-now | 

# **ChangeRepoMirrorConfig**
> ChangeRepoMirrorConfig(ctx, body, repository)


Allow users to modifying the repository's mirroring configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateMirrorConfig**](UpdateMirrorConfig.md)| Request body contents. | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepoMirrorConfig**
> CreateRepoMirrorConfig(ctx, body, repository)


Create a RepoMirrorConfig for a given Repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMirrorConfig**](CreateMirrorConfig.md)| Request body contents. | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoMirrorConfig**
> ViewMirrorConfig GetRepoMirrorConfig(ctx, repository)


Return the Mirror configuration for a given Repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

[**ViewMirrorConfig**](ViewMirrorConfig.md)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SyncCancel**
> SyncCancel(ctx, repository)


Update the sync_status for a given Repository's mirroring configuration.

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

# **SyncNow**
> SyncNow(ctx, repository)


Update the sync_status for a given Repository's mirroring configuration.

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

