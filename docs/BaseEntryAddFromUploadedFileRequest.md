# BaseEntryAddFromUploadedFileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entry** | [**KalturaBaseEntry**](KalturaBaseEntry.md) |  | 
**Type** | Pointer to **string** |  | [optional] 
**UploadTokenId** | **string** |  | 

## Methods

### NewBaseEntryAddFromUploadedFileRequest

`func NewBaseEntryAddFromUploadedFileRequest(entry KalturaBaseEntry, uploadTokenId string, ) *BaseEntryAddFromUploadedFileRequest`

NewBaseEntryAddFromUploadedFileRequest instantiates a new BaseEntryAddFromUploadedFileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryAddFromUploadedFileRequestWithDefaults

`func NewBaseEntryAddFromUploadedFileRequestWithDefaults() *BaseEntryAddFromUploadedFileRequest`

NewBaseEntryAddFromUploadedFileRequestWithDefaults instantiates a new BaseEntryAddFromUploadedFileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntry

`func (o *BaseEntryAddFromUploadedFileRequest) GetEntry() KalturaBaseEntry`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *BaseEntryAddFromUploadedFileRequest) GetEntryOk() (*KalturaBaseEntry, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *BaseEntryAddFromUploadedFileRequest) SetEntry(v KalturaBaseEntry)`

SetEntry sets Entry field to given value.


### GetType

`func (o *BaseEntryAddFromUploadedFileRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseEntryAddFromUploadedFileRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseEntryAddFromUploadedFileRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BaseEntryAddFromUploadedFileRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUploadTokenId

`func (o *BaseEntryAddFromUploadedFileRequest) GetUploadTokenId() string`

GetUploadTokenId returns the UploadTokenId field if non-nil, zero value otherwise.

### GetUploadTokenIdOk

`func (o *BaseEntryAddFromUploadedFileRequest) GetUploadTokenIdOk() (*string, bool)`

GetUploadTokenIdOk returns a tuple with the UploadTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTokenId

`func (o *BaseEntryAddFromUploadedFileRequest) SetUploadTokenId(v string)`

SetUploadTokenId sets UploadTokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


