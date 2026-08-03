import { useEffect, useState, useContext } from "react";
import api from "../api/axios";
import ProductCard from "../components/ProductCard";
import SearchBar from "../components/SearchBar";

function Home() {

    const [products, setProducts] = useState([]);
    const [search, setSearch] = useState("");
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        if (search.trim() === "") {
            getProducts();
            
        } else {
            searchProducts();
        }
    }, [search]);

    async function getProducts() {
        setLoading(true);

        try {
            const res = await api.get("/products");
            if(!res.data) {
                setProducts([])
            }
            else{  setProducts(res.data);

            }
          
            console.log(res);
        } catch (err) {
            console.log(err);
        } finally {
            setLoading(false);
        }
    }

    async function searchProducts() {
        setLoading(true);

        try {
            const res = await api.get(
                `/products/search?name=${search}`
            );

            setProducts(Array.isArray(res.data) ? res.data : []);
        } catch (err) {
            console.log(err);
        } finally {
            setLoading(false);
        }
    }

    async function addToCart(productID) {

        try {

            const userID = localStorage.getItem("userID");

            await api.post(
                `/user/add-product?id=${userID}&iD=${productID}`
            );

            alert("Added To Cart");

            setCartCount(prev => prev + 1);

        }
        catch (err) {

            alert(err.response?.data?.error);

        }
    }

    if (loading) {
        return (
            <div className="min-h-screen bg-gradient-to-b from-[#0a0f1c] via-slate-950 to-[#0a0f1c] flex items-center justify-center">
                <div className="flex flex-col items-center gap-4">
                    <div className="w-10 h-10 border-2 border-slate-700 border-t-cyan-400 rounded-full animate-spin" />
                    <h1 className="text-slate-400 text-lg font-medium">
                        Loading...
                    </h1>
                </div>
            </div>
        );
    }

    return (
        <div className="min-h-screen bg-gradient-to-b from-[#0a0f1c] via-slate-950 to-[#0a0f1c] relative overflow-hidden px-6 py-10">

            <div className="absolute -top-40 -left-40 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl pointer-events-none" />
            <div className="absolute top-1/3 -right-40 w-96 h-96 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none" />

            <div className="relative z-10">

                <div className="text-center py-12">
                    <h1 className="text-5xl font-bold text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-emerald-400">
                        Inventory Management
                    </h1>

                    <p className="mt-4 text-slate-400">
                        Browse products and manage your cart.
                    </p>
                </div>

                <SearchBar
                    value={search}
                    onChange={(e) => setSearch(e.target.value)}
                />

                <div className="grid md:grid-cols-3 gap-8 mt-8">
                    {products.length === 0 ? (
                        <h2 className="col-span-3 text-center text-slate-500 text-2xl">
                            No Products Found
                        </h2>
                    ) : (
                        products.map((product) => (
                            <ProductCard
                                key={product._id}
                                product={product}
                                onAdd={addToCart}
                            />
                        ))
                    )}
                </div>

            </div>

        </div>
    );
}

export default Home;