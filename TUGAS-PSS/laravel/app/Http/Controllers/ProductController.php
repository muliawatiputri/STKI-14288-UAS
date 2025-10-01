<?php
namespace App\Http\Controllers;
use App\Models\Product;
use Illuminate\Http\Request;

class ProductController extends Controller
{
    public function index() { return response()->json(Product::all()); }
    public function show($id) { return response()->json(Product::findOrFail($id)); }
    public function store(Request $r) {
        $p = Product::create($r->only(['name','category','price','stock','description','image_url']));
        return response()->json($p, 201);
    }
    public function update(Request $r, $id) {
        $p = Product::findOrFail($id);
        $p->update($r->only(['name','category','price','stock','description','image_url']));
        return response()->json($p);
    }
    public function destroy($id) {
        Product::findOrFail($id)->delete();
        return response()->json(['message'=>'deleted']);
    }
}
