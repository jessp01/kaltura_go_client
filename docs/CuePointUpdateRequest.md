# CuePointUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CuePoint** | [**KalturaCuePoint**](KalturaCuePoint.md) |  | 
**Id** | **string** |  | 

## Methods

### NewCuePointUpdateRequest

`func NewCuePointUpdateRequest(cuePoint KalturaCuePoint, id string, ) *CuePointUpdateRequest`

NewCuePointUpdateRequest instantiates a new CuePointUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCuePointUpdateRequestWithDefaults

`func NewCuePointUpdateRequestWithDefaults() *CuePointUpdateRequest`

NewCuePointUpdateRequestWithDefaults instantiates a new CuePointUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCuePoint

`func (o *CuePointUpdateRequest) GetCuePoint() KalturaCuePoint`

GetCuePoint returns the CuePoint field if non-nil, zero value otherwise.

### GetCuePointOk

`func (o *CuePointUpdateRequest) GetCuePointOk() (*KalturaCuePoint, bool)`

GetCuePointOk returns a tuple with the CuePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuePoint

`func (o *CuePointUpdateRequest) SetCuePoint(v KalturaCuePoint)`

SetCuePoint sets CuePoint field to given value.


### GetId

`func (o *CuePointUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CuePointUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CuePointUpdateRequest) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


