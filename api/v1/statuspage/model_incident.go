/*
 * Statuspage API
 *
 * # Code of Conduct Please don't abuse the API, and please report all feature requests and issues to https://help.statuspage.io/help/contact-us-30  # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://help.statuspage.io/help/contact-us-30  # Basics  ## HTTPS It's required  ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix:    https://api.statuspage.io/v1/  ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation.  * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element  ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry.  All examples are provided in JSON format, however they can easily be converted to form encoding if required.  Some examples of how to convert things are below:      // JSON     {       \"incident\": {         \"name\": \"test incident\",         \"components\": [\"8kbf7d35c070\", \"vtnh60py4yd7\"]       }     }      // Form Encoded (using curl as an example):     curl -X POST https://api.statuspage.io/v1/example \\       -d \"incident[name]=test incident\" \\       -d \"incident[components][]=8kbf7d35c070\" \\       -d \"incident[components][]=vtnh60py4yd7\"  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// Incident Get an incident
type Incident struct {
	// Incident Identifier
	Id *string `json:"id,omitempty"`
	// Incident components
	Components *[]Component `json:"components,omitempty"`
	// The timestamp when the incident was created at.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The impact of the incident.
	Impact *string `json:"impact,omitempty"`
	// value to override calculated impact value
	ImpactOverride *string `json:"impact_override,omitempty"`
	// The incident updates for incident.
	IncidentUpdates *[]IncidentUpdate `json:"incident_updates,omitempty"`
	// Metadata attached to the incident. Top level values must be objects.
	Metadata *Object `json:"metadata,omitempty"`
	// The timestamp when incident entered monitoring state.
	MonitoringAt *time.Time `json:"monitoring_at,omitempty"`
	// Incident Name
	Name *string `json:"name,omitempty"`
	// Incident Page Identifier
	PageId *string `json:"page_id,omitempty"`
	// Body of the Postmortem.
	PostmortemBody *string `json:"postmortem_body,omitempty"`
	// The timestamp when the incident postmortem body was last updated at.
	PostmortemBodyLastUpdatedAt *time.Time `json:"postmortem_body_last_updated_at,omitempty"`
	// Controls whether the incident will have postmortem.
	PostmortemIgnored *bool `json:"postmortem_ignored,omitempty"`
	// Indicates whether subscribers are already notificed about postmortem.
	PostmortemNotifiedSubscribers *bool `json:"postmortem_notified_subscribers,omitempty"`
	// Controls whether to decide if notify postmortem on twitter.
	PostmortemNotifiedTwitter *bool `json:"postmortem_notified_twitter,omitempty"`
	// The timestamp when the postmortem was published.
	PostmortemPublishedAt *bool `json:"postmortem_published_at,omitempty"`
	// The timestamp when incident was resolved.
	ResolvedAt *time.Time `json:"resolved_at,omitempty"`
	// Controls whether the incident is scheduled to automatically change to complete.
	ScheduledAutoCompleted *bool `json:"scheduled_auto_completed,omitempty"`
	// Controls whether the incident is scheduled to automatically change to in progress.
	ScheduledAutoInProgress *bool `json:"scheduled_auto_in_progress,omitempty"`
	// The timestamp the incident is scheduled for.
	ScheduledFor *time.Time `json:"scheduled_for,omitempty"`
	// Controls whether to remind subscribers prior to scheduled incidents.
	ScheduledRemindPrior *bool `json:"scheduled_remind_prior,omitempty"`
	// The timestamp when the scheduled incident reminder was sent at.
	ScheduledRemindedAt *time.Time `json:"scheduled_reminded_at,omitempty"`
	// The timestamp the incident is scheduled until.
	ScheduledUntil *time.Time `json:"scheduled_until,omitempty"`
	// Incident Shortlink.
	Shortlink *string `json:"shortlink,omitempty"`
	// The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed.
	Status *string `json:"status,omitempty"`
	// The timestamp when the incident was updated at.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewIncident instantiates a new Incident object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncident() *Incident {
	this := Incident{}
	return &this
}

// NewIncidentWithDefaults instantiates a new Incident object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncidentWithDefaults() *Incident {
	this := Incident{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Incident) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Incident) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Incident) SetId(v string) {
	o.Id = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *Incident) GetComponents() []Component {
	if o == nil || o.Components == nil {
		var ret []Component
		return ret
	}
	return *o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetComponentsOk() (*[]Component, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *Incident) HasComponents() bool {
	if o != nil && o.Components != nil {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []Component and assigns it to the Components field.
func (o *Incident) SetComponents(v []Component) {
	o.Components = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Incident) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Incident) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Incident) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *Incident) GetImpact() string {
	if o == nil || o.Impact == nil {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetImpactOk() (*string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *Incident) HasImpact() bool {
	if o != nil && o.Impact != nil {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *Incident) SetImpact(v string) {
	o.Impact = &v
}

// GetImpactOverride returns the ImpactOverride field value if set, zero value otherwise.
func (o *Incident) GetImpactOverride() string {
	if o == nil || o.ImpactOverride == nil {
		var ret string
		return ret
	}
	return *o.ImpactOverride
}

// GetImpactOverrideOk returns a tuple with the ImpactOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetImpactOverrideOk() (*string, bool) {
	if o == nil || o.ImpactOverride == nil {
		return nil, false
	}
	return o.ImpactOverride, true
}

// HasImpactOverride returns a boolean if a field has been set.
func (o *Incident) HasImpactOverride() bool {
	if o != nil && o.ImpactOverride != nil {
		return true
	}

	return false
}

// SetImpactOverride gets a reference to the given string and assigns it to the ImpactOverride field.
func (o *Incident) SetImpactOverride(v string) {
	o.ImpactOverride = &v
}

// GetIncidentUpdates returns the IncidentUpdates field value if set, zero value otherwise.
func (o *Incident) GetIncidentUpdates() []IncidentUpdate {
	if o == nil || o.IncidentUpdates == nil {
		var ret []IncidentUpdate
		return ret
	}
	return *o.IncidentUpdates
}

// GetIncidentUpdatesOk returns a tuple with the IncidentUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetIncidentUpdatesOk() (*[]IncidentUpdate, bool) {
	if o == nil || o.IncidentUpdates == nil {
		return nil, false
	}
	return o.IncidentUpdates, true
}

// HasIncidentUpdates returns a boolean if a field has been set.
func (o *Incident) HasIncidentUpdates() bool {
	if o != nil && o.IncidentUpdates != nil {
		return true
	}

	return false
}

// SetIncidentUpdates gets a reference to the given []IncidentUpdate and assigns it to the IncidentUpdates field.
func (o *Incident) SetIncidentUpdates(v []IncidentUpdate) {
	o.IncidentUpdates = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Incident) GetMetadata() Object {
	if o == nil || o.Metadata == nil {
		var ret Object
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetMetadataOk() (*Object, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Incident) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given Object and assigns it to the Metadata field.
func (o *Incident) SetMetadata(v Object) {
	o.Metadata = &v
}

// GetMonitoringAt returns the MonitoringAt field value if set, zero value otherwise.
func (o *Incident) GetMonitoringAt() time.Time {
	if o == nil || o.MonitoringAt == nil {
		var ret time.Time
		return ret
	}
	return *o.MonitoringAt
}

// GetMonitoringAtOk returns a tuple with the MonitoringAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetMonitoringAtOk() (*time.Time, bool) {
	if o == nil || o.MonitoringAt == nil {
		return nil, false
	}
	return o.MonitoringAt, true
}

// HasMonitoringAt returns a boolean if a field has been set.
func (o *Incident) HasMonitoringAt() bool {
	if o != nil && o.MonitoringAt != nil {
		return true
	}

	return false
}

// SetMonitoringAt gets a reference to the given time.Time and assigns it to the MonitoringAt field.
func (o *Incident) SetMonitoringAt(v time.Time) {
	o.MonitoringAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Incident) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Incident) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Incident) SetName(v string) {
	o.Name = &v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *Incident) GetPageId() string {
	if o == nil || o.PageId == nil {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetPageIdOk() (*string, bool) {
	if o == nil || o.PageId == nil {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *Incident) HasPageId() bool {
	if o != nil && o.PageId != nil {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *Incident) SetPageId(v string) {
	o.PageId = &v
}

// GetPostmortemBody returns the PostmortemBody field value if set, zero value otherwise.
func (o *Incident) GetPostmortemBody() string {
	if o == nil || o.PostmortemBody == nil {
		var ret string
		return ret
	}
	return *o.PostmortemBody
}

// GetPostmortemBodyOk returns a tuple with the PostmortemBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetPostmortemBodyOk() (*string, bool) {
	if o == nil || o.PostmortemBody == nil {
		return nil, false
	}
	return o.PostmortemBody, true
}

// HasPostmortemBody returns a boolean if a field has been set.
func (o *Incident) HasPostmortemBody() bool {
	if o != nil && o.PostmortemBody != nil {
		return true
	}

	return false
}

// SetPostmortemBody gets a reference to the given string and assigns it to the PostmortemBody field.
func (o *Incident) SetPostmortemBody(v string) {
	o.PostmortemBody = &v
}

// GetPostmortemBodyLastUpdatedAt returns the PostmortemBodyLastUpdatedAt field value if set, zero value otherwise.
func (o *Incident) GetPostmortemBodyLastUpdatedAt() time.Time {
	if o == nil || o.PostmortemBodyLastUpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PostmortemBodyLastUpdatedAt
}

// GetPostmortemBodyLastUpdatedAtOk returns a tuple with the PostmortemBodyLastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetPostmortemBodyLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.PostmortemBodyLastUpdatedAt == nil {
		return nil, false
	}
	return o.PostmortemBodyLastUpdatedAt, true
}

// HasPostmortemBodyLastUpdatedAt returns a boolean if a field has been set.
func (o *Incident) HasPostmortemBodyLastUpdatedAt() bool {
	if o != nil && o.PostmortemBodyLastUpdatedAt != nil {
		return true
	}

	return false
}

// SetPostmortemBodyLastUpdatedAt gets a reference to the given time.Time and assigns it to the PostmortemBodyLastUpdatedAt field.
func (o *Incident) SetPostmortemBodyLastUpdatedAt(v time.Time) {
	o.PostmortemBodyLastUpdatedAt = &v
}

// GetPostmortemIgnored returns the PostmortemIgnored field value if set, zero value otherwise.
func (o *Incident) GetPostmortemIgnored() bool {
	if o == nil || o.PostmortemIgnored == nil {
		var ret bool
		return ret
	}
	return *o.PostmortemIgnored
}

// GetPostmortemIgnoredOk returns a tuple with the PostmortemIgnored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetPostmortemIgnoredOk() (*bool, bool) {
	if o == nil || o.PostmortemIgnored == nil {
		return nil, false
	}
	return o.PostmortemIgnored, true
}

// HasPostmortemIgnored returns a boolean if a field has been set.
func (o *Incident) HasPostmortemIgnored() bool {
	if o != nil && o.PostmortemIgnored != nil {
		return true
	}

	return false
}

// SetPostmortemIgnored gets a reference to the given bool and assigns it to the PostmortemIgnored field.
func (o *Incident) SetPostmortemIgnored(v bool) {
	o.PostmortemIgnored = &v
}

// GetPostmortemNotifiedSubscribers returns the PostmortemNotifiedSubscribers field value if set, zero value otherwise.
func (o *Incident) GetPostmortemNotifiedSubscribers() bool {
	if o == nil || o.PostmortemNotifiedSubscribers == nil {
		var ret bool
		return ret
	}
	return *o.PostmortemNotifiedSubscribers
}

// GetPostmortemNotifiedSubscribersOk returns a tuple with the PostmortemNotifiedSubscribers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetPostmortemNotifiedSubscribersOk() (*bool, bool) {
	if o == nil || o.PostmortemNotifiedSubscribers == nil {
		return nil, false
	}
	return o.PostmortemNotifiedSubscribers, true
}

// HasPostmortemNotifiedSubscribers returns a boolean if a field has been set.
func (o *Incident) HasPostmortemNotifiedSubscribers() bool {
	if o != nil && o.PostmortemNotifiedSubscribers != nil {
		return true
	}

	return false
}

// SetPostmortemNotifiedSubscribers gets a reference to the given bool and assigns it to the PostmortemNotifiedSubscribers field.
func (o *Incident) SetPostmortemNotifiedSubscribers(v bool) {
	o.PostmortemNotifiedSubscribers = &v
}

// GetPostmortemNotifiedTwitter returns the PostmortemNotifiedTwitter field value if set, zero value otherwise.
func (o *Incident) GetPostmortemNotifiedTwitter() bool {
	if o == nil || o.PostmortemNotifiedTwitter == nil {
		var ret bool
		return ret
	}
	return *o.PostmortemNotifiedTwitter
}

// GetPostmortemNotifiedTwitterOk returns a tuple with the PostmortemNotifiedTwitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetPostmortemNotifiedTwitterOk() (*bool, bool) {
	if o == nil || o.PostmortemNotifiedTwitter == nil {
		return nil, false
	}
	return o.PostmortemNotifiedTwitter, true
}

// HasPostmortemNotifiedTwitter returns a boolean if a field has been set.
func (o *Incident) HasPostmortemNotifiedTwitter() bool {
	if o != nil && o.PostmortemNotifiedTwitter != nil {
		return true
	}

	return false
}

// SetPostmortemNotifiedTwitter gets a reference to the given bool and assigns it to the PostmortemNotifiedTwitter field.
func (o *Incident) SetPostmortemNotifiedTwitter(v bool) {
	o.PostmortemNotifiedTwitter = &v
}

// GetPostmortemPublishedAt returns the PostmortemPublishedAt field value if set, zero value otherwise.
func (o *Incident) GetPostmortemPublishedAt() bool {
	if o == nil || o.PostmortemPublishedAt == nil {
		var ret bool
		return ret
	}
	return *o.PostmortemPublishedAt
}

// GetPostmortemPublishedAtOk returns a tuple with the PostmortemPublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetPostmortemPublishedAtOk() (*bool, bool) {
	if o == nil || o.PostmortemPublishedAt == nil {
		return nil, false
	}
	return o.PostmortemPublishedAt, true
}

// HasPostmortemPublishedAt returns a boolean if a field has been set.
func (o *Incident) HasPostmortemPublishedAt() bool {
	if o != nil && o.PostmortemPublishedAt != nil {
		return true
	}

	return false
}

// SetPostmortemPublishedAt gets a reference to the given bool and assigns it to the PostmortemPublishedAt field.
func (o *Incident) SetPostmortemPublishedAt(v bool) {
	o.PostmortemPublishedAt = &v
}

// GetResolvedAt returns the ResolvedAt field value if set, zero value otherwise.
func (o *Incident) GetResolvedAt() time.Time {
	if o == nil || o.ResolvedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ResolvedAt
}

// GetResolvedAtOk returns a tuple with the ResolvedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetResolvedAtOk() (*time.Time, bool) {
	if o == nil || o.ResolvedAt == nil {
		return nil, false
	}
	return o.ResolvedAt, true
}

// HasResolvedAt returns a boolean if a field has been set.
func (o *Incident) HasResolvedAt() bool {
	if o != nil && o.ResolvedAt != nil {
		return true
	}

	return false
}

// SetResolvedAt gets a reference to the given time.Time and assigns it to the ResolvedAt field.
func (o *Incident) SetResolvedAt(v time.Time) {
	o.ResolvedAt = &v
}

// GetScheduledAutoCompleted returns the ScheduledAutoCompleted field value if set, zero value otherwise.
func (o *Incident) GetScheduledAutoCompleted() bool {
	if o == nil || o.ScheduledAutoCompleted == nil {
		var ret bool
		return ret
	}
	return *o.ScheduledAutoCompleted
}

// GetScheduledAutoCompletedOk returns a tuple with the ScheduledAutoCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetScheduledAutoCompletedOk() (*bool, bool) {
	if o == nil || o.ScheduledAutoCompleted == nil {
		return nil, false
	}
	return o.ScheduledAutoCompleted, true
}

// HasScheduledAutoCompleted returns a boolean if a field has been set.
func (o *Incident) HasScheduledAutoCompleted() bool {
	if o != nil && o.ScheduledAutoCompleted != nil {
		return true
	}

	return false
}

// SetScheduledAutoCompleted gets a reference to the given bool and assigns it to the ScheduledAutoCompleted field.
func (o *Incident) SetScheduledAutoCompleted(v bool) {
	o.ScheduledAutoCompleted = &v
}

// GetScheduledAutoInProgress returns the ScheduledAutoInProgress field value if set, zero value otherwise.
func (o *Incident) GetScheduledAutoInProgress() bool {
	if o == nil || o.ScheduledAutoInProgress == nil {
		var ret bool
		return ret
	}
	return *o.ScheduledAutoInProgress
}

// GetScheduledAutoInProgressOk returns a tuple with the ScheduledAutoInProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetScheduledAutoInProgressOk() (*bool, bool) {
	if o == nil || o.ScheduledAutoInProgress == nil {
		return nil, false
	}
	return o.ScheduledAutoInProgress, true
}

// HasScheduledAutoInProgress returns a boolean if a field has been set.
func (o *Incident) HasScheduledAutoInProgress() bool {
	if o != nil && o.ScheduledAutoInProgress != nil {
		return true
	}

	return false
}

// SetScheduledAutoInProgress gets a reference to the given bool and assigns it to the ScheduledAutoInProgress field.
func (o *Incident) SetScheduledAutoInProgress(v bool) {
	o.ScheduledAutoInProgress = &v
}

// GetScheduledFor returns the ScheduledFor field value if set, zero value otherwise.
func (o *Incident) GetScheduledFor() time.Time {
	if o == nil || o.ScheduledFor == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledFor
}

// GetScheduledForOk returns a tuple with the ScheduledFor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetScheduledForOk() (*time.Time, bool) {
	if o == nil || o.ScheduledFor == nil {
		return nil, false
	}
	return o.ScheduledFor, true
}

// HasScheduledFor returns a boolean if a field has been set.
func (o *Incident) HasScheduledFor() bool {
	if o != nil && o.ScheduledFor != nil {
		return true
	}

	return false
}

// SetScheduledFor gets a reference to the given time.Time and assigns it to the ScheduledFor field.
func (o *Incident) SetScheduledFor(v time.Time) {
	o.ScheduledFor = &v
}

// GetScheduledRemindPrior returns the ScheduledRemindPrior field value if set, zero value otherwise.
func (o *Incident) GetScheduledRemindPrior() bool {
	if o == nil || o.ScheduledRemindPrior == nil {
		var ret bool
		return ret
	}
	return *o.ScheduledRemindPrior
}

// GetScheduledRemindPriorOk returns a tuple with the ScheduledRemindPrior field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetScheduledRemindPriorOk() (*bool, bool) {
	if o == nil || o.ScheduledRemindPrior == nil {
		return nil, false
	}
	return o.ScheduledRemindPrior, true
}

// HasScheduledRemindPrior returns a boolean if a field has been set.
func (o *Incident) HasScheduledRemindPrior() bool {
	if o != nil && o.ScheduledRemindPrior != nil {
		return true
	}

	return false
}

// SetScheduledRemindPrior gets a reference to the given bool and assigns it to the ScheduledRemindPrior field.
func (o *Incident) SetScheduledRemindPrior(v bool) {
	o.ScheduledRemindPrior = &v
}

// GetScheduledRemindedAt returns the ScheduledRemindedAt field value if set, zero value otherwise.
func (o *Incident) GetScheduledRemindedAt() time.Time {
	if o == nil || o.ScheduledRemindedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledRemindedAt
}

// GetScheduledRemindedAtOk returns a tuple with the ScheduledRemindedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetScheduledRemindedAtOk() (*time.Time, bool) {
	if o == nil || o.ScheduledRemindedAt == nil {
		return nil, false
	}
	return o.ScheduledRemindedAt, true
}

// HasScheduledRemindedAt returns a boolean if a field has been set.
func (o *Incident) HasScheduledRemindedAt() bool {
	if o != nil && o.ScheduledRemindedAt != nil {
		return true
	}

	return false
}

// SetScheduledRemindedAt gets a reference to the given time.Time and assigns it to the ScheduledRemindedAt field.
func (o *Incident) SetScheduledRemindedAt(v time.Time) {
	o.ScheduledRemindedAt = &v
}

// GetScheduledUntil returns the ScheduledUntil field value if set, zero value otherwise.
func (o *Incident) GetScheduledUntil() time.Time {
	if o == nil || o.ScheduledUntil == nil {
		var ret time.Time
		return ret
	}
	return *o.ScheduledUntil
}

// GetScheduledUntilOk returns a tuple with the ScheduledUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetScheduledUntilOk() (*time.Time, bool) {
	if o == nil || o.ScheduledUntil == nil {
		return nil, false
	}
	return o.ScheduledUntil, true
}

// HasScheduledUntil returns a boolean if a field has been set.
func (o *Incident) HasScheduledUntil() bool {
	if o != nil && o.ScheduledUntil != nil {
		return true
	}

	return false
}

// SetScheduledUntil gets a reference to the given time.Time and assigns it to the ScheduledUntil field.
func (o *Incident) SetScheduledUntil(v time.Time) {
	o.ScheduledUntil = &v
}

// GetShortlink returns the Shortlink field value if set, zero value otherwise.
func (o *Incident) GetShortlink() string {
	if o == nil || o.Shortlink == nil {
		var ret string
		return ret
	}
	return *o.Shortlink
}

// GetShortlinkOk returns a tuple with the Shortlink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetShortlinkOk() (*string, bool) {
	if o == nil || o.Shortlink == nil {
		return nil, false
	}
	return o.Shortlink, true
}

// HasShortlink returns a boolean if a field has been set.
func (o *Incident) HasShortlink() bool {
	if o != nil && o.Shortlink != nil {
		return true
	}

	return false
}

// SetShortlink gets a reference to the given string and assigns it to the Shortlink field.
func (o *Incident) SetShortlink(v string) {
	o.Shortlink = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Incident) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Incident) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Incident) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Incident) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Incident) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Incident) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Incident) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o Incident) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Impact != nil {
		toSerialize["impact"] = o.Impact
	}
	if o.ImpactOverride != nil {
		toSerialize["impact_override"] = o.ImpactOverride
	}
	if o.IncidentUpdates != nil {
		toSerialize["incident_updates"] = o.IncidentUpdates
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MonitoringAt != nil {
		toSerialize["monitoring_at"] = o.MonitoringAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PageId != nil {
		toSerialize["page_id"] = o.PageId
	}
	if o.PostmortemBody != nil {
		toSerialize["postmortem_body"] = o.PostmortemBody
	}
	if o.PostmortemBodyLastUpdatedAt != nil {
		toSerialize["postmortem_body_last_updated_at"] = o.PostmortemBodyLastUpdatedAt
	}
	if o.PostmortemIgnored != nil {
		toSerialize["postmortem_ignored"] = o.PostmortemIgnored
	}
	if o.PostmortemNotifiedSubscribers != nil {
		toSerialize["postmortem_notified_subscribers"] = o.PostmortemNotifiedSubscribers
	}
	if o.PostmortemNotifiedTwitter != nil {
		toSerialize["postmortem_notified_twitter"] = o.PostmortemNotifiedTwitter
	}
	if o.PostmortemPublishedAt != nil {
		toSerialize["postmortem_published_at"] = o.PostmortemPublishedAt
	}
	if o.ResolvedAt != nil {
		toSerialize["resolved_at"] = o.ResolvedAt
	}
	if o.ScheduledAutoCompleted != nil {
		toSerialize["scheduled_auto_completed"] = o.ScheduledAutoCompleted
	}
	if o.ScheduledAutoInProgress != nil {
		toSerialize["scheduled_auto_in_progress"] = o.ScheduledAutoInProgress
	}
	if o.ScheduledFor != nil {
		toSerialize["scheduled_for"] = o.ScheduledFor
	}
	if o.ScheduledRemindPrior != nil {
		toSerialize["scheduled_remind_prior"] = o.ScheduledRemindPrior
	}
	if o.ScheduledRemindedAt != nil {
		toSerialize["scheduled_reminded_at"] = o.ScheduledRemindedAt
	}
	if o.ScheduledUntil != nil {
		toSerialize["scheduled_until"] = o.ScheduledUntil
	}
	if o.Shortlink != nil {
		toSerialize["shortlink"] = o.Shortlink
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableIncident struct {
	value *Incident
	isSet bool
}

func (v NullableIncident) Get() *Incident {
	return v.value
}

func (v *NullableIncident) Set(val *Incident) {
	v.value = val
	v.isSet = true
}

func (v NullableIncident) IsSet() bool {
	return v.isSet
}

func (v *NullableIncident) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncident(val *Incident) *NullableIncident {
	return &NullableIncident{value: val, isSet: true}
}

func (v NullableIncident) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncident) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


