package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/will7200/mda/da"
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
	return "", nil
}

// Implement the business logic of Change
func (md *stubMdaService) Change(ctx context.Context, id string, req da.DA) (message string, err error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return "", err
	}
	if err := md.db.Model(d).Update(md).Error; err != nil {
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
	log.Debugf("%+v", dd)
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
