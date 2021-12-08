## Build
```
go build
```

## Usage:
```
wallex_bot [OPTIONS]
```

**Options:**
```
  -t string
        telegram bot token
```

## Example
```
./wallex_bot -t AAAAAAAAAAAAAA
```
**Docker:**
```
docker run --name wallex_bot --restart always -d -e TOKEN=AAAAAAAAAAAAAA ghcr.io/doorbash/wallex_bot
```
**Docker-Compose:**
```
wallex_bot:
    restart: always
    environment:
      - TOKEN=AAAAAAAAAAAAAA
    image: ghcr.io/doorbash/wallex_bot
```