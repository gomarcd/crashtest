[![Release](https://github.com/gomarcd/crashtest/workflows/Release/badge.svg)](https://github.com/gomarcd/crashtest/actions/workflows/release.yml)

# Crashtest

A fast and lightweight cross-platform API client made in 🇨🇦 Canada 🇨🇦

![Crashtest screenshot](screenshot.png)

## Features

- Modern, minimal UI
- Free and open source
- No paywall and no registration required
- Designed for privacy and security
- Set request headers/parameters/body and view response headers
- Windows, macOS and Linux support

## Technology

Built in Golang with [Wails](https://wails.io).

## Security

:white_check_mark: No telemetry, ads or trackers

:white_check_mark: Runs locally on your machine, the only data ever going out are queries explicitly sent by user

:white_check_mark: Code signing: macOS binaries are signed with official Apple certificate issued by Developer ID Certification Authority, Windows binaries will be code signed via Azure Trusted Signing pending its identity validation

:white_check_mark: GPG signing: macOS, Windows and Linux binares are always GPG-signed with ed25519 key `A65E9AE2` (Fingerprint: `1353 E058 CB77 A738 F6AE  3362 883E 797A A65E 9AE2`), so you can verify the downloaded files are indeed from me. You can download the pubkey from:

- Here in this repo by clicking `gpg-pubkey.asc` above or click [here](https://github.com/gomarcd/crashtest/blob/main/gpg-pubkey.asc)
- From Ubuntu keyserver with `gpg --keyserver hkps://keyserver.ubuntu.com --recv-keys 1353E058CB77A738F6AE3362883E797AA65E9AE2` or by [clicking here](https://keyserver.ubuntu.com/pks/lookup?search=ci%40crashtest.app&fingerprint=on&op=index)
- From openpgp.org with `gpg --keyserver hkps://keys.openpgp.org --recv-keys 1353E058CB77A738F6AE3362883E797AA65E9AE2` or by [clicking here](https://keys.openpgp.org/search?q=1353E058CB77A738F6AE3362883E797AA65E9AE2)

:white_check_mark: SHA256 checksums accompany downloads for every release so you can verify the integrity of the file

:white_check_mark: Reproducible builds with instructions are coming, stay tuned 