package internal

import (
	"desktop/internal/api"
	"fmt"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
)

func GetImagesFromDir() []string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return []string{err.Error()}
	}
	const APP_NAME = ".picasa"
	base := fmt.Sprintf("%s/%s", homedir, APP_NAME)
	dir := path.Join(base, "/", "images")
	ent, err := os.ReadDir(dir)
	if err != nil {
		return []string{err.Error()}
	}
	images := []string{}
	for i, e := range ent {
		if i == 0 {
			continue
		}
		fpath := fmt.Sprintf("%s/%s", dir, e.Name())
		images = append(images, fpath)
	}
	return images

}

func FetchImages(conf api.ImageConfig) {
	apikey := conf.Apikey
	sp := conf.Path
	svc := api.NewImageDownload("unsplash", sp, apikey)
	err := svc.GetImages(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func WallpaperEvent(path string) {
	filepath := path
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", filepath))

	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
}

func GetAllFilesInDir(dir string) ([]string, error) {
	var images []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isImageFile(path) {
			if isImageFile(info.Name()) {
				images = append(images, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return images, nil
}

func isImageFile(file string) bool {
	extn := []string{".jpg", ".jpeg", ".png"}
	for _, ext := range extn {
		if strings.HasSuffix(strings.ToLower(file), ext) {
			return true
		}
	}
	return false
}

func GenerateGradientImage() error {
	w, h := 1900, 1200
	dc := gg.NewContext(w, h)
	stc := color.RGBA{255, 0, 0, 255}
	endc := color.RGBA{0, 0, 255, 255}

	for y := 0; y < h; y++ {
		t := float64(y) / float64(h)
		r := uint8((1-t)*float64(stc.R) + t*float64(endc.R))
		g := uint8((1-t)*float64(stc.G) + t*float64(endc.G))
		b := uint8((1-t)*float64(stc.B) + t*float64(endc.B))
		// a := uint8((1-t)*float64(stc.A) + t*float64(endc.A))

		dc.SetRGB(float64(r)/255, float64(g)/255, float64(b)/255)
		dc.DrawLine(0, float64(y), float64(w), float64(y))
		dc.Stroke()
	}

	f, err := os.Create("gradient.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	return png.Encode(f, dc.Image())
}

func HexToRGB(hex string) (int, int, int, error) {
	// Remove the '#' if present
	hex = strings.TrimPrefix(hex, "#")

	// Parse hex values
	r, err := strconv.ParseInt(hex[0:2], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}

	return int(r), int(g), int(b), nil
}
