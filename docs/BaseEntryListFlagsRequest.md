# BaseEntryListFlagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 

## Methods

### NewBaseEntryListFlagsRequest

`func NewBaseEntryListFlagsRequest(entryId string, ) *BaseEntryListFlagsRequest`

NewBaseEntryListFlagsRequest instantiates a new BaseEntryListFlagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEntryListFlagsRequestWithDefaults

`func NewBaseEntryListFlagsRequestWithDefaults() *BaseEntryListFlagsRequest`

NewBaseEntryListFlagsRequestWithDefaults instantiates a new BaseEntryListFlagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *BaseEntryListFlagsRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BaseEntryListFlagsRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BaseEntryListFlagsRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetPager

`func (o *BaseEntryListFlagsRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *BaseEntryListFlagsRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *BaseEntryListFlagsRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *BaseEntryListFlagsRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


