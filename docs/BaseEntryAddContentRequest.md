# BaseEntryAddContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Resource** | [**KalturaResource**](KalturaResource.md) |  | 

## Methods

### NewBaseEntryAddContentRequest

`func NewBaseEntryAddContentRequest(entryId string, resource KalturaResource, ) *BaseEntryAddContentRequest`

NewBaseEntryAddContentRequest instantiates a new BaseEntryAddContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryAddContentRequestWithDefaults

`func NewBaseEntryAddContentRequestWithDefaults() *BaseEntryAddContentRequest`

NewBaseEntryAddContentRequestWithDefaults instantiates a new BaseEntryAddContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *BaseEntryAddContentRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BaseEntryAddContentRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BaseEntryAddContentRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetResource

`func (o *BaseEntryAddContentRequest) GetResource() KalturaResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BaseEntryAddContentRequest) GetResourceOk() (*KalturaResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BaseEntryAddContentRequest) SetResource(v KalturaResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


