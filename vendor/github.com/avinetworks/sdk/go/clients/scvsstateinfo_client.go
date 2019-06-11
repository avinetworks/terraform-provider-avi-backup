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

// SCVsStateInfoClient is a client for avi SCVsStateInfo resource
type SCVsStateInfoClient struct {
	aviSession *session.AviSession
}

// NewSCVsStateInfoClient creates a new client for SCVsStateInfo resource
func NewSCVsStateInfoClient(aviSession *session.AviSession) *SCVsStateInfoClient {
	return &SCVsStateInfoClient{aviSession: aviSession}
}

func (client *SCVsStateInfoClient) getAPIPath(uuid string) string {
	path := "api/scvsstateinfo"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of SCVsStateInfo objects
func (client *SCVsStateInfoClient) GetAll(tenant ...string) ([]*models.SCVsStateInfo, error) {
	var plist []*models.SCVsStateInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing SCVsStateInfo by uuid
func (client *SCVsStateInfoClient) Get(uuid string, tenant ...string) (*models.SCVsStateInfo, error) {
	var obj *models.SCVsStateInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing SCVsStateInfo by name
func (client *SCVsStateInfoClient) GetByName(name string, tenant ...string) (*models.SCVsStateInfo, error) {
	var obj *models.SCVsStateInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("scvsstateinfo", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing SCVsStateInfo by filters like name, cloud, tenant
// Api creates SCVsStateInfo object with every call.
func (client *SCVsStateInfoClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.SCVsStateInfo, error) {
	var obj *models.SCVsStateInfo
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("scvsstateinfo", loc_tenant, newOptions...)
	return obj, err
}

// Create a new SCVsStateInfo object
func (client *SCVsStateInfoClient) Create(obj *models.SCVsStateInfo, tenant ...string) (*models.SCVsStateInfo, error) {
	var robj *models.SCVsStateInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing SCVsStateInfo object
func (client *SCVsStateInfoClient) Update(obj *models.SCVsStateInfo, tenant ...string) (*models.SCVsStateInfo, error) {
	var robj *models.SCVsStateInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing SCVsStateInfo object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.SCVsStateInfo
// or it should be json compatible of form map[string]interface{}
func (client *SCVsStateInfoClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.SCVsStateInfo, error) {
	var robj *models.SCVsStateInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing SCVsStateInfo object with a given UUID
func (client *SCVsStateInfoClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing SCVsStateInfo object with a given name
func (client *SCVsStateInfoClient) DeleteByName(name string, tenant ...string) error {
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
func (client *SCVsStateInfoClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
