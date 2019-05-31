package utils

import (
	"encoding/csv"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func WriteCsv(ctx *gin.Context, content []simplejson.Json, title []string, fileName string, headName []string) {
	fp, err := os.Create(fmt.Sprintf("%s_%d.csv", fileName, time.Now().Unix()))
	if err != nil {
		log.Error("create file error ", fileName, err)
		panic(err)
	}
	defer func() {
		err := fp.Close()
		if err != nil {
			log.Error("close file error ", fp.Name())
		}
		err = os.Remove(fp.Name())
		if err != nil {
			log.Error("file error: ", fileName, err)
		}
	}()
	csvWriter := csv.NewWriter(fp)
	records := [][]string{}
	head := []string{}
	for _, h := range headName {
		h, _ = UTF82GBK(h)
		head = append(head, h)
	}
	records = append(records, head)
	for _, c := range content {
		line := []string{}
		for _, t := range title {
			r, err := c.Get(t).String()
			r, _ = UTF82GBK(r)
			if err != nil {
				v, _ := c.Get(t).Int()
				r = strconv.Itoa(v)
				println(r)
			}
			line = append(line, r)
		}
		records = append(records, line)
	}
	err = csvWriter.WriteAll(records)
	if err != nil {
		log.Error("write csv error", err)
	}
	csvWriter.Flush()
	ctx.Header("Content-type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename= "+url.QueryEscape(fileName))
	ctx.File(fp.Name())
}

func UTF82GBK(src string) (string, error) {
	reader := transform.NewReader(strings.NewReader(src), simplifiedchinese.GBK.NewEncoder())
	if buf, err := ioutil.ReadAll(reader); err != nil {
		return "", err
	} else {
		return string(buf), nil
	}
}
