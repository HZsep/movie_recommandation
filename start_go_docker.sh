FILEPATH=$(dirname "$0")
docker run -t -i --rm -v $FILEPATH:/go golang:1.22rc2-bookworm bash


