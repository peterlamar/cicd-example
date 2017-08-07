pipeline {
    agent any

    stages {
        stage('Build Bin') {
            steps {
                sh 'build/build-bin.sh'
            }
        }
        stage('Build War') {
            steps {
                sh 'build/run-tests.sh'
            }
        }
    }
}
