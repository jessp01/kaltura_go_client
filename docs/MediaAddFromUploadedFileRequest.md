# MediaAddFromUploadedFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaEntry** | [**KalturaMediaEntry**](KalturaMediaEntry.md) |  | 
**UploadTokenId** | **string** |  | 

## Methods

### NewMediaAddFromUploadedFileRequest

`func NewMediaAddFromUploadedFileRequest(mediaEntry KalturaMediaEntry, uploadTokenId string, ) *MediaAddFromUploadedFileRequest`

NewMediaAddFromUploadedFileRequest instantiates a new MediaAddFromUploadedFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddFromUploadedFileRequestWithDefaults

`func NewMediaAddFromUploadedFileRequestWithDefaults() *MediaAddFromUploadedFileRequest`

NewMediaAddFromUploadedFileRequestWithDefaults instantiates a new MediaAddFromUploadedFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaEntry

`func (o *MediaAddFromUploadedFileRequest) GetMediaEntry() KalturaMediaEntry`

GetMediaEntry returns the MediaEntry field if non-nil, zero value otherwise.

### GetMediaEntryOk

`func (o *MediaAddFromUploadedFileRequest) GetMediaEntryOk() (*KalturaMediaEntry, bool)`

GetMediaEntryOk returns a tuple with the MediaEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaEntry

`func (o *MediaAddFromUploadedFileRequest) SetMediaEntry(v KalturaMediaEntry)`

SetMediaEntry sets MediaEntry field to given value.


### GetUploadTokenId

`func (o *MediaAddFromUploadedFileRequest) GetUploadTokenId() string`

GetUploadTokenId returns the UploadTokenId field if non-nil, zero value otherwise.

### GetUploadTokenIdOk

`func (o *MediaAddFromUploadedFileRequest) GetUploadTokenIdOk() (*string, bool)`

GetUploadTokenIdOk returns a tuple with the UploadTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTokenId

`func (o *MediaAddFromUploadedFileRequest) SetUploadTokenId(v string)`

SetUploadTokenId sets UploadTokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


