# Go Design Pattern

<!-- TOC -->

- [Creational](#creational)
    - [Builder](#builder)
    - [Factory](#factory)
    - [Singleton](#singleton)
        - [Basic Example](#basic-example)
        - [Thread-safe Example](#thread-safe-example)
- [Structural](#structural)
    - [Adapter](#adapter)
    - [Facade](#facade)
- [Behavioural](#behavioural)
    - [Observer](#observer)
    - [Iterator](#iterator)
        - [Callback Iteration](#callback-iteration)
        - [Pull Model Iteration](#pull-model-iteration)

<!-- /TOC -->

## Creational

### Builder

> The __Builder Pattern__ isolates an object's construction to increase your codes readability.

The following program creates a notification function in `.\Creational\Builder\builder.go` that can be called, and provided with the necessary variables, from your main function in `.\Creational\Builder\main.go`:

Bash:

```bash
go run ./Creational/Builder/*.go
Notification: &{title:New Notification subtitle: message:This is a level 3 alert! image:image.jpg icon:icon.png priority:3 notType:alert}
```

Powershell:

```cmd
go run .\Creational\Builder\builder.go .\Creational\Builder\notification.go .\Creational\Builder\main.go
Notification: &{title:New Notification subtitle: message:This is a level 3 alert! image:image.jpg icon:icon.png priority:3 notType:alert}
```

Remove the image or set the priority outside of the range between `1`-`5` to trigger the error checks:

```go
bldr.SetImage("")
bldr.SetPriority(0)
```


```bash
$ go run ./*.go
Error: You need to provide an image when using an icon!
$ go run ./*.go
Error: Priority has to be between 1 and 5!
```


### Factory

> The __Factory Pattern__ allows us to create an object when it's types are not yet determined.

The program allocates the creation of the object to the factory function. The factory then decides, based on the specified type, what creation function to use

```bash
go run ./Creational/Factory/*.go

Magazine name: Programmer.Inc
Type: *main.magazine
Name: Programmer.Inc
Pages: 47
Publisher: WTF! kubectl
----------------------

Magazine name: Rust Formation
Type: *main.magazine
Name: Rust Formation
Pages: 46
Publisher: New & Shiny
----------------------

Newspaper name: Vanuatu Daily
Type: *main.newspaper
Name: Vanuatu Daily
Pages: 44
Publisher: Island Life
----------------------

Newspaper name: clickBait
Type: *main.newspaper
Name: clickBait
Pages: 42
Publisher: The Sun
----------------------
```


### Singleton

> The __Singleton Pattern__ allows only one instance of a class and provides global access to it.

#### Basic Example

```bash
 go run ./Creational/Singelton/*.go

Creating logger instance
Providing logger instance
1 : Log level 1 message
Providing logger instance
2 : Log level 2 message
Providing logger instance
3 : Log level 3 message
```

Only one instance is (lazy) created on the first incoming call. The instance is then re-used for the 2 following calls.

#### Thread-safe Example

When the singleton is called on multiple threads we might end up with several instances - e.g.:


```go
for i := 1; i < 10; i++ {
    getLoggerInstance()
}
fmt.Scanln()
```

To prevent this create:

```go
var once sync.Once
```

And wrap your creation with it:


```go
once.Do(func ()  {
    // Lazy create logger on first call
    fmt.Println("Creating logger instance")
    logger = &AppLogger{}	
})
```


```bash
go run ./Creational/Singelton_Thread_Safe/*.go

Creating logger instance
Providing logger instance
Providing logger instance
Providing logger instance
Providing logger instance
Providing logger instance
Providing logger instance
Providing logger instance
Providing logger instance
Providing logger instance
```

And this makes sure that the creation step only runs once, even if there are concurrent requests coming in.


## Structural

### Adapter

> The __Adapter Pattern__ allows you to use an existing interface of an API as another interface without modifying the original API.

Example use case: adapting legacy code to a new API - e.g. integrating an old media center software into a new environment. The adapter allows us to send the same commands to each device to get the same, desired result:


```bash
go run ./Structural/Adapter/*.go

XBMC is On
XBMC Volume: 34
XBMC Volume: 35
XBMC Channel: 98
XBMC Channel: 97
XBMC Channel: 98
XBMC Channel: 1
XBMC is Off
-----------------------
OSMC is On
OSMC Volume is: 11
Setting OSMC Volume to: 10
OSMC Volume is: 10
Setting OSMC Volume to: 11
OSMC channel is: 21
Setting OSMC Channel to: 20
OSMC channel is: 20
Setting OSMC Channel to: 19
OSMC channel is: 19
Setting OSMC Channel to: 20
Setting OSMC Channel to: 21
OSMC is Off
```

### Facade

> The __Facade Pattern__ adds a simpler access point to your system hiding a more complex, low-level API.


```bash
go run ./Structural/Facade/*.go

Preparing an Americano
----------------------
Order for coffee with 146.875 grams of beans and grind level 6 .
Grinding 146.875 grams of beans at level 6 .
Adding 235 mL of hot water.
Not adding milk foam.
Order processed.
Americano is ready

Preparing a Latte
----------------------
Order for coffee with 131.25 grams of beans and grind level 2 .
Grinding 131.25 grams of beans at level 2 .
Adding 350 mL of hot water.
Adding 87.5 mL of milk.
Adding milk foam.
Order processed.
Latte is ready
```


## Behavioural

### Observer

> The __Observer Pattern__ introduces a Pub/Sub architecture where an object (__Subject__) maintains a list of other objects (__Observers__) that it needs to notify about changes. 


```bash
go run ./Behavioural/Observer/*.go

Brother Number One was registered
Brother Number Two was registered
Brother Number One received update: Update 1
Brother Number Two received update: Update 1
Brother Number One received update: Update 2
Brother Number Two received update: Update 2
Brother Number One received update: Update 3
Brother Number Two received update: Update 3
Brother Number One was deregistered
Brother Number Two received update: Update 4
```


### Iterator

> The __Iterator Pattern__ iterates over the elements inside an object without exposing the underlying implementation of the container


#### Callback Iteration

The library data is pushed to the callBack and then processed:

```bash
go run ./Behavioral/Iterator/Callback/*.go

Book title: Leviathan Wakes
Book title: Caliban's War
Book title: Abaddon's Gate
Book title: Cibola Burn
Book title: Nemesis Games
Book title: Babylon's Ashes
Book title: Persepolis Rising
Book title: Tiamat's Wrath
Book title: Leviathan Falls
Published: 2011
Published: 2012
Published: 2013
Published: 2014
Published: 2015
Published: 2016
Published: 2017
Published: 2019
Published: 2021
```

#### Pull Model Iteration

__WIP__