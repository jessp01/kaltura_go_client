# AttachmentAssetAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentAsset** | [**KalturaAttachmentAsset**](KalturaAttachmentAsset.md) |  | 
**EntryId** | **string** |  | 

## Methods

### NewAttachmentAssetAddRequest

`func NewAttachmentAssetAddRequest(attachmentAsset KalturaAttachmentAsset, entryId string, ) *AttachmentAssetAddRequest`

NewAttachmentAssetAddRequest instantiates a new AttachmentAssetAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentAssetAddRequestWithDefaults

`func NewAttachmentAssetAddRequestWithDefaults() *AttachmentAssetAddRequest`

NewAttachmentAssetAddRequestWithDefaults instantiates a new AttachmentAssetAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentAsset

`func (o *AttachmentAssetAddRequest) GetAttachmentAsset() KalturaAttachmentAsset`

GetAttachmentAsset returns the AttachmentAsset field if non-nil, zero value otherwise.

### GetAttachmentAssetOk

`func (o *AttachmentAssetAddRequest) GetAttachmentAssetOk() (*KalturaAttachmentAsset, bool)`

GetAttachmentAssetOk returns a tuple with the AttachmentAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentAsset

`func (o *AttachmentAssetAddRequest) SetAttachmentAsset(v KalturaAttachmentAsset)`

SetAttachmentAsset sets AttachmentAsset field to given value.


### GetEntryId

`func (o *AttachmentAssetAddRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *AttachmentAssetAddRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *AttachmentAssetAddRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


