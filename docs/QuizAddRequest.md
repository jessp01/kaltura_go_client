# QuizAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**Quiz** | [**KalturaQuiz**](KalturaQuiz.md) |  | 

## Methods

### NewQuizAddRequest

`func NewQuizAddRequest(entryId string, quiz KalturaQuiz, ) *QuizAddRequest`

NewQuizAddRequest instantiates a new QuizAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuizAddRequestWithDefaults

`func NewQuizAddRequestWithDefaults() *QuizAddRequest`

NewQuizAddRequestWithDefaults instantiates a new QuizAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *QuizAddRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *QuizAddRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *QuizAddRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetQuiz

`func (o *QuizAddRequest) GetQuiz() KalturaQuiz`

GetQuiz returns the Quiz field if non-nil, zero value otherwise.

### GetQuizOk

`func (o *QuizAddRequest) GetQuizOk() (*KalturaQuiz, bool)`

GetQuizOk returns a tuple with the Quiz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuiz

`func (o *QuizAddRequest) SetQuiz(v KalturaQuiz)`

SetQuiz sets Quiz field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


