## Build
```
go build
```

## Usage:
```
TOKEN=AAAAAAAAAAAAAA USERNAME=wallex_api_bot ./wallex_bot [OPTIONS]
```

**Options:**
```
  -i int
        fetch interval in seconds (default 5)
```
## Example
**Docker:**
```
docker run --name wallex_bot --restart always -d -e TOKEN=AAAAAAAAAAAAAA -e USERNAME=wallex_api_bot ghcr.io/doorbash/wallex_bot
```
**Docker-Compose:**
```
  wallex_bot:
    restart: always
    environment:
      - TOKEN=AAAAAAAAAAAAAA
      - USERNAME=wallex_api_bot
    image: ghcr.io/doorbash/wallex_bot
```