<p align="center">
  <img src="assets/lightbanner.svg#gh-light-mode-only" alt="Light Banner" width="600" style="display:block;">
  <img src="assets/darkbanner.svg#gh-dark-mode-only" alt="Dark Banner" width="600" style="display:block;">
  <br>
  A lightweight golang based concert notifier
  <br>
</p>


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