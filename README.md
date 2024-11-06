# Wallhaven Downloader

This Go program downloads wallpapers from Wallhaven using their API. It allows you to configure search queries, resolution, categories, and limits through a `config.json` file.

## Features

- Search and download wallpapers based on configurable criteria.
- Limit the number of wallpapers downloaded.
- Save wallpapers to a local directory.

## Configuration

Create a `config.json` file in the project directory with the following structure:

```json
{
  "api_key": "YOUR_API_KEY",
  "search_query": "nature",
  "resolution": "1920x1080",
  "categories": "111",
  "sorting": "random",
  "order": "desc",
  "limit": 10,
  "download_limit": 5
}