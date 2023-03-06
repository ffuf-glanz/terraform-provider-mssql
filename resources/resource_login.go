package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"terraform-provider-mssql/sql"
)

const usernameProp = "username"
const passwordProp = "password"
const typeProp = "usertype"
const connectionString = "connection_string"

// Login is the mssql_login resource
func Login() *schema.Resource {
	return &schema.Resource{
		Create: loginCreate,
		Read:   loginRead,
		Update: loginUpdate,
		Delete: loginDelete,

		Schema: map[string]*schema.Schema{
			usernameProp: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			typeProp: &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			passwordProp: &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			connectionString: &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func loginCreate(d *schema.ResourceData, meta interface{}) error {
	connector := meta.(sql.Connector)
	connector.ConnectionString = d.Get(connectionString).(string)
	username := d.Get(usernameProp).(string)
	password := d.Get(passwordProp).(string)
	usertype := d.Get(typeProp).(string)

	errDel := connector.DeleteLogin(username)
	if errDel != nil {
		return errDel
	}

	err := connector.CreateLogin(username, password, usertype)
	if err != nil {
		return err
	}

	d.SetId(username)

	return loginRead(d, meta)
}

func loginRead(d *schema.ResourceData, meta interface{}) error {
	connector := meta.(sql.Connector)
	connector.ConnectionString = d.Get(connectionString).(string)
	username := d.Id()

	login, err := connector.GetLogin(username)
	if err != nil {
		return err
	}
	if login == nil {
		d.SetId("")
	}

	return nil
}

func loginUpdate(d *schema.ResourceData, meta interface{}) error {
	connector := meta.(sql.Connector)
	connector.ConnectionString = d.Get(connectionString).(string)
	username := d.Get(usernameProp).(string)
	password := d.Get(passwordProp).(string)

	err := connector.UpdateLogin(username, password)
	if err != nil {
		return err
	}

	return loginRead(d, meta)
}

func loginDelete(d *schema.ResourceData, meta interface{}) error {
	connector := meta.(sql.Connector)
	connector.ConnectionString = d.Get(connectionString).(string)
	username := d.Id()

	return connector.DeleteLogin(username)
}
