package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	neturl "net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type UnleaseService struct {
	apikey string
	path   string
}

type Image struct {
	Urls struct {
		Raw  string `json:"raw"`
		Full string `json:"full"`
	} `json:"urls"`
	User struct {
		Links struct {
			Html string `json:"html"`
		} `json:"links"`
	} `json:"user"`
	Links struct {
		Download string `json:"download"`
	} `json:"links"`
}

type ImageConfig struct {
	Category               string
	TotalDownloadImage     int
	Apikey                 string
	Path                   string
	HasAutoDownloadEnabled bool
}

type RGBA struct {
	R int
	G int
	B int
	A int
}

const timeout = 60 * time.Second

func NewUnsplashService(apikey string, path string) *UnleaseService {
	return &UnleaseService{
		apikey: apikey,
		path:   path,
	}
}

func (u *UnleaseService) GetImages(imgConf ImageConfig) error {
	apiUrl := "https://api.unsplash.com"
	query := imgConf.Category
	imgCount := imgConf.TotalDownloadImage
	maxImage := strconv.Itoa(imgCount)
	accessKey := u.apikey
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var imagePath string = u.path

	url := fmt.Sprintf("%s/photos/random?client_id=%s&count=%s&orientation=landscape&query=%s", apiUrl, accessKey, maxImage, query)
	fmt.Println("Download...")
	fmt.Println(url)

	result, err := getImage(ctx, url)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup

	errCh := make(chan error, len(result))

	for key, v := range result {
		wg.Add(1)
		struri := v.User.Links.Html
		parsedURL, err := neturl.Parse(struri)

		if err != nil {
			return fmt.Errorf("Error parsing URL: %v", err)
		}

		path := parsedURL.Path
		parts := strings.Split(path, "/")
		username := parts[len(parts)-1]

		go (func(u string, k int, uname string) {
			defer wg.Done()
			if err := download(ctx, u, k, imagePath, uname); err != nil {
				errCh <- err
			}
		})(v.Urls.Full, key, username)
	}
	wg.Wait()
	close(errCh)

	for err := range errCh {
		return err
	}
	return nil
	// select {
	// case err := <-errCh:
	// 	return err
	// default:
	// 	return nil
	// }
}

func getImage(ctx context.Context, url string) ([]Image, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[GetImage]: request failed: %w", err)
	}
	defer resp.Body.Close()
	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var p []Image
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		return nil, err
	}
	return p, nil
}

func download(ctx context.Context, image string, index int, IMAGE_DIR string, upr string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, image, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("[Download]: request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	fn := fmt.Sprintf("%s/picasa_%v_%s.jpg", IMAGE_DIR, index, upr)

	info := fmt.Sprintf("Downloading: %s", fn)
	fmt.Println(info)

	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = io.Copy(f, resp.Body)

	if err != nil {
		return err
	}

	fmt.Println("Downloaded to: ", f.Name())

	return nil
}
