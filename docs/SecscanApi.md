# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRepoImageSecurity**](SecscanApi.md#GetRepoImageSecurity) | **Get** /api/v1/repository/{repository}/image/{imageid}/security | 
[**GetRepoManifestSecurity**](SecscanApi.md#GetRepoManifestSecurity) | **Get** /api/v1/repository/{repository}/manifest/{manifestref}/security | 

# **GetRepoImageSecurity**
> GetRepoImageSecurity(ctx, repository, imageid, optional)


Fetches the features and vulnerabilities (if any) for a repository image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
  **imageid** | **string**| The image ID | 
 **optional** | ***SecscanApiGetRepoImageSecurityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecscanApiGetRepoImageSecurityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vulnerabilities** | **optional.Bool**| Include vulnerabilities information | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoManifestSecurity**
> GetRepoManifestSecurity(ctx, manifestref, repository, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **manifestref** | **string**| The digest of the manifest | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
 **optional** | ***SecscanApiGetRepoManifestSecurityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecscanApiGetRepoManifestSecurityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vulnerabilities** | **optional.Bool**| Include vulnerabilities informations | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

