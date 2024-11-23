
### psql 명령어
```
\l: 현재 서버에 있는 모든 데이터베이스 목록을 조회한다.
\c database_name: 지정한 데이터베이스로 접속한다.
\d: 현재 데이터베이스에 있는 모든 테이블, 뷰, 시퀀스, 인덱스 등의 목록을 조회한다.
\dt: 현재 데이터베이스에 있는 모든 테이블의 목록을 조회한다.
\dv: 현재 데이터베이스에 있는 모든 뷰의 목록을 조회한다.
\di: 현재 데이터베이스에 있는 모든 인덱스의 목록을 조회한다.
\q: psql 클라이언트를 종료하고 터미널로 빠져나간다.
```

![img.png](img/psql.png)


```sql
INSERT INTO users (email, username, password) VALUES ('mail@email.com', 'pass', 'test');

INSERT INTO comments(post_id, user_id, content) VALUES (1, 1, 'test123');
```