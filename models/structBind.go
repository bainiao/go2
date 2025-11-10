package models

import "github.com/gin-gonic/gin"

type StructA struct {
	FieldA string `form:"field_a" binding:"required"`
}
type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}
type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}
type StructD struct {
	NestedAnnoyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	b := StructB{}
	c.Bind(&b)
	c.JSON(200, gin.H{"a": b.NestedStruct, "b": b.FieldB})
}
func GetDataC(c *gin.Context) {
	b := StructC{}
	c.Bind(&b)
	c.JSON(200, gin.H{"a": b.NestedStructPointer, "c": b.FieldC})
}
func GetDataD(c *gin.Context) {
	b := StructD{}
	c.Bind(&b)
	c.JSON(200, gin.H{"a": b.NestedAnnoyStruct, "d": b.FieldD})
}
