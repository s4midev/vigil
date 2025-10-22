# vigil
A lightweight golang based concert notifier

## Docker Compose Quick Start
```yaml
services:
  vigil:
    image: ghcr.io/s4midev/vigil:latest
    environment:
      - VIGIL_DATAPATH=/data
    volumes:
      - ./:/data
    user: "1000:1000"
```

## To-Do
- Custom CURL based notifications