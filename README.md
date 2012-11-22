hunspellgo
==========

Hunspell bindings for Golang

### Install 
    go get github.com/akhenakh/hunspellgo 

### Usage
    import "github.com/akhenakh/hunspellgo" 

    h := hunspellgo.Hunspell("/home/akh/dev/hunspell/fr.aff", "/home/akh/dev/hunspell/fr.dic")
    fmt.Println(h.Spell("Bonjour"))
    fmt.Println(h.Spell("Bonj"))
    fmt.Println(h.Spell("bébé"))
    
    true
    false
    true

Suggest not working yet
