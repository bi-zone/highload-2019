## About
 Посмотреть примерный алгоритм можно [тут](https://habr.com/ru/post/181654), но в статье есть проблемы, потому делать точно так не стоит

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
Best match: Modjo - Lady (Hear Me Tonight) (11% matched)
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
