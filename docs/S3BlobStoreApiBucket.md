# S3BlobStoreApiBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | The AWS region to create a new S3 bucket in or an existing S3 bucket&#39;s region | [readonly] 
**Name** | **string** | The name of the S3 bucket | [readonly] 
**Prefix** | Pointer to **string** | The S3 blob store (i.e S3 object) key prefix | [optional] [readonly] 
**Expiration** | **int32** | How many days until deleted blobs are finally removed from the S3 bucket (-1 to disable) | [readonly] 

## Methods

### NewS3BlobStoreApiBucket

`func NewS3BlobStoreApiBucket(region string, name string, expiration int32, ) *S3BlobStoreApiBucket`

NewS3BlobStoreApiBucket instantiates a new S3BlobStoreApiBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3BlobStoreApiBucketWithDefaults

`func NewS3BlobStoreApiBucketWithDefaults() *S3BlobStoreApiBucket`

NewS3BlobStoreApiBucketWithDefaults instantiates a new S3BlobStoreApiBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *S3BlobStoreApiBucket) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *S3BlobStoreApiBucket) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *S3BlobStoreApiBucket) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetName

`func (o *S3BlobStoreApiBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *S3BlobStoreApiBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *S3BlobStoreApiBucket) SetName(v string)`

SetName sets Name field to given value.


### GetPrefix

`func (o *S3BlobStoreApiBucket) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *S3BlobStoreApiBucket) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *S3BlobStoreApiBucket) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *S3BlobStoreApiBucket) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetExpiration

`func (o *S3BlobStoreApiBucket) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *S3BlobStoreApiBucket) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *S3BlobStoreApiBucket) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


