package main

import (
	"github.com/feilfeilundfeil/terraform-provider-mssql/resources"
	"github.com/feilfeilundfeil/terraform-provider-mssql/sql"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const connectionString = "connection_string"

// Provider exposes the provider to Terraform
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			connectionString: {
				Type:     schema.TypeString,
				Required: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"mssql_login": resources.Login(),
		},
	}
	p.ConfigureFunc = configure(p)
	return p
}

func configure(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		connStr := d.Get(connectionString).(string)
		return sql.Connector{
			ConnectionString: connStr,
		}, nil
	}
}
