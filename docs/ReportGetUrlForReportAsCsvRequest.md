# ReportGetUrlForReportAsCsvRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** |  | [optional] 
**Headers** | **string** |  | 
**ObjectIds** | Pointer to **string** |  | [optional] 
**Order** | Pointer to **string** |  | [optional] 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**ReportInputFilter** | [**KalturaReportInputFilter**](KalturaReportInputFilter.md) |  | 
**ReportText** | **string** |  | 
**ReportTitle** | **string** |  | 
**ReportType** | **string** |  | 
**ResponseOptions** | Pointer to [**KalturaReportResponseOptions**](KalturaReportResponseOptions.md) |  | [optional] 

## Methods

### NewReportGetUrlForReportAsCsvRequest

`func NewReportGetUrlForReportAsCsvRequest(headers string, reportInputFilter KalturaReportInputFilter, reportText string, reportTitle string, reportType string, ) *ReportGetUrlForReportAsCsvRequest`

NewReportGetUrlForReportAsCsvRequest instantiates a new ReportGetUrlForReportAsCsvRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportGetUrlForReportAsCsvRequestWithDefaults

`func NewReportGetUrlForReportAsCsvRequestWithDefaults() *ReportGetUrlForReportAsCsvRequest`

NewReportGetUrlForReportAsCsvRequestWithDefaults instantiates a new ReportGetUrlForReportAsCsvRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *ReportGetUrlForReportAsCsvRequest) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ReportGetUrlForReportAsCsvRequest) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *ReportGetUrlForReportAsCsvRequest) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetHeaders

`func (o *ReportGetUrlForReportAsCsvRequest) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ReportGetUrlForReportAsCsvRequest) SetHeaders(v string)`

SetHeaders sets Headers field to given value.


### GetObjectIds

`func (o *ReportGetUrlForReportAsCsvRequest) GetObjectIds() string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetObjectIdsOk() (*string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *ReportGetUrlForReportAsCsvRequest) SetObjectIds(v string)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *ReportGetUrlForReportAsCsvRequest) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### GetOrder

`func (o *ReportGetUrlForReportAsCsvRequest) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ReportGetUrlForReportAsCsvRequest) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ReportGetUrlForReportAsCsvRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetPager

`func (o *ReportGetUrlForReportAsCsvRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *ReportGetUrlForReportAsCsvRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *ReportGetUrlForReportAsCsvRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetReportInputFilter

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportInputFilter() KalturaReportInputFilter`

GetReportInputFilter returns the ReportInputFilter field if non-nil, zero value otherwise.

### GetReportInputFilterOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportInputFilterOk() (*KalturaReportInputFilter, bool)`

GetReportInputFilterOk returns a tuple with the ReportInputFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInputFilter

`func (o *ReportGetUrlForReportAsCsvRequest) SetReportInputFilter(v KalturaReportInputFilter)`

SetReportInputFilter sets ReportInputFilter field to given value.


### GetReportText

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportText() string`

GetReportText returns the ReportText field if non-nil, zero value otherwise.

### GetReportTextOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportTextOk() (*string, bool)`

GetReportTextOk returns a tuple with the ReportText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportText

`func (o *ReportGetUrlForReportAsCsvRequest) SetReportText(v string)`

SetReportText sets ReportText field to given value.


### GetReportTitle

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *ReportGetUrlForReportAsCsvRequest) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.


### GetReportType

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportGetUrlForReportAsCsvRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetResponseOptions

`func (o *ReportGetUrlForReportAsCsvRequest) GetResponseOptions() KalturaReportResponseOptions`

GetResponseOptions returns the ResponseOptions field if non-nil, zero value otherwise.

### GetResponseOptionsOk

`func (o *ReportGetUrlForReportAsCsvRequest) GetResponseOptionsOk() (*KalturaReportResponseOptions, bool)`

GetResponseOptionsOk returns a tuple with the ResponseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseOptions

`func (o *ReportGetUrlForReportAsCsvRequest) SetResponseOptions(v KalturaReportResponseOptions)`

SetResponseOptions sets ResponseOptions field to given value.

### HasResponseOptions

`func (o *ReportGetUrlForReportAsCsvRequest) HasResponseOptions() bool`

HasResponseOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


