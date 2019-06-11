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

// SystemConfigurationClient is a client for avi SystemConfiguration resource
type SystemConfigurationClient struct {
	aviSession *session.AviSession
}

// NewSystemConfigurationClient creates a new client for SystemConfiguration resource
func NewSystemConfigurationClient(aviSession *session.AviSession) *SystemConfigurationClient {
	return &SystemConfigurationClient{aviSession: aviSession}
}

func (client *SystemConfigurationClient) getAPIPath(uuid string) string {
	path := "api/systemconfiguration"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of SystemConfiguration objects
func (client *SystemConfigurationClient) GetAll(tenant ...string) ([]*models.SystemConfiguration, error) {
	var plist []*models.SystemConfiguration
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing SystemConfiguration by uuid
func (client *SystemConfigurationClient) Get(uuid string, tenant ...string) (*models.SystemConfiguration, error) {
	var obj *models.SystemConfiguration
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing SystemConfiguration by name
func (client *SystemConfigurationClient) GetByName(name string, tenant ...string) (*models.SystemConfiguration, error) {
	var obj *models.SystemConfiguration
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("systemconfiguration", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing SystemConfiguration by filters like name, cloud, tenant
// Api creates SystemConfiguration object with every call.
func (client *SystemConfigurationClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.SystemConfiguration, error) {
	var obj *models.SystemConfiguration
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("systemconfiguration", loc_tenant, newOptions...)
	return obj, err
}

// Create a new SystemConfiguration object
func (client *SystemConfigurationClient) Create(obj *models.SystemConfiguration, tenant ...string) (*models.SystemConfiguration, error) {
	var robj *models.SystemConfiguration
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing SystemConfiguration object
func (client *SystemConfigurationClient) Update(obj *models.SystemConfiguration, tenant ...string) (*models.SystemConfiguration, error) {
	var robj *models.SystemConfiguration
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing SystemConfiguration object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.SystemConfiguration
// or it should be json compatible of form map[string]interface{}
func (client *SystemConfigurationClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.SystemConfiguration, error) {
	var robj *models.SystemConfiguration
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing SystemConfiguration object with a given UUID
func (client *SystemConfigurationClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing SystemConfiguration object with a given name
func (client *SystemConfigurationClient) DeleteByName(name string, tenant ...string) error {
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
func (client *SystemConfigurationClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
