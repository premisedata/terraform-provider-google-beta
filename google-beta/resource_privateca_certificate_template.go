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
	privateca "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/privateca/beta"
)

func resourcePrivatecaCertificateTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourcePrivatecaCertificateTemplateCreate,
		Read:   resourcePrivatecaCertificateTemplateRead,
		Update: resourcePrivatecaCertificateTemplateUpdate,
		Delete: resourcePrivatecaCertificateTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: resourcePrivatecaCertificateTemplateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"identity_constraints": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplateIdentityConstraintsSchema(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"passthrough_extensions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplatePassthroughExtensionsSchema(),
			},

			"predefined_values": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesSchema(),
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"create_time": {
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

func PrivatecaCertificateTemplateIdentityConstraintsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"allow_subject_alt_names_passthrough": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: ``,
			},

			"allow_subject_passthrough": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: ``,
			},

			"cel_expression": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplateIdentityConstraintsCelExpressionSchema(),
			},
		},
	}
}

func PrivatecaCertificateTemplateIdentityConstraintsCelExpressionSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"expression": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},

			"title": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func PrivatecaCertificateTemplatePassthroughExtensionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_extensions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsSchema(),
			},

			"known_extensions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func PrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_id_path": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"additional_extensions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsSchema(),
			},

			"aia_ocsp_servers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"ca_options": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesCaOptionsSchema(),
			},

			"key_usage": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesKeyUsageSchema(),
			},

			"policy_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesPolicyIdsSchema(),
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_id": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdSchema(),
			},

			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: ``,
			},

			"critical": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_id_path": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesCaOptionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"is_ca": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"max_issuer_path_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesKeyUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"base_key_usage": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageSchema(),
			},

			"extended_key_usage": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageSchema(),
			},

			"unknown_extended_key_usages": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: ``,
				Elem:        PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesSchema(),
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cert_sign": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"content_commitment": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"crl_sign": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"data_encipherment": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"decipher_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"digital_signature": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"encipher_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"key_agreement": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"key_encipherment": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"code_signing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"email_protection": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"ocsp_signing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"server_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},

			"time_stamping": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: ``,
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_id_path": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func PrivatecaCertificateTemplatePredefinedValuesPolicyIdsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"object_id_path": {
				Type:        schema.TypeList,
				Required:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeInt},
			},
		},
	}
}

func resourcePrivatecaCertificateTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &privateca.CertificateTemplate{
		Location:              dcl.String(d.Get("location").(string)),
		Name:                  dcl.String(d.Get("name").(string)),
		Description:           dcl.String(d.Get("description").(string)),
		IdentityConstraints:   expandPrivatecaCertificateTemplateIdentityConstraints(d.Get("identity_constraints")),
		Labels:                checkStringMap(d.Get("labels")),
		PassthroughExtensions: expandPrivatecaCertificateTemplatePassthroughExtensions(d.Get("passthrough_extensions")),
		PredefinedValues:      expandPrivatecaCertificateTemplatePredefinedValues(d.Get("predefined_values")),
		Project:               dcl.String(project),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}")
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
	client := NewDCLPrivatecaClient(config, userAgent, billingProject)
	res, err := client.ApplyCertificateTemplate(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating CertificateTemplate: %s", err)
	}

	log.Printf("[DEBUG] Finished creating CertificateTemplate %q: %#v", d.Id(), res)

	return resourcePrivatecaCertificateTemplateRead(d, meta)
}

func resourcePrivatecaCertificateTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &privateca.CertificateTemplate{
		Location:              dcl.String(d.Get("location").(string)),
		Name:                  dcl.String(d.Get("name").(string)),
		Description:           dcl.String(d.Get("description").(string)),
		IdentityConstraints:   expandPrivatecaCertificateTemplateIdentityConstraints(d.Get("identity_constraints")),
		Labels:                checkStringMap(d.Get("labels")),
		PassthroughExtensions: expandPrivatecaCertificateTemplatePassthroughExtensions(d.Get("passthrough_extensions")),
		PredefinedValues:      expandPrivatecaCertificateTemplatePredefinedValues(d.Get("predefined_values")),
		Project:               dcl.String(project),
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
	client := NewDCLPrivatecaClient(config, userAgent, billingProject)
	res, err := client.GetCertificateTemplate(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("PrivatecaCertificateTemplate %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("description", res.Description); err != nil {
		return fmt.Errorf("error setting description in state: %s", err)
	}
	if err = d.Set("identity_constraints", flattenPrivatecaCertificateTemplateIdentityConstraints(res.IdentityConstraints)); err != nil {
		return fmt.Errorf("error setting identity_constraints in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("passthrough_extensions", flattenPrivatecaCertificateTemplatePassthroughExtensions(res.PassthroughExtensions)); err != nil {
		return fmt.Errorf("error setting passthrough_extensions in state: %s", err)
	}
	if err = d.Set("predefined_values", flattenPrivatecaCertificateTemplatePredefinedValues(res.PredefinedValues)); err != nil {
		return fmt.Errorf("error setting predefined_values in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}
func resourcePrivatecaCertificateTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &privateca.CertificateTemplate{
		Location:              dcl.String(d.Get("location").(string)),
		Name:                  dcl.String(d.Get("name").(string)),
		Description:           dcl.String(d.Get("description").(string)),
		IdentityConstraints:   expandPrivatecaCertificateTemplateIdentityConstraints(d.Get("identity_constraints")),
		Labels:                checkStringMap(d.Get("labels")),
		PassthroughExtensions: expandPrivatecaCertificateTemplatePassthroughExtensions(d.Get("passthrough_extensions")),
		PredefinedValues:      expandPrivatecaCertificateTemplatePredefinedValues(d.Get("predefined_values")),
		Project:               dcl.String(project),
	}
	// Construct state hint from old values
	old := &privateca.CertificateTemplate{
		Location:              dcl.String(oldValue(d.GetChange("location")).(string)),
		Name:                  dcl.String(oldValue(d.GetChange("name")).(string)),
		Description:           dcl.String(oldValue(d.GetChange("description")).(string)),
		IdentityConstraints:   expandPrivatecaCertificateTemplateIdentityConstraints(oldValue(d.GetChange("identity_constraints"))),
		Labels:                checkStringMap(oldValue(d.GetChange("labels"))),
		PassthroughExtensions: expandPrivatecaCertificateTemplatePassthroughExtensions(oldValue(d.GetChange("passthrough_extensions"))),
		PredefinedValues:      expandPrivatecaCertificateTemplatePredefinedValues(oldValue(d.GetChange("predefined_values"))),
		Project:               dcl.StringOrNil(oldValue(d.GetChange("project")).(string)),
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
	client := NewDCLPrivatecaClient(config, userAgent, billingProject)
	res, err := client.ApplyCertificateTemplate(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error updating CertificateTemplate: %s", err)
	}

	log.Printf("[DEBUG] Finished creating CertificateTemplate %q: %#v", d.Id(), res)

	return resourcePrivatecaCertificateTemplateRead(d, meta)
}

func resourcePrivatecaCertificateTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &privateca.CertificateTemplate{
		Location:              dcl.String(d.Get("location").(string)),
		Name:                  dcl.String(d.Get("name").(string)),
		Description:           dcl.String(d.Get("description").(string)),
		IdentityConstraints:   expandPrivatecaCertificateTemplateIdentityConstraints(d.Get("identity_constraints")),
		Labels:                checkStringMap(d.Get("labels")),
		PassthroughExtensions: expandPrivatecaCertificateTemplatePassthroughExtensions(d.Get("passthrough_extensions")),
		PredefinedValues:      expandPrivatecaCertificateTemplatePredefinedValues(d.Get("predefined_values")),
		Project:               dcl.String(project),
	}

	log.Printf("[DEBUG] Deleting CertificateTemplate %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLPrivatecaClient(config, userAgent, billingProject)
	if err := client.DeleteCertificateTemplate(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting CertificateTemplate: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting CertificateTemplate %q", d.Id())
	return nil
}

func resourcePrivatecaCertificateTemplateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/certificateTemplates/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandPrivatecaCertificateTemplateIdentityConstraints(o interface{}) *privateca.CertificateTemplateIdentityConstraints {
	if o == nil {
		return privateca.EmptyCertificateTemplateIdentityConstraints
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplateIdentityConstraints
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplateIdentityConstraints{
		AllowSubjectAltNamesPassthrough: dcl.Bool(obj["allow_subject_alt_names_passthrough"].(bool)),
		AllowSubjectPassthrough:         dcl.Bool(obj["allow_subject_passthrough"].(bool)),
		CelExpression:                   expandPrivatecaCertificateTemplateIdentityConstraintsCelExpression(obj["cel_expression"]),
	}
}

func flattenPrivatecaCertificateTemplateIdentityConstraints(obj *privateca.CertificateTemplateIdentityConstraints) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"allow_subject_alt_names_passthrough": obj.AllowSubjectAltNamesPassthrough,
		"allow_subject_passthrough":           obj.AllowSubjectPassthrough,
		"cel_expression":                      flattenPrivatecaCertificateTemplateIdentityConstraintsCelExpression(obj.CelExpression),
	}

	return []interface{}{transformed}

}

func expandPrivatecaCertificateTemplateIdentityConstraintsCelExpression(o interface{}) *privateca.CertificateTemplateIdentityConstraintsCelExpression {
	if o == nil {
		return privateca.EmptyCertificateTemplateIdentityConstraintsCelExpression
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplateIdentityConstraintsCelExpression
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplateIdentityConstraintsCelExpression{
		Description: dcl.String(obj["description"].(string)),
		Expression:  dcl.String(obj["expression"].(string)),
		Location:    dcl.String(obj["location"].(string)),
		Title:       dcl.String(obj["title"].(string)),
	}
}

func flattenPrivatecaCertificateTemplateIdentityConstraintsCelExpression(obj *privateca.CertificateTemplateIdentityConstraintsCelExpression) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"description": obj.Description,
		"expression":  obj.Expression,
		"location":    obj.Location,
		"title":       obj.Title,
	}

	return []interface{}{transformed}

}

func expandPrivatecaCertificateTemplatePassthroughExtensions(o interface{}) *privateca.CertificateTemplatePassthroughExtensions {
	if o == nil {
		return privateca.EmptyCertificateTemplatePassthroughExtensions
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplatePassthroughExtensions
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplatePassthroughExtensions{
		AdditionalExtensions: expandPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsArray(obj["additional_extensions"]),
		KnownExtensions:      expandPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsArray(obj["known_extensions"]),
	}
}

func flattenPrivatecaCertificateTemplatePassthroughExtensions(obj *privateca.CertificateTemplatePassthroughExtensions) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"additional_extensions": flattenPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsArray(obj.AdditionalExtensions),
		"known_extensions":      flattenPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsArray(obj.KnownExtensions),
	}

	return []interface{}{transformed}

}
func expandPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsArray(o interface{}) []privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions, 0, len(objs))
	for _, item := range objs {
		i := expandPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(item)
		items = append(items, *i)
	}

	return items
}

func expandPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(o interface{}) *privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions {
	if o == nil {
		return privateca.EmptyCertificateTemplatePassthroughExtensionsAdditionalExtensions
	}

	obj := o.(map[string]interface{})
	return &privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions{
		ObjectIdPath: expandIntegerArray(obj["object_id_path"]),
	}
}

func flattenPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsArray(objs []privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(&item)
		items = append(items, i)
	}

	return items
}

func flattenPrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions(obj *privateca.CertificateTemplatePassthroughExtensionsAdditionalExtensions) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"object_id_path": obj.ObjectIdPath,
	}

	return transformed

}

func expandPrivatecaCertificateTemplatePredefinedValues(o interface{}) *privateca.CertificateTemplatePredefinedValues {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValues
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplatePredefinedValues
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValues{
		AdditionalExtensions: expandPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsArray(obj["additional_extensions"]),
		AiaOcspServers:       expandStringArray(obj["aia_ocsp_servers"]),
		CaOptions:            expandPrivatecaCertificateTemplatePredefinedValuesCaOptions(obj["ca_options"]),
		KeyUsage:             expandPrivatecaCertificateTemplatePredefinedValuesKeyUsage(obj["key_usage"]),
		PolicyIds:            expandPrivatecaCertificateTemplatePredefinedValuesPolicyIdsArray(obj["policy_ids"]),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValues(obj *privateca.CertificateTemplatePredefinedValues) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"additional_extensions": flattenPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsArray(obj.AdditionalExtensions),
		"aia_ocsp_servers":      obj.AiaOcspServers,
		"ca_options":            flattenPrivatecaCertificateTemplatePredefinedValuesCaOptions(obj.CaOptions),
		"key_usage":             flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsage(obj.KeyUsage),
		"policy_ids":            flattenPrivatecaCertificateTemplatePredefinedValuesPolicyIdsArray(obj.PolicyIds),
	}

	return []interface{}{transformed}

}
func expandPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsArray(o interface{}) []privateca.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]privateca.CertificateTemplatePredefinedValuesAdditionalExtensions, 0, len(objs))
	for _, item := range objs {
		i := expandPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(item)
		items = append(items, *i)
	}

	return items
}

func expandPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(o interface{}) *privateca.CertificateTemplatePredefinedValuesAdditionalExtensions {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesAdditionalExtensions
	}

	obj := o.(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesAdditionalExtensions{
		ObjectId: expandPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(obj["object_id"]),
		Value:    dcl.String(obj["value"].(string)),
		Critical: dcl.Bool(obj["critical"].(bool)),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsArray(objs []privateca.CertificateTemplatePredefinedValuesAdditionalExtensions) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(&item)
		items = append(items, i)
	}

	return items
}

func flattenPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions(obj *privateca.CertificateTemplatePredefinedValuesAdditionalExtensions) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"object_id": flattenPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(obj.ObjectId),
		"value":     obj.Value,
		"critical":  obj.Critical,
	}

	return transformed

}

func expandPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(o interface{}) *privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId{
		ObjectIdPath: expandIntegerArray(obj["object_id_path"]),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId(obj *privateca.CertificateTemplatePredefinedValuesAdditionalExtensionsObjectId) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"object_id_path": obj.ObjectIdPath,
	}

	return []interface{}{transformed}

}

func expandPrivatecaCertificateTemplatePredefinedValuesCaOptions(o interface{}) *privateca.CertificateTemplatePredefinedValuesCaOptions {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesCaOptions
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplatePredefinedValuesCaOptions
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesCaOptions{
		IsCa:                dcl.Bool(obj["is_ca"].(bool)),
		MaxIssuerPathLength: dcl.Int64(int64(obj["max_issuer_path_length"].(int))),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesCaOptions(obj *privateca.CertificateTemplatePredefinedValuesCaOptions) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"is_ca":                  obj.IsCa,
		"max_issuer_path_length": obj.MaxIssuerPathLength,
	}

	return []interface{}{transformed}

}

func expandPrivatecaCertificateTemplatePredefinedValuesKeyUsage(o interface{}) *privateca.CertificateTemplatePredefinedValuesKeyUsage {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesKeyUsage
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplatePredefinedValuesKeyUsage
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesKeyUsage{
		BaseKeyUsage:             expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(obj["base_key_usage"]),
		ExtendedKeyUsage:         expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(obj["extended_key_usage"]),
		UnknownExtendedKeyUsages: expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray(obj["unknown_extended_key_usages"]),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsage(obj *privateca.CertificateTemplatePredefinedValuesKeyUsage) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"base_key_usage":              flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(obj.BaseKeyUsage),
		"extended_key_usage":          flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(obj.ExtendedKeyUsage),
		"unknown_extended_key_usages": flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray(obj.UnknownExtendedKeyUsages),
	}

	return []interface{}{transformed}

}

func expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(o interface{}) *privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage{
		CertSign:          dcl.Bool(obj["cert_sign"].(bool)),
		ContentCommitment: dcl.Bool(obj["content_commitment"].(bool)),
		CrlSign:           dcl.Bool(obj["crl_sign"].(bool)),
		DataEncipherment:  dcl.Bool(obj["data_encipherment"].(bool)),
		DecipherOnly:      dcl.Bool(obj["decipher_only"].(bool)),
		DigitalSignature:  dcl.Bool(obj["digital_signature"].(bool)),
		EncipherOnly:      dcl.Bool(obj["encipher_only"].(bool)),
		KeyAgreement:      dcl.Bool(obj["key_agreement"].(bool)),
		KeyEncipherment:   dcl.Bool(obj["key_encipherment"].(bool)),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage(obj *privateca.CertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cert_sign":          obj.CertSign,
		"content_commitment": obj.ContentCommitment,
		"crl_sign":           obj.CrlSign,
		"data_encipherment":  obj.DataEncipherment,
		"decipher_only":      obj.DecipherOnly,
		"digital_signature":  obj.DigitalSignature,
		"encipher_only":      obj.EncipherOnly,
		"key_agreement":      obj.KeyAgreement,
		"key_encipherment":   obj.KeyEncipherment,
	}

	return []interface{}{transformed}

}

func expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(o interface{}) *privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return privateca.EmptyCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage
	}
	obj := objArr[0].(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage{
		ClientAuth:      dcl.Bool(obj["client_auth"].(bool)),
		CodeSigning:     dcl.Bool(obj["code_signing"].(bool)),
		EmailProtection: dcl.Bool(obj["email_protection"].(bool)),
		OcspSigning:     dcl.Bool(obj["ocsp_signing"].(bool)),
		ServerAuth:      dcl.Bool(obj["server_auth"].(bool)),
		TimeStamping:    dcl.Bool(obj["time_stamping"].(bool)),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage(obj *privateca.CertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"client_auth":      obj.ClientAuth,
		"code_signing":     obj.CodeSigning,
		"email_protection": obj.EmailProtection,
		"ocsp_signing":     obj.OcspSigning,
		"server_auth":      obj.ServerAuth,
		"time_stamping":    obj.TimeStamping,
	}

	return []interface{}{transformed}

}
func expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray(o interface{}) []privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages, 0, len(objs))
	for _, item := range objs {
		i := expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(item)
		items = append(items, *i)
	}

	return items
}

func expandPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(o interface{}) *privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages
	}

	obj := o.(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages{
		ObjectIdPath: expandIntegerArray(obj["object_id_path"]),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesArray(objs []privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(&item)
		items = append(items, i)
	}

	return items
}

func flattenPrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages(obj *privateca.CertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"object_id_path": obj.ObjectIdPath,
	}

	return transformed

}
func expandPrivatecaCertificateTemplatePredefinedValuesPolicyIdsArray(o interface{}) []privateca.CertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]privateca.CertificateTemplatePredefinedValuesPolicyIds, 0, len(objs))
	for _, item := range objs {
		i := expandPrivatecaCertificateTemplatePredefinedValuesPolicyIds(item)
		items = append(items, *i)
	}

	return items
}

func expandPrivatecaCertificateTemplatePredefinedValuesPolicyIds(o interface{}) *privateca.CertificateTemplatePredefinedValuesPolicyIds {
	if o == nil {
		return privateca.EmptyCertificateTemplatePredefinedValuesPolicyIds
	}

	obj := o.(map[string]interface{})
	return &privateca.CertificateTemplatePredefinedValuesPolicyIds{
		ObjectIdPath: expandIntegerArray(obj["object_id_path"]),
	}
}

func flattenPrivatecaCertificateTemplatePredefinedValuesPolicyIdsArray(objs []privateca.CertificateTemplatePredefinedValuesPolicyIds) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenPrivatecaCertificateTemplatePredefinedValuesPolicyIds(&item)
		items = append(items, i)
	}

	return items
}

func flattenPrivatecaCertificateTemplatePredefinedValuesPolicyIds(obj *privateca.CertificateTemplatePredefinedValuesPolicyIds) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"object_id_path": obj.ObjectIdPath,
	}

	return transformed

}
func flattenPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsArray(obj []privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum) interface{} {
	if obj == nil {
		return nil
	}
	items := []string{}
	for _, item := range obj {
		items = append(items, string(item))
	}
	return items
}

func expandPrivatecaCertificateTemplatePassthroughExtensionsKnownExtensionsArray(o interface{}) []privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum {
	objs := o.([]interface{})
	items := make([]privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnum, 0, len(objs))
	for _, item := range objs {
		i := privateca.CertificateTemplatePassthroughExtensionsKnownExtensionsEnumRef(item.(string))
		items = append(items, *i)
	}
	return items
}
