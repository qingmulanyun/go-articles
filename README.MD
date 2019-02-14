# Article API

## App Information

- Language: Golang
- DB: postgres
- Test: Postman/curl

## System Setup Information

1. Install Golang/postgres, start postgres
2. Create database with `user: postgres` `db:goarticles`
3. `cd go-articles` and compile: `go build`
4. start server: `go-articles`

## Description

Basically it's building a Restful API server based Golang. It only has one resource ``articles`` with 3 operations:

1. Get with id
2. Post wit required params
3. Get index list with 2 params.

- Structure 

    It's a typical MVC  architectural pattern project. So the most import things would be `controllers` `models` `serializers` folders.
  
    `controllers` holds all the controllers of business logic.
    
    `models` holds all the data models.
    
    `serializers` here is standing for the `V(view)` because this is an API server. We only render data into JSON format.
    
    
- Dependencies

    Use `mux` to handler the router
    
    Use `gorm` to map model with database.
    
- Points
     
    1. Add version to the API. In this way, we can make api updates without breaking old clients. So the API endpoints have been changed to
          
          GET  http://localhost:8000/v1/articles/{id}
          
          GET  http://localhost:8000/v1/tags/{tag_name}/{date}
          
          POST http://localhost:8000/v1/articles
     
    2. Add index to tag and date to improve the final endpoint performance. 
    
    
- Things to improve if given more time

    1. Add user model and authentication/authorization. (with JWT)
    
    2. Add more validations(model level and parameters level)
    
    3. Use UUID instead of uint as ID.
    
    4. Couldn't hide all the timestamps fields of the articles model even add `gorm"-"`, need to do more research.