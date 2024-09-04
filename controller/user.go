package controller

import (
	"backend/model"
	"backend/pkg"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	
	"github.com/gin-gonic/gin"
)

// 给用户增加积分
func AddPoints(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "获取用户ID失败",
		})
		return
	}

	pointsStr := c.Query("points") // 或 c.PostForm("points")
	pointsToAdd, err := strconv.Atoi(pointsStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "积分值无效",
		})
		return
	}

	// 更新用户积分
	err = pkg.UpdatePoints(userID.(string), pointsToAdd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "更新积分失败：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "积分增加成功",
	})
}

// 用户注册
func Register(c *gin.Context) {
	// 将用户信息存入mysql数据库
	var user model.MysqlUser
	user.UserID = pkg.GenerateID()
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.RealInfo = pkg.EncryptByMD5(c.PostForm("realInfo"))
	user.Points = 5 // 初始化积分为 5

	err := pkg.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "register failed：" + err.Error(),
		})
		return
	}

	// 将用户信息存入区块链
	// userID string, userType string, realInfoHash string
	// 将post请求的参数封装成一个数组args
	var args []string
	args = append(args, user.UserID)
	args = append(args, c.PostForm("userType"))
	args = append(args, user.RealInfo)

	res, err := pkg.ChaincodeInvoke("RegisterUser", args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "register failed：" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "register success",
		"txid":    res,
	})
}

// 用户登录
func Login(c *gin.Context) {
	var user model.MysqlUser
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")

	// 获取用户ID
	var err error
	user.UserID, err = pkg.GetUserID(user.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "没有找到该用户",
		})
		return
	}

	userType, err := GetUserType(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "login failed:" + err.Error(),
		})
		return
	}

	err = pkg.Login(&user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed:" + err.Error(),
		})
		return
	}

	// 生成jwt
	jwt, err := pkg.GenToken(user.UserID, userType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "login failed:" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "login success",
		"jwt":     jwt,
	})
}

// 用户登出
func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "logout success",
	})
}

// 获取用户类型
func GetUserType(userID string) (string, error) {
	userType, err := pkg.ChaincodeQuery("GetUserType", userID)
	if err != nil {
		return "", err
	}
	return userType, nil
}

// 获取用户信息
func GetInfo(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "get user info failed",
		})
		return
	}

	userType, err := GetUserType(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get user type failed" + err.Error(),
		})
		return
	}

	username, err := pkg.GetUsername(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get user name failed" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"message":  "get user info success",
		"userType": userType,
		"username": username,
	})
}

// 上传模型
func UploadModel(c *gin.Context) {
    userID, exist := c.Get("userID")
    if !exist {
        c.JSON(http.StatusUnauthorized, gin.H{
            "message": "获取用户ID失败",
        })
        return
    }

    traceability_code := c.PostForm("traceability_code")
    if traceability_code == "" || len(traceability_code) != 18 {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "请检查溯源码是否正确!!",
        })
        return
    }

    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "获取文件失败：" + err.Error(),
        })
        return
    }

    // 打开上传的文件
    openedFile, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "无法打开文件：" + err.Error(),
        })
        return
    }
    defer openedFile.Close()

    // 创建保存文件
    filePath := filepath.Join("uploads", file.Filename)
    outFile, err := os.Create(filePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "创建文件失败：" + err.Error(),
        })
        return
    }
    defer outFile.Close()

    // 计算文件哈希值
    hasher := sha256.New()
    if _, err := io.Copy(io.MultiWriter(outFile, hasher), openedFile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "保存文件失败：" + err.Error(),
        })
        return
    }
    fileHash := hex.EncodeToString(hasher.Sum(nil))

    // 构建上链参数，包括溯源码和文件哈希
    args := []string{userID.(string), traceability_code, file.Filename, fileHash}

    // 上链
    res, err := pkg.ChaincodeInvoke("UploadModel", args)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "上链失败：" + err.Error(),
        })
        return
    }

    // 增加积分
    if err := pkg.UpdatePoints(userID.(string), 10); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "增加积分失败：" + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code":    http.StatusOK,
        "message": "模型上传成功",
        "txid":    res,
        "hash":    fileHash,
    })
}


// 下载模型
func DownloadModel(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "获取用户ID失败",
		})
		return
	}

	// 获取模型信息
	modelName := c.PostForm("modelName")

	// 获取当前用户的积分
	currentPoints, err := pkg.GetUserPoints(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "获取积分失败：" + err.Error(),
		})
		return
	}

	if currentPoints < 1 {
		c.JSON(http.StatusPaymentRequired, gin.H{
			"message": "积分不足，无法下载模型",
		})
		return
	}

	// 获取文件路径和哈希值
	filePath := c.Query("file")
	hash := c.Query("hash")

	file, err := os.Open(filepath.Join("uploads", filePath))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "文件未找到：" + err.Error(),
		})
		return
	}
	defer file.Close()

	// 计算下载文件的哈希值
	hasher := sha256.New()
	if _, err := io.Copy(io.MultiWriter(c.Writer, hasher), file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "下载文件时发生错误：" + err.Error(),
		})
		return
	}

	// 校验哈希值
	downloadedHash := hex.EncodeToString(hasher.Sum(nil))
	if downloadedHash != hash {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "文件完整性校验失败",
		})
		return
	}

	// 减少 1 个积分
	err = pkg.UpdatePoints(userID.(string), -1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "扣除积分失败：" + err.Error(),
		})
		return
	}

	// 查询模型信息
	modelData, err := pkg.ChaincodeQuery("DownloadModel", userID.(string), modelName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "下载模型失败：" + err.Error(),
		})
		return
	}

	// 发送文件
    c.File(filePath)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "模型下载成功",
		"data":    modelData,
	})
}

// 查询用户的积分
func GetUserPoints(c *gin.Context) {
	userID := c.PostForm("userID")
	points, err := pkg.GetUserPoints(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user points"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"points": points})
}
