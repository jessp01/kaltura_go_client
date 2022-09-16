# LiveReportsGetEventsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaLiveReportInputFilter**](KalturaLiveReportInputFilter.md) |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**ReportType** | **string** |  | 

## Methods

### NewLiveReportsGetEventsRequest

`func NewLiveReportsGetEventsRequest(reportType string, ) *LiveReportsGetEventsRequest`

NewLiveReportsGetEventsRequest instantiates a new LiveReportsGetEventsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveReportsGetEventsRequestWithDefaults

`func NewLiveReportsGetEventsRequestWithDefaults() *LiveReportsGetEventsRequest`

NewLiveReportsGetEventsRequestWithDefaults instantiates a new LiveReportsGetEventsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LiveReportsGetEventsRequest) GetFilter() KalturaLiveReportInputFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LiveReportsGetEventsRequest) GetFilterOk() (*KalturaLiveReportInputFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LiveReportsGetEventsRequest) SetFilter(v KalturaLiveReportInputFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LiveReportsGetEventsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPager

`func (o *LiveReportsGetEventsRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *LiveReportsGetEventsRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *LiveReportsGetEventsRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *LiveReportsGetEventsRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetReportType

`func (o *LiveReportsGetEventsRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *LiveReportsGetEventsRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *LiveReportsGetEventsRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


