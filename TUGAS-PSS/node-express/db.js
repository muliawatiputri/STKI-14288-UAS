const mysql = require('mysql2/promise');

const pool = mysql.createPool({
  host: 'localhost',
  user: 'root',
  password: '', // sesuaikan
  database: 'cute_store',
  waitForConnections: true,
  connectionLimit: 10
});

module.exports = pool;
