# PollGetVotesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerIds** | **string** |  | 
**PollId** | **string** |  | 

## Methods

### NewPollGetVotesRequest

`func NewPollGetVotesRequest(answerIds string, pollId string, ) *PollGetVotesRequest`

NewPollGetVotesRequest instantiates a new PollGetVotesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPollGetVotesRequestWithDefaults

`func NewPollGetVotesRequestWithDefaults() *PollGetVotesRequest`

NewPollGetVotesRequestWithDefaults instantiates a new PollGetVotesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerIds

`func (o *PollGetVotesRequest) GetAnswerIds() string`

GetAnswerIds returns the AnswerIds field if non-nil, zero value otherwise.

### GetAnswerIdsOk

`func (o *PollGetVotesRequest) GetAnswerIdsOk() (*string, bool)`

GetAnswerIdsOk returns a tuple with the AnswerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerIds

`func (o *PollGetVotesRequest) SetAnswerIds(v string)`

SetAnswerIds sets AnswerIds field to given value.


### GetPollId

`func (o *PollGetVotesRequest) GetPollId() string`

GetPollId returns the PollId field if non-nil, zero value otherwise.

### GetPollIdOk

`func (o *PollGetVotesRequest) GetPollIdOk() (*string, bool)`

GetPollIdOk returns a tuple with the PollId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollId

`func (o *PollGetVotesRequest) SetPollId(v string)`

SetPollId sets PollId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


