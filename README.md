<p  align="center">
<h1  align="center">PaperID Golang Backend</h1>
</p>
  

## Feature

- Clean code (DDD)
- Gin Framework
- Docker and docker-compose
- ORM with gorm
- JWT with jwt-go
  
 ## API
  
 - [x] Signin
 - [x] Signup
 - [x] Signup Admin
 - [x] Signout
 - [x] Update Profile
 - [x] Get Profile
 - [x] Create Finance Account
 - [x] Update Finance Account
 - [x] Delete Finance Account
 - [x] Get Finance Account
 - [x] Create Finance Account Type
 - [x] Update Finance Account Type
 - [x] Delete Finance Account Type
 - [x] Get Finance Account Type
 - [x] Create Finance Transaction
 - [x] Update Finance Transaction
 - [x] Delete Finance Transaction
 - [x] Get Finance Transaction

## How to run

  

There are two ways to run this application, with docker or without docker

  

```bash
# running with docker

# copy .env
cp .env.example .env

# you can use live-reload when safe file
# running in with build, 
make run-local

# running in without build
make run-local-up

# remove production container
make down-local
```

```bash
# running in local/without docker
# copy .env
cp .env.example .env
make install
make run

```

## Run migration
```bash
make migrate
```

  

## Stay in touch
* Author - [Syukri Husaibatul Khairi](https://www.linkedin.com/in/syukri-husaibatul-khairi-a6297314b/)

  

## License
PaperIdGolangBackend is [MIT](LICENSE).