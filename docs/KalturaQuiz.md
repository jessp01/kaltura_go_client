# KalturaQuiz

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAnswerUpdate** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**AllowDownload** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**AttemptsAllowed** | Pointer to **int32** |  | [optional] 
**ScoreType** | Pointer to **int32** | Enum Type: &#x60;KalturaScoreType&#x60; | [optional] 
**ShowCorrectAfterSubmission** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**ShowCorrectKeyOnAnswer** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**ShowGradeAfterSubmission** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**ShowResultOnAnswer** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60; | [optional] 
**UiAttributes** | Pointer to [**[]KalturaKeyValue**](KalturaKeyValue.md) |  | [optional] 
**Version** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaQuiz

`func NewKalturaQuiz() *KalturaQuiz`

NewKalturaQuiz instantiates a new KalturaQuiz object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaQuizWithDefaults

`func NewKalturaQuizWithDefaults() *KalturaQuiz`

NewKalturaQuizWithDefaults instantiates a new KalturaQuiz object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAnswerUpdate

`func (o *KalturaQuiz) GetAllowAnswerUpdate() int32`

GetAllowAnswerUpdate returns the AllowAnswerUpdate field if non-nil, zero value otherwise.

### GetAllowAnswerUpdateOk

`func (o *KalturaQuiz) GetAllowAnswerUpdateOk() (*int32, bool)`

GetAllowAnswerUpdateOk returns a tuple with the AllowAnswerUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAnswerUpdate

`func (o *KalturaQuiz) SetAllowAnswerUpdate(v int32)`

SetAllowAnswerUpdate sets AllowAnswerUpdate field to given value.

### HasAllowAnswerUpdate

`func (o *KalturaQuiz) HasAllowAnswerUpdate() bool`

HasAllowAnswerUpdate returns a boolean if a field has been set.

### GetAllowDownload

`func (o *KalturaQuiz) GetAllowDownload() int32`

GetAllowDownload returns the AllowDownload field if non-nil, zero value otherwise.

### GetAllowDownloadOk

`func (o *KalturaQuiz) GetAllowDownloadOk() (*int32, bool)`

GetAllowDownloadOk returns a tuple with the AllowDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDownload

`func (o *KalturaQuiz) SetAllowDownload(v int32)`

SetAllowDownload sets AllowDownload field to given value.

### HasAllowDownload

`func (o *KalturaQuiz) HasAllowDownload() bool`

HasAllowDownload returns a boolean if a field has been set.

### GetAttemptsAllowed

`func (o *KalturaQuiz) GetAttemptsAllowed() int32`

GetAttemptsAllowed returns the AttemptsAllowed field if non-nil, zero value otherwise.

### GetAttemptsAllowedOk

`func (o *KalturaQuiz) GetAttemptsAllowedOk() (*int32, bool)`

GetAttemptsAllowedOk returns a tuple with the AttemptsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptsAllowed

`func (o *KalturaQuiz) SetAttemptsAllowed(v int32)`

SetAttemptsAllowed sets AttemptsAllowed field to given value.

### HasAttemptsAllowed

`func (o *KalturaQuiz) HasAttemptsAllowed() bool`

HasAttemptsAllowed returns a boolean if a field has been set.

### GetScoreType

`func (o *KalturaQuiz) GetScoreType() int32`

GetScoreType returns the ScoreType field if non-nil, zero value otherwise.

### GetScoreTypeOk

`func (o *KalturaQuiz) GetScoreTypeOk() (*int32, bool)`

GetScoreTypeOk returns a tuple with the ScoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreType

`func (o *KalturaQuiz) SetScoreType(v int32)`

SetScoreType sets ScoreType field to given value.

### HasScoreType

`func (o *KalturaQuiz) HasScoreType() bool`

HasScoreType returns a boolean if a field has been set.

### GetShowCorrectAfterSubmission

`func (o *KalturaQuiz) GetShowCorrectAfterSubmission() int32`

GetShowCorrectAfterSubmission returns the ShowCorrectAfterSubmission field if non-nil, zero value otherwise.

### GetShowCorrectAfterSubmissionOk

`func (o *KalturaQuiz) GetShowCorrectAfterSubmissionOk() (*int32, bool)`

GetShowCorrectAfterSubmissionOk returns a tuple with the ShowCorrectAfterSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCorrectAfterSubmission

`func (o *KalturaQuiz) SetShowCorrectAfterSubmission(v int32)`

SetShowCorrectAfterSubmission sets ShowCorrectAfterSubmission field to given value.

### HasShowCorrectAfterSubmission

`func (o *KalturaQuiz) HasShowCorrectAfterSubmission() bool`

HasShowCorrectAfterSubmission returns a boolean if a field has been set.

### GetShowCorrectKeyOnAnswer

`func (o *KalturaQuiz) GetShowCorrectKeyOnAnswer() int32`

GetShowCorrectKeyOnAnswer returns the ShowCorrectKeyOnAnswer field if non-nil, zero value otherwise.

### GetShowCorrectKeyOnAnswerOk

`func (o *KalturaQuiz) GetShowCorrectKeyOnAnswerOk() (*int32, bool)`

GetShowCorrectKeyOnAnswerOk returns a tuple with the ShowCorrectKeyOnAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCorrectKeyOnAnswer

`func (o *KalturaQuiz) SetShowCorrectKeyOnAnswer(v int32)`

SetShowCorrectKeyOnAnswer sets ShowCorrectKeyOnAnswer field to given value.

### HasShowCorrectKeyOnAnswer

`func (o *KalturaQuiz) HasShowCorrectKeyOnAnswer() bool`

HasShowCorrectKeyOnAnswer returns a boolean if a field has been set.

### GetShowGradeAfterSubmission

`func (o *KalturaQuiz) GetShowGradeAfterSubmission() int32`

GetShowGradeAfterSubmission returns the ShowGradeAfterSubmission field if non-nil, zero value otherwise.

### GetShowGradeAfterSubmissionOk

`func (o *KalturaQuiz) GetShowGradeAfterSubmissionOk() (*int32, bool)`

GetShowGradeAfterSubmissionOk returns a tuple with the ShowGradeAfterSubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowGradeAfterSubmission

`func (o *KalturaQuiz) SetShowGradeAfterSubmission(v int32)`

SetShowGradeAfterSubmission sets ShowGradeAfterSubmission field to given value.

### HasShowGradeAfterSubmission

`func (o *KalturaQuiz) HasShowGradeAfterSubmission() bool`

HasShowGradeAfterSubmission returns a boolean if a field has been set.

### GetShowResultOnAnswer

`func (o *KalturaQuiz) GetShowResultOnAnswer() int32`

GetShowResultOnAnswer returns the ShowResultOnAnswer field if non-nil, zero value otherwise.

### GetShowResultOnAnswerOk

`func (o *KalturaQuiz) GetShowResultOnAnswerOk() (*int32, bool)`

GetShowResultOnAnswerOk returns a tuple with the ShowResultOnAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowResultOnAnswer

`func (o *KalturaQuiz) SetShowResultOnAnswer(v int32)`

SetShowResultOnAnswer sets ShowResultOnAnswer field to given value.

### HasShowResultOnAnswer

`func (o *KalturaQuiz) HasShowResultOnAnswer() bool`

HasShowResultOnAnswer returns a boolean if a field has been set.

### GetUiAttributes

`func (o *KalturaQuiz) GetUiAttributes() []KalturaKeyValue`

GetUiAttributes returns the UiAttributes field if non-nil, zero value otherwise.

### GetUiAttributesOk

`func (o *KalturaQuiz) GetUiAttributesOk() (*[]KalturaKeyValue, bool)`

GetUiAttributesOk returns a tuple with the UiAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiAttributes

`func (o *KalturaQuiz) SetUiAttributes(v []KalturaKeyValue)`

SetUiAttributes sets UiAttributes field to given value.

### HasUiAttributes

`func (o *KalturaQuiz) HasUiAttributes() bool`

HasUiAttributes returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaQuiz) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaQuiz) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaQuiz) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaQuiz) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


