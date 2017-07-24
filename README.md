# Данные о населении

Передача данных из csv файлов по TCP соединению

## Задача

Используется [DataSet](https://github.com/datasets/population-city)

Он же лежит [в каталоге](data)

Читаем данные и отправляем с определенной периодичностью на порт,
который будет прослушиваться **Spark** приложением

### Запуск приложения

Для написания **Go** кода использовался стандартная структура проектов, 
предлагаемая в официальной документации:

```
bin/
    hello                          # command executable
    outyet                         # command executable
    pupulation-data
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # package object
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go                   
    population-data/
        polulation-data.go
    ... (many more repositories and packages omitted) ...
```

Посему, все эти директории должны быть добавлены в переменную
**$GOPATH**, а для просто запуска бинарников в PATH добавляем **$GOPATH/bin** 

тогда, в директории **$GOPATH/src/population-data**: 
```bash
$ go install
```

Запуск для чтения файлов обоих полов:
```bash
$ population-data -data both 
```
Для чтения файла с записями о разных полах раздельно:
```bash
$ population-data -data diff
```

### Внимание!
При таком запуске в папку bin/population-data положить папку data репозитория!
ну или не выделываться и запускать так из папки, склонированной с гита: 
```bash
$ go run population-data
```