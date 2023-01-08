FROM mongo:4.4

# Cria o usuário administrador
COPY mongo.sh /docker-entrypoint-initdb.d/