import { useState } from "react";
import api from "../api/axios";

function Profile() {

  const [address, setAddress] = useState({
    house_name: "",
    house_street: "",
    house_city: "",
    house_pincode: ""
  });


  function handleChange(e) {

    setAddress({
      ...address,
      [e.target.name]: e.target.value
    });

  }


  async function addAddress(e) {

    e.preventDefault();

    try {

      const userID = localStorage.getItem("userID");


      await api.post(
        `/user/address?id=${userID}`,
        {
          ...address,
          house_pincode: Number(address.house_pincode)
        }
      );


      alert("Address Added Successfully");


      setAddress({
        house_name: "",
        house_street: "",
        house_city: "",
        house_pincode: ""
      });


    }
    catch (err) {

      console.log(err);

      alert(
        err.response?.data?.error || "Failed to add address"
      );

    }

  }


  return (

    <div className="min-h-screen bg-gradient-to-b from-[#0a0f1c] via-slate-950 to-[#0a0f1c] relative overflow-hidden px-6 py-10">

      <div className="absolute -top-40 -left-40 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl pointer-events-none" />
      <div className="absolute top-1/3 -right-40 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none" />

      <h1 className="text-4xl font-bold mb-8 text-white relative z-10 max-w-xl mx-auto">
        Profile
      </h1>

      <div className="max-w-xl mx-auto bg-white p-10 rounded-2xl shadow-2xl shadow-black/40 relative z-10">

        <h2 className="text-lg font-bold text-slate-900 mb-1">
          Add Delivery Address
        </h2>
        <p className="text-slate-500 text-sm mb-8">
          This is where we'll send your orders.
        </p>


        <form onSubmit={addAddress} className="space-y-4">


          <div>
            <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
              House / Flat
            </label>
            <input
              name="house_name"
              value={address.house_name}
              onChange={handleChange}
              placeholder="e.g. Flat 4B, Sunrise Apartments"
              className="w-full p-3 rounded-lg border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
            />
          </div>


          <div>
            <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
              Street
            </label>
            <input
              name="house_street"
              value={address.house_street}
              onChange={handleChange}
              placeholder="Street name"
              className="w-full p-3 rounded-lg border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
            />
          </div>


          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
                City
              </label>
              <input
                name="house_city"
                value={address.house_city}
                onChange={handleChange}
                placeholder="City"
                className="w-full p-3 rounded-lg border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              />
            </div>

            <div>
              <label className="block text-xs font-semibold text-slate-500 uppercase tracking-wide mb-1.5">
                Pincode
              </label>
              <input
                name="house_pincode"
                value={address.house_pincode}
                onChange={handleChange}
                placeholder="000000"
                type="number"
                className="w-full p-3 rounded-lg border border-slate-200 bg-slate-50 text-slate-900 placeholder-slate-400 outline-none focus:border-cyan-500 focus:ring-2 focus:ring-cyan-500/20 focus:bg-white transition-all"
              />
            </div>
          </div>


          <button
            type="submit"
            className="w-full bg-gradient-to-r from-cyan-500 to-emerald-500 hover:from-cyan-400 hover:to-emerald-400 text-slate-950 font-semibold py-3 rounded-lg transition-all shadow-lg shadow-cyan-500/30"
          >
            Save Address
          </button>


        </form>

      </div>

    </div>

  );
}

export default Profile;