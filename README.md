Task: Run Existing Docker Image as a Container


1. Run command in termiinal
  docker run -d -p 8080:80 --name welcome1 docker/welcome-to-docker
2. Check whether the welcome1 container can run on the desired port
3. See logs of running welcome1 container in docker desktop
4. Execute in exec command inside running welcome1 container
   cat usr/share/nginx/html/index.html
6. Stop container dan see lates logs
