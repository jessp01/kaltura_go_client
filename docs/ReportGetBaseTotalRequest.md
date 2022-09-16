# ReportGetBaseTotalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIds** | Pointer to **string** |  | [optional] 
**ReportInputFilter** | [**KalturaReportInputFilter**](KalturaReportInputFilter.md) |  | 
**ReportType** | **string** |  | 
**ResponseOptions** | Pointer to [**KalturaReportResponseOptions**](KalturaReportResponseOptions.md) |  | [optional] 

## Methods

### NewReportGetBaseTotalRequest

`func NewReportGetBaseTotalRequest(reportInputFilter KalturaReportInputFilter, reportType string, ) *ReportGetBaseTotalRequest`

NewReportGetBaseTotalRequest instantiates a new ReportGetBaseTotalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportGetBaseTotalRequestWithDefaults

`func NewReportGetBaseTotalRequestWithDefaults() *ReportGetBaseTotalRequest`

NewReportGetBaseTotalRequestWithDefaults instantiates a new ReportGetBaseTotalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectIds

`func (o *ReportGetBaseTotalRequest) GetObjectIds() string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *ReportGetBaseTotalRequest) GetObjectIdsOk() (*string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *ReportGetBaseTotalRequest) SetObjectIds(v string)`

SetObjectIds sets ObjectIds field to given value.

### HasObjectIds

`func (o *ReportGetBaseTotalRequest) HasObjectIds() bool`

HasObjectIds returns a boolean if a field has been set.

### GetReportInputFilter

`func (o *ReportGetBaseTotalRequest) GetReportInputFilter() KalturaReportInputFilter`

GetReportInputFilter returns the ReportInputFilter field if non-nil, zero value otherwise.

### GetReportInputFilterOk

`func (o *ReportGetBaseTotalRequest) GetReportInputFilterOk() (*KalturaReportInputFilter, bool)`

GetReportInputFilterOk returns a tuple with the ReportInputFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInputFilter

`func (o *ReportGetBaseTotalRequest) SetReportInputFilter(v KalturaReportInputFilter)`

SetReportInputFilter sets ReportInputFilter field to given value.


### GetReportType

`func (o *ReportGetBaseTotalRequest) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *ReportGetBaseTotalRequest) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *ReportGetBaseTotalRequest) SetReportType(v string)`

SetReportType sets ReportType field to given value.


### GetResponseOptions

`func (o *ReportGetBaseTotalRequest) GetResponseOptions() KalturaReportResponseOptions`

GetResponseOptions returns the ResponseOptions field if non-nil, zero value otherwise.

### GetResponseOptionsOk

`func (o *ReportGetBaseTotalRequest) GetResponseOptionsOk() (*KalturaReportResponseOptions, bool)`

GetResponseOptionsOk returns a tuple with the ResponseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseOptions

`func (o *ReportGetBaseTotalRequest) SetResponseOptions(v KalturaReportResponseOptions)`

SetResponseOptions sets ResponseOptions field to given value.

### HasResponseOptions

`func (o *ReportGetBaseTotalRequest) HasResponseOptions() bool`

HasResponseOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


