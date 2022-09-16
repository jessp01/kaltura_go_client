# BeaconAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Beacon** | [**KalturaBeacon**](KalturaBeacon.md) |  | 
**ShouldLog** | Pointer to **int32** |  | [optional] 

## Methods

### NewBeaconAddRequest

`func NewBeaconAddRequest(beacon KalturaBeacon, ) *BeaconAddRequest`

NewBeaconAddRequest instantiates a new BeaconAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeaconAddRequestWithDefaults

`func NewBeaconAddRequestWithDefaults() *BeaconAddRequest`

NewBeaconAddRequestWithDefaults instantiates a new BeaconAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeacon

`func (o *BeaconAddRequest) GetBeacon() KalturaBeacon`

GetBeacon returns the Beacon field if non-nil, zero value otherwise.

### GetBeaconOk

`func (o *BeaconAddRequest) GetBeaconOk() (*KalturaBeacon, bool)`

GetBeaconOk returns a tuple with the Beacon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeacon

`func (o *BeaconAddRequest) SetBeacon(v KalturaBeacon)`

SetBeacon sets Beacon field to given value.


### GetShouldLog

`func (o *BeaconAddRequest) GetShouldLog() int32`

GetShouldLog returns the ShouldLog field if non-nil, zero value otherwise.

### GetShouldLogOk

`func (o *BeaconAddRequest) GetShouldLogOk() (*int32, bool)`

GetShouldLogOk returns a tuple with the ShouldLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldLog

`func (o *BeaconAddRequest) SetShouldLog(v int32)`

SetShouldLog sets ShouldLog field to given value.

### HasShouldLog

`func (o *BeaconAddRequest) HasShouldLog() bool`

HasShouldLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


