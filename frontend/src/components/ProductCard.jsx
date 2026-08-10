import api from "../api/axios";
import {useDispatch} from "react-redux";
import { loadCart } from "../../redux/cartSlice";


function ProductCard({ product, onAdd , removeBool }) {


  const dispatch=useDispatch();

  
  async function remove(productID) {
    try {
      const id = localStorage.getItem("userID");
      console.log("id is", id)
      console.log("id is", productID)

      await api.delete(
        `/user/remove-product?id=${productID}&user_id=${id}`
      );

      // if (onRemove) {
      //   // onRemove();
      // }
      //instead of defining and calling function here i dispatch it
      dispatch(loadCart());


    } catch (err) {
      console.log(err);
    }
  }

  return (
    <div className="bg-slate-900 border border-slate-800 rounded-xl shadow-lg overflow-hidden hover:shadow-2xl hover:shadow-cyan-500/10 hover:border-slate-700 transition-all duration-300">
      <img
        src={product.image}
        alt={product.product_name}
        className="w-full h-52 object-cover"
      />

      <div className="p-4">
        <h2 className="text-xl font-bold text-white">
          {product.product_name}
        </h2>

        <p className="text-amber-400 mt-2 tracking-wide">
          {"⭐".repeat(product.product_rating)}
        </p>

        <p className="text-emerald-400 font-bold text-lg mt-2">
          ₹ {product.product_price}
        </p>

        {onAdd && (
          <button
            onClick={() => onAdd(product._id)}
            className="mt-4 w-full bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 text-slate-950 font-semibold py-2 rounded-lg transition-all shadow-md shadow-cyan-500/20"
          >
            Add To Cart
          </button>
        )}

        {removeBool && (
          <button
            className="mt-2 w-full bg-red-500/10 border border-red-500/30 text-red-400 hover:bg-red-500/20 hover:text-red-300 py-2 rounded-lg transition-colors"
            onClick={() => remove(product._id)}
          >
            Remove
          </button>
        )}
      </div>
    </div>
  );
}

export default ProductCard;