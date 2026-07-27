import { useState } from "react";
import api from "../api/axios";
import { useNavigate } from "react-router-dom";
function CreateProduct() {


  const navigate = useNavigate();


  const [product, setProduct] = useState({
    product_name: "",
    product_price: "",
    product_rating: "",
    image: ""
  });


  function handleChange(e) {

    setProduct({
      ...product,
      [e.target.name]: e.target.value
    });

  }


  async function addProduct(e) {

    e.preventDefault();

    try {

      const res = await api.post(
        "/products",
        {
          ...product,
          product_price: Number(product.product_price),
          product_rating: Number(product.product_rating)
        }
      );

      alert("Product Added Successfully");



      console.log(res.data);


      setProduct({
        product_name: "",
        product_price: "",
        product_rating: "",
        image: ""
      });


    }
    catch (err) {

      console.log(err);
      alert("Failed to add product");

    }
    navigate("/")
  }


  return (

    <div className="min-h-screen bg-gradient-to-b from-[#0a0f1c] via-slate-950 to-[#0a0f1c] relative overflow-hidden px-6 py-10">

      <div className="absolute -top-40 -left-40 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl pointer-events-none" />
      <div className="absolute top-1/3 -right-40 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none" />

      <div className="max-w-xl mx-auto mt-10 bg-white rounded-2xl shadow-2xl shadow-black/40 p-10 relative z-10">

        <h1 className="text-3xl font-bold mb-1 text-slate-900">
          Add Product
        </h1>
        <p className="text-slate-500 text-sm mb-8">
          Fill in the details below to list a new product.
        </p>


        <form onSubmit={addProduct} className="space-y-5">

          <div>
            <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
              Product Name
            </label>
            <input
              name="product_name"
              value={product.product_name}
              onChange={handleChange}
              placeholder="e.g. Wireless Mouse"
              className="border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 p-3 w-full rounded-lg outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
            />
          </div>

          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
                Price (₹)
              </label>
              <input
                name="product_price"
                value={product.product_price}
                onChange={handleChange}
                placeholder="0.00"
                type="number"
                className="border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 p-3 w-full rounded-lg outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              />
            </div>

            <div>
              <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
                Rating
              </label>
              <input
                name="product_rating"
                value={product.product_rating}
                onChange={handleChange}
                placeholder="1-5"
                type="number"
                className="border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 p-3 w-full rounded-lg outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              />
            </div>
          </div>

          <div>
            <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
              Image URL
            </label>
            <input
              name="image"
              value={product.image}
              onChange={handleChange}
              placeholder="https://..."
              className="border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 p-3 w-full rounded-lg outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
            />
          </div>


          <button
            type="submit"
            className="w-full bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 text-slate-950 font-semibold px-8 py-3 rounded-lg transition-all shadow-lg shadow-cyan-500/30"
          >
            Add Product
          </button>


        </form>

      </div>

    </div>

  );

}

export default CreateProduct;