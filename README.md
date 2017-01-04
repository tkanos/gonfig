# gonfig

github.com/tkanos/gonfig is a lightweight package, that give the ability to have all our json config file and env variables in one config object.


For that we have to define our configuration structure :

```golang
type Configuration struct {
	Port              int
	Connection_String string
}
```

and our json config file as well 

```json
{
	"Port": 8080
}
```

As you can see the Connection_String field is not provided on the config file, in order to follow the [Best practices of configuration file](https://medium.com/@tkanos/best-practices-for-configuration-file-in-your-code-2d6add3f4b86#.dze386j1t)
It will is an external configuration, no depend of our code, that will be provided on the env variables.

On Docker :
```bash
$ docker run [...] -e Connection_String="........" [...]
```

So in order to hide this 2 sources of information to our developpers, we will use the magical tkanos/gonfig package :D

```golang
configuration := Configuration{}
err := gonfig.GetConf(getFileName(), &configuration)
```

Now we can use configuration structure as if it was coming from one source.

## When do you have to use it ?

Well to be honest if we have to handle few env variables. It's better to set your configuration structure yourself :

```golang
configuration.Connection_String = os.Getenv("Connection_String")
```

Gonfig will be useful to have a static configuration that allow you to have all the config context everywhere in your app.
And have a generic way to get all env variables define in your struct.

## Sample

Yu can find a sample of the use of Gonfig project [HERE](https://github.com/Tkanos/gonfig-sample)


# Links
- [Best practices of configuration file](https://medium.com/@tkanos/best-practices-for-configuration-file-in-your-code-2d6add3f4b86#.dze386j1t)
- [Sample](https://github.com/Tkanos/gonfig-sample)
