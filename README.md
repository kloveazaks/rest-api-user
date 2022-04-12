# REST-API | Kloveazaks

# user-service

# REST API

GET /users -- list of users -- 200, 404, 500
GET /users/:id -- user by id -- 200, 404, 500
POST /users/:id -- create user -- 204, 4xx, Header Location: url
PUT /users/:id -- fully update user -- 204/200, 404, 400, 500
PATCH /users/:id -- partially update field user -- 204/200, 404, 400, 500
DELETE /users/:id -- delete user by id -- 204, 404, 400

200 -- Ok
204 -- No content
404 -- Not found
500 -- Internal Server 
