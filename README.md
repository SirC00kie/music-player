# music-player

Сервис на языке Go, моделирующий работу музыкального плейлиста.

### Технологии

`Go, HTTP, gRPC, Docker `

Функционал:
````
- AddSong - Добавление песни
- GetSong - Получение песни
- UpdateSong - Обновление песни
- DeleteSong - Удаление песни
- GetPlaylist - Получение плейлиста
- PlaySong - Начало воспроизведения
- PauseSong - Остановка воспроизведения
- NextSong - Воспроизведение следующей песни
- PrevSong - Воспроизведение предыдущей песни
````

## REST API

### AddSong
- POST `/api/v1/songs`

Request
````
    {
        "title":"title 1",
        "author":"author 1",
        "duration": 100000
    }
````
Response
````
    status 201
````

### GetSong
- GET `/api/v1/songs/:index`

Response
````
    {
        "title":"title 1",
        "author":"author 1",
        "duration": 100000
    }
````

### Update
- PUT `/api/v1/songs/:index`

Request
````
   {
        "title":"title 1",
        "author":"author 1",
        "duration": 100000
    }
````
Response
````
    status 204
````

### Delete
- DELETE `/api/v1/songs/:index`

Response
````
    status 204
````

### GetPlaylist
- GET `/api/v1/playlist`


Response
````
{
        "currentTime":0,
        "pausedTime":0,
        "playing":false,
        "songs":[
        {
            "author":"author 3",
            "duration":100000,
            "title":"title 3"
        },
        {
            "author":"author 2",
            "duration":100000,
            "title":"title 2"
        },
        {
            "author":"author 1",
            "duration":100000,
            "title":"title 1"
        }],
    }
````

### Play
- GET `/api/v1/play`

Response
````
    {
        "title":"title 1",
        "author":"author 1",
        "duration": 100000
    }
````

### Pause
- GET `/api/v1/pause`

Response
````
    {
        "title":"title 1",
        "author":"author 1",
        "duration": 100000
    }
````

### Next
- GET `/api/v1/next`

Response
````
    {
        "title":"title 1",
        "author":"author 1",
        "duration": 100000
    }
````

### Prev
- GET `/api/v1/prev`

Response
````
    {
        "title":"title 1",
        "author":"author 1",
        "duration": 100000
    }
````

## gRPC

Имеет аналогичный функционал. 

Клиент находиться в `/client/client.go`