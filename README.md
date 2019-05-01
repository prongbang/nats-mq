# Install NATS Server Docker

```
docker-compose up -d
```

### Start MQ

```
make start
```

### Test the NATS server to verify it is running.

```
telnet localhost 4222
```

Expected result:

```
Trying ::1...
Connected to localhost.
Escape character is '^]'.
INFO {"server_id":"fGRZg3AU5NEIuvbbQyRt7L","version":"1.3.0","proto":1,"git_commit":"eed4fbc","go":"go1.11","host":"0.0.0.0","port":4222,"max_payload":1048576,"client_id":1}
```

You can also test the monitoring endpoint with a browser.

```
http://localhost:8222
```

### Start Web server

```
make run
```

### Notification Promotion to App2

```
http://localhost:8081/api/v1/noti-promotion
```

### Download

```
go get -u github.com/prongbang/nats-mq
```