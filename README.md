# ghremind
Simple Golang server which serves open pull requests as JSON.

Designed for usage as a dashboard screen, to remind teams of open pull requests.

## Usage

To get the open pullrequests from a repo, call `GET /pullrequests` endpoint with parameters `owner` and `repos`.

Example: `GET http://localhost:8080?owner=nenadstojanovikj&repos=ghremind,zipzap`

## Docker

To set up the project, you can use the provided `docker-compose.yml` file. There is a watcher running, so any changes you make
to the `.go` files will trigger recompilation.
