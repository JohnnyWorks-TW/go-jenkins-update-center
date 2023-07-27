# Jenkins UpdateCenter Modifier

A small web server for replace Jenkins update mirror site

## How to use it

1. Modify `replacementText` to your own settings
    I use Yamagata university (Japan) as my default mirror site.

    ```
    https://ftp.yz.yamagata-u.ac.jp/pub/misc/jenkins/
    ```

    Use [this link](https://get.jenkins.io/plugins/git/5.2.0/git.hpi?mirrorlist) and find the proper mirror site for you.

2. Deploy this server to your site

3. Set your Jenkins ignore update center certificate

    ```
    -Dhudson.model.DownloadService.noSignatureCheck=true
    ```

    If you use docker, you can add this to your docker-compose.yml

    ```
    jenkins:
        image: jenkins/jenkins:lts
        container_name: jenkins
        ports:
            - 8080:8080
        volumes:
            - ./jenkins_home:/var/jenkins_home
        environment:
            - JAVA_OPTS=-Dhudson.model.DownloadService.noSignatureCheck=true
    ```

    If you use kubernetes, you can add this to your deployment.yaml

    ```
    spec:
      containers:
      - name: jenkins
        image: jenkins/jenkins:lts
        ports:
        - containerPort: 8080
        env:
        - name: JAVA_OPTS
          value: -Dhudson.model.DownloadService.noSignatureCheck=true
    ```

    If you start Jenkins by command line, you can add this to your command

    ```
    java -Dhudson.model.DownloadService.noSignatureCheck=true -Dhudson.model.UpdateCenter.updateCenterUrl=http://your.server.com/update-center.json -jar jenkins.war --httpPort=8080
    ```

    If you run Jenkins into Tomcat, you can add this to your JAVA_OPTS

    ```
    -Dhudson.model.DownloadService.noSignatureCheck=true -Dhudson.model.UpdateCenter.updateCenterUrl=http://your.server.com/update-center.json
    ```

3. Set your Jenkins update center url to your own server

    ```
    http://your.server.com/update-center.json
    ```

    It can be find by web UI:

    Manage Jenkins -> Manage Plugins -> Advanced -> Update Site

4. Restart Jenkins
