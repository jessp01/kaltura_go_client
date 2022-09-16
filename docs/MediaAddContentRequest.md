# MediaAddContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Resource** | Pointer to [**KalturaResource**](KalturaResource.md) |  | [optional] 

## Methods

### NewMediaAddContentRequest

`func NewMediaAddContentRequest(entryId string, ) *MediaAddContentRequest`

NewMediaAddContentRequest instantiates a new MediaAddContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaAddContentRequestWithDefaults

`func NewMediaAddContentRequestWithDefaults() *MediaAddContentRequest`

NewMediaAddContentRequestWithDefaults instantiates a new MediaAddContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *MediaAddContentRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *MediaAddContentRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *MediaAddContentRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetResource

`func (o *MediaAddContentRequest) GetResource() KalturaResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MediaAddContentRequest) GetResourceOk() (*KalturaResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MediaAddContentRequest) SetResource(v KalturaResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MediaAddContentRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


