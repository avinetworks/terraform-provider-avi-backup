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

// ApplicationProfileClient is a client for avi ApplicationProfile resource
type ApplicationProfileClient struct {
	aviSession *session.AviSession
}

// NewApplicationProfileClient creates a new client for ApplicationProfile resource
func NewApplicationProfileClient(aviSession *session.AviSession) *ApplicationProfileClient {
	return &ApplicationProfileClient{aviSession: aviSession}
}

func (client *ApplicationProfileClient) getAPIPath(uuid string) string {
	path := "api/applicationprofile"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of ApplicationProfile objects
func (client *ApplicationProfileClient) GetAll(tenant ...string) ([]*models.ApplicationProfile, error) {
	var plist []*models.ApplicationProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing ApplicationProfile by uuid
func (client *ApplicationProfileClient) Get(uuid string, tenant ...string) (*models.ApplicationProfile, error) {
	var obj *models.ApplicationProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing ApplicationProfile by name
func (client *ApplicationProfileClient) GetByName(name string, tenant ...string) (*models.ApplicationProfile, error) {
	var obj *models.ApplicationProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("applicationprofile", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing ApplicationProfile by filters like name, cloud, tenant
// Api creates ApplicationProfile object with every call.
func (client *ApplicationProfileClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.ApplicationProfile, error) {
	var obj *models.ApplicationProfile
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("applicationprofile", loc_tenant, newOptions...)
	return obj, err
}

// Create a new ApplicationProfile object
func (client *ApplicationProfileClient) Create(obj *models.ApplicationProfile, tenant ...string) (*models.ApplicationProfile, error) {
	var robj *models.ApplicationProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing ApplicationProfile object
func (client *ApplicationProfileClient) Update(obj *models.ApplicationProfile, tenant ...string) (*models.ApplicationProfile, error) {
	var robj *models.ApplicationProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing ApplicationProfile object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.ApplicationProfile
// or it should be json compatible of form map[string]interface{}
func (client *ApplicationProfileClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.ApplicationProfile, error) {
	var robj *models.ApplicationProfile
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing ApplicationProfile object with a given UUID
func (client *ApplicationProfileClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing ApplicationProfile object with a given name
func (client *ApplicationProfileClient) DeleteByName(name string, tenant ...string) error {
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
func (client *ApplicationProfileClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
