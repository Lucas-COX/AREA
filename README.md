# AREA
Action REAction application

## Setup database

> Make sure that the `area` user has all permissions over the `Area` database and that the database exists.
```bash
cat area.sql | mysql -u area -p Area
```

## Start project
To start the server use :
```bash
make run t=server
```

To start the web client use :
```bash
make run t=web
```

> To use the mobile client download the apk.

