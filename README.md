# hangry-ghc-2024

## Initial set up
1. Checkout to the branch and run command ```go mod init main.go```
2. Run ```go mod tidy```

## Run the service
1. Run cadence service ```docker-compose up``` . Cadence UI will be available at http://localhost:8088/domains
2. Register the domain ```docker run --network=host --rm ubercadence/cli:master --do hangry-ghc24-domain domain register -rd 1```
3. Run the service ```go run main.go```
4. Run the workflow ```cadence --domain hangry-ghc24-domain workflow start --et 600 --tl hangry-worker --workflow_type main.eatsOrderWorkflow --input '["Biryani", "Sanya", false]'```
