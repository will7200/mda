package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/will7200/mjs/job"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/will7200/mda/da"
	"github.com/will7200/mjs/apischeduler"
	"github.com/will7200/mjs/apischeduler/grpc/pb"
)

// Implement yor service methods methods.
// e.x: Foo(ctx context.Context,s string)(rs string, err error)
type MdaService interface {
	//METHODS: POST
	//PATH: /
	Add(ctx context.Context, req da.DA) (id string, err error)
	//METHODS: POST
	//PATH: /start/{id}
	Start(ctx context.Context, id string) (message string, err error)
	//METHODS: POST
	//PATH: /remove/{id}
	Remove(ctx context.Context, id string) (message string, err error)
	//METHODS: PUT
	//PATH: /change/{id}
	Change(ctx context.Context, id string, req da.DA) (message string, err error)
	//METHODS: GET
	//PATH: /{id}
	Get(ctx context.Context, id string) (result *da.DA, err error)
	//METHODS: GET
	//PATH: /
	List(ctx context.Context) (results *[]da.DA, err error)
	//METHODS: POST
	//PATH: /enable
	Enable(ctx context.Context, id string) (message string, err error)
	//METHODS: POST
	//PATH: /disable
	Disable(ctx context.Context, id string) (message string, err error)
	AddToSchedular(ctx context.Context, id string) error
}
type stubMdaService struct {
	db *gorm.DB
	da da.Downloader
}

var (
	ErrInvalidLocation = errors.New("Location is Invalid view supported Providers")
	ErrDaDNE           = errors.New("DA record does not exist")
	ErrDAUATS          = errors.New("Unable to save new Request")
)

// Get a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New(db *gorm.DB, d da.Downloader) (s MdaService) {
	s = &stubMdaService{db, d}
	return s
}

// Implement the business logic of Add
func (md *stubMdaService) Add(ctx context.Context, req da.DA) (id string, err error) {
	if req.Startdate == nil || req.Startdate.IsZero() {
		return id, fmt.Errorf("Start time cannot be left blank")
	}
	if req.URL == "" {
		return id, fmt.Errorf("URL IS Required")
	}
	if err := md.db.Create(&req).Error; err != nil {
		err = fmt.Errorf("Error: %s;\nDatabaseError:%s", ErrDAUATS, err.Error())
		return id, err
	}
	id = req.ID
	log.Debugf("%+v", req)
	err = md.AddToSchedular(ctx, id)
	if err != nil {
		log.Debug("DA created but could not be added to remote schedular for reason %s", err.Error())
		return id, nil
	}
	return id, err
}

// Implement the business logic of Start
func (md *stubMdaService) Start(ctx context.Context, id string) (message string, err error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return "", err
	}

	err = md.da.Add(d)
	if err != nil {
		return "", err
	}
	message = fmt.Sprintf("DA with id %s has been started", d.ID)
	return message, nil
}

// Implement the business logic of Remove
func (md *stubMdaService) Remove(ctx context.Context, id string) (message string, err error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if err := md.db.Delete(d).Error; err != nil {
		err = fmt.Errorf("Unable to delete from database\nError:%s", err.Error())
		return "", err
	}
	message = fmt.Sprintf("DA with id %s has been removed", d.ID)
	return message, nil
}

// Implement the business logic of Change
func (md *stubMdaService) Change(ctx context.Context, id string, req da.DA) (message string, err error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return "", err
	}
	log.Debugf("%+v", req.Parameters)
	if err := md.db.Model(d).Update(req).Error; err != nil {
		err = fmt.Errorf("Cannot Update record with id %s;Database Error:%s", id, err.Error())
		return "", err
	}
	//TODO : MAYBE IMPLEMENT TO GET THE AMOUNT OF FIELDS CHANGED
	message = fmt.Sprintf("DA with id %s has changed", d.ID)
	return message, err
}

// Implement the business logic of Get
func (md *stubMdaService) Get(ctx context.Context, id string) (result *da.DA, err error) {
	dd := &da.DA{}
	if err := md.db.Where(da.DA{ID: id}).First(dd).Error; err != nil {
		err = fmt.Errorf("Error: %s;\nDatabaseError:%s", ErrDaDNE, err.Error())
		return nil, err
	}
	result = dd
	return result, err
}

// Implement the business logic of List
func (md *stubMdaService) List(ctx context.Context) (results *[]da.DA, err error) {
	d := &[]da.DA{}
	if err := md.db.Find(d).Error; err != nil {
		return nil, err
	}
	results = d
	return results, nil
}

// Implement the business logic of Enable
func (md *stubMdaService) Enable(ctx context.Context, id string) (message string, err error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if err := md.db.Model(d).Update(da.DA{Enabled: true}).Error; err != nil {
		err = fmt.Errorf("Cannot Enable record with id %s;Database Error:%s", id, err.Error())
		return "", err
	}
	message = fmt.Sprintf("DA with id %s has been enabled", d.ID)
	return message, err
}

// Implement the business logic of Disable
func (md *stubMdaService) Disable(ctx context.Context, id string) (message string, err error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if err := md.db.Model(d).Update(da.DA{Enabled: false}).Error; err != nil {
		err = fmt.Errorf("Cannot Disable record with id %s;Database Error:%s", id, err.Error())
		return "", err
	}
	message = fmt.Sprintf("DA with id %s has been disabled", d.ID)
	return message, err
}

func (md *stubMdaService) AddToSchedular(ctx context.Context, id string) error {
	d, err := md.Get(ctx, id)
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(viper.Get("remote_schedular_rpc").(string), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := pb.NewAPISchedulerClient(conn)
	//rr, errr := c.Start(context.Background(), &pb.StartRequest{})
	//fmt.Printf("Error %+v\n", errr)
	// Contact the server and print out its response.
	ctxx := context.Background()
	ctxx = metadata.NewContext(ctx,
		metadata.Pairs(apischeduler.JobUniqueness, "UNIQUE"))
	_, err = c.Add(ctxx, &pb.AddRequest{Reqjob: &pb.Job{
		Name:        fmt.Sprintf("%s", d.ID),
		Command:     []string{"curl", "--request", "POST", fmt.Sprintf("%s/mda/start/%s", viper.Get("local_url").(string), d.ID)},
		Schedule:    fmt.Sprintf("R/%s/P1W", time.Now().Add(time.Second*10).UTC().Format(job.RFC3339WithoutTimezone)),
		Application: "MDA",
		Domain:      "Local Area",
	}})
	if err != nil {
		return err
	}
	return nil
}
