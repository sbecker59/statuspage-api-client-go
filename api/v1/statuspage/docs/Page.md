# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Page identifier | [optional] 
**CreatedAt** | Pointer to **time.Time** | Timestamp the record was created | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Timestamp the record was last updated | [optional] 
**Name** | Pointer to **string** | Name of your page to be displayed | [optional] 
**PageDescription** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Branding** | Pointer to **string** | The main template your statuspage will use | [optional] 
**Subdomain** | Pointer to **string** | Subdomain at which to access your status page | [optional] 
**Domain** | Pointer to **string** | CNAME alias for your status page | [optional] 
**Url** | Pointer to **string** | Website of your page.  Clicking on your statuspage image will link here. | [optional] 
**SupportUrl** | Pointer to **string** |  | [optional] 
**HiddenFromSearch** | Pointer to **bool** | Should your page hide itself from search engines | [optional] 
**AllowPageSubscribers** | Pointer to **bool** | Can your users subscribe to all notifications on the page | [optional] 
**AllowIncidentSubscribers** | Pointer to **bool** | Can your users subscribe to notifications for a single incident | [optional] 
**AllowEmailSubscribers** | Pointer to **bool** | Can your users choose to receive notifications via email | [optional] 
**AllowSmsSubscribers** | Pointer to **bool** | Can your users choose to receive notifications via SMS | [optional] 
**AllowRssAtomFeeds** | Pointer to **bool** | Can your users choose to access incident feeds via RSS/Atom | [optional] 
**AllowWebhookSubscribers** | Pointer to **bool** | Can your users choose to receive notifications via Webhooks | [optional] 
**NotificationsFromEmail** | Pointer to **string** | Allows you to customize the email address your page notifications come from | [optional] 
**NotificationsEmailFooter** | Pointer to **string** | Allows you to customize the footer appearing on your notification emails.  Accepts Markdown for formatting | [optional] 
**ActivityScore** | Pointer to **float32** |  | [optional] 
**TwitterUsername** | Pointer to **string** |  | [optional] 
**ViewersMustBeTeamMembers** | Pointer to **bool** |  | [optional] 
**IpRestrictions** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** | Timezone configured for your page | [optional] 
**CssBodyBackgroundColor** | Pointer to **string** | CSS Color | [optional] 
**CssFontColor** | Pointer to **string** | CSS Color | [optional] 
**CssLightFontColor** | Pointer to **string** | CSS Color | [optional] 
**CssGreens** | Pointer to **string** | CSS Color | [optional] 
**CssYellows** | Pointer to **string** | CSS Color | [optional] 
**CssOranges** | Pointer to **string** | CSS Color | [optional] 
**CssBlues** | Pointer to **string** | CSS Color | [optional] 
**CssReds** | Pointer to **string** | CSS Color | [optional] 
**CssBorderColor** | Pointer to **string** | CSS Color | [optional] 
**CssGraphColor** | Pointer to **string** | CSS Color | [optional] 
**CssLinkColor** | Pointer to **string** | CSS Color | [optional] 
**CssNoData** | Pointer to **string** | CSS Color | [optional] 
**FaviconLogo** | Pointer to **string** |  | [optional] 
**TransactionalLogo** | Pointer to **string** |  | [optional] 
**HeroCover** | Pointer to **string** |  | [optional] 
**EmailLogo** | Pointer to **string** |  | [optional] 
**TwitterLogo** | Pointer to **string** |  | [optional] 

## Methods

### NewPage

`func NewPage() *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Page) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Page) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Page) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Page) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Page) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Page) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Page) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Page) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Page) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Page) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Page) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Page) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *Page) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Page) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Page) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Page) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPageDescription

`func (o *Page) GetPageDescription() string`

GetPageDescription returns the PageDescription field if non-nil, zero value otherwise.

### GetPageDescriptionOk

`func (o *Page) GetPageDescriptionOk() (*string, bool)`

GetPageDescriptionOk returns a tuple with the PageDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageDescription

`func (o *Page) SetPageDescription(v string)`

SetPageDescription sets PageDescription field to given value.

### HasPageDescription

`func (o *Page) HasPageDescription() bool`

HasPageDescription returns a boolean if a field has been set.

### GetHeadline

`func (o *Page) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *Page) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *Page) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *Page) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetBranding

`func (o *Page) GetBranding() string`

GetBranding returns the Branding field if non-nil, zero value otherwise.

### GetBrandingOk

`func (o *Page) GetBrandingOk() (*string, bool)`

GetBrandingOk returns a tuple with the Branding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranding

`func (o *Page) SetBranding(v string)`

SetBranding sets Branding field to given value.

### HasBranding

`func (o *Page) HasBranding() bool`

HasBranding returns a boolean if a field has been set.

### GetSubdomain

`func (o *Page) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *Page) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *Page) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *Page) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetDomain

`func (o *Page) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Page) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Page) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Page) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUrl

`func (o *Page) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Page) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Page) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Page) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSupportUrl

`func (o *Page) GetSupportUrl() string`

GetSupportUrl returns the SupportUrl field if non-nil, zero value otherwise.

### GetSupportUrlOk

`func (o *Page) GetSupportUrlOk() (*string, bool)`

GetSupportUrlOk returns a tuple with the SupportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportUrl

`func (o *Page) SetSupportUrl(v string)`

SetSupportUrl sets SupportUrl field to given value.

### HasSupportUrl

`func (o *Page) HasSupportUrl() bool`

HasSupportUrl returns a boolean if a field has been set.

### GetHiddenFromSearch

`func (o *Page) GetHiddenFromSearch() bool`

GetHiddenFromSearch returns the HiddenFromSearch field if non-nil, zero value otherwise.

### GetHiddenFromSearchOk

`func (o *Page) GetHiddenFromSearchOk() (*bool, bool)`

GetHiddenFromSearchOk returns a tuple with the HiddenFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromSearch

`func (o *Page) SetHiddenFromSearch(v bool)`

SetHiddenFromSearch sets HiddenFromSearch field to given value.

### HasHiddenFromSearch

`func (o *Page) HasHiddenFromSearch() bool`

HasHiddenFromSearch returns a boolean if a field has been set.

### GetAllowPageSubscribers

`func (o *Page) GetAllowPageSubscribers() bool`

GetAllowPageSubscribers returns the AllowPageSubscribers field if non-nil, zero value otherwise.

### GetAllowPageSubscribersOk

`func (o *Page) GetAllowPageSubscribersOk() (*bool, bool)`

GetAllowPageSubscribersOk returns a tuple with the AllowPageSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPageSubscribers

`func (o *Page) SetAllowPageSubscribers(v bool)`

SetAllowPageSubscribers sets AllowPageSubscribers field to given value.

### HasAllowPageSubscribers

`func (o *Page) HasAllowPageSubscribers() bool`

HasAllowPageSubscribers returns a boolean if a field has been set.

### GetAllowIncidentSubscribers

`func (o *Page) GetAllowIncidentSubscribers() bool`

GetAllowIncidentSubscribers returns the AllowIncidentSubscribers field if non-nil, zero value otherwise.

### GetAllowIncidentSubscribersOk

`func (o *Page) GetAllowIncidentSubscribersOk() (*bool, bool)`

GetAllowIncidentSubscribersOk returns a tuple with the AllowIncidentSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowIncidentSubscribers

`func (o *Page) SetAllowIncidentSubscribers(v bool)`

SetAllowIncidentSubscribers sets AllowIncidentSubscribers field to given value.

### HasAllowIncidentSubscribers

`func (o *Page) HasAllowIncidentSubscribers() bool`

HasAllowIncidentSubscribers returns a boolean if a field has been set.

### GetAllowEmailSubscribers

`func (o *Page) GetAllowEmailSubscribers() bool`

GetAllowEmailSubscribers returns the AllowEmailSubscribers field if non-nil, zero value otherwise.

### GetAllowEmailSubscribersOk

`func (o *Page) GetAllowEmailSubscribersOk() (*bool, bool)`

GetAllowEmailSubscribersOk returns a tuple with the AllowEmailSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowEmailSubscribers

`func (o *Page) SetAllowEmailSubscribers(v bool)`

SetAllowEmailSubscribers sets AllowEmailSubscribers field to given value.

### HasAllowEmailSubscribers

`func (o *Page) HasAllowEmailSubscribers() bool`

HasAllowEmailSubscribers returns a boolean if a field has been set.

### GetAllowSmsSubscribers

`func (o *Page) GetAllowSmsSubscribers() bool`

GetAllowSmsSubscribers returns the AllowSmsSubscribers field if non-nil, zero value otherwise.

### GetAllowSmsSubscribersOk

`func (o *Page) GetAllowSmsSubscribersOk() (*bool, bool)`

GetAllowSmsSubscribersOk returns a tuple with the AllowSmsSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSmsSubscribers

`func (o *Page) SetAllowSmsSubscribers(v bool)`

SetAllowSmsSubscribers sets AllowSmsSubscribers field to given value.

### HasAllowSmsSubscribers

`func (o *Page) HasAllowSmsSubscribers() bool`

HasAllowSmsSubscribers returns a boolean if a field has been set.

### GetAllowRssAtomFeeds

`func (o *Page) GetAllowRssAtomFeeds() bool`

GetAllowRssAtomFeeds returns the AllowRssAtomFeeds field if non-nil, zero value otherwise.

### GetAllowRssAtomFeedsOk

`func (o *Page) GetAllowRssAtomFeedsOk() (*bool, bool)`

GetAllowRssAtomFeedsOk returns a tuple with the AllowRssAtomFeeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRssAtomFeeds

`func (o *Page) SetAllowRssAtomFeeds(v bool)`

SetAllowRssAtomFeeds sets AllowRssAtomFeeds field to given value.

### HasAllowRssAtomFeeds

`func (o *Page) HasAllowRssAtomFeeds() bool`

HasAllowRssAtomFeeds returns a boolean if a field has been set.

### GetAllowWebhookSubscribers

`func (o *Page) GetAllowWebhookSubscribers() bool`

GetAllowWebhookSubscribers returns the AllowWebhookSubscribers field if non-nil, zero value otherwise.

### GetAllowWebhookSubscribersOk

`func (o *Page) GetAllowWebhookSubscribersOk() (*bool, bool)`

GetAllowWebhookSubscribersOk returns a tuple with the AllowWebhookSubscribers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowWebhookSubscribers

`func (o *Page) SetAllowWebhookSubscribers(v bool)`

SetAllowWebhookSubscribers sets AllowWebhookSubscribers field to given value.

### HasAllowWebhookSubscribers

`func (o *Page) HasAllowWebhookSubscribers() bool`

HasAllowWebhookSubscribers returns a boolean if a field has been set.

### GetNotificationsFromEmail

`func (o *Page) GetNotificationsFromEmail() string`

GetNotificationsFromEmail returns the NotificationsFromEmail field if non-nil, zero value otherwise.

### GetNotificationsFromEmailOk

`func (o *Page) GetNotificationsFromEmailOk() (*string, bool)`

GetNotificationsFromEmailOk returns a tuple with the NotificationsFromEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsFromEmail

`func (o *Page) SetNotificationsFromEmail(v string)`

SetNotificationsFromEmail sets NotificationsFromEmail field to given value.

### HasNotificationsFromEmail

`func (o *Page) HasNotificationsFromEmail() bool`

HasNotificationsFromEmail returns a boolean if a field has been set.

### GetNotificationsEmailFooter

`func (o *Page) GetNotificationsEmailFooter() string`

GetNotificationsEmailFooter returns the NotificationsEmailFooter field if non-nil, zero value otherwise.

### GetNotificationsEmailFooterOk

`func (o *Page) GetNotificationsEmailFooterOk() (*string, bool)`

GetNotificationsEmailFooterOk returns a tuple with the NotificationsEmailFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsEmailFooter

`func (o *Page) SetNotificationsEmailFooter(v string)`

SetNotificationsEmailFooter sets NotificationsEmailFooter field to given value.

### HasNotificationsEmailFooter

`func (o *Page) HasNotificationsEmailFooter() bool`

HasNotificationsEmailFooter returns a boolean if a field has been set.

### GetActivityScore

`func (o *Page) GetActivityScore() float32`

GetActivityScore returns the ActivityScore field if non-nil, zero value otherwise.

### GetActivityScoreOk

`func (o *Page) GetActivityScoreOk() (*float32, bool)`

GetActivityScoreOk returns a tuple with the ActivityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityScore

`func (o *Page) SetActivityScore(v float32)`

SetActivityScore sets ActivityScore field to given value.

### HasActivityScore

`func (o *Page) HasActivityScore() bool`

HasActivityScore returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *Page) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *Page) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *Page) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *Page) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### GetViewersMustBeTeamMembers

`func (o *Page) GetViewersMustBeTeamMembers() bool`

GetViewersMustBeTeamMembers returns the ViewersMustBeTeamMembers field if non-nil, zero value otherwise.

### GetViewersMustBeTeamMembersOk

`func (o *Page) GetViewersMustBeTeamMembersOk() (*bool, bool)`

GetViewersMustBeTeamMembersOk returns a tuple with the ViewersMustBeTeamMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewersMustBeTeamMembers

`func (o *Page) SetViewersMustBeTeamMembers(v bool)`

SetViewersMustBeTeamMembers sets ViewersMustBeTeamMembers field to given value.

### HasViewersMustBeTeamMembers

`func (o *Page) HasViewersMustBeTeamMembers() bool`

HasViewersMustBeTeamMembers returns a boolean if a field has been set.

### GetIpRestrictions

`func (o *Page) GetIpRestrictions() string`

GetIpRestrictions returns the IpRestrictions field if non-nil, zero value otherwise.

### GetIpRestrictionsOk

`func (o *Page) GetIpRestrictionsOk() (*string, bool)`

GetIpRestrictionsOk returns a tuple with the IpRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRestrictions

`func (o *Page) SetIpRestrictions(v string)`

SetIpRestrictions sets IpRestrictions field to given value.

### HasIpRestrictions

`func (o *Page) HasIpRestrictions() bool`

HasIpRestrictions returns a boolean if a field has been set.

### GetCity

`func (o *Page) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Page) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Page) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Page) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Page) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Page) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Page) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Page) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCountry

`func (o *Page) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Page) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Page) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Page) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetTimeZone

`func (o *Page) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Page) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Page) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Page) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetCssBodyBackgroundColor

`func (o *Page) GetCssBodyBackgroundColor() string`

GetCssBodyBackgroundColor returns the CssBodyBackgroundColor field if non-nil, zero value otherwise.

### GetCssBodyBackgroundColorOk

`func (o *Page) GetCssBodyBackgroundColorOk() (*string, bool)`

GetCssBodyBackgroundColorOk returns a tuple with the CssBodyBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBodyBackgroundColor

`func (o *Page) SetCssBodyBackgroundColor(v string)`

SetCssBodyBackgroundColor sets CssBodyBackgroundColor field to given value.

### HasCssBodyBackgroundColor

`func (o *Page) HasCssBodyBackgroundColor() bool`

HasCssBodyBackgroundColor returns a boolean if a field has been set.

### GetCssFontColor

`func (o *Page) GetCssFontColor() string`

GetCssFontColor returns the CssFontColor field if non-nil, zero value otherwise.

### GetCssFontColorOk

`func (o *Page) GetCssFontColorOk() (*string, bool)`

GetCssFontColorOk returns a tuple with the CssFontColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssFontColor

`func (o *Page) SetCssFontColor(v string)`

SetCssFontColor sets CssFontColor field to given value.

### HasCssFontColor

`func (o *Page) HasCssFontColor() bool`

HasCssFontColor returns a boolean if a field has been set.

### GetCssLightFontColor

`func (o *Page) GetCssLightFontColor() string`

GetCssLightFontColor returns the CssLightFontColor field if non-nil, zero value otherwise.

### GetCssLightFontColorOk

`func (o *Page) GetCssLightFontColorOk() (*string, bool)`

GetCssLightFontColorOk returns a tuple with the CssLightFontColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLightFontColor

`func (o *Page) SetCssLightFontColor(v string)`

SetCssLightFontColor sets CssLightFontColor field to given value.

### HasCssLightFontColor

`func (o *Page) HasCssLightFontColor() bool`

HasCssLightFontColor returns a boolean if a field has been set.

### GetCssGreens

`func (o *Page) GetCssGreens() string`

GetCssGreens returns the CssGreens field if non-nil, zero value otherwise.

### GetCssGreensOk

`func (o *Page) GetCssGreensOk() (*string, bool)`

GetCssGreensOk returns a tuple with the CssGreens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssGreens

`func (o *Page) SetCssGreens(v string)`

SetCssGreens sets CssGreens field to given value.

### HasCssGreens

`func (o *Page) HasCssGreens() bool`

HasCssGreens returns a boolean if a field has been set.

### GetCssYellows

`func (o *Page) GetCssYellows() string`

GetCssYellows returns the CssYellows field if non-nil, zero value otherwise.

### GetCssYellowsOk

`func (o *Page) GetCssYellowsOk() (*string, bool)`

GetCssYellowsOk returns a tuple with the CssYellows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssYellows

`func (o *Page) SetCssYellows(v string)`

SetCssYellows sets CssYellows field to given value.

### HasCssYellows

`func (o *Page) HasCssYellows() bool`

HasCssYellows returns a boolean if a field has been set.

### GetCssOranges

`func (o *Page) GetCssOranges() string`

GetCssOranges returns the CssOranges field if non-nil, zero value otherwise.

### GetCssOrangesOk

`func (o *Page) GetCssOrangesOk() (*string, bool)`

GetCssOrangesOk returns a tuple with the CssOranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssOranges

`func (o *Page) SetCssOranges(v string)`

SetCssOranges sets CssOranges field to given value.

### HasCssOranges

`func (o *Page) HasCssOranges() bool`

HasCssOranges returns a boolean if a field has been set.

### GetCssBlues

`func (o *Page) GetCssBlues() string`

GetCssBlues returns the CssBlues field if non-nil, zero value otherwise.

### GetCssBluesOk

`func (o *Page) GetCssBluesOk() (*string, bool)`

GetCssBluesOk returns a tuple with the CssBlues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBlues

`func (o *Page) SetCssBlues(v string)`

SetCssBlues sets CssBlues field to given value.

### HasCssBlues

`func (o *Page) HasCssBlues() bool`

HasCssBlues returns a boolean if a field has been set.

### GetCssReds

`func (o *Page) GetCssReds() string`

GetCssReds returns the CssReds field if non-nil, zero value otherwise.

### GetCssRedsOk

`func (o *Page) GetCssRedsOk() (*string, bool)`

GetCssRedsOk returns a tuple with the CssReds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssReds

`func (o *Page) SetCssReds(v string)`

SetCssReds sets CssReds field to given value.

### HasCssReds

`func (o *Page) HasCssReds() bool`

HasCssReds returns a boolean if a field has been set.

### GetCssBorderColor

`func (o *Page) GetCssBorderColor() string`

GetCssBorderColor returns the CssBorderColor field if non-nil, zero value otherwise.

### GetCssBorderColorOk

`func (o *Page) GetCssBorderColorOk() (*string, bool)`

GetCssBorderColorOk returns a tuple with the CssBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssBorderColor

`func (o *Page) SetCssBorderColor(v string)`

SetCssBorderColor sets CssBorderColor field to given value.

### HasCssBorderColor

`func (o *Page) HasCssBorderColor() bool`

HasCssBorderColor returns a boolean if a field has been set.

### GetCssGraphColor

`func (o *Page) GetCssGraphColor() string`

GetCssGraphColor returns the CssGraphColor field if non-nil, zero value otherwise.

### GetCssGraphColorOk

`func (o *Page) GetCssGraphColorOk() (*string, bool)`

GetCssGraphColorOk returns a tuple with the CssGraphColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssGraphColor

`func (o *Page) SetCssGraphColor(v string)`

SetCssGraphColor sets CssGraphColor field to given value.

### HasCssGraphColor

`func (o *Page) HasCssGraphColor() bool`

HasCssGraphColor returns a boolean if a field has been set.

### GetCssLinkColor

`func (o *Page) GetCssLinkColor() string`

GetCssLinkColor returns the CssLinkColor field if non-nil, zero value otherwise.

### GetCssLinkColorOk

`func (o *Page) GetCssLinkColorOk() (*string, bool)`

GetCssLinkColorOk returns a tuple with the CssLinkColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssLinkColor

`func (o *Page) SetCssLinkColor(v string)`

SetCssLinkColor sets CssLinkColor field to given value.

### HasCssLinkColor

`func (o *Page) HasCssLinkColor() bool`

HasCssLinkColor returns a boolean if a field has been set.

### GetCssNoData

`func (o *Page) GetCssNoData() string`

GetCssNoData returns the CssNoData field if non-nil, zero value otherwise.

### GetCssNoDataOk

`func (o *Page) GetCssNoDataOk() (*string, bool)`

GetCssNoDataOk returns a tuple with the CssNoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssNoData

`func (o *Page) SetCssNoData(v string)`

SetCssNoData sets CssNoData field to given value.

### HasCssNoData

`func (o *Page) HasCssNoData() bool`

HasCssNoData returns a boolean if a field has been set.

### GetFaviconLogo

`func (o *Page) GetFaviconLogo() string`

GetFaviconLogo returns the FaviconLogo field if non-nil, zero value otherwise.

### GetFaviconLogoOk

`func (o *Page) GetFaviconLogoOk() (*string, bool)`

GetFaviconLogoOk returns a tuple with the FaviconLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaviconLogo

`func (o *Page) SetFaviconLogo(v string)`

SetFaviconLogo sets FaviconLogo field to given value.

### HasFaviconLogo

`func (o *Page) HasFaviconLogo() bool`

HasFaviconLogo returns a boolean if a field has been set.

### GetTransactionalLogo

`func (o *Page) GetTransactionalLogo() string`

GetTransactionalLogo returns the TransactionalLogo field if non-nil, zero value otherwise.

### GetTransactionalLogoOk

`func (o *Page) GetTransactionalLogoOk() (*string, bool)`

GetTransactionalLogoOk returns a tuple with the TransactionalLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionalLogo

`func (o *Page) SetTransactionalLogo(v string)`

SetTransactionalLogo sets TransactionalLogo field to given value.

### HasTransactionalLogo

`func (o *Page) HasTransactionalLogo() bool`

HasTransactionalLogo returns a boolean if a field has been set.

### GetHeroCover

`func (o *Page) GetHeroCover() string`

GetHeroCover returns the HeroCover field if non-nil, zero value otherwise.

### GetHeroCoverOk

`func (o *Page) GetHeroCoverOk() (*string, bool)`

GetHeroCoverOk returns a tuple with the HeroCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeroCover

`func (o *Page) SetHeroCover(v string)`

SetHeroCover sets HeroCover field to given value.

### HasHeroCover

`func (o *Page) HasHeroCover() bool`

HasHeroCover returns a boolean if a field has been set.

### GetEmailLogo

`func (o *Page) GetEmailLogo() string`

GetEmailLogo returns the EmailLogo field if non-nil, zero value otherwise.

### GetEmailLogoOk

`func (o *Page) GetEmailLogoOk() (*string, bool)`

GetEmailLogoOk returns a tuple with the EmailLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailLogo

`func (o *Page) SetEmailLogo(v string)`

SetEmailLogo sets EmailLogo field to given value.

### HasEmailLogo

`func (o *Page) HasEmailLogo() bool`

HasEmailLogo returns a boolean if a field has been set.

### GetTwitterLogo

`func (o *Page) GetTwitterLogo() string`

GetTwitterLogo returns the TwitterLogo field if non-nil, zero value otherwise.

### GetTwitterLogoOk

`func (o *Page) GetTwitterLogoOk() (*string, bool)`

GetTwitterLogoOk returns a tuple with the TwitterLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterLogo

`func (o *Page) SetTwitterLogo(v string)`

SetTwitterLogo sets TwitterLogo field to given value.

### HasTwitterLogo

`func (o *Page) HasTwitterLogo() bool`

HasTwitterLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


