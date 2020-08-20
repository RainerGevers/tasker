export version=$(cat VERSION | awk '{print $1}' | tr -d '\040\011\012\015' | tr -d \r\n)
docker build -t rainerza/tasker:latest -t rainerza/tasker:$version .

docker push rainerza/tasker:latest
docker push rainerza/tasker:$version
