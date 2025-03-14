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
	dataproc "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/dataproc/beta"
)

func resourceDataprocWorkflowTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDataprocWorkflowTemplateCreate,
		Read:   resourceDataprocWorkflowTemplateRead,
		Delete: resourceDataprocWorkflowTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDataprocWorkflowTemplateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"jobs": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        DataprocWorkflowTemplateJobsSchema(),
			},

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

			"placement": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementSchema(),
			},

			"dag_timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        DataprocWorkflowTemplateParametersSchema(),
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"version": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
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

func DataprocWorkflowTemplateJobsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"step_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"hadoop_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHadoopJobSchema(),
			},

			"hive_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHiveJobSchema(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"pig_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPigJobSchema(),
			},

			"prerequisite_step_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"presto_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPrestoJobSchema(),
			},

			"pyspark_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPysparkJobSchema(),
			},

			"scheduling": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSchedulingSchema(),
			},

			"spark_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkJobSchema(),
			},

			"spark_r_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkRJobSchema(),
			},

			"spark_sql_job": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkSqlJobSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHadoopJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHadoopJobLoggingConfigSchema(),
			},

			"main_class": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"main_jar_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHadoopJobLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHiveJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"continue_on_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsHiveJobQueryListSchema(),
			},

			"script_variables": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsHiveJobQueryListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPigJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"continue_on_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPigJobLoggingConfigSchema(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPigJobQueryListSchema(),
			},

			"script_variables": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPigJobLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPigJobQueryListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPrestoJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"client_tags": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"continue_on_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPrestoJobLoggingConfigSchema(),
			},

			"output_format": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPrestoJobQueryListSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPrestoJobLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPrestoJobQueryListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPysparkJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"main_python_file_uri": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsPysparkJobLoggingConfigSchema(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"python_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsPysparkJobLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSchedulingSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max_failures_per_hour": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"max_failures_total": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkJobLoggingConfigSchema(),
			},

			"main_class": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"main_jar_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkJobLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkRJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"main_r_file_uri": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"archive_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"args": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkRJobLoggingConfigSchema(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkRJobLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkSqlJobSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"jar_file_uris": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"logging_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigSchema(),
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"query_file_uri": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"query_list": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateJobsSparkSqlJobQueryListSchema(),
			},

			"script_variables": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"driver_log_levels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateJobsSparkSqlJobQueryListSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"queries": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_selector": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementClusterSelectorSchema(),
			},

			"managed_cluster": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplatePlacementManagedClusterSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementClusterSelectorSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_labels": {
				Type:        schema.TypeMap,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplatePlacementManagedClusterSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: ``,
			},

			"config": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigSchema(),
			},

			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateParametersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"fields": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
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
				ForceNew:    true,
				Description: ``,
			},

			"validation": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateParametersValidationSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateParametersValidationSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regex": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateParametersValidationRegexSchema(),
			},

			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateParametersValidationValuesSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateParametersValidationRegexSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"regexes": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateParametersValidationValuesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"values": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateClusterInstanceGroupConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerators": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        DataprocWorkflowTemplateClusterInstanceGroupConfigAcceleratorsSchema(),
			},

			"disk_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterInstanceGroupConfigDiskConfigSchema(),
			},

			"image": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"machine_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"min_cpu_platform": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"num_instances": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"preemptibility": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"PREEMPTIBILITY_UNSPECIFIED", "NON_PREEMPTIBLE", "PREEMPTIBLE", ""}, false),
			},

			"instance_names": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"is_preemptible": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: ``,
			},

			"managed_group_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: ``,
				Elem:        DataprocWorkflowTemplateClusterInstanceGroupConfigManagedGroupConfigSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateClusterInstanceGroupConfigAcceleratorsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"accelerator_count": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"accelerator_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterInstanceGroupConfigDiskConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"boot_disk_size_gb": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"boot_disk_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"num_local_ssds": {
				Type:        schema.TypeInt,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterInstanceGroupConfigManagedGroupConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"instance_group_manager_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},

			"instance_template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"autoscaling_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigAutoscalingConfigSchema(),
			},

			"encryption_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigEncryptionConfigSchema(),
			},

			"endpoint_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigEndpointConfigSchema(),
			},

			"gce_cluster_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigGceClusterConfigSchema(),
			},

			"gke_cluster_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigSchema(),
			},

			"initialization_actions": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigInitializationActionsSchema(),
			},

			"lifecycle_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigLifecycleConfigSchema(),
			},

			"master_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterInstanceGroupConfigSchema(),
			},

			"metastore_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigMetastoreConfigSchema(),
			},

			"secondary_worker_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterInstanceGroupConfigSchema(),
			},

			"security_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigSecurityConfigSchema(),
			},

			"software_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigSoftwareConfigSchema(),
			},

			"staging_bucket": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"temp_bucket": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"worker_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterInstanceGroupConfigSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigAutoscalingConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"policy": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigEncryptionConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"gce_pd_kms_key_name": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigEndpointConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enable_http_port_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"http_ports": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigGceClusterConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"internal_ip_only": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"metadata": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"network": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"node_group_affinity": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigGceClusterConfigNodeGroupAffinitySchema(),
			},

			"private_ipv6_google_access": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED", "INHERIT_FROM_SUBNETWORK", "OUTBOUND", "BIDIRECTIONAL", ""}, false),
			},

			"reservation_affinity": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigGceClusterConfigReservationAffinitySchema(),
			},

			"service_account": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"service_account_scopes": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"subnetwork": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},

			"zone": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigGceClusterConfigNodeGroupAffinitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"node_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigGceClusterConfigReservationAffinitySchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"consume_reservation_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Description:  ``,
				ValidateFunc: validation.StringInSlice([]string{"TYPE_UNSPECIFIED", "NO_RESERVATION", "ANY_RESERVATION", "SPECIFIC_RESERVATION", ""}, false),
			},

			"key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"values": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"namespaced_gke_deployment_target": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"target_gke_cluster": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigInitializationActionsSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"executable_file": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"execution_timeout": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigLifecycleConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"auto_delete_time": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"auto_delete_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"idle_delete_ttl": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"idle_start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigMetastoreConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dataproc_metastore_service": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigSecurityConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"kerberos_config": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				MaxItems:    1,
				Elem:        DataprocWorkflowTemplateClusterClusterConfigSecurityConfigKerberosConfigSchema(),
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigSecurityConfigKerberosConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cross_realm_trust_admin_server": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"cross_realm_trust_kdc": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"cross_realm_trust_realm": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"cross_realm_trust_shared_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"enable_kerberos": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"kdc_db_key": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"key_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"keystore": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"keystore_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"kms_key": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      ``,
			},

			"realm": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"root_principal_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"tgt_lifetime_hours": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"truststore": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"truststore_password": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},
		},
	}
}

func DataprocWorkflowTemplateClusterClusterConfigSoftwareConfigSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"image_version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
			},

			"properties": {
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
				Description: ``,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceDataprocWorkflowTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &dataproc.WorkflowTemplate{
		Jobs:       expandDataprocWorkflowTemplateJobsArray(d.Get("jobs")),
		Location:   dcl.String(d.Get("location").(string)),
		Name:       dcl.String(d.Get("name").(string)),
		Placement:  expandDataprocWorkflowTemplatePlacement(d.Get("placement")),
		DagTimeout: dcl.String(d.Get("dag_timeout").(string)),
		Labels:     checkStringMap(d.Get("labels")),
		Parameters: expandDataprocWorkflowTemplateParametersArray(d.Get("parameters")),
		Project:    dcl.String(project),
		Version:    dcl.Int64OrNil(int64(d.Get("version").(int))),
	}

	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}")
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
	client := NewDCLDataprocClient(config, userAgent, billingProject)
	res, err := client.ApplyWorkflowTemplate(context.Background(), obj, createDirective...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating WorkflowTemplate: %s", err)
	}

	log.Printf("[DEBUG] Finished creating WorkflowTemplate %q: %#v", d.Id(), res)

	return resourceDataprocWorkflowTemplateRead(d, meta)
}

func resourceDataprocWorkflowTemplateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &dataproc.WorkflowTemplate{
		Jobs:       expandDataprocWorkflowTemplateJobsArray(d.Get("jobs")),
		Location:   dcl.String(d.Get("location").(string)),
		Name:       dcl.String(d.Get("name").(string)),
		Placement:  expandDataprocWorkflowTemplatePlacement(d.Get("placement")),
		DagTimeout: dcl.String(d.Get("dag_timeout").(string)),
		Labels:     checkStringMap(d.Get("labels")),
		Parameters: expandDataprocWorkflowTemplateParametersArray(d.Get("parameters")),
		Project:    dcl.String(project),
		Version:    dcl.Int64OrNil(int64(d.Get("version").(int))),
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
	client := NewDCLDataprocClient(config, userAgent, billingProject)
	res, err := client.GetWorkflowTemplate(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("DataprocWorkflowTemplate %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("jobs", flattenDataprocWorkflowTemplateJobsArray(res.Jobs)); err != nil {
		return fmt.Errorf("error setting jobs in state: %s", err)
	}
	if err = d.Set("location", res.Location); err != nil {
		return fmt.Errorf("error setting location in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	if err = d.Set("placement", flattenDataprocWorkflowTemplatePlacement(res.Placement)); err != nil {
		return fmt.Errorf("error setting placement in state: %s", err)
	}
	if err = d.Set("dag_timeout", res.DagTimeout); err != nil {
		return fmt.Errorf("error setting dag_timeout in state: %s", err)
	}
	if err = d.Set("labels", res.Labels); err != nil {
		return fmt.Errorf("error setting labels in state: %s", err)
	}
	if err = d.Set("parameters", flattenDataprocWorkflowTemplateParametersArray(res.Parameters)); err != nil {
		return fmt.Errorf("error setting parameters in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("version", res.Version); err != nil {
		return fmt.Errorf("error setting version in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("update_time", res.UpdateTime); err != nil {
		return fmt.Errorf("error setting update_time in state: %s", err)
	}

	return nil
}

func resourceDataprocWorkflowTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &dataproc.WorkflowTemplate{
		Jobs:       expandDataprocWorkflowTemplateJobsArray(d.Get("jobs")),
		Location:   dcl.String(d.Get("location").(string)),
		Name:       dcl.String(d.Get("name").(string)),
		Placement:  expandDataprocWorkflowTemplatePlacement(d.Get("placement")),
		DagTimeout: dcl.String(d.Get("dag_timeout").(string)),
		Labels:     checkStringMap(d.Get("labels")),
		Parameters: expandDataprocWorkflowTemplateParametersArray(d.Get("parameters")),
		Project:    dcl.String(project),
		Version:    dcl.Int64OrNil(int64(d.Get("version").(int))),
	}

	log.Printf("[DEBUG] Deleting WorkflowTemplate %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLDataprocClient(config, userAgent, billingProject)
	if err := client.DeleteWorkflowTemplate(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting WorkflowTemplate: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting WorkflowTemplate %q", d.Id())
	return nil
}

func resourceDataprocWorkflowTemplateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/workflowTemplates/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandDataprocWorkflowTemplateJobsArray(o interface{}) []dataproc.WorkflowTemplateJobs {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]dataproc.WorkflowTemplateJobs, 0, len(objs))
	for _, item := range objs {
		i := expandDataprocWorkflowTemplateJobs(item)
		items = append(items, *i)
	}

	return items
}

func expandDataprocWorkflowTemplateJobs(o interface{}) *dataproc.WorkflowTemplateJobs {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobs
	}

	obj := o.(map[string]interface{})
	return &dataproc.WorkflowTemplateJobs{
		StepId:              dcl.String(obj["step_id"].(string)),
		HadoopJob:           expandDataprocWorkflowTemplateJobsHadoopJob(obj["hadoop_job"]),
		HiveJob:             expandDataprocWorkflowTemplateJobsHiveJob(obj["hive_job"]),
		Labels:              checkStringMap(obj["labels"]),
		PigJob:              expandDataprocWorkflowTemplateJobsPigJob(obj["pig_job"]),
		PrerequisiteStepIds: expandStringArray(obj["prerequisite_step_ids"]),
		PrestoJob:           expandDataprocWorkflowTemplateJobsPrestoJob(obj["presto_job"]),
		PysparkJob:          expandDataprocWorkflowTemplateJobsPysparkJob(obj["pyspark_job"]),
		Scheduling:          expandDataprocWorkflowTemplateJobsScheduling(obj["scheduling"]),
		SparkJob:            expandDataprocWorkflowTemplateJobsSparkJob(obj["spark_job"]),
		SparkRJob:           expandDataprocWorkflowTemplateJobsSparkRJob(obj["spark_r_job"]),
		SparkSqlJob:         expandDataprocWorkflowTemplateJobsSparkSqlJob(obj["spark_sql_job"]),
	}
}

func flattenDataprocWorkflowTemplateJobsArray(objs []dataproc.WorkflowTemplateJobs) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenDataprocWorkflowTemplateJobs(&item)
		items = append(items, i)
	}

	return items
}

func flattenDataprocWorkflowTemplateJobs(obj *dataproc.WorkflowTemplateJobs) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"step_id":               obj.StepId,
		"hadoop_job":            flattenDataprocWorkflowTemplateJobsHadoopJob(obj.HadoopJob),
		"hive_job":              flattenDataprocWorkflowTemplateJobsHiveJob(obj.HiveJob),
		"labels":                obj.Labels,
		"pig_job":               flattenDataprocWorkflowTemplateJobsPigJob(obj.PigJob),
		"prerequisite_step_ids": obj.PrerequisiteStepIds,
		"presto_job":            flattenDataprocWorkflowTemplateJobsPrestoJob(obj.PrestoJob),
		"pyspark_job":           flattenDataprocWorkflowTemplateJobsPysparkJob(obj.PysparkJob),
		"scheduling":            flattenDataprocWorkflowTemplateJobsScheduling(obj.Scheduling),
		"spark_job":             flattenDataprocWorkflowTemplateJobsSparkJob(obj.SparkJob),
		"spark_r_job":           flattenDataprocWorkflowTemplateJobsSparkRJob(obj.SparkRJob),
		"spark_sql_job":         flattenDataprocWorkflowTemplateJobsSparkSqlJob(obj.SparkSqlJob),
	}

	return transformed

}

func expandDataprocWorkflowTemplateJobsHadoopJob(o interface{}) *dataproc.WorkflowTemplateJobsHadoopJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsHadoopJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsHadoopJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsHadoopJob{
		ArchiveUris:    expandStringArray(obj["archive_uris"]),
		Args:           expandStringArray(obj["args"]),
		FileUris:       expandStringArray(obj["file_uris"]),
		JarFileUris:    expandStringArray(obj["jar_file_uris"]),
		LoggingConfig:  expandDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(obj["logging_config"]),
		MainClass:      dcl.String(obj["main_class"].(string)),
		MainJarFileUri: dcl.String(obj["main_jar_file_uri"].(string)),
		Properties:     checkStringMap(obj["properties"]),
	}
}

func flattenDataprocWorkflowTemplateJobsHadoopJob(obj *dataproc.WorkflowTemplateJobsHadoopJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"archive_uris":      obj.ArchiveUris,
		"args":              obj.Args,
		"file_uris":         obj.FileUris,
		"jar_file_uris":     obj.JarFileUris,
		"logging_config":    flattenDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(obj.LoggingConfig),
		"main_class":        obj.MainClass,
		"main_jar_file_uri": obj.MainJarFileUri,
		"properties":        obj.Properties,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(o interface{}) *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsHadoopJobLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsHadoopJobLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig{
		DriverLogLevels: checkStringMap(obj["driver_log_levels"]),
	}
}

func flattenDataprocWorkflowTemplateJobsHadoopJobLoggingConfig(obj *dataproc.WorkflowTemplateJobsHadoopJobLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"driver_log_levels": obj.DriverLogLevels,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsHiveJob(o interface{}) *dataproc.WorkflowTemplateJobsHiveJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsHiveJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsHiveJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsHiveJob{
		ContinueOnFailure: dcl.Bool(obj["continue_on_failure"].(bool)),
		JarFileUris:       expandStringArray(obj["jar_file_uris"]),
		Properties:        checkStringMap(obj["properties"]),
		QueryFileUri:      dcl.String(obj["query_file_uri"].(string)),
		QueryList:         expandDataprocWorkflowTemplateJobsHiveJobQueryList(obj["query_list"]),
		ScriptVariables:   checkStringMap(obj["script_variables"]),
	}
}

func flattenDataprocWorkflowTemplateJobsHiveJob(obj *dataproc.WorkflowTemplateJobsHiveJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"continue_on_failure": obj.ContinueOnFailure,
		"jar_file_uris":       obj.JarFileUris,
		"properties":          obj.Properties,
		"query_file_uri":      obj.QueryFileUri,
		"query_list":          flattenDataprocWorkflowTemplateJobsHiveJobQueryList(obj.QueryList),
		"script_variables":    obj.ScriptVariables,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsHiveJobQueryList(o interface{}) *dataproc.WorkflowTemplateJobsHiveJobQueryList {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsHiveJobQueryList
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsHiveJobQueryList
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsHiveJobQueryList{
		Queries: expandStringArray(obj["queries"]),
	}
}

func flattenDataprocWorkflowTemplateJobsHiveJobQueryList(obj *dataproc.WorkflowTemplateJobsHiveJobQueryList) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"queries": obj.Queries,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPigJob(o interface{}) *dataproc.WorkflowTemplateJobsPigJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPigJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPigJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPigJob{
		ContinueOnFailure: dcl.Bool(obj["continue_on_failure"].(bool)),
		JarFileUris:       expandStringArray(obj["jar_file_uris"]),
		LoggingConfig:     expandDataprocWorkflowTemplateJobsPigJobLoggingConfig(obj["logging_config"]),
		Properties:        checkStringMap(obj["properties"]),
		QueryFileUri:      dcl.String(obj["query_file_uri"].(string)),
		QueryList:         expandDataprocWorkflowTemplateJobsPigJobQueryList(obj["query_list"]),
		ScriptVariables:   checkStringMap(obj["script_variables"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPigJob(obj *dataproc.WorkflowTemplateJobsPigJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"continue_on_failure": obj.ContinueOnFailure,
		"jar_file_uris":       obj.JarFileUris,
		"logging_config":      flattenDataprocWorkflowTemplateJobsPigJobLoggingConfig(obj.LoggingConfig),
		"properties":          obj.Properties,
		"query_file_uri":      obj.QueryFileUri,
		"query_list":          flattenDataprocWorkflowTemplateJobsPigJobQueryList(obj.QueryList),
		"script_variables":    obj.ScriptVariables,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPigJobLoggingConfig(o interface{}) *dataproc.WorkflowTemplateJobsPigJobLoggingConfig {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPigJobLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPigJobLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPigJobLoggingConfig{
		DriverLogLevels: checkStringMap(obj["driver_log_levels"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPigJobLoggingConfig(obj *dataproc.WorkflowTemplateJobsPigJobLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"driver_log_levels": obj.DriverLogLevels,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPigJobQueryList(o interface{}) *dataproc.WorkflowTemplateJobsPigJobQueryList {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPigJobQueryList
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPigJobQueryList
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPigJobQueryList{
		Queries: expandStringArray(obj["queries"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPigJobQueryList(obj *dataproc.WorkflowTemplateJobsPigJobQueryList) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"queries": obj.Queries,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPrestoJob(o interface{}) *dataproc.WorkflowTemplateJobsPrestoJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPrestoJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPrestoJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPrestoJob{
		ClientTags:        expandStringArray(obj["client_tags"]),
		ContinueOnFailure: dcl.Bool(obj["continue_on_failure"].(bool)),
		LoggingConfig:     expandDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(obj["logging_config"]),
		OutputFormat:      dcl.String(obj["output_format"].(string)),
		Properties:        checkStringMap(obj["properties"]),
		QueryFileUri:      dcl.String(obj["query_file_uri"].(string)),
		QueryList:         expandDataprocWorkflowTemplateJobsPrestoJobQueryList(obj["query_list"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPrestoJob(obj *dataproc.WorkflowTemplateJobsPrestoJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"client_tags":         obj.ClientTags,
		"continue_on_failure": obj.ContinueOnFailure,
		"logging_config":      flattenDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(obj.LoggingConfig),
		"output_format":       obj.OutputFormat,
		"properties":          obj.Properties,
		"query_file_uri":      obj.QueryFileUri,
		"query_list":          flattenDataprocWorkflowTemplateJobsPrestoJobQueryList(obj.QueryList),
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(o interface{}) *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPrestoJobLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPrestoJobLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig{
		DriverLogLevels: checkStringMap(obj["driver_log_levels"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPrestoJobLoggingConfig(obj *dataproc.WorkflowTemplateJobsPrestoJobLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"driver_log_levels": obj.DriverLogLevels,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPrestoJobQueryList(o interface{}) *dataproc.WorkflowTemplateJobsPrestoJobQueryList {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPrestoJobQueryList
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPrestoJobQueryList
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPrestoJobQueryList{
		Queries: expandStringArray(obj["queries"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPrestoJobQueryList(obj *dataproc.WorkflowTemplateJobsPrestoJobQueryList) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"queries": obj.Queries,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPysparkJob(o interface{}) *dataproc.WorkflowTemplateJobsPysparkJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPysparkJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPysparkJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPysparkJob{
		MainPythonFileUri: dcl.String(obj["main_python_file_uri"].(string)),
		ArchiveUris:       expandStringArray(obj["archive_uris"]),
		Args:              expandStringArray(obj["args"]),
		FileUris:          expandStringArray(obj["file_uris"]),
		JarFileUris:       expandStringArray(obj["jar_file_uris"]),
		LoggingConfig:     expandDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(obj["logging_config"]),
		Properties:        checkStringMap(obj["properties"]),
		PythonFileUris:    expandStringArray(obj["python_file_uris"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPysparkJob(obj *dataproc.WorkflowTemplateJobsPysparkJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"main_python_file_uri": obj.MainPythonFileUri,
		"archive_uris":         obj.ArchiveUris,
		"args":                 obj.Args,
		"file_uris":            obj.FileUris,
		"jar_file_uris":        obj.JarFileUris,
		"logging_config":       flattenDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(obj.LoggingConfig),
		"properties":           obj.Properties,
		"python_file_uris":     obj.PythonFileUris,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(o interface{}) *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsPysparkJobLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsPysparkJobLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig{
		DriverLogLevels: checkStringMap(obj["driver_log_levels"]),
	}
}

func flattenDataprocWorkflowTemplateJobsPysparkJobLoggingConfig(obj *dataproc.WorkflowTemplateJobsPysparkJobLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"driver_log_levels": obj.DriverLogLevels,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsScheduling(o interface{}) *dataproc.WorkflowTemplateJobsScheduling {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsScheduling
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsScheduling
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsScheduling{
		MaxFailuresPerHour: dcl.Int64(int64(obj["max_failures_per_hour"].(int))),
		MaxFailuresTotal:   dcl.Int64(int64(obj["max_failures_total"].(int))),
	}
}

func flattenDataprocWorkflowTemplateJobsScheduling(obj *dataproc.WorkflowTemplateJobsScheduling) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"max_failures_per_hour": obj.MaxFailuresPerHour,
		"max_failures_total":    obj.MaxFailuresTotal,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsSparkJob(o interface{}) *dataproc.WorkflowTemplateJobsSparkJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsSparkJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsSparkJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsSparkJob{
		ArchiveUris:    expandStringArray(obj["archive_uris"]),
		Args:           expandStringArray(obj["args"]),
		FileUris:       expandStringArray(obj["file_uris"]),
		JarFileUris:    expandStringArray(obj["jar_file_uris"]),
		LoggingConfig:  expandDataprocWorkflowTemplateJobsSparkJobLoggingConfig(obj["logging_config"]),
		MainClass:      dcl.String(obj["main_class"].(string)),
		MainJarFileUri: dcl.String(obj["main_jar_file_uri"].(string)),
		Properties:     checkStringMap(obj["properties"]),
	}
}

func flattenDataprocWorkflowTemplateJobsSparkJob(obj *dataproc.WorkflowTemplateJobsSparkJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"archive_uris":      obj.ArchiveUris,
		"args":              obj.Args,
		"file_uris":         obj.FileUris,
		"jar_file_uris":     obj.JarFileUris,
		"logging_config":    flattenDataprocWorkflowTemplateJobsSparkJobLoggingConfig(obj.LoggingConfig),
		"main_class":        obj.MainClass,
		"main_jar_file_uri": obj.MainJarFileUri,
		"properties":        obj.Properties,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsSparkJobLoggingConfig(o interface{}) *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsSparkJobLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsSparkJobLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsSparkJobLoggingConfig{
		DriverLogLevels: checkStringMap(obj["driver_log_levels"]),
	}
}

func flattenDataprocWorkflowTemplateJobsSparkJobLoggingConfig(obj *dataproc.WorkflowTemplateJobsSparkJobLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"driver_log_levels": obj.DriverLogLevels,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsSparkRJob(o interface{}) *dataproc.WorkflowTemplateJobsSparkRJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsSparkRJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsSparkRJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsSparkRJob{
		MainRFileUri:  dcl.String(obj["main_r_file_uri"].(string)),
		ArchiveUris:   expandStringArray(obj["archive_uris"]),
		Args:          expandStringArray(obj["args"]),
		FileUris:      expandStringArray(obj["file_uris"]),
		LoggingConfig: expandDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(obj["logging_config"]),
		Properties:    checkStringMap(obj["properties"]),
	}
}

func flattenDataprocWorkflowTemplateJobsSparkRJob(obj *dataproc.WorkflowTemplateJobsSparkRJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"main_r_file_uri": obj.MainRFileUri,
		"archive_uris":    obj.ArchiveUris,
		"args":            obj.Args,
		"file_uris":       obj.FileUris,
		"logging_config":  flattenDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(obj.LoggingConfig),
		"properties":      obj.Properties,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(o interface{}) *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsSparkRJobLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsSparkRJobLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig{
		DriverLogLevels: checkStringMap(obj["driver_log_levels"]),
	}
}

func flattenDataprocWorkflowTemplateJobsSparkRJobLoggingConfig(obj *dataproc.WorkflowTemplateJobsSparkRJobLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"driver_log_levels": obj.DriverLogLevels,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsSparkSqlJob(o interface{}) *dataproc.WorkflowTemplateJobsSparkSqlJob {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsSparkSqlJob
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsSparkSqlJob
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsSparkSqlJob{
		JarFileUris:     expandStringArray(obj["jar_file_uris"]),
		LoggingConfig:   expandDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(obj["logging_config"]),
		Properties:      checkStringMap(obj["properties"]),
		QueryFileUri:    dcl.String(obj["query_file_uri"].(string)),
		QueryList:       expandDataprocWorkflowTemplateJobsSparkSqlJobQueryList(obj["query_list"]),
		ScriptVariables: checkStringMap(obj["script_variables"]),
	}
}

func flattenDataprocWorkflowTemplateJobsSparkSqlJob(obj *dataproc.WorkflowTemplateJobsSparkSqlJob) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"jar_file_uris":    obj.JarFileUris,
		"logging_config":   flattenDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(obj.LoggingConfig),
		"properties":       obj.Properties,
		"query_file_uri":   obj.QueryFileUri,
		"query_list":       flattenDataprocWorkflowTemplateJobsSparkSqlJobQueryList(obj.QueryList),
		"script_variables": obj.ScriptVariables,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(o interface{}) *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsSparkSqlJobLoggingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsSparkSqlJobLoggingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig{
		DriverLogLevels: checkStringMap(obj["driver_log_levels"]),
	}
}

func flattenDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig(obj *dataproc.WorkflowTemplateJobsSparkSqlJobLoggingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"driver_log_levels": obj.DriverLogLevels,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateJobsSparkSqlJobQueryList(o interface{}) *dataproc.WorkflowTemplateJobsSparkSqlJobQueryList {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateJobsSparkSqlJobQueryList
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateJobsSparkSqlJobQueryList
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateJobsSparkSqlJobQueryList{
		Queries: expandStringArray(obj["queries"]),
	}
}

func flattenDataprocWorkflowTemplateJobsSparkSqlJobQueryList(obj *dataproc.WorkflowTemplateJobsSparkSqlJobQueryList) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"queries": obj.Queries,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplatePlacement(o interface{}) *dataproc.WorkflowTemplatePlacement {
	if o == nil {
		return dataproc.EmptyWorkflowTemplatePlacement
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplatePlacement
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplatePlacement{
		ClusterSelector: expandDataprocWorkflowTemplatePlacementClusterSelector(obj["cluster_selector"]),
		ManagedCluster:  expandDataprocWorkflowTemplatePlacementManagedCluster(obj["managed_cluster"]),
	}
}

func flattenDataprocWorkflowTemplatePlacement(obj *dataproc.WorkflowTemplatePlacement) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cluster_selector": flattenDataprocWorkflowTemplatePlacementClusterSelector(obj.ClusterSelector),
		"managed_cluster":  flattenDataprocWorkflowTemplatePlacementManagedCluster(obj.ManagedCluster),
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplatePlacementClusterSelector(o interface{}) *dataproc.WorkflowTemplatePlacementClusterSelector {
	if o == nil {
		return dataproc.EmptyWorkflowTemplatePlacementClusterSelector
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplatePlacementClusterSelector
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplatePlacementClusterSelector{
		ClusterLabels: checkStringMap(obj["cluster_labels"]),
		Zone:          dcl.StringOrNil(obj["zone"].(string)),
	}
}

func flattenDataprocWorkflowTemplatePlacementClusterSelector(obj *dataproc.WorkflowTemplatePlacementClusterSelector) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cluster_labels": obj.ClusterLabels,
		"zone":           obj.Zone,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplatePlacementManagedCluster(o interface{}) *dataproc.WorkflowTemplatePlacementManagedCluster {
	if o == nil {
		return dataproc.EmptyWorkflowTemplatePlacementManagedCluster
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplatePlacementManagedCluster
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplatePlacementManagedCluster{
		ClusterName: dcl.String(obj["cluster_name"].(string)),
		Config:      expandDataprocWorkflowTemplateClusterClusterConfig(obj["config"]),
		Labels:      checkStringMap(obj["labels"]),
	}
}

func flattenDataprocWorkflowTemplatePlacementManagedCluster(obj *dataproc.WorkflowTemplatePlacementManagedCluster) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cluster_name": obj.ClusterName,
		"config":       flattenDataprocWorkflowTemplateClusterClusterConfig(obj.Config),
		"labels":       obj.Labels,
	}

	return []interface{}{transformed}

}
func expandDataprocWorkflowTemplateParametersArray(o interface{}) []dataproc.WorkflowTemplateParameters {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]dataproc.WorkflowTemplateParameters, 0, len(objs))
	for _, item := range objs {
		i := expandDataprocWorkflowTemplateParameters(item)
		items = append(items, *i)
	}

	return items
}

func expandDataprocWorkflowTemplateParameters(o interface{}) *dataproc.WorkflowTemplateParameters {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateParameters
	}

	obj := o.(map[string]interface{})
	return &dataproc.WorkflowTemplateParameters{
		Fields:      expandStringArray(obj["fields"]),
		Name:        dcl.String(obj["name"].(string)),
		Description: dcl.String(obj["description"].(string)),
		Validation:  expandDataprocWorkflowTemplateParametersValidation(obj["validation"]),
	}
}

func flattenDataprocWorkflowTemplateParametersArray(objs []dataproc.WorkflowTemplateParameters) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenDataprocWorkflowTemplateParameters(&item)
		items = append(items, i)
	}

	return items
}

func flattenDataprocWorkflowTemplateParameters(obj *dataproc.WorkflowTemplateParameters) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"fields":      obj.Fields,
		"name":        obj.Name,
		"description": obj.Description,
		"validation":  flattenDataprocWorkflowTemplateParametersValidation(obj.Validation),
	}

	return transformed

}

func expandDataprocWorkflowTemplateParametersValidation(o interface{}) *dataproc.WorkflowTemplateParametersValidation {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateParametersValidation
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateParametersValidation
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateParametersValidation{
		Regex:  expandDataprocWorkflowTemplateParametersValidationRegex(obj["regex"]),
		Values: expandDataprocWorkflowTemplateParametersValidationValues(obj["values"]),
	}
}

func flattenDataprocWorkflowTemplateParametersValidation(obj *dataproc.WorkflowTemplateParametersValidation) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"regex":  flattenDataprocWorkflowTemplateParametersValidationRegex(obj.Regex),
		"values": flattenDataprocWorkflowTemplateParametersValidationValues(obj.Values),
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateParametersValidationRegex(o interface{}) *dataproc.WorkflowTemplateParametersValidationRegex {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateParametersValidationRegex
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateParametersValidationRegex
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateParametersValidationRegex{
		Regexes: expandStringArray(obj["regexes"]),
	}
}

func flattenDataprocWorkflowTemplateParametersValidationRegex(obj *dataproc.WorkflowTemplateParametersValidationRegex) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"regexes": obj.Regexes,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateParametersValidationValues(o interface{}) *dataproc.WorkflowTemplateParametersValidationValues {
	if o == nil {
		return dataproc.EmptyWorkflowTemplateParametersValidationValues
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyWorkflowTemplateParametersValidationValues
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.WorkflowTemplateParametersValidationValues{
		Values: expandStringArray(obj["values"]),
	}
}

func flattenDataprocWorkflowTemplateParametersValidationValues(obj *dataproc.WorkflowTemplateParametersValidationValues) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"values": obj.Values,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterInstanceGroupConfig(o interface{}) *dataproc.ClusterInstanceGroupConfig {
	if o == nil {
		return dataproc.EmptyClusterInstanceGroupConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterInstanceGroupConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterInstanceGroupConfig{
		Accelerators:   expandDataprocWorkflowTemplateClusterInstanceGroupConfigAcceleratorsArray(obj["accelerators"]),
		DiskConfig:     expandDataprocWorkflowTemplateClusterInstanceGroupConfigDiskConfig(obj["disk_config"]),
		Image:          dcl.String(obj["image"].(string)),
		MachineType:    dcl.String(obj["machine_type"].(string)),
		MinCpuPlatform: dcl.String(obj["min_cpu_platform"].(string)),
		NumInstances:   dcl.Int64(int64(obj["num_instances"].(int))),
		Preemptibility: dataproc.ClusterInstanceGroupConfigPreemptibilityEnumRef(obj["preemptibility"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterInstanceGroupConfig(obj *dataproc.ClusterInstanceGroupConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"accelerators":         flattenDataprocWorkflowTemplateClusterInstanceGroupConfigAcceleratorsArray(obj.Accelerators),
		"disk_config":          flattenDataprocWorkflowTemplateClusterInstanceGroupConfigDiskConfig(obj.DiskConfig),
		"image":                obj.Image,
		"machine_type":         obj.MachineType,
		"min_cpu_platform":     obj.MinCpuPlatform,
		"num_instances":        obj.NumInstances,
		"preemptibility":       obj.Preemptibility,
		"instance_names":       obj.InstanceNames,
		"is_preemptible":       obj.IsPreemptible,
		"managed_group_config": flattenDataprocWorkflowTemplateClusterInstanceGroupConfigManagedGroupConfig(obj.ManagedGroupConfig),
	}

	return []interface{}{transformed}

}
func expandDataprocWorkflowTemplateClusterInstanceGroupConfigAcceleratorsArray(o interface{}) []dataproc.ClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]dataproc.ClusterInstanceGroupConfigAccelerators, 0, len(objs))
	for _, item := range objs {
		i := expandDataprocWorkflowTemplateClusterInstanceGroupConfigAccelerators(item)
		items = append(items, *i)
	}

	return items
}

func expandDataprocWorkflowTemplateClusterInstanceGroupConfigAccelerators(o interface{}) *dataproc.ClusterInstanceGroupConfigAccelerators {
	if o == nil {
		return dataproc.EmptyClusterInstanceGroupConfigAccelerators
	}

	obj := o.(map[string]interface{})
	return &dataproc.ClusterInstanceGroupConfigAccelerators{
		AcceleratorCount: dcl.Int64(int64(obj["accelerator_count"].(int))),
		AcceleratorType:  dcl.String(obj["accelerator_type"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterInstanceGroupConfigAcceleratorsArray(objs []dataproc.ClusterInstanceGroupConfigAccelerators) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenDataprocWorkflowTemplateClusterInstanceGroupConfigAccelerators(&item)
		items = append(items, i)
	}

	return items
}

func flattenDataprocWorkflowTemplateClusterInstanceGroupConfigAccelerators(obj *dataproc.ClusterInstanceGroupConfigAccelerators) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"accelerator_count": obj.AcceleratorCount,
		"accelerator_type":  obj.AcceleratorType,
	}

	return transformed

}

func expandDataprocWorkflowTemplateClusterInstanceGroupConfigDiskConfig(o interface{}) *dataproc.ClusterInstanceGroupConfigDiskConfig {
	if o == nil {
		return dataproc.EmptyClusterInstanceGroupConfigDiskConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterInstanceGroupConfigDiskConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterInstanceGroupConfigDiskConfig{
		BootDiskSizeGb: dcl.Int64(int64(obj["boot_disk_size_gb"].(int))),
		BootDiskType:   dcl.String(obj["boot_disk_type"].(string)),
		NumLocalSsds:   dcl.Int64OrNil(int64(obj["num_local_ssds"].(int))),
	}
}

func flattenDataprocWorkflowTemplateClusterInstanceGroupConfigDiskConfig(obj *dataproc.ClusterInstanceGroupConfigDiskConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"boot_disk_size_gb": obj.BootDiskSizeGb,
		"boot_disk_type":    obj.BootDiskType,
		"num_local_ssds":    obj.NumLocalSsds,
	}

	return []interface{}{transformed}

}

func flattenDataprocWorkflowTemplateClusterInstanceGroupConfigManagedGroupConfig(obj *dataproc.ClusterInstanceGroupConfigManagedGroupConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"instance_group_manager_name": obj.InstanceGroupManagerName,
		"instance_template_name":      obj.InstanceTemplateName,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfig(o interface{}) *dataproc.ClusterClusterConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfig{
		AutoscalingConfig:     expandDataprocWorkflowTemplateClusterClusterConfigAutoscalingConfig(obj["autoscaling_config"]),
		EncryptionConfig:      expandDataprocWorkflowTemplateClusterClusterConfigEncryptionConfig(obj["encryption_config"]),
		EndpointConfig:        expandDataprocWorkflowTemplateClusterClusterConfigEndpointConfig(obj["endpoint_config"]),
		GceClusterConfig:      expandDataprocWorkflowTemplateClusterClusterConfigGceClusterConfig(obj["gce_cluster_config"]),
		GkeClusterConfig:      expandDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfig(obj["gke_cluster_config"]),
		InitializationActions: expandDataprocWorkflowTemplateClusterClusterConfigInitializationActionsArray(obj["initialization_actions"]),
		LifecycleConfig:       expandDataprocWorkflowTemplateClusterClusterConfigLifecycleConfig(obj["lifecycle_config"]),
		MasterConfig:          expandDataprocWorkflowTemplateClusterInstanceGroupConfig(obj["master_config"]),
		MetastoreConfig:       expandDataprocWorkflowTemplateClusterClusterConfigMetastoreConfig(obj["metastore_config"]),
		SecondaryWorkerConfig: expandDataprocWorkflowTemplateClusterInstanceGroupConfig(obj["secondary_worker_config"]),
		SecurityConfig:        expandDataprocWorkflowTemplateClusterClusterConfigSecurityConfig(obj["security_config"]),
		SoftwareConfig:        expandDataprocWorkflowTemplateClusterClusterConfigSoftwareConfig(obj["software_config"]),
		StagingBucket:         dcl.String(obj["staging_bucket"].(string)),
		TempBucket:            dcl.String(obj["temp_bucket"].(string)),
		WorkerConfig:          expandDataprocWorkflowTemplateClusterInstanceGroupConfig(obj["worker_config"]),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfig(obj *dataproc.ClusterClusterConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"autoscaling_config":      flattenDataprocWorkflowTemplateClusterClusterConfigAutoscalingConfig(obj.AutoscalingConfig),
		"encryption_config":       flattenDataprocWorkflowTemplateClusterClusterConfigEncryptionConfig(obj.EncryptionConfig),
		"endpoint_config":         flattenDataprocWorkflowTemplateClusterClusterConfigEndpointConfig(obj.EndpointConfig),
		"gce_cluster_config":      flattenDataprocWorkflowTemplateClusterClusterConfigGceClusterConfig(obj.GceClusterConfig),
		"gke_cluster_config":      flattenDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfig(obj.GkeClusterConfig),
		"initialization_actions":  flattenDataprocWorkflowTemplateClusterClusterConfigInitializationActionsArray(obj.InitializationActions),
		"lifecycle_config":        flattenDataprocWorkflowTemplateClusterClusterConfigLifecycleConfig(obj.LifecycleConfig),
		"master_config":           flattenDataprocWorkflowTemplateClusterInstanceGroupConfig(obj.MasterConfig),
		"metastore_config":        flattenDataprocWorkflowTemplateClusterClusterConfigMetastoreConfig(obj.MetastoreConfig),
		"secondary_worker_config": flattenDataprocWorkflowTemplateClusterInstanceGroupConfig(obj.SecondaryWorkerConfig),
		"security_config":         flattenDataprocWorkflowTemplateClusterClusterConfigSecurityConfig(obj.SecurityConfig),
		"software_config":         flattenDataprocWorkflowTemplateClusterClusterConfigSoftwareConfig(obj.SoftwareConfig),
		"staging_bucket":          obj.StagingBucket,
		"temp_bucket":             obj.TempBucket,
		"worker_config":           flattenDataprocWorkflowTemplateClusterInstanceGroupConfig(obj.WorkerConfig),
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigAutoscalingConfig(o interface{}) *dataproc.ClusterClusterConfigAutoscalingConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigAutoscalingConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigAutoscalingConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigAutoscalingConfig{
		Policy: dcl.String(obj["policy"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigAutoscalingConfig(obj *dataproc.ClusterClusterConfigAutoscalingConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"policy": obj.Policy,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigEncryptionConfig(o interface{}) *dataproc.ClusterClusterConfigEncryptionConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigEncryptionConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigEncryptionConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigEncryptionConfig{
		GcePdKmsKeyName: dcl.String(obj["gce_pd_kms_key_name"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigEncryptionConfig(obj *dataproc.ClusterClusterConfigEncryptionConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"gce_pd_kms_key_name": obj.GcePdKmsKeyName,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigEndpointConfig(o interface{}) *dataproc.ClusterClusterConfigEndpointConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigEndpointConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigEndpointConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigEndpointConfig{
		EnableHttpPortAccess: dcl.Bool(obj["enable_http_port_access"].(bool)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigEndpointConfig(obj *dataproc.ClusterClusterConfigEndpointConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"enable_http_port_access": obj.EnableHttpPortAccess,
		"http_ports":              obj.HttpPorts,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigGceClusterConfig(o interface{}) *dataproc.ClusterClusterConfigGceClusterConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigGceClusterConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigGceClusterConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigGceClusterConfig{
		InternalIPOnly:          dcl.Bool(obj["internal_ip_only"].(bool)),
		Metadata:                checkStringMap(obj["metadata"]),
		Network:                 dcl.String(obj["network"].(string)),
		NodeGroupAffinity:       expandDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigNodeGroupAffinity(obj["node_group_affinity"]),
		PrivateIPv6GoogleAccess: dataproc.ClusterClusterConfigGceClusterConfigPrivateIPv6GoogleAccessEnumRef(obj["private_ipv6_google_access"].(string)),
		ReservationAffinity:     expandDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigReservationAffinity(obj["reservation_affinity"]),
		ServiceAccount:          dcl.String(obj["service_account"].(string)),
		ServiceAccountScopes:    expandStringArray(obj["service_account_scopes"]),
		Subnetwork:              dcl.String(obj["subnetwork"].(string)),
		Tags:                    expandStringArray(obj["tags"]),
		Zone:                    dcl.StringOrNil(obj["zone"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigGceClusterConfig(obj *dataproc.ClusterClusterConfigGceClusterConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"internal_ip_only":           obj.InternalIPOnly,
		"metadata":                   obj.Metadata,
		"network":                    obj.Network,
		"node_group_affinity":        flattenDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigNodeGroupAffinity(obj.NodeGroupAffinity),
		"private_ipv6_google_access": obj.PrivateIPv6GoogleAccess,
		"reservation_affinity":       flattenDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigReservationAffinity(obj.ReservationAffinity),
		"service_account":            obj.ServiceAccount,
		"service_account_scopes":     obj.ServiceAccountScopes,
		"subnetwork":                 obj.Subnetwork,
		"tags":                       obj.Tags,
		"zone":                       obj.Zone,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigNodeGroupAffinity(o interface{}) *dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigGceClusterConfigNodeGroupAffinity
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigGceClusterConfigNodeGroupAffinity
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity{
		NodeGroup: dcl.String(obj["node_group"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigNodeGroupAffinity(obj *dataproc.ClusterClusterConfigGceClusterConfigNodeGroupAffinity) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"node_group": obj.NodeGroup,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigReservationAffinity(o interface{}) *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigGceClusterConfigReservationAffinity
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigGceClusterConfigReservationAffinity
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity{
		ConsumeReservationType: dataproc.ClusterClusterConfigGceClusterConfigReservationAffinityConsumeReservationTypeEnumRef(obj["consume_reservation_type"].(string)),
		Key:                    dcl.String(obj["key"].(string)),
		Values:                 expandStringArray(obj["values"]),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigGceClusterConfigReservationAffinity(obj *dataproc.ClusterClusterConfigGceClusterConfigReservationAffinity) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"consume_reservation_type": obj.ConsumeReservationType,
		"key":                      obj.Key,
		"values":                   obj.Values,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfig(o interface{}) *dataproc.ClusterClusterConfigGkeClusterConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigGkeClusterConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigGkeClusterConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigGkeClusterConfig{
		NamespacedGkeDeploymentTarget: expandDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(obj["namespaced_gke_deployment_target"]),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfig(obj *dataproc.ClusterClusterConfigGkeClusterConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"namespaced_gke_deployment_target": flattenDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(obj.NamespacedGkeDeploymentTarget),
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(o interface{}) *dataproc.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget{
		ClusterNamespace: dcl.String(obj["cluster_namespace"].(string)),
		TargetGkeCluster: dcl.String(obj["target_gke_cluster"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget(obj *dataproc.ClusterClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cluster_namespace":  obj.ClusterNamespace,
		"target_gke_cluster": obj.TargetGkeCluster,
	}

	return []interface{}{transformed}

}
func expandDataprocWorkflowTemplateClusterClusterConfigInitializationActionsArray(o interface{}) []dataproc.ClusterClusterConfigInitializationActions {
	if o == nil {
		return nil
	}

	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}

	items := make([]dataproc.ClusterClusterConfigInitializationActions, 0, len(objs))
	for _, item := range objs {
		i := expandDataprocWorkflowTemplateClusterClusterConfigInitializationActions(item)
		items = append(items, *i)
	}

	return items
}

func expandDataprocWorkflowTemplateClusterClusterConfigInitializationActions(o interface{}) *dataproc.ClusterClusterConfigInitializationActions {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigInitializationActions
	}

	obj := o.(map[string]interface{})
	return &dataproc.ClusterClusterConfigInitializationActions{
		ExecutableFile:   dcl.String(obj["executable_file"].(string)),
		ExecutionTimeout: dcl.String(obj["execution_timeout"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigInitializationActionsArray(objs []dataproc.ClusterClusterConfigInitializationActions) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenDataprocWorkflowTemplateClusterClusterConfigInitializationActions(&item)
		items = append(items, i)
	}

	return items
}

func flattenDataprocWorkflowTemplateClusterClusterConfigInitializationActions(obj *dataproc.ClusterClusterConfigInitializationActions) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"executable_file":   obj.ExecutableFile,
		"execution_timeout": obj.ExecutionTimeout,
	}

	return transformed

}

func expandDataprocWorkflowTemplateClusterClusterConfigLifecycleConfig(o interface{}) *dataproc.ClusterClusterConfigLifecycleConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigLifecycleConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigLifecycleConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigLifecycleConfig{
		AutoDeleteTime: dcl.String(obj["auto_delete_time"].(string)),
		AutoDeleteTtl:  dcl.String(obj["auto_delete_ttl"].(string)),
		IdleDeleteTtl:  dcl.String(obj["idle_delete_ttl"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigLifecycleConfig(obj *dataproc.ClusterClusterConfigLifecycleConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"auto_delete_time": obj.AutoDeleteTime,
		"auto_delete_ttl":  obj.AutoDeleteTtl,
		"idle_delete_ttl":  obj.IdleDeleteTtl,
		"idle_start_time":  obj.IdleStartTime,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigMetastoreConfig(o interface{}) *dataproc.ClusterClusterConfigMetastoreConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigMetastoreConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigMetastoreConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigMetastoreConfig{
		DataprocMetastoreService: dcl.String(obj["dataproc_metastore_service"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigMetastoreConfig(obj *dataproc.ClusterClusterConfigMetastoreConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"dataproc_metastore_service": obj.DataprocMetastoreService,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigSecurityConfig(o interface{}) *dataproc.ClusterClusterConfigSecurityConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigSecurityConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigSecurityConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigSecurityConfig{
		KerberosConfig: expandDataprocWorkflowTemplateClusterClusterConfigSecurityConfigKerberosConfig(obj["kerberos_config"]),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigSecurityConfig(obj *dataproc.ClusterClusterConfigSecurityConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"kerberos_config": flattenDataprocWorkflowTemplateClusterClusterConfigSecurityConfigKerberosConfig(obj.KerberosConfig),
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigSecurityConfigKerberosConfig(o interface{}) *dataproc.ClusterClusterConfigSecurityConfigKerberosConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigSecurityConfigKerberosConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigSecurityConfigKerberosConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigSecurityConfigKerberosConfig{
		CrossRealmTrustAdminServer:    dcl.String(obj["cross_realm_trust_admin_server"].(string)),
		CrossRealmTrustKdc:            dcl.String(obj["cross_realm_trust_kdc"].(string)),
		CrossRealmTrustRealm:          dcl.String(obj["cross_realm_trust_realm"].(string)),
		CrossRealmTrustSharedPassword: dcl.String(obj["cross_realm_trust_shared_password"].(string)),
		EnableKerberos:                dcl.Bool(obj["enable_kerberos"].(bool)),
		KdcDbKey:                      dcl.String(obj["kdc_db_key"].(string)),
		KeyPassword:                   dcl.String(obj["key_password"].(string)),
		Keystore:                      dcl.String(obj["keystore"].(string)),
		KeystorePassword:              dcl.String(obj["keystore_password"].(string)),
		KmsKey:                        dcl.String(obj["kms_key"].(string)),
		Realm:                         dcl.String(obj["realm"].(string)),
		RootPrincipalPassword:         dcl.String(obj["root_principal_password"].(string)),
		TgtLifetimeHours:              dcl.Int64(int64(obj["tgt_lifetime_hours"].(int))),
		Truststore:                    dcl.String(obj["truststore"].(string)),
		TruststorePassword:            dcl.String(obj["truststore_password"].(string)),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigSecurityConfigKerberosConfig(obj *dataproc.ClusterClusterConfigSecurityConfigKerberosConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"cross_realm_trust_admin_server":    obj.CrossRealmTrustAdminServer,
		"cross_realm_trust_kdc":             obj.CrossRealmTrustKdc,
		"cross_realm_trust_realm":           obj.CrossRealmTrustRealm,
		"cross_realm_trust_shared_password": obj.CrossRealmTrustSharedPassword,
		"enable_kerberos":                   obj.EnableKerberos,
		"kdc_db_key":                        obj.KdcDbKey,
		"key_password":                      obj.KeyPassword,
		"keystore":                          obj.Keystore,
		"keystore_password":                 obj.KeystorePassword,
		"kms_key":                           obj.KmsKey,
		"realm":                             obj.Realm,
		"root_principal_password":           obj.RootPrincipalPassword,
		"tgt_lifetime_hours":                obj.TgtLifetimeHours,
		"truststore":                        obj.Truststore,
		"truststore_password":               obj.TruststorePassword,
	}

	return []interface{}{transformed}

}

func expandDataprocWorkflowTemplateClusterClusterConfigSoftwareConfig(o interface{}) *dataproc.ClusterClusterConfigSoftwareConfig {
	if o == nil {
		return dataproc.EmptyClusterClusterConfigSoftwareConfig
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 {
		return dataproc.EmptyClusterClusterConfigSoftwareConfig
	}
	obj := objArr[0].(map[string]interface{})
	return &dataproc.ClusterClusterConfigSoftwareConfig{
		ImageVersion: dcl.String(obj["image_version"].(string)),
		Properties:   checkStringMap(obj["properties"]),
	}
}

func flattenDataprocWorkflowTemplateClusterClusterConfigSoftwareConfig(obj *dataproc.ClusterClusterConfigSoftwareConfig) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"image_version": obj.ImageVersion,
		"properties":    obj.Properties,
	}

	return []interface{}{transformed}

}
