docker system prune -a && docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q)

# запустить контейнеры в фоне
(cd gateway && docker-compose up -d rabbitmq)
(cd auth && docker-compose up -d)
(cd polls && docker-compose up -d)
(cd gateway && docker-compose up -d gateway)

echo
echo
echo

# вывести информацию по контейнерам
(cd gateway && docker-compose ps)
(cd auth && docker-compose ps)
(cd polls && docker-compose ps)

# docker system prune -a && docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q)