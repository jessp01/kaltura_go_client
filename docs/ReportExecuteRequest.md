# ReportExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Params** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 

## Methods

### NewReportExecuteRequest

`func NewReportExecuteRequest(id int32, ) *ReportExecuteRequest`

NewReportExecuteRequest instantiates a new ReportExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportExecuteRequestWithDefaults

`func NewReportExecuteRequestWithDefaults() *ReportExecuteRequest`

NewReportExecuteRequestWithDefaults instantiates a new ReportExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportExecuteRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportExecuteRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportExecuteRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetParams

`func (o *ReportExecuteRequest) GetParams() []KalturaKeyValue`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ReportExecuteRequest) GetParamsOk() (*[]KalturaKeyValue, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ReportExecuteRequest) SetParams(v []KalturaKeyValue)`

SetParams sets Params field to given value.

### HasParams

`func (o *ReportExecuteRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


