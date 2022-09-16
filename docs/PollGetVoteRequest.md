# PollGetVoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PollId** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewPollGetVoteRequest

`func NewPollGetVoteRequest(pollId string, userId string, ) *PollGetVoteRequest`

NewPollGetVoteRequest instantiates a new PollGetVoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollGetVoteRequestWithDefaults

`func NewPollGetVoteRequestWithDefaults() *PollGetVoteRequest`

NewPollGetVoteRequestWithDefaults instantiates a new PollGetVoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPollId

`func (o *PollGetVoteRequest) GetPollId() string`

GetPollId returns the PollId field if non-nil, zero value otherwise.

### GetPollIdOk

`func (o *PollGetVoteRequest) GetPollIdOk() (*string, bool)`

GetPollIdOk returns a tuple with the PollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollId

`func (o *PollGetVoteRequest) SetPollId(v string)`

SetPollId sets PollId field to given value.


### GetUserId

`func (o *PollGetVoteRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PollGetVoteRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PollGetVoteRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


