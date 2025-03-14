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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	assuredworkloads "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/assuredworkloads/beta"
)

func resourceAssuredWorkloadsWorkload() *schema.Resource {
	return &schema.Resource{
		Create: resourceAssuredWorkloadsWorkloadCreate,
		Read:   resourceAssuredWorkloadsWorkloadRead,
		Update: resourceAssuredWorkloadsWorkloadUpdate,
		Delete: resourceAssuredWorkloadsWorkloadDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAssuredWorkloadsWorkloadImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"billing_account": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"compliance_regime": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"COMPLIANCE_REGIME_UNSPECIFIED", "IL4", "CJIS", "FEDRAMP_HIGH", "FEDRAMP_MODERATE", "US_REGIONAL_ACCESS", ""}, false),
			},

			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: ``,
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"organization": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"kms_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        AssuredWorkloadsWorkloadKmsSettingsSchema(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"provisioned_resources_parent": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"resource_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        AssuredWorkloadsWorkloadResourceSettingsSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"resources": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: ``,
				Elem:        AssuredWorkloadsWorkloadResourcesSchema(),
			},
		},
	}
}

func AssuredWorkloadsWorkloadKmsSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"next_rotation_time": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"rotation_period": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func AssuredWorkloadsWorkloadResourceSettingsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"resource_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"RESOURCE_TYPE_UNSPECIFIED", "CONSUMER_PROJECT", "ENCRYPTION_KEYS_PROJECT", "KEYRING", "CONSUMER_FOLDER", ""}, false),
			},
		},
	}
}

func AssuredWorkloadsWorkloadResourcesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: ``,
			},

			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func resourceAssuredWorkloadsWorkloadCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &assuredworkloads.Workload{
		BillingAccount:             dcl.String(d.Get("billing_account").(string)),
		ComplianceRegime:           assuredworkloads.WorkloadComplianceRegimeEnumRef(d.Get("compliance_regime").(string)),
		DisplayName:                dcl.String(d.Get("display_name").(string)),
		Location:                   dcl.String(d.Get("location").(string)),
		Organization:               dcl.String(d.Get("organization").(string)),
		KmsSettings:                expandAssuredWorkloadsWorkloadKmsSettings(d.Get("kms_settings")),
		Labels:                     checkStringMap(d.Get("labels")),
		ProvisionedResourcesParent: dcl.String(d.Get("provisioned_resources_parent").(string)),
		ResourceSettings:           expandAssuredWorkloadsWorkloadResourceSettingsArray(d.Get("resource_settings")),
	}

	id, err := replaceVarsForId(d, config, "organizations/{{organization}}/locations/{{location}}/workloads/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	createDirective := CreateDirective
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLAssuredWorkloadsClient(config, userAgent, billingProject)
	res, err := client.ApplyWorkload(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Workload: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Workload %q: %#v", d.Id(), res)

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	// Id has a server-generated value, set again after creation
	id, err = replaceVarsForId(d, config, "organizations/{{organization}}/locations/{{location}}/workloads/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return resourceAssuredWorkloadsWorkloadRead(d, meta)
}

func resourceAssuredWorkloadsWorkloadRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &assuredworkloads.Workload{
		BillingAccount:             dcl.String(d.Get("billing_account").(string)),
		ComplianceRegime:           assuredworkloads.WorkloadComplianceRegimeEnumRef(d.Get("compliance_regime").(string)),
		DisplayName:                dcl.String(d.Get("display_name").(string)),
		Location:                   dcl.String(d.Get("location").(string)),
		Organization:               dcl.String(d.Get("organization").(string)),
		KmsSettings:                expandAssuredWorkloadsWorkloadKmsSettings(d.Get("kms_settings")),
		Labels:                     checkStringMap(d.Get("labels")),
		ProvisionedResourcesParent: dcl.String(d.Get("provisioned_resources_parent").(string)),
		ResourceSettings:           expandAssuredWorkloadsWorkloadResourceSettingsArray(d.Get("resource_settings")),
		Name:                       dcl.StringOrNil(d.Get("name").(string)),
	}

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLAssuredWorkloadsClient(config, userAgent, billingProject)
	res, err := client.GetWorkload(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("AssuredWorkloadsWorkload %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("billing_account", res.BillingAccount); err != nil {
		return fmt.Errorf("error setting billing_account in state: %s", err)
	}
	if err = d.Set("compliance_regime", res.ComplianceRegime); err != nil {
		return fmt.Errorf("error setting compliance_regime in state: %s", err)
	}
	if err = d.Set("display_name", res.DisplayName); err != nil {
		return fmt.Errorf("error setting display_name in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("organization", res.Organization); err != nil {
		return fmt.Errorf("error setting organization in state: %s", err)
	}
	if err = d.Set("kms_settings", flattenAssuredWorkloadsWorkloadKmsSettings(res.KmsSettings)); err != nil {
		return fmt.Errorf("error setting kms_settings in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("provisioned_resources_parent", res.ProvisionedResourcesParent); err != nil {
		return fmt.Errorf("error setting provisioned_resources_parent in state: %s", err)
	}
	if err = d.Set("resource_settings", flattenAssuredWorkloadsWorkloadResourceSettingsArray(res.ResourceSettings)); err != nil {
		return fmt.Errorf("error setting resource_settings in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("resources", flattenAssuredWorkloadsWorkloadResourcesArray(res.Resources)); err != nil {
		return fmt.Errorf("error setting resources in state: %s", err)
	}

	return nil
}
func resourceAssuredWorkloadsWorkloadUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &assuredworkloads.Workload{
		BillingAccount:             dcl.String(d.Get("billing_account").(string)),
		ComplianceRegime:           assuredworkloads.WorkloadComplianceRegimeEnumRef(d.Get("compliance_regime").(string)),
		DisplayName:                dcl.String(d.Get("display_name").(string)),
		Location:                   dcl.String(d.Get("location").(string)),
		Organization:               dcl.String(d.Get("organization").(string)),
		KmsSettings:                expandAssuredWorkloadsWorkloadKmsSettings(d.Get("kms_settings")),
		Labels:                     checkStringMap(d.Get("labels")),
		ProvisionedResourcesParent: dcl.String(d.Get("provisioned_resources_parent").(string)),
		ResourceSettings:           expandAssuredWorkloadsWorkloadResourceSettingsArray(d.Get("resource_settings")),
		Name:                       dcl.StringOrNil(d.Get("name").(string)),
	}
	// Construct state hint from old values
	old := &assuredworkloads.Workload{
		BillingAccount:             dcl.String(oldValue(d.GetChange("billing_account")).(string)),
		ComplianceRegime:           assuredworkloads.WorkloadComplianceRegimeEnumRef(oldValue(d.GetChange("compliance_regime")).(string)),
		DisplayName:                dcl.String(oldValue(d.GetChange("display_name")).(string)),
		Location:                   dcl.String(oldValue(d.GetChange("location")).(string)),
		Organization:               dcl.String(oldValue(d.GetChange("organization")).(string)),
		KmsSettings:                expandAssuredWorkloadsWorkloadKmsSettings(oldValue(d.GetChange("kms_settings"))),
		Labels:                     checkStringMap(oldValue(d.GetChange("labels"))),
		ProvisionedResourcesParent: dcl.String(oldValue(d.GetChange("provisioned_resources_parent")).(string)),
		ResourceSettings:           expandAssuredWorkloadsWorkloadResourceSettingsArray(oldValue(d.GetChange("resource_settings"))),
		Name:                       dcl.StringOrNil(oldValue(d.GetChange("name")).(string)),
	}
	directive := UpdateDirective
	directive = append(directive, dcl.WithStateHint(old))
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLAssuredWorkloadsClient(config, userAgent, billingProject)
	res, err := client.ApplyWorkload(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Workload: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Workload %q: %#v", d.Id(), res)

	return resourceAssuredWorkloadsWorkloadRead(d, meta)
}

func resourceAssuredWorkloadsWorkloadDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &assuredworkloads.Workload{
		BillingAccount:             dcl.String(d.Get("billing_account").(string)),
		ComplianceRegime:           assuredworkloads.WorkloadComplianceRegimeEnumRef(d.Get("compliance_regime").(string)),
		DisplayName:                dcl.String(d.Get("display_name").(string)),
		Location:                   dcl.String(d.Get("location").(string)),
		Organization:               dcl.String(d.Get("organization").(string)),
		KmsSettings:                expandAssuredWorkloadsWorkloadKmsSettings(d.Get("kms_settings")),
		Labels:                     checkStringMap(d.Get("labels")),
		ProvisionedResourcesParent: dcl.String(d.Get("provisioned_resources_parent").(string)),
		ResourceSettings:           expandAssuredWorkloadsWorkloadResourceSettingsArray(d.Get("resource_settings")),
		Name:                       dcl.StringOrNil(d.Get("name").(string)),
	}

	log.Printf("[DEBUG] Deleting Workload %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLAssuredWorkloadsClient(config, userAgent, billingProject)
	if err := client.DeleteWorkload(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Workload: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Workload %q", d.Id())
	return nil
}

func resourceAssuredWorkloadsWorkloadImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"organizations/(?P<organization>[^/]+)/locations/(?P<location>[^/]+)/workloads/(?P<name>[^/]+)",
		"(?P<organization>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "organizations/{{organization}}/locations/{{location}}/workloads/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandAssuredWorkloadsWorkloadKmsSettings(o interface{}) *assuredworkloads.WorkloadKmsSettings {
	if o == nil {
		return assuredworkloads.EmptyWorkloadKmsSettings
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return assuredworkloads.EmptyWorkloadKmsSettings
	}
	obj := objArr[0].(map[string]interface{})
	return &assuredworkloads.WorkloadKmsSettings{
		NextRotationTime: dcl.String(obj["next_rotation_time"].(string)),
		RotationPeriod:   dcl.String(obj["rotation_period"].(string)),
	}
}

func flattenAssuredWorkloadsWorkloadKmsSettings(obj *assuredworkloads.WorkloadKmsSettings) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"next_rotation_time": obj.NextRotationTime,
		"rotation_period":    obj.RotationPeriod,
	}

	return []interface{}{transformed}

}
func expandAssuredWorkloadsWorkloadResourceSettingsArray(o interface{}) []assuredworkloads.WorkloadResourceSettings {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]assuredworkloads.WorkloadResourceSettings, 0, len(objs))
	for _, item := range objs {
		i := expandAssuredWorkloadsWorkloadResourceSettings(item)
		items = append(items, *i)
	}

	return items
}

func expandAssuredWorkloadsWorkloadResourceSettings(o interface{}) *assuredworkloads.WorkloadResourceSettings {
	if o == nil {
		return assuredworkloads.EmptyWorkloadResourceSettings
	}

	obj := o.(map[string]interface{})
	return &assuredworkloads.WorkloadResourceSettings{
		ResourceId:   dcl.String(obj["resource_id"].(string)),
		ResourceType: assuredworkloads.WorkloadResourceSettingsResourceTypeEnumRef(obj["resource_type"].(string)),
	}
}

func flattenAssuredWorkloadsWorkloadResourceSettingsArray(objs []assuredworkloads.WorkloadResourceSettings) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenAssuredWorkloadsWorkloadResourceSettings(&item)
		items = append(items, i)
	}

	return items
}

func flattenAssuredWorkloadsWorkloadResourceSettings(obj *assuredworkloads.WorkloadResourceSettings) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"resource_id":   obj.ResourceId,
		"resource_type": obj.ResourceType,
	}

	return transformed

}

func flattenAssuredWorkloadsWorkloadResourcesArray(objs []assuredworkloads.WorkloadResources) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenAssuredWorkloadsWorkloadResources(&item)
		items = append(items, i)
	}

	return items
}

func flattenAssuredWorkloadsWorkloadResources(obj *assuredworkloads.WorkloadResources) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"resource_id":   obj.ResourceId,
		"resource_type": obj.ResourceType,
	}

	return transformed

}
