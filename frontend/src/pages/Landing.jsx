import { Link } from "react-router-dom";
import {
  FaBoxOpen,
  FaShoppingCart,
  FaChartLine,
} from "react-icons/fa";

function Landing() {
  const token = localStorage.getItem("token");
  return (
    <div className="min-h-screen bg-gradient-to-b from-[#0a0f1c] via-slate-950 to-black text-white relative overflow-hidden">

      {/* Ambient glows — consistent with the rest of the app */}
      <div className="absolute -top-40 -left-40 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl pointer-events-none" />
      <div className="absolute top-1/4 -right-40 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none" />

      {/* Hero Section */}
      <section className="max-w-7xl mx-auto px-6 py-20 grid md:grid-cols-2 gap-10 items-center relative z-10">

        <div>

          <h1 className="text-6xl font-bold leading-tight">
            Eveything You Need, All In 
            <span className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-emerald-400">
              {" "}One Place
            </span>
          </h1>

          <p className="mt-6 text-lg text-slate-400">
            InventoryPro helps you discover, manage products,
            place orders and simplify
            the entire shopping experience.
          </p>

          {!token && <div className="flex gap-5 mt-8">

            <Link
              to="/signup"
              className="bg-gradient-to-r from-cyan-500 to-emerald-500 text-slate-950 font-bold px-8 py-3 rounded-lg hover:scale-105 hover:shadow-lg hover:shadow-cyan-500/30 transition-all"
            >
              Get Started
            </Link>


            <Link
              to="/login"
              className="border border-slate-700 px-8 py-3 rounded-lg hover:bg-slate-800 hover:border-slate-600 transition-colors"
            >
              Login
            </Link>

          </div> }
          {/* <div className="flex gap-5 mt-8">

            <Link
              to="/signup"
              className="bg-gradient-to-r from-cyan-500 to-emerald-500 text-slate-950 font-bold px-8 py-3 rounded-lg hover:scale-105 hover:shadow-lg hover:shadow-cyan-500/30 transition-all"
            >
              Get Started
            </Link>


            <Link
              to="/login"
              className="border border-slate-700 px-8 py-3 rounded-lg hover:bg-slate-800 hover:border-slate-600 transition-colors"
            >
              Login
            </Link>

          </div> */}

        </div>


        {/* Image */}
        <div className="relative">
          <div className="absolute -inset-4 bg-gradient-to-r from-cyan-500/20 to-emerald-500/20 rounded-3xl blur-2xl" />
          <img
            src="https://images.unsplash.com/photo-1749244768351-2726dc23d26c?w=1200&auto=format&fit=crop&q=80"
            alt="Organized warehouse shelving"
            className="relative rounded-2xl shadow-2xl border border-slate-800 w-full h-full object-cover"
          />
        </div>


      </section>


      {/* Features */}
      <section className="max-w-7xl mx-auto px-6 py-16 relative z-10">

        <h2 className="text-4xl font-bold text-center mb-12">
          Everything You Need
        </h2>


        <div className="grid md:grid-cols-3 gap-8">


          <div className="bg-slate-900/60 border border-slate-800 p-8 rounded-xl hover:border-cyan-500/40 hover:-translate-y-1 transition-all">
            <div className="w-12 h-12 rounded-lg bg-cyan-500/10 flex items-center justify-center mb-5">
              <FaBoxOpen className="text-cyan-400 text-xl" />
            </div>
            <h3 className="text-2xl font-bold text-cyan-400">
              Product Management
            </h3>
            <p className="mt-3 text-slate-400">
              Add, update and organize products easily.
            </p>
          </div>


          <div className="bg-slate-900/60 border border-slate-800 p-8 rounded-xl hover:border-emerald-500/40 hover:-translate-y-1 transition-all">
            <div className="w-12 h-12 rounded-lg bg-emerald-500/10 flex items-center justify-center mb-5">
              <FaShoppingCart className="text-emerald-400 text-xl" />
            </div>
            <h3 className="text-2xl font-bold text-emerald-400">
              Smart Cart
            </h3>
            <p className="mt-3 text-slate-400">
              Manage customer carts and complete purchases smoothly.
            </p>
          </div>


          <div className="bg-slate-900/60 border border-slate-800 p-8 rounded-xl hover:border-purple-500/40 hover:-translate-y-1 transition-all">
            <div className="w-12 h-12 rounded-lg bg-purple-500/10 flex items-center justify-center mb-5">
              <FaChartLine className="text-purple-400 text-xl" />
            </div>
            <h3 className="text-2xl font-bold text-purple-400">
              Order Tracking
            </h3>
            <p className="mt-3 text-slate-400">
              Keep track of completed orders and customer activity.
            </p>
          </div>


        </div>

      </section>


      {/* Footer */}
      <footer className="border-t border-slate-800 relative z-10">
        <div className="max-w-7xl mx-auto px-6 py-8 flex flex-col md:flex-row justify-between items-center gap-4">
          <p className="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-emerald-400 font-bold text-lg">
            InventoryPro
          </p>
          <p className="text-slate-500 text-sm">
            © {new Date().getFullYear()} InventoryPro. All rights reserved.
          </p>
        </div>
      </footer>


    </div>
  );
}

export default Landing;