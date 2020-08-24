# HackU_2020_vol.2_sysken

## ローカルリポジトリに持ってくる
1. git clone https://github.com/KK56ken/HackU_2020_vol.2_sysken.git
2. git pull 

## vue起動コマンド
1. docker-compose build 
2. docker-compose up

エラーが出た場合
1. dockerfileの下2行をコメントアウト
2. docker-compose build 
3. docker-compose up -d 
4. docker exec -it hacku_vue /bin/ash
5. cd raise_todo
6. npm install --save
7. exit
8. docker-compose down
9. dockerfileの下2行をコメント外す
10. docker-compose build
11. docker-compose up 
