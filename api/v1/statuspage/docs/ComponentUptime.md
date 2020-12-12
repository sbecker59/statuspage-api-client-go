# ComponentUptime

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RangeStart** | [**time.Time**](time.Time.md) | Start date used for uptime calculation (see the warnings field in the response if this value does not match the start parameter you provided). | [optional] 
**RangeEnd** | [**time.Time**](time.Time.md) | End date used for uptime calculation (see the warnings field in the response if this value does not match the end parameter you provided). | [optional] 
**UptimePercentage** | **float32** | Uptime percentage for a component | [optional] 
**MajorOutage** | **int32** | Seconds of major outage | [optional] 
**PartialOutage** | **int32** | Seconds of partial outage | [optional] 
**Warnings** | **string** | Warning messages related to the uptime query that may occur | [optional] 
**Id** | **string** | Component identifier | [optional] 
**Name** | **string** | Component display name | [optional] 
**RelatedEvents** | [***ComponentUptimeRelatedEvents**](ComponentUptime_related_events.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


