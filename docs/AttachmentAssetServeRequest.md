# AttachmentAssetServeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentAssetId** | **string** |  | 
**ServeOptions** | Pointer to [**KalturaAttachmentServeOptions**](KalturaAttachmentServeOptions.md) |  | [optional] 

## Methods

### NewAttachmentAssetServeRequest

`func NewAttachmentAssetServeRequest(attachmentAssetId string, ) *AttachmentAssetServeRequest`

NewAttachmentAssetServeRequest instantiates a new AttachmentAssetServeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentAssetServeRequestWithDefaults

`func NewAttachmentAssetServeRequestWithDefaults() *AttachmentAssetServeRequest`

NewAttachmentAssetServeRequestWithDefaults instantiates a new AttachmentAssetServeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentAssetId

`func (o *AttachmentAssetServeRequest) GetAttachmentAssetId() string`

GetAttachmentAssetId returns the AttachmentAssetId field if non-nil, zero value otherwise.

### GetAttachmentAssetIdOk

`func (o *AttachmentAssetServeRequest) GetAttachmentAssetIdOk() (*string, bool)`

GetAttachmentAssetIdOk returns a tuple with the AttachmentAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentAssetId

`func (o *AttachmentAssetServeRequest) SetAttachmentAssetId(v string)`

SetAttachmentAssetId sets AttachmentAssetId field to given value.


### GetServeOptions

`func (o *AttachmentAssetServeRequest) GetServeOptions() KalturaAttachmentServeOptions`

GetServeOptions returns the ServeOptions field if non-nil, zero value otherwise.

### GetServeOptionsOk

`func (o *AttachmentAssetServeRequest) GetServeOptionsOk() (*KalturaAttachmentServeOptions, bool)`

GetServeOptionsOk returns a tuple with the ServeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServeOptions

`func (o *AttachmentAssetServeRequest) SetServeOptions(v KalturaAttachmentServeOptions)`

SetServeOptions sets ServeOptions field to given value.

### HasServeOptions

`func (o *AttachmentAssetServeRequest) HasServeOptions() bool`

HasServeOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


