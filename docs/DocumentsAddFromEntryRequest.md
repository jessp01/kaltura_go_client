# DocumentsAddFromEntryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentEntry** | Pointer to [**KalturaDocumentEntry**](KalturaDocumentEntry.md) |  | [optional] 
**SourceEntryId** | **string** |  | 
**SourceFlavorParamsId** | Pointer to **int32** |  | [optional] 

## Methods

### NewDocumentsAddFromEntryRequest

`func NewDocumentsAddFromEntryRequest(sourceEntryId string, ) *DocumentsAddFromEntryRequest`

NewDocumentsAddFromEntryRequest instantiates a new DocumentsAddFromEntryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsAddFromEntryRequestWithDefaults

`func NewDocumentsAddFromEntryRequestWithDefaults() *DocumentsAddFromEntryRequest`

NewDocumentsAddFromEntryRequestWithDefaults instantiates a new DocumentsAddFromEntryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentEntry

`func (o *DocumentsAddFromEntryRequest) GetDocumentEntry() KalturaDocumentEntry`

GetDocumentEntry returns the DocumentEntry field if non-nil, zero value otherwise.

### GetDocumentEntryOk

`func (o *DocumentsAddFromEntryRequest) GetDocumentEntryOk() (*KalturaDocumentEntry, bool)`

GetDocumentEntryOk returns a tuple with the DocumentEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentEntry

`func (o *DocumentsAddFromEntryRequest) SetDocumentEntry(v KalturaDocumentEntry)`

SetDocumentEntry sets DocumentEntry field to given value.

### HasDocumentEntry

`func (o *DocumentsAddFromEntryRequest) HasDocumentEntry() bool`

HasDocumentEntry returns a boolean if a field has been set.

### GetSourceEntryId

`func (o *DocumentsAddFromEntryRequest) GetSourceEntryId() string`

GetSourceEntryId returns the SourceEntryId field if non-nil, zero value otherwise.

### GetSourceEntryIdOk

`func (o *DocumentsAddFromEntryRequest) GetSourceEntryIdOk() (*string, bool)`

GetSourceEntryIdOk returns a tuple with the SourceEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEntryId

`func (o *DocumentsAddFromEntryRequest) SetSourceEntryId(v string)`

SetSourceEntryId sets SourceEntryId field to given value.


### GetSourceFlavorParamsId

`func (o *DocumentsAddFromEntryRequest) GetSourceFlavorParamsId() int32`

GetSourceFlavorParamsId returns the SourceFlavorParamsId field if non-nil, zero value otherwise.

### GetSourceFlavorParamsIdOk

`func (o *DocumentsAddFromEntryRequest) GetSourceFlavorParamsIdOk() (*int32, bool)`

GetSourceFlavorParamsIdOk returns a tuple with the SourceFlavorParamsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFlavorParamsId

`func (o *DocumentsAddFromEntryRequest) SetSourceFlavorParamsId(v int32)`

SetSourceFlavorParamsId sets SourceFlavorParamsId field to given value.

### HasSourceFlavorParamsId

`func (o *DocumentsAddFromEntryRequest) HasSourceFlavorParamsId() bool`

HasSourceFlavorParamsId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


