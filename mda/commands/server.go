package commands

import (
	"context"
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
	"github.com/will7200/mda/mda/endpoints"
	mdahttp "github.com/will7200/mda/mda/http"
	"github.com/will7200/mda/mda/service"
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
	showHTTPDir       bool
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
	servercmd.Flags().String("homedir", "./mda/", "home directory to download into")
	servercmd.Flags().BoolVar(&showHTTPDir, "httpdir", false, "Output the http directory")
	//servercmd.Flags().Int("workers", 4, "amount of workers in pool")
	viper.BindPFlag("verbose", servercmd.Flags().Lookup("verbose"))
	viper.BindPFlag("interface.port", servercmd.Flags().Lookup("port"))
	viper.BindPFlag("database.dbname", servercmd.Flags().Lookup("dbname"))
	viper.BindPFlag("database.connection", servercmd.Flags().Lookup("connection"))
	viper.BindPFlag("interface.workers", servercmd.Flags().Lookup("workers"))
	viper.BindPFlag("interface.home", servercmd.Flags().Lookup("homedir"))
	viper.SetEnvPrefix("MDA") // will be uppercased automatically
	viper.BindEnv("verbose")
	viper.BindEnv("mjs_service_grpc")
	viper.BindEnv("acl_token")
	viper.BindEnv("consul_address")
}
func server(cmd *cobra.Command, args []string) error {
	verbose = viper.GetBool("verbose") || verbose
	if verbose {
		log.SetLevel(log.DebugLevel)
	}
	var parsedPort string
	if viper.GetInt("interface.port") != 0 {
		parsedPort = fmt.Sprintf(":%d", viper.GetInt("interface.port"))
	} else {
		parsedPort = ":4004"
	}
	db, err = gorm.Open(viper.GetString("database.dbname"), viper.GetString("database.connection"))
	if verbose {
		db.LogMode(true)
	}
	if err != nil {
		panic(fmt.Sprintf("failed to connect database \ntype %s with connection %s", viper.GetString("database.dbname"), viper.GetString("database.connection")))
	}
	if errors := db.AutoMigrate(&da.DA{}, &da.Stats{}).GetErrors(); len(errors) != 0 {
		fmt.Printf("Cound not auto migrate tables for reasons below %v", errors)
		fmt.Println()
		panic("Could not make/migrate tables")
	}
	d := da.NewDownloader(viper.GetString("interface.home"), db)
	svc := service.New(db, d)
	ep := endpoints.New(svc)
	r := mdahttp.NewHTTPHandler(ep)
	if verbose || showHTTPDir {
		showHTTPPaths(r)
	}
	log.Infof("Starting Server on port %d", viper.GetInt("interface.port"))
	go AddtoSchedular(db, &svc)
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 7 * time.Second,
		Addr:         parsedPort,
		Handler:      r,
	}
	return server.ListenAndServe()
}

func AddtoSchedular(db *gorm.DB, s *service.MdaService) {
	das := []da.DA{}
	db.Find(&das)
	ss := *s
	for _, da := range das {
		err := ss.AddToSchedular(context.Background(), da.ID)
		if err != nil && !strings.Contains(err.Error(), "Job already exists") {
			log.Error(err)
		}
	}
}

func showHTTPPaths(r *mux.Router) {
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
}
