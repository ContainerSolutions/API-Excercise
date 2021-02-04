alias psql='docker exec -it "$(docker ps -f 'name=postgres' -q)" psql -U postgres'
psql -c 'CREATE DATABASE titanic_dev;'
psql -c 'CREATE DATABASE titanic_prod;'
psql -c 'CREATE DATABASE titanic_test;'
