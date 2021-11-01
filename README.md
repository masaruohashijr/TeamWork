## List of tasks I would need more time to complete
- Documentation
- Testing 
- Deployment
- The rest of the repositories: Postgres, Redis, SQLite, etc.
- The rest of the api technologies: GraphQL, gRPC, RSocket, REST, SOAP.

## Instructions

The goal of this exercise is to create a demo RESTful application using Golang.

### The Task

In this task, we are building backend of an application that helps us managing our team using REST.

### Features and Requirements
- A member has a name and a type the late one can be an employee or a contractor.
- - If it's a contractor, we want to store the the duration of the contract as an integer.
- - If it's an employee, we need to store their role, for instance: Software Engineer, Project Manager and so on.
- A member can be tagged, for instance: C#, Angular, General Frontend, Seasoned Leader and so on. (Tags will likely be used as filters later, so keep that in mind)
- We need to offer a REST CRUD for all the information above.

#### Notes:

1. You can use any Golang framework
2. Make sure to provide a tutorial on how to run your application
3. Feel free to use any database

## Evaluation
| Functionality     |                                                                | Possible Points |
|-------------------|----------------------------------------------------------------|-----------------|
|                   | Matches the proposed requirements                              |              20 |
|                   | Implements REST correctly                                      |              15 |
|                   | Separation of business logic and persistence layers            |              15 |
|                   | Input validations                                              |               5 |
|                   | Standard HTTP error codes                                      |               5 |
| **Code Quality**  |                                                                |                 |
|                   | Code formatting, readability, maintainability, etc             |               5 |
|                   | Folders and packages structure                                 |               5 |
| **DevOps**        |                                                                |                 |
|                   | Docker image to build/run the project                          |              10 |
|                   | DB migrations                                                  |               5 |
| **Documentation** |                                                                |                 |
|                   | Documentation about the work done, how to run the project, etc |               5 |
| **Testing**       |                                                                |                 |
|                   | Has tests                                                      |              10 |
| **Total**         |                                                                |             100 |


### Bonus Points:
1. If you deploy the application in any server and share the link with us
2. If provide thoughts on what you could improve on your code given more time and incentives

### Is it necessary build a frontend?
No, this is a simply backend exercise.

### How do you evaluate the exercise?
For every exercise we have two senior backend engineers from our team reviewing the code and the functionality and giving a score for each line item as shown in the previous table.

### How can I deliver the exercise?
To deliver the exercise, you should clone this repository and work on a new branch. When you'll consider it completed, just push the branch and open a Merge Request.

### Will I have access to the evaluation?
By default we only send the result, however you can feel free to request the full evaluation and we will share it with you as well as the final score.
