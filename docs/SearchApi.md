# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConductRepoSearch**](SearchApi.md#ConductRepoSearch) | **Get** /api/v1/find/repositories | 
[**ConductSearch**](SearchApi.md#ConductSearch) | **Get** /api/v1/find/all | 
[**GetMatchingEntities**](SearchApi.md#GetMatchingEntities) | **Get** /api/v1/entities/{prefix} | 

# **ConductRepoSearch**
> ConductRepoSearch(ctx, optional)


Get a list of apps and repositories that match the specified query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiConductRepoSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiConductRepoSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The page. | 
 **query** | **optional.String**| The search query. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConductSearch**
> ConductSearch(ctx, optional)


Get a list of entities and resources that match the specified query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiConductSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiConductSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **optional.String**| The search query. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMatchingEntities**
> GetMatchingEntities(ctx, prefix, optional)


Get a list of entities that match the specified prefix.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **prefix** | **string**|  | 
 **optional** | ***SearchApiGetMatchingEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiGetMatchingEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeOrgs** | **optional.Bool**| Whether to include orgs names. | 
 **includeTeams** | **optional.Bool**| Whether to include team names. | 
 **namespace** | **optional.String**| Namespace to use when querying for org entities. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

