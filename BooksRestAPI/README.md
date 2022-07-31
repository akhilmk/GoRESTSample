GET 
    http://localhost:8000/api/books
    http://localhost:8000/api/books/1
POST
    http://localhost:8000/api/books 
   
    Headers
    Content-Type : application/json
   
    Body : raw
    {
    "isbn":"1234",
    "Title":"posted book",
     "Author":{"firstname":"fn","lastname":"ln"}    
    }
Delete
    http://localhost:8000/api/books/1
PUT
    http://localhost:8000/api/books/{id}