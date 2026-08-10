import { useState, useContext } from "react";
import { useNavigate } from "react-router-dom";

import { useDispatch } from "react-redux" //helps connect react with redux


import api from "../api/axios";
// import { AuthContext } from "../context/AuthContext";

import { login } from "../../redux/authSlice"

function Login() {

  const navigate = useNavigate();

  const dispatch = useDispatch();

  // const { login } = useContext(AuthContext);

  const [email, setEmail] = useState("");

  const [password, setPassword] = useState("");

  async function handleLogin(e) {

    e.preventDefault();

    try {

      const res = await api.post("/user/login", {
        email,
        password,
      });

      console.log(res.data);
      // login(res.data.token, res.data.user_id);
      //here all values in login go in action as payload
      dispatch(login({
        jwt: res.data.token,
        userId: res.data.user_id
      }))
      alert("Login Successful");

      navigate("/");

    } catch (err) {

      alert(err.response?.data?.error || "Login Failed");

    }

  }

  return (

    <div className="min-h-screen bg-gradient-to-b from-[#0a0f1c] via-slate-950 to-[#0a0f1c] relative overflow-hidden flex items-center justify-center px-6 py-10">

      <div className="absolute -top-40 -left-40 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl pointer-events-none" />
      <div className="absolute top-1/3 -right-40 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none" />

      <div className="w-full max-w-md bg-white shadow-2xl shadow-black/40 rounded-2xl p-10 relative z-10">

        <h1 className="text-3xl font-bold mb-1 text-center text-slate-900">
          Welcome back
        </h1>
        <p className="text-slate-500 text-sm text-center mb-8">
          Log in to manage your inventory.
        </p>

        <form
          onSubmit={handleLogin}
          className="space-y-5"
        >

          <div>
            <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
              Email
            </label>
            <input
              type="email"
              placeholder="you@example.com"
              className="border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 rounded-lg w-full p-3 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <div>
            <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
              Password
            </label>
            <input
              type="password"
              placeholder="••••••••"
              className="border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 rounded-lg w-full p-3 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          <button
            className="bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 text-slate-950 font-semibold w-full p-3 rounded-lg transition-all shadow-lg shadow-cyan-500/30"
          >
            Login
          </button>

        </form>

      </div>

    </div>

  );

}

export default Login;