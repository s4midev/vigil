![Light Banner](assets/lightbanner.svg#gh-light-mode-only)
![Dark Banner](assets/darkbanner.svg#gh-dark-mode-only)

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

## Attributions
- Icon originally from [GameIcons](https://github.com/game-icons/icons)