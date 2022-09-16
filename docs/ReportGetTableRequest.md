# ReportGetTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIds** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Pager** | [**KalturaFilterPager**](KalturaFilterPager.md) |  | 
**ReportInputFilter** | [**KalturaReportInputFilter**](KalturaReportInputFilter.md) |  | 
**ReportType** | **string** |  | 
**ResponseOptions** | Pointer to [**KalturaReportResponseOptions**](KalturaReportResponseOptions.md) |  | [optional] 

## Methods

### NewReportGetTableRequest

`func NewReportGetTableRequest(pager KalturaFilterPager, reportInputFilter KalturaReportInputFilter, reportType string, ) *ReportGetTableRequest`

NewReportGetTableRequest instantiates a new ReportGetTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportGetTableRequestWithDefaults

`func NewReportGetTableRequestWithDefaults() *ReportGetTableRequest`

NewReportGetTableRequestWithDefaults instantiates a new ReportGetTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectIds

`func (o *ReportGetTableRequest) GetObjectIds() string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *ReportGetTableRequest) GetObjectIdsOk() (*string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *ReportGetTableRequest) SetObjectIds(v string)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *ReportGetTableRequest) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### GetOrder

`func (o *ReportGetTableRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReportGetTableRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReportGetTableRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReportGetTableRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPager

`func (o *ReportGetTableRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ReportGetTableRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ReportGetTableRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.


### GetReportInputFilter

`func (o *ReportGetTableRequest) GetReportInputFilter() KalturaReportInputFilter`

GetReportInputFilter returns the ReportInputFilter field if non-nil, zero value otherwise.

### GetReportInputFilterOk

`func (o *ReportGetTableRequest) GetReportInputFilterOk() (*KalturaReportInputFilter, bool)`

GetReportInputFilterOk returns a tuple with the ReportInputFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInputFilter

`func (o *ReportGetTableRequest) SetReportInputFilter(v KalturaReportInputFilter)`

SetReportInputFilter sets ReportInputFilter field to given value.


### GetReportType

`func (o *ReportGetTableRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportGetTableRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportGetTableRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetResponseOptions

`func (o *ReportGetTableRequest) GetResponseOptions() KalturaReportResponseOptions`

GetResponseOptions returns the ResponseOptions field if non-nil, zero value otherwise.

### GetResponseOptionsOk

`func (o *ReportGetTableRequest) GetResponseOptionsOk() (*KalturaReportResponseOptions, bool)`

GetResponseOptionsOk returns a tuple with the ResponseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseOptions

`func (o *ReportGetTableRequest) SetResponseOptions(v KalturaReportResponseOptions)`

SetResponseOptions sets ResponseOptions field to given value.

### HasResponseOptions

`func (o *ReportGetTableRequest) HasResponseOptions() bool`

HasResponseOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


