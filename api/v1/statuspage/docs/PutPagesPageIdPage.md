# PutPagesPageIdPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of your page to be displayed | [optional] 
**Domain** | Pointer to **string** | CNAME alias for your status page | [optional] 
**Subdomain** | Pointer to **string** | Subdomain at which to access your status page | [optional] 
**Url** | Pointer to **string** | Website of your page.  Clicking on your statuspage image will link here. | [optional] 
**Branding** | Pointer to **string** | The main template your statuspage will use | [optional] 
**CssBodyBackgroundColor** | Pointer to **string** | CSS Color | [optional] 
**CssFontColor** | Pointer to **string** | CSS Color | [optional] 
**CssLightFontColor** | Pointer to **string** | CSS Color | [optional] 
**CssGreens** | Pointer to **string** | CSS Color | [optional] 
**CssYellows** | Pointer to **string** | CSS Color | [optional] 
**CssOranges** | Pointer to **string** | CSS Color | [optional] 
**CssReds** | Pointer to **string** | CSS Color | [optional] 
**CssBlues** | Pointer to **string** | CSS Color | [optional] 
**CssBorderColor** | Pointer to **string** | CSS Color | [optional] 
**CssGraphColor** | Pointer to **string** | CSS Color | [optional] 
**CssLinkColor** | Pointer to **string** | CSS Color | [optional] 
**CssNoData** | Pointer to **string** | CSS Color | [optional] 
**HiddenFromSearch** | Pointer to **bool** | Should your page hide itself from search engines | [optional] 
**ViewersMustBeTeamMembers** | Pointer to **bool** |  | [optional] 
**AllowPageSubscribers** | Pointer to **bool** | Can your users subscribe to all notifications on the page | [optional] 
**AllowIncidentSubscribers** | Pointer to **bool** | Can your users subscribe to notifications for a single incident | [optional] 
**AllowEmailSubscribers** | Pointer to **bool** | Can your users choose to receive notifications via email | [optional] 
**AllowSmsSubscribers** | Pointer to **bool** | Can your users choose to receive notifications via SMS | [optional] 
**AllowRssAtomFeeds** | Pointer to **bool** | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) | [optional] 
**AllowWebhookSubscribers** | Pointer to **bool** | Can your users choose to receive notifications via Webhooks | [optional] 
**NotificationsFromEmail** | Pointer to **string** | Allows you to customize the email address your page notifications come from | [optional] 
**TimeZone** | Pointer to **string** | Timezone configured for your page | [optional] 
**NotificationsEmailFooter** | Pointer to **string** | Allows you to customize the footer appearing on your notification emails.  Accepts Markdown for formatting | [optional] 

## Methods

### NewPutPagesPageIdPage

`func NewPutPagesPageIdPage() *PutPagesPageIdPage`

NewPutPagesPageIdPage instantiates a new PutPagesPageIdPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPagesPageIdPageWithDefaults

`func NewPutPagesPageIdPageWithDefaults() *PutPagesPageIdPage`

NewPutPagesPageIdPageWithDefaults instantiates a new PutPagesPageIdPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PutPagesPageIdPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutPagesPageIdPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutPagesPageIdPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PutPagesPageIdPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *PutPagesPageIdPage) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PutPagesPageIdPage) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PutPagesPageIdPage) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PutPagesPageIdPage) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSubdomain

`func (o *PutPagesPageIdPage) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *PutPagesPageIdPage) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *PutPagesPageIdPage) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *PutPagesPageIdPage) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUrl

`func (o *PutPagesPageIdPage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PutPagesPageIdPage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PutPagesPageIdPage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PutPagesPageIdPage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBranding

`func (o *PutPagesPageIdPage) GetBranding() string`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *PutPagesPageIdPage) GetBrandingOk() (*string, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *PutPagesPageIdPage) SetBranding(v string)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *PutPagesPageIdPage) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetCssBodyBackgroundColor

`func (o *PutPagesPageIdPage) GetCssBodyBackgroundColor() string`

GetCssBodyBackgroundColor returns the CssBodyBackgroundColor field if non-nil, zero value otherwise.

### GetCssBodyBackgroundColorOk

`func (o *PutPagesPageIdPage) GetCssBodyBackgroundColorOk() (*string, bool)`

GetCssBodyBackgroundColorOk returns a tuple with the CssBodyBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBodyBackgroundColor

`func (o *PutPagesPageIdPage) SetCssBodyBackgroundColor(v string)`

SetCssBodyBackgroundColor sets CssBodyBackgroundColor field to given value.

### HasCssBodyBackgroundColor

`func (o *PutPagesPageIdPage) HasCssBodyBackgroundColor() bool`

HasCssBodyBackgroundColor returns a boolean if a field has been set.

### GetCssFontColor

`func (o *PutPagesPageIdPage) GetCssFontColor() string`

GetCssFontColor returns the CssFontColor field if non-nil, zero value otherwise.

### GetCssFontColorOk

`func (o *PutPagesPageIdPage) GetCssFontColorOk() (*string, bool)`

GetCssFontColorOk returns a tuple with the CssFontColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFontColor

`func (o *PutPagesPageIdPage) SetCssFontColor(v string)`

SetCssFontColor sets CssFontColor field to given value.

### HasCssFontColor

`func (o *PutPagesPageIdPage) HasCssFontColor() bool`

HasCssFontColor returns a boolean if a field has been set.

### GetCssLightFontColor

`func (o *PutPagesPageIdPage) GetCssLightFontColor() string`

GetCssLightFontColor returns the CssLightFontColor field if non-nil, zero value otherwise.

### GetCssLightFontColorOk

`func (o *PutPagesPageIdPage) GetCssLightFontColorOk() (*string, bool)`

GetCssLightFontColorOk returns a tuple with the CssLightFontColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLightFontColor

`func (o *PutPagesPageIdPage) SetCssLightFontColor(v string)`

SetCssLightFontColor sets CssLightFontColor field to given value.

### HasCssLightFontColor

`func (o *PutPagesPageIdPage) HasCssLightFontColor() bool`

HasCssLightFontColor returns a boolean if a field has been set.

### GetCssGreens

`func (o *PutPagesPageIdPage) GetCssGreens() string`

GetCssGreens returns the CssGreens field if non-nil, zero value otherwise.

### GetCssGreensOk

`func (o *PutPagesPageIdPage) GetCssGreensOk() (*string, bool)`

GetCssGreensOk returns a tuple with the CssGreens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssGreens

`func (o *PutPagesPageIdPage) SetCssGreens(v string)`

SetCssGreens sets CssGreens field to given value.

### HasCssGreens

`func (o *PutPagesPageIdPage) HasCssGreens() bool`

HasCssGreens returns a boolean if a field has been set.

### GetCssYellows

`func (o *PutPagesPageIdPage) GetCssYellows() string`

GetCssYellows returns the CssYellows field if non-nil, zero value otherwise.

### GetCssYellowsOk

`func (o *PutPagesPageIdPage) GetCssYellowsOk() (*string, bool)`

GetCssYellowsOk returns a tuple with the CssYellows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssYellows

`func (o *PutPagesPageIdPage) SetCssYellows(v string)`

SetCssYellows sets CssYellows field to given value.

### HasCssYellows

`func (o *PutPagesPageIdPage) HasCssYellows() bool`

HasCssYellows returns a boolean if a field has been set.

### GetCssOranges

`func (o *PutPagesPageIdPage) GetCssOranges() string`

GetCssOranges returns the CssOranges field if non-nil, zero value otherwise.

### GetCssOrangesOk

`func (o *PutPagesPageIdPage) GetCssOrangesOk() (*string, bool)`

GetCssOrangesOk returns a tuple with the CssOranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssOranges

`func (o *PutPagesPageIdPage) SetCssOranges(v string)`

SetCssOranges sets CssOranges field to given value.

### HasCssOranges

`func (o *PutPagesPageIdPage) HasCssOranges() bool`

HasCssOranges returns a boolean if a field has been set.

### GetCssReds

`func (o *PutPagesPageIdPage) GetCssReds() string`

GetCssReds returns the CssReds field if non-nil, zero value otherwise.

### GetCssRedsOk

`func (o *PutPagesPageIdPage) GetCssRedsOk() (*string, bool)`

GetCssRedsOk returns a tuple with the CssReds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssReds

`func (o *PutPagesPageIdPage) SetCssReds(v string)`

SetCssReds sets CssReds field to given value.

### HasCssReds

`func (o *PutPagesPageIdPage) HasCssReds() bool`

HasCssReds returns a boolean if a field has been set.

### GetCssBlues

`func (o *PutPagesPageIdPage) GetCssBlues() string`

GetCssBlues returns the CssBlues field if non-nil, zero value otherwise.

### GetCssBluesOk

`func (o *PutPagesPageIdPage) GetCssBluesOk() (*string, bool)`

GetCssBluesOk returns a tuple with the CssBlues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBlues

`func (o *PutPagesPageIdPage) SetCssBlues(v string)`

SetCssBlues sets CssBlues field to given value.

### HasCssBlues

`func (o *PutPagesPageIdPage) HasCssBlues() bool`

HasCssBlues returns a boolean if a field has been set.

### GetCssBorderColor

`func (o *PutPagesPageIdPage) GetCssBorderColor() string`

GetCssBorderColor returns the CssBorderColor field if non-nil, zero value otherwise.

### GetCssBorderColorOk

`func (o *PutPagesPageIdPage) GetCssBorderColorOk() (*string, bool)`

GetCssBorderColorOk returns a tuple with the CssBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBorderColor

`func (o *PutPagesPageIdPage) SetCssBorderColor(v string)`

SetCssBorderColor sets CssBorderColor field to given value.

### HasCssBorderColor

`func (o *PutPagesPageIdPage) HasCssBorderColor() bool`

HasCssBorderColor returns a boolean if a field has been set.

### GetCssGraphColor

`func (o *PutPagesPageIdPage) GetCssGraphColor() string`

GetCssGraphColor returns the CssGraphColor field if non-nil, zero value otherwise.

### GetCssGraphColorOk

`func (o *PutPagesPageIdPage) GetCssGraphColorOk() (*string, bool)`

GetCssGraphColorOk returns a tuple with the CssGraphColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssGraphColor

`func (o *PutPagesPageIdPage) SetCssGraphColor(v string)`

SetCssGraphColor sets CssGraphColor field to given value.

### HasCssGraphColor

`func (o *PutPagesPageIdPage) HasCssGraphColor() bool`

HasCssGraphColor returns a boolean if a field has been set.

### GetCssLinkColor

`func (o *PutPagesPageIdPage) GetCssLinkColor() string`

GetCssLinkColor returns the CssLinkColor field if non-nil, zero value otherwise.

### GetCssLinkColorOk

`func (o *PutPagesPageIdPage) GetCssLinkColorOk() (*string, bool)`

GetCssLinkColorOk returns a tuple with the CssLinkColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLinkColor

`func (o *PutPagesPageIdPage) SetCssLinkColor(v string)`

SetCssLinkColor sets CssLinkColor field to given value.

### HasCssLinkColor

`func (o *PutPagesPageIdPage) HasCssLinkColor() bool`

HasCssLinkColor returns a boolean if a field has been set.

### GetCssNoData

`func (o *PutPagesPageIdPage) GetCssNoData() string`

GetCssNoData returns the CssNoData field if non-nil, zero value otherwise.

### GetCssNoDataOk

`func (o *PutPagesPageIdPage) GetCssNoDataOk() (*string, bool)`

GetCssNoDataOk returns a tuple with the CssNoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssNoData

`func (o *PutPagesPageIdPage) SetCssNoData(v string)`

SetCssNoData sets CssNoData field to given value.

### HasCssNoData

`func (o *PutPagesPageIdPage) HasCssNoData() bool`

HasCssNoData returns a boolean if a field has been set.

### GetHiddenFromSearch

`func (o *PutPagesPageIdPage) GetHiddenFromSearch() bool`

GetHiddenFromSearch returns the HiddenFromSearch field if non-nil, zero value otherwise.

### GetHiddenFromSearchOk

`func (o *PutPagesPageIdPage) GetHiddenFromSearchOk() (*bool, bool)`

GetHiddenFromSearchOk returns a tuple with the HiddenFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromSearch

`func (o *PutPagesPageIdPage) SetHiddenFromSearch(v bool)`

SetHiddenFromSearch sets HiddenFromSearch field to given value.

### HasHiddenFromSearch

`func (o *PutPagesPageIdPage) HasHiddenFromSearch() bool`

HasHiddenFromSearch returns a boolean if a field has been set.

### GetViewersMustBeTeamMembers

`func (o *PutPagesPageIdPage) GetViewersMustBeTeamMembers() bool`

GetViewersMustBeTeamMembers returns the ViewersMustBeTeamMembers field if non-nil, zero value otherwise.

### GetViewersMustBeTeamMembersOk

`func (o *PutPagesPageIdPage) GetViewersMustBeTeamMembersOk() (*bool, bool)`

GetViewersMustBeTeamMembersOk returns a tuple with the ViewersMustBeTeamMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewersMustBeTeamMembers

`func (o *PutPagesPageIdPage) SetViewersMustBeTeamMembers(v bool)`

SetViewersMustBeTeamMembers sets ViewersMustBeTeamMembers field to given value.

### HasViewersMustBeTeamMembers

`func (o *PutPagesPageIdPage) HasViewersMustBeTeamMembers() bool`

HasViewersMustBeTeamMembers returns a boolean if a field has been set.

### GetAllowPageSubscribers

`func (o *PutPagesPageIdPage) GetAllowPageSubscribers() bool`

GetAllowPageSubscribers returns the AllowPageSubscribers field if non-nil, zero value otherwise.

### GetAllowPageSubscribersOk

`func (o *PutPagesPageIdPage) GetAllowPageSubscribersOk() (*bool, bool)`

GetAllowPageSubscribersOk returns a tuple with the AllowPageSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPageSubscribers

`func (o *PutPagesPageIdPage) SetAllowPageSubscribers(v bool)`

SetAllowPageSubscribers sets AllowPageSubscribers field to given value.

### HasAllowPageSubscribers

`func (o *PutPagesPageIdPage) HasAllowPageSubscribers() bool`

HasAllowPageSubscribers returns a boolean if a field has been set.

### GetAllowIncidentSubscribers

`func (o *PutPagesPageIdPage) GetAllowIncidentSubscribers() bool`

GetAllowIncidentSubscribers returns the AllowIncidentSubscribers field if non-nil, zero value otherwise.

### GetAllowIncidentSubscribersOk

`func (o *PutPagesPageIdPage) GetAllowIncidentSubscribersOk() (*bool, bool)`

GetAllowIncidentSubscribersOk returns a tuple with the AllowIncidentSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIncidentSubscribers

`func (o *PutPagesPageIdPage) SetAllowIncidentSubscribers(v bool)`

SetAllowIncidentSubscribers sets AllowIncidentSubscribers field to given value.

### HasAllowIncidentSubscribers

`func (o *PutPagesPageIdPage) HasAllowIncidentSubscribers() bool`

HasAllowIncidentSubscribers returns a boolean if a field has been set.

### GetAllowEmailSubscribers

`func (o *PutPagesPageIdPage) GetAllowEmailSubscribers() bool`

GetAllowEmailSubscribers returns the AllowEmailSubscribers field if non-nil, zero value otherwise.

### GetAllowEmailSubscribersOk

`func (o *PutPagesPageIdPage) GetAllowEmailSubscribersOk() (*bool, bool)`

GetAllowEmailSubscribersOk returns a tuple with the AllowEmailSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmailSubscribers

`func (o *PutPagesPageIdPage) SetAllowEmailSubscribers(v bool)`

SetAllowEmailSubscribers sets AllowEmailSubscribers field to given value.

### HasAllowEmailSubscribers

`func (o *PutPagesPageIdPage) HasAllowEmailSubscribers() bool`

HasAllowEmailSubscribers returns a boolean if a field has been set.

### GetAllowSmsSubscribers

`func (o *PutPagesPageIdPage) GetAllowSmsSubscribers() bool`

GetAllowSmsSubscribers returns the AllowSmsSubscribers field if non-nil, zero value otherwise.

### GetAllowSmsSubscribersOk

`func (o *PutPagesPageIdPage) GetAllowSmsSubscribersOk() (*bool, bool)`

GetAllowSmsSubscribersOk returns a tuple with the AllowSmsSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSmsSubscribers

`func (o *PutPagesPageIdPage) SetAllowSmsSubscribers(v bool)`

SetAllowSmsSubscribers sets AllowSmsSubscribers field to given value.

### HasAllowSmsSubscribers

`func (o *PutPagesPageIdPage) HasAllowSmsSubscribers() bool`

HasAllowSmsSubscribers returns a boolean if a field has been set.

### GetAllowRssAtomFeeds

`func (o *PutPagesPageIdPage) GetAllowRssAtomFeeds() bool`

GetAllowRssAtomFeeds returns the AllowRssAtomFeeds field if non-nil, zero value otherwise.

### GetAllowRssAtomFeedsOk

`func (o *PutPagesPageIdPage) GetAllowRssAtomFeedsOk() (*bool, bool)`

GetAllowRssAtomFeedsOk returns a tuple with the AllowRssAtomFeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRssAtomFeeds

`func (o *PutPagesPageIdPage) SetAllowRssAtomFeeds(v bool)`

SetAllowRssAtomFeeds sets AllowRssAtomFeeds field to given value.

### HasAllowRssAtomFeeds

`func (o *PutPagesPageIdPage) HasAllowRssAtomFeeds() bool`

HasAllowRssAtomFeeds returns a boolean if a field has been set.

### GetAllowWebhookSubscribers

`func (o *PutPagesPageIdPage) GetAllowWebhookSubscribers() bool`

GetAllowWebhookSubscribers returns the AllowWebhookSubscribers field if non-nil, zero value otherwise.

### GetAllowWebhookSubscribersOk

`func (o *PutPagesPageIdPage) GetAllowWebhookSubscribersOk() (*bool, bool)`

GetAllowWebhookSubscribersOk returns a tuple with the AllowWebhookSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWebhookSubscribers

`func (o *PutPagesPageIdPage) SetAllowWebhookSubscribers(v bool)`

SetAllowWebhookSubscribers sets AllowWebhookSubscribers field to given value.

### HasAllowWebhookSubscribers

`func (o *PutPagesPageIdPage) HasAllowWebhookSubscribers() bool`

HasAllowWebhookSubscribers returns a boolean if a field has been set.

### GetNotificationsFromEmail

`func (o *PutPagesPageIdPage) GetNotificationsFromEmail() string`

GetNotificationsFromEmail returns the NotificationsFromEmail field if non-nil, zero value otherwise.

### GetNotificationsFromEmailOk

`func (o *PutPagesPageIdPage) GetNotificationsFromEmailOk() (*string, bool)`

GetNotificationsFromEmailOk returns a tuple with the NotificationsFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsFromEmail

`func (o *PutPagesPageIdPage) SetNotificationsFromEmail(v string)`

SetNotificationsFromEmail sets NotificationsFromEmail field to given value.

### HasNotificationsFromEmail

`func (o *PutPagesPageIdPage) HasNotificationsFromEmail() bool`

HasNotificationsFromEmail returns a boolean if a field has been set.

### GetTimeZone

`func (o *PutPagesPageIdPage) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PutPagesPageIdPage) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PutPagesPageIdPage) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PutPagesPageIdPage) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetNotificationsEmailFooter

`func (o *PutPagesPageIdPage) GetNotificationsEmailFooter() string`

GetNotificationsEmailFooter returns the NotificationsEmailFooter field if non-nil, zero value otherwise.

### GetNotificationsEmailFooterOk

`func (o *PutPagesPageIdPage) GetNotificationsEmailFooterOk() (*string, bool)`

GetNotificationsEmailFooterOk returns a tuple with the NotificationsEmailFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEmailFooter

`func (o *PutPagesPageIdPage) SetNotificationsEmailFooter(v string)`

SetNotificationsEmailFooter sets NotificationsEmailFooter field to given value.

### HasNotificationsEmailFooter

`func (o *PutPagesPageIdPage) HasNotificationsEmailFooter() bool`

HasNotificationsEmailFooter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


