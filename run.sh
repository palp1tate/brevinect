
nohup go run service/user/main.go > service/user/output.log &

nohup go run service/admin/main.go > service/admin/output.log &

nohup go run service/meeting/main.go > service/meeting/output.log &

nohup go run service/third/main.go > service/third/output.log &

nohup go run api/main.go -p 8080 > api/output.log &

sudo nginx -s reload -c /code/brevinect/nginx.conf

