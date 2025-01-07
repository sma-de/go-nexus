/*
Nexus Repository Manager REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.75.1-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nexus

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TagXO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagXO{}

// TagXO struct for TagXO
type TagXO struct {
	Name string `json:"name" validate:"regexp=^[a-zA-Z0-9\\\\-]{1}[a-zA-Z0-9_\\\\-\\\\.]*$"`
	Attributes map[string]map[string]interface{} `json:"attributes,omitempty"`
	FirstCreated *time.Time `json:"firstCreated,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
}

type _TagXO TagXO

// NewTagXO instantiates a new TagXO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagXO(name string) *TagXO {
	this := TagXO{}
	this.Name = name
	return &this
}

// NewTagXOWithDefaults instantiates a new TagXO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagXOWithDefaults() *TagXO {
	this := TagXO{}
	return &this
}

// GetName returns the Name field value
func (o *TagXO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TagXO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TagXO) SetName(v string) {
	o.Name = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TagXO) GetAttributes() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagXO) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attributes) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TagXO) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *TagXO) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = v
}

// GetFirstCreated returns the FirstCreated field value if set, zero value otherwise.
func (o *TagXO) GetFirstCreated() time.Time {
	if o == nil || IsNil(o.FirstCreated) {
		var ret time.Time
		return ret
	}
	return *o.FirstCreated
}

// GetFirstCreatedOk returns a tuple with the FirstCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagXO) GetFirstCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FirstCreated) {
		return nil, false
	}
	return o.FirstCreated, true
}

// HasFirstCreated returns a boolean if a field has been set.
func (o *TagXO) HasFirstCreated() bool {
	if o != nil && !IsNil(o.FirstCreated) {
		return true
	}

	return false
}

// SetFirstCreated gets a reference to the given time.Time and assigns it to the FirstCreated field.
func (o *TagXO) SetFirstCreated(v time.Time) {
	o.FirstCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *TagXO) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagXO) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *TagXO) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *TagXO) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o TagXO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagXO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.FirstCreated) {
		toSerialize["firstCreated"] = o.FirstCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return toSerialize, nil
}

func (o *TagXO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTagXO := _TagXO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagXO)

	if err != nil {
		return err
	}

	*o = TagXO(varTagXO)

	return err
}

type NullableTagXO struct {
	value *TagXO
	isSet bool
}

func (v NullableTagXO) Get() *TagXO {
	return v.value
}

func (v *NullableTagXO) Set(val *TagXO) {
	v.value = val
	v.isSet = true
}

func (v NullableTagXO) IsSet() bool {
	return v.isSet
}

func (v *NullableTagXO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagXO(val *TagXO) *NullableTagXO {
	return &NullableTagXO{value: val, isSet: true}
}

func (v NullableTagXO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagXO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


