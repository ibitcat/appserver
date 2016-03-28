package apiservice

import (
	"fmt"
	"strconv"

	"app-server/logic"

	"github.com/gin-gonic/gin"
)

func AuditRedpacket(c *gin.Context) error {
	c.Request.ParseForm()
	userId := c.MustGet("userId").(string)
	user, err := logic.GetUserData(userId)
	if err != nil {
		return err
	}
	if !user.IsGm {
		return fmt.Errorf("not gm")
	}

	redpktId := c.Request.FormValue("redpktId")

	status, err := strconv.Atoi(c.Request.FormValue("status"))
	if err != nil {
		return err
	}

	logic.UpdateVerify(redpktId, status)
	return nil
}

func ManualCharge(c *gin.Context) error {
	c.Request.ParseForm()

	userId := c.MustGet("userId").(string)
	user, err := logic.GetUserData(userId)
	if err != nil {
		return err
	}
	if !user.IsGm {
		return fmt.Errorf("not gm")
	}

	targetPhone := c.Request.FormValue("target_phone")
	targetUserId := logic.FindUserIdByAccount(targetPhone)
	if targetUserId == "" {
		fmt.Println("invalid phone", targetPhone)
		return fmt.Errorf("invalid phone %s", targetPhone)
	}

	amount, err := strconv.ParseInt(c.Request.FormValue("amount"), 10, 64)
	if err != nil {
		return err
	}

	targetUser, err := logic.GetUserData(targetUserId)
	if err != nil {
		return err
	}

	nickname := c.Request.FormValue("target_nickname")
	if targetUser.NickName != nickname {
		fmt.Println("invalid nick name", nickname)
		return fmt.Errorf("invalid nick name %s", nickname)
	}
	account := c.Request.FormValue("target_account")
	if targetUser.Account != account {
		fmt.Println("invalid account", account)
		return fmt.Errorf("invalid account %s", account)
	}

	logic.UpdateMoney(targetUserId, amount)
	return nil
}
