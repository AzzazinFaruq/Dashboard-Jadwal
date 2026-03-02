package handlers

import (
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

type MataKuliahHandler struct {
	DB *bun.DB
}

func NewMataKuliahHandler(db *bun.DB) *MataKuliahHandler {
	return &MataKuliahHandler{DB: db}
}

func (h *MataKuliahHandler) CreateMataKuliah(c *fiber.Ctx) error {

	type CreateMataKuliahRequest struct {
		NamaMataKuliah string `json:"nama_mata_kuliah"`
		SKS int `json:"sks"`
		DosenPengajar string `json:"dosen_pengajar"`
	}

	req := new(CreateMataKuliahRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	mataKuliah := &models.MataKuliah{
		NamaMataKuliah: req.NamaMataKuliah,
		SKS: req.SKS,
		DosenPengajar: req.DosenPengajar,
	}

	_, err := h.DB.NewInsert().Model(mataKuliah).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Mata Kuliah berhasil dibuat",
		"data":    mataKuliah,
	})
}

func (h *MataKuliahHandler) GetMataKuliah(c *fiber.Ctx) error {
	var mataKuliah []models.MataKuliah
	err := h.DB.NewSelect().Model(&mataKuliah).Scan(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(mataKuliah)
}