# Record Appender

## Introduction

This project wants to bulk insert records into the database and then provide an interface for fetchting them.
The idea is comming from an interview question which wants you to use database, knwoing about bulk insert and
also writing HTTP endpoints.

Actually written by [Elahe Dastan](https://github.com/elahe-dastan/) at Spring 2020.

## How to run?

First you need to build the project:

```bash
go build
```

Then you need to do the migration:

```bash
./record-appender migrate
```

After that you can insert the data as bulk:

```bash
./record-appender setup
```

At the end you can serve the requests:

```bash
./record-appender serve
```
