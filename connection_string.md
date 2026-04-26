Finding your PostgreSQL connection string depends on whether you are connecting to a local database or a cloud-hosted one. A connection string (or URI) generally follows this standard format:

`postgresql://[user]:[password]@[host]:[port]/[dbname]`

1. Finding Details for a Local Database
If you installed PostgreSQL on your machine (or via WSL), the defaults are almost always the same.

***Host***: localhost (or 127.0.0.1)

***Port***: 5432 (this is the industry standard for Postgres)

***User***: postgres (the default superuser created during installation)

***Password***: Whatever you set during the installation process.

***Database***: postgres (the default initial database)

***Typical Local String***:
postgresql://postgres:your_password@localhost:5432/postgres

***Using the Terminal to Check***
If you are already logged into the psql terminal and want to see exactly how you are connected, run this command:

`\conninfo`

It will return a message like: "You are connected to database "postgres" as user "postgres" via socket in "/var/run/postgresql" at port "5432"..."