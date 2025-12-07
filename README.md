# gimme

A lightweight, extensible command-line downloader written in Go for grabbing direct media files from various websites.

## Installation

```sh
go install github.com/valsaven/gimme@latest
```

Or clone and build manually:

```sh
git clone https://github.com/valsaven/gimme.git
cd gimme
go build -o gimme
```

## Usage

```sh
gimme <url>
```

### Examples

```bash
gimme https://z0r.de/3483?flash
gimme https://z0r.de/493?flash
```

The file is saved to the current directory using its original filename.

## List of supported sites

- <http://z0r.de/>

## License

[GNU GPL v3](LICENSE)
