# SimpleServer

# which additional requirement you chose
`Add user-based access control to your files such that only the user that originally
uploaded the file can access it.`

####Have I chosen the others
1. `Add token-based access control to your files such that instead of the identifier,
   files can be accessed with a token that expires after a set period of time.`
   - I would have leveraged s3 buckets Presigned Object Url
2. `Add an endpoint that returns a list of all files in the system, their identifier, original
    filename, and the byte size of the file.`
    - I just need to create this endpoint without the username constrains
3. `Automate the setup of all infrastructure (servers, cloud services, code, etc) such
    that you could easily deploy a second complete, working copy of your app in a
    command or two.`
    - I would have used docker, cloud formations and jenkins and have them trigger on code change

# how to compile/build/run the code locally
**Prerequisits**

Set up Aws CLI with valid credentials for (read/write) for s3 bucket and (read/write/create) for DynamoDB.

Dynamodb Table will be created automatically if it does not exist.

S3 bucket should be renamed in project. 
- `service/fileRecordS3Service.go`

Side note: if I had more time I would have added local dynamodb and s3 mock to this project.

Compile/build/run

 1. `dep ensure`
 2. `go run main.go`

# where to access the deployed version of the project

**Server**:  `simpleserver.alejand.com:8080`

# all design / architectural / technical decisions

##Architecture
Full disclosure this is my first time using Go. 

I decided to use gorilla mux as my router, based on my research it seemed easy. 
I separated the file structure similar to how I would do it with java spring. 
Config files are instantiated on start up and we use dependency injection to supply the endpoints.
Once I realized that Go treats package as the API I separated each endpoint to their own file to keep the code clean. 
I created services for my AWS calls and seperated out reusable structs.

I added an authentication interceptor in the router in order to achieve basic authentication. Normally I would use JWT but with time constrains I didn't want to spent to much time here.

I decided to leverage AWS entirely, I am very comfortable with the cloud and it really let me concentrate on learning GO.

I added an arbitrary 5MB limit to the file upload so you don't take down my server.

I did not add any unit tests, normally I would but because of my lack of knowledge of GO I spent most of my time stumbling around learning it.

## Aws Resources
- DB: Dynamodb, an easy to use scalable data store for quick projects
- S3: Easy to use scalable storage service

## Rest API
**Server**:  `simpleserver.alejand.com:8080`

- Authentication is Basic auth 
    - Example: `Authorization: Basic YmFzaWNBdXRoOmJhc2VQYXNzd29yZA==`
    - Password does not matter
    - Username is attached to file
    

| **URL** | **Method** | **Auth required** | **SampleResponse** |
| --- | --- | --- | --- |
| `/file` | `GET` | YES | [sample](#get-all-files-success-response) |
| `/file` | `POST` | YES | [sample](#upload-file) |
| `/file{FileId}` | `GET` | YES | Downloads a file |

### Get All Files Success Response

For a Username basicAuth

```json
[
    {
        "FileId": "QYb1v5iYKF8IRULkTVWvLa",
        "User": "basicAuth",
        "FileName": "Abstract-Colors-4K-Wallpaper-3840x2160.jpg",
        "FileSize": 1736545,
        "FileLocation": "https://s3.us-east-2.amazonaws.com/gigamog.simple.server.2018/QYb1v5iYKF8IRULkTVWvLa"
    }
]
```

### Upload File Success Response
```json
{
    "FileId": "QYb1v5iYKF8IRULkTVWvLa",
    "User": "basicAuth",
    "FileName": "Abstract-Colors-4K-Wallpaper-3840x2160.jpg",
    "FileSize": 1736545,
    "FileLocation": "https://s3.us-east-2.amazonaws.com/gigamog.simple.server.2018/QYb1v5iYKF8IRULkTVWvLa"
}
```

