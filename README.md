# URL-shortener

## Description

This is an URL shortener using Go programming language.

There are 2 APIs.

* POST /api/v1/urls
  * Request

    ```json
    {
        "url": "<original_url>",
        "expireAt": "2021-02-08T09:20:41Z"
    }
    ```

  * Response
  
    ```json

    {
        "id": "<url_id>",
        "shortUrl": "http://localhost/<url_id>"
    }
    ```

  * Detailed description: After sending the request, server will check url validity, and then generate a random string with 6 letters to be url_id. Finally, insert this url object to database, and response to client.

* GET /:url_id
  * Redirect to original URL
  * Detailed description: Server will first check url_id exists or not. If exists, check expired time and redirect to original url.

## Set up

1. Clone project

    ```shell
    git clone https://github.com/gary1030/URL-shortener.git
    ```

2. Change directory

   ```shell
   cd URL-shortener
   ```

3. Copy configuration files

    ```shell
    cp .env.example .env
    ```

4. Edit `.env` file

    ```txt
    PG_HOST=localhost
    PG_PORT=5432
    PG_USERNAME=
    PG_PASSWORD=
    PG_DBNAME=URL
    DOMAIN_NAME=localhost
    ```
    
5. Install packages

    ```shell
    go install
    ```

6. Start backend service

    ```shell
    go run .
    ```

7. Test

   You may go to `http://localhost:8080/ABCDEF` to check whether setting is correct.

## Technique

* Go
* Gorm
  * Save quite a bit of tedious sql coding
  <!-- * Slow down application rather than SQLite -->
* Gin
  * Known for high-performing APIs
  * A lean and simple framework
  * Awful router pattern matching, but we have only two APIs, so it's fine.
* PostgreSQL
  * Familiar to me
