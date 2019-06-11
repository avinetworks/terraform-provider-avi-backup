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

// L4PolicySetClient is a client for avi L4PolicySet resource
type L4PolicySetClient struct {
	aviSession *session.AviSession
}

// NewL4PolicySetClient creates a new client for L4PolicySet resource
func NewL4PolicySetClient(aviSession *session.AviSession) *L4PolicySetClient {
	return &L4PolicySetClient{aviSession: aviSession}
}

func (client *L4PolicySetClient) getAPIPath(uuid string) string {
	path := "api/l4policyset"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of L4PolicySet objects
func (client *L4PolicySetClient) GetAll(tenant ...string) ([]*models.L4PolicySet, error) {
	var plist []*models.L4PolicySet
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing L4PolicySet by uuid
func (client *L4PolicySetClient) Get(uuid string, tenant ...string) (*models.L4PolicySet, error) {
	var obj *models.L4PolicySet
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing L4PolicySet by name
func (client *L4PolicySetClient) GetByName(name string, tenant ...string) (*models.L4PolicySet, error) {
	var obj *models.L4PolicySet
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("l4policyset", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing L4PolicySet by filters like name, cloud, tenant
// Api creates L4PolicySet object with every call.
func (client *L4PolicySetClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.L4PolicySet, error) {
	var obj *models.L4PolicySet
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("l4policyset", loc_tenant, newOptions...)
	return obj, err
}

// Create a new L4PolicySet object
func (client *L4PolicySetClient) Create(obj *models.L4PolicySet, tenant ...string) (*models.L4PolicySet, error) {
	var robj *models.L4PolicySet
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing L4PolicySet object
func (client *L4PolicySetClient) Update(obj *models.L4PolicySet, tenant ...string) (*models.L4PolicySet, error) {
	var robj *models.L4PolicySet
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing L4PolicySet object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.L4PolicySet
// or it should be json compatible of form map[string]interface{}
func (client *L4PolicySetClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.L4PolicySet, error) {
	var robj *models.L4PolicySet
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing L4PolicySet object with a given UUID
func (client *L4PolicySetClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing L4PolicySet object with a given name
func (client *L4PolicySetClient) DeleteByName(name string, tenant ...string) error {
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
func (client *L4PolicySetClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
