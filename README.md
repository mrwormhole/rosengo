# Rosengo

## - How to install and run on desktops

### Quick start
```
git clone github.com/MrWormHole/rosengo
cd rosengo
go run github.com/MrWormHole/rosengo
```

## - How to build for Android

### Requirements
- Ensure that you have ebitenmobile CLI installed globally
```
go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@latest
```
- Ensure that you have java and javac installed and configured in your env variables.
- Ensure that you have the Android SDK installed and configured in your env variables.
- Ensure that you have the Android NDK installed and configured in your env variables.

### Build .aar file and sources.jar
```
git clone github.com/MrWormHole/rosengo
cd rosengo
go run github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile bind -target android -javapkg com.goldenhand.rosengo -o ./mobile/android/rosengo/rosengo.aar ./mobile
```

and run the Android Studio project in `./mobile/android`