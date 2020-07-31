# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeRepoState**](RepositoryApi.md#ChangeRepoState) | **Put** /api/v1/repository/{repository}/changestate | 
[**ChangeRepoVisibility**](RepositoryApi.md#ChangeRepoVisibility) | **Post** /api/v1/repository/{repository}/changevisibility | 
[**CreateRepo**](RepositoryApi.md#CreateRepo) | **Post** /api/v1/repository | 
[**DeleteRepository**](RepositoryApi.md#DeleteRepository) | **Delete** /api/v1/repository/{repository} | 
[**GetRepo**](RepositoryApi.md#GetRepo) | **Get** /api/v1/repository/{repository} | 
[**ListRepos**](RepositoryApi.md#ListRepos) | **Get** /api/v1/repository | 
[**UpdateRepo**](RepositoryApi.md#UpdateRepo) | **Put** /api/v1/repository/{repository} | 

# **ChangeRepoState**
> ChangeRepoState(ctx, body, repository)


Change the state of a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeRepoState**](ChangeRepoState.md)| Request body contents. | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeRepoVisibility**
> ChangeRepoVisibility(ctx, body, repository)


Change the visibility of a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeVisibility**](ChangeVisibility.md)| Request body contents. | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateRepo**
> CreateRepo(ctx, body)


Create a new repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NewRepo**](NewRepo.md)| Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteRepository**
> DeleteRepository(ctx, repository)


Delete a repository.

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

# **GetRepo**
> GetRepo(ctx, repository, optional)


Fetch the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
 **optional** | ***RepositoryApiGetRepoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiGetRepoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTags** | **optional.Bool**| Whether to include repository tags | 
 **includeStats** | **optional.Bool**| Whether to include action statistics | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRepos**
> ListRepos(ctx, optional)


Fetch the list of repositories visible to the current user under a variety of situations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoryApiListReposOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoryApiListReposOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPage** | **optional.String**| The page token for the next page | 
 **repoKind** | **optional.String**| The kind of repositories to return | 
 **popularity** | **optional.Bool**| Whether to include the repository&#x27;s popularity metric. | 
 **lastModified** | **optional.Bool**| Whether to include when the repository was last modified. | 
 **public** | **optional.Bool**| Adds any repositories visible to the user by virtue of being public | 
 **starred** | **optional.Bool**| Filters the repositories returned to those starred by the user | 
 **namespace** | **optional.String**| Filters the repositories returned to this namespace | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRepo**
> UpdateRepo(ctx, body, repository)


Update the description in the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RepoUpdate**](RepoUpdate.md)| Request body contents. | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

