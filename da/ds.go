package da

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type DA struct {
	ID          string `gorm:"primary_key"`
	Location    string
	URL         string
	Frequency   string
	Owner       string
	Enabled     bool
	Parameters  Metadata `sql:"Type:bytea"`
	Startdate   *time.Time
	Currentdate *time.Time
}
type Stats struct {
	Session string `gorm:"primary_key"`
	ID      string `gorm:"ForeignKey:ID;AssociationForeignKey:ID"`
	Success bool
	Error   string
	RanAt   *time.Time
}
type Metadata map[string]string

func newStats(id string) *Stats {
	t := time.Now()
	return &Stats{ID: id, RanAt: &t,
		Session: uuid.NewV4().String(), Success: true}
}
func (d Metadata) Value() (driver.Value, error) {
	if d == nil {
		return nil, nil
	}
	dd, err := json.Marshal(d)
	return dd, err
}

func (d *Metadata) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		var i interface{}
		err := json.Unmarshal(src, &i)
		if err != nil {
			return err
		}
		var ok bool
		*d, ok = i.(map[string]string)
		if !ok {
			return errors.New("Type assertion .(map[string]string) failed.")
		}

		return nil
	case string:
		var s map[string]string
		//d = make(map[string]string)
		err := json.Unmarshal([]byte(src), &d)
		if err != nil {
			return err
		}
		*d = s
		return nil
	case nil:
		return nil
	}
	return fmt.Errorf("Metadate: cannot convert %T to Metadata", src)
}

func (d *DA) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	if d.Currentdate == nil {
		scope.SetColumn("Currentdate", time.Time{})
	}
	return nil
}
func (d *DA) GetDA() DA {
	return *d
}
func CreateDatabaseTables(db *gorm.DB) {
	db.AutoMigrate(&DA{}, &Stats{})
}
