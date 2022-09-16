# PollAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PollType** | Pointer to **string** |  | [optional] [default to "SINGLE_ANONYMOUS"]

## Methods

### NewPollAddRequest

`func NewPollAddRequest() *PollAddRequest`

NewPollAddRequest instantiates a new PollAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollAddRequestWithDefaults

`func NewPollAddRequestWithDefaults() *PollAddRequest`

NewPollAddRequestWithDefaults instantiates a new PollAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPollType

`func (o *PollAddRequest) GetPollType() string`

GetPollType returns the PollType field if non-nil, zero value otherwise.

### GetPollTypeOk

`func (o *PollAddRequest) GetPollTypeOk() (*string, bool)`

GetPollTypeOk returns a tuple with the PollType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollType

`func (o *PollAddRequest) SetPollType(v string)`

SetPollType sets PollType field to given value.

### HasPollType

`func (o *PollAddRequest) HasPollType() bool`

HasPollType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


