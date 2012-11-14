hunspellgo
==========

Hunspell bidings for Golang

### Usage
  h, _ := hunspellgo.Hunspell("/home/akh/dev/hunspell/fr.aff", "/home/akh/dev/hunspell/fr.dic")
  fmt.Println(hunspellgo.Spell(h, "Bonjour"))
  fmt.Println(hunspellgo.Spell(h, "Bonj"))
  fmt.Println(hunspellgo.Spell(h, "bébé"))
  hunspellgo.Destroy(h)

  true
  false
  true

Suggest not working yet
