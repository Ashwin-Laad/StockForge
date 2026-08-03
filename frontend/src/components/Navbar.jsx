import { Link } from "react-router-dom";
import { FaShoppingCart } from "react-icons/fa";
import { useContext } from "react";
import { AuthContext } from "../context/AuthContext";




function Navbar() {

  const { token, logout } = useContext(AuthContext);
  return (
    <nav className="bg-slate-950/80 backdrop-blur-md border-b border-slate-800 sticky top-0 z-50">
      <div className="max-w-7xl mx-auto flex justify-between items-center h-16 px-6">

        <Link
          to="/"
          className="text-2xl font-bold tracking-tight text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-emerald-400 hover:from-cyan-300 hover:to-emerald-300 transition-all"
        >
          InventoryPro
        </Link>

        <div className="flex gap-8 items-center">
          {
            token && (
              <>
                <Link to="/products" className="text-slate-300 hover:text-white transition-colors text-sm font-medium">Home</Link>

                <Link to="/profile" className="text-slate-300 hover:text-white transition-colors text-sm font-medium">
                  Profile
                </Link>
                <Link to="/cart" className="text-slate-300 hover:text-white transition-colors text-sm font-medium">
                  Your Picks
                </Link>
                <Link to="/create-product" className="text-slate-300 hover:text-white transition-colors text-sm font-medium">
                  Add Product
                </Link>

               
              </>
            )
          }
        </div>
        {
          token ?
            <button
              onClick={logout}
              className="bg-red-500/10 text-red-400 border border-red-500/30 px-4 py-2 rounded-lg text-sm font-medium hover:bg-red-500/20 hover:text-red-300 transition-colors"
            >
              Logout
            </button>
            :
            (<div className="flex gap-4 items-center">
              <Link
                to="/login"
                className="text-slate-300 hover:text-white transition-colors text-sm font-medium"
              >
                Login
              </Link>
              <Link
                to="/signup"
                className="bg-gradient-to-r from-cyan-500 to-emerald-500 text-slate-950 px-4 py-2 rounded-lg text-sm font-semibold hover:from-cyan-400 hover:to-emerald-400 transition-all shadow-lg shadow-cyan-500/20"
              >
                Sign Up
              </Link>
            </div>)
        }
      </div>
    </nav>
  );
}

export default Navbar;