# PollVoteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerIds** | **string** |  | 
**PollId** | **string** |  | 
**UserId** | **string** |  | 

## Methods

### NewPollVoteRequest

`func NewPollVoteRequest(answerIds string, pollId string, userId string, ) *PollVoteRequest`

NewPollVoteRequest instantiates a new PollVoteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollVoteRequestWithDefaults

`func NewPollVoteRequestWithDefaults() *PollVoteRequest`

NewPollVoteRequestWithDefaults instantiates a new PollVoteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerIds

`func (o *PollVoteRequest) GetAnswerIds() string`

GetAnswerIds returns the AnswerIds field if non-nil, zero value otherwise.

### GetAnswerIdsOk

`func (o *PollVoteRequest) GetAnswerIdsOk() (*string, bool)`

GetAnswerIdsOk returns a tuple with the AnswerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerIds

`func (o *PollVoteRequest) SetAnswerIds(v string)`

SetAnswerIds sets AnswerIds field to given value.


### GetPollId

`func (o *PollVoteRequest) GetPollId() string`

GetPollId returns the PollId field if non-nil, zero value otherwise.

### GetPollIdOk

`func (o *PollVoteRequest) GetPollIdOk() (*string, bool)`

GetPollIdOk returns a tuple with the PollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollId

`func (o *PollVoteRequest) SetPollId(v string)`

SetPollId sets PollId field to given value.


### GetUserId

`func (o *PollVoteRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PollVoteRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PollVoteRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


