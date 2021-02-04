## Post struct
```json
    {
        "id": string,
        "userId": string,
        "envData":
        {
            "humidity": string,
            "temperature": string,
            "light": string,
        },
        "images": string[],
        "date": string
    }
```

## Get single post
- method: `GET`
- path: `/post/{postId}`
- return: 
    ```json
    {
        "Posts": []post, //use first element of array
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```


## Get user's all posts
- method: `GET`
- path: `/post/user/{userID}`
- return: 
    ```json
    {
        "Posts": []post,
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```

## Create new post
- method: `POST`
- path: `/post/new`
- parameters: 
    ```json
        {
            "id": string,
            "userId": string,
            "envData":
            {
                "humidity": string,
                "temperature": string,
                "light": string,
            },
            "images": string[],
            "date": string
        }
    ```
- return:
    ```json
    {
        "Posts": []post, // Will return NULL
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```

## Delete post
- method: `DELETE`
- path: `/post/{postId}`
- return:
    ```json
    {
        "Posts": []post, // Will return NULL
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```