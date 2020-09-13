package logone

import (
	"time"

	"github.com/go.standard.project.layout/project.web/fiber/crud.mongoa/pkg/fmts"
)

func Error(org_id, org_name, org_unity, profile_env,
	system_id, hostname, app_name, message_id,
	parent_correlation_id, correlation_id,
	logged_user,
	caller_ip, service_method, service_parameters,
	service_returned, msg string) {
	fmts.Stdout(
		`{`, org_id, `}`, ` | `,
		`{`, org_name, `}`, ` | `,
		`{`, org_unity, `}`, ` | `,
		`{`, profile_env, `}`, ` | `,
		`{`, system_id, `}`, ` | `,
		`{`, hostname, `}`, ` | `,
		`{`, app_name, `}`, ` | `,
		`{`, message_id, `}`, ` | `,
		`{`, parent_correlation_id, `}`, ` | `,
		`{`, correlation_id, `}`, ` | `,
		`{`, time.Now().Format("2006-01-02 15:04:05"), `}`, ` | `,
		`{`, logged_user, `}`, ` | `,
		`{`, caller_ip, `}`, ` | `,
		`{`, service_method, `}`, ` | `,
		`{`, service_parameters, `}`, ` | `,
		`{`, service_returned, `}`, ` | `,
		`{`, msg, `}`, "\n",
	)
}
