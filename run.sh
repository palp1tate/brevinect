
nohup go run service/user/main.go -p 8081 > service/user/output.log &

nohup go run service/admin/main.go -p 8082 > service/admin/output.log &

nohup go run service/meeting/main.go -p 8083 > service/meeting/output.log &

nohup go run service/third/main.go -p 8084 > service/third/output.log &

nohup go run api/main.go -p 8080 > api/output.log &

sudo nginx -s reload -c /code/brevinect/nginx.conf

