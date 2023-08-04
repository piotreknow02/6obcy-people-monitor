package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/corpix/uarand"
	"github.com/glebarez/sqlite"
	"github.com/piotreknow02/6obcy-people-monitor/server/models"
	"gorm.io/gorm"
)

var url string = "https://6obcy.org/ajax/online"



func main() {
	// Check environment

	var outfile = flag.String("out", "./6obcy.db", "output database file")
	if _, err := os.Stat(*outfile); errors.Is(err, os.ErrNotExist) {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		fmt.Errorf("DB file does not exist. Creaing %s", path.Join(cwd, *outfile))
		_, err = os.Create(*outfile)
		if err != nil {
			panic(err)
		}
	}

	// Connect to database
	db, err := gorm.Open(sqlite.Open(*outfile), &gorm.Config{})
	if err != nil {
		panic(err)
	}	
	db.AutoMigrate(
		&models.Log{},
	)

	// Http Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", uarand.GetRandom())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	scount, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	// Save to database
	count, err := strconv.Atoi(string(scount))
	if err != nil {
		panic(err)
	}
	newLog := models.Log{Count: uint16(count)}
	db.Create(&newLog)
	
	fmt.Printf("%d people at time %s\n", count, time.Now().Format(time.RFC850))
}
