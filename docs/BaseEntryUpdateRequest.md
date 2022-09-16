# BaseEntryUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseEntry** | [**KalturaBaseEntry**](KalturaBaseEntry.md) |  | 
**EntryId** | **string** |  | 

## Methods

### NewBaseEntryUpdateRequest

`func NewBaseEntryUpdateRequest(baseEntry KalturaBaseEntry, entryId string, ) *BaseEntryUpdateRequest`

NewBaseEntryUpdateRequest instantiates a new BaseEntryUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryUpdateRequestWithDefaults

`func NewBaseEntryUpdateRequestWithDefaults() *BaseEntryUpdateRequest`

NewBaseEntryUpdateRequestWithDefaults instantiates a new BaseEntryUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseEntry

`func (o *BaseEntryUpdateRequest) GetBaseEntry() KalturaBaseEntry`

GetBaseEntry returns the BaseEntry field if non-nil, zero value otherwise.

### GetBaseEntryOk

`func (o *BaseEntryUpdateRequest) GetBaseEntryOk() (*KalturaBaseEntry, bool)`

GetBaseEntryOk returns a tuple with the BaseEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseEntry

`func (o *BaseEntryUpdateRequest) SetBaseEntry(v KalturaBaseEntry)`

SetBaseEntry sets BaseEntry field to given value.


### GetEntryId

`func (o *BaseEntryUpdateRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BaseEntryUpdateRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BaseEntryUpdateRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


