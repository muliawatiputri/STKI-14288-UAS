const express = require('express');
const router = express.Router();
const db = require('../db');

router.get('/', async (req, res) => {
  try {
    const [rows] = await db.query('SELECT * FROM products');
    res.json(rows);
  } catch (err) { res.status(500).json({error: err.message}); }
});

router.get('/:id', async (req, res) => {
  try {
    const [rows] = await db.query('SELECT * FROM products WHERE id = ?', [req.params.id]);
    if (!rows.length) return res.status(404).json({error: 'Not found'});
    res.json(rows[0]);
  } catch (err) { res.status(500).json({error: err.message}); }
});

router.post('/', async (req, res) => {
  try {
    const {name, category, price, stock, description, image_url} = req.body;
    const [result] = await db.query(
      'INSERT INTO products (name, category, price, stock, description, image_url) VALUES (?, ?, ?, ?, ?, ?)',
      [name, category, price, stock, description, image_url]
    );
    const [newRow] = await db.query('SELECT * FROM products WHERE id = ?', [result.insertId]);
    res.status(201).json(newRow[0]);
  } catch (err) { res.status(500).json({error: err.message}); }
});

router.put('/:id', async (req, res) => {
  try {
    const {name, category, price, stock, description, image_url} = req.body;
    await db.query(
      `UPDATE products SET name=?, category=?, price=?, stock=?, description=?, image_url=? WHERE id=?`,
      [name, category, price, stock, description, image_url, req.params.id]
    );
    const [rows] = await db.query('SELECT * FROM products WHERE id = ?', [req.params.id]);
    if (!rows.length) return res.status(404).json({error: 'Not found'});
    res.json(rows[0]);
  } catch (err) { res.status(500).json({error: err.message}); }
});

router.delete('/:id', async (req, res) => {
  try {
    const [result] = await db.query('DELETE FROM products WHERE id = ?', [req.params.id]);
    if (result.affectedRows === 0) return res.status(404).json({error: 'Not found'});
    res.json({message: 'Deleted'});
  } catch (err) { res.status(500).json({error: err.message}); }
});

module.exports = router;
