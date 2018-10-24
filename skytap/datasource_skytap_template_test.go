package skytap

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccDataSourceSkytapTemplate_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSkytapTemplateConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.skytap_template.foo", "id"),
				),
			},
		},
	})
}

func testAccDataSourceSkytapTemplateConfig_basic() string {
	return fmt.Sprintf(`
data "skytap_template" "foo" {
	name = "CentOS 6.10 Desktop Firstboot"
}

output "id" {
  value = "${data.skytap_template.foo.id}"
}`)
}

func TestAccDataSourceSkytapTemplate_RegexMostRecent(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSkytapTemplateConfig_regexMostRecent(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.skytap_template.foo", "id"),
					resource.TestCheckResourceAttr("data.skytap_template.foo", "name", "Advanced Import Appliance on Ubuntu 18.04.1"),
				),
			},
		},
	})
}

func testAccDataSourceSkytapTemplateConfig_regexMostRecent() string {
	return fmt.Sprintf(`
data "skytap_template" "foo" {
	name = "18.04"
    most_recent = true
}

output "id" {
  value = "${data.skytap_template.foo.id}"
}`)
}