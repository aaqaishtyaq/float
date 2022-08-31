# Float

Minimalist Configurable Homelab Start Page.

![float-banner](https://user-images.githubusercontent.com/22131756/187611815-29ec322a-40e5-4cbe-9596-49d14ad0bfef.png)

## Usage

``` console
Usage of ./float:
  -file string
        Config file with the float page content (default "/etc/float/float.yml")
  -port string
        Port to be used by HTTP Server (default "8080")
```

## Defaults

| **Flag** |        **Value**       |
|:---------:|:----------------------:|
|   `file`  | `/etc/float/float.yml` |
|   `port`  |         `8080`         |

## How to run using docker

### Docker Compose

```yaml
version: "3.9"
services:
  float:
    container_name: float
    image: ghcr.io/aaqaishtyaq/float:latest
    restart: unless-stopped
    ports:
      - "80:8080"
    volumes:
      - /mnt/data/float:/etc/float/
```

### Docker CLI

```console
docker run \
  -v /mnt/data/float:/etc/float/ \
  -p 80:8080 \
  --restart unless-stopped \
  --name float \
  ghcr.io/aaqaishtyaq/float:latest
```

## Example Configuration

```yaml
title: "Vessel"
page_data:
  - name: "Grafana"
    port: 3000
  - name: "Portainer"
    port: 9000
```
