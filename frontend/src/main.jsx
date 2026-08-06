import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";
import "./index.css";
import {Provider} from "react-redux";
import { store } from "../redux/store";
import AuthProvider from "./context/AuthContext";

ReactDOM.createRoot(document.getElementById("root")).render(
  <BrowserRouter>
    <AuthProvider>
        <Provider store={store}>
              <App />
        </Provider>      
    </AuthProvider>
  </BrowserRouter>
);