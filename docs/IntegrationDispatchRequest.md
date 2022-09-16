# IntegrationDispatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**KalturaIntegrationJobData**](KalturaIntegrationJobData.md) |  | 
**ObjectId** | **string** |  | 
**ObjectType** | **string** |  | 

## Methods

### NewIntegrationDispatchRequest

`func NewIntegrationDispatchRequest(data KalturaIntegrationJobData, objectId string, objectType string, ) *IntegrationDispatchRequest`

NewIntegrationDispatchRequest instantiates a new IntegrationDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationDispatchRequestWithDefaults

`func NewIntegrationDispatchRequestWithDefaults() *IntegrationDispatchRequest`

NewIntegrationDispatchRequestWithDefaults instantiates a new IntegrationDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *IntegrationDispatchRequest) GetData() KalturaIntegrationJobData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IntegrationDispatchRequest) GetDataOk() (*KalturaIntegrationJobData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IntegrationDispatchRequest) SetData(v KalturaIntegrationJobData)`

SetData sets Data field to given value.


### GetObjectId

`func (o *IntegrationDispatchRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *IntegrationDispatchRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *IntegrationDispatchRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.


### GetObjectType

`func (o *IntegrationDispatchRequest) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IntegrationDispatchRequest) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IntegrationDispatchRequest) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


