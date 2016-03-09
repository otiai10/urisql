urisql
=======

```sh
% urisql -uri=mysql://root:xxx@localhost/mydb?reconnect=true
# mysql --user=root --host=localhost --database=mydb --password=xxx
```

```sh
% heroku config | grep DATABASE_URL | awk '{print $2}' | urisql
# mysql --user=abcde --host=region.cleardb.net --database=heroku_uuid --password=xxxxx
```
