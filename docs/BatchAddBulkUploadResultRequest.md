# BatchAddBulkUploadResultRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkUploadResult** | [**KalturaBulkUploadResult**](KalturaBulkUploadResult.md) |  | 
**PluginDataArray** | Pointer to [**[]KalturaBulkUploadPluginData**](KalturaBulkUploadPluginData.md) |  | [optional] 

## Methods

### NewBatchAddBulkUploadResultRequest

`func NewBatchAddBulkUploadResultRequest(bulkUploadResult KalturaBulkUploadResult, ) *BatchAddBulkUploadResultRequest`

NewBatchAddBulkUploadResultRequest instantiates a new BatchAddBulkUploadResultRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchAddBulkUploadResultRequestWithDefaults

`func NewBatchAddBulkUploadResultRequestWithDefaults() *BatchAddBulkUploadResultRequest`

NewBatchAddBulkUploadResultRequestWithDefaults instantiates a new BatchAddBulkUploadResultRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkUploadResult

`func (o *BatchAddBulkUploadResultRequest) GetBulkUploadResult() KalturaBulkUploadResult`

GetBulkUploadResult returns the BulkUploadResult field if non-nil, zero value otherwise.

### GetBulkUploadResultOk

`func (o *BatchAddBulkUploadResultRequest) GetBulkUploadResultOk() (*KalturaBulkUploadResult, bool)`

GetBulkUploadResultOk returns a tuple with the BulkUploadResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadResult

`func (o *BatchAddBulkUploadResultRequest) SetBulkUploadResult(v KalturaBulkUploadResult)`

SetBulkUploadResult sets BulkUploadResult field to given value.


### GetPluginDataArray

`func (o *BatchAddBulkUploadResultRequest) GetPluginDataArray() []KalturaBulkUploadPluginData`

GetPluginDataArray returns the PluginDataArray field if non-nil, zero value otherwise.

### GetPluginDataArrayOk

`func (o *BatchAddBulkUploadResultRequest) GetPluginDataArrayOk() (*[]KalturaBulkUploadPluginData, bool)`

GetPluginDataArrayOk returns a tuple with the PluginDataArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginDataArray

`func (o *BatchAddBulkUploadResultRequest) SetPluginDataArray(v []KalturaBulkUploadPluginData)`

SetPluginDataArray sets PluginDataArray field to given value.

### HasPluginDataArray

`func (o *BatchAddBulkUploadResultRequest) HasPluginDataArray() bool`

HasPluginDataArray returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


