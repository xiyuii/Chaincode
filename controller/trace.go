package controller

import (
	"backend/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
)

// 保存了大模型上链与查询的函数

func Uplink(c *gin.Context) {
    traceability_code := c.PostForm("traceability_code")
    if traceability_code == "" || len(traceability_code) != 18 {
        c.JSON(200, gin.H{
            "message": "请检查溯源码是否正确!!",
        })
        return
    }
    args := buildArgs(c, traceability_code)
    if args == nil {
        return
    }
    res, err := pkg.ChaincodeInvoke("Uplink", args)
    if err != nil {
        c.JSON(200, gin.H{
            "message": "uplink failed" + err.Error(),
        })
        return
    }

    // 增加积分
    userID, _ := c.Get("userID")
    if err := pkg.UpdatePoints(userID.(string), 10); err != nil {  // 假设上传成功增加 10 积分
        c.JSON(200, gin.H{
            "message": "增加积分失败：" + err.Error(),
        })
        return
    }

    c.JSON(200, gin.H{
        "code":              200,
        "message":           "uplink success",
        "txid":              res,
        "traceability_code": args[1],
    })
}


// 获取大模型的上链信息
func GetModelInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetModelInfo", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "查询失败：" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})

}

// 获取用户的大模型ID列表
func GetModelList(c *gin.Context) {
	userID, _ := c.Get("userID")
	res, err := pkg.ChaincodeQuery("GetModelList", userID.(string))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取所有的大模型信息
func GetAllModelInfo(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetAllModelInfo", "")
	fmt.Print("res", res)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

// 获取大模型上链历史
// func (s *SmartContract) GetModelHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
func GetModelHistory(c *gin.Context) {
	res, err := pkg.ChaincodeQuery("GetModelHistory", c.PostForm("traceability_code"))
	if err != nil {
		c.JSON(200, gin.H{
			"message": "query failed" + err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"code":    200,
		"message": "query success",
		"data":    res,
	})
}

func buildArgs(c *gin.Context, traceability_code string) []string {
	var args []string
	userID, _ := c.Get("userID")
	args = append(args, userID.(string))
	// 所有用户都需要提供溯源码
	traceability_code = c.PostForm("traceability_code")
	if traceability_code == "" || len(traceability_code) != 18 {
		c.JSON(200, gin.H{
			"message": "请检查溯源码是否正确!!",
		})
		return nil
	}
	args = append(args, traceability_code)
	// 添加其他参数
	args = append(args, c.PostForm("arg1"))
	args = append(args, c.PostForm("arg2"))
	args = append(args, c.PostForm("arg3"))
	args = append(args, c.PostForm("arg4"))
	args = append(args, c.PostForm("arg5"))
	return args
}
