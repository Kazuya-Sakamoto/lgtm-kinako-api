## 概要

## 対応 issue

- #39

## API 動作確認

■ db（API 実行前）

```sql
mysql> select * from albums;
+----+----------+----------------------------------------------------------+-------------------------+-------------------------+---------+
| id | title    | image                                                    | created_at              | updated_at              | user_id |
+----+----------+----------------------------------------------------------+-------------------------+-------------------------+---------+
|  1 | ???????1 | https://d18g0hf2wnz3gs.cloudfront.net/20230926223208.JPG | NULL                    | NULL                    |       1 |
|  2 | 2        | https://d18g0hf2wnz3gs.cloudfront.net/20231118132545.JPG | NULL                    | NULL                    |       1 |
|  3 | ????     | https://d18g0hf2wnz3gs.cloudfront.net/20240501142223.JPG | 2024-05-01 14:22:23.793 | 2024-05-01 14:22:23.793 |       1 |
+----+----------+----------------------------------------------------------+-------------------------+-------------------------+---------+
3 rows in set (0.01 sec)

```

■ API

```
http://localhost:8081/api/v1/albums
```

■ Request Parameter

```json
なし
```

■ Response

```json
[
  {
    "id": 2,
    "title": "2",
    "image": "https://d18g0hf2wnz3gs.cloudfront.net/20231118132545.JPG",
    "tags": [],
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": 1,
    "title": "タイトルテスト1",
    "image": "https://d18g0hf2wnz3gs.cloudfront.net/20230926223208.JPG",
    "tags": [
      {
        "id": 1,
        "name": "おすすめ"
      },
      {
        "id": 2,
        "name": "へんてこ"
      },
      {
        "id": 3,
        "name": "ぽーとれーと"
      },
      {
        "id": 4,
        "name": "いべんと"
      }
    ],
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  },
  {
    "id": 3,
    "title": "3",
    "image": "https://d18g0hf2wnz3gs.cloudfront.net/20231028001149.JPG",
    "tags": [],
    "created_at": "0001-01-01T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  }
]
```

■ db（API 実行後）

```sql
GETのため不要
```

## テスト内容

- [ ] 画像の通り画面が正常にレンダリングすること

## テスト内容（開発者向け）

- [ ] test と build が通ること
