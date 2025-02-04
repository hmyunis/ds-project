## How to Run the App From A Freshly Cloned Repository

**Step 1: Clone the repository or pull the latest version**

```bash
git clone https://github.com/hmyunis/ds-project.git
```

**Step 2: Open your terminal inside the repo**

```bash
cd ds-project
```

**Step 3: Setup your backend**

**3.1 Setup your database**

* Open MySQL Workbench and create a database named "chatapp".  You can run the following query:

```sql
CREATE DATABASE chatapp;
```

* On this newly created database, create a "users" table by running the SQL script found in `/server/db/migrations` for creating the table.

* Go to `/server/db/db.go` and edit line 16 to replace the placeholder username and password with your MySQL credentials.

```go
// Example:
sql.Open("mysql", "your_username:your_password@tcp(127.0.0.1:3306)/chatapp") 
```


**3.2 Setup your Go server**

```bash
cd server
go mod tidy
cd cmd
go run main.go
```

Now your backend server is live at `http://localhost:8080`.


**Step 4: Setup the frontend**

```bash
cd client
npm install
npm run dev
```

Open `http://localhost:3000` in your browser. Signup, then log in, then chat etc.


## Group Members

| Name             | ID No.      |
|-----------------|-------------|
| ESTIFANOS TAYE  | UGR/7285/14 |
| HAMDI MOHAMMED  | UGR/8929/14 |
| MOTI LEGGESE    | UGR/5389/14 |
| YORDANOS ZEGEYE | UGR/6316/14 |
| ZEAMANUEL ADMASU| UGR/8908/14 |
