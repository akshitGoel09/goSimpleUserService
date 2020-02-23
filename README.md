# goSimpleUserService
Creating a simple user service with CURD operations for user with first and last name as part of tutorial to learn Go


## Running application
To run the project clone it and from inside the directory run
```
    go build .
```
The above command should create a goSimpleUserService in case of mac or goSimpleUserService.exe in case of windows.
 Run it in terminal in case of mac like as below
```
./goSimpleUserService
```
The application then start on **3000** port on **localhost**.

## Application Details
The user object the application works with
### user object
```
{
  "ID" : 1,
  "FirstName" : "Akshit",
  "LastName" : "Goel"
}
```

The operations supported by the Application

### Add user

```
Post Request
URL - http://localhost:3000/users
Body -
{
	"FirstName" : "Akshit",
	"LastName" : "Goel"
}

Sample Response -
{
  "ID":1,
  "FirstName":"Akshit",
  "LastName":"Goel"
}
```

### Get All Users

```
Get Request
URL - http://localhost:3000/users

Sample Response -
[
  {
    "ID":1,
    "FirstName":"Akshit",
    "LastName":"Goel"
  }, {
    "ID":2,
    "FirstName":"Akshit2",
    "LastName":"Goel2"
  }
]
```

### Get user by ID

```
Get Request
URL - http://localhost:3000/users/1

Sample Response -
{
  "ID":1,
  "FirstName":"Akshit",
  "LastName":"Goel"
}
```

### Update user

```
Put Request
URL - http://localhost:3000/users/1
Body -
{
  "ID" : 1,
	"FirstName" : "UpdatedUserFirstName",
	"LastName" : "UpdatedUserLastName"
}

Sample Response -
{
  "ID":1,
  "FirstName":"UpdatedUserFirstName",
  "LastName":"UpdatedUserLastName"
}
```

### Delete user

```
Delete Request
URL - http://localhost:3000/users/1
Sample Response -
Status 200 ok
```

The application also handles various errors as and when encountered.
