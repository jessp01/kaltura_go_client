# ReportGetCsvRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedFields** | Pointer to **string** |  | [optional] 
**Id** | **int32** |  | 
**Params** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 

## Methods

### NewReportGetCsvRequest

`func NewReportGetCsvRequest(id int32, ) *ReportGetCsvRequest`

NewReportGetCsvRequest instantiates a new ReportGetCsvRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportGetCsvRequestWithDefaults

`func NewReportGetCsvRequestWithDefaults() *ReportGetCsvRequest`

NewReportGetCsvRequestWithDefaults instantiates a new ReportGetCsvRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedFields

`func (o *ReportGetCsvRequest) GetExcludedFields() string`

GetExcludedFields returns the ExcludedFields field if non-nil, zero value otherwise.

### GetExcludedFieldsOk

`func (o *ReportGetCsvRequest) GetExcludedFieldsOk() (*string, bool)`

GetExcludedFieldsOk returns a tuple with the ExcludedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedFields

`func (o *ReportGetCsvRequest) SetExcludedFields(v string)`

SetExcludedFields sets ExcludedFields field to given value.

### HasExcludedFields

`func (o *ReportGetCsvRequest) HasExcludedFields() bool`

HasExcludedFields returns a boolean if a field has been set.

### GetId

`func (o *ReportGetCsvRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportGetCsvRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportGetCsvRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetParams

`func (o *ReportGetCsvRequest) GetParams() []KalturaKeyValue`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ReportGetCsvRequest) GetParamsOk() (*[]KalturaKeyValue, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ReportGetCsvRequest) SetParams(v []KalturaKeyValue)`

SetParams sets Params field to given value.

### HasParams

`func (o *ReportGetCsvRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


