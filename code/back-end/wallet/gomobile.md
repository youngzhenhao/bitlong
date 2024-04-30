# gomobile

## Install latest version of Golang

- https://go.dev/dl/

```bash
go version
```

## Install Android Studio

- Install the **latest version** of the `SDK` and **version** 21 of the `NDK` in Android Studio

### Ubuntu

```bash
apt install sdkmanager
sdkmanager --install "ndk;21.4.7075529"
sdkmanager --install "build-tools;35.0.0-rc3"
sdkmanager --install "platforms;android-34"
mkdir /root/Android
mkdir /root/Android/Sdk
mkdir /root/Android/Sdk/ndk
mkdir /root/Android/Sdk/platforms
cp -r /opt/android-sdk/build-tools/35.0.0-rc3/ /root/Android/Sdk/
cp -r /opt/android-sdk/ndk/21.4.7075529/ /root/Android/Sdk/ndk/
cp -r /opt/android-sdk/platforms/android-34/ /root/Android/Sdk/platforms/
```

## Install JDK 11

- https://www.oracle.com/java/technologies/javase-downloads.html

### Ubuntu

```bash
apt install openjdk-11-jdk
java --version
```

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

## Fix bug

```bash
 vim /root/bitlong/code/back-end/wallet/lib/lightning-terminal@v0.12.2/terminal.go
```

- delete line 94

## Use gomobile

- https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile#hdr-Build_a_library_for_Android_and_iOS

*To package a specific sub-module, you need to go to that directory first.*

```bash
cd api
gomobile bind -target android 
```
