# UserScoreUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GameObjectId** | **string** |  | 
**GameObjectType** | **string** |  | 
**Score** | **int32** |  | 
**UserId** | **string** |  | 

## Methods

### NewUserScoreUpdateRequest

`func NewUserScoreUpdateRequest(gameObjectId string, gameObjectType string, score int32, userId string, ) *UserScoreUpdateRequest`

NewUserScoreUpdateRequest instantiates a new UserScoreUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserScoreUpdateRequestWithDefaults

`func NewUserScoreUpdateRequestWithDefaults() *UserScoreUpdateRequest`

NewUserScoreUpdateRequestWithDefaults instantiates a new UserScoreUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGameObjectId

`func (o *UserScoreUpdateRequest) GetGameObjectId() string`

GetGameObjectId returns the GameObjectId field if non-nil, zero value otherwise.

### GetGameObjectIdOk

`func (o *UserScoreUpdateRequest) GetGameObjectIdOk() (*string, bool)`

GetGameObjectIdOk returns a tuple with the GameObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameObjectId

`func (o *UserScoreUpdateRequest) SetGameObjectId(v string)`

SetGameObjectId sets GameObjectId field to given value.


### GetGameObjectType

`func (o *UserScoreUpdateRequest) GetGameObjectType() string`

GetGameObjectType returns the GameObjectType field if non-nil, zero value otherwise.

### GetGameObjectTypeOk

`func (o *UserScoreUpdateRequest) GetGameObjectTypeOk() (*string, bool)`

GetGameObjectTypeOk returns a tuple with the GameObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameObjectType

`func (o *UserScoreUpdateRequest) SetGameObjectType(v string)`

SetGameObjectType sets GameObjectType field to given value.


### GetScore

`func (o *UserScoreUpdateRequest) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *UserScoreUpdateRequest) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *UserScoreUpdateRequest) SetScore(v int32)`

SetScore sets Score field to given value.


### GetUserId

`func (o *UserScoreUpdateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserScoreUpdateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserScoreUpdateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


