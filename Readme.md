How to Run the App From A Freshly Cloned Repository

Step 1: Clone the repository or pull the latest version

Step 2: Open your terminal inside the repo
 * Good job, you have almost completed half the steps 😆

Step 3: Setup your backend
3.1 Setup your database
* Open MySQL workbench and create a database named "chatapp".
    You can run the query "create database chatapp"
* On this created database, create a "users" table by running the SQL found in /server/db/migrations for creating the table.
* Go to /server/db/db.go and edit line 16 to your MySQL username and password. Replace "root" with your username, and "password" with your password.

3.2 Setup your Go server 
cd server
go mod tidy
cd cmd
go run main.go

Now your backend server is live at :8080

Step 5: Setup the frontend
cd client
npm install
npm run dev

Open localhost:3000 on your browser and play around. Signup, then log in, then chat.