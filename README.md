# ZCOM

rodar mongodb

instalar dependencias mongod e mongosh
caso acontença algum problema 

utilize o comando na raiz: 
sudo rm -rf /tmp/mongodb-27017.sock

fluxo natural:
service mongod start
mongosh


# Inicializando Todo Fluxo

Para inicializar as instancias com o kafka local e o mongodb:
    Na pasta /Zcom/docker rode no terminal o seguinte comando:
    docker-compose up
    (Tenha o docker baixado corretamente em sua maquina)

Na pasta raiz /Zcom
no terminal rode o seguinte comando
    go run main.go

Na pasta ZCOM/consumer, onde iniciará o Consumidor das mensagens kafka
no terminal rode o seguinte comando:
    go run main.go
