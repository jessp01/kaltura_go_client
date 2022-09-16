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

// KalturaVirtualEvent struct for KalturaVirtualEvent
type KalturaVirtualEvent struct {
	AdminsGroupIds *string `json:"adminsGroupIds,omitempty"`
	AgendaScheduleEventId *int32 `json:"agendaScheduleEventId,omitempty"`
	AttendeesGroupIds *string `json:"attendeesGroupIds,omitempty"`
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	DeletionDueDate *int32 `json:"deletionDueDate,omitempty"`
	Description *string `json:"description,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	MainEventScheduleEventId *int32 `json:"mainEventScheduleEventId,omitempty"`
	Name *string `json:"name,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// JSON-Schema of the Registration Form
	RegistrationFormSchema *string `json:"registrationFormSchema,omitempty"`
	RegistrationScheduleEventId *int32 `json:"registrationScheduleEventId,omitempty"`
	// `readOnly`  Enum Type: `KalturaVirtualEventStatus`
	Status *int32 `json:"status,omitempty"`
	Tags *string `json:"tags,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
}

// NewKalturaVirtualEvent instantiates a new KalturaVirtualEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVirtualEvent() *KalturaVirtualEvent {
	this := KalturaVirtualEvent{}
	return &this
}

// NewKalturaVirtualEventWithDefaults instantiates a new KalturaVirtualEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVirtualEventWithDefaults() *KalturaVirtualEvent {
	this := KalturaVirtualEvent{}
	return &this
}

// GetAdminsGroupIds returns the AdminsGroupIds field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetAdminsGroupIds() string {
	if o == nil || o.AdminsGroupIds == nil {
		var ret string
		return ret
	}
	return *o.AdminsGroupIds
}

// GetAdminsGroupIdsOk returns a tuple with the AdminsGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetAdminsGroupIdsOk() (*string, bool) {
	if o == nil || o.AdminsGroupIds == nil {
		return nil, false
	}
	return o.AdminsGroupIds, true
}

// HasAdminsGroupIds returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasAdminsGroupIds() bool {
	if o != nil && o.AdminsGroupIds != nil {
		return true
	}

	return false
}

// SetAdminsGroupIds gets a reference to the given string and assigns it to the AdminsGroupIds field.
func (o *KalturaVirtualEvent) SetAdminsGroupIds(v string) {
	o.AdminsGroupIds = &v
}

// GetAgendaScheduleEventId returns the AgendaScheduleEventId field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetAgendaScheduleEventId() int32 {
	if o == nil || o.AgendaScheduleEventId == nil {
		var ret int32
		return ret
	}
	return *o.AgendaScheduleEventId
}

// GetAgendaScheduleEventIdOk returns a tuple with the AgendaScheduleEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetAgendaScheduleEventIdOk() (*int32, bool) {
	if o == nil || o.AgendaScheduleEventId == nil {
		return nil, false
	}
	return o.AgendaScheduleEventId, true
}

// HasAgendaScheduleEventId returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasAgendaScheduleEventId() bool {
	if o != nil && o.AgendaScheduleEventId != nil {
		return true
	}

	return false
}

// SetAgendaScheduleEventId gets a reference to the given int32 and assigns it to the AgendaScheduleEventId field.
func (o *KalturaVirtualEvent) SetAgendaScheduleEventId(v int32) {
	o.AgendaScheduleEventId = &v
}

// GetAttendeesGroupIds returns the AttendeesGroupIds field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetAttendeesGroupIds() string {
	if o == nil || o.AttendeesGroupIds == nil {
		var ret string
		return ret
	}
	return *o.AttendeesGroupIds
}

// GetAttendeesGroupIdsOk returns a tuple with the AttendeesGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetAttendeesGroupIdsOk() (*string, bool) {
	if o == nil || o.AttendeesGroupIds == nil {
		return nil, false
	}
	return o.AttendeesGroupIds, true
}

// HasAttendeesGroupIds returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasAttendeesGroupIds() bool {
	if o != nil && o.AttendeesGroupIds != nil {
		return true
	}

	return false
}

// SetAttendeesGroupIds gets a reference to the given string and assigns it to the AttendeesGroupIds field.
func (o *KalturaVirtualEvent) SetAttendeesGroupIds(v string) {
	o.AttendeesGroupIds = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaVirtualEvent) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDeletionDueDate returns the DeletionDueDate field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetDeletionDueDate() int32 {
	if o == nil || o.DeletionDueDate == nil {
		var ret int32
		return ret
	}
	return *o.DeletionDueDate
}

// GetDeletionDueDateOk returns a tuple with the DeletionDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetDeletionDueDateOk() (*int32, bool) {
	if o == nil || o.DeletionDueDate == nil {
		return nil, false
	}
	return o.DeletionDueDate, true
}

// HasDeletionDueDate returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasDeletionDueDate() bool {
	if o != nil && o.DeletionDueDate != nil {
		return true
	}

	return false
}

// SetDeletionDueDate gets a reference to the given int32 and assigns it to the DeletionDueDate field.
func (o *KalturaVirtualEvent) SetDeletionDueDate(v int32) {
	o.DeletionDueDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaVirtualEvent) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaVirtualEvent) SetId(v int32) {
	o.Id = &v
}

// GetMainEventScheduleEventId returns the MainEventScheduleEventId field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetMainEventScheduleEventId() int32 {
	if o == nil || o.MainEventScheduleEventId == nil {
		var ret int32
		return ret
	}
	return *o.MainEventScheduleEventId
}

// GetMainEventScheduleEventIdOk returns a tuple with the MainEventScheduleEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetMainEventScheduleEventIdOk() (*int32, bool) {
	if o == nil || o.MainEventScheduleEventId == nil {
		return nil, false
	}
	return o.MainEventScheduleEventId, true
}

// HasMainEventScheduleEventId returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasMainEventScheduleEventId() bool {
	if o != nil && o.MainEventScheduleEventId != nil {
		return true
	}

	return false
}

// SetMainEventScheduleEventId gets a reference to the given int32 and assigns it to the MainEventScheduleEventId field.
func (o *KalturaVirtualEvent) SetMainEventScheduleEventId(v int32) {
	o.MainEventScheduleEventId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaVirtualEvent) SetName(v string) {
	o.Name = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaVirtualEvent) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetRegistrationFormSchema returns the RegistrationFormSchema field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetRegistrationFormSchema() string {
	if o == nil || o.RegistrationFormSchema == nil {
		var ret string
		return ret
	}
	return *o.RegistrationFormSchema
}

// GetRegistrationFormSchemaOk returns a tuple with the RegistrationFormSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetRegistrationFormSchemaOk() (*string, bool) {
	if o == nil || o.RegistrationFormSchema == nil {
		return nil, false
	}
	return o.RegistrationFormSchema, true
}

// HasRegistrationFormSchema returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasRegistrationFormSchema() bool {
	if o != nil && o.RegistrationFormSchema != nil {
		return true
	}

	return false
}

// SetRegistrationFormSchema gets a reference to the given string and assigns it to the RegistrationFormSchema field.
func (o *KalturaVirtualEvent) SetRegistrationFormSchema(v string) {
	o.RegistrationFormSchema = &v
}

// GetRegistrationScheduleEventId returns the RegistrationScheduleEventId field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetRegistrationScheduleEventId() int32 {
	if o == nil || o.RegistrationScheduleEventId == nil {
		var ret int32
		return ret
	}
	return *o.RegistrationScheduleEventId
}

// GetRegistrationScheduleEventIdOk returns a tuple with the RegistrationScheduleEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetRegistrationScheduleEventIdOk() (*int32, bool) {
	if o == nil || o.RegistrationScheduleEventId == nil {
		return nil, false
	}
	return o.RegistrationScheduleEventId, true
}

// HasRegistrationScheduleEventId returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasRegistrationScheduleEventId() bool {
	if o != nil && o.RegistrationScheduleEventId != nil {
		return true
	}

	return false
}

// SetRegistrationScheduleEventId gets a reference to the given int32 and assigns it to the RegistrationScheduleEventId field.
func (o *KalturaVirtualEvent) SetRegistrationScheduleEventId(v int32) {
	o.RegistrationScheduleEventId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaVirtualEvent) SetStatus(v int32) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *KalturaVirtualEvent) SetTags(v string) {
	o.Tags = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaVirtualEvent) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaVirtualEvent) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaVirtualEvent) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaVirtualEvent) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

func (o KalturaVirtualEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminsGroupIds != nil {
		toSerialize["adminsGroupIds"] = o.AdminsGroupIds
	}
	if o.AgendaScheduleEventId != nil {
		toSerialize["agendaScheduleEventId"] = o.AgendaScheduleEventId
	}
	if o.AttendeesGroupIds != nil {
		toSerialize["attendeesGroupIds"] = o.AttendeesGroupIds
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DeletionDueDate != nil {
		toSerialize["deletionDueDate"] = o.DeletionDueDate
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MainEventScheduleEventId != nil {
		toSerialize["mainEventScheduleEventId"] = o.MainEventScheduleEventId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.RegistrationFormSchema != nil {
		toSerialize["registrationFormSchema"] = o.RegistrationFormSchema
	}
	if o.RegistrationScheduleEventId != nil {
		toSerialize["registrationScheduleEventId"] = o.RegistrationScheduleEventId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVirtualEvent struct {
	value *KalturaVirtualEvent
	isSet bool
}

func (v NullableKalturaVirtualEvent) Get() *KalturaVirtualEvent {
	return v.value
}

func (v *NullableKalturaVirtualEvent) Set(val *KalturaVirtualEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVirtualEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVirtualEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVirtualEvent(val *KalturaVirtualEvent) *NullableKalturaVirtualEvent {
	return &NullableKalturaVirtualEvent{value: val, isSet: true}
}

func (v NullableKalturaVirtualEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVirtualEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

