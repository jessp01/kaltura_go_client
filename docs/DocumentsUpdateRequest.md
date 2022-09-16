# DocumentsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentEntry** | [**KalturaDocumentEntry**](KalturaDocumentEntry.md) |  | 
**EntryId** | **string** |  | 

## Methods

### NewDocumentsUpdateRequest

`func NewDocumentsUpdateRequest(documentEntry KalturaDocumentEntry, entryId string, ) *DocumentsUpdateRequest`

NewDocumentsUpdateRequest instantiates a new DocumentsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsUpdateRequestWithDefaults

`func NewDocumentsUpdateRequestWithDefaults() *DocumentsUpdateRequest`

NewDocumentsUpdateRequestWithDefaults instantiates a new DocumentsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentEntry

`func (o *DocumentsUpdateRequest) GetDocumentEntry() KalturaDocumentEntry`

GetDocumentEntry returns the DocumentEntry field if non-nil, zero value otherwise.

### GetDocumentEntryOk

`func (o *DocumentsUpdateRequest) GetDocumentEntryOk() (*KalturaDocumentEntry, bool)`

GetDocumentEntryOk returns a tuple with the DocumentEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentEntry

`func (o *DocumentsUpdateRequest) SetDocumentEntry(v KalturaDocumentEntry)`

SetDocumentEntry sets DocumentEntry field to given value.


### GetEntryId

`func (o *DocumentsUpdateRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *DocumentsUpdateRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *DocumentsUpdateRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


