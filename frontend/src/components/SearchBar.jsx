import { FaSearch } from "react-icons/fa";

function SearchBar({ value, onChange }) {
  return (
    <div className="max-w-xl mx-auto mb-10">
      <div className="flex items-center bg-slate-100 rounded-xl shadow-lg px-4 py-3 ring-1 ring-transparent focus-within:ring-2 focus-within:ring-cyan-500 transition-all">
        <FaSearch className="text-slate-400 mr-3" />

        <input
          type="text"
          placeholder="Search Products..."
          value={value}
          onChange={onChange}
          className="w-full outline-none bg-transparent text-slate-900 placeholder-slate-400"
        />
      </div>
    </div>
  );
}

export default SearchBar;