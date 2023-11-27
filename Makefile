PORTS = 8081 8082 8083 8084 8080

.PHONY: run stop restart-nginx

run:
	nohup go run service/user/main.go -p 8081 > service/user/output.log 2>&1 &

	nohup go run service/admin/main.go -p 8082 > service/admin/output.log 2>&1 &

	nohup go run service/meeting/main.go -p 8083 > service/meeting/output.log 2>&1 &

	nohup go run service/third/main.go -p 8084 > service/third/output.log 2>&1 &

	nohup go run api/main.go -p 8080 > api/output.log 2>&1 &

stop:
	for port in $(PORTS); do \
		pids=$$(lsof -t -i:"$$port"); \
		if [ -n "$$pids" ]; then \
			for pid in $$pids; do \
				if kill -9 "$$pid"; then \
					echo "Closed service running on port $$port with pid $$pid"; \
				else \
					echo "Failed to close service on port $$port with pid $$pid"; \
				fi; \
			done; \
		else \
			echo "No process found on port $$port"; \
		fi; \
	done; \
	echo "Services closed."

restart-nginx:
	sudo nginx -s reload -c /code/brevinect/nginx.conf
