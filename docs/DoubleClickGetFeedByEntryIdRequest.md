# DoubleClickGetFeedByEntryIdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistributionProfileId** | **int32** |  | 
**EntryId** | **string** |  | 
**Hash** | **string** |  | 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewDoubleClickGetFeedByEntryIdRequest

`func NewDoubleClickGetFeedByEntryIdRequest(distributionProfileId int32, entryId string, hash string, ) *DoubleClickGetFeedByEntryIdRequest`

NewDoubleClickGetFeedByEntryIdRequest instantiates a new DoubleClickGetFeedByEntryIdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDoubleClickGetFeedByEntryIdRequestWithDefaults

`func NewDoubleClickGetFeedByEntryIdRequestWithDefaults() *DoubleClickGetFeedByEntryIdRequest`

NewDoubleClickGetFeedByEntryIdRequestWithDefaults instantiates a new DoubleClickGetFeedByEntryIdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistributionProfileId

`func (o *DoubleClickGetFeedByEntryIdRequest) GetDistributionProfileId() int32`

GetDistributionProfileId returns the DistributionProfileId field if non-nil, zero value otherwise.

### GetDistributionProfileIdOk

`func (o *DoubleClickGetFeedByEntryIdRequest) GetDistributionProfileIdOk() (*int32, bool)`

GetDistributionProfileIdOk returns a tuple with the DistributionProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributionProfileId

`func (o *DoubleClickGetFeedByEntryIdRequest) SetDistributionProfileId(v int32)`

SetDistributionProfileId sets DistributionProfileId field to given value.


### GetEntryId

`func (o *DoubleClickGetFeedByEntryIdRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *DoubleClickGetFeedByEntryIdRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *DoubleClickGetFeedByEntryIdRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetHash

`func (o *DoubleClickGetFeedByEntryIdRequest) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *DoubleClickGetFeedByEntryIdRequest) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *DoubleClickGetFeedByEntryIdRequest) SetHash(v string)`

SetHash sets Hash field to given value.


### GetVersion

`func (o *DoubleClickGetFeedByEntryIdRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DoubleClickGetFeedByEntryIdRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DoubleClickGetFeedByEntryIdRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DoubleClickGetFeedByEntryIdRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


