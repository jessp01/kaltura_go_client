# KalturaUploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Enum Type: &#x60;KalturaUploadErrorCode&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **int32** |  | [optional] 
**UploadTokenId** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaUploadResponse

`func NewKalturaUploadResponse() *KalturaUploadResponse`

NewKalturaUploadResponse instantiates a new KalturaUploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaUploadResponseWithDefaults

`func NewKalturaUploadResponseWithDefaults() *KalturaUploadResponse`

NewKalturaUploadResponseWithDefaults instantiates a new KalturaUploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *KalturaUploadResponse) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *KalturaUploadResponse) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *KalturaUploadResponse) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *KalturaUploadResponse) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *KalturaUploadResponse) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *KalturaUploadResponse) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *KalturaUploadResponse) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *KalturaUploadResponse) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetFileSize

`func (o *KalturaUploadResponse) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *KalturaUploadResponse) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *KalturaUploadResponse) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *KalturaUploadResponse) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetUploadTokenId

`func (o *KalturaUploadResponse) GetUploadTokenId() string`

GetUploadTokenId returns the UploadTokenId field if non-nil, zero value otherwise.

### GetUploadTokenIdOk

`func (o *KalturaUploadResponse) GetUploadTokenIdOk() (*string, bool)`

GetUploadTokenIdOk returns a tuple with the UploadTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTokenId

`func (o *KalturaUploadResponse) SetUploadTokenId(v string)`

SetUploadTokenId sets UploadTokenId field to given value.

### HasUploadTokenId

`func (o *KalturaUploadResponse) HasUploadTokenId() bool`

HasUploadTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


