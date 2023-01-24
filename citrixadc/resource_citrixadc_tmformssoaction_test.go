/*
Copyright 2016 Citrix Systems, Inc

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package citrixadc

import (
	"fmt"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccTmformssoaction_basic = `

	resource "citrixadc_tmformssoaction" "tf_tmformssoaction" {
		name           = "my_formsso_action"
		actionurl      = "/logon.php"
		userfield      = "loginID"
		passwdfield    = "passwd"
		ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
	}
`
const testAccTmformssoaction_update = `

	resource "citrixadc_tmformssoaction" "tf_tmformssoaction" {
		name           = "my_formsso_action"
		actionurl      = "/main/logon.php"
		userfield      = "loginID2"
		passwdfield    = "passwd2"
		ssosuccessrule = "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"
	}
`

func TestAccTmformssoaction_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTmformssoactionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTmformssoaction_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_tmformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "name", "my_formsso_action"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "actionurl", "/logon.php"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "userfield", "loginID"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "passwdfield", "passwd"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "ssosuccessrule", "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"),
				),
			},
			resource.TestStep{
				Config: testAccTmformssoaction_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmformssoactionExist("citrixadc_tmformssoaction.tf_tmformssoaction", nil),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "name", "my_formsso_action"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "actionurl", "/main/logon.php"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "userfield", "loginID2"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "passwdfield", "passwd2"),
					resource.TestCheckResourceAttr("citrixadc_tmformssoaction.tf_tmformssoaction", "ssosuccessrule", "HTTP.RES.HEADER(\"Set-Cookie\").CONTAINS(\"LogonID\")"),
				),
			},
		},
	})
}

func testAccCheckTmformssoactionExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmformssoaction name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Tmformssoaction.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmformssoaction %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmformssoactionDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmformssoaction" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Tmformssoaction.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmformssoaction %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
