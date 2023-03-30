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
#### Create 2 tables 
```mysql
create table if not exists categories(
    id   bigint auto_increment primary key,
    name varchar(100) not null
);

create table if not exists products(
    id          bigint auto_increment primary key,
    name        varchar(100) not null,
    price       double       not null,
    category_id bigint       null,
    created_at  datetime(3)  null,
    updated_at  datetime(3)  null,
    deleted_at  datetime(3)  null,
    constraint fk_products_category foreign key (category_id) references goexpert.categories (id)
);

create index idx_products_deleted_at on goexpert.products (deleted_at);
```

#### Select
```mysql
select p.id         as id,
       p.name       as name,
       p.price      as price,
       c.name       as category,
       p.created_at as reated_at,
       p.updated_at as updated_at,
       p.deleted_at as deleted_at
from products p join categories c on p.category_id = c.id
```

