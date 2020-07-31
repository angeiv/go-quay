# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddManifestLabel**](ManifestApi.md#AddManifestLabel) | **Post** /api/v1/repository/{repository}/manifest/{manifestref}/labels | 
[**DeleteManifestLabel**](ManifestApi.md#DeleteManifestLabel) | **Delete** /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid} | 
[**GetManifestLabel**](ManifestApi.md#GetManifestLabel) | **Get** /api/v1/repository/{repository}/manifest/{manifestref}/labels/{labelid} | 
[**GetRepoManifest**](ManifestApi.md#GetRepoManifest) | **Get** /api/v1/repository/{repository}/manifest/{manifestref} | 
[**ListManifestLabels**](ManifestApi.md#ListManifestLabels) | **Get** /api/v1/repository/{repository}/manifest/{manifestref}/labels | 

# **AddManifestLabel**
> AddManifestLabel(ctx, body, manifestref, repository)


Adds a new label into the tag manifest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddLabel**](AddLabel.md)| Request body contents. | 
  **manifestref** | **string**| The digest of the manifest | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteManifestLabel**
> DeleteManifestLabel(ctx, labelid, manifestref, repository)


Deletes an existing label from a manifest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelid** | **string**| The ID of the label | 
  **manifestref** | **string**| The digest of the manifest | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetManifestLabel**
> GetManifestLabel(ctx, labelid, manifestref, repository)


Retrieves the label with the specific ID under the manifest.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **labelid** | **string**| The ID of the label | 
  **manifestref** | **string**| The digest of the manifest | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoManifest**
> GetRepoManifest(ctx, manifestref, repository)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **manifestref** | **string**| The digest of the manifest | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListManifestLabels**
> ListManifestLabels(ctx, manifestref, repository, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **manifestref** | **string**| The digest of the manifest | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
 **optional** | ***ManifestApiListManifestLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManifestApiListManifestLabelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| If specified, only labels matching the given prefix will be returned | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

