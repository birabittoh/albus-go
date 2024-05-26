# Albus
A Telegram Bot that converts documents to more readable formats.

## Usage

```
git submodule init
git submodule update
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
