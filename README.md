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
$ make migrate-up
```

**down**
```sh
$ make migrate-down
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
-d '{"name": "yuhi", "email": "s7i.yuhi@gmail.com", "password": "password"}'
```