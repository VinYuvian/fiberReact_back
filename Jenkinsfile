pipeline {
  options{
    skipDefaultCheckout()
  }
  environment{
      db_user = credentials('db_user')
      db_info = credentials('mongo_db_details')
      image_name = 'vin1711/fiber_react-backend'
    }
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
    stage('checkout'){
      steps{
        git branch:'main',url:'https://github.com/VinYuvian/fiberReact_back.git'
        //stash 'workspace'
      }
    }
    stage('Build') {
      steps {  // no container directive is needed as the maven container is the default
        sh "go version"   
        sh "go build -o Fiber_${BUILD_ID}"
      }
    }
    stage('Build Docker Image') {
      steps {
          container('docker') {  
            sh "docker build -t ${image_name} -t ${image_name}:${BUILD_ID} --build-arg BUILD_NUMBER=${BUILD_ID} ."
            script{
                env.choice = input message:"please select how to proceed",parameters:[choice(name:'build_type',
                                                                                         choices:'test\nprod\nstage\ntestPipeline',
                                                                                         description:'Is it a pipeline check or a deployment step?')]
            }
         }
      }
    }
    stage('Push Image'){
      when{
        expression{
          env.choice == 'prod'
        }
      }
      steps{
        container('docker'){
            withCredentials([usernamePassword(credentialsId:'dockerCred',usernameVariable:'user',passwordVariable:'password')]){
                sh 'docker login -u $user -p $password'
                sh "docker push ${image_name}:${BUILD_ID}"
                sh "docker push ${image_name}:latest"
            }
        }
      }
    }
    stage('deploy to kubernetes'){
      steps{ 
         //unstash 'workspace'
           kubernetesDeploy(configs: '**/*.yaml', kubeconfigId:'kubeConfig',secretNamespace:'jenkins',enableConfigSubstitution:true,deleteResource:true)
           //kubernetesDeploy(configs: '**/*.yaml', kubeconfigId:'kubeConfig',secretNamespace:'jenkins',enableConfigSubstitution:true)
      }
    }
   }
 }


