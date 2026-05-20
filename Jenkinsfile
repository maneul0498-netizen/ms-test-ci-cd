pipeline {

    agent any

    environment {
        GOCACHE = "${WORKSPACE}/.gocache"
    }

    stages {

        stage('Checkout') {
            steps {

                deleteDir()

                git branch: 'main',
                url: 'https://github.com/maneul0498-netizen/ms-test-ci-cd'
            }
        }

        stage('Build user-service') {
            steps {
                sh '''
                    mkdir -p $GOCACHE

                    cd user-service

                    go mod tidy

                    go build -o app .
                '''
            }
        }

        stage('Build notification-service') {
            steps {
                sh '''
                    mkdir -p $GOCACHE

                    cd notification-service

                    go mod tidy

                    go build -o app .
                '''
            }
        }

        stage('Run Containers') {

            steps {
                sh '''
                    docker-compose down || true

                    CHANGED_FILES=$(git diff --name-only HEAD~1 HEAD)

                    echo "=== CHANGED FILES ==="
                    echo "$CHANGED_FILES"

                    if echo "$CHANGED_FILES" | grep '^user-service/'; then
                        echo "Building user-service..."
                        docker-compose build user-service
                    fi

                    if echo "$CHANGED_FILES" | grep '^notification-service/'; then
                        echo "Building notification-service..."
                        docker-compose build notification-service
                    fi

                    docker-compose up -d

                    docker ps
                '''
            }
        }

        stage('Wait Services') {

            steps {
                sh '''
                    sleep 10

                    docker ps
                '''
            }
        }

        stage('Integration Test') {

            steps {


                sh """
                    curl -v \
                    -H 'Content-Type: application/json' \
                    -d '{\"name\":\"manuel\"}' \
                    http://host.docker.internal:8081/users

                """
            }
        }
    }

    post {

        always {

            sh '''
                docker ps

                docker-compose logs

                docker-compose down
            '''
        }
    }
}