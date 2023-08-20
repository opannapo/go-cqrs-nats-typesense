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