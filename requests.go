package main

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/luopengift/framework"
	"github.com/luopengift/log"
)

func main() {
	framework.MainFunc(mainThread)
	framework.Run()
}

func mainThread(ctx context.Context) error {
	client := http.Client{}
	req, err := http.NewRequest("get", "http://dingyue.nosdn.127.net/qko87BAx0VO6XWgZA5WcWNgzYPBi=ZVzKbx6yH=8KggKY1502361463759compressflag.jpg", nil)
	//req, err := http.NewRequest("get", "https://t66y.com/htm_data/8/1810/3322488.html", nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Info("%s", string(bytes))
	return nil
}

// Download download file
func Download(name string) error {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}
