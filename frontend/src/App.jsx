import { Routes, Route } from "react-router-dom";

import Navbar from "./components/Navbar";

import Home from "./pages/Home";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import Cart from "./pages/Cart";
import Landing from "./pages/Landing";
import Profile from "./pages/Profile"
import CreateProduct from "./pages/CreateProduct";


function App() {
  return (
    <div className="min-h-screen bg-slate-950">

      <Navbar />

      <div className="max-w-7xl mx-auto mt-8 px-4">

        <Routes>

          <Route path="/" element={<Landing />} />
          <Route path="/products" element={<Home />} />
          <Route path="/profile" element={<Profile />} />
          <Route path="/login" element={<Login />} />

          <Route path="/signup" element={<Signup />} />

          <Route path="/cart" element={<Cart />} />

          <Route
            path="/create-product"
            element={<CreateProduct />}
          />

        </Routes>

      </div>
    </div>
  );
}

export default App;