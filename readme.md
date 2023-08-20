# Background
Project ini didasarkan pada penerapan/implementasi design CQRS (Command Query Responsibility Segregation).
Dalam beberapa sumber dijelaskan bahwa cqrs adalah penerapan even sourcing yang membagi tanggung jawab antara (minimal 2 service) proses write & read.
Dalam hal ini proses write biasa disebut dengan command, sementara proses read disebut dengan query. 
Anda bisa mencari tau lebih banyak dan detail terkait pro cons design ini diinternet. Pada project ini tidak dijelaskan lebih detail tentang apa itu CQRS. 
 
# Design
Pada project ini sinkronisasi data antar service di distribusikan menggunakan message broker menggunakan N.A.T.S.
Berikut ini tech stack yang digunakan dalam project:
- Microservices dikemas kedalam konsep MonoRepo
- Event Driven dengan message broker
- Golang (Programming Language)
- Service Command :
    - Fiber (Http Framework)
    - Mysql (Database)
    - N.A.T.S (Message Broker)
    - GORM (Sql ORM)
    - Database Migration Script
    - Etc...
- Service Query :
    - Fiber (Http Framework)
    - TypeSense (Search Engine)
    - N.A.T.S (Message Broker)
    - Etc...


![Screenshot from 2023-08-21 00-50-15](https://github.com/opannapo/go-cqrs-nats-typesense/assets/18698574/e7b0653b-0d09-482c-a84b-fb34f056f2e6)


# Demo
Demo applikasi Command dan Query.
Skenario :
- Command Service : Create Article (Via Postman)
- Query Service : Consume message from (subject) article created, lalu lakukan insert data ke TypeSense.
- Monitoring Dashboard TypeSense (https://bfritscher.github.io/typesense-dashboard/)
- Get Article By ID (Via Postman)
- Search Article (Filter author & query:title & body) (Via Postman)

### Output:
[Screencast from 21-08-23 01:00:41.webm](https://github.com/opannapo/go-cqrs-nats-typesense/assets/18698574/00e091f9-2429-48d1-92bb-5923b8be5b06)



# Installasi applikasi pihak ke tiga & Konfigurasi.
Work directory : **_deployment_**
1. Masuk ke lokasi docker file
```shell
cd deployment
```

2. build docker images : Mysql
```shell
sudo docker-compose -f 'docker-compose.mysql.yaml' up
```

3. build docker images : NATS
```shell
sudo docker-compose -f 'docker-compose.nats.yaml' up
```

4. build docker images : TypeSense
```shell
mkdir typesense-data
sudo docker-compose -f 'docker-compose.typesense.yaml' up
```

5. Konfigurasi TypeSense
Setelah proses installasi typesense pada docker berhasil, lakukan langkah pembuatan collection.
Anda dapat melakukan ini dengan mudah melalui link https://bfritscher.github.io/typesense-dashboard/#/
Login dengan API KEY yang tertera di configurasi **.env** file.

Buat collection dengan schema :
```json
{
  "name": "articles",
  "fields": [
    {
      "name": "author",
      "type": "string",
      "facet": false,
      "optional": false,
      "index": true,
      "sort": false,
      "infix": false,
      "locale": ""
    },
    {
      "name": "title",
      "type": "string",
      "facet": false,
      "optional": false,
      "index": true,
      "sort": false,
      "infix": false,
      "locale": ""
    },
    {
      "name": "body",
      "type": "string",
      "facet": false,
      "optional": false,
      "index": true,
      "sort": false,
      "infix": false,
      "locale": ""
    },
    {
      "name": "created",
      "type": "int64",
      "facet": false,
      "optional": false,
      "index": true,
      "sort": true,
      "infix": false,
      "locale": ""
    }
  ],
  "default_sorting_field": "",
  "enable_nested_fields": false,
  "symbols_to_index": [],
  "token_separators": []
}
```
#

## Menjalankan Applikasi Command
Work directory : command
1. Masuk ke folder command
```shell
cd command
```

2. Execute script migration untuk pembuatan table _**article**_
```shell
go run main.go migrateup
```

3. Jalankan Applikasi Command Service
```shell
go run main.go command
```
4. Output : Akan tampil display seperti ini pada terminal.
```shell
 ┌───────────────────────────────────────────────────┐ 
 │               GCNT-COMMAND-SERVICE                │ 
 │                   Fiber v2.48.0                   │ 
 │               http://127.0.0.1:1111               │ 
 │       (bound on host 0.0.0.0 and port 1111)       │ 
 │                                                   │ 
 │ Handlers ............. 5  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 28610 │ 
 └───────────────────────────────────────────────────┘ 
```

#
## Menjalankan Applikasi Query
Work directory : query
1. Masuk ke folder query
```shell
cd query
```

2. Jalankan Applikasi Command Service
```shell
go run main.go query
```
4. Output : Akan tampil display seperti ini pada terminal.
```shell
 ┌───────────────────────────────────────────────────┐ 
 │                GCNT-QUERY-SERVICE                 │ 
 │                   Fiber v2.48.0                   │ 
 │               http://127.0.0.1:2222               │ 
 │       (bound on host 0.0.0.0 and port 2222)       │ 
 │                                                   │ 
 │ Handlers ............. 8  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 29127 │ 
 └───────────────────────────────────────────────────┘ 
```


