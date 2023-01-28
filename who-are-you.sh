curl -s https://01.alem.school/assets/superhero/all.json --data '{"query":"{user(where:{login:{_eq:\"choumi\"}}){id}}"}' | jq ".[52] | .name"

