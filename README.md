# structura

## Список задач

### Инфраструктура

- [x] Написать docker окружение
    - [x] Написать dockerfile для сервиса
    - [x] Написать docker-compose файл
    - [x] Написать Makefile для сборки и запуска
- [ ] Написать CI/CD
    - [ ] Написать конфигурацию для github actions
- [x] Добавить линтер

### Хранилище

- [ ] Интегрироваться с postgresql
    - [ ] Написать скрипт для создания базы данных и схемы
    - [ ] Написать миграции для таблиц используя goose
    - [ ] Написать репозиторий для работы с базой данных
- [ ] Написать In Memory Cache в самом процессе
    - [ ] Написать интерфейс для работы с кэшем
    - [ ] Написать реализацию кэша на основе map + sync.RWMutex
    - [ ] Добавить репозиторий для работы с кэшем
        - [ ] Обновление данных в кеше должно происходить по cron раз в минуту
        - [ ] Прогрев данных должен происходить в момент старта приложения

### Наблюдаемость

- [ ] Добавить поддержку трейсинга opentelemetry
- [x] Добавить поддержку логирования
    - [x] Написать конфигурацию для логирования в json
    - [x] Написать middleware для обработки ошибок

### Основной функционал

- [ ] Написать тесты
    - [ ] Написать юнит тесты
    - [ ] Написать интеграционные тесты для REST с использованием testcontainers

#### Crawler

- [ ] Написать crawler для парсинга сайта
    - [ ] Добавить возможность подсчета кол-ва страниц в домене
    - [ ] Добавить возможность скачивания страниц в домене

#### Parser

- [ ] Написать парсер для поиска header на странице
    - [ ] Добавить возможность поиска в Tilda
    - [ ] Добавить возможность поиска в Wordpress
    - [ ] Добавить возможность поиска в Bitrix
    - [ ] Добавить возможность поиска в HTML5
- [ ] Написать парсер для поиска footer на странице
    - [ ] Добавить возможность поиска в Tilda
    - [ ] Добавить возможность поиска в Wordpress
    - [ ] Добавить возможность поиска в Bitrix
    - [ ] Добавить возможность поиска в HTML5

#### API

- [x] Написать REST интерфейс
