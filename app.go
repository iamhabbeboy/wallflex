package main

import (
	"context"
	"desktop/internal"
	"desktop/internal/api"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

var appConf = internal.AppConfig{}

type Conf struct {
	ImageCategory            string
	TotalImage               int
	Interval                 string
	DefaultPath              string
	Apikey                   string
	ScheduleDownloadInterval string
	HasAutoDownloadEnabled   bool
}

const APP_NAME = "wallflex"

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	menu := SetMenuItem(ctx, a)
	runtime.MenuSetApplicationMenu(ctx, menu)
	appConf.Init(fmt.Sprintf("$HOME/.%s", APP_NAME))

	home, _ := os.UserHomeDir()
	srcPath := "com.user.wallflex.plist"
	bin := "/usr/local/bin/wallflex_scheduler"
	destDir := filepath.Join(home, "Library/LaunchAgents")
	destPath := filepath.Join(destDir, filepath.Base(srcPath))

	if !hasFile("./" + srcPath) {
		println("Launch script does not exist")
		return
	}

	if !hasFile(bin) {
		println("Wallflex binary does not exist")
		return
	}

	if !hasFile(destPath) {
		if err := setupLaunchctl(srcPath, destPath); err != nil {
			println(err.Error())
		}
	}
}

func setupLaunchctl(srcPath, destPath string) error {
	cmd := exec.Command("cp", srcPath, destPath)
	if err := cmd.Run(); err != nil {
		return errors.New("unable to copy the launchctl agent script")
	}

	if err := exec.Command("launchctl", "load", destPath).Run(); err != nil {
		return errors.New("unable to load the launchctl agent script")
	}

	if err := exec.Command("launchctl", "start", destPath).Run(); err != nil {
		return errors.New("unable to start the launchctl agent script")
	}
	return nil
}

func hasFile(dir string) bool {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

func (a *App) GetDownloadedImages() ([]string, error) {
	selectedPath, err := appConf.Get("image.selected_abs_path")

	if err != nil {
		return nil, fmt.Errorf("path_not_found:Error getting selected path: %v", err)
	}
	path := selectedPath.(string)
	var fp string = path
	img, err := internal.GetAllFilesInDir(fp)
	if err != nil {
		return nil, fmt.Errorf("images_not_found:Images not found: %v", err)
	}

	return img, nil
}

func (a *App) SelectImageDir() []string {
	dir, err := OpenNativeDir(a.ctx)
	if err != nil {
		println(err.Error())
		return []string{""}
	}

	// store the path selected.
	appConf.Set("image.selected_abs_path", dir)

	imgs, err := internal.GetAllFilesInDir(dir)

	if err != nil {
		println(err)
		return []string{""}
	}

	return imgs
}

func (a *App) DownloadImages() error {
	apikey, _ := appConf.Get("api.unsplash_apikey")
	dp, _ := appConf.Get("image.selected_abs_path")
	tot, _ := appConf.Get("api.download_limit")
	cat, _ := appConf.Get("api.image_category")

	if apikey == nil || dp == nil {
		return fmt.Errorf("image path not set")
	}

	var ct int
	if tot == nil {
		ct = 10
	} else {
		ct = tot.(int)
	}

	var ccat string
	if cat == nil {
		ccat = "abstract"
	} else {
		ccat = cat.(string)
	}

	var imagePath = dp.(string)
	if strings.Contains(imagePath, "."+APP_NAME) {
		home, _ := os.UserHomeDir()
		fp := fmt.Sprintf("%s/.%s/images", home, APP_NAME)
		imagePath = fp
	}

	c := api.ImageConfig{
		Category:           ccat,
		TotalDownloadImage: ct,
		Path:               imagePath,
		Apikey:             apikey.(string),
	}

	err := deleteFilesWithPrefix(imagePath, APP_NAME+"_")

	if err != nil {
		return fmt.Errorf("error deleting images: %v ", err.Error())
	}
	if err := internal.FetchImages(c); err != nil {
		return err
	}

	return nil
}

func (a *App) SetWallpaper(path string) {
	internal.WallpaperEvent(path)
}

func (a *App) GetConfig() Conf {
	imgCat, _ := appConf.Get("api.image_category")
	totalImg, _ := appConf.Get("api.download_limit")
	intvl, _ := appConf.Get("image.interval")
	dp, _ := appConf.Get("image.selected_abs_path")
	akey, _ := appConf.Get("api.unsplash_apikey")
	autoDownload, _ := appConf.Get("image.auto_download")
	scheduleImageDownload, _ := appConf.Get("api.download_interval")
	var img, intv, d, scImgDown string
	var tot int
	if imgCat == nil {
		img = ""
	} else {
		img = imgCat.(string)
	}

	if totalImg == nil {
		tot = 0
	} else {
		tot = totalImg.(int)
	}

	if intvl == nil {
		intv = ""
	} else {
		intv = intvl.(string)
	}

	if scheduleImageDownload == nil {
		scImgDown = ""
	} else {
		scImgDown = scheduleImageDownload.(string)
	}

	if dp == nil {
		d = "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY" //TODO: hardcoding api key is bad, but what can i say... user can be funny, this will help restore the key even when deleted
	} else {
		d = dp.(string)
	}

	var k string
	if akey == nil {
		k = ""
	} else {
		k = akey.(string)
	}

	c := Conf{
		ImageCategory:            img,
		TotalImage:               tot,
		Interval:                 intv,
		DefaultPath:              d,
		Apikey:                   k,
		ScheduleDownloadInterval: scImgDown,
		HasAutoDownloadEnabled:   autoDownload.(bool),
	}

	return c
}

func (a *App) SetConfig(conf Conf) {
	appConf.Set("api.image_category", conf.ImageCategory)
	appConf.Set("api.download_limit", conf.TotalImage)
	appConf.Set("image.selected_abs_path", conf.DefaultPath)
	appConf.Set("image.interval", conf.Interval)
	appConf.Set("image.auto_download", conf.HasAutoDownloadEnabled)
	appConf.Set("api.download_interval", conf.ScheduleDownloadInterval)
}

func (a *App) OpenDirDialogWindow() string {
	dir, err := OpenNativeDir(a.ctx)
	if err != nil {
		println(err.Error())
		return ""
	}
	return dir
}

func (a *App) MessageDialog(m string) (string, error) {
	return MessageBox(a.ctx, m)
}

func deleteFilesWithPrefix(dir, prf string) error {
	imgs, err := internal.GetAllFilesInDir(dir)
	if err != nil {
		return err
	}

	for _, v := range imgs {
		if strings.Contains(v, prf) {
			// deletedFiles = append(deletedFiles, v)
			if err := os.Remove(v); err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *App) GetGradientImage(c []api.RGBA) error {
	if len(c) == 0 {
		return errors.New("no colors added")
	}

	err := internal.GenerateGradientImage(c[0], c[1])
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetHexToRGBA(color string) (api.RGBA, error) {
	fmt.Println(color)
	r, _ := internal.HexToRGBA(color)
	fmt.Println(r)
	return api.RGBA{}, nil
}

func (a *App) Testament() error {
	return fmt.Errorf("Testament...")
}

// func (ax *App) GGradient(color string) interface{} {
//
// 	r, g, b, a, err := internal.HexToRGBA(color)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	res := struct {
// 		R int `json:"r"`
// 		G int `json:"g"`
// 		B int `json:"b"`
// 		A int `json:"a"`
// 	}{
// 		R: r,
// 		G: g,
// 		B: b,
// 		A: a,
// 	}
// 	return res
// }

// https://gist.github.com/stupidbodo/0db61fa874213a31dc57 - replacement for cronjob
// https://gist.github.com/harubaru/f727cedacae336d1f7877c4bbe2196e1#model-overview
