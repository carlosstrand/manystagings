package app

import (
	"github.com/carlosstrand/manystagings/consts"
	"github.com/carlosstrand/manystagings/core/service"
	"github.com/go-zepto/zepto"
	"github.com/go-zepto/zepto/plugins/linker"
	"github.com/go-zepto/zepto/web"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	DB        *gorm.DB
	Zepto     *zepto.Zepto
	Linker    *linker.Linker
	Service   *service.Service
	apiRouter *web.Router
}

type Options struct {
	DB     *gorm.DB
	Logger *log.Logger
}

func NewApp(opts Options) *App {
	if opts.DB == nil {
		db, err := CreateDB()
		if err != nil {
			panic(err)
		}
		opts.DB = db
	}
	zopts := []zepto.Option{
		zepto.Name("manystagings"),
		zepto.Version(consts.VERSION),
		zepto.Logger(log.New()),
	}
	if opts.Logger != nil {
		zopts = append(zopts, zepto.Logger(opts.Logger))
	}
	zapp := zepto.NewZepto(
		zopts...,
	)
	apiRouter := zapp.Router("/api")
	return &App{
		Zepto:     zapp,
		DB:        opts.DB,
		apiRouter: apiRouter,
	}
}

func (app *App) Init() {
	app.setupMailer()
	app.setupAuth()
	app.setupLinker()
	app.setupAdmin()
	app.setupService()
	app.setupControllers()
	app.bootstrap()
	app.Zepto.SetupHTTP("0.0.0.0:8000")
}

func (app *App) Start() {
	app.Init()
	app.Zepto.Start()
}
