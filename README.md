# movearr

Simple CLI tool to perform Sonnar/Radarr moves based on specific criteria.

## Sample Configuration

```yml
radarr:
  database: /opt/radarr/app/radarr.db
  url: https://radarr.domain.com
  api_key: token
  root_folder: /mnt/unionfs/Media/Movies
sonarr:
  database: /opt/sonarr/app/sonarr.db
  url: https://sonarr.domain.com
  api_key: token
  root_folder: /mnt/unionfs/Media/TV
```

## Assumptions

### Movies

Movie folders contain years within ( ), e.g. `(2020)`

Movie folders contain an IMDb **AND** TMDb ID (must have `imdb:` **AND** `tmdb:`).

Example: `/mnt/unionfs/Media/Movies/The Midnight Sky (2020) - [imdb:tt10539608] [tmdb:614911]`

### Series

Series folders contain years within ( ), e.g. `(2020)`

Series folders contain a TVDb ID (must have `tvdb:`).

Example: `/mnt/unionfs/Media/TV/24 (2001) - [tvdb:76290]`

## Commands

- `fix-ids` - Move series or movies where the folder contains old metadata id(s).

- `fix-years` - Move series or movies where the folder does not contain the correct year.

- `missing-ids` - Move series or movies where the folder does not have a metadata id(s).

All commands above support the `--dry-run` flag where ultimately no move request will be sent.

If no `--limit 5` is specified, ALL will be sent which is likely not want you want (if there is a large amount to move).

## Sample Commands

`movearr fix-years --pvr sonarr --limit 5`

`movearr fix-ids --pvr radarr --limit 5`

`movearr missing-ids --pvr radarr --limit 5`
