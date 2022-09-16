# BatchCountBulkUploadEntriesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkUploadJobId** | **int32** |  | 
**BulkUploadObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewBatchCountBulkUploadEntriesRequest

`func NewBatchCountBulkUploadEntriesRequest(bulkUploadJobId int32, ) *BatchCountBulkUploadEntriesRequest`

NewBatchCountBulkUploadEntriesRequest instantiates a new BatchCountBulkUploadEntriesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCountBulkUploadEntriesRequestWithDefaults

`func NewBatchCountBulkUploadEntriesRequestWithDefaults() *BatchCountBulkUploadEntriesRequest`

NewBatchCountBulkUploadEntriesRequestWithDefaults instantiates a new BatchCountBulkUploadEntriesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkUploadJobId

`func (o *BatchCountBulkUploadEntriesRequest) GetBulkUploadJobId() int32`

GetBulkUploadJobId returns the BulkUploadJobId field if non-nil, zero value otherwise.

### GetBulkUploadJobIdOk

`func (o *BatchCountBulkUploadEntriesRequest) GetBulkUploadJobIdOk() (*int32, bool)`

GetBulkUploadJobIdOk returns a tuple with the BulkUploadJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadJobId

`func (o *BatchCountBulkUploadEntriesRequest) SetBulkUploadJobId(v int32)`

SetBulkUploadJobId sets BulkUploadJobId field to given value.


### GetBulkUploadObjectType

`func (o *BatchCountBulkUploadEntriesRequest) GetBulkUploadObjectType() string`

GetBulkUploadObjectType returns the BulkUploadObjectType field if non-nil, zero value otherwise.

### GetBulkUploadObjectTypeOk

`func (o *BatchCountBulkUploadEntriesRequest) GetBulkUploadObjectTypeOk() (*string, bool)`

GetBulkUploadObjectTypeOk returns a tuple with the BulkUploadObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadObjectType

`func (o *BatchCountBulkUploadEntriesRequest) SetBulkUploadObjectType(v string)`

SetBulkUploadObjectType sets BulkUploadObjectType field to given value.

### HasBulkUploadObjectType

`func (o *BatchCountBulkUploadEntriesRequest) HasBulkUploadObjectType() bool`

HasBulkUploadObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


