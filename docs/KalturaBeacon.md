# KalturaBeacon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60;  Beacon id | [optional] [readonly] 
**IndexType** | Pointer to **string** | &#x60;readOnly&#x60;  Beacon indexType | [optional] [readonly] 
**ObjectId** | Pointer to **string** |  | [optional] 
**PrivateData** | Pointer to **string** |  | [optional] 
**RawData** | Pointer to **string** |  | [optional] 
**RelatedObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaBeaconObjectTypes&#x60;  The object which this beacon belongs to | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Beacon update date as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaBeacon

`func NewKalturaBeacon() *KalturaBeacon`

NewKalturaBeacon instantiates a new KalturaBeacon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBeaconWithDefaults

`func NewKalturaBeaconWithDefaults() *KalturaBeacon`

NewKalturaBeaconWithDefaults instantiates a new KalturaBeacon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *KalturaBeacon) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *KalturaBeacon) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *KalturaBeacon) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *KalturaBeacon) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *KalturaBeacon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaBeacon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaBeacon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaBeacon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexType

`func (o *KalturaBeacon) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *KalturaBeacon) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *KalturaBeacon) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *KalturaBeacon) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetObjectId

`func (o *KalturaBeacon) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *KalturaBeacon) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *KalturaBeacon) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *KalturaBeacon) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetPrivateData

`func (o *KalturaBeacon) GetPrivateData() string`

GetPrivateData returns the PrivateData field if non-nil, zero value otherwise.

### GetPrivateDataOk

`func (o *KalturaBeacon) GetPrivateDataOk() (*string, bool)`

GetPrivateDataOk returns a tuple with the PrivateData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateData

`func (o *KalturaBeacon) SetPrivateData(v string)`

SetPrivateData sets PrivateData field to given value.

### HasPrivateData

`func (o *KalturaBeacon) HasPrivateData() bool`

HasPrivateData returns a boolean if a field has been set.

### GetRawData

`func (o *KalturaBeacon) GetRawData() string`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *KalturaBeacon) GetRawDataOk() (*string, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *KalturaBeacon) SetRawData(v string)`

SetRawData sets RawData field to given value.

### HasRawData

`func (o *KalturaBeacon) HasRawData() bool`

HasRawData returns a boolean if a field has been set.

### GetRelatedObjectType

`func (o *KalturaBeacon) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *KalturaBeacon) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *KalturaBeacon) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *KalturaBeacon) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaBeacon) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaBeacon) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaBeacon) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaBeacon) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


