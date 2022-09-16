# DocumentsAddFromFlavorAssetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentEntry** | Pointer to [**KalturaDocumentEntry**](KalturaDocumentEntry.md) |  | [optional] 
**SourceFlavorAssetId** | **string** |  | 

## Methods

### NewDocumentsAddFromFlavorAssetRequest

`func NewDocumentsAddFromFlavorAssetRequest(sourceFlavorAssetId string, ) *DocumentsAddFromFlavorAssetRequest`

NewDocumentsAddFromFlavorAssetRequest instantiates a new DocumentsAddFromFlavorAssetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsAddFromFlavorAssetRequestWithDefaults

`func NewDocumentsAddFromFlavorAssetRequestWithDefaults() *DocumentsAddFromFlavorAssetRequest`

NewDocumentsAddFromFlavorAssetRequestWithDefaults instantiates a new DocumentsAddFromFlavorAssetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentEntry

`func (o *DocumentsAddFromFlavorAssetRequest) GetDocumentEntry() KalturaDocumentEntry`

GetDocumentEntry returns the DocumentEntry field if non-nil, zero value otherwise.

### GetDocumentEntryOk

`func (o *DocumentsAddFromFlavorAssetRequest) GetDocumentEntryOk() (*KalturaDocumentEntry, bool)`

GetDocumentEntryOk returns a tuple with the DocumentEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentEntry

`func (o *DocumentsAddFromFlavorAssetRequest) SetDocumentEntry(v KalturaDocumentEntry)`

SetDocumentEntry sets DocumentEntry field to given value.

### HasDocumentEntry

`func (o *DocumentsAddFromFlavorAssetRequest) HasDocumentEntry() bool`

HasDocumentEntry returns a boolean if a field has been set.

### GetSourceFlavorAssetId

`func (o *DocumentsAddFromFlavorAssetRequest) GetSourceFlavorAssetId() string`

GetSourceFlavorAssetId returns the SourceFlavorAssetId field if non-nil, zero value otherwise.

### GetSourceFlavorAssetIdOk

`func (o *DocumentsAddFromFlavorAssetRequest) GetSourceFlavorAssetIdOk() (*string, bool)`

GetSourceFlavorAssetIdOk returns a tuple with the SourceFlavorAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFlavorAssetId

`func (o *DocumentsAddFromFlavorAssetRequest) SetSourceFlavorAssetId(v string)`

SetSourceFlavorAssetId sets SourceFlavorAssetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


