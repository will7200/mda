package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/will7200/mda/da"
)

// Implement yor service methods methods.
// e.x: Foo(ctx context.Context,s string)(rs string, err error)
type MdaService interface {
	Add(ctx context.Context, mdd da.DA) (string, error)
	Start(ctx context.Context, id string) error
	Remove(ctx context.Context, id string) error
	Change(ctx context.Context, id string, mdd da.DA) error
	Get(ctx context.Context, id string) (d *da.DA, err error)
	List(ctx context.Context) (result *[]da.DA, e error)
	Enable(ctx context.Context, id string) error
	Disable(ctx context.Context, id string) error
	//METHODS: GET,PUT,DELETE
	Try(ctx context.Context, id string) error
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
func (md *stubMdaService) Add(ctx context.Context, mdd da.DA) (s0 string, e1 error) {
	if err := md.db.Create(&mdd).Error; err != nil {
		err = fmt.Errorf("Error: %s;\nDatabaseError:%s", ErrDAUATS, err.Error())
		return "", err
	}
	log.Debugf("%+v", md)
	s0 = mdd.ID
	return s0, e1
}

// Implement the business logic of Start
func (md *stubMdaService) Start(ctx context.Context, id string) (e0 error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return err
	}

	err = md.da.Add(d)
	if err != nil {
		return err
	}
	return e0
}

// Implement the business logic of Remove
func (md *stubMdaService) Remove(ctx context.Context, id string) (e0 error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return err
	}
	if err := md.db.Delete(d).Error; err != nil {
		e0 = fmt.Errorf("Unable to delete from database\nError:%s", err.Error())
	}
	return e0
}

// Implement the business logic of Change
func (md *stubMdaService) Change(ctx context.Context, id string, mdd da.DA) (e0 error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return err
	}
	if err := md.db.Model(d).Update(md).Error; err != nil {
		err = fmt.Errorf("Cannot Update record with id %s;Database Error:%s", id, err.Error())
		return err
	}
	return e0
}

// Implement the business logic of Get
func (md *stubMdaService) Get(ctx context.Context, id string) (d *da.DA, err error) {
	d = &da.DA{}
	if err = md.db.Where(da.DA{ID: id}).First(d).Error; err != nil {
		err = fmt.Errorf("Error: %s;\nDatabaseError:%s", ErrDaDNE, err.Error())
		return nil, err
	}
	log.Debugf("%+v", d)
	return d, err
}

// Implement the business logic of List
func (md *stubMdaService) List(ctx context.Context) (result *[]da.DA, e error) {
	result = &[]da.DA{}
	if err := md.db.Find(result).Error; err != nil {
		return nil, err
	}
	return result, e
}

// Implement the business logic of Enable
func (md *stubMdaService) Enable(ctx context.Context, id string) (e0 error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return err
	}
	if err := md.db.Model(d).Update(da.DA{Enabled: true}).Error; err != nil {
		err = fmt.Errorf("Cannot Enable record with id %s;Database Error:%s", id, err.Error())
		return err
	}
	return e0
}

// Implement the business logic of Disable
func (md *stubMdaService) Disable(ctx context.Context, id string) (e0 error) {
	d, err := md.Get(ctx, id)
	if err != nil {
		return err
	}
	if err := md.db.Model(d).Update(da.DA{Enabled: false}).Error; err != nil {
		err = fmt.Errorf("Cannot Disable record with id %s;Database Error:%s", id, err.Error())
		return err
	}
	return e0
}

// Implement the business logic of Try
func (md *stubMdaService) Try(ctx context.Context, id string) (e0 error) {
	return e0
}
