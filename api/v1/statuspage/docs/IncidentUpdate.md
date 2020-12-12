# IncidentUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Incident Update Identifier. | [optional] 
**IncidentId** | **string** | Incident Identifier. | [optional] 
**AffectedComponents** | [**[]interface{}**](interface{}.md) | Affected components associated with the incident update. | [optional] 
**Body** | **string** | Incident update body. | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) | The timestamp when the incident update was created at. | [optional] 
**CustomTweet** | **string** | An optional customized tweet message for incident postmortem. | [optional] 
**DeliverNotifications** | **bool** | Controls whether to delivery notifications. | [optional] 
**DisplayAt** | [**time.Time**](time.Time.md) | Timestamp when incident update is happened. | [optional] 
**Status** | **string** | The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed. | [optional] 
**TweetId** | **string** | Tweet identifier associated to this incident update. | [optional] 
**TwitterUpdatedAt** | [**time.Time**](time.Time.md) | The timestamp when twitter updated at. | [optional] 
**UpdatedAt** | [**time.Time**](time.Time.md) | The timestamp when the incident update is updated. | [optional] 
**WantsTwitterUpdate** | **bool** | Controls whether to create twitter update. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


