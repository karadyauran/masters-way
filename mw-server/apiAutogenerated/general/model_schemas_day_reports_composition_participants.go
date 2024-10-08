/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasDayReportsCompositionParticipants type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasDayReportsCompositionParticipants{}

// SchemasDayReportsCompositionParticipants struct for SchemasDayReportsCompositionParticipants
type SchemasDayReportsCompositionParticipants struct {
	DayReportId string `json:"dayReportId"`
	WayId string `json:"wayId"`
	WayName string `json:"wayName"`
}

type _SchemasDayReportsCompositionParticipants SchemasDayReportsCompositionParticipants

// NewSchemasDayReportsCompositionParticipants instantiates a new SchemasDayReportsCompositionParticipants object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasDayReportsCompositionParticipants(dayReportId string, wayId string, wayName string) *SchemasDayReportsCompositionParticipants {
	this := SchemasDayReportsCompositionParticipants{}
	this.DayReportId = dayReportId
	this.WayId = wayId
	this.WayName = wayName
	return &this
}

// NewSchemasDayReportsCompositionParticipantsWithDefaults instantiates a new SchemasDayReportsCompositionParticipants object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasDayReportsCompositionParticipantsWithDefaults() *SchemasDayReportsCompositionParticipants {
	this := SchemasDayReportsCompositionParticipants{}
	return &this
}

// GetDayReportId returns the DayReportId field value
func (o *SchemasDayReportsCompositionParticipants) GetDayReportId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DayReportId
}

// GetDayReportIdOk returns a tuple with the DayReportId field value
// and a boolean to check if the value has been set.
func (o *SchemasDayReportsCompositionParticipants) GetDayReportIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DayReportId, true
}

// SetDayReportId sets field value
func (o *SchemasDayReportsCompositionParticipants) SetDayReportId(v string) {
	o.DayReportId = v
}

// GetWayId returns the WayId field value
func (o *SchemasDayReportsCompositionParticipants) GetWayId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WayId
}

// GetWayIdOk returns a tuple with the WayId field value
// and a boolean to check if the value has been set.
func (o *SchemasDayReportsCompositionParticipants) GetWayIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WayId, true
}

// SetWayId sets field value
func (o *SchemasDayReportsCompositionParticipants) SetWayId(v string) {
	o.WayId = v
}

// GetWayName returns the WayName field value
func (o *SchemasDayReportsCompositionParticipants) GetWayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WayName
}

// GetWayNameOk returns a tuple with the WayName field value
// and a boolean to check if the value has been set.
func (o *SchemasDayReportsCompositionParticipants) GetWayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WayName, true
}

// SetWayName sets field value
func (o *SchemasDayReportsCompositionParticipants) SetWayName(v string) {
	o.WayName = v
}

func (o SchemasDayReportsCompositionParticipants) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasDayReportsCompositionParticipants) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dayReportId"] = o.DayReportId
	toSerialize["wayId"] = o.WayId
	toSerialize["wayName"] = o.WayName
	return toSerialize, nil
}

func (o *SchemasDayReportsCompositionParticipants) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dayReportId",
		"wayId",
		"wayName",
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

	varSchemasDayReportsCompositionParticipants := _SchemasDayReportsCompositionParticipants{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasDayReportsCompositionParticipants)

	if err != nil {
		return err
	}

	*o = SchemasDayReportsCompositionParticipants(varSchemasDayReportsCompositionParticipants)

	return err
}

type NullableSchemasDayReportsCompositionParticipants struct {
	value *SchemasDayReportsCompositionParticipants
	isSet bool
}

func (v NullableSchemasDayReportsCompositionParticipants) Get() *SchemasDayReportsCompositionParticipants {
	return v.value
}

func (v *NullableSchemasDayReportsCompositionParticipants) Set(val *SchemasDayReportsCompositionParticipants) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasDayReportsCompositionParticipants) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasDayReportsCompositionParticipants) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasDayReportsCompositionParticipants(val *SchemasDayReportsCompositionParticipants) *NullableSchemasDayReportsCompositionParticipants {
	return &NullableSchemasDayReportsCompositionParticipants{value: val, isSet: true}
}

func (v NullableSchemasDayReportsCompositionParticipants) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasDayReportsCompositionParticipants) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


