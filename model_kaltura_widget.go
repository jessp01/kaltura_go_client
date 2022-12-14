/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// KalturaWidget struct for KalturaWidget
type KalturaWidget struct {
	// Addes the HTML5 script line to the widget's embed code
	AddEmbedHtml5Support *bool `json:"addEmbedHtml5Support,omitempty"`
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// Should enforce entitlement on feed entries
	EnforceEntitlement *bool `json:"enforceEntitlement,omitempty"`
	EntryId *string `json:"entryId,omitempty"`
	// `readOnly`
	Id *string `json:"id,omitempty"`
	// Can be used to store various partner related data as a string
	PartnerData *string `json:"partnerData,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// Set privacy context for search entries that assiged to private and public categories within a category privacy context.
	PrivacyContext *string `json:"privacyContext,omitempty"`
	Privileges *string `json:"privileges,omitempty"`
	Roles *string `json:"roles,omitempty"`
	// `readOnly`
	RootWidgetId *string `json:"rootWidgetId,omitempty"`
	SecurityPolicy *int32 `json:"securityPolicy,omitempty"`
	// Enum Type: `KalturaWidgetSecurityType`
	SecurityType *int32 `json:"securityType,omitempty"`
	SourceWidgetId *string `json:"sourceWidgetId,omitempty"`
	UiConfId *int32 `json:"uiConfId,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	// `readOnly`
	WidgetHTML *string `json:"widgetHTML,omitempty"`
}

// NewKalturaWidget instantiates a new KalturaWidget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaWidget() *KalturaWidget {
	this := KalturaWidget{}
	return &this
}

// NewKalturaWidgetWithDefaults instantiates a new KalturaWidget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaWidgetWithDefaults() *KalturaWidget {
	this := KalturaWidget{}
	return &this
}

// GetAddEmbedHtml5Support returns the AddEmbedHtml5Support field value if set, zero value otherwise.
func (o *KalturaWidget) GetAddEmbedHtml5Support() bool {
	if o == nil || o.AddEmbedHtml5Support == nil {
		var ret bool
		return ret
	}
	return *o.AddEmbedHtml5Support
}

// GetAddEmbedHtml5SupportOk returns a tuple with the AddEmbedHtml5Support field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetAddEmbedHtml5SupportOk() (*bool, bool) {
	if o == nil || o.AddEmbedHtml5Support == nil {
		return nil, false
	}
	return o.AddEmbedHtml5Support, true
}

// HasAddEmbedHtml5Support returns a boolean if a field has been set.
func (o *KalturaWidget) HasAddEmbedHtml5Support() bool {
	if o != nil && o.AddEmbedHtml5Support != nil {
		return true
	}

	return false
}

// SetAddEmbedHtml5Support gets a reference to the given bool and assigns it to the AddEmbedHtml5Support field.
func (o *KalturaWidget) SetAddEmbedHtml5Support(v bool) {
	o.AddEmbedHtml5Support = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaWidget) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaWidget) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaWidget) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetEnforceEntitlement returns the EnforceEntitlement field value if set, zero value otherwise.
func (o *KalturaWidget) GetEnforceEntitlement() bool {
	if o == nil || o.EnforceEntitlement == nil {
		var ret bool
		return ret
	}
	return *o.EnforceEntitlement
}

// GetEnforceEntitlementOk returns a tuple with the EnforceEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetEnforceEntitlementOk() (*bool, bool) {
	if o == nil || o.EnforceEntitlement == nil {
		return nil, false
	}
	return o.EnforceEntitlement, true
}

// HasEnforceEntitlement returns a boolean if a field has been set.
func (o *KalturaWidget) HasEnforceEntitlement() bool {
	if o != nil && o.EnforceEntitlement != nil {
		return true
	}

	return false
}

// SetEnforceEntitlement gets a reference to the given bool and assigns it to the EnforceEntitlement field.
func (o *KalturaWidget) SetEnforceEntitlement(v bool) {
	o.EnforceEntitlement = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaWidget) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaWidget) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaWidget) SetEntryId(v string) {
	o.EntryId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaWidget) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaWidget) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaWidget) SetId(v string) {
	o.Id = &v
}

// GetPartnerData returns the PartnerData field value if set, zero value otherwise.
func (o *KalturaWidget) GetPartnerData() string {
	if o == nil || o.PartnerData == nil {
		var ret string
		return ret
	}
	return *o.PartnerData
}

// GetPartnerDataOk returns a tuple with the PartnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetPartnerDataOk() (*string, bool) {
	if o == nil || o.PartnerData == nil {
		return nil, false
	}
	return o.PartnerData, true
}

// HasPartnerData returns a boolean if a field has been set.
func (o *KalturaWidget) HasPartnerData() bool {
	if o != nil && o.PartnerData != nil {
		return true
	}

	return false
}

// SetPartnerData gets a reference to the given string and assigns it to the PartnerData field.
func (o *KalturaWidget) SetPartnerData(v string) {
	o.PartnerData = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaWidget) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaWidget) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaWidget) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPrivacyContext returns the PrivacyContext field value if set, zero value otherwise.
func (o *KalturaWidget) GetPrivacyContext() string {
	if o == nil || o.PrivacyContext == nil {
		var ret string
		return ret
	}
	return *o.PrivacyContext
}

// GetPrivacyContextOk returns a tuple with the PrivacyContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetPrivacyContextOk() (*string, bool) {
	if o == nil || o.PrivacyContext == nil {
		return nil, false
	}
	return o.PrivacyContext, true
}

// HasPrivacyContext returns a boolean if a field has been set.
func (o *KalturaWidget) HasPrivacyContext() bool {
	if o != nil && o.PrivacyContext != nil {
		return true
	}

	return false
}

// SetPrivacyContext gets a reference to the given string and assigns it to the PrivacyContext field.
func (o *KalturaWidget) SetPrivacyContext(v string) {
	o.PrivacyContext = &v
}

// GetPrivileges returns the Privileges field value if set, zero value otherwise.
func (o *KalturaWidget) GetPrivileges() string {
	if o == nil || o.Privileges == nil {
		var ret string
		return ret
	}
	return *o.Privileges
}

// GetPrivilegesOk returns a tuple with the Privileges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetPrivilegesOk() (*string, bool) {
	if o == nil || o.Privileges == nil {
		return nil, false
	}
	return o.Privileges, true
}

// HasPrivileges returns a boolean if a field has been set.
func (o *KalturaWidget) HasPrivileges() bool {
	if o != nil && o.Privileges != nil {
		return true
	}

	return false
}

// SetPrivileges gets a reference to the given string and assigns it to the Privileges field.
func (o *KalturaWidget) SetPrivileges(v string) {
	o.Privileges = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *KalturaWidget) GetRoles() string {
	if o == nil || o.Roles == nil {
		var ret string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetRolesOk() (*string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *KalturaWidget) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given string and assigns it to the Roles field.
func (o *KalturaWidget) SetRoles(v string) {
	o.Roles = &v
}

// GetRootWidgetId returns the RootWidgetId field value if set, zero value otherwise.
func (o *KalturaWidget) GetRootWidgetId() string {
	if o == nil || o.RootWidgetId == nil {
		var ret string
		return ret
	}
	return *o.RootWidgetId
}

// GetRootWidgetIdOk returns a tuple with the RootWidgetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetRootWidgetIdOk() (*string, bool) {
	if o == nil || o.RootWidgetId == nil {
		return nil, false
	}
	return o.RootWidgetId, true
}

// HasRootWidgetId returns a boolean if a field has been set.
func (o *KalturaWidget) HasRootWidgetId() bool {
	if o != nil && o.RootWidgetId != nil {
		return true
	}

	return false
}

// SetRootWidgetId gets a reference to the given string and assigns it to the RootWidgetId field.
func (o *KalturaWidget) SetRootWidgetId(v string) {
	o.RootWidgetId = &v
}

// GetSecurityPolicy returns the SecurityPolicy field value if set, zero value otherwise.
func (o *KalturaWidget) GetSecurityPolicy() int32 {
	if o == nil || o.SecurityPolicy == nil {
		var ret int32
		return ret
	}
	return *o.SecurityPolicy
}

// GetSecurityPolicyOk returns a tuple with the SecurityPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetSecurityPolicyOk() (*int32, bool) {
	if o == nil || o.SecurityPolicy == nil {
		return nil, false
	}
	return o.SecurityPolicy, true
}

// HasSecurityPolicy returns a boolean if a field has been set.
func (o *KalturaWidget) HasSecurityPolicy() bool {
	if o != nil && o.SecurityPolicy != nil {
		return true
	}

	return false
}

// SetSecurityPolicy gets a reference to the given int32 and assigns it to the SecurityPolicy field.
func (o *KalturaWidget) SetSecurityPolicy(v int32) {
	o.SecurityPolicy = &v
}

// GetSecurityType returns the SecurityType field value if set, zero value otherwise.
func (o *KalturaWidget) GetSecurityType() int32 {
	if o == nil || o.SecurityType == nil {
		var ret int32
		return ret
	}
	return *o.SecurityType
}

// GetSecurityTypeOk returns a tuple with the SecurityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetSecurityTypeOk() (*int32, bool) {
	if o == nil || o.SecurityType == nil {
		return nil, false
	}
	return o.SecurityType, true
}

// HasSecurityType returns a boolean if a field has been set.
func (o *KalturaWidget) HasSecurityType() bool {
	if o != nil && o.SecurityType != nil {
		return true
	}

	return false
}

// SetSecurityType gets a reference to the given int32 and assigns it to the SecurityType field.
func (o *KalturaWidget) SetSecurityType(v int32) {
	o.SecurityType = &v
}

// GetSourceWidgetId returns the SourceWidgetId field value if set, zero value otherwise.
func (o *KalturaWidget) GetSourceWidgetId() string {
	if o == nil || o.SourceWidgetId == nil {
		var ret string
		return ret
	}
	return *o.SourceWidgetId
}

// GetSourceWidgetIdOk returns a tuple with the SourceWidgetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetSourceWidgetIdOk() (*string, bool) {
	if o == nil || o.SourceWidgetId == nil {
		return nil, false
	}
	return o.SourceWidgetId, true
}

// HasSourceWidgetId returns a boolean if a field has been set.
func (o *KalturaWidget) HasSourceWidgetId() bool {
	if o != nil && o.SourceWidgetId != nil {
		return true
	}

	return false
}

// SetSourceWidgetId gets a reference to the given string and assigns it to the SourceWidgetId field.
func (o *KalturaWidget) SetSourceWidgetId(v string) {
	o.SourceWidgetId = &v
}

// GetUiConfId returns the UiConfId field value if set, zero value otherwise.
func (o *KalturaWidget) GetUiConfId() int32 {
	if o == nil || o.UiConfId == nil {
		var ret int32
		return ret
	}
	return *o.UiConfId
}

// GetUiConfIdOk returns a tuple with the UiConfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetUiConfIdOk() (*int32, bool) {
	if o == nil || o.UiConfId == nil {
		return nil, false
	}
	return o.UiConfId, true
}

// HasUiConfId returns a boolean if a field has been set.
func (o *KalturaWidget) HasUiConfId() bool {
	if o != nil && o.UiConfId != nil {
		return true
	}

	return false
}

// SetUiConfId gets a reference to the given int32 and assigns it to the UiConfId field.
func (o *KalturaWidget) SetUiConfId(v int32) {
	o.UiConfId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaWidget) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaWidget) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaWidget) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetWidgetHTML returns the WidgetHTML field value if set, zero value otherwise.
func (o *KalturaWidget) GetWidgetHTML() string {
	if o == nil || o.WidgetHTML == nil {
		var ret string
		return ret
	}
	return *o.WidgetHTML
}

// GetWidgetHTMLOk returns a tuple with the WidgetHTML field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaWidget) GetWidgetHTMLOk() (*string, bool) {
	if o == nil || o.WidgetHTML == nil {
		return nil, false
	}
	return o.WidgetHTML, true
}

// HasWidgetHTML returns a boolean if a field has been set.
func (o *KalturaWidget) HasWidgetHTML() bool {
	if o != nil && o.WidgetHTML != nil {
		return true
	}

	return false
}

// SetWidgetHTML gets a reference to the given string and assigns it to the WidgetHTML field.
func (o *KalturaWidget) SetWidgetHTML(v string) {
	o.WidgetHTML = &v
}

func (o KalturaWidget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddEmbedHtml5Support != nil {
		toSerialize["addEmbedHtml5Support"] = o.AddEmbedHtml5Support
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.EnforceEntitlement != nil {
		toSerialize["enforceEntitlement"] = o.EnforceEntitlement
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PartnerData != nil {
		toSerialize["partnerData"] = o.PartnerData
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.PrivacyContext != nil {
		toSerialize["privacyContext"] = o.PrivacyContext
	}
	if o.Privileges != nil {
		toSerialize["privileges"] = o.Privileges
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.RootWidgetId != nil {
		toSerialize["rootWidgetId"] = o.RootWidgetId
	}
	if o.SecurityPolicy != nil {
		toSerialize["securityPolicy"] = o.SecurityPolicy
	}
	if o.SecurityType != nil {
		toSerialize["securityType"] = o.SecurityType
	}
	if o.SourceWidgetId != nil {
		toSerialize["sourceWidgetId"] = o.SourceWidgetId
	}
	if o.UiConfId != nil {
		toSerialize["uiConfId"] = o.UiConfId
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.WidgetHTML != nil {
		toSerialize["widgetHTML"] = o.WidgetHTML
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaWidget struct {
	value *KalturaWidget
	isSet bool
}

func (v NullableKalturaWidget) Get() *KalturaWidget {
	return v.value
}

func (v *NullableKalturaWidget) Set(val *KalturaWidget) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaWidget) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaWidget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaWidget(val *KalturaWidget) *NullableKalturaWidget {
	return &NullableKalturaWidget{value: val, isSet: true}
}

func (v NullableKalturaWidget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaWidget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


