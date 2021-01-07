pipeline {
  environment{
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
        stash 'workspace'
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
            
            //sh "docker tag fiber_react-backend ${image_name} ${image_name}:${BUILD_ID}"
            // "docker login -u ${cred_USR} -p ${cred_PSW}"
            //sh "docker push vin1711/fiber_react-backend:${BUILD_ID}"
            //sh "docker push vin1711/fiber_react-backend"*//
            //when we run docker in this step, we're running it via a shell on the docker build-pod container, 
           //sh "docker push vividseats/promo-app:dev"        // which is just connecting to the host docker deaemon
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
        withCredentials([usernamePassword(credentialsId:'dockerCred',usernameVariable:'user',passwordVariable:'password')]){
              sh 'docker login -u $user -p $password'
              sh "docker push vin1711/fiber_react-backend:${BUILD_ID}"
              sh "docker push vin1711/fiber_react-backend"
            }
      }
    }
    stage('deploy to kubernetes'){
      steps{
         //unstash 'workspace'
         withCredentials([file(credentialsId:'fiberBackend',variable:'file')]){
           script{
             env.conf=sh(returnStdout:true,script:"cat $file")
           }
           script{
             data=readYaml(file:'kube/config-map.yaml.template')
             echo "${data}"
             data.data="[${conf}]"
             //datas="${data.data}"
             echo "${data.data}"
             //echo "${data}"
             //sh "rm -f kube/config-map.yaml"
             writeYaml(file:'kube/config-maps.yaml.temp',data:"${data}",overwrite:true)
             //writeYaml(file:'kube/config-map.yaml',data:"${data}",overwrite:true,charset:'collection')
             //writeYaml charset: 'string', data: "${data}", file: 'kube/config-map.yaml'
             data1=readYaml(file:'kube/config-maps.yaml.temp')
             //env.datas=sh(returnStdout:true,script:"cat /kube/config-maps.yaml")
             echo "${data1}"
             //echo "${datas}"
             //data=readYaml(file:'kube/config-maps.yaml')
           }    
           //kubernetesDeploy(configs: '**/*.yaml', kubeconfigId:'kubeConfig',secretNamespace:'jenkins',enableConfigSubstitution:true)
        }
      }
    }
  }
}

