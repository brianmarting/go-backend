# Go-backend

## What does the application do?

The application manages all kinds of cryptocurrency related data. \
By default when running, it will:
- create tables
- try to connect to psql
- try to connect to rabbitmq and create queues in order to listen for withdrawal requests
- sets up a tcp server on the given port, listening for withdrawal requests from external parties

## Running

When running the application, you will need to provide env variables. The following can be used:

```
DB_HOST=localhost;DB_PASSWORD=postgres;DB_USERNAME=postgres;RABBITMQ_HOST=localhost;RABBITMQ_PASSWORD=guest;RABBITMQ_USERNAME=guest;TCP_WITHDRAWAL_LISTENER_PORT=8000;ENVIRONMENT=dev
```
You also need to run the `docker-compose.yml` file in order to start the psql and the rabbitmq instance so that the application can start.

## Tracing

Tracing is being done by using the `otlphttptrace` package, meaning that the application will send traces via http to the set location. \
This being done in the otlp format, and is being sent to an otel collector. In the collector it is being logged and being forwarded in the
same otlp format to the given ingester, in this case being jaeger. \
This can be viewed when running the `docker-compose.yml` and going to `http://localhost:16686/` (jaeger UI).
