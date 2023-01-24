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
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

const testAccRdpclientprofile_basic = `


resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile" {
	name              = "my_rdpclientprofile"
	rdpurloverride    = "ENABLE"
	redirectclipboard = "ENABLE"
	redirectdrives    = "ENABLE"
  }
  
`
const testAccRdpclientprofile_update = `


resource "citrixadc_rdpclientprofile" "tf_rdpclientprofile" {
	name              = "my_rdpclientprofile"
	rdpurloverride    = "DISABLE"
	redirectclipboard = "DISABLE"
	redirectdrives    = "DISABLE"
  }
  
`

func TestAccRdpclientprofile_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRdpclientprofileDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccRdpclientprofile_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "name", "my_rdpclientprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "rdpurloverride", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectclipboard", "ENABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectdrives", "ENABLE"),
				),
			},
			resource.TestStep{
				Config: testAccRdpclientprofile_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRdpclientprofileExist("citrixadc_rdpclientprofile.tf_rdpclientprofile", nil),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "name", "my_rdpclientprofile"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "rdpurloverride", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectclipboard", "DISABLE"),
					resource.TestCheckResourceAttr("citrixadc_rdpclientprofile.tf_rdpclientprofile", "redirectdrives", "DISABLE"),
				),
			},
		},
	})
}

func testAccCheckRdpclientprofileExist(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rdpclientprofile name is set")
		}

		if id != nil {
			if *id != "" && *id != rs.Primary.ID {
				return fmt.Errorf("Resource ID has changed!")
			}

			*id = rs.Primary.ID
		}

		nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client
		data, err := nsClient.FindResource("rdpclientprofile", rs.Primary.ID)

		if err != nil {
			return err
		}

		if data == nil {
			return fmt.Errorf("rdpclientprofile %s not found", n)
		}

		return nil
	}
}

func testAccCheckRdpclientprofileDestroy(s *terraform.State) error {
	nsClient := testAccProvider.Meta().(*NetScalerNitroClient).client

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "citrixadc_rdpclientprofile" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No name is set")
		}

		_, err := nsClient.FindResource("rdpclientprofile", rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("rdpclientprofile %s still exists", rs.Primary.ID)
		}

	}

	return nil
}
