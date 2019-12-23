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

// CreateSnapshotRequest struct for CreateSnapshotRequest
type CreateSnapshotRequest struct {
	// A description for the snapshot.
	Description *string `json:"Description,omitempty"`
	// If `true`, checks whether you have the required permissions to perform the action.
	DryRun *bool `json:"DryRun,omitempty"`
	// The pre-signed URL of the snapshot you want to import from the OSU bucket.
	FileLocation *string `json:"FileLocation,omitempty"`
	// The size of the snapshot created in your account, in gibibytes (GiB). This size must be exactly the same as the source snapshot one. The maximum allowed size is 14,901 GiB.
	SnapshotSize *int64 `json:"SnapshotSize,omitempty"`
	// The name of the source Region, which must be the same as the Region of your account.
	SourceRegionName *string `json:"SourceRegionName,omitempty"`
	// The ID of the snapshot you want to copy.
	SourceSnapshotId *string `json:"SourceSnapshotId,omitempty"`
	// The ID of the volume you want to create a snapshot of.
	VolumeId *string `json:"VolumeId,omitempty"`
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateSnapshotRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDryRun returns the DryRun field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetDryRun() bool {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret
	}
	return *o.DryRun
}

// GetDryRunOk returns a tuple with the DryRun field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetDryRunOk() (bool, bool) {
	if o == nil || o.DryRun == nil {
		var ret bool
		return ret, false
	}
	return *o.DryRun, true
}

// HasDryRun returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasDryRun() bool {
	if o != nil && o.DryRun != nil {
		return true
	}

	return false
}

// SetDryRun gets a reference to the given bool and assigns it to the DryRun field.
func (o *CreateSnapshotRequest) SetDryRun(v bool) {
	o.DryRun = &v
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetFileLocation() string {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetFileLocationOk() (string, bool) {
	if o == nil || o.FileLocation == nil {
		var ret string
		return ret, false
	}
	return *o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasFileLocation() bool {
	if o != nil && o.FileLocation != nil {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *CreateSnapshotRequest) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetSnapshotSize returns the SnapshotSize field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetSnapshotSize() int64 {
	if o == nil || o.SnapshotSize == nil {
		var ret int64
		return ret
	}
	return *o.SnapshotSize
}

// GetSnapshotSizeOk returns a tuple with the SnapshotSize field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetSnapshotSizeOk() (int64, bool) {
	if o == nil || o.SnapshotSize == nil {
		var ret int64
		return ret, false
	}
	return *o.SnapshotSize, true
}

// HasSnapshotSize returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasSnapshotSize() bool {
	if o != nil && o.SnapshotSize != nil {
		return true
	}

	return false
}

// SetSnapshotSize gets a reference to the given int64 and assigns it to the SnapshotSize field.
func (o *CreateSnapshotRequest) SetSnapshotSize(v int64) {
	o.SnapshotSize = &v
}

// GetSourceRegionName returns the SourceRegionName field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetSourceRegionName() string {
	if o == nil || o.SourceRegionName == nil {
		var ret string
		return ret
	}
	return *o.SourceRegionName
}

// GetSourceRegionNameOk returns a tuple with the SourceRegionName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetSourceRegionNameOk() (string, bool) {
	if o == nil || o.SourceRegionName == nil {
		var ret string
		return ret, false
	}
	return *o.SourceRegionName, true
}

// HasSourceRegionName returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasSourceRegionName() bool {
	if o != nil && o.SourceRegionName != nil {
		return true
	}

	return false
}

// SetSourceRegionName gets a reference to the given string and assigns it to the SourceRegionName field.
func (o *CreateSnapshotRequest) SetSourceRegionName(v string) {
	o.SourceRegionName = &v
}

// GetSourceSnapshotId returns the SourceSnapshotId field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetSourceSnapshotId() string {
	if o == nil || o.SourceSnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SourceSnapshotId
}

// GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetSourceSnapshotIdOk() (string, bool) {
	if o == nil || o.SourceSnapshotId == nil {
		var ret string
		return ret, false
	}
	return *o.SourceSnapshotId, true
}

// HasSourceSnapshotId returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasSourceSnapshotId() bool {
	if o != nil && o.SourceSnapshotId != nil {
		return true
	}

	return false
}

// SetSourceSnapshotId gets a reference to the given string and assigns it to the SourceSnapshotId field.
func (o *CreateSnapshotRequest) SetSourceSnapshotId(v string) {
	o.SourceSnapshotId = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *CreateSnapshotRequest) GetVolumeId() string {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotRequest) GetVolumeIdOk() (string, bool) {
	if o == nil || o.VolumeId == nil {
		var ret string
		return ret, false
	}
	return *o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *CreateSnapshotRequest) HasVolumeId() bool {
	if o != nil && o.VolumeId != nil {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given string and assigns it to the VolumeId field.
func (o *CreateSnapshotRequest) SetVolumeId(v string) {
	o.VolumeId = &v
}

type NullableCreateSnapshotRequest struct {
	Value        CreateSnapshotRequest
	ExplicitNull bool
}

func (v NullableCreateSnapshotRequest) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCreateSnapshotRequest) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}