from confluent_kafka import Consumer

def kafka_consume(topic, group_id="ai-engine", bootstrap_servers="kafka:9092"):
    consumer = Consumer({
        'bootstrap.servers': bootstrap_servers,
        'group.id': group_id,
        'auto.offset.reset': 'earliest'
    })
    consumer.subscribe([topic])
    while True:
        msg = consumer.poll(1.0)
        if msg is None:
            continue
        if msg.error():
            print("Consumer error: {}".format(msg.error()))
            continue
        yield msg.value().decode('utf-8') 