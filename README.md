## plural spanish

##### Libreria utilitaria que incluye reglas de trasformación de palabras plurales y singulares en español


#### Agregar a  go modules
```
# First line is optional if your project is already defined as a Go module
go mod init <YOUR_PROJECT_NAME>
go get github.com/pzentenoe/plural-spanish v1.0.2
```

#### Modo de uso

````go
package main

import (
    "fmt"
    pluralSpanish "github.com/pzentenoe/plural-spanish"
)

func main() {

	pluralWords := pluralSpanish.NewPluralWords()

	value := pluralWords.ToPlural("carne")
	fmt.Println(value)
    
    //Validate if word is plural
    fmt.Println(pluralWords.IsPlural("carnes"))
	
	value = pluralWords.ToSingular("leches")
	fmt.Println(value)
    
    //Validate if word is singular    
    fmt.Println(pluralWords.IsSingular("carne"))

}
````

#### Outputs
````text
carnes
true
leche
true

````


## Authors
* **[Pablo Zenteno](https://github.com/pzentenoe)** - *Full Stack developer*