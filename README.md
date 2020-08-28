## plural spanish

##### Libreria utilitaria que incluye reglas de trasformación de palabras plurales y singulares en español

### Modo de uso

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
    
    //Validate if word is singular
    fmt.Println(pluralWords.IsPlural("carne"))
	
	value = pluralWords.ToSingular("leches")
	fmt.Println(value)
    
    //Validate if word is plural
    fmt.Println(pluralWords.IsPlural("carne"))

}
````

#### Outputs
````text
carnes
true
leche
true

````