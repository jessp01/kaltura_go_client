# KalturaAnalyticsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to **string** | Comma separated dimensions list | [optional] 
**Filters** | Pointer to [**[]KalturaReportFilter**](KalturaReportFilter.md) |  | [optional] 
**FromTime** | Pointer to **string** | Query start time (in local time) MM/dd/yyyy HH:mi | [optional] 
**Metrics** | Pointer to **string** | Comma separated metrics list | [optional] 
**OrderBy** | Pointer to **string** | Query order by metric/dimension | [optional] 
**ToTime** | Pointer to **string** | Query end time (in local time) MM/dd/yyyy HH:mi | [optional] 
**UtcOffset** | Pointer to **float32** | Timezone offset from UTC (in minutes) | [optional] 

## Methods

### NewKalturaAnalyticsFilter

`func NewKalturaAnalyticsFilter() *KalturaAnalyticsFilter`

NewKalturaAnalyticsFilter instantiates a new KalturaAnalyticsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAnalyticsFilterWithDefaults

`func NewKalturaAnalyticsFilterWithDefaults() *KalturaAnalyticsFilter`

NewKalturaAnalyticsFilterWithDefaults instantiates a new KalturaAnalyticsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *KalturaAnalyticsFilter) GetDimensions() string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *KalturaAnalyticsFilter) GetDimensionsOk() (*string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *KalturaAnalyticsFilter) SetDimensions(v string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *KalturaAnalyticsFilter) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetFilters

`func (o *KalturaAnalyticsFilter) GetFilters() []KalturaReportFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *KalturaAnalyticsFilter) GetFiltersOk() (*[]KalturaReportFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *KalturaAnalyticsFilter) SetFilters(v []KalturaReportFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *KalturaAnalyticsFilter) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFromTime

`func (o *KalturaAnalyticsFilter) GetFromTime() string`

GetFromTime returns the FromTime field if non-nil, zero value otherwise.

### GetFromTimeOk

`func (o *KalturaAnalyticsFilter) GetFromTimeOk() (*string, bool)`

GetFromTimeOk returns a tuple with the FromTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromTime

`func (o *KalturaAnalyticsFilter) SetFromTime(v string)`

SetFromTime sets FromTime field to given value.

### HasFromTime

`func (o *KalturaAnalyticsFilter) HasFromTime() bool`

HasFromTime returns a boolean if a field has been set.

### GetMetrics

`func (o *KalturaAnalyticsFilter) GetMetrics() string`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *KalturaAnalyticsFilter) GetMetricsOk() (*string, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *KalturaAnalyticsFilter) SetMetrics(v string)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *KalturaAnalyticsFilter) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetOrderBy

`func (o *KalturaAnalyticsFilter) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *KalturaAnalyticsFilter) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *KalturaAnalyticsFilter) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *KalturaAnalyticsFilter) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetToTime

`func (o *KalturaAnalyticsFilter) GetToTime() string`

GetToTime returns the ToTime field if non-nil, zero value otherwise.

### GetToTimeOk

`func (o *KalturaAnalyticsFilter) GetToTimeOk() (*string, bool)`

GetToTimeOk returns a tuple with the ToTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToTime

`func (o *KalturaAnalyticsFilter) SetToTime(v string)`

SetToTime sets ToTime field to given value.

### HasToTime

`func (o *KalturaAnalyticsFilter) HasToTime() bool`

HasToTime returns a boolean if a field has been set.

### GetUtcOffset

`func (o *KalturaAnalyticsFilter) GetUtcOffset() float32`

GetUtcOffset returns the UtcOffset field if non-nil, zero value otherwise.

### GetUtcOffsetOk

`func (o *KalturaAnalyticsFilter) GetUtcOffsetOk() (*float32, bool)`

GetUtcOffsetOk returns a tuple with the UtcOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcOffset

`func (o *KalturaAnalyticsFilter) SetUtcOffset(v float32)`

SetUtcOffset sets UtcOffset field to given value.

### HasUtcOffset

`func (o *KalturaAnalyticsFilter) HasUtcOffset() bool`

HasUtcOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


