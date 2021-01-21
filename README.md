# Quasar-fire

---

### Service config

***Satellites position***

Name      | X   | Y
----------|-----|-----
kenobi    | -500| -200
skywalker | 100 | -100
sato      | 500 | 100

***Environment values***

Name        |   Default value
------------|----------------
GRPC_PORT   |   50051
REST_PORT   |   8080

**in heroku use PORT env to set REST_PORT value*

###CLI

*__serve quasar__*: start grpc and rest services

### Make files from proto
```shell
sh ./hack/make-protos.sh
```

### Run application

***Golang***

```shell
  go run ./cmd serve quasar
```

***docker-compose***

```shell
  docker-compose build && docker-compose run serve
```

### Deployment

***Heroku***

```shell
  heroku login
  heroku create "your application name" 
  heroku stack:set container
  heroku container:login
  heroku container:push web
  heroku container:release web
  git push heroku master
```

***API***

- documentation: https://quasar-fire-alkapa.herokuapp.com/swagger

- endpoints:  https://quasar-fire-alkapa.herokuapp.com/
    ```
    POST → /topsecret/
    {
        "satellites": [
            {
                “name”: "kenobi",
                “distance”: 100.0,
                “message”: ["este", "", "", "mensaje", ""]
            },
            {
                “name”: "skywalker",
                “distance”: 115.5
                “message”: ["", "es", "", "", "secreto"]
            },
            {
                “name”: "sato",
                “distance”: 142.7
                “message”: ["este", "", "un", "", ""]
            }
        ]
    }
    
    RESPONSE CODE: 200
    {
        "position": {
            "x": -100.0,
            "y": 75.5
        },
        "message": "este es un mensaje secreto"
    }
    
    ERROR CODE: 404
    {
        "error":"message quantity not valid"
        "code":404
    }
    ```

    ```
    POST → /topsecret_split/{satellite_name}
        {
            "distance": 100.0,
            "message": ["este", "", "", "mensaje", ""]
        }
  
    RESPONSE CODE: 200
        {
            "position": {
                "x": -100.0,
                "y": 75.5
            },
            "message": "este es un mensaje secreto"
        }
    
    ERROR CODE: 404
    {
        "error":"insufficient data"
        "code":404
    }
    ```
  
    ```
    GET → /topsecret_split/
  
    RESPONSE CODE: 200
        {
            "position": {
                "x": -100.0,
                "y": 75.5
            },
            "message": "este es un mensaje secreto"
        }
    
    ERROR CODE: 404
    {
        "error":"insufficient data"
        "code":404
    }
    ```
  - error messages 404:
    - _satellite alliance name it's duplicate_
    - _insufficient data_
    - _messages len not equal_
    - _can't get localization_
    - _can't get message_