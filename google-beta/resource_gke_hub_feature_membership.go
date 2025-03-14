// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	gkehub "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/gkehub/beta"
)

func resourceGkeHubFeatureMembership() *schema.Resource {
	return &schema.Resource{
		Create: resourceGkeHubFeatureMembershipCreate,
		Read:   resourceGkeHubFeatureMembershipRead,
		Update: resourceGkeHubFeatureMembershipUpdate,
		Delete: resourceGkeHubFeatureMembershipDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGkeHubFeatureMembershipImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"configmanagement": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkeHubFeatureMembershipConfigmanagementSchema(),
			},

			"feature": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"membership": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},
		},
	}
}

func GkeHubFeatureMembershipConfigmanagementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"binauthz": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkeHubFeatureMembershipConfigmanagementBinauthzSchema(),
			},

			"config_sync": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkeHubFeatureMembershipConfigmanagementConfigSyncSchema(),
			},

			"hierarchy_controller": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkeHubFeatureMembershipConfigmanagementHierarchyControllerSchema(),
			},

			"policy_controller": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkeHubFeatureMembershipConfigmanagementPolicyControllerSchema(),
			},

			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func GkeHubFeatureMembershipConfigmanagementBinauthzSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func GkeHubFeatureMembershipConfigmanagementConfigSyncSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"git": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        GkeHubFeatureMembershipConfigmanagementConfigSyncGitSchema(),
			},

			"source_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func GkeHubFeatureMembershipConfigmanagementConfigSyncGitSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"https_proxy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"policy_dir": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"secret_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"sync_branch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"sync_repo": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"sync_rev": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"sync_wait_secs": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func GkeHubFeatureMembershipConfigmanagementHierarchyControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_hierarchical_resource_quota": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"enable_pod_tree_labels": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func GkeHubFeatureMembershipConfigmanagementPolicyControllerSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"audit_interval_seconds": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"exemptable_namespaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"log_denies_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"referential_rules_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"template_library_installed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func resourceGkeHubFeatureMembershipCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.FeatureMembership{
		Configmanagement: expandGkeHubFeatureMembershipConfigmanagement(d.Get("configmanagement")),
		Feature:          dcl.String(d.Get("feature").(string)),
		Location:         dcl.String(d.Get("location").(string)),
		Membership:       dcl.String(d.Get("membership").(string)),
		Project:          dcl.String(project),
	}
	lockName, err := replaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/features/{{feature}}/membershipId/{{membership}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	res, err := client.ApplyFeatureMembership(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating FeatureMembership: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FeatureMembership %q: %#v", d.Id(), res)

	return resourceGkeHubFeatureMembershipRead(d, meta)
}

func resourceGkeHubFeatureMembershipRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.FeatureMembership{
		Configmanagement: expandGkeHubFeatureMembershipConfigmanagement(d.Get("configmanagement")),
		Feature:          dcl.String(d.Get("feature").(string)),
		Location:         dcl.String(d.Get("location").(string)),
		Membership:       dcl.String(d.Get("membership").(string)),
		Project:          dcl.String(project),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	res, err := client.GetFeatureMembership(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("GkeHubFeatureMembership %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("configmanagement", flattenGkeHubFeatureMembershipConfigmanagement(res.Configmanagement)); err != nil {
		return fmt.Errorf("error setting configmanagement in state: %s", err)
	}
	if err = d.Set("feature", res.Feature); err != nil {
		return fmt.Errorf("error setting feature in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("membership", res.Membership); err != nil {
		return fmt.Errorf("error setting membership in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}

	return nil
}
func resourceGkeHubFeatureMembershipUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.FeatureMembership{
		Configmanagement: expandGkeHubFeatureMembershipConfigmanagement(d.Get("configmanagement")),
		Feature:          dcl.String(d.Get("feature").(string)),
		Location:         dcl.String(d.Get("location").(string)),
		Membership:       dcl.String(d.Get("membership").(string)),
		Project:          dcl.String(project),
	}
	lockName, err := replaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	directive := UpdateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	res, err := client.ApplyFeatureMembership(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating FeatureMembership: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FeatureMembership %q: %#v", d.Id(), res)

	return resourceGkeHubFeatureMembershipRead(d, meta)
}

func resourceGkeHubFeatureMembershipDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &gkehub.FeatureMembership{
		Configmanagement: expandGkeHubFeatureMembershipConfigmanagement(d.Get("configmanagement")),
		Feature:          dcl.String(d.Get("feature").(string)),
		Location:         dcl.String(d.Get("location").(string)),
		Membership:       dcl.String(d.Get("membership").(string)),
		Project:          dcl.String(project),
	}
	lockName, err := replaceVarsForId(d, config, "{{project}}/{{location}}/{{feature}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	log.Printf("[DEBUG] Deleting FeatureMembership %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLGkeHubClient(config, userAgent, billingProject)
	if err := client.DeleteFeatureMembership(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting FeatureMembership: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting FeatureMembership %q", d.Id())
	return nil
}

func resourceGkeHubFeatureMembershipImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/features/(?P<feature>[^/]+)/membershipId/(?P<membership>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<feature>[^/]+)/(?P<membership>[^/]+)",
		"(?P<location>[^/]+)/(?P<feature>[^/]+)/(?P<membership>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/features/{{feature}}/membershipId/{{membership}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandGkeHubFeatureMembershipConfigmanagement(o interface{}) *gkehub.FeatureMembershipConfigmanagement {
	if o == nil {
		return gkehub.EmptyFeatureMembershipConfigmanagement
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureMembershipConfigmanagement
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureMembershipConfigmanagement{
		Binauthz:            expandGkeHubFeatureMembershipConfigmanagementBinauthz(obj["binauthz"]),
		ConfigSync:          expandGkeHubFeatureMembershipConfigmanagementConfigSync(obj["config_sync"]),
		HierarchyController: expandGkeHubFeatureMembershipConfigmanagementHierarchyController(obj["hierarchy_controller"]),
		PolicyController:    expandGkeHubFeatureMembershipConfigmanagementPolicyController(obj["policy_controller"]),
		Version:             dcl.String(obj["version"].(string)),
	}
}

func flattenGkeHubFeatureMembershipConfigmanagement(obj *gkehub.FeatureMembershipConfigmanagement) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"binauthz":             flattenGkeHubFeatureMembershipConfigmanagementBinauthz(obj.Binauthz),
		"config_sync":          flattenGkeHubFeatureMembershipConfigmanagementConfigSync(obj.ConfigSync),
		"hierarchy_controller": flattenGkeHubFeatureMembershipConfigmanagementHierarchyController(obj.HierarchyController),
		"policy_controller":    flattenGkeHubFeatureMembershipConfigmanagementPolicyController(obj.PolicyController),
		"version":              obj.Version,
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureMembershipConfigmanagementBinauthz(o interface{}) *gkehub.FeatureMembershipConfigmanagementBinauthz {
	if o == nil {
		return gkehub.EmptyFeatureMembershipConfigmanagementBinauthz
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureMembershipConfigmanagementBinauthz
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureMembershipConfigmanagementBinauthz{
		Enabled: dcl.Bool(obj["enabled"].(bool)),
	}
}

func flattenGkeHubFeatureMembershipConfigmanagementBinauthz(obj *gkehub.FeatureMembershipConfigmanagementBinauthz) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"enabled": obj.Enabled,
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureMembershipConfigmanagementConfigSync(o interface{}) *gkehub.FeatureMembershipConfigmanagementConfigSync {
	if o == nil {
		return gkehub.EmptyFeatureMembershipConfigmanagementConfigSync
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureMembershipConfigmanagementConfigSync
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureMembershipConfigmanagementConfigSync{
		Git:          expandGkeHubFeatureMembershipConfigmanagementConfigSyncGit(obj["git"]),
		SourceFormat: dcl.String(obj["source_format"].(string)),
	}
}

func flattenGkeHubFeatureMembershipConfigmanagementConfigSync(obj *gkehub.FeatureMembershipConfigmanagementConfigSync) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"git":           flattenGkeHubFeatureMembershipConfigmanagementConfigSyncGit(obj.Git),
		"source_format": obj.SourceFormat,
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureMembershipConfigmanagementConfigSyncGit(o interface{}) *gkehub.FeatureMembershipConfigmanagementConfigSyncGit {
	if o == nil {
		return gkehub.EmptyFeatureMembershipConfigmanagementConfigSyncGit
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureMembershipConfigmanagementConfigSyncGit
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureMembershipConfigmanagementConfigSyncGit{
		HttpsProxy:   dcl.String(obj["https_proxy"].(string)),
		PolicyDir:    dcl.String(obj["policy_dir"].(string)),
		SecretType:   dcl.String(obj["secret_type"].(string)),
		SyncBranch:   dcl.String(obj["sync_branch"].(string)),
		SyncRepo:     dcl.String(obj["sync_repo"].(string)),
		SyncRev:      dcl.String(obj["sync_rev"].(string)),
		SyncWaitSecs: dcl.String(obj["sync_wait_secs"].(string)),
	}
}

func flattenGkeHubFeatureMembershipConfigmanagementConfigSyncGit(obj *gkehub.FeatureMembershipConfigmanagementConfigSyncGit) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"https_proxy":    obj.HttpsProxy,
		"policy_dir":     obj.PolicyDir,
		"secret_type":    obj.SecretType,
		"sync_branch":    obj.SyncBranch,
		"sync_repo":      obj.SyncRepo,
		"sync_rev":       obj.SyncRev,
		"sync_wait_secs": obj.SyncWaitSecs,
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureMembershipConfigmanagementHierarchyController(o interface{}) *gkehub.FeatureMembershipConfigmanagementHierarchyController {
	if o == nil {
		return gkehub.EmptyFeatureMembershipConfigmanagementHierarchyController
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureMembershipConfigmanagementHierarchyController
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureMembershipConfigmanagementHierarchyController{
		EnableHierarchicalResourceQuota: dcl.Bool(obj["enable_hierarchical_resource_quota"].(bool)),
		EnablePodTreeLabels:             dcl.Bool(obj["enable_pod_tree_labels"].(bool)),
		Enabled:                         dcl.Bool(obj["enabled"].(bool)),
	}
}

func flattenGkeHubFeatureMembershipConfigmanagementHierarchyController(obj *gkehub.FeatureMembershipConfigmanagementHierarchyController) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"enable_hierarchical_resource_quota": obj.EnableHierarchicalResourceQuota,
		"enable_pod_tree_labels":             obj.EnablePodTreeLabels,
		"enabled":                            obj.Enabled,
	}

	return []interface{}{transformed}

}

func expandGkeHubFeatureMembershipConfigmanagementPolicyController(o interface{}) *gkehub.FeatureMembershipConfigmanagementPolicyController {
	if o == nil {
		return gkehub.EmptyFeatureMembershipConfigmanagementPolicyController
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return gkehub.EmptyFeatureMembershipConfigmanagementPolicyController
	}
	obj := objArr[0].(map[string]interface{})
	return &gkehub.FeatureMembershipConfigmanagementPolicyController{
		AuditIntervalSeconds:     dcl.String(obj["audit_interval_seconds"].(string)),
		Enabled:                  dcl.Bool(obj["enabled"].(bool)),
		ExemptableNamespaces:     expandStringArray(obj["exemptable_namespaces"]),
		LogDeniesEnabled:         dcl.Bool(obj["log_denies_enabled"].(bool)),
		ReferentialRulesEnabled:  dcl.Bool(obj["referential_rules_enabled"].(bool)),
		TemplateLibraryInstalled: dcl.Bool(obj["template_library_installed"].(bool)),
	}
}

func flattenGkeHubFeatureMembershipConfigmanagementPolicyController(obj *gkehub.FeatureMembershipConfigmanagementPolicyController) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"audit_interval_seconds":     obj.AuditIntervalSeconds,
		"enabled":                    obj.Enabled,
		"exemptable_namespaces":      obj.ExemptableNamespaces,
		"log_denies_enabled":         obj.LogDeniesEnabled,
		"referential_rules_enabled":  obj.ReferentialRulesEnabled,
		"template_library_installed": obj.TemplateLibraryInstalled,
	}

	return []interface{}{transformed}

}
