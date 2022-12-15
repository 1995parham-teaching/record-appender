<h1 align="center">Record Appender</h1>

<h6 align="center">Append in Bulk, then Serve</h6>

<p align="center">
  <img src="https://img.shields.io/github/actions/workflow/status/1995parham-teaching/record-appender/ci.yaml?label=ci&logo=github&style=for-the-badge&branch=main" alt="GitHub Workflow Status">
  <img alt="GitHub" src="https://img.shields.io/github/license/1995parham-teaching/record-appender?logo=gnu&style=for-the-badge">
</p>

## Introduction

This project wants to bulk insert records into the database and then provide an interface for fetching them.
The idea is coming from an interview question which wants you to use database, knowing about bulk insert and
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
