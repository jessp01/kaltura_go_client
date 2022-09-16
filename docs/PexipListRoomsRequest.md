# PexipListRoomsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveOnly** | Pointer to **bool** |  | [optional] [default to false]
**Offset** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewPexipListRoomsRequest

`func NewPexipListRoomsRequest() *PexipListRoomsRequest`

NewPexipListRoomsRequest instantiates a new PexipListRoomsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPexipListRoomsRequestWithDefaults

`func NewPexipListRoomsRequestWithDefaults() *PexipListRoomsRequest`

NewPexipListRoomsRequestWithDefaults instantiates a new PexipListRoomsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveOnly

`func (o *PexipListRoomsRequest) GetActiveOnly() bool`

GetActiveOnly returns the ActiveOnly field if non-nil, zero value otherwise.

### GetActiveOnlyOk

`func (o *PexipListRoomsRequest) GetActiveOnlyOk() (*bool, bool)`

GetActiveOnlyOk returns a tuple with the ActiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOnly

`func (o *PexipListRoomsRequest) SetActiveOnly(v bool)`

SetActiveOnly sets ActiveOnly field to given value.

### HasActiveOnly

`func (o *PexipListRoomsRequest) HasActiveOnly() bool`

HasActiveOnly returns a boolean if a field has been set.

### GetOffset

`func (o *PexipListRoomsRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PexipListRoomsRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PexipListRoomsRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PexipListRoomsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetPageSize

`func (o *PexipListRoomsRequest) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PexipListRoomsRequest) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PexipListRoomsRequest) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PexipListRoomsRequest) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


