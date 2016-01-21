#ansible-slack



Test:
curl -H "Content-Type: application/json" -d '{"Target":"localhost","Module":"shell","Command":"df -h"}' http://localhost:8080/execute
