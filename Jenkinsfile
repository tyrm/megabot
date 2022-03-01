pipeline {
  environment {
    registry = 'tyrm/megabot'
    registryCredential = 'docker-io-tyrm'
    dockerImage = ''
    gitDescribe = ''
  }

  agent any

  stages {

    stage('Test') {
      agent {
        docker {
          image 'gobuild:1.17'
          args '-e GOCACHE=/gocache -e HOME=${WORKSPACE} -v /var/lib/jenkins/gocache:/gocache -v /var/lib/jenkins/go:/go'
        }
      }
      steps {
        script {
          withCredentials([
            string(credentialsId: 'codecov-tyrm-supreme-robot', variable: 'CODECOV_TOKEN'),
            usernamePassword(credentialsId: 'integration-postgres-test', usernameVariable: 'POSTGRES_USER', passwordVariable: 'POSTGRES_PASSWORD'),
            string(credentialsId: 'integration-redis-test', variable: 'REDIS_PASSWORD')
          ]) {
            sh """#!/bin/bash
            go get -t -v ./...
            go test -race -coverprofile=coverage.txt -covermode=atomic ./...
            RESULT=\$?
            gosec -fmt=junit-xml -out=gosec.xml  ./...
            bash <(curl -s https://codecov.io/bash)
            exit \$RESULT
            """
          }
          junit allowEmptyResults: true, checksName: 'Security', testResults: "gosec.xml"
        }
      }
    }

    stage('Release') {
      when {
        buildingTag()
      }
      environment {
        GITHUB_TOKEN = credentials('gihub-tyrm-pat')
      }
      steps {
        script {
          sh 'goreleaser'
        }
      }
    }

  }

}
