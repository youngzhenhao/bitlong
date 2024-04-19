# gomobile

## Install latest version of Golang

- https://go.dev/dl/

## Install Android Studio

- Install the **latest version** of the `SDK` and **version** 21 of the `NDK` in Android Studio

## Install JDK 11

- https://www.oracle.com/java/technologies/javase-downloads.html

## Install gomobile and gobind

```bash
go install golang.org/x/mobile/cmd/gomobile@latest
gomobile init
go install golang.org/x/mobile/cmd/gobind@latest
```

- get bind

```bash
go get golang.org/x/mobile/bind
```

## Use gomobile

```bash
cd api
gomobile bind -target android 
```
