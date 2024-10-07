package device

import (
	"context"
	"fmt"

	"github.com/apenella/go-ansible/v2/pkg/execute"
	"github.com/apenella/go-ansible/v2/pkg/playbook"
)

type ciscoRouter struct {
	PlaybookDIR string
	

}

func (c *ciscoRouter) AddStaticRoute(ctx context.Context) {
	role := fmt.Sprintf("%s/add_static_route.router.yml", c.PlaybookDIR)

	playbookCMD := playbook.NewAnsiblePlaybookCmd(
		playbook.WithPlaybooks(role),
		playbook.WithPlaybookOptions(&playbook.AnsiblePlaybookOptions{
			Inventory: "",
		}),
	)

	exec := execute.NewDefaultExecute(
		execute.WithCmd(playbookCMD),
	)

	if err := exec.Execute(context.TODO()); err != nil {
		fmt.Println("failed to exec")
		return
	}
}

func (c *ciscoRouter) Ping(ctx context.Context)             {}
func (c *ciscoRouter) AddACL(ctx context.Context)           {}
func (c *ciscoRouter) ShowInterface(ctx context.Context)    {}
func (c *ciscoRouter) EnableInterface(ctx context.Context)  {}
func (c *ciscoRouter) DisableInterface(ctx context.Context) {}
func (c *ciscoRouter) ShowVLAN(ctx context.Context)         {}
