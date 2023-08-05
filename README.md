# 6obcy people monitor

### Simple service to monitor how many people are on [6obcy](https://6obcy.org/) at selected time

---

### Stack

**backend**: go (server in fiber and scraper)

**frontend**: solidjs

---

### Configure environment

If you are running this project locally you should set variables in .env.example file, otherwise defaults would be loaded which are fine for testing most of the time

If you are running on docker all variables are set in docker-compose.yml

### Running

### Scraper

To run backend type following commands

```bash
cd server
go mod download
go run cmd/scraper/main.go
```

#### Backend

To run backend type following commands

```bash
cd server
go mod download
go run main.go
```

#### Frontend

```bash
cd client
pnpm install
pnpm run dev
```

### Running on docker

To start this project on docker just run

```bash
docker-compose up --build
```

### Requirements

> Note that you should have [go](https://go.dev/) and [node](https://nodejs.org/) with [corepack enabled](https://pnpm.io/installation#using-corepack)


