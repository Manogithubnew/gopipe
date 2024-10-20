pipeline {
   agent any


   tools {
      go '1.23.2'
   }
   environment {
       DOCKERHUB_CREDENTIALS = credentials('doocker-hub-credential')
       DOCKER_IMAGE = 'ephraimaudu/test-app'
       GITHUB_CREDENTIALS = 'github'
       SONAR_TOKEN = credentials('SONAR_TOKEN')
       SNAP_REPO = 'vprofile-snapshot'
       NEXUS_USER = 'admin'
       NEXUS_PASS = 'admin@123'
       RELEASE_REPO = 'vprofile-release'
       CENTRAL_REPO = 'vpro-maven-central'
       NEXUS_IP = '192.168.29.68'
       NEXUS_PORT = '8081'
       NEXUS_GRP_REPO = 'vpro-maven-group'
       NEXUS_LOGIN = 'nexuslogin'
       SONARSERVER = 'sonarserver'
       SONARSCANNER = 'sonarscanner'
   }


   stages{
       stage('Checkout'){
           steps{
               echo "checking out repo"
               git url: 'https://github.com/Manogithubnew/gopipe.git', branch: 'main',
               credentialsId: "${GITHUB_CREDENTIALS}"
           }
       }
       stage('Run SonarQube Analysis') {
           steps {
               script {
                   echo 'starting analysis'
                   sh '/usr/local/sonar/bin/sonar-scanner -X -Dsonar.organization=eph-test-app -Dsonar.projectKey=eph-test-app-test-go-app -Dsonar.sources=. -Dsonar.host.url=https://sonarcloud.io'
               }
           }
       }
       stage('Run Docker Build'){
           steps{
               script{
                    echo "starting docker build"
                    sh "docker build build -t ${DOCKER_IMAGE}:${env.BUILD_ID} ."
                    echo "docker built successfully"
               }
           }
       }
       stage('push to docker hub'){
           steps{
               echo "pushing to docker hub"
               script{
                   docker.withRegistry('https://index.docker.io/v1/', 'dockerhub'){
                       docker.image("${DOCKER_IMAGE}:${env.BUILD_ID}").push()
                   }
               }
               echo "done"
           }
       }
   }


   post {
       always{
           cleanWs()
       }
   }
}
