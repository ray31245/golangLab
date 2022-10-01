# TODO: try to use auth, curennt code is not avilbale
# docker container run -d --network mongo_work --name my_mongo	-e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=bW9uZ29hZG1pbgo= mongo

# throw test the code is work on mongo:5.0.6
docker container run -d --network mongo_work --name my_mongo mongo:5.0.6

# create container mongosh to operation cli on mongodb
docker run -it --network mongo_work --rm mongo mongosh --host my_mongo test

# cli refrence
https://blog.gtwang.org/programming/getting-started-with-mongodb-shell-1/