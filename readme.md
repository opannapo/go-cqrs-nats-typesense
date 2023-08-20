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
    - Etc...
- Service Query :
    - Fiber (Http Framework)
    - TypeSense (Search Engine)
    - N.A.T.S (Message Broker)
    - Etc...

### (TODO) UPLOAD GAMBAR DIAGRAM




# Command Services

### build docker images : mysql,nats
**go to application directory :** ```cd command/deployment```
```dockerfile
sudo docker-compose -f 'docker-compose.mysql.yaml' up
```


# Query Services
build docker images : typesense
**go to application directory :** cd ```query/deployment```
```dockerfile
mkdir typesense-data
sudo docker-compose -f 'docker-compose.typesense.yaml' up
```