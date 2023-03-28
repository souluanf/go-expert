### Subir database
```shell
docker compose up -d
```

#### Acessar database
```shell
docker exec -it mysql mysql -u root -p
```

####  Select database 
```mysql
use goexpert;
```

#### Create table
```mysql
create table produtos (id varchar(255) not null,name varchar(80) not null,preco decimal(10,2) not null,primary key (id));
```