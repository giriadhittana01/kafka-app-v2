Cheat Sheet

=> Connect to zookeeper
zookeeper-shell localhost:2181

=> List all brokers
ls /brokers/ids

=> List all topics
ls /brokers/topics
kafka-topics --list --bootstrap-server localhost:8097,localhost:8098,localhost:8099

=> Get broker status by ID
get /brokers/ids/[brokerid]

=> Describe topic
kafka-topics --describe --bootstrap-server localhost:8097,localhost:8098,localhost:8099 --topic [name]

=> Delete Specific Topic
kafka-topics --delete --bootstrap-server localhost:8097,localhost:8098,localhost:8098 --topic [name]

=> Create Topic
kafka-topics --bootstrap-server localhost:8097 --create --topic [name] --partitions [n] --replication-factor [n]

=> Create Topic with Retention Time
kafka-topics --bootstrap-server localhost:8097 --create --topic foo --partitions 4 --replication-factor 2 --config retention.ms=[ms]

=> Listen specific topic partition
kcat -C -b localhost:8097,localhost:8098,localhost:8099 -t foo -p [n]

=> Produce message
kafka-console-producer --bootstrap-server localhost:8098,localhost:8099,localhost:8097 --topic [name]

=> Consume message 
kafka-console-consumer --bootstrap-server localhost:8097,localhost:8098,localhost:8099 --topic [name] --from-beginning