# Subscriber

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Subscriber Identifier | [optional] 
**SkipConfirmationNotification** | **bool** | If this is true, do not notify the user with changes to their subscription. | [optional] 
**Mode** | **string** | The communication mode of the subscriber. | [optional] 
**Email** | **string** | The email address to use to contact the subscriber. Used for Email and Webhook subscribers. | [optional] 
**Endpoint** | **string** | The URL where a webhook subscriber elects to receive updates. | [optional] 
**PhoneNumber** | **string** | The phone number used to contact an SMS subscriber | [optional] 
**PhoneCountry** | **string** | The two-character country code representing the country of which the phone_number is a part. | [optional] 
**DisplayPhoneNumber** | **string** | A formatted version of the phone_number and phone_country pair, nicely formatted for display. | [optional] 
**ObfuscatedChannelName** | **string** | Obfuscated slack channel name | [optional] 
**WorkspaceName** | **string** | The workspace name of the slack subscriber. | [optional] 
**QuarantinedAt** | [**time.Time**](time.Time.md) | The timestamp when the subscriber was quarantined due to an issue reaching them. | [optional] 
**PurgeAt** | [**time.Time**](time.Time.md) | The timestamp when a quarantined subscriber will be purged (unsubscribed). | [optional] 
**Components** | **string** | The components for which the subscriber has elected to receive updates. | [optional] 
**PageAccessUserId** | **string** | The Page Access user this subscriber belongs to (only for audience-specific pages). | [optional] 
**CreatedAt** | [**time.Time**](time.Time.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


