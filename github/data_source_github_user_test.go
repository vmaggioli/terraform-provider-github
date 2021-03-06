package github

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccGithubUserDataSource_noMatchReturnsError(t *testing.T) {
	username := "admin"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckGithubUserDataSourceConfig(username),
				ExpectError: regexp.MustCompile(`Not Found`),
			},
		},
	})
}

func TestAccGithubUserDataSource_existing(t *testing.T) {
	username := "raphink"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGithubUserDataSourceConfig(username),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.github_user.test", "name"),
					resource.TestCheckResourceAttr("data.github_user.test", "id", "650430"),
					resource.TestCheckResourceAttr("data.github_user.test", "name", "Raphaël Pinson"),
				),
			},
		},
	})
}

func testAccCheckGithubUserDataSourceConfig(username string) string {
	return fmt.Sprintf(`
data "github_user" "test" {
  username = "%s"
}
`, username)
}
