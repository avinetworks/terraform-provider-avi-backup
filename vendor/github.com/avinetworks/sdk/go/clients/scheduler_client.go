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

// SchedulerClient is a client for avi Scheduler resource
type SchedulerClient struct {
	aviSession *session.AviSession
}

// NewSchedulerClient creates a new client for Scheduler resource
func NewSchedulerClient(aviSession *session.AviSession) *SchedulerClient {
	return &SchedulerClient{aviSession: aviSession}
}

func (client *SchedulerClient) getAPIPath(uuid string) string {
	path := "api/scheduler"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of Scheduler objects
func (client *SchedulerClient) GetAll(tenant ...string) ([]*models.Scheduler, error) {
	var plist []*models.Scheduler
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist, loc_tenant)
	return plist, err
}

// Get an existing Scheduler by uuid
func (client *SchedulerClient) Get(uuid string, tenant ...string) (*models.Scheduler, error) {
	var obj *models.Scheduler
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj, loc_tenant)
	return obj, err
}

// GetByName - Get an existing Scheduler by name
func (client *SchedulerClient) GetByName(name string, tenant ...string) (*models.Scheduler, error) {
	var obj *models.Scheduler
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.GetObjectByName("scheduler", name, &obj, loc_tenant)
	return obj, err
}

// GetObject - Get an existing Scheduler by filters like name, cloud, tenant
// Api creates Scheduler object with every call.
func (client *SchedulerClient) GetObject(tenant string, options ...session.ApiOptionsParams) (*models.Scheduler, error) {
	var obj *models.Scheduler
	loc_tenant := ""
	if tenant != "" {
		loc_tenant = tenant
	}
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("scheduler", loc_tenant, newOptions...)
	return obj, err
}

// Create a new Scheduler object
func (client *SchedulerClient) Create(obj *models.Scheduler, tenant ...string) (*models.Scheduler, error) {
	var robj *models.Scheduler
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj, loc_tenant)
	return robj, err
}

// Update an existing Scheduler object
func (client *SchedulerClient) Update(obj *models.Scheduler, tenant ...string) (*models.Scheduler, error) {
	var robj *models.Scheduler
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj, loc_tenant)
	return robj, err
}

// Patch an existing Scheduler object specified using uuid
// patchOp: Patch operation - add, replace, or delete
// patch: Patch payload should be compatible with the models.Scheduler
// or it should be json compatible of form map[string]interface{}
func (client *SchedulerClient) Patch(uuid string, patch interface{}, patchOp string, tenant ...string) (*models.Scheduler, error) {
	var robj *models.Scheduler
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	path := client.getAPIPath(uuid)
	err := client.aviSession.Patch(path, patch, patchOp, &robj, loc_tenant)
	return robj, err
}

// Delete an existing Scheduler object with a given UUID
func (client *SchedulerClient) Delete(uuid string, tenant ...string) error {
	loc_tenant := ""
	if len(tenant) != 0 {
		loc_tenant = tenant[0]
	}
	return client.aviSession.Delete(client.getAPIPath(uuid), loc_tenant)
}

// DeleteByName - Delete an existing Scheduler object with a given name
func (client *SchedulerClient) DeleteByName(name string, tenant ...string) error {
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
func (client *SchedulerClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
