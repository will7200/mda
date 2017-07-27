package da

import (
	"bufio"
	"errors"
	"os/exec"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var (
	pdefault          map[string]string
	youtubeDL         = []string{"youtube-dl"}
	queue             map[string]*DA
	ErrAlreadyInQueue = errors.New("DA is currently in queue, Please wait until finished")
	timeFormat        = "20060102"
)

func init() {
	pdefault = make(map[string]string)
	pdefault["-f"] = "mp4"
	pdefault["-x"] = ""
	pdefault["--audio-format"] = "m4a"
	pdefault["--audio-quality"] = "9"
	pdefault["--embed-thumbnail"] = ""
	queue = make(map[string]*DA)
}

type Downloader interface {
	YoutubeDL(url string, parameters Metadata, da *DA)
	Add(da *DA) error
}

type downloader struct {
	Home string
	p    map[string]string
	db   *gorm.DB
}

func NewDownloader(home string, db *gorm.DB) Downloader {
	def := make(map[string]string)
	for index, value := range pdefault {
		def[index] = value
	}
	def["-o"] = home + "%(playlist)s\\%(upload_date)s\\%(id)s__%(title)s.%(ext)s"
	return downloader{home, def, db}
}

func (d downloader) Add(da *DA) error {
	if _, ok := queue[da.ID]; ok {
		logrus.Debug("Not adding already in queue")
		return ErrAlreadyInQueue
	}
	queue[da.ID] = da
	go d.YoutubeDL(da.URL, da.Parameters, da)
	return nil
}
func (d downloader) YoutubeDL(url string, parameters Metadata, da *DA) {
	args := make([]string, 0)
	args = append(args, youtubeDL...)
	args = append(args, url)
	v := combineMap(d.p, parameters)
	for index, value := range v {
		if value != "" {
			args = append(args, index, value)
		} else {
			args = append(args, index)
		}
	}
	if da.Currentdate.Before(*da.Startdate) {
		args = append(args, "--dateafter", da.Startdate.Format(timeFormat))
	} else {
		args = append(args, "--dateafter", da.Currentdate.Format(timeFormat))
	}
	cmd := exec.Command(args[0], args[1:]...)
	logrus.Debug("Command executing with ", args)
	stdpipe, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Info("Cannot open pipe")
		return
	}
	stats := newStats(da.ID)
	done := make(chan error)
	notinrange := make(chan bool)

	err = cmd.Start()
	if err != nil {
		logrus.Info("Process exiting with error ", err)
		logrus.Info("Could not continue with job")
		logrus.Debug("This will not be logged")
		delete(queue, da.ID)
		return
	}
	//file, _ = os.Open("log_mjs.txt")
	fs := bufio.NewScanner(stdpipe)
	go func() {
		for fs.Scan() {
			t := fs.Text()
			//fmt.Println(t)
			if strings.Contains(t, "not in range") {
				notinrange <- true
			}
		}
	}()
	go func() { done <- cmd.Wait() }()
	select {
	case err := <-done:
		logrus.Info("Process exiting with error ", err)
		if err != nil {
			stats.Success = false
			stats.Error = err.Error()
		} else {
			*da.Currentdate = time.Now()
			d.db.Model(da).Update(da)
		}
	case _ = <-notinrange:
		cmd.Process.Kill()
		*da.Currentdate = time.Now()
		d.db.Model(da).Update(da)
	}

	d.db.Create(stats)
	delete(queue, da.ID)
}

func combineMap(a, b map[string]string) map[string]string {
	v := make(map[string]string)
	for index, value := range a {
		v[index] = value
	}
	for index, value := range b {
		v[index] = value
	}
	return v
}
