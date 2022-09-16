# DocumentsAddFromUploadedFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentEntry** | [**KalturaDocumentEntry**](KalturaDocumentEntry.md) |  | 
**UploadTokenId** | **string** |  | 

## Methods

### NewDocumentsAddFromUploadedFileRequest

`func NewDocumentsAddFromUploadedFileRequest(documentEntry KalturaDocumentEntry, uploadTokenId string, ) *DocumentsAddFromUploadedFileRequest`

NewDocumentsAddFromUploadedFileRequest instantiates a new DocumentsAddFromUploadedFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsAddFromUploadedFileRequestWithDefaults

`func NewDocumentsAddFromUploadedFileRequestWithDefaults() *DocumentsAddFromUploadedFileRequest`

NewDocumentsAddFromUploadedFileRequestWithDefaults instantiates a new DocumentsAddFromUploadedFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentEntry

`func (o *DocumentsAddFromUploadedFileRequest) GetDocumentEntry() KalturaDocumentEntry`

GetDocumentEntry returns the DocumentEntry field if non-nil, zero value otherwise.

### GetDocumentEntryOk

`func (o *DocumentsAddFromUploadedFileRequest) GetDocumentEntryOk() (*KalturaDocumentEntry, bool)`

GetDocumentEntryOk returns a tuple with the DocumentEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentEntry

`func (o *DocumentsAddFromUploadedFileRequest) SetDocumentEntry(v KalturaDocumentEntry)`

SetDocumentEntry sets DocumentEntry field to given value.


### GetUploadTokenId

`func (o *DocumentsAddFromUploadedFileRequest) GetUploadTokenId() string`

GetUploadTokenId returns the UploadTokenId field if non-nil, zero value otherwise.

### GetUploadTokenIdOk

`func (o *DocumentsAddFromUploadedFileRequest) GetUploadTokenIdOk() (*string, bool)`

GetUploadTokenIdOk returns a tuple with the UploadTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTokenId

`func (o *DocumentsAddFromUploadedFileRequest) SetUploadTokenId(v string)`

SetUploadTokenId sets UploadTokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


