# 開発環境

### db, adminerを立ち上げる
```sh
$ docker-compose up -d
```

### mainプログラムを動かす
```sh
$ air
```

### migration
**up**
```sh
$ make migrate-up $C={count}
```

**down**
```sh
$ make migrate-down $C={count}
```

**fix dirty**
```sh
$ make migrate-force $V={version}
```


## api

### register
```sh
$ curl localhost:5050/register  \
-XPOST -H "content-type: application/json" \
-d '{"name": "yuhi", "email": "foo@gmail.com", "password": "password"}'
```