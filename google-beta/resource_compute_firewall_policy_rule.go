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
	compute "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/compute/beta"
)

func resourceComputeFirewallPolicyRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeFirewallPolicyRuleCreate,
		Read:   resourceComputeFirewallPolicyRuleRead,
		Update: resourceComputeFirewallPolicyRuleUpdate,
		Delete: resourceComputeFirewallPolicyRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeFirewallPolicyRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Required:    true,
				Description: ``,
			},

			"direction": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"INGRESS", "EGRESS", ""}, false),
			},

			"firewall_policy": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"match": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        ComputeFirewallPolicyRuleMatchSchema(),
			},

			"priority": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"disabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"enable_logging": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"target_resources": {
				Type:             schema.TypeList,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
				Elem:             &schema.Schema{Type: schema.TypeString},
			},

			"target_service_accounts": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"kind": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"rule_tuple_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func ComputeFirewallPolicyRuleMatchSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"layer4_configs": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				Elem:        ComputeFirewallPolicyRuleMatchLayer4ConfigsSchema(),
			},

			"dest_ip_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"src_ip_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func ComputeFirewallPolicyRuleMatchLayer4ConfigsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ip_protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: ``,
			},

			"ports": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceComputeFirewallPolicyRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &compute.FirewallPolicyRule{
		Action:                dcl.String(d.Get("action").(string)),
		Direction:             compute.FirewallPolicyRuleDirectionEnumRef(d.Get("direction").(string)),
		FirewallPolicy:        dcl.String(d.Get("firewall_policy").(string)),
		Match:                 expandComputeFirewallPolicyRuleMatch(d.Get("match")),
		Priority:              dcl.Int64(int64(d.Get("priority").(int))),
		Description:           dcl.String(d.Get("description").(string)),
		Disabled:              dcl.Bool(d.Get("disabled").(bool)),
		EnableLogging:         dcl.Bool(d.Get("enable_logging").(bool)),
		TargetResources:       expandStringArray(d.Get("target_resources")),
		TargetServiceAccounts: expandStringArray(d.Get("target_service_accounts")),
	}

	id, err := replaceVarsForId(d, config, "locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}")
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
	client := NewDCLComputeClient(config, userAgent, billingProject)
	res, err := client.ApplyFirewallPolicyRule(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating FirewallPolicyRule: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FirewallPolicyRule %q: %#v", d.Id(), res)

	return resourceComputeFirewallPolicyRuleRead(d, meta)
}

func resourceComputeFirewallPolicyRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &compute.FirewallPolicyRule{
		Action:                dcl.String(d.Get("action").(string)),
		Direction:             compute.FirewallPolicyRuleDirectionEnumRef(d.Get("direction").(string)),
		FirewallPolicy:        dcl.String(d.Get("firewall_policy").(string)),
		Match:                 expandComputeFirewallPolicyRuleMatch(d.Get("match")),
		Priority:              dcl.Int64(int64(d.Get("priority").(int))),
		Description:           dcl.String(d.Get("description").(string)),
		Disabled:              dcl.Bool(d.Get("disabled").(bool)),
		EnableLogging:         dcl.Bool(d.Get("enable_logging").(bool)),
		TargetResources:       expandStringArray(d.Get("target_resources")),
		TargetServiceAccounts: expandStringArray(d.Get("target_service_accounts")),
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
	client := NewDCLComputeClient(config, userAgent, billingProject)
	res, err := client.GetFirewallPolicyRule(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("ComputeFirewallPolicyRule %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("action", res.Action); err != nil {
		return fmt.Errorf("error setting action in state: %s", err)
	}
	if err = d.Set("direction", res.Direction); err != nil {
		return fmt.Errorf("error setting direction in state: %s", err)
	}
	if err = d.Set("firewall_policy", res.FirewallPolicy); err != nil {
		return fmt.Errorf("error setting firewall_policy in state: %s", err)
	}
	if err = d.Set("match", flattenComputeFirewallPolicyRuleMatch(res.Match)); err != nil {
		return fmt.Errorf("error setting match in state: %s", err)
	}
	if err = d.Set("priority", res.Priority); err != nil {
		return fmt.Errorf("error setting priority in state: %s", err)
	}
	if err = d.Set("description", res.Description); err != nil {
		return fmt.Errorf("error setting description in state: %s", err)
	}
	if err = d.Set("disabled", res.Disabled); err != nil {
		return fmt.Errorf("error setting disabled in state: %s", err)
	}
	if err = d.Set("enable_logging", res.EnableLogging); err != nil {
		return fmt.Errorf("error setting enable_logging in state: %s", err)
	}
	if err = d.Set("target_resources", res.TargetResources); err != nil {
		return fmt.Errorf("error setting target_resources in state: %s", err)
	}
	if err = d.Set("target_service_accounts", res.TargetServiceAccounts); err != nil {
		return fmt.Errorf("error setting target_service_accounts in state: %s", err)
	}
	if err = d.Set("kind", res.Kind); err != nil {
		return fmt.Errorf("error setting kind in state: %s", err)
	}
	if err = d.Set("rule_tuple_count", res.RuleTupleCount); err != nil {
		return fmt.Errorf("error setting rule_tuple_count in state: %s", err)
	}

	return nil
}
func resourceComputeFirewallPolicyRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &compute.FirewallPolicyRule{
		Action:                dcl.String(d.Get("action").(string)),
		Direction:             compute.FirewallPolicyRuleDirectionEnumRef(d.Get("direction").(string)),
		FirewallPolicy:        dcl.String(d.Get("firewall_policy").(string)),
		Match:                 expandComputeFirewallPolicyRuleMatch(d.Get("match")),
		Priority:              dcl.Int64(int64(d.Get("priority").(int))),
		Description:           dcl.String(d.Get("description").(string)),
		Disabled:              dcl.Bool(d.Get("disabled").(bool)),
		EnableLogging:         dcl.Bool(d.Get("enable_logging").(bool)),
		TargetResources:       expandStringArray(d.Get("target_resources")),
		TargetServiceAccounts: expandStringArray(d.Get("target_service_accounts")),
	}
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
	client := NewDCLComputeClient(config, userAgent, billingProject)
	res, err := client.ApplyFirewallPolicyRule(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating FirewallPolicyRule: %s", err)
	}

	log.Printf("[DEBUG] Finished creating FirewallPolicyRule %q: %#v", d.Id(), res)

	return resourceComputeFirewallPolicyRuleRead(d, meta)
}

func resourceComputeFirewallPolicyRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	obj := &compute.FirewallPolicyRule{
		Action:                dcl.String(d.Get("action").(string)),
		Direction:             compute.FirewallPolicyRuleDirectionEnumRef(d.Get("direction").(string)),
		FirewallPolicy:        dcl.String(d.Get("firewall_policy").(string)),
		Match:                 expandComputeFirewallPolicyRuleMatch(d.Get("match")),
		Priority:              dcl.Int64(int64(d.Get("priority").(int))),
		Description:           dcl.String(d.Get("description").(string)),
		Disabled:              dcl.Bool(d.Get("disabled").(bool)),
		EnableLogging:         dcl.Bool(d.Get("enable_logging").(bool)),
		TargetResources:       expandStringArray(d.Get("target_resources")),
		TargetServiceAccounts: expandStringArray(d.Get("target_service_accounts")),
	}

	log.Printf("[DEBUG] Deleting FirewallPolicyRule %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := ""
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLComputeClient(config, userAgent, billingProject)
	if err := client.DeleteFirewallPolicyRule(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting FirewallPolicyRule: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting FirewallPolicyRule %q", d.Id())
	return nil
}

func resourceComputeFirewallPolicyRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"locations/global/firewallPolicies/(?P<firewall_policy>[^/]+)/rules/(?P<priority>[^/]+)",
		"(?P<firewall_policy>[^/]+)/(?P<priority>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "locations/global/firewallPolicies/{{firewall_policy}}/rules/{{priority}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandComputeFirewallPolicyRuleMatch(o interface{}) *compute.FirewallPolicyRuleMatch {
	if o == nil {
		return compute.EmptyFirewallPolicyRuleMatch
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return compute.EmptyFirewallPolicyRuleMatch
	}
	obj := objArr[0].(map[string]interface{})
	return &compute.FirewallPolicyRuleMatch{
		Layer4Configs: expandComputeFirewallPolicyRuleMatchLayer4ConfigsArray(obj["layer4_configs"]),
		DestIPRanges:  expandStringArray(obj["dest_ip_ranges"]),
		SrcIPRanges:   expandStringArray(obj["src_ip_ranges"]),
	}
}

func flattenComputeFirewallPolicyRuleMatch(obj *compute.FirewallPolicyRuleMatch) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"layer4_configs": flattenComputeFirewallPolicyRuleMatchLayer4ConfigsArray(obj.Layer4Configs),
		"dest_ip_ranges": obj.DestIPRanges,
		"src_ip_ranges":  obj.SrcIPRanges,
	}

	return []interface{}{transformed}

}
func expandComputeFirewallPolicyRuleMatchLayer4ConfigsArray(o interface{}) []compute.FirewallPolicyRuleMatchLayer4Configs {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]compute.FirewallPolicyRuleMatchLayer4Configs, 0, len(objs))
	for _, item := range objs {
		i := expandComputeFirewallPolicyRuleMatchLayer4Configs(item)
		items = append(items, *i)
	}

	return items
}

func expandComputeFirewallPolicyRuleMatchLayer4Configs(o interface{}) *compute.FirewallPolicyRuleMatchLayer4Configs {
	if o == nil {
		return compute.EmptyFirewallPolicyRuleMatchLayer4Configs
	}

	obj := o.(map[string]interface{})
	return &compute.FirewallPolicyRuleMatchLayer4Configs{
		IPProtocol: dcl.String(obj["ip_protocol"].(string)),
		Ports:      expandStringArray(obj["ports"]),
	}
}

func flattenComputeFirewallPolicyRuleMatchLayer4ConfigsArray(objs []compute.FirewallPolicyRuleMatchLayer4Configs) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenComputeFirewallPolicyRuleMatchLayer4Configs(&item)
		items = append(items, i)
	}

	return items
}

func flattenComputeFirewallPolicyRuleMatchLayer4Configs(obj *compute.FirewallPolicyRuleMatchLayer4Configs) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"ip_protocol": obj.IPProtocol,
		"ports":       obj.Ports,
	}

	return transformed

}
