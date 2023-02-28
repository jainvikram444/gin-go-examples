package form

import "github.com/gin-gonic/gin"

type StrucA struct {
	FieldA string `form:"field_a"`
}

type StrucB struct {
	NestedStruct StrucA
	FieldB       string `form:"field_b"`
}

type StrucC struct {
	NestedStructPointer *StrucA
	FieldC              string `form:"field_c"`
}

type StrucD struct {
	NestedAnnonuStruc struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StrucB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StrucC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StrucD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedAnnonuStruc,
		"c": b.FieldD,
	})
}
