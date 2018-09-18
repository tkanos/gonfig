# gonfig

gonfig is a lightweight Golang package for intergrating both JSON configs and enviornment variables into one config object.

## Usage

First define a configuration structure:

```golang
type Configuration struct {
	Port              int
	Connection_String string
}
```

Then fill in our JSON file:

```json
{
	"Port": 8080
}
```

We do not define `Connection_String` in the JSON as we would prefer to define that through an enviornment variable.

[Best practices of configuration file](https://medium.com/@tkanos/best-practices-for-configuration-file-in-your-code-2d6add3f4b86#.dze386j1t)

using Docker:
```bash
$ docker run [...] -e Connection_String="..." [...]
```

To make this simple for developers we can use gonfig to easily fill in our struct.

```bash
$ go get github.com/tkanos/gonfig
```

```golang
import "github.com/tkanos/gonfig"

configuration := Configuration{}
err := gonfig.GetConf("pathtomyjonfile.json", &configuration)
if err != nil {
	panic(err)
}
```

Now we can use the configuration as if it was coming from one source.

```golang
// pseudo code
if configuration.Port == 8080 {
	return true
}
if configuration.Connection_String != nil {
	return true
}
```

### using different environment variables name

If your env variable has a different name than the json one, you can just define an env attribute

```golang
type Configuration struct {
	Port              int  `env:"MYAPP_PORT"`
	Connection_String string
}
```

## When should gonfig be used?

If you have a limited number of enviornment configuration variables, it's probably better to set the struct values yourself.

```golang
configuration.Connection_String = os.Getenv("Connection_String")
```

gonfig makes it easier to combine JSON and enviornment variables into one struct automatically.

## Sample

You can find a sample of the use of Gonfig project [HERE](https://github.com/Tkanos/gonfig-sample)


# Links
- [Best practices of configuration file](https://medium.com/@tkanos/best-practices-for-configuration-file-in-your-code-2d6add3f4b86#.dze386j1t)
- [Sample](https://github.com/Tkanos/gonfig-sample)
