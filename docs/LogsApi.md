# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportOrgLogs**](LogsApi.md#ExportOrgLogs) | **Post** /api/v1/organization/{orgname}/exportlogs | 
[**ExportRepoLogs**](LogsApi.md#ExportRepoLogs) | **Post** /api/v1/repository/{repository}/exportlogs | 
[**ExportUserLogs**](LogsApi.md#ExportUserLogs) | **Post** /api/v1/user/exportlogs | 
[**GetAggregateOrgLogs**](LogsApi.md#GetAggregateOrgLogs) | **Get** /api/v1/organization/{orgname}/aggregatelogs | 
[**GetAggregateRepoLogs**](LogsApi.md#GetAggregateRepoLogs) | **Get** /api/v1/repository/{repository}/aggregatelogs | 
[**GetAggregateUserLogs**](LogsApi.md#GetAggregateUserLogs) | **Get** /api/v1/user/aggregatelogs | 
[**ListOrgLogs**](LogsApi.md#ListOrgLogs) | **Get** /api/v1/organization/{orgname}/logs | 
[**ListRepoLogs**](LogsApi.md#ListRepoLogs) | **Get** /api/v1/repository/{repository}/logs | 
[**ListUserLogs**](LogsApi.md#ListUserLogs) | **Get** /api/v1/user/logs | 

# **ExportOrgLogs**
> ExportOrgLogs(ctx, body, orgname, optional)


Exports the logs for the specified organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportLogs**](ExportLogs.md)| Request body contents. | 
  **orgname** | **string**| The name of the organization | 
 **optional** | ***LogsApiExportOrgLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiExportOrgLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endtime** | **optional.**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportRepoLogs**
> ExportRepoLogs(ctx, body, repository, optional)


Queues an export of the logs for the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportLogs**](ExportLogs.md)| Request body contents. | 
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
 **optional** | ***LogsApiExportRepoLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiExportRepoLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endtime** | **optional.**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExportUserLogs**
> ExportUserLogs(ctx, body, optional)


Returns the aggregated logs for the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExportLogs**](ExportLogs.md)| Request body contents. | 
 **optional** | ***LogsApiExportUserLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiExportUserLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endtime** | **optional.**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregateOrgLogs**
> GetAggregateOrgLogs(ctx, orgname, optional)


Gets the aggregated logs for the specified organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgname** | **string**| The name of the organization | 
 **optional** | ***LogsApiGetAggregateOrgLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiGetAggregateOrgLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **performer** | **optional.String**| Username for which to filter logs. | 
 **endtime** | **optional.String**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.String**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregateRepoLogs**
> GetAggregateRepoLogs(ctx, repository, optional)


Returns the aggregated logs for the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
 **optional** | ***LogsApiGetAggregateRepoLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiGetAggregateRepoLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endtime** | **optional.String**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.String**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAggregateUserLogs**
> GetAggregateUserLogs(ctx, optional)


Returns the aggregated logs for the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogsApiGetAggregateUserLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiGetAggregateUserLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **performer** | **optional.String**| Username for which to filter logs. | 
 **endtime** | **optional.String**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.String**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOrgLogs**
> ListOrgLogs(ctx, orgname, optional)


List the logs for the specified organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgname** | **string**| The name of the organization | 
 **optional** | ***LogsApiListOrgLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiListOrgLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPage** | **optional.String**| The page token for the next page | 
 **performer** | **optional.String**| Username for which to filter logs. | 
 **endtime** | **optional.String**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.String**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRepoLogs**
> ListRepoLogs(ctx, repository, optional)


List the logs for the specified repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repository** | **string**| The full path of the repository. e.g. namespace/name | 
 **optional** | ***LogsApiListRepoLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiListRepoLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPage** | **optional.String**| The page token for the next page | 
 **endtime** | **optional.String**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.String**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUserLogs**
> ListUserLogs(ctx, optional)


List the logs for the current user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogsApiListUserLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsApiListUserLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPage** | **optional.String**| The page token for the next page | 
 **performer** | **optional.String**| Username for which to filter logs. | 
 **endtime** | **optional.String**| Latest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 
 **starttime** | **optional.String**| Earliest time for logs. Format: \&quot;%m/%d/%Y\&quot; in UTC. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

