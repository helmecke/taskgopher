# Taskgopher

[![license](https://img.shields.io/github/license/helmecke/taskgopher.svg)](LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> Personal task manager inspired by taskwarrior.

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Background

I love taskwarrior but there some quirks which bug me.

## Install

As it's early in development install it via `go install`:
```bash
go install github.com/helmecke/taskgopher
```

## Usage

Create, list and complete a task:
```bash
alias t=taskgopher

t add very important task due:2022-07-01

t list

t complete 1
```

## Contributing

See [the contributing file](CONTRIBUTING.md)!

PRs accepted.

Small note: If editing the Readme, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

[MIT © Jakob Helmecke.](LICENSE)
