# CreateMirrorConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsEnabled** | **bool** | Used to enable or disable synchronizations. | [optional] [default to null]
**ExternalRegistryConfig** | [***UpdateMirrorConfigExternalRegistryConfig**](UpdateMirrorConfig_external_registry_config.md) |  | [optional] [default to null]
**SyncStartDate** | **string** | Determines the next time this repository is ready for synchronization. | [default to null]
**ExternalReference** | **string** | Location of the external repository. | [default to null]
**RootRule** | [***UpdateMirrorConfigRootRule**](UpdateMirrorConfig_root_rule.md) |  | [default to null]
**SyncInterval** | **int32** | Number of seconds after next_start_date to begin synchronizing. | [default to null]
**RobotUsername** | **string** | Username of robot which will be used for image pushes. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

