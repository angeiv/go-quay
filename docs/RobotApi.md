# {{classname}}

All URIs are relative to *http://localhost/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgRobot**](RobotApi.md#CreateOrgRobot) | **Put** /api/v1/organization/{orgname}/robots/{robot_shortname} | 
[**CreateUserRobot**](RobotApi.md#CreateUserRobot) | **Put** /api/v1/user/robots/{robot_shortname} | 
[**DeleteOrgRobot**](RobotApi.md#DeleteOrgRobot) | **Delete** /api/v1/organization/{orgname}/robots/{robot_shortname} | 
[**DeleteUserRobot**](RobotApi.md#DeleteUserRobot) | **Delete** /api/v1/user/robots/{robot_shortname} | 
[**GetOrgRobot**](RobotApi.md#GetOrgRobot) | **Get** /api/v1/organization/{orgname}/robots/{robot_shortname} | 
[**GetOrgRobotPermissions**](RobotApi.md#GetOrgRobotPermissions) | **Get** /api/v1/organization/{orgname}/robots/{robot_shortname}/permissions | 
[**GetOrgRobots**](RobotApi.md#GetOrgRobots) | **Get** /api/v1/organization/{orgname}/robots | 
[**GetUserRobot**](RobotApi.md#GetUserRobot) | **Get** /api/v1/user/robots/{robot_shortname} | 
[**GetUserRobotPermissions**](RobotApi.md#GetUserRobotPermissions) | **Get** /api/v1/user/robots/{robot_shortname}/permissions | 
[**GetUserRobots**](RobotApi.md#GetUserRobots) | **Get** /api/v1/user/robots | 
[**RegenerateOrgRobotToken**](RobotApi.md#RegenerateOrgRobotToken) | **Post** /api/v1/organization/{orgname}/robots/{robot_shortname}/regenerate | 
[**RegenerateUserRobotToken**](RobotApi.md#RegenerateUserRobotToken) | **Post** /api/v1/user/robots/{robot_shortname}/regenerate | 

# **CreateOrgRobot**
> CreateOrgRobot(ctx, body, orgname, robotShortname)


Create a new robot in the organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRobot**](CreateRobot.md)| Request body contents. | 
  **orgname** | **string**| The name of the organization | 
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUserRobot**
> CreateUserRobot(ctx, body, robotShortname)


Create a new user robot with the specified name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateRobot**](CreateRobot.md)| Request body contents. | 
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrgRobot**
> DeleteOrgRobot(ctx, orgname, robotShortname)


Delete an existing organization robot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgname** | **string**| The name of the organization | 
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUserRobot**
> DeleteUserRobot(ctx, robotShortname)


Delete an existing robot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgRobot**
> GetOrgRobot(ctx, orgname, robotShortname)


Returns the organization's robot with the specified name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgname** | **string**| The name of the organization | 
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgRobotPermissions**
> GetOrgRobotPermissions(ctx, orgname, robotShortname)


Returns the list of repository permissions for the org's robot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgname** | **string**| The name of the organization | 
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrgRobots**
> GetOrgRobots(ctx, orgname, optional)


List the organization's robots.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgname** | **string**| The name of the organization | 
 **optional** | ***RobotApiGetOrgRobotsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RobotApiGetOrgRobotsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| If specified, the number of robots to return. | 
 **token** | **optional.Bool**| If false, the robot&#x27;s token is not returned. | 
 **permissions** | **optional.Bool**| Whether to include repostories and teams in which the robots have permission. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserRobot**
> GetUserRobot(ctx, robotShortname)


Returns the user's robot with the specified name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserRobotPermissions**
> GetUserRobotPermissions(ctx, robotShortname)


Returns the list of repository permissions for the user's robot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserRobots**
> GetUserRobots(ctx, optional)


List the available robots for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RobotApiGetUserRobotsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RobotApiGetUserRobotsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| If specified, the number of robots to return. | 
 **token** | **optional.Bool**| If false, the robot&#x27;s token is not returned. | 
 **permissions** | **optional.Bool**| Whether to include repositories and teams in which the robots have permission. | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateOrgRobotToken**
> RegenerateOrgRobotToken(ctx, orgname, robotShortname)


Regenerates the token for an organization robot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgname** | **string**| The name of the organization | 
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegenerateUserRobotToken**
> RegenerateUserRobotToken(ctx, robotShortname)


Regenerates the token for a user's robot.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **robotShortname** | **string**| The short name for the robot, without any user or organization prefix | 

### Return type

 (empty response body)

### Authorization

[oauth2_implicit](../README.md#oauth2_implicit)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

