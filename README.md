# Iruyan Task コマンドによる環境構築手順

このプロジェクトでは、Taskfile によって定義されたコマンドを使って、開発環境を効率的に構築・操作できます。

---

## 🛠 Task コマンドを使った環境構築手順

事前に [`task`](https://taskfile.dev) コマンドがインストールされていることを確認してください。

---

### ✅ 1. リポジトリをクローン

```bash
git clone https://github.com/OverAgers/iruyan.git
cd iruyan
```

---

### ✅ 2. Docker コンテナを起動（バックグラウンド）

```bash
task run
```

- コンテナがバックグラウンドで起動されます（`docker-compose up -d` 相当）

---

### ✅ 3. アプリの確認

- フロントエンド（Next.js）: http://localhost:3000
- バックエンド（Go API）: http://localhost:8080

---

### ✅ 4. コンテナの停止と削除

```bash
task down
```

- コンテナの停止と削除を行います（`docker-compose down` 相当）

---

### ✅ 5. ボリューム（データ）ごと完全削除

```bash
task clean
```

- コンテナだけでなく、ボリューム（DB など）も完全に削除します（`docker-compose down -v` 相当）

---

### ✅ 6. タスク一覧の表示

```bash
task
```

- 利用可能な task コマンド一覧を表示します

---

以上のコマンドで、簡単に開発環境の立ち上げ・停止・初期化が可能です。
