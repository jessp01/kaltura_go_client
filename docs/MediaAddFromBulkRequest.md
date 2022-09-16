# MediaAddFromBulkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkUploadId** | **int32** |  | 
**MediaEntry** | [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | 
**Url** | **string** |  | 

## Methods

### NewMediaAddFromBulkRequest

`func NewMediaAddFromBulkRequest(bulkUploadId int32, mediaEntry KalturaMediaEntry, url string, ) *MediaAddFromBulkRequest`

NewMediaAddFromBulkRequest instantiates a new MediaAddFromBulkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddFromBulkRequestWithDefaults

`func NewMediaAddFromBulkRequestWithDefaults() *MediaAddFromBulkRequest`

NewMediaAddFromBulkRequestWithDefaults instantiates a new MediaAddFromBulkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkUploadId

`func (o *MediaAddFromBulkRequest) GetBulkUploadId() int32`

GetBulkUploadId returns the BulkUploadId field if non-nil, zero value otherwise.

### GetBulkUploadIdOk

`func (o *MediaAddFromBulkRequest) GetBulkUploadIdOk() (*int32, bool)`

GetBulkUploadIdOk returns a tuple with the BulkUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkUploadId

`func (o *MediaAddFromBulkRequest) SetBulkUploadId(v int32)`

SetBulkUploadId sets BulkUploadId field to given value.


### GetMediaEntry

`func (o *MediaAddFromBulkRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaAddFromBulkRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaAddFromBulkRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.


### GetUrl

`func (o *MediaAddFromBulkRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaAddFromBulkRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaAddFromBulkRequest) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


