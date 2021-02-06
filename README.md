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
- returns: 
    ```js
    {
        "Posts": []post, //use first element of array
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```


## Get user's all posts
- method: `GET`
- path: `/post/user/{userID}`
- returns: 
    ```js
    {
        "Posts": []post,
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```

## Create
- path: `/post/new`
- method: `POST`
- parameters: 
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | x-auth-token | authentication token of the user  | String |

    ```js
        {
            "userId": string,
            "envData":
            {
                "humidity": string,
                "temperature": string,
                "light": string,
            },
            "images": string[],
        }
    ```
- returns:
    ```js
    {
        "Posts": []post, // Will return NULL
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```

## Delete
- path: `/post/{postId}`
- method: `DELETE`
- params:
   * Header:
  
        |  Name | Description                           | Type   |
        |:---------:|---------------------------------------|--------|
        | x-auth-token | authentication token of the user  | String |

- returns:
    ```js
    {
        "Posts": []post, // Will return NULL
        "Success": integer, // 1 for success, 0 for failure
        "Message": string
    }
    ```
