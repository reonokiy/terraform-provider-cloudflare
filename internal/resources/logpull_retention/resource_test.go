package logpull_retention_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stainless-sdks/cloudflare-terraform/internal/acctest"
	"github.com/stainless-sdks/cloudflare-terraform/internal/consts"
	"github.com/stainless-sdks/cloudflare-terraform/internal/utils"
)

func TestAccLogpullRetentionSetStatus(t *testing.T) {
	// Temporarily unset CLOUDFLARE_API_TOKEN if it is set as the Logpull
	// service is throwing authentication errors despite it being marked as
	// available.
	if os.Getenv("CLOUDFLARE_API_TOKEN") != "" {
		t.Setenv("CLOUDFLARE_API_TOKEN", "")
	}

	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_logpull_retention." + rnd
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testLogpullRetentionSetConfig(rnd, zoneID, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, consts.ZoneIDSchemaKey, zoneID),
					resource.TestCheckResourceAttr(name, "enabled", "false"),
				),
			},
		},
	})
}

func testLogpullRetentionSetConfig(id, zoneID, enabled string) string {
	return fmt.Sprintf(`
  resource "cloudflare_logpull_retention" "%[1]s" {
    zone_id = "%[2]s"
	  enabled = "%[3]s"
  }`, id, zoneID, enabled)
}
