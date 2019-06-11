/***************************************************************************
 *
 * AVI CONFIDENTIAL
 * __________________
 *
 * [2013] - [2018] Avi Networks Incorporated
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is, and remains the property
 * of Avi Networks Incorporated and its suppliers, if any. The intellectual
 * and technical concepts contained herein are proprietary to Avi Networks
 * Incorporated, and its suppliers and are covered by U.S. and Foreign
 * Patents, patents in process, and are protected by trade secret or
 * copyright law, and other laws. Dissemination of this information or
 * reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Avi Networks Incorporated.
 */

package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// GslbGeoDbProfileClient is a client for avi GslbGeoDbProfile resource
type GslbGeoDbProfileClient struct {
	aviSession *session.AviSession
}

// NewGslbGeoDbProfileClient creates a new client for GslbGeoDbProfile resource
func NewGslbGeoDbProfileClient(aviSession *session.AviSession) *GslbGeoDbProfileClient {
	return &GslbGeoDbProfileClient{aviSession: aviSession}
}

func (client *GslbGeoDbProfileClient) getAPIPath(uuid string) string {
	path := "api/gslbgeodbprofile"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of GslbGeoDbProfile objects
func (client *GslbGeoDbProfileClient) GetAll(tenant ...string) ([]*models.GslbGeoDbProfile, error) {
	var plist []*models.GslbGeoDbProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing GslbGeoDbProfile by uuid
func (client *GslbGeoDbProfileClient) Get(uuid string, tenant ...string) (*models.GslbGeoDbProfile, error) {
	var obj *models.GslbGeoDbProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing GslbGeoDbProfile by name
func (client *GslbGeoDbProfileClient) GetByName(name string, tenant ...string) (*models.GslbGeoDbProfile, error) {
	var obj *models.GslbGeoDbProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("gslbgeodbprofile", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing GslbGeoDbProfile by filters like name, cloud, tenant
// Api creates GslbGeoDbProfile object with every call.
func (client *GslbGeoDbProfileClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.GslbGeoDbProfile, error) {
	var obj *models.GslbGeoDbProfile
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("gslbgeodbprofile", loc_tenant, newOptions...)
	return obj, err
}

// Create a new GslbGeoDbProfile object
func (client *GslbGeoDbProfileClient) Create(obj *models.GslbGeoDbProfile, tenant ...string) (*models.GslbGeoDbProfile, error) {
	var robj *models.GslbGeoDbProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing GslbGeoDbProfile object
func (client *GslbGeoDbProfileClient) Update(obj *models.GslbGeoDbProfile, tenant ...string) (*models.GslbGeoDbProfile, error) {
	var robj *models.GslbGeoDbProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing GslbGeoDbProfile object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.GslbGeoDbProfile
// or it should be json compatible of form map[string]interface{}
func (client *GslbGeoDbProfileClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.GslbGeoDbProfile, error) {
	var robj *models.GslbGeoDbProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing GslbGeoDbProfile object with a given UUID
func (client *GslbGeoDbProfileClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing GslbGeoDbProfile object with a given name
func (client *GslbGeoDbProfileClient) DeleteByName(name string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	res, err := client.GetByName(name, loc_tenant)
	if err != nil {
		return err
	}
	return client.Delete(*res.UUID, loc_tenant)
}

// GetAviSession
func (client *GslbGeoDbProfileClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
