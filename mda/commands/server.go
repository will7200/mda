package commands

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/will7200/mda/da"
	"github.com/will7200/mdar/mda/endpoints"
	mdahttp "github.com/will7200/mdar/mda/http"
	"github.com/will7200/mdar/mda/service"
)

var (
	disableLiveReload bool
	renderToDisk      bool
	serverAppend      bool
	serverInterface   string
	port              int
	serverWatch       bool
	verbose           bool
	db                *gorm.DB
	err               error
)
var servercmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"run"},
	Short:   "Whip up a instance",
	RunE:    server,
}

func init() {
	servercmd.Flags().IntVarP(&port, "port", "p", 4004, "port on which to listen to")
	servercmd.Flags().BoolVar(&verbose, "verbose", false, "output log verbose")
	servercmd.Flags().String("dbname", "sqlite3", "database type")
	servercmd.Flags().String("connection", "./temp_db.db", "database connection string")
	//servercmd.Flags().Int("workers", 4, "amount of workers in pool")
	viper.BindPFlag("port", servercmd.Flags().Lookup("port"))
	viper.BindPFlag("database.dbname", servercmd.Flags().Lookup("dbname"))
	viper.BindPFlag("database.connection", servercmd.Flags().Lookup("connection"))
	viper.BindPFlag("interface.workers", servercmd.Flags().Lookup("workers"))
}
func server(cmd *cobra.Command, args []string) error {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
	var parsedPort string
	if port != 0 {
		parsedPort = fmt.Sprintf(":%d", port)
	} else {
		parsedPort = ":4004"
	}
	//Dispatch = &job.Dispatcher{}
	//Dispatch.StartDispatcher(viper.GetInt("interface.workers"))
	db, err = gorm.Open(viper.GetString("database.dbname"), viper.GetString("database.connection"))
	if err != nil {
		panic(fmt.Sprintf("failed to connect database \ntype %s with connection %s", viper.GetString("database.dbname"), viper.GetString("database.connection")))
	}
	//api.CreateDatabaseTables(db)
	db.LogMode(true)
	fmt.Println(db.AutoMigrate(&da.DA{}, &da.Stats{}).GetErrors())
	//Dispatch.SetPersistStorage(db)
	//fmt.Println(db.AutoMigrate(&job.Job{}, &job.JobStats{}).GetErrors())
	d := da.NewDownloader("G:\\music\\", db)
	svc := service.New(db, d)
	ep := endpoints.New(svc)
	r := mdahttp.NewHTTPHandler(ep)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		// p will contain regular expression is compatible with regular expression in Perl, Python, and other languages.
		// for instance the regular expression for path '/articles/{id}' will be '^/articles/(?P<v0>[^/]+)$'
		p, err := route.GetPathRegexp()
		if err != nil {
			return err
		}
		m, err := route.GetMethods()
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(m, ","), t, p)
		return nil
	})
	log.Infof("Starting Server on port %d", port)
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 7 * time.Second,
		Addr:         parsedPort,
		Handler:      r,
	}
	return server.ListenAndServe()
}