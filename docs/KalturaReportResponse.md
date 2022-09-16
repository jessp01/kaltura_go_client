# KalturaReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**[]KalturaString**](KalturaString.md) |  | [optional] 

## Methods

### NewKalturaReportResponse

`func NewKalturaReportResponse() *KalturaReportResponse`

NewKalturaReportResponse instantiates a new KalturaReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaReportResponseWithDefaults

`func NewKalturaReportResponseWithDefaults() *KalturaReportResponse`

NewKalturaReportResponseWithDefaults instantiates a new KalturaReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *KalturaReportResponse) GetColumns() string`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *KalturaReportResponse) GetColumnsOk() (*string, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *KalturaReportResponse) SetColumns(v string)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *KalturaReportResponse) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetResults

`func (o *KalturaReportResponse) GetResults() []KalturaString`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *KalturaReportResponse) GetResultsOk() (*[]KalturaString, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *KalturaReportResponse) SetResults(v []KalturaString)`

SetResults sets Results field to given value.

### HasResults

`func (o *KalturaReportResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


