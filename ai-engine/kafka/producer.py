from confluent_kafka import Producer

def kafka_produce(topic, message, bootstrap_servers="kafka:9092"):
    producer = Producer({'bootstrap.servers': bootstrap_servers})
    producer.produce(topic, message.encode('utf-8'))
    producer.flush() 