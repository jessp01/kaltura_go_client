# ReportGetGraphsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | Pointer to **string** |  | [optional] 
**ObjectIds** | Pointer to **string** |  | [optional] 
**ReportInputFilter** | [**KalturaReportInputFilter**](KalturaReportInputFilter.md) |  | 
**ReportType** | **string** |  | 
**ResponseOptions** | Pointer to [**KalturaReportResponseOptions**](KalturaReportResponseOptions.md) |  | [optional] 

## Methods

### NewReportGetGraphsRequest

`func NewReportGetGraphsRequest(reportInputFilter KalturaReportInputFilter, reportType string, ) *ReportGetGraphsRequest`

NewReportGetGraphsRequest instantiates a new ReportGetGraphsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportGetGraphsRequestWithDefaults

`func NewReportGetGraphsRequestWithDefaults() *ReportGetGraphsRequest`

NewReportGetGraphsRequestWithDefaults instantiates a new ReportGetGraphsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *ReportGetGraphsRequest) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ReportGetGraphsRequest) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ReportGetGraphsRequest) SetDimension(v string)`

SetDimension sets Dimension field to given value.

### HasDimension

`func (o *ReportGetGraphsRequest) HasDimension() bool`

HasDimension returns a boolean if a field has been set.

### GetObjectIds

`func (o *ReportGetGraphsRequest) GetObjectIds() string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *ReportGetGraphsRequest) GetObjectIdsOk() (*string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *ReportGetGraphsRequest) SetObjectIds(v string)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *ReportGetGraphsRequest) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### GetReportInputFilter

`func (o *ReportGetGraphsRequest) GetReportInputFilter() KalturaReportInputFilter`

GetReportInputFilter returns the ReportInputFilter field if non-nil, zero value otherwise.

### GetReportInputFilterOk

`func (o *ReportGetGraphsRequest) GetReportInputFilterOk() (*KalturaReportInputFilter, bool)`

GetReportInputFilterOk returns a tuple with the ReportInputFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInputFilter

`func (o *ReportGetGraphsRequest) SetReportInputFilter(v KalturaReportInputFilter)`

SetReportInputFilter sets ReportInputFilter field to given value.


### GetReportType

`func (o *ReportGetGraphsRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportGetGraphsRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportGetGraphsRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetResponseOptions

`func (o *ReportGetGraphsRequest) GetResponseOptions() KalturaReportResponseOptions`

GetResponseOptions returns the ResponseOptions field if non-nil, zero value otherwise.

### GetResponseOptionsOk

`func (o *ReportGetGraphsRequest) GetResponseOptionsOk() (*KalturaReportResponseOptions, bool)`

GetResponseOptionsOk returns a tuple with the ResponseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseOptions

`func (o *ReportGetGraphsRequest) SetResponseOptions(v KalturaReportResponseOptions)`

SetResponseOptions sets ResponseOptions field to given value.

### HasResponseOptions

`func (o *ReportGetGraphsRequest) HasResponseOptions() bool`

HasResponseOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


