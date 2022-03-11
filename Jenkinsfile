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
          args '-e GOCACHE=/gocache -e HOME=${WORKSPACE} -v /var/lib/jenkins/gocache:/gocache -v /var/lib/jenkins/go:/go -v /var/lib/jenkins/.npm:/.npm'
        }
      }
      steps {
        script {
          withCredentials([
            string(credentialsId: 'codecov-tyrm-megabot', variable: 'CODECOV_TOKEN'),
            usernamePassword(credentialsId: 'integration-postgres-test', usernameVariable: 'POSTGRES_USER', passwordVariable: 'POSTGRES_PASSWORD'),
            string(credentialsId: 'integration-redis-test', variable: 'REDIS_PASSWORD')
          ]) {
            sh """#!/bin/bash
            make clean
            make clean-npm
            make npm-install-jenkins
            make npm-scss
            go get -t -v ./...
            go test -race -coverprofile=coverage.txt -covermode=atomic ./...
            RESULT=\$?
            /go/bin/gosec -fmt=junit-xml -out=gosec.xml  ./...
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
      steps {
        script {
          withCredentials([
            usernamePassword(credentialsId: 'gihub-tyrm-pat', usernameVariable: 'GITHUB_USER', passwordVariable: 'GITHUB_TOKEN')
          ]) {
            sh '~/go/bin/goreleaser'
          }
        }
      }
    }

  }

}
