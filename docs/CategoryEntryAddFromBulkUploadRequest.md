# CategoryEntryAddFromBulkUploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkUploadCategoryEntryData** | Pointer to [**KalturaBulkUploadCategoryEntryData**](KalturaBulkUploadCategoryEntryData.md) |  | [optional] 
**BulkUploadData** | [**KalturaBulkServiceData**](KalturaBulkServiceData.md) |  | 

## Methods

### NewCategoryEntryAddFromBulkUploadRequest

`func NewCategoryEntryAddFromBulkUploadRequest(bulkUploadData KalturaBulkServiceData, ) *CategoryEntryAddFromBulkUploadRequest`

NewCategoryEntryAddFromBulkUploadRequest instantiates a new CategoryEntryAddFromBulkUploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryEntryAddFromBulkUploadRequestWithDefaults

`func NewCategoryEntryAddFromBulkUploadRequestWithDefaults() *CategoryEntryAddFromBulkUploadRequest`

NewCategoryEntryAddFromBulkUploadRequestWithDefaults instantiates a new CategoryEntryAddFromBulkUploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkUploadCategoryEntryData

`func (o *CategoryEntryAddFromBulkUploadRequest) GetBulkUploadCategoryEntryData() KalturaBulkUploadCategoryEntryData`

GetBulkUploadCategoryEntryData returns the BulkUploadCategoryEntryData field if non-nil, zero value otherwise.

### GetBulkUploadCategoryEntryDataOk

`func (o *CategoryEntryAddFromBulkUploadRequest) GetBulkUploadCategoryEntryDataOk() (*KalturaBulkUploadCategoryEntryData, bool)`

GetBulkUploadCategoryEntryDataOk returns a tuple with the BulkUploadCategoryEntryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadCategoryEntryData

`func (o *CategoryEntryAddFromBulkUploadRequest) SetBulkUploadCategoryEntryData(v KalturaBulkUploadCategoryEntryData)`

SetBulkUploadCategoryEntryData sets BulkUploadCategoryEntryData field to given value.

### HasBulkUploadCategoryEntryData

`func (o *CategoryEntryAddFromBulkUploadRequest) HasBulkUploadCategoryEntryData() bool`

HasBulkUploadCategoryEntryData returns a boolean if a field has been set.

### GetBulkUploadData

`func (o *CategoryEntryAddFromBulkUploadRequest) GetBulkUploadData() KalturaBulkServiceData`

GetBulkUploadData returns the BulkUploadData field if non-nil, zero value otherwise.

### GetBulkUploadDataOk

`func (o *CategoryEntryAddFromBulkUploadRequest) GetBulkUploadDataOk() (*KalturaBulkServiceData, bool)`

GetBulkUploadDataOk returns a tuple with the BulkUploadData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadData

`func (o *CategoryEntryAddFromBulkUploadRequest) SetBulkUploadData(v KalturaBulkServiceData)`

SetBulkUploadData sets BulkUploadData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


