package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Wallpaper struct {
	ID     string `json:"id"`
	Path   string `json:"path"`
	Thumbs struct {
		Original string `json:"original"`
	} `json:"thumbs"`
}

type Response struct {
	Data []Wallpaper `json:"data"`
}

func fetchWallpapers(config *Config) ([]Wallpaper, error) {
	url := fmt.Sprintf("https://wallhaven.cc/api/v1/search?q=%s&resolutions=%s&categories=%s&sorting=%s&order=%s&apikey=%s&purity=%s",
		config.SearchQuery, config.Resolution, config.Categories, config.Sorting, config.Order, config.APIKey, config.Purity)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

func downloadWallpaper(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	wallpapers, err := fetchWallpapers(config)
	if err != nil {
		fmt.Println("Error fetching wallpapers:", err)
		return
	}

	if err := os.MkdirAll("wallpapers", os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	downloadCount := 0
	for _, wallpaper := range wallpapers {
		if downloadCount >= config.DownloadLimit {
			break
		}
		filepath := filepath.Join("wallpapers", wallpaper.ID+".jpg")
		fmt.Println("Downloading:", wallpaper.Path)
		if err := downloadWallpaper(wallpaper.Path, filepath); err != nil {
			fmt.Println("Error downloading wallpaper:", err)
		} else {
			downloadCount++
		}
	}
}
