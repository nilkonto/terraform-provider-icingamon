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
				ForceNew: true,
			},
			"check_command": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_notifications": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"enable_active_checks": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"volatile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"notes_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"action_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
			},
			"check_period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceIcinga2HostCreate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)
	var enableNotifications string
	var checkCommand string
	var enableActiveChecks string
	var volatile string
	var notes string
	var notesURL string
	var actionURL string
	var checkPeriod string
	var address string

	hostname := d.Get("hostname").(string)

	address = d.Get("address").(string)

	zone := d.Get("zone").(string)

	if d.Get("check_command") != nil {

		checkCommand = d.Get("check_command").(string)

	} else {

		checkCommand = "hostalive"
	}

	vars, err := flatmap.Flatten(d.Get("vars").(map[string]interface{}))

	templates := make([]string, len(d.Get("templates").([]interface{})))

	for i, v := range d.Get("templates").([]interface{}) {
		templates[i] = v.(string)
	}

	if d.Get("enable_notifications") != nil {
		enableNotifications = d.Get("enable_notifications").(string)
	}

	if d.Get("enable_active_checks") != nil {
		enableActiveChecks = d.Get("enable_active_checks").(string)
	}

	if d.Get("volatile") != nil {
		volatile = d.Get("volatile").(string)
	}

	if d.Get("notes") != nil {
		notes = d.Get("notes").(string)
	}

	if d.Get("notes_url") != nil {
		notesURL = d.Get("notes_url").(string)
	}

	if d.Get("action_url") != nil {
		actionURL = d.Get("action_url").(string)
	}

	if d.Get("check_period") != nil {
		checkPeriod = d.Get("check_period").(string)
	}

	// Call CreateHost with normalized data
	hosts, errCreateHost := client.CreateHost(hostname, address, zone, checkCommand, checkPeriod, actionURL, notesURL, notes, volatile, enableActiveChecks, enableNotifications, vars, templates)

	if errCreateHost != nil {
		return errCreateHost
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

	// Set enable_notifications flag
	errSetNotifications := client.SetNotifications(hostname, enableNotifications)

	if errSetNotifications != nil {
		return errSetNotifications
	}

	return nil
}

func resourceIcinga2HostUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*iapi.Server)
	var enableNotifications string
	var checkCommand string
	var enableActiveChecks string
	var volatile string
	var notes string
	var notesURL string
	var actionURL string
	var checkPeriod string
	var address string

	hostname := d.Get("hostname").(string)

	address = d.Get("address").(string)

	zone := d.Get("zone").(string)

	if d.Get("check_command") != nil {
		checkCommand = d.Get("check_command").(string)
	} else {
		checkCommand = "hostalive"
	}

	vars, err := flatmap.Flatten(d.Get("vars").(map[string]interface{}))

	templates := make([]string, len(d.Get("templates").([]interface{})))
	for i, v := range d.Get("templates").([]interface{}) {
		templates[i] = v.(string)
	}

	if d.Get("enable_notifications") != nil {
		enableNotifications = d.Get("enable_notifications").(string)
	}

	if d.Get("enable_active_checks") != nil {
		enableActiveChecks = d.Get("enable_active_checks").(string)
	}

	if d.Get("volatile") != nil {
		volatile = d.Get("volatile").(string)
	}

	if d.Get("notes") != nil {
		notes = d.Get("notes").(string)
	}

	if d.Get("notes_url") != nil {
		notesURL = d.Get("notes_url").(string)
	}

	if d.Get("action_url") != nil {
		actionURL = d.Get("action_url").(string)
	}

	if d.Get("check_period") != nil {
		checkPeriod = d.Get("check_period").(string)
	}

	// Call UpdateHost with normalized data
	hosts, errUpdateHost := client.UpdateHost(hostname, address, zone, checkCommand, checkPeriod, actionURL, notesURL, notes, volatile, enableActiveChecks, enableNotifications, vars, templates)

	if errUpdateHost != nil {
		return errUpdateHost
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

	// Set enable_notifications flag
	errSetNotifications := client.SetNotifications(hostname, enableNotifications)

	if errSetNotifications != nil {
		return errSetNotifications
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
			// d.Set("vars", host.Attrs.Vars)
			d.Set("zone", host.Attrs.Zone)
			d.Set("enable_notifications", host.Attrs.EnableNotifications)
			d.Set("enable_active_checks", host.Attrs.EnableActiveChecks)
			d.Set("volatile", host.Attrs.Volatile)
			d.Set("notes", host.Attrs.Notes)
			d.Set("notes_url", host.Attrs.NotesURL)
			d.Set("action_url", host.Attrs.ActionURL)
			d.Set("check_period", host.Attrs.CheckPeriod)

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
