# \HeadscaleServiceApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HeadscaleServiceCreateApiKey**](HeadscaleServiceApi.md#HeadscaleServiceCreateApiKey) | **Post** /api/v1/apikey | --- ApiKeys start ---
[**HeadscaleServiceCreatePreAuthKey**](HeadscaleServiceApi.md#HeadscaleServiceCreatePreAuthKey) | **Post** /api/v1/preauthkey | --- PreAuthKeys start ---
[**HeadscaleServiceCreateUser**](HeadscaleServiceApi.md#HeadscaleServiceCreateUser) | **Post** /api/v1/user | 
[**HeadscaleServiceDebugCreateMachine**](HeadscaleServiceApi.md#HeadscaleServiceDebugCreateMachine) | **Post** /api/v1/debug/machine | --- Machine start ---
[**HeadscaleServiceDeleteMachine**](HeadscaleServiceApi.md#HeadscaleServiceDeleteMachine) | **Delete** /api/v1/machine/{machineId} | 
[**HeadscaleServiceDeleteRoute**](HeadscaleServiceApi.md#HeadscaleServiceDeleteRoute) | **Delete** /api/v1/routes/{routeId} | 
[**HeadscaleServiceDeleteUser**](HeadscaleServiceApi.md#HeadscaleServiceDeleteUser) | **Delete** /api/v1/user/{name} | 
[**HeadscaleServiceDisableRoute**](HeadscaleServiceApi.md#HeadscaleServiceDisableRoute) | **Post** /api/v1/routes/{routeId}/disable | 
[**HeadscaleServiceEnableRoute**](HeadscaleServiceApi.md#HeadscaleServiceEnableRoute) | **Post** /api/v1/routes/{routeId}/enable | 
[**HeadscaleServiceExpireApiKey**](HeadscaleServiceApi.md#HeadscaleServiceExpireApiKey) | **Post** /api/v1/apikey/expire | 
[**HeadscaleServiceExpireMachine**](HeadscaleServiceApi.md#HeadscaleServiceExpireMachine) | **Post** /api/v1/machine/{machineId}/expire | 
[**HeadscaleServiceExpirePreAuthKey**](HeadscaleServiceApi.md#HeadscaleServiceExpirePreAuthKey) | **Post** /api/v1/preauthkey/expire | 
[**HeadscaleServiceGetMachine**](HeadscaleServiceApi.md#HeadscaleServiceGetMachine) | **Get** /api/v1/machine/{machineId} | 
[**HeadscaleServiceGetMachineRoutes**](HeadscaleServiceApi.md#HeadscaleServiceGetMachineRoutes) | **Get** /api/v1/machine/{machineId}/routes | 
[**HeadscaleServiceGetRoutes**](HeadscaleServiceApi.md#HeadscaleServiceGetRoutes) | **Get** /api/v1/routes | --- Route start ---
[**HeadscaleServiceGetUser**](HeadscaleServiceApi.md#HeadscaleServiceGetUser) | **Get** /api/v1/user/{name} | --- User start ---
[**HeadscaleServiceListApiKeys**](HeadscaleServiceApi.md#HeadscaleServiceListApiKeys) | **Get** /api/v1/apikey | 
[**HeadscaleServiceListMachines**](HeadscaleServiceApi.md#HeadscaleServiceListMachines) | **Get** /api/v1/machine | 
[**HeadscaleServiceListPreAuthKeys**](HeadscaleServiceApi.md#HeadscaleServiceListPreAuthKeys) | **Get** /api/v1/preauthkey | 
[**HeadscaleServiceListUsers**](HeadscaleServiceApi.md#HeadscaleServiceListUsers) | **Get** /api/v1/user | 
[**HeadscaleServiceMoveMachine**](HeadscaleServiceApi.md#HeadscaleServiceMoveMachine) | **Post** /api/v1/machine/{machineId}/user | 
[**HeadscaleServiceRegisterMachine**](HeadscaleServiceApi.md#HeadscaleServiceRegisterMachine) | **Post** /api/v1/machine/register | 
[**HeadscaleServiceRenameMachine**](HeadscaleServiceApi.md#HeadscaleServiceRenameMachine) | **Post** /api/v1/machine/{machineId}/rename/{newName} | 
[**HeadscaleServiceRenameUser**](HeadscaleServiceApi.md#HeadscaleServiceRenameUser) | **Post** /api/v1/user/{oldName}/rename/{newName} | 
[**HeadscaleServiceSetTags**](HeadscaleServiceApi.md#HeadscaleServiceSetTags) | **Post** /api/v1/machine/{machineId}/tags | 


# **HeadscaleServiceCreateApiKey**
> V1CreateApiKeyResponse HeadscaleServiceCreateApiKey(ctx, body)
--- ApiKeys start ---

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateApiKeyRequest**](V1CreateApiKeyRequest.md)|  | 

### Return type

[**V1CreateApiKeyResponse**](v1CreateApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceCreatePreAuthKey**
> V1CreatePreAuthKeyResponse HeadscaleServiceCreatePreAuthKey(ctx, body)
--- PreAuthKeys start ---

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreatePreAuthKeyRequest**](V1CreatePreAuthKeyRequest.md)|  | 

### Return type

[**V1CreatePreAuthKeyResponse**](v1CreatePreAuthKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceCreateUser**
> V1CreateUserResponse HeadscaleServiceCreateUser(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateUserRequest**](V1CreateUserRequest.md)|  | 

### Return type

[**V1CreateUserResponse**](v1CreateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceDebugCreateMachine**
> V1DebugCreateMachineResponse HeadscaleServiceDebugCreateMachine(ctx, body)
--- Machine start ---

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1DebugCreateMachineRequest**](V1DebugCreateMachineRequest.md)|  | 

### Return type

[**V1DebugCreateMachineResponse**](v1DebugCreateMachineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceDeleteMachine**
> V1DeleteMachineResponse HeadscaleServiceDeleteMachine(ctx, machineId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **machineId** | **string**|  | 

### Return type

[**V1DeleteMachineResponse**](v1DeleteMachineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceDeleteRoute**
> V1DeleteRouteResponse HeadscaleServiceDeleteRoute(ctx, routeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeId** | **string**|  | 

### Return type

[**V1DeleteRouteResponse**](v1DeleteRouteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceDeleteUser**
> V1DeleteUserResponse HeadscaleServiceDeleteUser(ctx, name)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**V1DeleteUserResponse**](v1DeleteUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceDisableRoute**
> V1DisableRouteResponse HeadscaleServiceDisableRoute(ctx, routeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeId** | **string**|  | 

### Return type

[**V1DisableRouteResponse**](v1DisableRouteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceEnableRoute**
> V1EnableRouteResponse HeadscaleServiceEnableRoute(ctx, routeId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **routeId** | **string**|  | 

### Return type

[**V1EnableRouteResponse**](v1EnableRouteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceExpireApiKey**
> V1ExpireApiKeyResponse HeadscaleServiceExpireApiKey(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ExpireApiKeyRequest**](V1ExpireApiKeyRequest.md)|  | 

### Return type

[**V1ExpireApiKeyResponse**](v1ExpireApiKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceExpireMachine**
> V1ExpireMachineResponse HeadscaleServiceExpireMachine(ctx, machineId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **machineId** | **string**|  | 

### Return type

[**V1ExpireMachineResponse**](v1ExpireMachineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceExpirePreAuthKey**
> V1ExpirePreAuthKeyResponse HeadscaleServiceExpirePreAuthKey(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ExpirePreAuthKeyRequest**](V1ExpirePreAuthKeyRequest.md)|  | 

### Return type

[**V1ExpirePreAuthKeyResponse**](v1ExpirePreAuthKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceGetMachine**
> V1GetMachineResponse HeadscaleServiceGetMachine(ctx, machineId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **machineId** | **string**|  | 

### Return type

[**V1GetMachineResponse**](v1GetMachineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceGetMachineRoutes**
> V1GetMachineRoutesResponse HeadscaleServiceGetMachineRoutes(ctx, machineId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **machineId** | **string**|  | 

### Return type

[**V1GetMachineRoutesResponse**](v1GetMachineRoutesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceGetRoutes**
> V1GetRoutesResponse HeadscaleServiceGetRoutes(ctx, )
--- Route start ---

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**V1GetRoutesResponse**](v1GetRoutesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceGetUser**
> V1GetUserResponse HeadscaleServiceGetUser(ctx, name)
--- User start ---

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**|  | 

### Return type

[**V1GetUserResponse**](v1GetUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceListApiKeys**
> V1ListApiKeysResponse HeadscaleServiceListApiKeys(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListApiKeysResponse**](v1ListApiKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceListMachines**
> V1ListMachinesResponse HeadscaleServiceListMachines(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HeadscaleServiceApiHeadscaleServiceListMachinesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HeadscaleServiceApiHeadscaleServiceListMachinesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 

### Return type

[**V1ListMachinesResponse**](v1ListMachinesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceListPreAuthKeys**
> V1ListPreAuthKeysResponse HeadscaleServiceListPreAuthKeys(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HeadscaleServiceApiHeadscaleServiceListPreAuthKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HeadscaleServiceApiHeadscaleServiceListPreAuthKeysOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 

### Return type

[**V1ListPreAuthKeysResponse**](v1ListPreAuthKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceListUsers**
> V1ListUsersResponse HeadscaleServiceListUsers(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**V1ListUsersResponse**](v1ListUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceMoveMachine**
> V1MoveMachineResponse HeadscaleServiceMoveMachine(ctx, machineId, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **machineId** | **string**|  | 
 **optional** | ***HeadscaleServiceApiHeadscaleServiceMoveMachineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HeadscaleServiceApiHeadscaleServiceMoveMachineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **optional.String**|  | 

### Return type

[**V1MoveMachineResponse**](v1MoveMachineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceRegisterMachine**
> V1RegisterMachineResponse HeadscaleServiceRegisterMachine(ctx, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HeadscaleServiceApiHeadscaleServiceRegisterMachineOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HeadscaleServiceApiHeadscaleServiceRegisterMachineOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **optional.String**|  | 
 **key** | **optional.String**|  | 

### Return type

[**V1RegisterMachineResponse**](v1RegisterMachineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceRenameMachine**
> V1RenameMachineResponse HeadscaleServiceRenameMachine(ctx, machineId, newName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **machineId** | **string**|  | 
  **newName** | **string**|  | 

### Return type

[**V1RenameMachineResponse**](v1RenameMachineResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceRenameUser**
> V1RenameUserResponse HeadscaleServiceRenameUser(ctx, oldName, newName)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oldName** | **string**|  | 
  **newName** | **string**|  | 

### Return type

[**V1RenameUserResponse**](v1RenameUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HeadscaleServiceSetTags**
> V1SetTagsResponse HeadscaleServiceSetTags(ctx, machineId, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **machineId** | **string**|  | 
  **body** | [**Body**](Body.md)|  | 

### Return type

[**V1SetTagsResponse**](v1SetTagsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

