pipeline {
  environment {
    PATH = '/go/bin:~/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin'
    composeFile = "deployments/docker-compose-integration.yaml"
    networkName = "network-${env.BUILD_TAG}"
    registry = 'tyrm/megabot'
    registryCredential = 'docker-io-tyrm'
    dockerImage = ''
    gitDescribe = ''
  }

  agent any

  stages {

    stage('Build Static Assets') {
      steps {
        script {
          sh """#!/bin/bash
          make clean
          make npm-scss
          make stage-static
          """
        }
      }
    }

    stage('Start External Test Requirements'){
      steps{
        script{
          retry(2) {
            sh """NETWORK_NAME="${networkName}" docker-compose -f ${composeFile} pull
            NETWORK_NAME="${networkName}" docker-compose -p ${env.BUILD_TAG} -f ${composeFile} up -d"""
          }
          parallel(
            mysql: {
              retry(30) {
                sleep 1
                sh "docker run -t --rm --network=${networkName} subfuzion/netcat -z mariadb 3306"
              }
            },
            postgres: {
              retry(30) {
                sleep 1
                sh "docker run -t --rm --network=${networkName} subfuzion/netcat -z posgres 5432"
              }
            },
            redis: {
              retry(30) {
                sleep 1
                sh "docker run -t --rm --network=${networkName} subfuzion/netcat -z redis 6379"
              }
            }
          )
        }
      }
    }

    stage('Test') {
      agent {
        docker {
          image 'gobuild:1.18'
          args '--network ${networkName} -e HOME=${WORKSPACE} -v /var/lib/jenkins/go:/go'
          reuseNode true
        }
      }
      steps {
        script {
          withCredentials([
            string(credentialsId: 'codecov-tyrm-megabot', variable: 'CODECOV_TOKEN'),
            file(credentialsId: 'tls-localhost-crt', variable: 'MB_TLS_CERT'),
            file(credentialsId: 'tls-localhost-key', variable: 'MB_TLS_KEY')
          ]) {
            sh """#!/bin/bash
            go test --tags=postgres,redis,mysql,mssql -race -coverprofile=coverage.txt -covermode=atomic ./...
            RESULT=\$?
            #gosec -fmt=junit-xml -out=gosec.xml  ./...
            bash <(curl -s https://codecov.io/bash)
            exit \$RESULT
            """
          }
          junit allowEmptyResults: true, checksName: 'Security', testResults: "gosec.xml"
        }
      }
    }

    stage('Build Release') {
      agent {
        docker {
          image 'gobuild:1.18'
          args '--network ${networkName} -e HOME=${WORKSPACE} -v /var/lib/jenkins/go:/go'
          reuseNode true
        }
      }
      when {
        buildingTag()
      }
      steps {
        script {
          withCredentials([
            usernamePassword(credentialsId: 'gihub-tyrm-pat', usernameVariable: 'GITHUB_USER', passwordVariable: 'GITHUB_TOKEN')
          ]) {
            sh 'make build'
          }
        }
      }
    }

    stage('Build Snapshot') {
      agent {
        docker {
          image 'gobuild:1.18'
          args '--network ${networkName} -e HOME=${WORKSPACE} -v /var/lib/jenkins/go:/go'
          reuseNode true
        }
      }
      when {
        not {
          buildingTag()
        }
      }
      steps {
        script {
          sh 'make build-snapshot'
        }
      }
    }

  }

  post {
    always {
      sh """NETWORK_NAME="${networkName}" docker-compose -p ${env.BUILD_TAG} -f ${composeFile} down"""
    }
  }

}
