# PatchPagesPage

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
**AllowRssAtomFeeds** | Pointer to **bool** | Can your users choose to access incident feeds via RSS/Atom | [optional] 
**AllowWebhookSubscribers** | Pointer to **bool** | Can your users choose to receive notifications via Webhooks | [optional] 
**NotificationsFromEmail** | Pointer to **string** | Allows you to customize the email address your page notifications come from | [optional] 
**TimeZone** | Pointer to **string** | Timezone configured for your page | [optional] 
**NotificationsEmailFooter** | Pointer to **string** | Allows you to customize the footer appearing on your notification emails.  Accepts Markdown for formatting | [optional] 

## Methods

### NewPatchPagesPage

`func NewPatchPagesPage() *PatchPagesPage`

NewPatchPagesPage instantiates a new PatchPagesPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPagesPageWithDefaults

`func NewPatchPagesPageWithDefaults() *PatchPagesPage`

NewPatchPagesPageWithDefaults instantiates a new PatchPagesPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchPagesPage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchPagesPage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchPagesPage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchPagesPage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDomain

`func (o *PatchPagesPage) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PatchPagesPage) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PatchPagesPage) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PatchPagesPage) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetSubdomain

`func (o *PatchPagesPage) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *PatchPagesPage) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *PatchPagesPage) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *PatchPagesPage) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetUrl

`func (o *PatchPagesPage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PatchPagesPage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PatchPagesPage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PatchPagesPage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBranding

`func (o *PatchPagesPage) GetBranding() string`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *PatchPagesPage) GetBrandingOk() (*string, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *PatchPagesPage) SetBranding(v string)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *PatchPagesPage) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetCssBodyBackgroundColor

`func (o *PatchPagesPage) GetCssBodyBackgroundColor() string`

GetCssBodyBackgroundColor returns the CssBodyBackgroundColor field if non-nil, zero value otherwise.

### GetCssBodyBackgroundColorOk

`func (o *PatchPagesPage) GetCssBodyBackgroundColorOk() (*string, bool)`

GetCssBodyBackgroundColorOk returns a tuple with the CssBodyBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBodyBackgroundColor

`func (o *PatchPagesPage) SetCssBodyBackgroundColor(v string)`

SetCssBodyBackgroundColor sets CssBodyBackgroundColor field to given value.

### HasCssBodyBackgroundColor

`func (o *PatchPagesPage) HasCssBodyBackgroundColor() bool`

HasCssBodyBackgroundColor returns a boolean if a field has been set.

### GetCssFontColor

`func (o *PatchPagesPage) GetCssFontColor() string`

GetCssFontColor returns the CssFontColor field if non-nil, zero value otherwise.

### GetCssFontColorOk

`func (o *PatchPagesPage) GetCssFontColorOk() (*string, bool)`

GetCssFontColorOk returns a tuple with the CssFontColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFontColor

`func (o *PatchPagesPage) SetCssFontColor(v string)`

SetCssFontColor sets CssFontColor field to given value.

### HasCssFontColor

`func (o *PatchPagesPage) HasCssFontColor() bool`

HasCssFontColor returns a boolean if a field has been set.

### GetCssLightFontColor

`func (o *PatchPagesPage) GetCssLightFontColor() string`

GetCssLightFontColor returns the CssLightFontColor field if non-nil, zero value otherwise.

### GetCssLightFontColorOk

`func (o *PatchPagesPage) GetCssLightFontColorOk() (*string, bool)`

GetCssLightFontColorOk returns a tuple with the CssLightFontColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLightFontColor

`func (o *PatchPagesPage) SetCssLightFontColor(v string)`

SetCssLightFontColor sets CssLightFontColor field to given value.

### HasCssLightFontColor

`func (o *PatchPagesPage) HasCssLightFontColor() bool`

HasCssLightFontColor returns a boolean if a field has been set.

### GetCssGreens

`func (o *PatchPagesPage) GetCssGreens() string`

GetCssGreens returns the CssGreens field if non-nil, zero value otherwise.

### GetCssGreensOk

`func (o *PatchPagesPage) GetCssGreensOk() (*string, bool)`

GetCssGreensOk returns a tuple with the CssGreens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssGreens

`func (o *PatchPagesPage) SetCssGreens(v string)`

SetCssGreens sets CssGreens field to given value.

### HasCssGreens

`func (o *PatchPagesPage) HasCssGreens() bool`

HasCssGreens returns a boolean if a field has been set.

### GetCssYellows

`func (o *PatchPagesPage) GetCssYellows() string`

GetCssYellows returns the CssYellows field if non-nil, zero value otherwise.

### GetCssYellowsOk

`func (o *PatchPagesPage) GetCssYellowsOk() (*string, bool)`

GetCssYellowsOk returns a tuple with the CssYellows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssYellows

`func (o *PatchPagesPage) SetCssYellows(v string)`

SetCssYellows sets CssYellows field to given value.

### HasCssYellows

`func (o *PatchPagesPage) HasCssYellows() bool`

HasCssYellows returns a boolean if a field has been set.

### GetCssOranges

`func (o *PatchPagesPage) GetCssOranges() string`

GetCssOranges returns the CssOranges field if non-nil, zero value otherwise.

### GetCssOrangesOk

`func (o *PatchPagesPage) GetCssOrangesOk() (*string, bool)`

GetCssOrangesOk returns a tuple with the CssOranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssOranges

`func (o *PatchPagesPage) SetCssOranges(v string)`

SetCssOranges sets CssOranges field to given value.

### HasCssOranges

`func (o *PatchPagesPage) HasCssOranges() bool`

HasCssOranges returns a boolean if a field has been set.

### GetCssReds

`func (o *PatchPagesPage) GetCssReds() string`

GetCssReds returns the CssReds field if non-nil, zero value otherwise.

### GetCssRedsOk

`func (o *PatchPagesPage) GetCssRedsOk() (*string, bool)`

GetCssRedsOk returns a tuple with the CssReds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssReds

`func (o *PatchPagesPage) SetCssReds(v string)`

SetCssReds sets CssReds field to given value.

### HasCssReds

`func (o *PatchPagesPage) HasCssReds() bool`

HasCssReds returns a boolean if a field has been set.

### GetCssBlues

`func (o *PatchPagesPage) GetCssBlues() string`

GetCssBlues returns the CssBlues field if non-nil, zero value otherwise.

### GetCssBluesOk

`func (o *PatchPagesPage) GetCssBluesOk() (*string, bool)`

GetCssBluesOk returns a tuple with the CssBlues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBlues

`func (o *PatchPagesPage) SetCssBlues(v string)`

SetCssBlues sets CssBlues field to given value.

### HasCssBlues

`func (o *PatchPagesPage) HasCssBlues() bool`

HasCssBlues returns a boolean if a field has been set.

### GetCssBorderColor

`func (o *PatchPagesPage) GetCssBorderColor() string`

GetCssBorderColor returns the CssBorderColor field if non-nil, zero value otherwise.

### GetCssBorderColorOk

`func (o *PatchPagesPage) GetCssBorderColorOk() (*string, bool)`

GetCssBorderColorOk returns a tuple with the CssBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBorderColor

`func (o *PatchPagesPage) SetCssBorderColor(v string)`

SetCssBorderColor sets CssBorderColor field to given value.

### HasCssBorderColor

`func (o *PatchPagesPage) HasCssBorderColor() bool`

HasCssBorderColor returns a boolean if a field has been set.

### GetCssGraphColor

`func (o *PatchPagesPage) GetCssGraphColor() string`

GetCssGraphColor returns the CssGraphColor field if non-nil, zero value otherwise.

### GetCssGraphColorOk

`func (o *PatchPagesPage) GetCssGraphColorOk() (*string, bool)`

GetCssGraphColorOk returns a tuple with the CssGraphColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssGraphColor

`func (o *PatchPagesPage) SetCssGraphColor(v string)`

SetCssGraphColor sets CssGraphColor field to given value.

### HasCssGraphColor

`func (o *PatchPagesPage) HasCssGraphColor() bool`

HasCssGraphColor returns a boolean if a field has been set.

### GetCssLinkColor

`func (o *PatchPagesPage) GetCssLinkColor() string`

GetCssLinkColor returns the CssLinkColor field if non-nil, zero value otherwise.

### GetCssLinkColorOk

`func (o *PatchPagesPage) GetCssLinkColorOk() (*string, bool)`

GetCssLinkColorOk returns a tuple with the CssLinkColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLinkColor

`func (o *PatchPagesPage) SetCssLinkColor(v string)`

SetCssLinkColor sets CssLinkColor field to given value.

### HasCssLinkColor

`func (o *PatchPagesPage) HasCssLinkColor() bool`

HasCssLinkColor returns a boolean if a field has been set.

### GetCssNoData

`func (o *PatchPagesPage) GetCssNoData() string`

GetCssNoData returns the CssNoData field if non-nil, zero value otherwise.

### GetCssNoDataOk

`func (o *PatchPagesPage) GetCssNoDataOk() (*string, bool)`

GetCssNoDataOk returns a tuple with the CssNoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssNoData

`func (o *PatchPagesPage) SetCssNoData(v string)`

SetCssNoData sets CssNoData field to given value.

### HasCssNoData

`func (o *PatchPagesPage) HasCssNoData() bool`

HasCssNoData returns a boolean if a field has been set.

### GetHiddenFromSearch

`func (o *PatchPagesPage) GetHiddenFromSearch() bool`

GetHiddenFromSearch returns the HiddenFromSearch field if non-nil, zero value otherwise.

### GetHiddenFromSearchOk

`func (o *PatchPagesPage) GetHiddenFromSearchOk() (*bool, bool)`

GetHiddenFromSearchOk returns a tuple with the HiddenFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromSearch

`func (o *PatchPagesPage) SetHiddenFromSearch(v bool)`

SetHiddenFromSearch sets HiddenFromSearch field to given value.

### HasHiddenFromSearch

`func (o *PatchPagesPage) HasHiddenFromSearch() bool`

HasHiddenFromSearch returns a boolean if a field has been set.

### GetViewersMustBeTeamMembers

`func (o *PatchPagesPage) GetViewersMustBeTeamMembers() bool`

GetViewersMustBeTeamMembers returns the ViewersMustBeTeamMembers field if non-nil, zero value otherwise.

### GetViewersMustBeTeamMembersOk

`func (o *PatchPagesPage) GetViewersMustBeTeamMembersOk() (*bool, bool)`

GetViewersMustBeTeamMembersOk returns a tuple with the ViewersMustBeTeamMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewersMustBeTeamMembers

`func (o *PatchPagesPage) SetViewersMustBeTeamMembers(v bool)`

SetViewersMustBeTeamMembers sets ViewersMustBeTeamMembers field to given value.

### HasViewersMustBeTeamMembers

`func (o *PatchPagesPage) HasViewersMustBeTeamMembers() bool`

HasViewersMustBeTeamMembers returns a boolean if a field has been set.

### GetAllowPageSubscribers

`func (o *PatchPagesPage) GetAllowPageSubscribers() bool`

GetAllowPageSubscribers returns the AllowPageSubscribers field if non-nil, zero value otherwise.

### GetAllowPageSubscribersOk

`func (o *PatchPagesPage) GetAllowPageSubscribersOk() (*bool, bool)`

GetAllowPageSubscribersOk returns a tuple with the AllowPageSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPageSubscribers

`func (o *PatchPagesPage) SetAllowPageSubscribers(v bool)`

SetAllowPageSubscribers sets AllowPageSubscribers field to given value.

### HasAllowPageSubscribers

`func (o *PatchPagesPage) HasAllowPageSubscribers() bool`

HasAllowPageSubscribers returns a boolean if a field has been set.

### GetAllowIncidentSubscribers

`func (o *PatchPagesPage) GetAllowIncidentSubscribers() bool`

GetAllowIncidentSubscribers returns the AllowIncidentSubscribers field if non-nil, zero value otherwise.

### GetAllowIncidentSubscribersOk

`func (o *PatchPagesPage) GetAllowIncidentSubscribersOk() (*bool, bool)`

GetAllowIncidentSubscribersOk returns a tuple with the AllowIncidentSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIncidentSubscribers

`func (o *PatchPagesPage) SetAllowIncidentSubscribers(v bool)`

SetAllowIncidentSubscribers sets AllowIncidentSubscribers field to given value.

### HasAllowIncidentSubscribers

`func (o *PatchPagesPage) HasAllowIncidentSubscribers() bool`

HasAllowIncidentSubscribers returns a boolean if a field has been set.

### GetAllowEmailSubscribers

`func (o *PatchPagesPage) GetAllowEmailSubscribers() bool`

GetAllowEmailSubscribers returns the AllowEmailSubscribers field if non-nil, zero value otherwise.

### GetAllowEmailSubscribersOk

`func (o *PatchPagesPage) GetAllowEmailSubscribersOk() (*bool, bool)`

GetAllowEmailSubscribersOk returns a tuple with the AllowEmailSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmailSubscribers

`func (o *PatchPagesPage) SetAllowEmailSubscribers(v bool)`

SetAllowEmailSubscribers sets AllowEmailSubscribers field to given value.

### HasAllowEmailSubscribers

`func (o *PatchPagesPage) HasAllowEmailSubscribers() bool`

HasAllowEmailSubscribers returns a boolean if a field has been set.

### GetAllowSmsSubscribers

`func (o *PatchPagesPage) GetAllowSmsSubscribers() bool`

GetAllowSmsSubscribers returns the AllowSmsSubscribers field if non-nil, zero value otherwise.

### GetAllowSmsSubscribersOk

`func (o *PatchPagesPage) GetAllowSmsSubscribersOk() (*bool, bool)`

GetAllowSmsSubscribersOk returns a tuple with the AllowSmsSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSmsSubscribers

`func (o *PatchPagesPage) SetAllowSmsSubscribers(v bool)`

SetAllowSmsSubscribers sets AllowSmsSubscribers field to given value.

### HasAllowSmsSubscribers

`func (o *PatchPagesPage) HasAllowSmsSubscribers() bool`

HasAllowSmsSubscribers returns a boolean if a field has been set.

### GetAllowRssAtomFeeds

`func (o *PatchPagesPage) GetAllowRssAtomFeeds() bool`

GetAllowRssAtomFeeds returns the AllowRssAtomFeeds field if non-nil, zero value otherwise.

### GetAllowRssAtomFeedsOk

`func (o *PatchPagesPage) GetAllowRssAtomFeedsOk() (*bool, bool)`

GetAllowRssAtomFeedsOk returns a tuple with the AllowRssAtomFeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRssAtomFeeds

`func (o *PatchPagesPage) SetAllowRssAtomFeeds(v bool)`

SetAllowRssAtomFeeds sets AllowRssAtomFeeds field to given value.

### HasAllowRssAtomFeeds

`func (o *PatchPagesPage) HasAllowRssAtomFeeds() bool`

HasAllowRssAtomFeeds returns a boolean if a field has been set.

### GetAllowWebhookSubscribers

`func (o *PatchPagesPage) GetAllowWebhookSubscribers() bool`

GetAllowWebhookSubscribers returns the AllowWebhookSubscribers field if non-nil, zero value otherwise.

### GetAllowWebhookSubscribersOk

`func (o *PatchPagesPage) GetAllowWebhookSubscribersOk() (*bool, bool)`

GetAllowWebhookSubscribersOk returns a tuple with the AllowWebhookSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWebhookSubscribers

`func (o *PatchPagesPage) SetAllowWebhookSubscribers(v bool)`

SetAllowWebhookSubscribers sets AllowWebhookSubscribers field to given value.

### HasAllowWebhookSubscribers

`func (o *PatchPagesPage) HasAllowWebhookSubscribers() bool`

HasAllowWebhookSubscribers returns a boolean if a field has been set.

### GetNotificationsFromEmail

`func (o *PatchPagesPage) GetNotificationsFromEmail() string`

GetNotificationsFromEmail returns the NotificationsFromEmail field if non-nil, zero value otherwise.

### GetNotificationsFromEmailOk

`func (o *PatchPagesPage) GetNotificationsFromEmailOk() (*string, bool)`

GetNotificationsFromEmailOk returns a tuple with the NotificationsFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsFromEmail

`func (o *PatchPagesPage) SetNotificationsFromEmail(v string)`

SetNotificationsFromEmail sets NotificationsFromEmail field to given value.

### HasNotificationsFromEmail

`func (o *PatchPagesPage) HasNotificationsFromEmail() bool`

HasNotificationsFromEmail returns a boolean if a field has been set.

### GetTimeZone

`func (o *PatchPagesPage) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *PatchPagesPage) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *PatchPagesPage) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *PatchPagesPage) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetNotificationsEmailFooter

`func (o *PatchPagesPage) GetNotificationsEmailFooter() string`

GetNotificationsEmailFooter returns the NotificationsEmailFooter field if non-nil, zero value otherwise.

### GetNotificationsEmailFooterOk

`func (o *PatchPagesPage) GetNotificationsEmailFooterOk() (*string, bool)`

GetNotificationsEmailFooterOk returns a tuple with the NotificationsEmailFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEmailFooter

`func (o *PatchPagesPage) SetNotificationsEmailFooter(v string)`

SetNotificationsEmailFooter sets NotificationsEmailFooter field to given value.

### HasNotificationsEmailFooter

`func (o *PatchPagesPage) HasNotificationsEmailFooter() bool`

HasNotificationsEmailFooter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


