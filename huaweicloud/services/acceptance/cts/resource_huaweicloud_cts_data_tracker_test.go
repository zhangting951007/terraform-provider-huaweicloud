package cts

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	cts "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/cts/v3/model"

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/config"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud/services/acceptance"
)

func getCTSDataTrackerResourceObj(conf *config.Config, state *terraform.ResourceState) (interface{}, error) {
	client, err := conf.HcCtsV3Client(acceptance.HW_REGION_NAME)
	if err != nil {
		return nil, fmt.Errorf("error creating CTS client: %s", err)
	}

	name := state.Primary.ID
	trackerType := cts.GetListTrackersRequestTrackerTypeEnum().DATA
	listOpts := &cts.ListTrackersRequest{
		TrackerName: &name,
		TrackerType: &trackerType,
	}

	response, err := client.ListTrackers(listOpts)
	if err != nil {
		return nil, fmt.Errorf("error retrieving CTS tracker: %s", err)
	}

	if response.Trackers == nil || len(*response.Trackers) == 0 {
		return nil, fmt.Errorf("can not find the CTS tracker %s", name)
	}

	allTrackers := *response.Trackers
	ctsTracker := allTrackers[0]

	return ctsTracker, nil
}

func TestAccCTSDataTracker_basic(t *testing.T) {
	var dataTracker cts.TrackerResponseBody
	rName := acceptance.RandomAccResourceNameWithDash()
	resourceName := "huaweicloud_cts_data_tracker.tracker"

	rc := acceptance.InitResourceCheck(
		resourceName,
		&dataTracker,
		getCTSDataTrackerResourceObj,
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acceptance.TestAccPreCheck(t) },
		ProviderFactories: acceptance.TestAccProviderFactories,
		CheckDestroy:      rc.CheckResourceDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccCTSDataTracker_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					rc.CheckResourceExists(),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "lts_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "transfer_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "type", "data"),
					resource.TestCheckResourceAttr(resourceName, "status", "enabled"),
					resource.TestCheckResourceAttr(resourceName, "data_operation.#", "2"),
					resource.TestCheckResourceAttrPair(resourceName, "data_bucket",
						"huaweicloud_obs_bucket.data_bucket", "bucket"),
				),
			},
			{
				Config: testAccCTSDataTracker_update(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "transfer_enabled", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "bucket_name",
						"huaweicloud_obs_bucket.trans_bucket", "bucket"),
					resource.TestCheckResourceAttr(resourceName, "file_prefix", "cts"),
					resource.TestCheckResourceAttr(resourceName, "obs_retention_period", "30"),
					resource.TestCheckResourceAttr(resourceName, "validate_file", "false"),
					resource.TestCheckResourceAttr(resourceName, "lts_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "status", "enabled"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCTSDataTracker_basic(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_obs_bucket" "data_bucket" {
  bucket = "%[1]s-data"
  acl    = "public-read"
}

resource "huaweicloud_cts_data_tracker" "tracker" {
  name        = "%[1]s"
  data_bucket = huaweicloud_obs_bucket.data_bucket.bucket
  lts_enabled = true
}
`, rName)
}

func testAccCTSDataTracker_update(rName string) string {
	return fmt.Sprintf(`
resource "huaweicloud_obs_bucket" "data_bucket" {
  bucket = "%[1]s-data"
  acl    = "public-read"
}

resource "huaweicloud_obs_bucket" "trans_bucket" {
  bucket        = "%[1]s-log"
  acl           = "private"
  force_destroy = true

  lifecycle {
    ignore_changes = [lifecycle_rule]
  }
}

resource "huaweicloud_cts_data_tracker" "tracker" {
  name                 = "%[1]s"
  data_bucket          = huaweicloud_obs_bucket.data_bucket.bucket
  bucket_name          = huaweicloud_obs_bucket.trans_bucket.bucket
  obs_retention_period = 30
  file_prefix          = "cts"
  validate_file        = false
  lts_enabled          = false
}
`, rName)
}
