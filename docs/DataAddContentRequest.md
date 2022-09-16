# DataAddContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Resource** | [**KalturaGenericDataCenterContentResource**](KalturaGenericDataCenterContentResource.md) |  | 

## Methods

### NewDataAddContentRequest

`func NewDataAddContentRequest(entryId string, resource KalturaGenericDataCenterContentResource, ) *DataAddContentRequest`

NewDataAddContentRequest instantiates a new DataAddContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataAddContentRequestWithDefaults

`func NewDataAddContentRequestWithDefaults() *DataAddContentRequest`

NewDataAddContentRequestWithDefaults instantiates a new DataAddContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *DataAddContentRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *DataAddContentRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *DataAddContentRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetResource

`func (o *DataAddContentRequest) GetResource() KalturaGenericDataCenterContentResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DataAddContentRequest) GetResourceOk() (*KalturaGenericDataCenterContentResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DataAddContentRequest) SetResource(v KalturaGenericDataCenterContentResource)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


