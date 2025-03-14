// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceComputeGlobalForwardingRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeGlobalForwardingRuleCreate,
		Read:   resourceComputeGlobalForwardingRuleRead,
		Update: resourceComputeGlobalForwardingRuleUpdate,
		Delete: resourceComputeGlobalForwardingRuleDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeGlobalForwardingRuleImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(4 * time.Minute),
			Update: schema.DefaultTimeout(4 * time.Minute),
			Delete: schema.DefaultTimeout(4 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource; provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"target": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: compareSelfLinkRelativePaths,
				Description: `The URL of the target resource to receive the matched traffic.
The forwarded traffic must be of a type appropriate to the target object.
For INTERNAL_SELF_MANAGED load balancing, only HTTP and HTTPS targets
are valid.

([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html) only) For global address with a purpose of PRIVATE_SERVICE_CONNECT and
addressType of INTERNAL, only "all-apis" and "vpc-sc" are valid.`,
			},
			"ip_address": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: internalIpDiffSuppress,
				Description: `The IP address that this forwarding rule serves. When a client sends
traffic to this IP address, the forwarding rule directs the traffic to
the target that you specify in the forwarding rule. The
loadBalancingScheme and the forwarding rule's target determine the
type of IP address that you can use. For detailed information, refer
to [IP address specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#ip_address_specifications).

An address can be specified either by a literal IP address or a
reference to an existing Address resource. If you don't specify a
reserved IP address, an ephemeral IP address is assigned.

The value must be set to 0.0.0.0 when the target is a targetGrpcProxy
that has validateForProxyless field set to true.

For Private Service Connect forwarding rules that forward traffic to
Google APIs, IP address must be provided.`,
			},
			"ip_protocol": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"TCP", "UDP", "ESP", "AH", "SCTP", "ICMP", ""}, false),
				DiffSuppressFunc: caseDiffSuppress,
				Description: `The IP protocol to which this rule applies. When the load balancing scheme is
INTERNAL_SELF_MANAGED, only TCP is valid. This field must not be set if the
global address is configured as a purpose of PRIVATE_SERVICE_CONNECT
and addressType of INTERNAL Possible values: ["TCP", "UDP", "ESP", "AH", "SCTP", "ICMP"]`,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `An optional description of this resource. Provide this property when
you create the resource.`,
			},
			"ip_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"IPV4", "IPV6", ""}, false),
				Description:  `The IP Version that will be used by this global forwarding rule. Possible values: ["IPV4", "IPV6"]`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Labels to apply to this forwarding rule.  A list of key->value pairs.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"load_balancing_scheme": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"EXTERNAL", "INTERNAL_SELF_MANAGED", ""}, false),
				Description: `This signifies what the GlobalForwardingRule will be used for.
The value of INTERNAL_SELF_MANAGED means that this will be used for
Internal Global HTTP(S) LB. The value of EXTERNAL means that this
will be used for External Global Load Balancing (HTTP(S) LB,
External TCP/UDP LB, SSL Proxy)

([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html) only) Note: This field must be set "" if the global address is
configured as a purpose of PRIVATE_SERVICE_CONNECT and addressType of INTERNAL. Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL_SELF_MANAGED"]`,
				Default: "EXTERNAL",
			},
			"metadata_filters": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Opaque filter criteria used by Loadbalancer to restrict routing
configuration to a limited set xDS compliant clients. In their xDS
requests to Loadbalancer, xDS clients present node metadata. If a
match takes place, the relevant routing configuration is made available
to those proxies.

For each metadataFilter in this list, if its filterMatchCriteria is set
to MATCH_ANY, at least one of the filterLabels must match the
corresponding label provided in the metadata. If its filterMatchCriteria
is set to MATCH_ALL, then all of its filterLabels must match with
corresponding labels in the provided metadata.

metadataFilters specified here can be overridden by those specified in
the UrlMap that this ForwardingRule references.

metadataFilters only applies to Loadbalancers that have their
loadBalancingScheme set to INTERNAL_SELF_MANAGED.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter_labels": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Description: `The list of label value pairs that must match labels in the
provided metadata based on filterMatchCriteria

This list must not be empty and can have at the most 64 entries.`,
							MinItems: 1,
							MaxItems: 64,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: `Name of the metadata label. The length must be between
1 and 1024 characters, inclusive.`,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
										Description: `The value that the label must match. The value has a maximum
length of 1024 characters.`,
									},
								},
							},
						},
						"filter_match_criteria": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"MATCH_ANY", "MATCH_ALL"}, false),
							Description: `Specifies how individual filterLabel matches within the list of
filterLabels contribute towards the overall metadataFilter match.

MATCH_ANY - At least one of the filterLabels must have a matching
label in the provided metadata.
MATCH_ALL - All filterLabels must have matching labels in the
provided metadata. Possible values: ["MATCH_ANY", "MATCH_ALL"]`,
						},
					},
				},
			},
			"network": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `This field is not used for external load balancing.
For INTERNAL_SELF_MANAGED load balancing, this field
identifies the network that the load balanced IP should belong to
for this global forwarding rule. If this field is not specified,
the default network will be used.`,
			},
			"port_range": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: portRangeDiffSuppress,
				Description: `This field is used along with the target field for TargetHttpProxy,
TargetHttpsProxy, TargetSslProxy, TargetTcpProxy, TargetVpnGateway,
TargetPool, TargetInstance.

Applicable only when IPProtocol is TCP, UDP, or SCTP, only packets
addressed to ports in the specified range will be forwarded to target.
Forwarding rules with the same [IPAddress, IPProtocol] pair must have
disjoint port ranges.

Some types of forwarding target have constraints on the acceptable
ports:

* TargetHttpProxy: 80, 8080
* TargetHttpsProxy: 443
* TargetTcpProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
                  1883, 5222
* TargetSslProxy: 25, 43, 110, 143, 195, 443, 465, 587, 700, 993, 995,
                  1883, 5222
* TargetVpnGateway: 500, 4500`,
			},
			"label_fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The fingerprint used for optimistic locking of this resource.  Used
internally during updates.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeGlobalForwardingRuleCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeGlobalForwardingRuleDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	IPAddressProp, err := expandComputeGlobalForwardingRuleIPAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(IPAddressProp)) && (ok || !reflect.DeepEqual(v, IPAddressProp)) {
		obj["IPAddress"] = IPAddressProp
	}
	IPProtocolProp, err := expandComputeGlobalForwardingRuleIPProtocol(d.Get("ip_protocol"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_protocol"); !isEmptyValue(reflect.ValueOf(IPProtocolProp)) && (ok || !reflect.DeepEqual(v, IPProtocolProp)) {
		obj["IPProtocol"] = IPProtocolProp
	}
	ipVersionProp, err := expandComputeGlobalForwardingRuleIpVersion(d.Get("ip_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_version"); !isEmptyValue(reflect.ValueOf(ipVersionProp)) && (ok || !reflect.DeepEqual(v, ipVersionProp)) {
		obj["ipVersion"] = ipVersionProp
	}
	labelsProp, err := expandComputeGlobalForwardingRuleLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	labelFingerprintProp, err := expandComputeGlobalForwardingRuleLabelFingerprint(d.Get("label_fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(labelFingerprintProp)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
		obj["labelFingerprint"] = labelFingerprintProp
	}
	loadBalancingSchemeProp, err := expandComputeGlobalForwardingRuleLoadBalancingScheme(d.Get("load_balancing_scheme"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("load_balancing_scheme"); !isEmptyValue(reflect.ValueOf(loadBalancingSchemeProp)) && (ok || !reflect.DeepEqual(v, loadBalancingSchemeProp)) {
		obj["loadBalancingScheme"] = loadBalancingSchemeProp
	}
	metadataFiltersProp, err := expandComputeGlobalForwardingRuleMetadataFilters(d.Get("metadata_filters"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata_filters"); !isEmptyValue(reflect.ValueOf(metadataFiltersProp)) && (ok || !reflect.DeepEqual(v, metadataFiltersProp)) {
		obj["metadataFilters"] = metadataFiltersProp
	}
	nameProp, err := expandComputeGlobalForwardingRuleName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networkProp, err := expandComputeGlobalForwardingRuleNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	portRangeProp, err := expandComputeGlobalForwardingRulePortRange(d.Get("port_range"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port_range"); !isEmptyValue(reflect.ValueOf(portRangeProp)) && (ok || !reflect.DeepEqual(v, portRangeProp)) {
		obj["portRange"] = portRangeProp
	}
	targetProp, err := expandComputeGlobalForwardingRuleTarget(d.Get("target"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(targetProp)) && (ok || !reflect.DeepEqual(v, targetProp)) {
		obj["target"] = targetProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GlobalForwardingRule: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalForwardingRule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GlobalForwardingRule: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating GlobalForwardingRule", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GlobalForwardingRule: %s", err)
	}

	if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		// Labels cannot be set in a create.  We'll have to set them here.
		err = resourceComputeGlobalForwardingRuleRead(d, meta)
		if err != nil {
			return err
		}

		obj := make(map[string]interface{})
		// d.Get("labels") will have been overridden by the Read call.
		labelsProp, err := expandComputeGlobalForwardingRuleLabels(v, d, config)
		if err != nil {
			return err
		}
		obj["labels"] = labelsProp
		labelFingerprintProp := d.Get("label_fingerprint")
		obj["labelFingerprint"] = labelFingerprintProp

		url, err = replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}/setLabels")
		if err != nil {
			return err
		}
		res, err = sendRequest(config, "POST", project, url, userAgent, obj)
		if err != nil {
			return fmt.Errorf("Error adding labels to ComputeGlobalForwardingRule %q: %s", d.Id(), err)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating ComputeGlobalForwardingRule Labels", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}

	}

	log.Printf("[DEBUG] Finished creating GlobalForwardingRule %q: %#v", d.Id(), res)

	return resourceComputeGlobalForwardingRuleRead(d, meta)
}

func resourceComputeGlobalForwardingRuleRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalForwardingRule: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeGlobalForwardingRule %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}

	if err := d.Set("description", flattenComputeGlobalForwardingRuleDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("ip_address", flattenComputeGlobalForwardingRuleIPAddress(res["IPAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("ip_protocol", flattenComputeGlobalForwardingRuleIPProtocol(res["IPProtocol"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("ip_version", flattenComputeGlobalForwardingRuleIpVersion(res["ipVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("labels", flattenComputeGlobalForwardingRuleLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("label_fingerprint", flattenComputeGlobalForwardingRuleLabelFingerprint(res["labelFingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("load_balancing_scheme", flattenComputeGlobalForwardingRuleLoadBalancingScheme(res["loadBalancingScheme"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("metadata_filters", flattenComputeGlobalForwardingRuleMetadataFilters(res["metadataFilters"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("name", flattenComputeGlobalForwardingRuleName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("network", flattenComputeGlobalForwardingRuleNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("port_range", flattenComputeGlobalForwardingRulePortRange(res["portRange"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("target", flattenComputeGlobalForwardingRuleTarget(res["target"], d, config)); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading GlobalForwardingRule: %s", err)
	}

	return nil
}

func resourceComputeGlobalForwardingRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalForwardingRule: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("labels") || d.HasChange("label_fingerprint") {
		obj := make(map[string]interface{})

		labelsProp, err := expandComputeGlobalForwardingRuleLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}
		labelFingerprintProp, err := expandComputeGlobalForwardingRuleLabelFingerprint(d.Get("label_fingerprint"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("label_fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelFingerprintProp)) {
			obj["labelFingerprint"] = labelFingerprintProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}/setLabels")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating GlobalForwardingRule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating GlobalForwardingRule %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating GlobalForwardingRule", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}
	if d.HasChange("target") {
		obj := make(map[string]interface{})

		targetProp, err := expandComputeGlobalForwardingRuleTarget(d.Get("target"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("target"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, targetProp)) {
			obj["target"] = targetProp
		}

		url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}/setTarget")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating GlobalForwardingRule %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating GlobalForwardingRule %q: %#v", d.Id(), res)
		}

		err = computeOperationWaitTime(
			config, res, project, "Updating GlobalForwardingRule", userAgent,
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceComputeGlobalForwardingRuleRead(d, meta)
}

func resourceComputeGlobalForwardingRuleDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GlobalForwardingRule: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GlobalForwardingRule %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GlobalForwardingRule")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting GlobalForwardingRule", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GlobalForwardingRule %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeGlobalForwardingRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/global/forwardingRules/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/global/forwardingRules/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeGlobalForwardingRuleDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIPAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIPProtocol(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleIpVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleLabelFingerprint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleMetadataFilters(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"filter_match_criteria": flattenComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(original["filterMatchCriteria"], d, config),
			"filter_labels":         flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabels(original["filterLabels"], d, config),
		})
	}
	return transformed
}
func flattenComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":  flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(original["name"], d, config),
			"value": flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(original["value"], d, config),
		})
	}
	return transformed
}
func flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenComputeGlobalForwardingRulePortRange(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeGlobalForwardingRuleTarget(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandComputeGlobalForwardingRuleDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIPProtocol(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleIpVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandComputeGlobalForwardingRuleLabelFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleLoadBalancingScheme(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFilters(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFilterMatchCriteria, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(original["filter_match_criteria"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilterMatchCriteria); val.IsValid() && !isEmptyValue(val) {
			transformed["filterMatchCriteria"] = transformedFilterMatchCriteria
		}

		transformedFilterLabels, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabels(original["filter_labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFilterLabels); val.IsValid() && !isEmptyValue(val) {
			transformed["filterLabels"] = transformedFilterLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterMatchCriteria(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabels(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedValue, err := expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(original["value"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedValue); val.IsValid() && !isEmptyValue(val) {
			transformed["value"] = transformedValue
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleMetadataFiltersFilterLabelsValue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("networks", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for network: %s", err)
	}
	return f.RelativeLink(), nil
}

func expandComputeGlobalForwardingRulePortRange(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeGlobalForwardingRuleTarget(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
