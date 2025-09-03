FROM rabbitmq

RUN rabbitmq-plugins enable --offline rabbitmq_management

# management console: localhost:15672
EXPOSE 15671 15672