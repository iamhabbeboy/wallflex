package internal

import (
	"desktop/internal/api"
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
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

func GenerateGradientImage(c1 api.RGBA, c2 api.RGBA) error {

	w, h := 1900, 1200
	dc := gg.NewContext(w, h)
	// stc := color.RGBA{255, 0, 0, 255}
	// endc := color.RGBA{0, 0, 255, 255}

	stc := color.RGBA{uint8(c1.R), uint8(c1.G), uint8(c1.B), uint8(c1.A)}
	endc := color.RGBA{uint8(c2.R), uint8(c2.G), uint8(c2.B), uint8(c2.A)}

	for y := 0; y < h; y++ {
		t := float64(y) / float64(h)
		r := uint8((1-t)*float64(stc.R) + t*float64(endc.R))
		g := uint8((1-t)*float64(stc.G) + t*float64(endc.G))
		b := uint8((1-t)*float64(stc.B) + t*float64(endc.B))
		// a := uint8((1-t)*float64(stc.A) + t*float64(endc.A))
		//
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

func HexToRGBA(hex string) (api.RGBA, error) {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	var r, g, b, a int

	switch len(hex) {
	case 6: // #RRGGBB
		r64, err := strconv.ParseInt(hex[0:2], 16, 8)
		if err != nil {
			return api.RGBA{}, err
		}
		g64, err := strconv.ParseInt(hex[2:4], 16, 8)
		if err != nil {
			return api.RGBA{}, err
		}
		b64, err := strconv.ParseInt(hex[4:6], 16, 8)
		if err != nil {
			return api.RGBA{}, err
		}

		r, g, b, a = int(r64), int(g64), int(b64), 255 // Default alpha to 255

		fmt.Println(r, g, b, a)
	case 8: // #RRGGBBAA
		r64, err := strconv.ParseInt(hex[0:2], 16, 8)
		if err != nil {
			return api.RGBA{}, err
		}
		g64, err := strconv.ParseInt(hex[2:4], 16, 8)
		if err != nil {
			return api.RGBA{}, err
		}
		b64, err := strconv.ParseInt(hex[4:6], 16, 8)
		if err != nil {
			return api.RGBA{}, err
		}
		a64, err := strconv.ParseInt(hex[6:8], 16, 8)
		if err != nil {
			return api.RGBA{}, err
		}

		r, g, b, a = int(r64), int(g64), int(b64), int(a64)
		fmt.Println(r, g, b, a)

	default:
		return api.RGBA{}, fmt.Errorf("invalid hex color format")
	}

	fmt.Println(r, g, b, a)

	res := api.RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}

	return res, nil
}
