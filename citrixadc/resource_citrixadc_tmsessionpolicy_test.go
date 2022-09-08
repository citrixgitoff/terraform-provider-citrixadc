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
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"testing"
)

const testAccTmsessionpolicy_basic = `

	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "true"
		action = "tf_tmsessaction"
	}
`
const testAccTmsessionpolicy_update= `

	resource "citrixadc_tmsessionpolicy" "tf_tmsessionpolicy" {
		name   = "my_tmsession_policy"
		rule   = "false"
		action = "tf_sessionaction2"
	}
`

func TestAccTmsessionpolicy_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTmsessionpolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccTmsessionpolicy_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionpolicyExist("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "name", "my_tmsession_policy"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "rule", "true"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "action", "tf_tmsessaction"),
				),
			},
			resource.TestStep{
				Config: testAccTmsessionpolicy_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTmsessionpolicyExist("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", nil),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "name", "my_tmsession_policy"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "rule", "false"),
					resource.TestCheckResourceAttr("citrixadc_tmsessionpolicy.tf_tmsessionpolicy", "action", "tf_sessionaction2"),
				),
			},
		},
	})
}

func testAccCheckTmsessionpolicyExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No tmsessionpolicy name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource(service.Tmsessionpolicy.Type(), rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("tmsessionpolicy %s not found", n)
		}

		return nil
	}
}

func testAccCheckTmsessionpolicyDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_tmsessionpolicy" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource(service.Tmsessionpolicy.Type(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("tmsessionpolicy %s still exists", rs.Primary.ID)
		}

	}

	return nil
}