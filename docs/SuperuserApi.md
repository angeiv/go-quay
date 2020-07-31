# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveServiceKey**](SuperuserApi.md#ApproveServiceKey) | **Post** /api/v1/superuser/approvedkeys/{kid} | 
[**ChangeOrganization**](SuperuserApi.md#ChangeOrganization) | **Put** /api/v1/superuser/organizations/{name} | 
[**CreateInstallUser**](SuperuserApi.md#CreateInstallUser) | **Post** /api/v1/superuser/users/ | 
[**CreateServiceKey**](SuperuserApi.md#CreateServiceKey) | **Post** /api/v1/superuser/keys | 
[**DeleteOrganization**](SuperuserApi.md#DeleteOrganization) | **Delete** /api/v1/superuser/organizations/{name} | 
[**DeleteServiceKey**](SuperuserApi.md#DeleteServiceKey) | **Delete** /api/v1/superuser/keys/{kid} | 
[**GetRepoBuildLogsSuperUser**](SuperuserApi.md#GetRepoBuildLogsSuperUser) | **Get** /api/v1/superuser/{build_uuid}/logs | 
[**GetRepoBuildStatusSuperUser**](SuperuserApi.md#GetRepoBuildStatusSuperUser) | **Get** /api/v1/superuser/{build_uuid}/status | 
[**GetRepoBuildSuperUser**](SuperuserApi.md#GetRepoBuildSuperUser) | **Get** /api/v1/superuser/{build_uuid}/build | 
[**GetServiceKey**](SuperuserApi.md#GetServiceKey) | **Get** /api/v1/superuser/keys/{kid} | 
[**ListAllUsers**](SuperuserApi.md#ListAllUsers) | **Get** /api/v1/superuser/users/ | 
[**ListServiceKeys**](SuperuserApi.md#ListServiceKeys) | **Get** /api/v1/superuser/keys | 
[**UpdateServiceKey**](SuperuserApi.md#UpdateServiceKey) | **Put** /api/v1/superuser/keys/{kid} | 

# **ApproveServiceKey**
> ApproveServiceKey(ctx, body, kid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApproveServiceKey**](ApproveServiceKey.md)| Request body contents. | 
  **kid** | **string**| The unique identifier for a service key | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChangeOrganization**
> ChangeOrganization(ctx, body, name)


Updates information about the specified user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOrg**](UpdateOrg.md)| Request body contents. | 
  **name** | **string**| The name of the organizaton being managed | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateInstallUser**
> CreateInstallUser(ctx, body)


Creates a new user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateInstallUser**](CreateInstallUser.md)| Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateServiceKey**
> CreateServiceKey(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateServiceKey**](CreateServiceKey.md)| Request body contents. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrganization**
> DeleteOrganization(ctx, name)


Deletes the specified organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The name of the organizaton being managed | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteServiceKey**
> DeleteServiceKey(ctx, kid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kid** | **string**| The unique identifier for a service key | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoBuildLogsSuperUser**
> GetRepoBuildLogsSuperUser(ctx, buildUuid)


Return the build logs for the build specified by the build uuid.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buildUuid** | **string**| The UUID of the build | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoBuildStatusSuperUser**
> GetRepoBuildStatusSuperUser(ctx, repository, buildUuid)


Return the status for the builds specified by the build uuids.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
  **buildUuid** | **string**| The UUID of the build | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRepoBuildSuperUser**
> GetRepoBuildSuperUser(ctx, repository, buildUuid)


Returns information about a build.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
  **buildUuid** | **string**| The UUID of the build | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServiceKey**
> GetServiceKey(ctx, kid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **kid** | **string**| The unique identifier for a service key | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllUsers**
> ListAllUsers(ctx, optional)


Returns a list of all users in the system.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SuperuserApiListAllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SuperuserApiListAllUsersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **disabled** | **optional.Bool**| If false, only enabled users will be returned. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListServiceKeys**
> ListServiceKeys(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateServiceKey**
> UpdateServiceKey(ctx, body, kid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PutServiceKey**](PutServiceKey.md)| Request body contents. | 
  **kid** | **string**| The unique identifier for a service key | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

