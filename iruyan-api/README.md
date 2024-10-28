## Docker コンテナ内での確認

PostgreSQL コンテナに接続して直接データベースを操作する方法です。

### 手順

1. PostgreSQL コンテナにアクセス

```
docker-compose exec db psql -U user -d exampledb
```

上記コマンドで、データベース名が exampledb で、ユーザーが user の場合に接続します。

2. テーブルの一覧を表示

```
\dt
```

3. テーブルの内容を確認

```
SELECT * FROM users; -- ユーザーテーブルを例としています
```

4. 終了

```
\q
```
