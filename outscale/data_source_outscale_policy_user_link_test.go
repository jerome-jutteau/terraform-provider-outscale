package outscale

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/terraform-providers/terraform-provider-outscale/osc/eim"
)

func TestAccOutscaleDSPolicyUserLink_basic(t *testing.T) {
	var out eim.ListAttachedUserPoliciesOutput
	rName := acctest.RandString(10)
	policyName1 := fmt.Sprintf("test-policy-%s", acctest.RandString(10))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccOutscaleDSUserPolicyAttachConfig(rName, policyName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOutscaleDSPolicyUserLinkExists("data.outscale_policy_user_link.outscale_policy_user_link", 1, &out),
				),
			},
		},
	})
}

func testAccCheckOutscaleDSPolicyUserLinkExists(n string, c int, out *eim.ListAttachedUserPoliciesOutput) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No policy DS is set")
		}

		return nil
	}
}

func testAccOutscaleDSUserPolicyAttachConfig(rName, policyName string) string {
	return fmt.Sprintf(`
resource "outscale_user" "user" {
    user_name = "test-user-%s"
}

resource "outscale_policy" "policy" {
    policy_name = "%s"
    #description = "A test policy"
    policy_document = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "eim:ChangePassword"
      ],
      "Resource": "*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "outscale_policy_user_link" "test-attach" {
    user_name = "${outscale_user.user.user_name}"
    policy_arn = "${outscale_policy.policy.arn}"
}

data "outscale_policy_user_link" "outscale_policy_user_link" {
    user_name = "${outscale_user.user.user_name}"
}`, rName, policyName)
}
