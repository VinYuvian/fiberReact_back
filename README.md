# fiberReact_back


This project uses Fiber framework which is an express inspired web framework designed for Golang.
Main purpose of this project is to create rest api's for CRUD operations on a mongo-db databse.

Dockerfile is used to specify the image specifications for docker build.
Jenkinsfile is used for setting up a CI/CD pipeline to deploy the application to a kubernetes cluster hosted in Digital Ocean using a kubernetes-cd plugin.
Jenkins is hosted in kubernetes cluster and has dynamic container set up as agents for the golang app build,docker image creation and push.
The CI/CD flow as below :
1. A commit is made (github)
2. Jenkins is notified with a webhook
3. Source code checkout (Jenkins)
4. build process (Jenkins slave container - golang container)
5. Image build process (Jenkins slave container - docker)
6. Image push process (Jenkins slave container - docker - dockerhub)
7. Deployment ( Jenkins kubernetes-cd plugin)
