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

// DNSPolicyClient is a client for avi DNSPolicy resource
type DNSPolicyClient struct {
	aviSession *session.AviSession
}

// NewDNSPolicyClient creates a new client for DNSPolicy resource
func NewDNSPolicyClient(aviSession *session.AviSession) *DNSPolicyClient {
	return &DNSPolicyClient{aviSession: aviSession}
}

func (client *DNSPolicyClient) getAPIPath(uuid string) string {
	path := "api/dnspolicy"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of DNSPolicy objects
func (client *DNSPolicyClient) GetAll(tenant ...string) ([]*models.DNSPolicy, error) {
	var plist []*models.DNSPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing DNSPolicy by uuid
func (client *DNSPolicyClient) Get(uuid string, tenant ...string) (*models.DNSPolicy, error) {
	var obj *models.DNSPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing DNSPolicy by name
func (client *DNSPolicyClient) GetByName(name string, tenant ...string) (*models.DNSPolicy, error) {
	var obj *models.DNSPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("dnspolicy", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing DNSPolicy by filters like name, cloud, tenant
// Api creates DNSPolicy object with every call.
func (client *DNSPolicyClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.DNSPolicy, error) {
	var obj *models.DNSPolicy
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("dnspolicy", loc_tenant, newOptions...)
	return obj, err
}

// Create a new DNSPolicy object
func (client *DNSPolicyClient) Create(obj *models.DNSPolicy, tenant ...string) (*models.DNSPolicy, error) {
	var robj *models.DNSPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing DNSPolicy object
func (client *DNSPolicyClient) Update(obj *models.DNSPolicy, tenant ...string) (*models.DNSPolicy, error) {
	var robj *models.DNSPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing DNSPolicy object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.DNSPolicy
// or it should be json compatible of form map[string]interface{}
func (client *DNSPolicyClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.DNSPolicy, error) {
	var robj *models.DNSPolicy
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing DNSPolicy object with a given UUID
func (client *DNSPolicyClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing DNSPolicy object with a given name
func (client *DNSPolicyClient) DeleteByName(name string, tenant ...string) error {
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
func (client *DNSPolicyClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
