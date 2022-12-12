study-simple-bank
---
Scripture in https://github.com/techschool/simplebank .

## Requirement
This repository depends following tools.
- docker
- [migrate](https://github.com/golang-migrate/migrate)

## Usage
```bash
$ git clone https://github.com/n-hachi/study-simple-bank
$ cd study-simple-bank
$ make postgres
$ make createdb
$ make migrateup
$ make test
```
