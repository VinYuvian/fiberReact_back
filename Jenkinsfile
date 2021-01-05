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
        sh "go build -o Fiber_${BUILD_ID}"
      }
    }
    stage('Build Docker Image') {
      steps {
        container('docker') {  
          withCredentials([usernamePassword(credentialsId:'dockerCred',usernameVariable:'user',passwordVariable:'password')]){
            sh "docker build -t vin1711/fiber_react-backend -t vin1711/fiber_react-backend:${BUILD_ID} --build-arg BUILD_NUMBER=${BUILD_ID} ." 
            //sh "docker tag fiber_react-backend vin1711/fiber_react-backend vin1711/fiber_react-backend:${BUILD_ID}"
            sh "docker login -u ${user} -p ${password}"
            sh "docker push vin1711/fiber_react-backend:${BUILD_ID}"
            sh "docker push vin1711/fiber_react-backend"/// when we run docker in this step, we're running it via a shell on the docker build-pod container, 
          //sh "docker push vividseats/promo-app:dev"        // which is just connecting to the host docker deaemon
          }
        }
      }
    }
  }
}

