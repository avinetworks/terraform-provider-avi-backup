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

// VirtualServiceClient is a client for avi VirtualService resource
type VirtualServiceClient struct {
	aviSession *session.AviSession
}

// NewVirtualServiceClient creates a new client for VirtualService resource
func NewVirtualServiceClient(aviSession *session.AviSession) *VirtualServiceClient {
	return &VirtualServiceClient{aviSession: aviSession}
}

func (client *VirtualServiceClient) getAPIPath(uuid string) string {
	path := "api/virtualservice"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of VirtualService objects
func (client *VirtualServiceClient) GetAll(tenant ...string) ([]*models.VirtualService, error) {
	var plist []*models.VirtualService
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing VirtualService by uuid
func (client *VirtualServiceClient) Get(uuid string, tenant ...string) (*models.VirtualService, error) {
	var obj *models.VirtualService
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing VirtualService by name
func (client *VirtualServiceClient) GetByName(name string, tenant ...string) (*models.VirtualService, error) {
	var obj *models.VirtualService
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("virtualservice", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing VirtualService by filters like name, cloud, tenant
// Api creates VirtualService object with every call.
func (client *VirtualServiceClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.VirtualService, error) {
	var obj *models.VirtualService
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("virtualservice", loc_tenant, newOptions...)
	return obj, err
}

// Create a new VirtualService object
func (client *VirtualServiceClient) Create(obj *models.VirtualService, tenant ...string) (*models.VirtualService, error) {
	var robj *models.VirtualService
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing VirtualService object
func (client *VirtualServiceClient) Update(obj *models.VirtualService, tenant ...string) (*models.VirtualService, error) {
	var robj *models.VirtualService
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing VirtualService object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.VirtualService
// or it should be json compatible of form map[string]interface{}
func (client *VirtualServiceClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.VirtualService, error) {
	var robj *models.VirtualService
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing VirtualService object with a given UUID
func (client *VirtualServiceClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing VirtualService object with a given name
func (client *VirtualServiceClient) DeleteByName(name string, tenant ...string) error {
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
func (client *VirtualServiceClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
