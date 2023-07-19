1. Kafka Setup
    ```
   docker compose up -d
    ```
   
2. Configuration
    ``` 
   copy bootstrap.servers=localhost:9092
   ```

3. Create Topic
   ```
   docker compose exec broker \
      kafka-topics --create \
      --topic purchases \
      --bootstrap-server localhost:9092 \
      --replication-factor 1 \
      --partitions 1
   ```