curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"raqazhet\"}}){id}}"}' | jq ".data.user | .[] | .id"
