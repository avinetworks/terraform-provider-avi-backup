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

// VIDCInfoClient is a client for avi VIDCInfo resource
type VIDCInfoClient struct {
	aviSession *session.AviSession
}

// NewVIDCInfoClient creates a new client for VIDCInfo resource
func NewVIDCInfoClient(aviSession *session.AviSession) *VIDCInfoClient {
	return &VIDCInfoClient{aviSession: aviSession}
}

func (client *VIDCInfoClient) getAPIPath(uuid string) string {
	path := "api/vidcinfo"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of VIDCInfo objects
func (client *VIDCInfoClient) GetAll(tenant ...string) ([]*models.VIDCInfo, error) {
	var plist []*models.VIDCInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing VIDCInfo by uuid
func (client *VIDCInfoClient) Get(uuid string, tenant ...string) (*models.VIDCInfo, error) {
	var obj *models.VIDCInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing VIDCInfo by name
func (client *VIDCInfoClient) GetByName(name string, tenant ...string) (*models.VIDCInfo, error) {
	var obj *models.VIDCInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("vidcinfo", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing VIDCInfo by filters like name, cloud, tenant
// Api creates VIDCInfo object with every call.
func (client *VIDCInfoClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.VIDCInfo, error) {
	var obj *models.VIDCInfo
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("vidcinfo", loc_tenant, newOptions...)
	return obj, err
}

// Create a new VIDCInfo object
func (client *VIDCInfoClient) Create(obj *models.VIDCInfo, tenant ...string) (*models.VIDCInfo, error) {
	var robj *models.VIDCInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing VIDCInfo object
func (client *VIDCInfoClient) Update(obj *models.VIDCInfo, tenant ...string) (*models.VIDCInfo, error) {
	var robj *models.VIDCInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing VIDCInfo object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.VIDCInfo
// or it should be json compatible of form map[string]interface{}
func (client *VIDCInfoClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.VIDCInfo, error) {
	var robj *models.VIDCInfo
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing VIDCInfo object with a given UUID
func (client *VIDCInfoClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing VIDCInfo object with a given name
func (client *VIDCInfoClient) DeleteByName(name string, tenant ...string) error {
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
func (client *VIDCInfoClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
