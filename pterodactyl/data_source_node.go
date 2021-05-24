package pterodactyl

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	croc "github.com/krickler/crocgodyl"
)

func dataSourceNode() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceNodeRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"scheme": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"behind_proxy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"maintenance_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"memory_overallocate": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"disk": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"disk_overallocate": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upload_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"daemon_listen": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"daemon_sftp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"daemon_base": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	panel := m.(*croc.AppConfig)

	nodeID := d.Get("id").(int)

	node, err := panel.GetNode(nodeID)
	if err != nil {
		return diag.FromErr(err)
	}

	attributes := node.Attributes
	if err := d.Set("attributes", attributes); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(attributes.ID))
	d.Set("uuid", attributes.Uuid)
	d.Set("public", attributes.Uuid)
	d.Set("name", attributes.Name)
	d.Set("description", attributes.Description)
	d.Set("location_id", attributes.LocationID)
	d.Set("fqdn", attributes.FQDN)
	d.Set("scheme", attributes.Scheme)
	d.Set("behind_proxy", attributes.BehindProxy)
	d.Set("maintenance_mode", attributes.MaintenanceMode)
	d.Set("memory", attributes.Memory)
	d.Set("memory_overallocate", attributes.MemoryOverAlloc)
	d.Set("disk", attributes.Disk)
	d.Set("disk_overallocate", attributes.DiskOverAlloc)
	d.Set("upload_size", attributes.UploadSize)
	d.Set("daemon_listen", attributes.DaemonListen)
	d.Set("daemon_sftp", attributes.DaemonSftp)
	d.Set("daemon_base", attributes.DaemonBase)
	d.Set("created_at", attributes.CreatedAt)
	d.Set("updated_at", attributes.UpdatedAt)

	return diag.Errorf("please specify a id, name or a selector to lookup the Server")
}
