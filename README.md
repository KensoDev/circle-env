# .env swiss army knife for CircleCI (but not only)

[![Circle CI](https://circleci.com/gh/KensoDev/circleenv/tree/master.svg?style=svg)](https://circleci.com/gh/KensoDev/circleenv/tree/master)

## What?

This is a swiss army knife CLI to work with environment variables and CircleCI

### What it does

#### Upload environment variables from `.env` to circleCI

Locally, you can work with `.env` and when you execute `circle-env sync` it will "upload" all the variabvles to CircleCI.

#### Replace all ENV variables in template files and rename them

For example, lets say you have a template file that looks like this

`config/database.yml.template`


```
production:
	gmail_username: <GMAIL_USERNAME>
	gmail_password: <GMAIL_PASSWORD>
```

When you execute `circle-env replace` it will scan the file for everything between `<>` and replace the value with the ENV variable.

If no ENV variables with that name exists, it will leave it as is.


## Why?

Environment variables are a good way to store "secrets" without checking them
into source control.

Things like database passwords, API keys and others do not belong in your
source code.

`.env` and the ruby project [dotenv](https://github.com/bkeepers/dotenv) are a
good way to automatically load project specific environment variable.

However, when you run your project in CI environment, those environment
variables don't exist there by default.

This is where this project comes in.

Also, it's easy to create template file and only use the "real" files on CI. Say for a Docker image build

## How?

1. Download the release from the [releases](https://github.com/KensoDev/circleenv/releases) page.
2. execute (refer to the usage section)

## Usage

```
NAME:
   circle-env - Circle CI commands

USAGE:
   circle-env [global options] command [command options] [arguments...]

VERSION:
   0.0.3

COMMANDS:
   sync, p      Sync Environment Variables to CircleCI
   replace, p   Replace ENV variables in files
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version

```

## `circle-env sync`

Syncs your `.env` file to the CircleCI project

* `project-name` your CircleCI project name (eg: `circleenv`)
* `username` Your CircleCI (and github) username (eg: `KensoDev`)
* `token` Your CircleCI API token

Optional flag is the filename where you store your environment variables. By
default it's `.env` in the current working directory (Which is the easiest to
follow/remember)

## `circle-env replace`

Scans your current working directory for `.template` files, scans the env variable in them and writes the file, replacing the content.

This will turn `database.yml.template` to `database.yml` with the variables replaces from env variables.

## License

The MIT License (MIT)

Copyright (c) 2016 Avi Tzurel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
