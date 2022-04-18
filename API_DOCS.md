All the APIs except register api are intercepted and checked for token either via cookie as below: 
```
access_token=********; Path=/; Domain=jira-clone.com; HttpOnly; Expires=Mon, 04 Apr 2022 15:55:06 GMT;
```

OR as a **Authentication Header**
```
Authorization: Bearer ********.
```
# Updated Backend APIs Endpoints
- Any request that goes without token gets 401 status code in reply and response as follows:
    ```
    {
    "message": "ACTION_NOT_AUTHORIZED",
    "status": false
    }
    ```
- Any request with expired token gets 400 status code in reply and response as follows:
    ```
    {
    "message": "EXPIRED_TOKEN",
    "status": false
    }
    ```

-  **Create Projects**  
	*Request Format* -
``` 
	POST /api/project/create
	{
        "name" :"Project 6"
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
    {
        "message": "",
        "resp": {
            "project_name": "Project 6",
            "project_id": 7,
            "created_at": "2022-04-02T01:41:11.449013-04:00"
        },
    "status": true
    }
```  
<br/>

-  **List Projects**  
	*Request Format* -
``` 
	POST /api/project/list
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
         {
        "message": "",
        "resp": {
            "projects": [
                {
                    "name": "Project 1",
                    "id": 1,
                    "created_at": "2022-03-23T15:05:23.536313-04:00",
                    "user_role": 1
                },
                {
                    "name": "Project 5",
                    "id": 5,
                    "created_at": "2022-03-31T19:12:09.045909-04:00",
                    "user_role": 1
                },
                {
                    "name": "Project Sample PS1",
                    "id": 6,
                    "created_at": "2022-04-01T18:33:53.976626-04:00",
                    "user_role": 1
                },
                {
                    "name": "Project 6",
                    "id": 7,
                    "created_at": "2022-04-02T01:41:11.449013-04:00",
                    "user_role": 1
                }
            ]
        },
        "status": true
    }
```  
<br/>

-  **To get a project's info**  
	*Request Format* -
``` 
    GET /api/project/info
    {
     "project_id": 6,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "",
        "resp": {
            "project_id": 6,
            "project_name": "Project Sample PS1",
            "owner_uname": "pypalkar23",
            "owner_id": 1,
            "owner_fname": "Mandar",
            "owner_lname": "Palkar",
            "created_at": "2022-04-01T18:33:53.976626-04:00"
        },
        "status": true
    }
```  

<br/>

- **Delete a Project**	  
  *Request Format* -
``` 
    POST /api/project/delete
    {
        "project_id": 7,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "PROJECT_DELETE_SUCCESS",
        "resp": "",
        "status": true
    }
```  

- **List Members**	  
  *Request Format* -
``` 
    POST /api/project/members
    {
        "project_id": 1,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
    "message": "",
    "resp": {
        "members": [
            {
                "user_id": 1,
                "first_name": "Mandar",
                "last_name": "Palkar",
                "user_role": 1
            },
            {
                "user_id": 2,
                "first_name": "Ashish",
                "last_name": "Mhaske",
                "user_role": 2
            }
        ]
    },
    "status": true
}
```  

-  **Create Sprints**  
	*Request Format* -
``` 
	POST /api/sprint/create
	{   
        "sprint_name" :"Sprint 5",
        "project_id" : 1,
        "start_date": "2021-03-07",
        "end_date" : "2021-03-21"
    }

``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
	{
    "message": "",
    "resp": {
        "sprint_name": "Sprint 5",
        "sprint_id": 6
    },
    "status": true
    }
```  
<br/>

-  **List Sprints**  
	*Request Format* -
``` 
	GET /api/sprint/list
	{
    	  "project_id":1
	}
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
    
``` 
{
    "message": "",
    "resp": {
        "sprints": [
            {
                "name": "Sprint 1",
                "id": 1,
                "start_date": "2021-03-06T00:00:00Z",
                "end_date": "2021-03-21T00:00:00Z",
                "created_at": "2022-03-23T15:23:01.116584-04:00",
                "project_id": 1
            },
            {
                "name": "Sprint 3",
                "id": 3,
                "start_date": "2021-03-07T00:00:00Z",
                "end_date": "2021-03-21T00:00:00Z",
                "created_at": "2022-03-29T19:55:12.21721-04:00",
                "project_id": 1
            }
        ]
    },
    "status": true
}
```  

<br/>

-  **To get a sprint's info**  
	*Request Format* -
``` 
    POST /api/sprint/info
    {   
        "sprint_id" :3
    }

``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "",
        "resp": {
            "sprint_id": 3,
            "sprint_name": "Sprint 1",
            "created_at": "2022-03-29T19:55:12.21721-04:00",
            "start_date": "2021-03-07T00:00:00Z",
            "end_date": "2021-03-21T00:00:00Z"
        },
        "status": true
    }
```  
<br/>

- **Delete a Sprint**  	  
	*Request Format* -
``` 
    POST /api/sprint/delete
    {
     "sprint_id": 6,
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "SPRINT_DELETE_SUCCESS",
        "resp": "",
        "status": true
    }
``` 
<br/> 

-  **To Create an Issue**. 
	*Request Format* -
```
    POST /api/issue/create
   {    
        "issue_title":"sample issue title",
        "issue_description": "sample issue description",
        "issue_type": 1,
        "creator": 2,
        "assignee": 1,
        "sprint_id": 1,
        "project_id": 1
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "issue_id": 0,
            "issue_title": "sample issue title",
            "issue_description": "sample issue description",
            "issue_type": 1,
            "creator_name": "amhaske32",
            "assignee_name": "pypalkar23",
            "created_at": "2022-04-02T02:01:04.593536-04:00",
            "issue_status": 1
        },
        "status": true
    }

```
<br/>

- **To get a issue’s info**  
	*Request Format* -
```
    POST api/issues/info
    {
        "issue_id": 3
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "issue_id": 3,
            "title": "Sample Issue Edited 1",
            "type": 1,
            "sprint_id": 3,
            "sprint_name": "Sprint 1",
            "project_id": 1,
            "description": "Sample Text Edited 5",
            "assignee_id": 1,
            "assignee_name": "pypalkar23",
            "creator_id": 2,
            "creator_name": "amhaske32",
            "created_at": "2022-03-23T18:05:00.58531-04:00",
            "project_name": "Project 1",
            "issue_status": 2
        },
        "status": true
    }
```
<br/>

- **To list issues for sprint**  
 	*Request Format* -
```
    GET api/issue/list
    {
        "sprint_id": 1
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "issues": [
                {
                    "issue_id": 1,
                    "title": "issue title",
                    "status": 1,
                    "created_at": "2022-03-23T16:39:17.532173-04:00"
                },
                {
                    "issue_id": 4,
                    "title": "issue title 4",
                    "status": 1,
                    "created_at": "2022-03-29T22:53:55.005758-04:00"
                },
                {
                    "issue_id": 5,
                    "title": "sample issue title",
                    "status": 1,
                    "created_at": "2022-04-02T02:01:04.593536-04:00"
                }
            ]
        },
        "status": true
    }
```
<br/>

- **To delete an issue**  	  
 	*Request Format* -
``` 
    POST /api/issue/delete
    {
        "issue_id": 4
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "ISSUE_DELETE_SUCCESS",
        "resp": "",
        "status": true
    }
``` 

- **To update an issue**  	  
 	*Request Format* -
``` 
    POST /api/issue/delete
    {
        "issue_id": 3,
        "issue_title": "Sample Issue Edited 1",
        "issue_description": "Sample Text Edited 5",
        "issue_type": 1,
        "creator": 1,
        "assignee": 2,
        "sprint_id": 3,
        "project_id": 1
    }   
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "ISSUE_UPDATE_SUCCESS",
        "resp": {},
        "status": true
    }
``` 

- **To update an issue**  	  
	*Request Format* -
``` 
    POST /api/issue/move
    {
        "issue_id":3,
	    "project_id":1,
	    "status":2
    }
``` 

&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -

```    
    {
        "message": "ISSUE_UPDATE_SUCCESS",
        "resp": {},
        "status": true
    }
``` 

- **To Register a User**  
 	*Request Format* -
```
    POST /register
    {
        "username": "dipesh97",
        "password": "1c1d15237b2433f18f588d8bf6984751",
        "firstname": "Dipesh",
        "lastname": "Palkar",
        "email_id": "dipeshpalkar@gmail.com"
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "REGISTERATION_SUCCESSFUL",
        "resp": "",
        "status": true
    }
```
<br/>


- **To Login**  
	*Request Format* -
```
    POST /login
    {
        "username":"pypalkar23",
        "password":"**********"
    }
```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
        "message": "LOGIN_SUCCESSFUL",
        "resp": {
            "access_token": "******"
        },
        "status": true
    }
```
<br/>


- **To Add Comments To An Issue**  
    *Request Format* -
```
    POST /api/comment/add
    {
        "issue_id": 3,
        "comment": "This is a comment for issue_id: 3"
    }

```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
    {
        "message": "",
        "resp": {
            "comment" : "This is a comment for issue_id: 3"
        },
        "status": true
    }
```
<br/>

- **View Comments**  
    *Request Format* -
```
    POST /api/comment/view
    {
        "issue_id": 3
    }

```
&nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;*Response Format* -
```
   {
        "message": "",
        "resp": {
            "comments": [
                {
                    "comment_id": 1,
                    "comment": "This is a comment for issue_id: 3"
                }
            ]
        },
        "status": true
    }
```
<br/>