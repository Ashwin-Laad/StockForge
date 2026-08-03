import axios from "axios";

const api = axios.create({
    baseURL: "http://localhost:8080",
});

//for the request object to inject the token in it so backend can verify

api.interceptors.request.use((config) => {
    const token = localStorage.getItem("token");

    if (token) {
        config.headers.Token = token;
    }

    return config;
});

export default api;