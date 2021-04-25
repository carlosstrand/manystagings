package app

import (
	"github.com/carlosstrand/manystagings/consts"
	"github.com/carlosstrand/manystagings/core/orchestrator"
	"github.com/carlosstrand/manystagings/core/orchestrator/providers/kubernetes"
	"github.com/carlosstrand/manystagings/core/service"
	"github.com/go-zepto/zepto"
	"github.com/go-zepto/zepto/plugins/linker"
	"github.com/go-zepto/zepto/web"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type App struct {
	DB           *gorm.DB
	Orchestrator orchestrator.Orchestrator
	Zepto        *zepto.Zepto
	Linker       *linker.Linker
	Service      *service.Service
	apiRouter    *web.Router
}

type Options struct {
	DB           *gorm.DB
	Orchestrator orchestrator.Orchestrator
	Logger       *log.Logger
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
	if opts.Orchestrator == nil {
		opts.Orchestrator = kubernetes.NewKubernetesProvider(kubernetes.Options{
			LogLevel: logrus.DebugLevel,
		})
	}
	zapp := zepto.NewZepto(
		zopts...,
	)
	apiRouter := zapp.Router("/api")
	return &App{
		Zepto:        zapp,
		DB:           opts.DB,
		Orchestrator: opts.Orchestrator,
		apiRouter:    apiRouter,
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
	app.setupWebapp()
	app.Zepto.SetupHTTP("0.0.0.0:8000")
}

func (app *App) Start() {
	app.Init()
	app.Zepto.Start()
}
