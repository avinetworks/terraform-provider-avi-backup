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

// NatPolicyClient is a client for avi NatPolicy resource
type NatPolicyClient struct {
	aviSession *session.AviSession
}

// NewNatPolicyClient creates a new client for NatPolicy resource
func NewNatPolicyClient(aviSession *session.AviSession) *NatPolicyClient {
	return &NatPolicyClient{aviSession: aviSession}
}

func (client *NatPolicyClient) getAPIPath(uuid string) string {
	path := "api/natpolicy"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of NatPolicy objects
func (client *NatPolicyClient) GetAll(tenant ...string) ([]*models.NatPolicy, error) {
	var plist []*models.NatPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing NatPolicy by uuid
func (client *NatPolicyClient) Get(uuid string, tenant ...string) (*models.NatPolicy, error) {
	var obj *models.NatPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing NatPolicy by name
func (client *NatPolicyClient) GetByName(name string, tenant ...string) (*models.NatPolicy, error) {
	var obj *models.NatPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("natpolicy", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing NatPolicy by filters like name, cloud, tenant
// Api creates NatPolicy object with every call.
func (client *NatPolicyClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.NatPolicy, error) {
	var obj *models.NatPolicy
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("natpolicy", loc_tenant, newOptions...)
	return obj, err
}

// Create a new NatPolicy object
func (client *NatPolicyClient) Create(obj *models.NatPolicy, tenant ...string) (*models.NatPolicy, error) {
	var robj *models.NatPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing NatPolicy object
func (client *NatPolicyClient) Update(obj *models.NatPolicy, tenant ...string) (*models.NatPolicy, error) {
	var robj *models.NatPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing NatPolicy object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.NatPolicy
// or it should be json compatible of form map[string]interface{}
func (client *NatPolicyClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.NatPolicy, error) {
	var robj *models.NatPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing NatPolicy object with a given UUID
func (client *NatPolicyClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing NatPolicy object with a given name
func (client *NatPolicyClient) DeleteByName(name string, tenant ...string) error {
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
func (client *NatPolicyClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
