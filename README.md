# Rosengo

## How to install and run on desktops

```
mkdir rosengo
cd rosengo
git close github.com/MrWormHole/rosengo
go run github.com/MrWormHole/rosengo
```

## How to build for Android

```
go run github.com/MrWormHole/ebiten/v2/cmd/ebitenmobile bind -target android -javapkg com.goldenhand.rosengo -o ./mobile/android/rosengo/rosengo.aar ./mobile
```

and run the Android Studio project in `./mobile/android`