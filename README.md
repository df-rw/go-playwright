# go-playwright

Go web application and playwright

## Demo application

* Simple todo application here. Nothing special.
* `make dev` to make it go (requires https://github.com/air-verse/air for live
  reloading). Click click click.
* `make test-unit` to run unit tests using the go [testing
  library](https://pkg.go.dev/testing).
* `make test-browser` to run the stand-alone playwright test app.
* `make test-browser-with-testing` to run playwright from the go testing library.

## Install playwright

Install [Playwright for Go](https://github.com/playwright-community/playwright-go):

```shell
go get -u github.com/playwright-community/playwright-go
```

### Install browsers

The browsers can be installed automatically by running:

```go
err := playwright.Install()
```

Output of the installation may look something like this:

```
go run ./cmd/playwright/do-tests.go
2025/01/29 13:30:29 INFO Downloading driver path=/Users/dfogarty/Library/Caches/ms-playwright-go/1.49.1
2025/01/29 13:31:06 INFO Downloaded driver successfully
2025/01/29 13:31:06 INFO Downloading browsers...
Removing unused browser at /Users/dfogarty/Library/Caches/ms-playwright/chromium-1071
Removing unused browser at /Users/dfogarty/Library/Caches/ms-playwright/firefox-1419
Removing unused browser at /Users/dfogarty/Library/Caches/ms-playwright/webkit-1869
Downloading Chromium 131.0.6778.33 (playwright build v1148) from https://playwright.azureedge.net/builds/chromium/1148/chromium-mac-arm64.zip
121.6 MiB [====================] 100% 0.0s
Chromium 131.0.6778.33 (playwright build v1148) downloaded to /Users/dfogarty/Library/Caches/ms-playwright/chromium-1148
Downloading Chromium Headless Shell 131.0.6778.33 (playwright build v1148) from https://playwright.azureedge.net/builds/chromium/1148/chromium-headless-shell-mac-arm64.zip
77.5 MiB [====================] 100% 0.0s
Chromium Headless Shell 131.0.6778.33 (playwright build v1148) downloaded to /Users/dfogarty/Library/Caches/ms-playwright/chromium_headless_shell-1148
Downloading Firefox 132.0 (playwright build v1466) from https://playwright.azureedge.net/builds/firefox/1466/firefox-mac-arm64.zip
81.6 MiB [====================] 100% 0.0s
Firefox 132.0 (playwright build v1466) downloaded to /Users/dfogarty/Library/Caches/ms-playwright/firefox-1466
Downloading Webkit 18.2 (playwright build v2104) from https://playwright.azureedge.net/builds/webkit/2104/webkit-mac-15-arm64.zip
62.8 MiB [====================] 100% 0.0s
Webkit 18.2 (playwright build v2104) downloaded to /Users/dfogarty/Library/Caches/ms-playwright/webkit-2104
Downloading FFMPEG playwright build v1010 from https://playwright.azureedge.net/builds/ffmpeg/1010/ffmpeg-mac-arm64.zip
1.1 MiB [====================] 100% 0.0s
FFMPEG playwright build v1010 downloaded to /Users/dfogarty/Library/Caches/ms-playwright/ffmpeg-1010
2025/01/29 13:32:15 INFO Downloaded browsers successfully
```

## Running tests

Browser tests run by playwright can exist as stand-alone executables, or be
wrapped with the Go [testing library](https://pkg.go.dev/testing).

### Stand-alone

* Make a new package under `/cmd`.
* Create a `main.go` file that calls the [Playwright
  API](https://playwright.dev/docs/api/class-playwright).

An example stand-alone test is in `cmd/playwright/do-tests.go`.

### Wrapped in testing

TODO

## References

* https://playwright-community.github.io/playwright-go/
