package handlers

import (
	"backend/models"
	"backend/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type UserHandler struct {
	DB *bun.DB
}

func NewUserHandler(db *bun.DB) *UserHandler {
	return &UserHandler{DB: db}
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Role     string `json:"role"`
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

	req := new(CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	hashedPwd, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal enkripsi password"})
	}

	user := &models.User{
		Username: req.Username,
		Password: hashedPwd,
		Fullname: req.Fullname,
		Role:     req.Role,
	}

	_, err = h.DB.NewInsert().Model(user).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User berhasil dibuat",
		"data":    user,
	})
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	var users []models.User

	err := h.DB.NewSelect().
		Model(&users).
		Order("id ASC").
		Scan(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := new(models.User)
	
	err := h.DB.NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(c.Context())

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	return c.JSON(user)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.ParseInt(c.Params("id"), 10, 64)

	type UpdateUserRequest struct {
		Username string `json:"username"`
		Fullname string `json:"fullname"`
	}

	req := new(UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input salah"})
	}

	user := &models.User{
		ID:       id,
		Fullname: req.Fullname,
		Username: req.Username,
		UpdatedAt: time.Now(),
	}

	q := h.DB.NewUpdate().
		Model(user).
		Column("fullname", "username", "updated_at").
		Where("id = ?", id)

	result, err := q.Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan atau akses ditolak"})
	}

	return c.JSON(fiber.Map{"message": "User berhasil diupdate"})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := h.DB.NewDelete().
		Model((*models.User)(nil)).
		Where("id = ?", id).
		Exec(c.Context())

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "User tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"message": "User berhasil dihapus"})
}