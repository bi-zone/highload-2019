## About

 Задача: написать модуль, распознающий аудиофайлы. Примеры использования см ниже.
 
 Посмотреть примерный алгоритм можно [тут](https://habr.com/ru/post/181654), но в статье есть проблемы, потому делать точно так не стоит

 При решении учитывать, что проиндексированый файл и файл для распознавания могут не совпадать не только побайтово, но и в обычном смысле, на слух: на аудиозаписе может быть шум или даже пара треков наложеных друг на друга, подряд может быть проиндексировано несколько похожих композиций и всякие прочие осложнения.

<<<<<<< HEAD
 Треки для тестов https://gofile.io/?c=0Wq0kw
=======
>>>>>>> upstream/master
 
## Sources Description

 В репозитории лежит четыре небольших модуля: decode, fingerprint, models, musiclibrary и отдельный файл shell.go;
 1. decode: Отвечает за декодинг mp3 файлов с помощью сишной библиотеки mpg123. Код модуля уже написан, менять его необходиомости нет. Библиотека выбрана путем ресерча и тестирования, если есть желание использовать другую, все в ваших руках.
 1. fingerprint: основной предмет вашего интереса, модуль предоставляет функцию Fingerprint которая должна вернуть хеш аудиозаписи. Далее этот хеш должен помещаться в некоторое хранилище вместе с названием трека.
 1. models: обертки над postgres db, вместо постгри можно использовать in-memory хранилище.
 1. musiclibrary: модуль, который будет экспортироваться при тестировании. Содержит в себе api
 1. shell.go: обертка над musiclibrary, в тестировании решений участия не принимает но может быть полезна для дебага.
 
## Как решать
 1. Форкнуть репозиторий на гитхабе
 1. Реализовать недостающие функции (можно переписать все с нуля, главное, чтобы у пакета musicLibrary было ожидаемое апи)
 1. Открыть pull request
 1. Задать возникшие вопросы в конференции BI.ZONE в телеграме 
 1. PROFIT
 
 В середине каждого дня конференции все пул реквесты будут клонированы и протестированы на стендах. Ожидаемое апи модуля описано ниже.

## Dependencies
Для декодирования mp3 рекомендуется использовать библиотеку mpg123 (враперы над сишным кодом в decode) либо внимательно отнестись к подбору альтернативы

#### macOS (Homebrew)
```
$ brew install mpg123
```
#### Debian
```
$ sudo apt-get install libmpg123-dev
```

## Setup
Перед запуском тестов будет создана таблица в postgres следующими командами
```
$ createdb -O user databasename
$ psql -f createdb.sql databasename
```

Использование базы данных не является необходимым для решения задачи, доступ к ней предоставляется для удобства.

## Usage
#### Shell mode
```
$ DBUSER=username DBNAME=databasename go run shell.go
Initializing library...

MusicLibrary interactive shell
>>> help

Commands:
  clear             clear the screen
  delete            delete audio from database
  exit              exit the program
  help              display help
  index             index audiofile
  indexdir          index directory
  recognize         recognize audiofile
  recognizedir      recognize directory


>>> index resources/Modjo\ -\ Lady\ \(Hear\ Me\ Tonight\).mp3
Indexing 'resources/Modjo - Lady (Hear Me Tonight).mp3'...
>>> recognize samples/modjo_lady_sample.mp3
Recognizing 'samples/modjo_lady_sample.mp3'...
Best match: Modjo - Lady (Hear Me Tonight)
```

#### API
```golang
package main

import (
	"fmt"
	"proj/musiclibrary"
	_ "github.com/lib/pq"
)

func main() {
	cfg := &musiclibrary.Config{
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		DBname:   os.Getenv("DBNAME"),
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
	}

	musicLib, _ := musiclibrary.Open(cfg)
	defer musicLib.Close()

	musicLib.Index("resources/Modjo - Lady (Hear Me Tonight).mp3")
	result := musicLib.Recognize("samples/modjo_lady_sample.mp3")
	fmt.Println(result)
}
```

## Особые благодарности
 Особая благодарность Московскому хип-хоп-постпанк-ньювейв-постпродиджи андерграунду в лице группы Китай-Брусника за предоставленные треки с аутентичным сведением.
 https://vk.com/kitaybrusnika
