# Postmortem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviewKey** | Pointer to **string** | Preview Key | [optional] 
**Body** | Pointer to **string** | Postmortem body | [optional] 
**BodyUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**BodyDraft** | Pointer to **string** | Body draft | [optional] 
**BodyDraftUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**NotifySubscribers** | Pointer to **bool** | Should email subscribers be notified. | [optional] 
**NotifyTwitter** | Pointer to **bool** | Should Twitter followers be notified. | [optional] 
**CustomTweet** | Pointer to **string** | Custom tweet for Incident Postmortem | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPostmortem

`func NewPostmortem() *Postmortem`

NewPostmortem instantiates a new Postmortem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostmortemWithDefaults

`func NewPostmortemWithDefaults() *Postmortem`

NewPostmortemWithDefaults instantiates a new Postmortem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviewKey

`func (o *Postmortem) GetPreviewKey() string`

GetPreviewKey returns the PreviewKey field if non-nil, zero value otherwise.

### GetPreviewKeyOk

`func (o *Postmortem) GetPreviewKeyOk() (*string, bool)`

GetPreviewKeyOk returns a tuple with the PreviewKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewKey

`func (o *Postmortem) SetPreviewKey(v string)`

SetPreviewKey sets PreviewKey field to given value.

### HasPreviewKey

`func (o *Postmortem) HasPreviewKey() bool`

HasPreviewKey returns a boolean if a field has been set.

### GetBody

`func (o *Postmortem) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Postmortem) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Postmortem) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Postmortem) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetBodyUpdatedAt

`func (o *Postmortem) GetBodyUpdatedAt() time.Time`

GetBodyUpdatedAt returns the BodyUpdatedAt field if non-nil, zero value otherwise.

### GetBodyUpdatedAtOk

`func (o *Postmortem) GetBodyUpdatedAtOk() (*time.Time, bool)`

GetBodyUpdatedAtOk returns a tuple with the BodyUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyUpdatedAt

`func (o *Postmortem) SetBodyUpdatedAt(v time.Time)`

SetBodyUpdatedAt sets BodyUpdatedAt field to given value.

### HasBodyUpdatedAt

`func (o *Postmortem) HasBodyUpdatedAt() bool`

HasBodyUpdatedAt returns a boolean if a field has been set.

### GetBodyDraft

`func (o *Postmortem) GetBodyDraft() string`

GetBodyDraft returns the BodyDraft field if non-nil, zero value otherwise.

### GetBodyDraftOk

`func (o *Postmortem) GetBodyDraftOk() (*string, bool)`

GetBodyDraftOk returns a tuple with the BodyDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyDraft

`func (o *Postmortem) SetBodyDraft(v string)`

SetBodyDraft sets BodyDraft field to given value.

### HasBodyDraft

`func (o *Postmortem) HasBodyDraft() bool`

HasBodyDraft returns a boolean if a field has been set.

### GetBodyDraftUpdatedAt

`func (o *Postmortem) GetBodyDraftUpdatedAt() time.Time`

GetBodyDraftUpdatedAt returns the BodyDraftUpdatedAt field if non-nil, zero value otherwise.

### GetBodyDraftUpdatedAtOk

`func (o *Postmortem) GetBodyDraftUpdatedAtOk() (*time.Time, bool)`

GetBodyDraftUpdatedAtOk returns a tuple with the BodyDraftUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyDraftUpdatedAt

`func (o *Postmortem) SetBodyDraftUpdatedAt(v time.Time)`

SetBodyDraftUpdatedAt sets BodyDraftUpdatedAt field to given value.

### HasBodyDraftUpdatedAt

`func (o *Postmortem) HasBodyDraftUpdatedAt() bool`

HasBodyDraftUpdatedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *Postmortem) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Postmortem) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Postmortem) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Postmortem) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetNotifySubscribers

`func (o *Postmortem) GetNotifySubscribers() bool`

GetNotifySubscribers returns the NotifySubscribers field if non-nil, zero value otherwise.

### GetNotifySubscribersOk

`func (o *Postmortem) GetNotifySubscribersOk() (*bool, bool)`

GetNotifySubscribersOk returns a tuple with the NotifySubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySubscribers

`func (o *Postmortem) SetNotifySubscribers(v bool)`

SetNotifySubscribers sets NotifySubscribers field to given value.

### HasNotifySubscribers

`func (o *Postmortem) HasNotifySubscribers() bool`

HasNotifySubscribers returns a boolean if a field has been set.

### GetNotifyTwitter

`func (o *Postmortem) GetNotifyTwitter() bool`

GetNotifyTwitter returns the NotifyTwitter field if non-nil, zero value otherwise.

### GetNotifyTwitterOk

`func (o *Postmortem) GetNotifyTwitterOk() (*bool, bool)`

GetNotifyTwitterOk returns a tuple with the NotifyTwitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTwitter

`func (o *Postmortem) SetNotifyTwitter(v bool)`

SetNotifyTwitter sets NotifyTwitter field to given value.

### HasNotifyTwitter

`func (o *Postmortem) HasNotifyTwitter() bool`

HasNotifyTwitter returns a boolean if a field has been set.

### GetCustomTweet

`func (o *Postmortem) GetCustomTweet() string`

GetCustomTweet returns the CustomTweet field if non-nil, zero value otherwise.

### GetCustomTweetOk

`func (o *Postmortem) GetCustomTweetOk() (*string, bool)`

GetCustomTweetOk returns a tuple with the CustomTweet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTweet

`func (o *Postmortem) SetCustomTweet(v string)`

SetCustomTweet sets CustomTweet field to given value.

### HasCustomTweet

`func (o *Postmortem) HasCustomTweet() bool`

HasCustomTweet returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Postmortem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Postmortem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Postmortem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Postmortem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Postmortem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Postmortem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Postmortem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Postmortem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


