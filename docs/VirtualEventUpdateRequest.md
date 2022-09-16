# VirtualEventUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**VirtualEvent** | [**KalturaVirtualEvent**](KalturaVirtualEvent.md) |  | 

## Methods

### NewVirtualEventUpdateRequest

`func NewVirtualEventUpdateRequest(id int32, virtualEvent KalturaVirtualEvent, ) *VirtualEventUpdateRequest`

NewVirtualEventUpdateRequest instantiates a new VirtualEventUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualEventUpdateRequestWithDefaults

`func NewVirtualEventUpdateRequestWithDefaults() *VirtualEventUpdateRequest`

NewVirtualEventUpdateRequestWithDefaults instantiates a new VirtualEventUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualEventUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualEventUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualEventUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVirtualEvent

`func (o *VirtualEventUpdateRequest) GetVirtualEvent() KalturaVirtualEvent`

GetVirtualEvent returns the VirtualEvent field if non-nil, zero value otherwise.

### GetVirtualEventOk

`func (o *VirtualEventUpdateRequest) GetVirtualEventOk() (*KalturaVirtualEvent, bool)`

GetVirtualEventOk returns a tuple with the VirtualEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualEvent

`func (o *VirtualEventUpdateRequest) SetVirtualEvent(v KalturaVirtualEvent)`

SetVirtualEvent sets VirtualEvent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


