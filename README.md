# fita

I use existing libs :

 - Chi Router
 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - Gorm, for ORM


# For setup after cloning the repo:
> cd fita
> go mod tidy

# to do a unit test :
> go to the package you want to testing then run a command "go test"
> you can see the coverage testing in each package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package

# for db table :
> in folder db, there is a .sql file with the create table command and insert command. I use postgresql for this case. you can run the command in your sql editor page.

# the endpoint
> here is the curl for the endpoint :
> curl --location --request POST 'http://localhost:8080/api/checkout' \
--header 'x-api-key: fita' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        {
            "sku": "120P90",
            "qty": 3
        }
    ] 
}'
> i use x-api-key in header, you can fill it with anything in example I fill it with "fita". it just an example if in assumsion the endp0int using a header middleware
> here is the postman link if you want to use postman instead : 
> https://www.getpostman.com/collections/dc6ab4a511f2a7d689b2
> or you can open apispec.yaml as well, for API Documentation
