#ansible-slack

This is simply something I wanted to see if I could do. If there's serious consideration for putting this into production .. please make sure you add token-auth and limit the commands that can be run


Test:

curl -H "Content-Type: application/json" -d '{"Target":"localhost","Module":"shell","Command":"df -h"}' http://localhost:8080/execute
