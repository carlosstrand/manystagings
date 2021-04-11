package app

import "github.com/go-zepto/zepto/plugins/mailer"

// TODO: Add real mailer provider
type mailerStub struct{}

func (m *mailerStub) Init(opts *mailer.InitOptions) {}
func (m *mailerStub) SendFromHTML(html string, opts *mailer.SendOptions) error {
	return nil
}
func (m *mailerStub) SendFromTemplate(template string, opts *mailer.SendOptions) error {
	return nil
}

func (app *App) setupMailer() {
	app.Zepto.AddPlugin(mailer.NewMailerPlugin(mailer.Options{
		Mailer: &mailerStub{},
	}))
}
