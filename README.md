# totext

A Go wrapper library to convert different types of documents to plain text.

## Dependencies

### To convert MS word doc files, install `wv`

For Ubuntu/Debian:

```bash
sudo apt install wv
```

For MacOs:

```bash
brew install wv
```

### To convert PDF files, install `poppler`

For Ubuntu/Debian:

```bash
sudo apt install poppler-utils
```

For MacOs:

```bash
brew install poppler
```

### To convert RTF files, install `unrtf`

For Ubuntu/Debian:

```bash
sudo apt install unrtf
```

For MacOs:

```bash
brew install unrtf
```

### To convert HTML files, `prettier` is required

```bash
npm init
npm install --save-dev --save-exact prettier
```

### To fetch remote web page and extract text

When a remote page is requested to be fetched by the application
for the first time, it will download the latest version of
[Chromium browser](https://commondatastorage.googleapis.com/chromium-browser-snapshots/index.html)
automatically.

`prettier` is required

```bash
npm init
npm install --save-dev --save-exact prettier
```

## Building command line tool

```bash
go mod tidy
chmod +x compile.sh
./compile.sh
```
