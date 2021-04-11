package app

import (
	"github.com/go-zepto/zepto/plugins/linkeradmin"
	"github.com/go-zepto/zepto/plugins/linkeradmin/fields"
)

func (a *App) setupAdmin() {
	admin := linkeradmin.NewGuesserAdmin(a.Linker)

	admin.Menu.Links = []linkeradmin.MenuLink{
		{
			Icon:               "widgets",
			Label:              "Environments",
			LinkToResourceName: "Environment",
		},
		{
			Icon:               "settings",
			Label:              "Config",
			LinkToResourceName: "Config",
		},
	}

	// Environment
	environment := admin.Resource("Environment")
	environment.Update.
		AddInput(fields.Input(fields.NewReferenceListInput("Application", "environment_id", &fields.ReferenceListInputOptions{})))

	// Application
	application := admin.Resource("Application")
	application.Update.
		AddInput(fields.Input(fields.NewReferenceListInput("ApplicationEnvVar", "application_id", &fields.ReferenceListInputOptions{})))

	// Application Env Var
	applicationEnvVar := admin.Resource("ApplicationEnvVar")
	applicationEnvVar.List.
		RemoveField("id").
		RemoveField("created_at").
		RemoveField("updated_at").
		AddField(fields.NewTextField("key", nil)).
		AddField(fields.NewTextField("value", nil))

	// Config
	admin.Resource("Config")

	a.Zepto.AddPlugin(linkeradmin.NewLinkerAdminPlugin(linkeradmin.Options{
		Admin: admin,
		Path:  "/admin",
	}))
}
