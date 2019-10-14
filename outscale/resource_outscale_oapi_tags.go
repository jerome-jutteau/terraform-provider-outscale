package outscale

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-outscale/osc/fcu"
	"github.com/terraform-providers/terraform-provider-outscale/osc/oapi"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOutscaleOAPITags() *schema.Resource {
	return &schema.Resource{
		Create: resourceOutscaleOAPITagsCreate,
		Read:   resourceOutscaleOAPITagsRead,
		Delete: resourceOutscaleOAPITagsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: getOAPITagsSchema(),
	}
}

func resourceOutscaleOAPITagsCreate(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OAPI

	request := oapi.CreateTagsRequest{}

	tag, tagsOk := d.GetOk("tag")

	resourceIds, resourceIdsOk := d.GetOk("resource_ids")

	if tagsOk == false && resourceIdsOk == false {
		return fmt.Errorf("One tag and resource id, must be assigned")
	}

	if tagsOk {
		request.Tags = tagsOAPIFromMap(tag.(map[string]interface{}))
	}
	if resourceIdsOk {
		var rids []string
		sgs := resourceIds.(*schema.Set).List()
		for _, v := range sgs {
			str := v.(string)
			rids = append(rids, str)
		}

		request.ResourceIds = rids
	}

	err := resource.Retry(60*time.Second, func() *resource.RetryError {
		_, err := conn.POST_CreateTags(request)
		if err != nil {
			if strings.Contains(fmt.Sprint(err), ".NotFound") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	d.SetId(resource.UniqueId())

	return resourceOutscaleOAPITagsRead(d, meta)
}

func resourceOutscaleOAPITagsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).OAPI

	// Build up search parameters
	params := oapi.ReadTagsRequest{
		Filters: oapi.FiltersTag{},
	}

	tag, tagsOk := d.GetOk("tag")
	if tagsOk {
		tgs := tagsOAPIFromMap(tag.(map[string]interface{}))
		keys := make([]string, 0, len(tgs))
		values := make([]string, 0, len(tgs))
		for _, t := range tgs {
			keys = append(keys, t.Key)
			values = append(values, t.Value)
		}

		params.Filters.Keys = keys
		params.Filters.Values = values

	}

	resourceIds, resourceIdsOk := d.GetOk("resource_ids")
	if resourceIdsOk {
		var rids []string
		sgs := resourceIds.(*schema.Set).List()
		for _, v := range sgs {
			str := v.(string)
			rids = append(rids, str)
		}

		params.Filters.ResourceIds = rids
	}

	var resp *oapi.POST_ReadTagsResponses
	var err error

	err = resource.Retry(60*time.Second, func() *resource.RetryError {
		resp, err = conn.POST_ReadTags(params)
		return resource.RetryableError(err)
	})

	if err != nil {
		return err
	}

	tg := oapiTagsDescToList(resp.OK.Tags)
	err = d.Set("tags", tg)

	return err
}

func resourceOutscaleOAPITagsDelete(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*OutscaleClient).FCU

	request := &fcu.DeleteTagsInput{}

	tag, tagsOk := d.GetOk("tag")

	resourceIds, resourceIdsOk := d.GetOk("resource_ids")

	if tagsOk == false && resourceIdsOk == false {
		return fmt.Errorf("One tag and resource id, must be assigned")
	}

	if tagsOk {
		request.Tags = tagsFromMap(tag.(map[string]interface{}))
	}
	if resourceIdsOk {
		var rids []*string
		sgs := resourceIds.(*schema.Set).List()
		for _, v := range sgs {
			str := v.(string)
			rids = append(rids, aws.String(str))
		}

		request.Resources = rids
	}

	err := resource.Retry(60*time.Second, func() *resource.RetryError {
		_, err := conn.VM.DeleteTags(request)
		if err != nil {
			ec2err, ok := err.(awserr.Error)
			if ok && strings.Contains(ec2err.Code(), ".NotFound") {
				return resource.RetryableError(err) // retry
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func getOAPITagsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"resource_ids": {
			Type:     schema.TypeSet,
			Optional: true,
			ForceNew: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"tag": {
			Type:     schema.TypeMap,
			Optional: true,
			ForceNew: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key": {
						Type:     schema.TypeString,
						Computed: true,
						Optional: true,
					},
					"value": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
		},
		"tags": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"key": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"value": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"resource_id": {
						Type:     schema.TypeString,
						Computed: true,
					},
					"resource_type": {
						Type:     schema.TypeString,
						Computed: true,
					},
				},
			},
		},
	}
}

//TODO: OAPI
// func setOAPITags(conn *fcu.Client, d *schema.ResourceData) error {

// 	if d.HasChange("tag") {
// 		oraw, nraw := d.GetChange("tag")
// 		o := oraw.(map[string]interface{})
// 		n := nraw.(map[string]interface{})
// 		create, remove := diffTags(tagsFromMap(o), tagsFromMap(n))

// 		// Set tag
// 		if len(remove) > 0 {
// 			err := resource.Retry(60*time.Second, func() *resource.RetryError {
// 				log.Printf("[DEBUG] Removing tag: %#v from %s", remove, d.Id())
// 				_, err := conn.VM.DeleteTags(&fcu.DeleteTagsInput{
// 					Resources: []*string{aws.String(d.Id())},
// 					Tags:      remove,
// 				})
// 				if err != nil {
// 					ec2err, ok := err.(awserr.Error)
// 					if ok && strings.Contains(ec2err.Code(), ".NotFound") {
// 						return resource.RetryableError(err) // retry
// 					}
// 					return resource.NonRetryableError(err)
// 				}
// 				return nil
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		if len(create) > 0 {
// 			err := resource.Retry(60*time.Second, func() *resource.RetryError {
// 				log.Printf("[DEBUG] Creating tag: %v for %s", create, d.Id())
// 				_, err := conn.VM.CreateTags(&fcu.CreateTagsInput{
// 					Resources: []*string{aws.String(d.Id())},
// 					Tags:      create,
// 				})
// 				if err != nil {
// 					ec2err, ok := err.(awserr.Error)
// 					if ok && strings.Contains(ec2err.Code(), ".NotFound") {
// 						return resource.RetryableError(err) // retry
// 					}
// 					return resource.NonRetryableError(err)
// 				}
// 				return nil
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

func oapiTagsDescToList(ts []oapi.Tag) []map[string]string {
	result := make([]map[string]string, len(ts))
	for k, t := range ts {
		if !oapiTagDescIgnored(&t) {
			r := map[string]string{}
			r["load_balancer_name"] = t.Key
			r["value"] = t.Value
			r["resource_id"] = t.ResourceId
			r["resource_type"] = t.ResourceType

			result[k] = r
		}
	}

	return result
}

func oapiTagDescIgnored(t *oapi.Tag) bool {
	filter := []string{"^outscale:"}
	for _, v := range filter {
		if r, _ := regexp.MatchString(v, t.Key); r == true {
			return true
		}
	}
	return false
}

func assignOapiTags(tag []interface{}, resourceID string, conn *oapi.Client) error {
	request := oapi.CreateTagsRequest{}
	request.Tags = tagsOAPIFromSliceMap(tag)
	request.ResourceIds = []string{resourceID}
	err := resource.Retry(60*time.Second, func() *resource.RetryError {
		_, err := conn.POST_CreateTags(request)
		if err != nil {
			if strings.Contains(fmt.Sprint(err), ".NotFound") {
				return resource.RetryableError(err)
			}
			return resource.NonRetryableError(err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// diffOAPITags takes our tag locally and the ones remotely and returns
// the set of tag that must be created, and the set of tag that must
// be destroyed.
func diffOAPITags(oldTags, newTags []oapi.ResourceTag) ([]oapi.ResourceTag, []oapi.ResourceTag) {
	// First, we're creating everything we have
	create := make(map[string]interface{})
	for _, t := range newTags {
		create[t.Key] = t.Value
	}

	// Build the list of what to remove
	var remove []oapi.ResourceTag
	for _, t := range oldTags {
		old, ok := create[t.Key]
		if !ok || old != t.Value {
			remove = append(remove, t)
		}
	}

	return tagsOAPIFromMap(create), remove
}

func tagsOAPIFromMap(m map[string]interface{}) []oapi.ResourceTag {
	result := make([]oapi.ResourceTag, 0, len(m))
	for k, v := range m {
		t := oapi.ResourceTag{
			Key:   k,
			Value: v.(string),
		}
		result = append(result, t)
	}

	return result
}

func tagsOAPIFromSliceMap(m []interface{}) []oapi.ResourceTag {
	result := make([]oapi.ResourceTag, 0, len(m))
	for _, v := range m {
		tag := v.(map[string]interface{})
		t := oapi.ResourceTag{
			Key:   tag["key"].(string),
			Value: tag["value"].(string),
		}
		result = append(result, t)
	}

	return result
}

// tagsOAPI	ToMap turns the list of tag into a map.
func tagsOAPIToMap(ts []oapi.ResourceTag) []map[string]string {
	result := make([]map[string]string, len(ts))
	if len(ts) > 0 {
		for k, t := range ts {
			tag := make(map[string]string)
			tag["key"] = t.Key
			tag["value"] = t.Value
			result[k] = tag
		}
	} else {
		result = make([]map[string]string, 0)
	}

	return result
}

func tagsOAPIToMapString(ts []oapi.ResourceTag) map[string]string {
	tags := make(map[string]string)
	if len(ts) > 0 {
		for _, t := range ts {
			tags[t.Key] = t.Value
		}
	}
	return tags
}

func tagsOAPISchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
		Computed: true,
		ForceNew: true,
	}
}

func tagsListOAPISchema() *schema.Schema {
	return &schema.Schema{
		Type: schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
				"value": {
					Type:     schema.TypeString,
					Computed: true,
					Optional: true,
				},
			},
		},
		Computed: true,
		Optional: true,
	}
}

func tagsListOAPISchemaForceNew() *schema.Schema {
	tagsSchema := tagsListOAPISchema()
	tagsSchema.ForceNew = true
	return tagsSchema
}

func tagsOAPIListSchemaComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"value": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	}
}

func tagsOAPISchemaComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Computed: true,
	}
}

func setOAPITags(conn *oapi.Client, d *schema.ResourceData) error {

	if d.HasChange("tags") {
		oraw, nraw := d.GetChange("tags")
		o := oraw.([]interface{})
		n := nraw.([]interface{})
		create, remove := diffOAPITags(tagsOAPIFromSliceMap(o), tagsOAPIFromSliceMap(n))

		// Set tag
		if len(remove) > 0 {
			err := resource.Retry(60*time.Second, func() *resource.RetryError {
				_, err := conn.POST_DeleteTags(oapi.DeleteTagsRequest{
					ResourceIds: []string{d.Id()},
					Tags:        remove,
				})
				if err != nil {
					if strings.Contains(fmt.Sprint(err), ".NotFound") {
						return resource.RetryableError(err) // retry
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		if len(create) > 0 {
			err := resource.Retry(60*time.Second, func() *resource.RetryError {
				_, err := conn.POST_CreateTags(oapi.CreateTagsRequest{
					ResourceIds: []string{d.Id()},
					Tags:        create,
				})
				if err != nil {
					if strings.Contains(fmt.Sprint(err), ".NotFound") {
						return resource.RetryableError(err) // retry
					}
					return resource.NonRetryableError(err)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
