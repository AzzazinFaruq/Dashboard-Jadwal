package handlers

import (
	"backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
	"strconv"
)

type MataKuliahHandler struct {
	DB *bun.DB
}

func NewMataKuliahHandler(db *bun.DB) *MataKuliahHandler {
	return &MataKuliahHandler{DB: db}
}

func (h *MataKuliahHandler) CreateMataKuliah(c *fiber.Ctx) error {

	type CreateMataKuliahRequest struct {
		KodeMK string `json:"kode_mk"`
		NamaMataKuliah string `json:"nama_mata_kuliah"`
		SKS int `json:"sks"`
	}

	req := new(CreateMataKuliahRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	mataKuliah := &models.MataKuliah{
		KodeMK: req.KodeMK,
		NamaMataKuliah: req.NamaMataKuliah,
		SKS: req.SKS,
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

func (h *MataKuliahHandler) UpdateMataKuliah(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}

	type UpdateMataKuliahRequest struct {
		KodeMK string `json:"kode_mk"`
		NamaMataKuliah string `json:"nama_mata_kuliah"`
		SKS           int    `json:"sks"`
	}

	req := new(UpdateMataKuliahRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	mataKuliah := &models.MataKuliah{
		ID:            id,
		KodeMK: req.KodeMK,
		NamaMataKuliah: req.NamaMataKuliah,
		SKS:           req.SKS,

	}

	res, err := h.DB.NewUpdate().Model(mataKuliah).
		Column("kode_mk", "nama_mata_kuliah", "sks").
		WherePK().
		Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Mata kuliah tidak ditemukan"})
	}

	return c.JSON(fiber.Map{
		"message": "Mata kuliah berhasil diperbarui",
		"data":    mataKuliah,
	})
}

func (h *MataKuliahHandler) DeleteMataKuliah(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
	}
	res, err := h.DB.NewDelete().Model((*models.MataKuliah)(nil)).Where("id = ?", id).Exec(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Mata kuliah tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"message": "Mata kuliah berhasil dihapus"})
}