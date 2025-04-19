package main

import (
	"desktop/internal"
	"desktop/internal/api"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	conf := &internal.AppConfig{}
	conf.Init("$HOME/.wallflex")

	cint, _ := conf.Get("image.interval")
	imgs, _ := conf.Get("image.selected_abs_path")

	apikey, _ := conf.Get("api.unsplash_apikey")
	cat, _ := conf.Get("api.image_category")
	totimg, _ := conf.Get("api.download_limit")

	if apikey == nil || imgs == nil {
		log.Fatal("Image path not set")
	}

	var ccat, k string
	var timg int

	if cat == nil {
		ccat = "technology"
	} else {
		ccat = cat.(string)
	}

	if apikey == nil {
		log.Fatal("API key not found")
	} else {
		k = apikey.(string)
	}

	if totimg == nil {
		timg = 10
	} else {
		timg = totimg.(int)
	}

	cimgs := imgs.(string)

	tmx := getDuration(cint.(string))

	deskw := time.NewTicker(tmx)
	down := time.NewTicker(20 * time.Minute)
	//
	defer deskw.Stop()
	defer down.Stop()

	c := api.ImageConfig{
		Category:           ccat,
		TotalDownloadImage: timg,
		Apikey:             k,
		Path:               cimgs,
	}

	for {
		select {
		case t := <-deskw.C:
			scheduleSetDesktopWallpaper(t, cimgs)
		case s := <-down.C:
			scheduleDownloadImages(s, c)
		}
	}
}

func scheduleDownloadImages(t time.Time, c api.ImageConfig) error {
	println("------Downloading....")
	println(t.String())
	println("------Downloading....")
	fmt.Println(c)
	if !c.HasAutoDownloadEnabled {
		return nil
	}
	println("Downloading images...")
	j, _ := json.Marshal(c)
	println(string(j))
	if err := internal.FetchImages(c); err != nil {
		return err
	}
	return nil
}

func scheduleSetDesktopWallpaper(t time.Time, cnf string) error {
	println(t.String())
	if cnf == "" {
		log.Fatal("Image directory not set")
	}
	fp := cnf

	imgs := getImages(fp)

	random := rand.Intn(len(imgs))
	f := imgs[random]

	log.Println("Current Wallpaper Path: ", f)
	internal.WallpaperEvent(f)

	return nil
}

func getImages(path string) []string {
	var fp string = path
	img, err := internal.GetAllFilesInDir(fp)
	if err != nil {
		println(err.Error())
	}
	return img
}

func getDuration(tm string) time.Duration {
	if tm == "" {
		return 30 * time.Minute
	}

	var tmx time.Duration
	var seq rune
	var dur []rune

	rtm := []rune(tm)

	if tm != "" {
		indx := len(rtm) - 1
		seq = rtm[indx]
		dur = rtm[:indx]
	}

	cs := string(seq)
	var t time.Duration

	if cs == "m" {
		t = time.Minute
	} else if cs == "s" {
		t = time.Second
	} else if cs == "h" {
		t = time.Hour
	} else if cs == "w" {
		t = 7 * 24 * time.Hour
	}

	idur := string(dur)
	n, _ := strconv.Atoi(idur)

	tmx = time.Duration(n) * t

	return tmx
}
