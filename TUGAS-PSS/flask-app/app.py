from flask import Flask, request, jsonify
from flask_sqlalchemy import SQLAlchemy
from datetime import datetime

app = Flask(__name__)
app.config['SQLALCHEMY_DATABASE_URI'] = 'mysql+pymysql://root:@localhost/cute_store'
app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
db = SQLAlchemy(app)

class Product(db.Model):
    __tablename__ = 'products'
    id = db.Column(db.Integer, primary_key=True)
    name = db.Column(db.String(255), nullable=False)
    category = db.Column(db.String(100), default='aksesoris')
    price = db.Column(db.Numeric(10,2), nullable=False)
    stock = db.Column(db.Integer, default=0)
    description = db.Column(db.Text)
    image_url = db.Column(db.String(500))
    created_at = db.Column(db.DateTime, default=datetime.utcnow)
    updated_at = db.Column(db.DateTime, default=datetime.utcnow, onupdate=datetime.utcnow)

    def to_dict(self):
        return {c.name: getattr(self, c.name) for c in self.__table__.columns}

@app.route('/products', methods=['GET'])
def list_products():
    return jsonify([p.to_dict() for p in Product.query.all()])

@app.route('/products/<int:id>', methods=['GET'])
def get_product(id):
    p = Product.query.get_or_404(id)
    return jsonify(p.to_dict())

@app.route('/products', methods=['POST'])
def create_product():
    data = request.json
    p = Product(**{k:data[k] for k in ['name','category','price','stock','description','image_url'] if k in data})
    db.session.add(p)
    db.session.commit()
    return jsonify(p.to_dict()), 201

@app.route('/products/<int:id>', methods=['PUT'])
def update_product(id):
    p = Product.query.get_or_404(id)
    data = request.json
    for k in ['name','category','price','stock','description','image_url']:
        if k in data: setattr(p, k, data[k])
    db.session.commit()
    return jsonify(p.to_dict())

@app.route('/products/<int:id>', methods=['DELETE'])
def delete_product(id):
    p = Product.query.get_or_404(id)
    db.session.delete(p)
    db.session.commit()
    return jsonify({'message':'deleted'})

if __name__ == '__main__':
    app.run(debug=True)
