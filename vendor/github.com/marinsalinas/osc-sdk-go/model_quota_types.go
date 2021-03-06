/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 0.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package oscgo

import (
	"bytes"
	"encoding/json"
)

// QuotaTypes One or more quotas.
type QuotaTypes struct {
	// The resource ID if it is a resource-specific quota, `global` if it is not.
	QuotaType *string `json:"QuotaType,omitempty"`
	// One or more quotas associated with the user.
	Quotas *[]Quota `json:"Quotas,omitempty"`
}

// GetQuotaType returns the QuotaType field value if set, zero value otherwise.
func (o *QuotaTypes) GetQuotaType() string {
	if o == nil || o.QuotaType == nil {
		var ret string
		return ret
	}
	return *o.QuotaType
}

// GetQuotaTypeOk returns a tuple with the QuotaType field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *QuotaTypes) GetQuotaTypeOk() (string, bool) {
	if o == nil || o.QuotaType == nil {
		var ret string
		return ret, false
	}
	return *o.QuotaType, true
}

// HasQuotaType returns a boolean if a field has been set.
func (o *QuotaTypes) HasQuotaType() bool {
	if o != nil && o.QuotaType != nil {
		return true
	}

	return false
}

// SetQuotaType gets a reference to the given string and assigns it to the QuotaType field.
func (o *QuotaTypes) SetQuotaType(v string) {
	o.QuotaType = &v
}

// GetQuotas returns the Quotas field value if set, zero value otherwise.
func (o *QuotaTypes) GetQuotas() []Quota {
	if o == nil || o.Quotas == nil {
		var ret []Quota
		return ret
	}
	return *o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *QuotaTypes) GetQuotasOk() ([]Quota, bool) {
	if o == nil || o.Quotas == nil {
		var ret []Quota
		return ret, false
	}
	return *o.Quotas, true
}

// HasQuotas returns a boolean if a field has been set.
func (o *QuotaTypes) HasQuotas() bool {
	if o != nil && o.Quotas != nil {
		return true
	}

	return false
}

// SetQuotas gets a reference to the given []Quota and assigns it to the Quotas field.
func (o *QuotaTypes) SetQuotas(v []Quota) {
	o.Quotas = &v
}

type NullableQuotaTypes struct {
	Value        QuotaTypes
	ExplicitNull bool
}

func (v NullableQuotaTypes) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableQuotaTypes) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
