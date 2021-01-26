# Go-Yaml

This is a repository dedicated to displaying demo code using the `gopkg.in/yaml` Golang package.

# Table of Contents

<details><summary>What Is YAML?</summary>

# Minimum Viable Go-Yaml Product
In my opinion, the most basic thing you can do with the `gopkg.in/yaml` package is read a config file, edit the config, and re-write config file.

After a quick explanation of YAML, you'll find example Go code demonstrating this in the next chapter.

### <ins>What is YAML?</ins>

YAML is a superset of something called JSON (JavaScript Object-Notation).

JSON is a key-value structure that allows you to organize data (values) by an index (key).

```json
//EXAMPLE:
//Key      Value
{"username": "my_username"}
```

When using a programming language to work with an object notation language (JSON, YAML), the notation gets converted into a data structure that is designed to work with key-value oriented data.

For example, when we process YAML data in Python, we extrapolate that data into a Python data structure called a "<a href="https://docs.python.org/3/tutorial/datastructures.html#dictionaries">Dictionary</a>".

```Python
Alex = {"firstname": "Alex"}
print(Alex["firstname"])
# Output == 'Alex'
```

For the Go programming lanauge, it is extrapolated into a data structure called a "<a href="https://blog.golang.org/maps">Map</a>".

```Go
Alex := map[string]string{"firstname": "Alex"}
fmt.Println(Alex["firstname"])
// Output == 'Alex'
```
</details>


<details><summary>What Does Go Need To Read YAML?</summary>

# Requirements To Work With YAML

Golang is a statically typed, and compiled language. 

When working with Go and YAML, the type of data that comes from YAML must be known to Go so the compiler can compile the right data to the right types.

Since the compiler won't allow data that doesn't have a statically-defined type to be compiled, we must explain to Go (via structs) exactly what we will be dealing with when we obtain our data from YAML.

### Structs Must Reflect The Data You Retrieve From YAML

If you are looking for a string and an integer from a YAML file, you must define a `string` and `int` type in Go and scan the respective YAML values to these types accordingly.

Let's take a look at how YAML values can be reflected in Go.

Below is a YAML document which reflects a `programmer`, firstname `Alex`, who is age `31`.

```yaml
programmer:
    firstname:  Alex
    age:        31
```

If we wanted to reflect this YAML structure in Go, we can create a user-defined struct. To handle the recursive nature of YAML, we will embed a struct for every sub-level of data that needs to be expressed.

Notice how we define 2 structs in the example below. One for the `programmer` level, and another struct for the levels beneath. The types for the YAML data fields are first defined using Go data-types (`string`), then further defined to explain how the struct fields will be regarded with respect to the YAML specification. This format of `language_type:"key_defining_string"` works for JSON as well: `json:"firstname"`

```Go
type Programmer struct {
	Programmer struct {
		Firstname   string  `yaml:"firstname"`
		SOA         int     `yaml:"age"`
	}
}
```

While it is possible to define these structs separately, I find that embedding the structs produces a more readable format, and also reduces the amount amount of code required to be written.

See the next chapters for examples of how these structs can be used in conjuction with `go-yaml`
</details>







<details><summary>Reading a YAML Config File In Go</summary>

# Reading a YAML File into Memory

As mentioned in the previous chapter, a struct must be used to explain to Go what data types will be extracted from the YAML.

In the below code, we will read a YAML file called `config.yaml` into memory.

### Define the structure of the YAML File

Here is the "config.yaml" file we will be working with.

```yaml
domain:
  hostname: example.com
  soa: 
    - ns12.domaincontrol.com.
    - dns.jomax.net.  
  ip: 157.230.49.66
  nameservers:	
    - ns11.domaincontrol.com.
    - ns12.domaincontrol.com.
```

Here is the structure written as a Go struct.

```go
type Domain struct {
    Domain struct {
        Hostname    string      `yaml:"hostname"`
        SOA         []string    `yaml:"soa,flow"` // ",flow" signifies an array/list
        IP          string      `yaml:"ip"`
        Nameservers []string    `yaml:"nameservers,flow"` // ",flow" signifies an array/list
    }
}
```
### Read YAML File

Next, we can begin the process of reading the YAML config file into memory.

When we read a file into memory using Go, it is typically read an array of bytes.

Once it is in memory as bytes, it can be written to another file, or it can be converted into characters, or utilized according to the purpose of your program.

Here is a basic example of reading a YAML file into memory as an array of bytes.

```go
	//Variable for the name of the file we wish to read
	configFile := "./config.yaml"

	//Read configFile into a slice/array of bytes ([]byte)
	sliceOfBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}
```

When `ioutil.ReadFile(configFile)` executes successfully, it will return a slice of bytes to the specified variable `sliceOfBytes`.

### Load YAML Bytes into Our Go Struct Using yaml.Unmarshal

Now that the file contents are in memory, we can now load that data into the fields defined in our Go struct.

Data can be unloaded from memory, into a struct, using the `gopkg.in/yaml` function called `Unmarshal`.

The term "Unmarshal" was confusing to me. It helped when I began to associate the term "Unmarshal" with the term "unload," as in: `unload from memory`.

The below Go code shows an example of taking the byte data from memory, and loading it into a `Config` struct.

```go
    //Create an empty Config object
    var yamlConfig Config

    //Unload the yaml 'config.yaml' bytes into our empty Config object
	err = yaml.Unmarshal(sliceOfBytes, &yamlConfig)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
    }
```

### Why Load YAML into Memory & into a Struct?

You may be asking your self why we have an array of usable data, and a struct with the exact same data duplicated into it.

The answer is that it is far easier to work with a struct than it is to manually handle bytes in memory.

The Go struct will give you greater ease of use, better labeling of the data, and help other programmers understand your code.

# Full Program To Read 'config.yaml' into Memory

```Go
// Read file into Memory
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config
type Config struct {
	Domain struct {
		Hostname    string `yaml:"hostname"`
		SOA         []string  `yaml:"soa,flow"`
		IP          string    `yaml:"ip"`
		Nameservers []string  `yaml:"nameservers,flow"`
	}
}

func main() {
	//Parse options
    configFile := "./config.yaml"
    var yamlConfig Config

	//Read configFile into a slice of bytes ([]byte)
	sliceOfBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}

	//Unmarshal bytes into 'yamlConfig' (Type == Config)
	err = yaml.Unmarshal(sliceOfBytes, &yamlConfig)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}

	fmt.Println(yamlConfig)
}

```

For information on how to write YAML to a file, see the next chapter.
</details>

<details><summary>Writing YAML To A File</summary>

# How To Write YAML To A File

You may find yourself in a situation where you need to write YAML configuration to a file.

Let's a look a how we can do this.

### Utilizing a Go Struct for YAML

As mentioned in previous chapters, we use Go structs, based on the YAML we wish to work with, to control YAML input and output.

We will reuse the config and structs from the previous chapter

```go
// YAML structure we wish to write
type Config struct {
	Domain struct {
		Hostname    string   `yaml:"hostname"`
		SOA         []string `yaml:"soa,flow"`
		IP          string   `yaml:"ip"`
		Nameservers []string `yaml:"nameservers,flow"`
	}
}

// Populate the Go struct with data
	yamlConfig.Domain.Hostname = "example.com"
	yamlConfig.Domain.SOA = []string{"ns11.example.com", "ns12.example.com"}
	yamlConfig.Domain.IP = "10.0.0.80"
	yamlConfig.Domain.Nameservers = []string{"ns11.example.com", "ns12.example.com"}
```

### Getting a Slice/Array of Bytes to Write Our YAML

As explained previously, we use an array of bytes to work with YAML in the raw.

Last time we worked with an array of bytes, we moved data from the YAML side of things, into the Go side of things using `yaml.Unmarshal`.

`config.yaml => Memory =[Unmarshal]=> GoStruct`

Now we shall do the opposite. We will move data from the Go side of things, to the YAML side of things.

This will be accomplished using a similar yaml function called `yaml.Marshal`.

In essence, we will be reversing the flow of data as listed below:

`goStruct =[Marshal]=> Memory => GoStruct`

Another way to look at this operation is:

```java
Unmarshal:  "Unloads from memory/array"
Marshal:    "Loads into memory/array"
```

Below, you will see us use the `yaml.Marshall` command to load bytes into memory from Go.

```go
	//Create []bytes from yaml struct 'yamlConfig'
	sliceOfBytes, err := yaml.Marshal(&yamlConfig)
	if err != nil {
		log.Fatalf("We could not Marshal new struct: %s\n", err)
	}
```

### Creating a File and Writing Our YAML to it.

Now that we have the YAML data `marshal`ed into memory, we can use Go native libraries to create a file and write the YAML structure into that file which was created.

```go
	////create new YAML config 'b.yaml'
	f, err := os.Create("./b.yaml")
	if err != nil {
		log.Fatalf("We could not Create new file 'b.yaml': %s\n", err)
	}

	//If anything fails, defer will execute a close before program terminates no matter what
	defer f.Close()

	////write to b.yaml
	f.Write(sliceOfBytes)
```

# Full Program To Write YAML to 'b.yaml'

You may notice that when we write our data to this b.yaml file, it doesn't exactly resemble what we used before.

The reason for this is due to how tricky the YAML spec is.

YAML is very human friendly, but that ease comes at a cost. It is not trivial to write program libraries which can preserve structure and follow stylistic norms.

Because you can write valid YAML several different ways, `go-yaml` defaults to a YAML structure which closely resembles the `map[string]string` data structure that is utilized in Go.

The thing to bear in mind is that the YAML data found in `b.yaml` is valid and compatible with the other format we saw in the previous lesson file `config.yaml`.
</details>