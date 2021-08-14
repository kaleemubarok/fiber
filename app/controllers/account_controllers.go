package controllers

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kaleemubarok/fiber/app/models"
	"github.com/kaleemubarok/fiber/app/queries"
	"github.com/kaleemubarok/fiber/pkg/utils"
	"github.com/kaleemubarok/fiber/platform/database"
)

func OpenDBConnection() (*queries.AccountQueries, error) {
	db, err := database.PostgresSQLConnetion()
	if err != nil {
		return nil, err
	}

	return &queries.AccountQueries{db}, nil
}

func GetAccounts(c *fiber.Ctx) error {
	db, err := OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	accounts, err := db.GetAccounts()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "account were not found" + err.Error(),
			"count":   0,
			"account": nil,
		})
	}
	for _, accont := range accounts {
		accont.Salt = ""
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"count":    len(accounts),
		"accounts": accounts,
	})
}

// GetAccount func for retrieve account details.
// @Description get account details.
// @Summary get account details
// @Tags Account
// @Accept json
// @Produce json
// @Param id body string true "ID"
// @Success 200 {object} models.AccountSwag
// @Security ApiKeyAuth
// @Router /v1/account/{id} [get]
func GetAccount(c *fiber.Ctx) error {
	fmt.Println("sample  UUID:", uuid.New())

	accountID := c.Get("id")
	fmt.Println("1.", accountID)
	if accountID == "" {
		accountID = c.Params("id")
		fmt.Println("2.", accountID)
	}

	accountUUID, err := uuid.Parse(accountID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	db, err := OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	account, err := db.GetAccount(accountUUID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "account were not found " + err.Error(),
			"count":   0,
			"account": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     "",
		"count":   1,
		"account": account,
	})

}

// CreeateAccount func for cretae an account.
// @Description create an account.
// @Summary create an account
// @Tags Account
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 200 {object} models.AccountSwag "response will be: account created."
// @Router /v1/account [post]
func CreateAccount(c *fiber.Ctx) error {
	type accountRequest struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	sEmail := &accountRequest{}
	if err := c.BodyParser(sEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	validate := utils.NewValidator()
	if err := validate.Struct(sEmail); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err), //todo: update error msg
		})
	}

	uuid := uuid.New()
	email := sEmail.Email
	password := sEmail.Password
	create_date := time.Now()
	salt := RandStringBytes(32)

	password += salt

	hash := sha256.New()
	hash.Write([]byte(password))
	password = fmt.Sprintf("%x", hash.Sum(nil))

	db, err := OpenDBConnection()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	account := models.Account{
		ID:            uuid,
		CreatedAt:     create_date,
		Email:         email,
		Salt:          salt,
		Password:      password,
		AccountStatus: 1,
	}

	if err := db.CreateAccount(account); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     "account created.",
		"account": account,
	})

}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
