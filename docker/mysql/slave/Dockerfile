FROM mysql:8.0.31

# EXPOSE 3306

COPY ./mysql/config/my.conf /etc/mysql/conf.d/my_slave.cnf
# 起動時に実行してくれるディレクトリ
COPY ./mysql/migrations /docker-entrypoint-initdb.d
