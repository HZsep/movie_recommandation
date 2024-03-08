FILEPATH=$(dirname "$0")
docker run -t -i -p 8000:8080 --rm -v $FILEPATH:/go golang:1.22rc2-bookworm bash


