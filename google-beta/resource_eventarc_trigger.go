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
	eventarc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/eventarc/beta"
)

func resourceEventarcTrigger() *schema.Resource {
	return &schema.Resource{
		Create: resourceEventarcTriggerCreate,
		Read:   resourceEventarcTriggerRead,
		Update: resourceEventarcTriggerUpdate,
		Delete: resourceEventarcTriggerDelete,

		Importer: &schema.ResourceImporter{
			State: resourceEventarcTriggerImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"destination": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        EventarcTriggerDestinationSchema(),
			},

			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"matching_criteria": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: ``,
				Elem:        EventarcTriggerMatchingCriteriaSchema(),
				Set:         schema.HashResource(EventarcTriggerMatchingCriteriaSchema()),
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: ``,
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"service_account": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"transport": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        EventarcTriggerTransportSchema(),
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func EventarcTriggerDestinationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cloud_run_service": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        EventarcTriggerDestinationCloudRunServiceSchema(),
			},
		},
	}
}

func EventarcTriggerDestinationCloudRunServiceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"service": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"region": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func EventarcTriggerMatchingCriteriaSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"attribute": {
				Type:        schema.TypeString,
				Required:    true,
				Description: ``,
			},

			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: ``,
			},
		},
	}
}

func EventarcTriggerTransportSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pubsub": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        EventarcTriggerTransportPubsubSchema(),
			},
		},
	}
}

func EventarcTriggerTransportPubsubSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"topic": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"subscription": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func resourceEventarcTriggerCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &eventarc.Trigger{
		Destination:      expandEventarcTriggerDestination(d.Get("destination")),
		Location:         dcl.String(d.Get("location").(string)),
		MatchingCriteria: expandEventarcTriggerMatchingCriteriaArray(d.Get("matching_criteria")),
		Name:             dcl.String(d.Get("name").(string)),
		Labels:           checkStringMap(d.Get("labels")),
		Project:          dcl.String(project),
		ServiceAccount:   dcl.String(d.Get("service_account").(string)),
		Transport:        expandEventarcTriggerTransport(d.Get("transport")),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/triggers/{{name}}")
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
	client := NewDCLEventarcClient(config, userAgent, billingProject)
	res, err := client.ApplyTrigger(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Trigger: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Trigger %q: %#v", d.Id(), res)

	return resourceEventarcTriggerRead(d, meta)
}

func resourceEventarcTriggerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &eventarc.Trigger{
		Destination:      expandEventarcTriggerDestination(d.Get("destination")),
		Location:         dcl.String(d.Get("location").(string)),
		MatchingCriteria: expandEventarcTriggerMatchingCriteriaArray(d.Get("matching_criteria")),
		Name:             dcl.String(d.Get("name").(string)),
		Labels:           checkStringMap(d.Get("labels")),
		Project:          dcl.String(project),
		ServiceAccount:   dcl.String(d.Get("service_account").(string)),
		Transport:        expandEventarcTriggerTransport(d.Get("transport")),
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
	client := NewDCLEventarcClient(config, userAgent, billingProject)
	res, err := client.GetTrigger(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("EventarcTrigger %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("destination", flattenEventarcTriggerDestination(res.Destination)); err != nil {
		return fmt.Errorf("error setting destination in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("matching_criteria", flattenEventarcTriggerMatchingCriteriaArray(res.MatchingCriteria)); err != nil {
		return fmt.Errorf("error setting matching_criteria in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("service_account", res.ServiceAccount); err != nil {
		return fmt.Errorf("error setting service_account in state: %s", err)
	}
	if err = d.Set("transport", flattenEventarcTriggerTransport(res.Transport)); err != nil {
		return fmt.Errorf("error setting transport in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("etag", res.Etag); err != nil {
		return fmt.Errorf("error setting etag in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourceEventarcTriggerUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &eventarc.Trigger{
		Destination:      expandEventarcTriggerDestination(d.Get("destination")),
		Location:         dcl.String(d.Get("location").(string)),
		MatchingCriteria: expandEventarcTriggerMatchingCriteriaArray(d.Get("matching_criteria")),
		Name:             dcl.String(d.Get("name").(string)),
		Labels:           checkStringMap(d.Get("labels")),
		Project:          dcl.String(project),
		ServiceAccount:   dcl.String(d.Get("service_account").(string)),
		Transport:        expandEventarcTriggerTransport(d.Get("transport")),
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
	client := NewDCLEventarcClient(config, userAgent, billingProject)
	res, err := client.ApplyTrigger(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating Trigger: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Trigger %q: %#v", d.Id(), res)

	return resourceEventarcTriggerRead(d, meta)
}

func resourceEventarcTriggerDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &eventarc.Trigger{
		Destination:      expandEventarcTriggerDestination(d.Get("destination")),
		Location:         dcl.String(d.Get("location").(string)),
		MatchingCriteria: expandEventarcTriggerMatchingCriteriaArray(d.Get("matching_criteria")),
		Name:             dcl.String(d.Get("name").(string)),
		Labels:           checkStringMap(d.Get("labels")),
		Project:          dcl.String(project),
		ServiceAccount:   dcl.String(d.Get("service_account").(string)),
		Transport:        expandEventarcTriggerTransport(d.Get("transport")),
	}

	log.Printf("[DEBUG] Deleting Trigger %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLEventarcClient(config, userAgent, billingProject)
	if err := client.DeleteTrigger(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Trigger: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Trigger %q", d.Id())
	return nil
}

func resourceEventarcTriggerImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/triggers/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/triggers/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandEventarcTriggerDestination(o interface{}) *eventarc.TriggerDestination {
	if o == nil {
		return eventarc.EmptyTriggerDestination
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return eventarc.EmptyTriggerDestination
	}
	obj := objArr[0].(map[string]interface{})
	return &eventarc.TriggerDestination{
		CloudRunService: expandEventarcTriggerDestinationCloudRunService(obj["cloud_run_service"]),
	}
}

func flattenEventarcTriggerDestination(obj *eventarc.TriggerDestination) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cloud_run_service": flattenEventarcTriggerDestinationCloudRunService(obj.CloudRunService),
	}

	return []interface{}{transformed}

}

func expandEventarcTriggerDestinationCloudRunService(o interface{}) *eventarc.TriggerDestinationCloudRunService {
	if o == nil {
		return eventarc.EmptyTriggerDestinationCloudRunService
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return eventarc.EmptyTriggerDestinationCloudRunService
	}
	obj := objArr[0].(map[string]interface{})
	return &eventarc.TriggerDestinationCloudRunService{
		Service: dcl.String(obj["service"].(string)),
		Path:    dcl.String(obj["path"].(string)),
		Region:  dcl.StringOrNil(obj["region"].(string)),
	}
}

func flattenEventarcTriggerDestinationCloudRunService(obj *eventarc.TriggerDestinationCloudRunService) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"service": obj.Service,
		"path":    obj.Path,
		"region":  obj.Region,
	}

	return []interface{}{transformed}

}
func expandEventarcTriggerMatchingCriteriaArray(o interface{}) []eventarc.TriggerMatchingCriteria {
	if o == nil {
		return nil
	}

	o = o.(*schema.Set).List()

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]eventarc.TriggerMatchingCriteria, 0, len(objs))
	for _, item := range objs {
		i := expandEventarcTriggerMatchingCriteria(item)
		items = append(items, *i)
	}

	return items
}

func expandEventarcTriggerMatchingCriteria(o interface{}) *eventarc.TriggerMatchingCriteria {
	if o == nil {
		return eventarc.EmptyTriggerMatchingCriteria
	}

	obj := o.(map[string]interface{})
	return &eventarc.TriggerMatchingCriteria{
		Attribute: dcl.String(obj["attribute"].(string)),
		Value:     dcl.String(obj["value"].(string)),
	}
}

func flattenEventarcTriggerMatchingCriteriaArray(objs []eventarc.TriggerMatchingCriteria) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenEventarcTriggerMatchingCriteria(&item)
		items = append(items, i)
	}

	return items
}

func flattenEventarcTriggerMatchingCriteria(obj *eventarc.TriggerMatchingCriteria) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"attribute": obj.Attribute,
		"value":     obj.Value,
	}

	return transformed

}

func expandEventarcTriggerTransport(o interface{}) *eventarc.TriggerTransport {
	if o == nil {
		return nil
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return nil
	}
	obj := objArr[0].(map[string]interface{})
	return &eventarc.TriggerTransport{
		Pubsub: expandEventarcTriggerTransportPubsub(obj["pubsub"]),
	}
}

func flattenEventarcTriggerTransport(obj *eventarc.TriggerTransport) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"pubsub": flattenEventarcTriggerTransportPubsub(obj.Pubsub),
	}

	return []interface{}{transformed}

}

func expandEventarcTriggerTransportPubsub(o interface{}) *eventarc.TriggerTransportPubsub {
	if o == nil {
		return eventarc.EmptyTriggerTransportPubsub
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return eventarc.EmptyTriggerTransportPubsub
	}
	obj := objArr[0].(map[string]interface{})
	return &eventarc.TriggerTransportPubsub{
		Topic: dcl.String(obj["topic"].(string)),
	}
}

func flattenEventarcTriggerTransportPubsub(obj *eventarc.TriggerTransportPubsub) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"topic":        obj.Topic,
		"subscription": obj.Subscription,
	}

	return []interface{}{transformed}

}
