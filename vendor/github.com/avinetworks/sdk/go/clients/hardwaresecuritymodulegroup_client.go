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

// HardwareSecurityModuleGroupClient is a client for avi HardwareSecurityModuleGroup resource
type HardwareSecurityModuleGroupClient struct {
	aviSession *session.AviSession
}

// NewHardwareSecurityModuleGroupClient creates a new client for HardwareSecurityModuleGroup resource
func NewHardwareSecurityModuleGroupClient(aviSession *session.AviSession) *HardwareSecurityModuleGroupClient {
	return &HardwareSecurityModuleGroupClient{aviSession: aviSession}
}

func (client *HardwareSecurityModuleGroupClient) getAPIPath(uuid string) string {
	path := "api/hardwaresecuritymodulegroup"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of HardwareSecurityModuleGroup objects
func (client *HardwareSecurityModuleGroupClient) GetAll(tenant ...string) ([]*models.HardwareSecurityModuleGroup, error) {
	var plist []*models.HardwareSecurityModuleGroup
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing HardwareSecurityModuleGroup by uuid
func (client *HardwareSecurityModuleGroupClient) Get(uuid string, tenant ...string) (*models.HardwareSecurityModuleGroup, error) {
	var obj *models.HardwareSecurityModuleGroup
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing HardwareSecurityModuleGroup by name
func (client *HardwareSecurityModuleGroupClient) GetByName(name string, tenant ...string) (*models.HardwareSecurityModuleGroup, error) {
	var obj *models.HardwareSecurityModuleGroup
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("hardwaresecuritymodulegroup", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing HardwareSecurityModuleGroup by filters like name, cloud, tenant
// Api creates HardwareSecurityModuleGroup object with every call.
func (client *HardwareSecurityModuleGroupClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.HardwareSecurityModuleGroup, error) {
	var obj *models.HardwareSecurityModuleGroup
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("hardwaresecuritymodulegroup", loc_tenant, newOptions...)
	return obj, err
}

// Create a new HardwareSecurityModuleGroup object
func (client *HardwareSecurityModuleGroupClient) Create(obj *models.HardwareSecurityModuleGroup, tenant ...string) (*models.HardwareSecurityModuleGroup, error) {
	var robj *models.HardwareSecurityModuleGroup
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing HardwareSecurityModuleGroup object
func (client *HardwareSecurityModuleGroupClient) Update(obj *models.HardwareSecurityModuleGroup, tenant ...string) (*models.HardwareSecurityModuleGroup, error) {
	var robj *models.HardwareSecurityModuleGroup
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing HardwareSecurityModuleGroup object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.HardwareSecurityModuleGroup
// or it should be json compatible of form map[string]interface{}
func (client *HardwareSecurityModuleGroupClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.HardwareSecurityModuleGroup, error) {
	var robj *models.HardwareSecurityModuleGroup
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing HardwareSecurityModuleGroup object with a given UUID
func (client *HardwareSecurityModuleGroupClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing HardwareSecurityModuleGroup object with a given name
func (client *HardwareSecurityModuleGroupClient) DeleteByName(name string, tenant ...string) error {
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
func (client *HardwareSecurityModuleGroupClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
