package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GenshinCharacter struct {
	Name          string `json:"name"`
	Weapon        string `json:"weapon"`
	Constellation string `json:"constellation"`
	Region        string `json:"region"`
	Vision        string `json:"vision"`
}

	var characters = []GenshinCharacter{
	{Name: "Diluc", Weapon: "Claymore", Constellation: "Noctua", Region: "Mondstadt", Vision: "Pyro"},
	{Name: "Jean", Weapon: "Sword", Constellation: "Leo Minor", Region: "Mondstadt", Vision: "Anemo"},
	{Name: "Keqing", Weapon: "Sword", Constellation: "Trulla Cementarii", Region: "Liyue", Vision: "Electro"},
	{Name: "Klee", Weapon: "Catalyst", Constellation: "Trifolium", Region: "Mondstadt", Vision: "Pyro"},
	{Name: "Mona", Weapon: "Catalyst", Constellation: "Astrolabos", Region: "Mondstadt", Vision: "Hydro"},
	{Name: "Qiqi", Weapon: "Sword", Constellation: "Pristina Nola", Region: "Liyue", Vision: "Cryo"},
	{Name: "Venti", Weapon: "Bow", Constellation: "Carmen Dei", Region: "Mondstadt", Vision: "Anemo"},
	{Name: "Xiao", Weapon: "Polearm", Constellation: "Alatus Nemeseos", Region: "Liyue", Vision: "Anemo"},
	{Name: "Zhongli", Weapon: "Polearm", Constellation: "Lapis Dei", Region: "Liyue", Vision: "Geo"},
}

func getCharacters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, characters)

}

func addCharacters(c *gin.Context) {
	var newCharacter GenshinCharacter

	if err := c.BindJSON(&newCharacter); err != nil {
		return
	}

	characters = append(characters, newCharacter)
	c.IndentedJSON(http.StatusCreated, newCharacter)

}

func main() {
	router := gin.Default()
	router.GET("/genshin/characters", getCharacters)
	router.POST("/genshin/add-characters", addCharacters)
	router.GET("/genshin/characters/:name", func(c *gin.Context) {
		name := c.Param("name")
		for _, a := range characters {
			if a.Name == name {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "character not found"})
	})
	router.Run()

}
