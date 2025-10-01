const express = require('express');
const bodyParser = require('body-parser');
const products = require('./routes/products');

const app = express();
app.use(bodyParser.json());
app.use('/products', products);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => console.log(`Server listening ${PORT}`));
