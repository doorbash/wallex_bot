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
docker run --name wallex_bot --restart always -d ghcr.io/doorbash/wallex_bot -t AAAAAAAAAAAAAA
```