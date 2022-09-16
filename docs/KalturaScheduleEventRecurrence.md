# KalturaScheduleEventRecurrence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ByDay** | Pointer to **string** | Comma separated of KalturaScheduleEventRecurrenceDay  Each byDay value can also be preceded by a positive (+n) or negative (-n) integer.  If present, this indicates the nth occurrence of the specific day within the MONTHLY or YEARLY RRULE.  For example, within a MONTHLY rule, +1MO (or simply 1MO) represents the first Monday within the month, whereas -1MO represents the last Monday of the month.  If an integer modifier is not present, it means all days of this type within the specified frequency.  For example, within a MONTHLY rule, MO represents all Mondays within the month. | [optional] 
**ByHour** | Pointer to **string** | Comma separated numbers between 0 to 23 | [optional] 
**ByMinute** | Pointer to **string** | Comma separated numbers between 0 to 59 | [optional] 
**ByMonth** | Pointer to **string** | Comma separated numbers between 1 to 12 | [optional] 
**ByMonthDay** | Pointer to **string** | Comma separated of numbers between -31 to 31, excluding 0.  For example, -10 represents the tenth to the last day of the month. | [optional] 
**ByOffset** | Pointer to **string** | Comma separated of numbers between -366 to 366, excluding 0.  Corresponds to the nth occurrence within the set of events specified by the rule.  It must only be used in conjunction with another byrule part.  For example \&quot;the last work day of the month\&quot; could be represented as: frequency&#x3D;MONTHLY;byDay&#x3D;MO,TU,WE,TH,FR;byOffset&#x3D;-1  Each byOffset value can include a positive (+n) or negative (-n) integer.  If present, this indicates the nth occurrence of the specific occurrence within the set of events specified by the rule. | [optional] 
**BySecond** | Pointer to **string** | Comma separated numbers between 0 to 59 | [optional] 
**ByWeekNumber** | Pointer to **string** | Comma separated of numbers between -53 to 53, excluding 0.  This corresponds to weeks according to week numbering.  A week is defined as a seven day period, starting on the day of the week defined to be the week start.  Week number one of the calendar year is the first week which contains at least four (4) days in that calendar year.  This rule part is only valid for YEARLY frequency.  For example, 3 represents the third week of the year. | [optional] 
**ByYearDay** | Pointer to **string** | Comma separated of numbers between -366 to 366, excluding 0.  For example, -1 represents the last day of the year (December 31st) and -306 represents the 306th to the last day of the year (March 1st). | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Frequency** | Pointer to **string** | Enum Type: &#x60;KalturaScheduleEventRecurrenceFrequency&#x60; | [optional] 
**Interval** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**TimeZone** | Pointer to **string** | TimeZone String | [optional] 
**Until** | Pointer to **int32** |  | [optional] 
**WeekStartDay** | Pointer to **string** | Enum Type: &#x60;KalturaScheduleEventRecurrenceDay&#x60; | [optional] 

## Methods

### NewKalturaScheduleEventRecurrence

`func NewKalturaScheduleEventRecurrence() *KalturaScheduleEventRecurrence`

NewKalturaScheduleEventRecurrence instantiates a new KalturaScheduleEventRecurrence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaScheduleEventRecurrenceWithDefaults

`func NewKalturaScheduleEventRecurrenceWithDefaults() *KalturaScheduleEventRecurrence`

NewKalturaScheduleEventRecurrenceWithDefaults instantiates a new KalturaScheduleEventRecurrence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetByDay

`func (o *KalturaScheduleEventRecurrence) GetByDay() string`

GetByDay returns the ByDay field if non-nil, zero value otherwise.

### GetByDayOk

`func (o *KalturaScheduleEventRecurrence) GetByDayOk() (*string, bool)`

GetByDayOk returns a tuple with the ByDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByDay

`func (o *KalturaScheduleEventRecurrence) SetByDay(v string)`

SetByDay sets ByDay field to given value.

### HasByDay

`func (o *KalturaScheduleEventRecurrence) HasByDay() bool`

HasByDay returns a boolean if a field has been set.

### GetByHour

`func (o *KalturaScheduleEventRecurrence) GetByHour() string`

GetByHour returns the ByHour field if non-nil, zero value otherwise.

### GetByHourOk

`func (o *KalturaScheduleEventRecurrence) GetByHourOk() (*string, bool)`

GetByHourOk returns a tuple with the ByHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByHour

`func (o *KalturaScheduleEventRecurrence) SetByHour(v string)`

SetByHour sets ByHour field to given value.

### HasByHour

`func (o *KalturaScheduleEventRecurrence) HasByHour() bool`

HasByHour returns a boolean if a field has been set.

### GetByMinute

`func (o *KalturaScheduleEventRecurrence) GetByMinute() string`

GetByMinute returns the ByMinute field if non-nil, zero value otherwise.

### GetByMinuteOk

`func (o *KalturaScheduleEventRecurrence) GetByMinuteOk() (*string, bool)`

GetByMinuteOk returns a tuple with the ByMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByMinute

`func (o *KalturaScheduleEventRecurrence) SetByMinute(v string)`

SetByMinute sets ByMinute field to given value.

### HasByMinute

`func (o *KalturaScheduleEventRecurrence) HasByMinute() bool`

HasByMinute returns a boolean if a field has been set.

### GetByMonth

`func (o *KalturaScheduleEventRecurrence) GetByMonth() string`

GetByMonth returns the ByMonth field if non-nil, zero value otherwise.

### GetByMonthOk

`func (o *KalturaScheduleEventRecurrence) GetByMonthOk() (*string, bool)`

GetByMonthOk returns a tuple with the ByMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByMonth

`func (o *KalturaScheduleEventRecurrence) SetByMonth(v string)`

SetByMonth sets ByMonth field to given value.

### HasByMonth

`func (o *KalturaScheduleEventRecurrence) HasByMonth() bool`

HasByMonth returns a boolean if a field has been set.

### GetByMonthDay

`func (o *KalturaScheduleEventRecurrence) GetByMonthDay() string`

GetByMonthDay returns the ByMonthDay field if non-nil, zero value otherwise.

### GetByMonthDayOk

`func (o *KalturaScheduleEventRecurrence) GetByMonthDayOk() (*string, bool)`

GetByMonthDayOk returns a tuple with the ByMonthDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByMonthDay

`func (o *KalturaScheduleEventRecurrence) SetByMonthDay(v string)`

SetByMonthDay sets ByMonthDay field to given value.

### HasByMonthDay

`func (o *KalturaScheduleEventRecurrence) HasByMonthDay() bool`

HasByMonthDay returns a boolean if a field has been set.

### GetByOffset

`func (o *KalturaScheduleEventRecurrence) GetByOffset() string`

GetByOffset returns the ByOffset field if non-nil, zero value otherwise.

### GetByOffsetOk

`func (o *KalturaScheduleEventRecurrence) GetByOffsetOk() (*string, bool)`

GetByOffsetOk returns a tuple with the ByOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByOffset

`func (o *KalturaScheduleEventRecurrence) SetByOffset(v string)`

SetByOffset sets ByOffset field to given value.

### HasByOffset

`func (o *KalturaScheduleEventRecurrence) HasByOffset() bool`

HasByOffset returns a boolean if a field has been set.

### GetBySecond

`func (o *KalturaScheduleEventRecurrence) GetBySecond() string`

GetBySecond returns the BySecond field if non-nil, zero value otherwise.

### GetBySecondOk

`func (o *KalturaScheduleEventRecurrence) GetBySecondOk() (*string, bool)`

GetBySecondOk returns a tuple with the BySecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBySecond

`func (o *KalturaScheduleEventRecurrence) SetBySecond(v string)`

SetBySecond sets BySecond field to given value.

### HasBySecond

`func (o *KalturaScheduleEventRecurrence) HasBySecond() bool`

HasBySecond returns a boolean if a field has been set.

### GetByWeekNumber

`func (o *KalturaScheduleEventRecurrence) GetByWeekNumber() string`

GetByWeekNumber returns the ByWeekNumber field if non-nil, zero value otherwise.

### GetByWeekNumberOk

`func (o *KalturaScheduleEventRecurrence) GetByWeekNumberOk() (*string, bool)`

GetByWeekNumberOk returns a tuple with the ByWeekNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByWeekNumber

`func (o *KalturaScheduleEventRecurrence) SetByWeekNumber(v string)`

SetByWeekNumber sets ByWeekNumber field to given value.

### HasByWeekNumber

`func (o *KalturaScheduleEventRecurrence) HasByWeekNumber() bool`

HasByWeekNumber returns a boolean if a field has been set.

### GetByYearDay

`func (o *KalturaScheduleEventRecurrence) GetByYearDay() string`

GetByYearDay returns the ByYearDay field if non-nil, zero value otherwise.

### GetByYearDayOk

`func (o *KalturaScheduleEventRecurrence) GetByYearDayOk() (*string, bool)`

GetByYearDayOk returns a tuple with the ByYearDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByYearDay

`func (o *KalturaScheduleEventRecurrence) SetByYearDay(v string)`

SetByYearDay sets ByYearDay field to given value.

### HasByYearDay

`func (o *KalturaScheduleEventRecurrence) HasByYearDay() bool`

HasByYearDay returns a boolean if a field has been set.

### GetCount

`func (o *KalturaScheduleEventRecurrence) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *KalturaScheduleEventRecurrence) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *KalturaScheduleEventRecurrence) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *KalturaScheduleEventRecurrence) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFrequency

`func (o *KalturaScheduleEventRecurrence) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *KalturaScheduleEventRecurrence) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *KalturaScheduleEventRecurrence) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *KalturaScheduleEventRecurrence) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetInterval

`func (o *KalturaScheduleEventRecurrence) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *KalturaScheduleEventRecurrence) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *KalturaScheduleEventRecurrence) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *KalturaScheduleEventRecurrence) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetName

`func (o *KalturaScheduleEventRecurrence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaScheduleEventRecurrence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaScheduleEventRecurrence) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaScheduleEventRecurrence) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeZone

`func (o *KalturaScheduleEventRecurrence) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *KalturaScheduleEventRecurrence) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *KalturaScheduleEventRecurrence) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *KalturaScheduleEventRecurrence) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUntil

`func (o *KalturaScheduleEventRecurrence) GetUntil() int32`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *KalturaScheduleEventRecurrence) GetUntilOk() (*int32, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *KalturaScheduleEventRecurrence) SetUntil(v int32)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *KalturaScheduleEventRecurrence) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetWeekStartDay

`func (o *KalturaScheduleEventRecurrence) GetWeekStartDay() string`

GetWeekStartDay returns the WeekStartDay field if non-nil, zero value otherwise.

### GetWeekStartDayOk

`func (o *KalturaScheduleEventRecurrence) GetWeekStartDayOk() (*string, bool)`

GetWeekStartDayOk returns a tuple with the WeekStartDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekStartDay

`func (o *KalturaScheduleEventRecurrence) SetWeekStartDay(v string)`

SetWeekStartDay sets WeekStartDay field to given value.

### HasWeekStartDay

`func (o *KalturaScheduleEventRecurrence) HasWeekStartDay() bool`

HasWeekStartDay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


