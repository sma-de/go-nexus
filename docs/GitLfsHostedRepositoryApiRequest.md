# GitLfsHostedRepositoryApiRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique identifier for this repository | 
**Online** | **bool** | Whether this repository accepts incoming requests | 
**Storage** | [**HostedStorageAttributes**](HostedStorageAttributes.md) |  | 
**Cleanup** | Pointer to [**CleanupPolicyAttributes**](CleanupPolicyAttributes.md) |  | [optional] 
**Component** | Pointer to [**ComponentAttributes**](ComponentAttributes.md) |  | [optional] 

## Methods

### NewGitLfsHostedRepositoryApiRequest

`func NewGitLfsHostedRepositoryApiRequest(name string, online bool, storage HostedStorageAttributes, ) *GitLfsHostedRepositoryApiRequest`

NewGitLfsHostedRepositoryApiRequest instantiates a new GitLfsHostedRepositoryApiRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitLfsHostedRepositoryApiRequestWithDefaults

`func NewGitLfsHostedRepositoryApiRequestWithDefaults() *GitLfsHostedRepositoryApiRequest`

NewGitLfsHostedRepositoryApiRequestWithDefaults instantiates a new GitLfsHostedRepositoryApiRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitLfsHostedRepositoryApiRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitLfsHostedRepositoryApiRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitLfsHostedRepositoryApiRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOnline

`func (o *GitLfsHostedRepositoryApiRequest) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *GitLfsHostedRepositoryApiRequest) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *GitLfsHostedRepositoryApiRequest) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetStorage

`func (o *GitLfsHostedRepositoryApiRequest) GetStorage() HostedStorageAttributes`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *GitLfsHostedRepositoryApiRequest) GetStorageOk() (*HostedStorageAttributes, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *GitLfsHostedRepositoryApiRequest) SetStorage(v HostedStorageAttributes)`

SetStorage sets Storage field to given value.


### GetCleanup

`func (o *GitLfsHostedRepositoryApiRequest) GetCleanup() CleanupPolicyAttributes`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *GitLfsHostedRepositoryApiRequest) GetCleanupOk() (*CleanupPolicyAttributes, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *GitLfsHostedRepositoryApiRequest) SetCleanup(v CleanupPolicyAttributes)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *GitLfsHostedRepositoryApiRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetComponent

`func (o *GitLfsHostedRepositoryApiRequest) GetComponent() ComponentAttributes`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *GitLfsHostedRepositoryApiRequest) GetComponentOk() (*ComponentAttributes, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *GitLfsHostedRepositoryApiRequest) SetComponent(v ComponentAttributes)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *GitLfsHostedRepositoryApiRequest) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


