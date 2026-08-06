import { useEffect, useState } from "react";
import api from "../api/axios";
import ProductCard from "../components/ProductCard";
import { useNavigate } from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import { loadCart } from "../../redux/cartSlice";

function Cart() {

  const navigate = useNavigate();
const dispatch=useDispatch();
  // const [cart, setCart] = useState([]);
 const cart= useSelector(state=>state.cart.items)
  // const [total, setTotal] = useState(0);
  const total=useSelector(state=>state.cart.total)

  useEffect(() => {
  dispatch(loadCart())

  }, []);

  async function checkout() {

    const id = localStorage.getItem("userID");

    await api.post(

      `/user/buy?id=${id}`

    );

    alert("Order Placed");

    navigate("/");

  }

  // async function loadCart() {

  //   const id = localStorage.getItem("userID");

  //   const res = await api.get(`/get-cart?id=${id}`);

  //   setCart(res.data.cart);

  //   setTotal(res.data.total);

  // }

  return (

    <div className="min-h-screen bg-slate-950 px-6 py-10">

      <h1 className="text-4xl font-bold mb-8 text-white">

        My Cart

      </h1>

      <div className="grid md:grid-cols-3 gap-8">
        {cart.map((product) => (
          <ProductCard
            key={product._id}
            product={product}
            removeBool={true}
          />
        ))}
      </div>
      <h2 className="text-2xl font-bold mt-8 text-slate-200">

        Total : <span className="text-emerald-400">₹ {total}</span>

      </h2>

      <button

        className="bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 text-slate-950 font-semibold px-8 py-3 rounded-lg mt-6 transition-all shadow-lg shadow-cyan-500/20"

        onClick={checkout}

      >

        Checkout

      </button>

    </div>

  );

}

export default Cart;