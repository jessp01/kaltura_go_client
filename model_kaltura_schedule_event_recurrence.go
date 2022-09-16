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

// KalturaScheduleEventRecurrence struct for KalturaScheduleEventRecurrence
type KalturaScheduleEventRecurrence struct {
	// Comma separated of KalturaScheduleEventRecurrenceDay  Each byDay value can also be preceded by a positive (+n) or negative (-n) integer.  If present, this indicates the nth occurrence of the specific day within the MONTHLY or YEARLY RRULE.  For example, within a MONTHLY rule, +1MO (or simply 1MO) represents the first Monday within the month, whereas -1MO represents the last Monday of the month.  If an integer modifier is not present, it means all days of this type within the specified frequency.  For example, within a MONTHLY rule, MO represents all Mondays within the month.
	ByDay *string `json:"byDay,omitempty"`
	// Comma separated numbers between 0 to 23
	ByHour *string `json:"byHour,omitempty"`
	// Comma separated numbers between 0 to 59
	ByMinute *string `json:"byMinute,omitempty"`
	// Comma separated numbers between 1 to 12
	ByMonth *string `json:"byMonth,omitempty"`
	// Comma separated of numbers between -31 to 31, excluding 0.  For example, -10 represents the tenth to the last day of the month.
	ByMonthDay *string `json:"byMonthDay,omitempty"`
	// Comma separated of numbers between -366 to 366, excluding 0.  Corresponds to the nth occurrence within the set of events specified by the rule.  It must only be used in conjunction with another byrule part.  For example \"the last work day of the month\" could be represented as: frequency=MONTHLY;byDay=MO,TU,WE,TH,FR;byOffset=-1  Each byOffset value can include a positive (+n) or negative (-n) integer.  If present, this indicates the nth occurrence of the specific occurrence within the set of events specified by the rule.
	ByOffset *string `json:"byOffset,omitempty"`
	// Comma separated numbers between 0 to 59
	BySecond *string `json:"bySecond,omitempty"`
	// Comma separated of numbers between -53 to 53, excluding 0.  This corresponds to weeks according to week numbering.  A week is defined as a seven day period, starting on the day of the week defined to be the week start.  Week number one of the calendar year is the first week which contains at least four (4) days in that calendar year.  This rule part is only valid for YEARLY frequency.  For example, 3 represents the third week of the year.
	ByWeekNumber *string `json:"byWeekNumber,omitempty"`
	// Comma separated of numbers between -366 to 366, excluding 0.  For example, -1 represents the last day of the year (December 31st) and -306 represents the 306th to the last day of the year (March 1st).
	ByYearDay *string `json:"byYearDay,omitempty"`
	Count *int32 `json:"count,omitempty"`
	// Enum Type: `KalturaScheduleEventRecurrenceFrequency`
	Frequency *string `json:"frequency,omitempty"`
	Interval *int32 `json:"interval,omitempty"`
	Name *string `json:"name,omitempty"`
	// TimeZone String
	TimeZone *string `json:"timeZone,omitempty"`
	Until *int32 `json:"until,omitempty"`
	// Enum Type: `KalturaScheduleEventRecurrenceDay`
	WeekStartDay *string `json:"weekStartDay,omitempty"`
}

// NewKalturaScheduleEventRecurrence instantiates a new KalturaScheduleEventRecurrence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaScheduleEventRecurrence() *KalturaScheduleEventRecurrence {
	this := KalturaScheduleEventRecurrence{}
	return &this
}

// NewKalturaScheduleEventRecurrenceWithDefaults instantiates a new KalturaScheduleEventRecurrence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaScheduleEventRecurrenceWithDefaults() *KalturaScheduleEventRecurrence {
	this := KalturaScheduleEventRecurrence{}
	return &this
}

// GetByDay returns the ByDay field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByDay() string {
	if o == nil || o.ByDay == nil {
		var ret string
		return ret
	}
	return *o.ByDay
}

// GetByDayOk returns a tuple with the ByDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByDayOk() (*string, bool) {
	if o == nil || o.ByDay == nil {
		return nil, false
	}
	return o.ByDay, true
}

// HasByDay returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByDay() bool {
	if o != nil && o.ByDay != nil {
		return true
	}

	return false
}

// SetByDay gets a reference to the given string and assigns it to the ByDay field.
func (o *KalturaScheduleEventRecurrence) SetByDay(v string) {
	o.ByDay = &v
}

// GetByHour returns the ByHour field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByHour() string {
	if o == nil || o.ByHour == nil {
		var ret string
		return ret
	}
	return *o.ByHour
}

// GetByHourOk returns a tuple with the ByHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByHourOk() (*string, bool) {
	if o == nil || o.ByHour == nil {
		return nil, false
	}
	return o.ByHour, true
}

// HasByHour returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByHour() bool {
	if o != nil && o.ByHour != nil {
		return true
	}

	return false
}

// SetByHour gets a reference to the given string and assigns it to the ByHour field.
func (o *KalturaScheduleEventRecurrence) SetByHour(v string) {
	o.ByHour = &v
}

// GetByMinute returns the ByMinute field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByMinute() string {
	if o == nil || o.ByMinute == nil {
		var ret string
		return ret
	}
	return *o.ByMinute
}

// GetByMinuteOk returns a tuple with the ByMinute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByMinuteOk() (*string, bool) {
	if o == nil || o.ByMinute == nil {
		return nil, false
	}
	return o.ByMinute, true
}

// HasByMinute returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByMinute() bool {
	if o != nil && o.ByMinute != nil {
		return true
	}

	return false
}

// SetByMinute gets a reference to the given string and assigns it to the ByMinute field.
func (o *KalturaScheduleEventRecurrence) SetByMinute(v string) {
	o.ByMinute = &v
}

// GetByMonth returns the ByMonth field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByMonth() string {
	if o == nil || o.ByMonth == nil {
		var ret string
		return ret
	}
	return *o.ByMonth
}

// GetByMonthOk returns a tuple with the ByMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByMonthOk() (*string, bool) {
	if o == nil || o.ByMonth == nil {
		return nil, false
	}
	return o.ByMonth, true
}

// HasByMonth returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByMonth() bool {
	if o != nil && o.ByMonth != nil {
		return true
	}

	return false
}

// SetByMonth gets a reference to the given string and assigns it to the ByMonth field.
func (o *KalturaScheduleEventRecurrence) SetByMonth(v string) {
	o.ByMonth = &v
}

// GetByMonthDay returns the ByMonthDay field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByMonthDay() string {
	if o == nil || o.ByMonthDay == nil {
		var ret string
		return ret
	}
	return *o.ByMonthDay
}

// GetByMonthDayOk returns a tuple with the ByMonthDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByMonthDayOk() (*string, bool) {
	if o == nil || o.ByMonthDay == nil {
		return nil, false
	}
	return o.ByMonthDay, true
}

// HasByMonthDay returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByMonthDay() bool {
	if o != nil && o.ByMonthDay != nil {
		return true
	}

	return false
}

// SetByMonthDay gets a reference to the given string and assigns it to the ByMonthDay field.
func (o *KalturaScheduleEventRecurrence) SetByMonthDay(v string) {
	o.ByMonthDay = &v
}

// GetByOffset returns the ByOffset field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByOffset() string {
	if o == nil || o.ByOffset == nil {
		var ret string
		return ret
	}
	return *o.ByOffset
}

// GetByOffsetOk returns a tuple with the ByOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByOffsetOk() (*string, bool) {
	if o == nil || o.ByOffset == nil {
		return nil, false
	}
	return o.ByOffset, true
}

// HasByOffset returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByOffset() bool {
	if o != nil && o.ByOffset != nil {
		return true
	}

	return false
}

// SetByOffset gets a reference to the given string and assigns it to the ByOffset field.
func (o *KalturaScheduleEventRecurrence) SetByOffset(v string) {
	o.ByOffset = &v
}

// GetBySecond returns the BySecond field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetBySecond() string {
	if o == nil || o.BySecond == nil {
		var ret string
		return ret
	}
	return *o.BySecond
}

// GetBySecondOk returns a tuple with the BySecond field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetBySecondOk() (*string, bool) {
	if o == nil || o.BySecond == nil {
		return nil, false
	}
	return o.BySecond, true
}

// HasBySecond returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasBySecond() bool {
	if o != nil && o.BySecond != nil {
		return true
	}

	return false
}

// SetBySecond gets a reference to the given string and assigns it to the BySecond field.
func (o *KalturaScheduleEventRecurrence) SetBySecond(v string) {
	o.BySecond = &v
}

// GetByWeekNumber returns the ByWeekNumber field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByWeekNumber() string {
	if o == nil || o.ByWeekNumber == nil {
		var ret string
		return ret
	}
	return *o.ByWeekNumber
}

// GetByWeekNumberOk returns a tuple with the ByWeekNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByWeekNumberOk() (*string, bool) {
	if o == nil || o.ByWeekNumber == nil {
		return nil, false
	}
	return o.ByWeekNumber, true
}

// HasByWeekNumber returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByWeekNumber() bool {
	if o != nil && o.ByWeekNumber != nil {
		return true
	}

	return false
}

// SetByWeekNumber gets a reference to the given string and assigns it to the ByWeekNumber field.
func (o *KalturaScheduleEventRecurrence) SetByWeekNumber(v string) {
	o.ByWeekNumber = &v
}

// GetByYearDay returns the ByYearDay field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetByYearDay() string {
	if o == nil || o.ByYearDay == nil {
		var ret string
		return ret
	}
	return *o.ByYearDay
}

// GetByYearDayOk returns a tuple with the ByYearDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetByYearDayOk() (*string, bool) {
	if o == nil || o.ByYearDay == nil {
		return nil, false
	}
	return o.ByYearDay, true
}

// HasByYearDay returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasByYearDay() bool {
	if o != nil && o.ByYearDay != nil {
		return true
	}

	return false
}

// SetByYearDay gets a reference to the given string and assigns it to the ByYearDay field.
func (o *KalturaScheduleEventRecurrence) SetByYearDay(v string) {
	o.ByYearDay = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *KalturaScheduleEventRecurrence) SetCount(v int32) {
	o.Count = &v
}

// GetFrequency returns the Frequency field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetFrequency() string {
	if o == nil || o.Frequency == nil {
		var ret string
		return ret
	}
	return *o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetFrequencyOk() (*string, bool) {
	if o == nil || o.Frequency == nil {
		return nil, false
	}
	return o.Frequency, true
}

// HasFrequency returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasFrequency() bool {
	if o != nil && o.Frequency != nil {
		return true
	}

	return false
}

// SetFrequency gets a reference to the given string and assigns it to the Frequency field.
func (o *KalturaScheduleEventRecurrence) SetFrequency(v string) {
	o.Frequency = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetInterval() int32 {
	if o == nil || o.Interval == nil {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetIntervalOk() (*int32, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *KalturaScheduleEventRecurrence) SetInterval(v int32) {
	o.Interval = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaScheduleEventRecurrence) SetName(v string) {
	o.Name = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetTimeZone() string {
	if o == nil || o.TimeZone == nil {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetTimeZoneOk() (*string, bool) {
	if o == nil || o.TimeZone == nil {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasTimeZone() bool {
	if o != nil && o.TimeZone != nil {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *KalturaScheduleEventRecurrence) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetUntil returns the Until field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetUntil() int32 {
	if o == nil || o.Until == nil {
		var ret int32
		return ret
	}
	return *o.Until
}

// GetUntilOk returns a tuple with the Until field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetUntilOk() (*int32, bool) {
	if o == nil || o.Until == nil {
		return nil, false
	}
	return o.Until, true
}

// HasUntil returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasUntil() bool {
	if o != nil && o.Until != nil {
		return true
	}

	return false
}

// SetUntil gets a reference to the given int32 and assigns it to the Until field.
func (o *KalturaScheduleEventRecurrence) SetUntil(v int32) {
	o.Until = &v
}

// GetWeekStartDay returns the WeekStartDay field value if set, zero value otherwise.
func (o *KalturaScheduleEventRecurrence) GetWeekStartDay() string {
	if o == nil || o.WeekStartDay == nil {
		var ret string
		return ret
	}
	return *o.WeekStartDay
}

// GetWeekStartDayOk returns a tuple with the WeekStartDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaScheduleEventRecurrence) GetWeekStartDayOk() (*string, bool) {
	if o == nil || o.WeekStartDay == nil {
		return nil, false
	}
	return o.WeekStartDay, true
}

// HasWeekStartDay returns a boolean if a field has been set.
func (o *KalturaScheduleEventRecurrence) HasWeekStartDay() bool {
	if o != nil && o.WeekStartDay != nil {
		return true
	}

	return false
}

// SetWeekStartDay gets a reference to the given string and assigns it to the WeekStartDay field.
func (o *KalturaScheduleEventRecurrence) SetWeekStartDay(v string) {
	o.WeekStartDay = &v
}

func (o KalturaScheduleEventRecurrence) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ByDay != nil {
		toSerialize["byDay"] = o.ByDay
	}
	if o.ByHour != nil {
		toSerialize["byHour"] = o.ByHour
	}
	if o.ByMinute != nil {
		toSerialize["byMinute"] = o.ByMinute
	}
	if o.ByMonth != nil {
		toSerialize["byMonth"] = o.ByMonth
	}
	if o.ByMonthDay != nil {
		toSerialize["byMonthDay"] = o.ByMonthDay
	}
	if o.ByOffset != nil {
		toSerialize["byOffset"] = o.ByOffset
	}
	if o.BySecond != nil {
		toSerialize["bySecond"] = o.BySecond
	}
	if o.ByWeekNumber != nil {
		toSerialize["byWeekNumber"] = o.ByWeekNumber
	}
	if o.ByYearDay != nil {
		toSerialize["byYearDay"] = o.ByYearDay
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Frequency != nil {
		toSerialize["frequency"] = o.Frequency
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TimeZone != nil {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Until != nil {
		toSerialize["until"] = o.Until
	}
	if o.WeekStartDay != nil {
		toSerialize["weekStartDay"] = o.WeekStartDay
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaScheduleEventRecurrence struct {
	value *KalturaScheduleEventRecurrence
	isSet bool
}

func (v NullableKalturaScheduleEventRecurrence) Get() *KalturaScheduleEventRecurrence {
	return v.value
}

func (v *NullableKalturaScheduleEventRecurrence) Set(val *KalturaScheduleEventRecurrence) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaScheduleEventRecurrence) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaScheduleEventRecurrence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaScheduleEventRecurrence(val *KalturaScheduleEventRecurrence) *NullableKalturaScheduleEventRecurrence {
	return &NullableKalturaScheduleEventRecurrence{value: val, isSet: true}
}

func (v NullableKalturaScheduleEventRecurrence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaScheduleEventRecurrence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

