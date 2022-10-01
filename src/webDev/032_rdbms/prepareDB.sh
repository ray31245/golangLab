# it will fail,may MYSQL_ROOT_PASSWORD is too simple
docker container run -d -p 3306:3306 --name db -e MYSQL_ROOT_PASSWORD=BevisAwesome mysql

# it work, run `docker logs db` to check password
docker container run -d -p 3306:3306 --name db -e MYSQL_RANDOM_ROOT_PASSWORD=yes -e MYSQL_DATABASE=test mysql
O/14jiQ6LOG4x5r5niA6RHkwKssek/yJ