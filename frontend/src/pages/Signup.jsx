import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import api from "../api/axios";

function Signup() {

  const navigate = useNavigate();

  const [user, setUser] = useState({
    first_name: "",
    last_name: "",
    email: "",
    phone: "",
    password: "",
  });

  async function handleSignup(e) {

    e.preventDefault();

    try {

      await api.post("/user/signup", user);

      alert("Account Created Successfully");

      navigate("/login");

    } catch (err) {

      alert(err.response?.data?.error || "Signup Failed");

    }

  }

  function handleChange(e) {

    setUser({
      ...user,
      [e.target.name]: e.target.value,
    });

  }

  return (

    <div className="min-h-screen flex items-center justify-center bg-gradient-to-b from-[#0a0f1c] via-slate-950 to-[#0a0f1c] relative overflow-hidden px-6 py-10">

      <div className="absolute -top-40 -left-40 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl pointer-events-none" />
      <div className="absolute top-1/3 -right-40 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none" />

      <div className="bg-white p-10 rounded-2xl shadow-2xl shadow-black/40 w-full max-w-md relative z-10">

        <h1 className="text-3xl font-bold text-center mb-1 text-slate-900">

          Create account

        </h1>
        <p className="text-slate-500 text-sm text-center mb-8">
          Set up your account to get started.
        </p>

        <form
          onSubmit={handleSignup}
          className="space-y-4"
        >

          <div className="grid grid-cols-2 gap-4">
            <input
              name="first_name"
              placeholder="First Name"
              className="w-full border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 rounded-lg p-3 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              onChange={handleChange}
            />

            <input
              name="last_name"
              placeholder="Last Name"
              className="w-full border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 rounded-lg p-3 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              onChange={handleChange}
            />
          </div>

          <input
            type="email"
            name="email"
            placeholder="Email"
            className="w-full border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 rounded-lg p-3 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
            onChange={handleChange}
          />

          <input
            name="phone"
            placeholder="Phone"
            className="w-full border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 rounded-lg p-3 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
            onChange={handleChange}
          />

          <input
            type="password"
            name="password"
            placeholder="Password"
            className="w-full border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 rounded-lg p-3 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
            onChange={handleChange}
          />

          <button
            className="w-full bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 text-slate-950 font-semibold py-3 rounded-lg transition-all shadow-lg shadow-cyan-500/30"
          >
            Sign Up
          </button>

        </form>

        <p className="text-center mt-6 text-slate-500 text-sm">

          Already have an account?

          <Link
            to="/login"
            className="text-cyan-600 hover:text-cyan-700 font-medium ml-2"
          >
            Login
          </Link>

        </p>

      </div>

    </div>

  );

}

export default Signup;