# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateBuildTrigger**](TriggerApi.md#ActivateBuildTrigger) | **Post** /api/v1/repository/{repository}/trigger/{trigger_uuid}/activate | 
[**DeleteBuildTrigger**](TriggerApi.md#DeleteBuildTrigger) | **Delete** /api/v1/repository/{repository}/trigger/{trigger_uuid} | 
[**GetBuildTrigger**](TriggerApi.md#GetBuildTrigger) | **Get** /api/v1/repository/{repository}/trigger/{trigger_uuid} | 
[**ListBuildTriggers**](TriggerApi.md#ListBuildTriggers) | **Get** /api/v1/repository/{repository}/trigger/ | 
[**ListTriggerRecentBuilds**](TriggerApi.md#ListTriggerRecentBuilds) | **Get** /api/v1/repository/{repository}/trigger/{trigger_uuid}/builds | 
[**ManuallyStartBuildTrigger**](TriggerApi.md#ManuallyStartBuildTrigger) | **Post** /api/v1/repository/{repository}/trigger/{trigger_uuid}/start | 
[**UpdateBuildTrigger**](TriggerApi.md#UpdateBuildTrigger) | **Put** /api/v1/repository/{repository}/trigger/{trigger_uuid} | 

# **ActivateBuildTrigger**
> ActivateBuildTrigger(ctx, body, triggerUuid, repository)


Activate the specified build trigger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BuildTriggerActivateRequest**](BuildTriggerActivateRequest.md)| Request body contents. | 
  **triggerUuid** | **string**| The UUID of the build trigger | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBuildTrigger**
> DeleteBuildTrigger(ctx, triggerUuid, repository)


Delete the specified build trigger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerUuid** | **string**| The UUID of the build trigger | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBuildTrigger**
> GetBuildTrigger(ctx, triggerUuid, repository)


Get information for the specified build trigger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerUuid** | **string**| The UUID of the build trigger | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBuildTriggers**
> ListBuildTriggers(ctx, repository)


List the triggers for the specified repository.

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

# **ListTriggerRecentBuilds**
> ListTriggerRecentBuilds(ctx, triggerUuid, repository, optional)


List the builds started by the specified trigger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **triggerUuid** | **string**| The UUID of the build trigger | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
 **optional** | ***TriggerApiListTriggerRecentBuildsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TriggerApiListTriggerRecentBuildsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| The maximum number of builds to return | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManuallyStartBuildTrigger**
> ManuallyStartBuildTrigger(ctx, body, triggerUuid, repository)


Manually start a build from the specified trigger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RunParameters**](RunParameters.md)| Request body contents. | 
  **triggerUuid** | **string**| The UUID of the build trigger | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBuildTrigger**
> UpdateBuildTrigger(ctx, body, triggerUuid, repository)


Updates the specified build trigger.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTrigger**](UpdateTrigger.md)| Request body contents. | 
  **triggerUuid** | **string**| The UUID of the build trigger | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

