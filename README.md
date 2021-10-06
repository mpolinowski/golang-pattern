# Go Design Pattern

<!-- TOC -->

- [Creational](#creational)
    - [Builder](#builder)
    - [Factory](#factory)
    - [Singleton](#singleton)
        - [Basic Example](#basic-example)
        - [Thread-safe Example](#thread-safe-example)
- [Structural](#structural)
- [Behavioural](#behavioural)

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

* Adapter
* Facade


## Behavioural

* Iterator
* Observer