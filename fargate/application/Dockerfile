FROM dockercloud/hello-world

RUN sed -i -e "s/world/Fargate/g" /www/index.php
RUN sed -i -e "s/hostname/project ID/g" /www/index.php
RUN sed -i -e "s/HOSTNAME/PROJECT_ID/g" /www/index.php
