## _"Clinic" test application_

Prerequisite:
- [docker](https://gdevillele.github.io/)
- [go](https://go.dev/doc/install)
- [Makefile](https://dev.to/skypy/linux-make-install-command-2dd6)
- internet :)

## Public repo: [Clinic](https://github.com/404th/clinic) on GitHub.

## Running application after cloning it.

```
git clone git@github.com:404th/clinic.git
```

After cloning, global environment file ```.env``` must be created and example environments from ```.env.example``` are copied and pasted to ```.env```:
```
cd clinic
```
```
touch .env
```
Install packages:
```
go mod tidy
```
Install swaggo:
```
go get -u github.com/swaggo/swag 
```
```
make swag-init
```
If this is not running or returning error, make sure your GO Path is on the PATH environment variable: 
```
export PATH=$(go env GOPATH)/bin:$PATH 
```

Then, database should be created in docker container. To build db in container built-in commands are written in Makefile. We should run the command step-by-step:
```
make psqlcontainer
```
```
make createdb
```
Now, we successfully run container and created database =)

Next, we have to run migration to create tables in postgresql.

If you have not installed golang-migration yet, install [golang-migrate](https://github.com/golang-migrate/migrate) using command below:
```
brew install golang-migrate
```

Running migration up:
```
make migration-up
```

Yahooooo 😎
our application are ready to start: 
```
make start
```

After running app, API document available on local swagger:
For example: http://127.0.0.1:5050/swagger/index.html#/

###### Swaggo:
```http://{host}:{port}/swagger/index.html#/```
