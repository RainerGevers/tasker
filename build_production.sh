export version=$(cat VERSION | grep "^VERSION=" | cut -d'=' -f 2 | awk '{print $1}' | tr -d '\040\011\012\015' | tr -d \r)
docker build -t rainerza/tasker:latest -t rainerza/tasker:0.0.5 .
#docker push rainerza/tasker:latest rainerza/tasker:"$(version | tr -d "\r")" .
