# Image Gallery API

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

## Evaluation

Your solution will be evaluated on the following criteria:

- Does it implement all required endpoints with correct functionality?
- Is the code well-organized, modular, and readable?
- Is the error handling robust and informative?
- Is the database schema well-designed and efficient?
- Is the API documentation clear and concise?
- Are best practices for RESTful API design followed?

Copyright, Max Base 2023
