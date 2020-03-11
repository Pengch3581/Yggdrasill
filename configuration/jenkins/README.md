# jenkins 搭建

### jar 包启动

> 建议限制 java 虚拟机内存
nohup /usr/bin/java -server -Xms2000m -Xmx4000m -Xmn500m -Xss512k -XX:MetaspaceSize=64m -XX:MaxMetaspaceSize=128m -Djava.awt.headless=true -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:ConcGCThreads=1 -jar ./jenkins.war --httpListenAddress=127.0.0.1 &

### slave 节点配置

Manager 页面新增 node，打开安全中的 tcp 端口，配置 ssh-key

启动命令
nohup /usr/bin/java -server -Xms2000m -Xmx4000m -Xmn500m -Xss512k -XX:MetaspaceSize=64m -XX:MaxMetaspaceSize=128m -Djava.awt.headless=true -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:ConcGCThreads=1 -jar agent.jar -jnlpUrl http://127.0.0.1:8080/computer/slave1/slave-agent.jnlp -secret xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx &