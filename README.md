# restaurant

## Creating the docker containers

```sh
cp .env.example .env
docker-compose up -d --build
```

## Docker Env

If in APP_ENV=prod the **docker-compose up -d --build** command will use the [prod.dockerfile](docker/prod.dockerfile) file, in case APP_ENV=dev will use the [dev.dockerfile](docker/dev.dockerfile) file

## Create a Todo

```sh
curl --request POST 'localhost:3001/v1/todo' \
    -H "Content-Type: application/x-www-form-urlencoded" \
    --data-urlencode "title=Fazer alguma coisa" \
    --data-urlencode "description=Fazer alguma coisa que eu goste"
# or
curl -X POST http://localhost:3001/v1/todo \
-H "Content-Type: application/json" \
-d '{"title":"Fazer alguma coisa", "description":"Fazer alguma coisa que eu goste"}'


```
