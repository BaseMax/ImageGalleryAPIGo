# Image Gallery API

This project is a RESTful API built with Go for managing an image gallery. It allows users to upload images, view a list of all images, view a single image, update image metadata, and delete images. The metadata for each image includes a title, description, and tags, and is stored in a database of your choice. The API follows best practices for RESTful API design, using appropriate HTTP methods and status codes and handling errors gracefully. The code is well-organized, modular, and readable, with clear documentation for the API endpoints. Use this project as a starting point for building your own Go-based RESTful APIs for image management.

## Project

You will build a RESTful API that allows users to manage and view an image gallery. The API will allow users to upload images, view a list of all images, view a single image, update image metadata, and delete images. The metadata for each image will include a title, description, and tags.

## Requirements

- Use the net/http package to handle HTTP requests and responses.
- Use a database of your choice to store image metadata.
- Use the multipart/form-data content type to handle image uploads.
- Implement the following endpoints:
  - `POST /api/images`: Upload a new image and its metadata.
  - `GET /api/images`: Retrieve a list of all images with their metadata.
  - `GET /api/images/:id`: Retrieve a single image with its metadata.
  - `PUT /api/images/:id`: Update the metadata for a single image.
  - `DELETE /api/images/:id`: Delete a single image.
- Handle errors gracefully and provide appropriate HTTP status codes in responses.
- Write clear and concise documentation for the API endpoints.
- Use best practices for RESTful API design, including proper HTTP methods and status codes.

## Endpoints

This API have the following endpoints:

- `GET /api/images/`

Returns a JSON array of all images and it's metadata.

```console
curl -s http://localhost:8080/api/images/ | jq
```

Response body:

```json
[
  {
    "id": 1,
    "title": "My Image",
    "description": "This is my image",
    "tags": "tag1,tag2",
    "file_path": "uploads/1680919323489361965.png",
    "created_at": "2023-04-08 05:32:03"
  },
  {
    "id": 2,
    "title": "My Image",
    "description": "This is my image",
    "tags": "tag1,tag2",
    "file_path": "uploads/1680919362135848916.png",
    "created_at": "2023-04-08 05:32:42"
  },
  {
    "id": 3,
    "title": "My Image",
    "description": "This is my image",
    "tags": "tag1,tag2",
    "file_path": "uploads/1680919369929026576.png",
    "created_at": "2023-04-08 05:32:49"
  }
]
```

- `GET /api/images/{id:[1-9]+}`

Returns the JSON representation of a single image and it's metadata specified by the id parameter.

```console
curl -s http://localhost:8080/api/images/1 | jq
```

Response Body:

```json
{
  "id": 1,
  "title": "My Image",
  "description": "This is my image",
  "tags": "tag1,tag2",
  "file_path": "uploads/1680919323489361965.png",
  "created_at": "2023-04-08 05:32:03"
}
```

- `POST /api/images/`

Adds a new image with it's metadata.

Request command(curl):

```console
curl -s -X POST -H "Content-Type: multipart/form-data" \
            -F "image=@/home/max/Pictures/golang/1.png" \
            -F "title=My Image" \
            -F "description=This is my image" \
            -F "tags=tag1,tag2" http://localhost:8080/api/images/ | jq
```

Response Body:

```json
{
  "id": 1,
  "title": "My Image",
  "description": "This is my image",
  "tags": "tag1,tag2",
  "file_path": "uploads/1680919323489361965.png",
  "created_at": "2023-04-08 05:32:03"
}
```

- `PUT /api/images/{id:[1-9]+}`

Updates an existing image metadata specified by the id parameter.

Request command(curl):

```console
curl -X PUT -H "Content-Type: application/json" \
            -d '{"title":"New Title", "description":"New Description", "tags":"new, tags"}' \
            http://localhost:8080/api/images/1

```

Response Body:

```json
{
  "id": 0,
  "title": "New Title",
  "description": "New Description",
  "tags": "new, tags",
  "file_path": "",
  "created_at": ""
}
```

- `DELETE /api/images/{id:[1-9]+}`

Deletes an existing image and it's metadata specified by the id parameter.

```console
curl -X DELETE http://localhost:8080/api/images/1
```

Response Body:

```json
{
  "message": "Image with ID 1 has been deleted"
}
```

## Evaluation

Your solution will be evaluated on the following criteria:

- Does it implement all required endpoints with correct functionality?
- Is the code well-organized, modular, and readable?
- Is the error handling robust and informative?
- Is the database schema well-designed and efficient?
- Is the API documentation clear and concise?
- Are best practices for RESTful API design followed?

Copyright, Max Base, MaxianEdison 2023
