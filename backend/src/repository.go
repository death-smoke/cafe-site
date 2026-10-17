package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

// Repository - весь SQL приложения.
type Repository struct{ db *sql.DB }

func NewRepository(db *sql.DB) *Repository { return &Repository{db: db} }

func (r *Repository) Ping() error { return r.db.Ping() }

func (r *Repository) ListCategories() ([]Category, error) {
	rows, err := r.db.Query(`SELECT id, name, COALESCE(description,'') FROM categories ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := []Category{}
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.ID, &c.Name, &c.Description); err != nil {
			return nil, err
		}
		out = append(out, c)
	}
	return out, rows.Err()
}

const selectMenu = `SELECT m.id, m.name, m.category_id, c.name,
	COALESCE(m.description,''), COALESCE(m.ingredients,''), m.image_id,
	COALESCE(i.path,''), COALESCE(i.attributes::text,'{}'), m.price
	FROM menu m
	JOIN categories c ON c.id = m.category_id
	LEFT JOIN images i ON i.id = m.image_id`

type scanner interface{ Scan(dest ...any) error }

func scanItem(s scanner) (MenuItem, error) {
	var it MenuItem
	var path, attrs string
	err := s.Scan(&it.ID, &it.Name, &it.CategoryID, &it.CategoryName, &it.Description,
		&it.Ingredients, &it.ImageID, &path, &attrs, &it.Price)
	if err != nil {
		return it, err
	}
	if path != "" {
		it.ImageURL = "/api/" + path
	}
	it.ImageAttributes = json.RawMessage(attrs)
	return it, nil
}

func (r *Repository) GetMenuItem(id int) (MenuItem, error) {
	return scanItem(r.db.QueryRow(selectMenu+" WHERE m.id = $1", id))
}

// ListMenu возвращает страницу блюд и общее число записей под фильтр.
func (r *Repository) ListMenu(f MenuFilter) ([]MenuItem, int, error) {
	where := []string{"1=1"}
	args := []any{}

	if f.CategoryID > 0 {
		args = append(args, f.CategoryID)
		where = append(where, fmt.Sprintf("m.category_id = $%d", len(args)))
	}
	if f.Search != "" {
		args = append(args, "%"+f.Search+"%")
		n := len(args)
		where = append(where, fmt.Sprintf(
			"(m.name ILIKE $%d OR m.description ILIKE $%d OR m.ingredients ILIKE $%d)", n, n, n))
	}
	whereSQL := strings.Join(where, " AND ")

	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM menu m JOIN categories c ON c.id = m.category_id WHERE "+whereSQL, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	args = append(args, f.Limit, (f.Page-1)*f.Limit)
	q := fmt.Sprintf("%s WHERE %s ORDER BY %s %s, m.id ASC LIMIT $%d OFFSET $%d",
		selectMenu, whereSQL, f.SortCol, f.Dir, len(args)-1, len(args))
	rows, err := r.db.Query(q, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	items := []MenuItem{}
	for rows.Next() {
		it, err := scanItem(rows)
		if err != nil {
			return nil, 0, err
		}
		items = append(items, it)
	}
	return items, total, rows.Err()
}

func (r *Repository) CreateMenuItem(in MenuInput) (int, error) {
	var id int
	err := r.db.QueryRow(`INSERT INTO menu (name, category_id, description, ingredients, image_id, price)
		VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`,
		in.Name, in.CategoryID, in.Description, in.Ingredients, in.ImageID, in.Price).Scan(&id)
	return id, err
}

// MenuImageID возвращает image_id блюда (sql.ErrNoRows, если блюда нет).
func (r *Repository) MenuImageID(id int) (*int, error) {
	var img *int
	err := r.db.QueryRow(`SELECT image_id FROM menu WHERE id = $1`, id).Scan(&img)
	return img, err
}

func (r *Repository) UpdateMenuItem(id int, in MenuInput) error {
	_, err := r.db.Exec(`UPDATE menu SET name=$1, category_id=$2, description=$3, ingredients=$4, image_id=$5, price=$6
		WHERE id=$7`, in.Name, in.CategoryID, in.Description, in.Ingredients, in.ImageID, in.Price, id)
	return err
}

// DeleteMenuItem удаляет блюдо и возвращает его image_id (sql.ErrNoRows, если блюда нет).
func (r *Repository) DeleteMenuItem(id int) (*int, error) {
	var img *int
	err := r.db.QueryRow(`DELETE FROM menu WHERE id = $1 RETURNING image_id`, id).Scan(&img)
	return img, err
}

// ---------- images ----------

func (r *Repository) InsertImage(path string, attrs []byte) (int, error) {
	var id int
	err := r.db.QueryRow(`INSERT INTO images (path, attributes) VALUES ($1, $2::jsonb) RETURNING id`,
		path, string(attrs)).Scan(&id)
	return id, err
}

// DeleteImageIfUnused удаляет запись картинки, если на неё никто не ссылается; возвращает её path.
func (r *Repository) DeleteImageIfUnused(id int) (string, error) {
	var path string
	err := r.db.QueryRow(`DELETE FROM images i WHERE i.id = $1
		AND NOT EXISTS (SELECT 1 FROM menu m WHERE m.image_id = i.id) RETURNING i.path`, id).Scan(&path)
	return path, err
}
