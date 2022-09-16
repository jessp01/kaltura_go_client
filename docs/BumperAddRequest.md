# BumperAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bumper** | [**KalturaBumper**](KalturaBumper.md) |  | 
**EntryId** | **string** |  | 

## Methods

### NewBumperAddRequest

`func NewBumperAddRequest(bumper KalturaBumper, entryId string, ) *BumperAddRequest`

NewBumperAddRequest instantiates a new BumperAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBumperAddRequestWithDefaults

`func NewBumperAddRequestWithDefaults() *BumperAddRequest`

NewBumperAddRequestWithDefaults instantiates a new BumperAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBumper

`func (o *BumperAddRequest) GetBumper() KalturaBumper`

GetBumper returns the Bumper field if non-nil, zero value otherwise.

### GetBumperOk

`func (o *BumperAddRequest) GetBumperOk() (*KalturaBumper, bool)`

GetBumperOk returns a tuple with the Bumper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBumper

`func (o *BumperAddRequest) SetBumper(v KalturaBumper)`

SetBumper sets Bumper field to given value.


### GetEntryId

`func (o *BumperAddRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *BumperAddRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *BumperAddRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


