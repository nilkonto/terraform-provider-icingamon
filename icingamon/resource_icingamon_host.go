package icingamon

import (
	"fmt"

	"github.com/astaxie/flatmap"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nilkonto/icingamon/iapi"
)

func resourceIcinga2Host() *schema.Resource {

	return &schema.Resource{
		Create: resourceIcinga2HostCreate,
		Update: resourceIcinga2HostUpdate,
		Read:   resourceIcinga2HostRead,
		Delete: resourceIcinga2HostDelete,
		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Hostname",
				ForceNew:    true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"check_command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"vars": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: false,
			},
			"templates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceIcinga2HostCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	hostname := d.Get("hostname").(string)
	address := d.Get("address").(string)
	zone := d.Get("zone").(string)
	var checkCommand string

	//checkCommand := d.Get("check_command").(string)
	if d.Get("check_command") != nil {
		checkCommand = d.Get("check_command").(string)
	} else {
		checkCommand = "hostalive"
	}

	//var vars map[string]string

	//if d.Get("vars") != nil {
	vars, err := flatmap.Flatten(d.Get("vars").(map[string]interface{}))

	//Normalize from map[string]interface{} to map[string]string
	// iterator := d.Get("vars").(map[string]interface{})
	// for key, value := range iterator {
	// 	vars[key] = value.(string)
	// }
	//}

	templates := make([]string, len(d.Get("templates").([]interface{})))
	for i, v := range d.Get("templates").([]interface{}) {
		templates[i] = v.(string)
	}

	// Call CreateHost with normalized data
	hosts, err := client.CreateHost(hostname, address, zone, checkCommand, vars, templates)
	if err != nil {
		return err
	}

	found := false
	for _, host := range hosts {
		if host.Name == hostname {
			d.SetId(hostname)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Failed to Create Host %s : %s", hostname, err)
	}

	return nil
}

func resourceIcinga2HostUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	hostname := d.Get("hostname").(string)
	address := d.Get("address").(string)
	zone := d.Get("zone").(string)
	var checkCommand string

	//checkCommand := d.Get("check_command").(string)
	if d.Get("check_command") != nil {
		checkCommand = d.Get("check_command").(string)
	} else {
		checkCommand = "hostalive"
	}

	//var vars map[string]string

	//if d.Get("vars") != nil {
	vars, err := flatmap.Flatten(d.Get("vars").(map[string]interface{}))

	//Normalize from map[string]interface{} to map[string]string
	// iterator := d.Get("vars").(map[string]interface{})
	// for key, value := range iterator {
	// 	vars[key] = value.(string)
	// }
	//}

	templates := make([]string, len(d.Get("templates").([]interface{})))
	for i, v := range d.Get("templates").([]interface{}) {
		templates[i] = v.(string)
	}

	// Call CreateHost with normalized data
	hosts, err := client.UpdateHost(hostname, address, zone, checkCommand, vars, templates)
	if err != nil {
		return err
	}

	found := false
	for _, host := range hosts {
		if host.Name == hostname {
			d.SetId(hostname)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("Failed to Create Host %s : %s", hostname, err)
	}

	return nil
}

func resourceIcinga2HostRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)

	hostname := d.Get("hostname").(string)

	hosts, err := client.GetHostTr(hostname)
	if err != nil {
		return err
	}

	found := false
	for _, host := range hosts {
		if host.Name == hostname {
			d.SetId(hostname)
			d.Set("hostname", host.Name)
			d.Set("address", host.Attrs.Address)
			d.Set("check_command", host.Attrs.CheckCommand)
			//d.Set("vars", host.Attrs.Vars)
			d.Set("zone", host.Attrs.Zone)
			found = true
		}
	}

	if !found {
		d.SetId("")
		// return fmt.Errorf("Failed to Read Host %s : %s", hostname, err)
	}

	return nil
}

func resourceIcinga2HostDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)
	hostname := d.Get("hostname").(string)

	err := client.DeleteHost(hostname)
	if err != nil {
		return fmt.Errorf("Failed to Delete Host %s : %s", hostname, err)
	}
	return nil
}
