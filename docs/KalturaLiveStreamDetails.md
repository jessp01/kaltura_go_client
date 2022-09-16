# KalturaLiveStreamDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BroadcastStatus** | Pointer to **int32** | Enum Type: &#x60;KalturaLiveStreamBroadcastStatus&#x60; | [optional] 
**PrimaryStreamStatus** | Pointer to **int32** | Enum Type: &#x60;KalturaEntryServerNodeStatus&#x60;  The status of the primary stream | [optional] 
**SecondaryStreamStatus** | Pointer to **int32** | Enum Type: &#x60;KalturaEntryServerNodeStatus&#x60;  The status of the secondary stream | [optional] 
**ViewMode** | Pointer to **int32** | Enum Type: &#x60;KalturaViewMode&#x60; | [optional] 
**WasBroadcast** | Pointer to **bool** |  | [optional] 

## Methods

### NewKalturaLiveStreamDetails

`func NewKalturaLiveStreamDetails() *KalturaLiveStreamDetails`

NewKalturaLiveStreamDetails instantiates a new KalturaLiveStreamDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaLiveStreamDetailsWithDefaults

`func NewKalturaLiveStreamDetailsWithDefaults() *KalturaLiveStreamDetails`

NewKalturaLiveStreamDetailsWithDefaults instantiates a new KalturaLiveStreamDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroadcastStatus

`func (o *KalturaLiveStreamDetails) GetBroadcastStatus() int32`

GetBroadcastStatus returns the BroadcastStatus field if non-nil, zero value otherwise.

### GetBroadcastStatusOk

`func (o *KalturaLiveStreamDetails) GetBroadcastStatusOk() (*int32, bool)`

GetBroadcastStatusOk returns a tuple with the BroadcastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastStatus

`func (o *KalturaLiveStreamDetails) SetBroadcastStatus(v int32)`

SetBroadcastStatus sets BroadcastStatus field to given value.

### HasBroadcastStatus

`func (o *KalturaLiveStreamDetails) HasBroadcastStatus() bool`

HasBroadcastStatus returns a boolean if a field has been set.

### GetPrimaryStreamStatus

`func (o *KalturaLiveStreamDetails) GetPrimaryStreamStatus() int32`

GetPrimaryStreamStatus returns the PrimaryStreamStatus field if non-nil, zero value otherwise.

### GetPrimaryStreamStatusOk

`func (o *KalturaLiveStreamDetails) GetPrimaryStreamStatusOk() (*int32, bool)`

GetPrimaryStreamStatusOk returns a tuple with the PrimaryStreamStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryStreamStatus

`func (o *KalturaLiveStreamDetails) SetPrimaryStreamStatus(v int32)`

SetPrimaryStreamStatus sets PrimaryStreamStatus field to given value.

### HasPrimaryStreamStatus

`func (o *KalturaLiveStreamDetails) HasPrimaryStreamStatus() bool`

HasPrimaryStreamStatus returns a boolean if a field has been set.

### GetSecondaryStreamStatus

`func (o *KalturaLiveStreamDetails) GetSecondaryStreamStatus() int32`

GetSecondaryStreamStatus returns the SecondaryStreamStatus field if non-nil, zero value otherwise.

### GetSecondaryStreamStatusOk

`func (o *KalturaLiveStreamDetails) GetSecondaryStreamStatusOk() (*int32, bool)`

GetSecondaryStreamStatusOk returns a tuple with the SecondaryStreamStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryStreamStatus

`func (o *KalturaLiveStreamDetails) SetSecondaryStreamStatus(v int32)`

SetSecondaryStreamStatus sets SecondaryStreamStatus field to given value.

### HasSecondaryStreamStatus

`func (o *KalturaLiveStreamDetails) HasSecondaryStreamStatus() bool`

HasSecondaryStreamStatus returns a boolean if a field has been set.

### GetViewMode

`func (o *KalturaLiveStreamDetails) GetViewMode() int32`

GetViewMode returns the ViewMode field if non-nil, zero value otherwise.

### GetViewModeOk

`func (o *KalturaLiveStreamDetails) GetViewModeOk() (*int32, bool)`

GetViewModeOk returns a tuple with the ViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMode

`func (o *KalturaLiveStreamDetails) SetViewMode(v int32)`

SetViewMode sets ViewMode field to given value.

### HasViewMode

`func (o *KalturaLiveStreamDetails) HasViewMode() bool`

HasViewMode returns a boolean if a field has been set.

### GetWasBroadcast

`func (o *KalturaLiveStreamDetails) GetWasBroadcast() bool`

GetWasBroadcast returns the WasBroadcast field if non-nil, zero value otherwise.

### GetWasBroadcastOk

`func (o *KalturaLiveStreamDetails) GetWasBroadcastOk() (*bool, bool)`

GetWasBroadcastOk returns a tuple with the WasBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasBroadcast

`func (o *KalturaLiveStreamDetails) SetWasBroadcast(v bool)`

SetWasBroadcast sets WasBroadcast field to given value.

### HasWasBroadcast

`func (o *KalturaLiveStreamDetails) HasWasBroadcast() bool`

HasWasBroadcast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


