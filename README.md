# PlayFabMultiplayerServersCLI

A totally unofficial cross-platform CLI to manage [Azure PlayFab MultiPlayer Servers](https://playfab.com/features/game-services/multiplayer/) written in Go.

Uses

- [Go Modules](https://github.com/golang/go/wiki/Modules)
- [PlayFab's Multiplayer Server API](https://api.playfab.com/documentation/Multiplayer#MultiplayerServer)
- [A totally unofficial Golang PlayFab API that I wrote](https://github.com/dgkanatsios/playfabsdk-go) (there is an [open Pull Request](https://github.com/PlayFab/SDKGenerator/pull/392) for that on PlayFab's official SDK repo)

Heavy work in progress. No guarantees. Why Go? Because cross compilation to native executables is [really easy](https://golangcookbook.com/chapters/running/cross-compiling/).

## Supported commands

* th login --title <titleID> --secret <secret key> // or th login - <titleID> -s <secret key>
* th server enable
* th asset create -f c:\winrunnerSample6.zip
* th asset delete winrunnerSample6.zip
* th build list
* th build get -b <buildID>
* th build delete -b <buildID>

---
This is NOT an official Microsoft/PlayFab product.