# golang-test-task

To run:
    docker build -t datacollector .
    docker run -p 8081:8081 --name test --rm datacollector

To stop:
    docker stop test

Request Example:
    curl -H "Content-Type: application/json" -X POST -d '["https://golang.org/pkg/regexp/", "https://golang.org/pkg/fmt/", "http://symfony.com/", "https://launchpad.net/gnuflag"]' localhost:8081