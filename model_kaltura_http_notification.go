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

// KalturaHttpNotification Wrapper for sent notifications
type KalturaHttpNotification struct {
	// ID of the batch job that execute the notification
	EventNotificationJobId *int32 `json:"eventNotificationJobId,omitempty"`
	// Enum Type: `KalturaEventNotificationEventObjectType`  Object type that triggered the notification
	EventObjectType *string `json:"eventObjectType,omitempty"`
	// Enum Type: `KalturaEventNotificationEventType`  Ecent type that triggered the notification
	EventType *string `json:"eventType,omitempty"`
	Object map[string]interface{} `json:"object,omitempty"`
	// ID of the template that triggered the notification
	TemplateId *int32 `json:"templateId,omitempty"`
	// Name of the template that triggered the notification
	TemplateName *string `json:"templateName,omitempty"`
	// System name of the template that triggered the notification
	TemplateSystemName *string `json:"templateSystemName,omitempty"`
}

// NewKalturaHttpNotification instantiates a new KalturaHttpNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaHttpNotification() *KalturaHttpNotification {
	this := KalturaHttpNotification{}
	return &this
}

// NewKalturaHttpNotificationWithDefaults instantiates a new KalturaHttpNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaHttpNotificationWithDefaults() *KalturaHttpNotification {
	this := KalturaHttpNotification{}
	return &this
}

// GetEventNotificationJobId returns the EventNotificationJobId field value if set, zero value otherwise.
func (o *KalturaHttpNotification) GetEventNotificationJobId() int32 {
	if o == nil || o.EventNotificationJobId == nil {
		var ret int32
		return ret
	}
	return *o.EventNotificationJobId
}

// GetEventNotificationJobIdOk returns a tuple with the EventNotificationJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaHttpNotification) GetEventNotificationJobIdOk() (*int32, bool) {
	if o == nil || o.EventNotificationJobId == nil {
		return nil, false
	}
	return o.EventNotificationJobId, true
}

// HasEventNotificationJobId returns a boolean if a field has been set.
func (o *KalturaHttpNotification) HasEventNotificationJobId() bool {
	if o != nil && o.EventNotificationJobId != nil {
		return true
	}

	return false
}

// SetEventNotificationJobId gets a reference to the given int32 and assigns it to the EventNotificationJobId field.
func (o *KalturaHttpNotification) SetEventNotificationJobId(v int32) {
	o.EventNotificationJobId = &v
}

// GetEventObjectType returns the EventObjectType field value if set, zero value otherwise.
func (o *KalturaHttpNotification) GetEventObjectType() string {
	if o == nil || o.EventObjectType == nil {
		var ret string
		return ret
	}
	return *o.EventObjectType
}

// GetEventObjectTypeOk returns a tuple with the EventObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaHttpNotification) GetEventObjectTypeOk() (*string, bool) {
	if o == nil || o.EventObjectType == nil {
		return nil, false
	}
	return o.EventObjectType, true
}

// HasEventObjectType returns a boolean if a field has been set.
func (o *KalturaHttpNotification) HasEventObjectType() bool {
	if o != nil && o.EventObjectType != nil {
		return true
	}

	return false
}

// SetEventObjectType gets a reference to the given string and assigns it to the EventObjectType field.
func (o *KalturaHttpNotification) SetEventObjectType(v string) {
	o.EventObjectType = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *KalturaHttpNotification) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaHttpNotification) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *KalturaHttpNotification) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *KalturaHttpNotification) SetEventType(v string) {
	o.EventType = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *KalturaHttpNotification) GetObject() map[string]interface{} {
	if o == nil || o.Object == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaHttpNotification) GetObjectOk() (map[string]interface{}, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *KalturaHttpNotification) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given map[string]interface{} and assigns it to the Object field.
func (o *KalturaHttpNotification) SetObject(v map[string]interface{}) {
	o.Object = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *KalturaHttpNotification) GetTemplateId() int32 {
	if o == nil || o.TemplateId == nil {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaHttpNotification) GetTemplateIdOk() (*int32, bool) {
	if o == nil || o.TemplateId == nil {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *KalturaHttpNotification) HasTemplateId() bool {
	if o != nil && o.TemplateId != nil {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *KalturaHttpNotification) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *KalturaHttpNotification) GetTemplateName() string {
	if o == nil || o.TemplateName == nil {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaHttpNotification) GetTemplateNameOk() (*string, bool) {
	if o == nil || o.TemplateName == nil {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *KalturaHttpNotification) HasTemplateName() bool {
	if o != nil && o.TemplateName != nil {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *KalturaHttpNotification) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTemplateSystemName returns the TemplateSystemName field value if set, zero value otherwise.
func (o *KalturaHttpNotification) GetTemplateSystemName() string {
	if o == nil || o.TemplateSystemName == nil {
		var ret string
		return ret
	}
	return *o.TemplateSystemName
}

// GetTemplateSystemNameOk returns a tuple with the TemplateSystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaHttpNotification) GetTemplateSystemNameOk() (*string, bool) {
	if o == nil || o.TemplateSystemName == nil {
		return nil, false
	}
	return o.TemplateSystemName, true
}

// HasTemplateSystemName returns a boolean if a field has been set.
func (o *KalturaHttpNotification) HasTemplateSystemName() bool {
	if o != nil && o.TemplateSystemName != nil {
		return true
	}

	return false
}

// SetTemplateSystemName gets a reference to the given string and assigns it to the TemplateSystemName field.
func (o *KalturaHttpNotification) SetTemplateSystemName(v string) {
	o.TemplateSystemName = &v
}

func (o KalturaHttpNotification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventNotificationJobId != nil {
		toSerialize["eventNotificationJobId"] = o.EventNotificationJobId
	}
	if o.EventObjectType != nil {
		toSerialize["eventObjectType"] = o.EventObjectType
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.TemplateId != nil {
		toSerialize["templateId"] = o.TemplateId
	}
	if o.TemplateName != nil {
		toSerialize["templateName"] = o.TemplateName
	}
	if o.TemplateSystemName != nil {
		toSerialize["templateSystemName"] = o.TemplateSystemName
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaHttpNotification struct {
	value *KalturaHttpNotification
	isSet bool
}

func (v NullableKalturaHttpNotification) Get() *KalturaHttpNotification {
	return v.value
}

func (v *NullableKalturaHttpNotification) Set(val *KalturaHttpNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaHttpNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaHttpNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaHttpNotification(val *KalturaHttpNotification) *NullableKalturaHttpNotification {
	return &NullableKalturaHttpNotification{value: val, isSet: true}
}

func (v NullableKalturaHttpNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaHttpNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

