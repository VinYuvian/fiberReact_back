pipeline {
  agent {
    kubernetes {
      cloud 'kubernetes'
      label 'promo-app'  // all your pods will be named with this prefix, followed by a unique id
      idleMinutes 5  // how long the pod will live after no jobs have run on it
      yamlFile 'pod.yaml'  // path to the pod definition relative to the root of our project 
      defaultContainer 'golang'  // define a default container if more than a few stages use it, will default to jnlp container
      podRetention never()
    }
  }
  stages {
    stage('Build') {
      steps {  // no container directive is needed as the maven container is the default
        sh "go version"   
        sh "go build"
      }
    }
    stage('Build Docker Image') {
      steps {
        container('docker') {  
          sh "docker build -t vin1711/fiberReact-backend:${env.BUILD_ID} ."  // when we run docker in this step, we're running it via a shell on the docker build-pod container, 
          //sh "docker push vividseats/promo-app:dev"        // which is just connecting to the host docker deaemon
        }
      }
    }
  }
}

