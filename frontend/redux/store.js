import { configureStore } from "@reduxjs/toolkit";
import cartReducer from "./cartSlice";

export const store=configureStore({

    //telling the component to use this reducer which is named as cart to manage your cart section of global state
    reducer:{
        cart:cartReducer
    }
})