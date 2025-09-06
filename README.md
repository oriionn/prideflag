# prideflag.fun
A simple test about pride flags !!

## Installation (Standalone)
### Prerequisites
- `Golang`

### Installation
1. Clone the repository
```sh
git clone https://github.com/oriionn/prideflag.git
cd prideflag
```

2. Build the program
```sh
make
```

3. Run the program
```sh
./prideflag.fun
```

## Installation (Docker)
### Prerequisites
- `Docker`

### Installation
1. Clone the repository
```sh
git clone https://github.com/oriionn/prideflag.git
cd prideflag
```

2. Edit the `docker-compose.yml` like you want
```sh
make
```

3. Run with docker compose
```sh
docker compose up -d --build
```

## CLI Params
To customize program behavior, the program includes several CLI parameters, listed below:
- `--port` / `-p`: To change your program's port
- `--database` / `-d` : To change the path of your SQLite database
