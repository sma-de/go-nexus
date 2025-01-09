/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.76.0-03
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"time"
)

// checks if the TaskXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskXO{}

// TaskXO struct for TaskXO
type TaskXO struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Message *string `json:"message,omitempty"`
	CurrentState *string `json:"currentState,omitempty"`
	LastRunResult *string `json:"lastRunResult,omitempty"`
	NextRun *time.Time `json:"nextRun,omitempty"`
	LastRun *time.Time `json:"lastRun,omitempty"`
}

// NewTaskXO instantiates a new TaskXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskXO() *TaskXO {
	this := TaskXO{}
	return &this
}

// NewTaskXOWithDefaults instantiates a new TaskXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskXOWithDefaults() *TaskXO {
	this := TaskXO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskXO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskXO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaskXO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaskXO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaskXO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaskXO) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TaskXO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TaskXO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TaskXO) SetType(v string) {
	o.Type = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TaskXO) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TaskXO) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TaskXO) SetMessage(v string) {
	o.Message = &v
}

// GetCurrentState returns the CurrentState field value if set, zero value otherwise.
func (o *TaskXO) GetCurrentState() string {
	if o == nil || IsNil(o.CurrentState) {
		var ret string
		return ret
	}
	return *o.CurrentState
}

// GetCurrentStateOk returns a tuple with the CurrentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetCurrentStateOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentState) {
		return nil, false
	}
	return o.CurrentState, true
}

// HasCurrentState returns a boolean if a field has been set.
func (o *TaskXO) HasCurrentState() bool {
	if o != nil && !IsNil(o.CurrentState) {
		return true
	}

	return false
}

// SetCurrentState gets a reference to the given string and assigns it to the CurrentState field.
func (o *TaskXO) SetCurrentState(v string) {
	o.CurrentState = &v
}

// GetLastRunResult returns the LastRunResult field value if set, zero value otherwise.
func (o *TaskXO) GetLastRunResult() string {
	if o == nil || IsNil(o.LastRunResult) {
		var ret string
		return ret
	}
	return *o.LastRunResult
}

// GetLastRunResultOk returns a tuple with the LastRunResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetLastRunResultOk() (*string, bool) {
	if o == nil || IsNil(o.LastRunResult) {
		return nil, false
	}
	return o.LastRunResult, true
}

// HasLastRunResult returns a boolean if a field has been set.
func (o *TaskXO) HasLastRunResult() bool {
	if o != nil && !IsNil(o.LastRunResult) {
		return true
	}

	return false
}

// SetLastRunResult gets a reference to the given string and assigns it to the LastRunResult field.
func (o *TaskXO) SetLastRunResult(v string) {
	o.LastRunResult = &v
}

// GetNextRun returns the NextRun field value if set, zero value otherwise.
func (o *TaskXO) GetNextRun() time.Time {
	if o == nil || IsNil(o.NextRun) {
		var ret time.Time
		return ret
	}
	return *o.NextRun
}

// GetNextRunOk returns a tuple with the NextRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetNextRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.NextRun) {
		return nil, false
	}
	return o.NextRun, true
}

// HasNextRun returns a boolean if a field has been set.
func (o *TaskXO) HasNextRun() bool {
	if o != nil && !IsNil(o.NextRun) {
		return true
	}

	return false
}

// SetNextRun gets a reference to the given time.Time and assigns it to the NextRun field.
func (o *TaskXO) SetNextRun(v time.Time) {
	o.NextRun = &v
}

// GetLastRun returns the LastRun field value if set, zero value otherwise.
func (o *TaskXO) GetLastRun() time.Time {
	if o == nil || IsNil(o.LastRun) {
		var ret time.Time
		return ret
	}
	return *o.LastRun
}

// GetLastRunOk returns a tuple with the LastRun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskXO) GetLastRunOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastRun) {
		return nil, false
	}
	return o.LastRun, true
}

// HasLastRun returns a boolean if a field has been set.
func (o *TaskXO) HasLastRun() bool {
	if o != nil && !IsNil(o.LastRun) {
		return true
	}

	return false
}

// SetLastRun gets a reference to the given time.Time and assigns it to the LastRun field.
func (o *TaskXO) SetLastRun(v time.Time) {
	o.LastRun = &v
}

func (o TaskXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.CurrentState) {
		toSerialize["currentState"] = o.CurrentState
	}
	if !IsNil(o.LastRunResult) {
		toSerialize["lastRunResult"] = o.LastRunResult
	}
	if !IsNil(o.NextRun) {
		toSerialize["nextRun"] = o.NextRun
	}
	if !IsNil(o.LastRun) {
		toSerialize["lastRun"] = o.LastRun
	}
	return toSerialize, nil
}

type NullableTaskXO struct {
	value *TaskXO
	isSet bool
}

func (v NullableTaskXO) Get() *TaskXO {
	return v.value
}

func (v *NullableTaskXO) Set(val *TaskXO) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskXO) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskXO(val *TaskXO) *NullableTaskXO {
	return &NullableTaskXO{value: val, isSet: true}
}

func (v NullableTaskXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


