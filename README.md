# Albus
A Telegram Bot that converts documents to more readable formats.

## Usage

```
git submodule update --init --recursive
cp .env.example .env
```

Put your Telegram Bot Token inside `.env`.

### Docker
```
make
make run
```

### Poetry
Ensure `pandoc` and `tectonic` are installed.

```
go run .
```
