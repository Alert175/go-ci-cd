# gb.ru - модуль Tech Lead: Автоматизация цикла разработки, урок 3

## Что было сделано:
- написан стартовый проект, который подключается к СУБД
- прописаны файлы для контейнеризации: Dockerfile, docker-compose.yaml, .dockerignore
- далее собраны образы, причем для проекта __app__, использую тэг в формате semver
- для имитации сервера, создал папку __server__, в ней прописал docker-compose.yaml и скопировал необходимый __.env__

Единственное с чем не смог до конца разобраться, как правильно настроить docker-compose.yaml с использованием __healthcheck__ для __app__ сервиса, тк при запуске ругается на нарушение синтаксиса.

`ERROR: The Compose file './docker-compose.yaml' is invalid because:
services.app.depends_on contains an invalid type, it should be an array`