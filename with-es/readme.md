1. run elasticsearch in docker.
    ```
    docker network create elastic
    docker run --name elasticsearch --net elastic -p 9200:9200 -e discovery.type=single-node -e ES_JAVA_OPTS="-Xms1g -Xmx1g" -e xpack.security.enabled=false -it docker.elastic.co/elasticsearch/elasticsearch:8.8.2
    ```
   
2. http://localhost:9200/_cat/indices


Ref: 
   1. https://medium.com/@vikaskumar4793/golang-with-elasticsearch-c6e274f4838b
   2. https://medium.com/@dinhhuy_67517/x%C3%A2y-d%E1%BB%B1ng-d%E1%BB%8Bch-v%E1%BB%A5-t%C3%ACm-ki%E1%BA%BFm-v%E1%BB%9Bi-go-v%C3%A0-elaticsearch-e3b24a12aea7
   3. https://dev.to/aleksk1ng/go-and-elasticsearch-full-text-search-microservice-in-k8s-2o7g
   4. https://viblo.asia/p/docker-la-gi-kien-thuc-co-ban-ve-docker-maGK7qeelj2