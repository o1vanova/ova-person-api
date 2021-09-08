# ova-person-api
This project implements API for persons

This project uses Postgres database.
You can find the configuration in the folder <b>configs</b>, where connection path is located.

Before the first deploy you should run these scripts:

<code>create database exampledb

create user test with encrypted password 'ozon'

grant all privileges on database exampledb to test
</code>

<b>exampledb</b> - database name<br>
<b>test</b> - role name<br>
<b>ozon</b> - password
