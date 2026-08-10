import {createSlice, createAsyncThunk} from "@reduxjs/toolkit"
import api from "../src/api/axios"

export const loadCart= createAsyncThunk("cart/loadCart",

    async ()=>{
    const id = localStorage.getItem("userId");

    const res = await api.get(`/get-cart?id=${id}`);
        
   console.log("function called and returned value", res.data);
    return res.data; //automatically becomes the payload of the calling fucntion
    }
    
)
const cartSlice= createSlice({
    //identifies the slice

    name:"cart", 
    //replaces all the state var
    initialState:{
        items:[],
        total:0,
        loading:false,
        error:null
    },
    //name helps to access as state.cart.items ///so we need the slice name
    reducers:{},
    //here our operations are async hence not using normal reducers
    extraReducers:(builder)=>{
        builder.addCase(
            loadCart.pending,(state)=>{state.loading=true;
            state.error=null;

        })

         builder.addCase(
            loadCart.fulfilled,(state,action)=>{
            state.loading=false;
            //like setcart and settotal
            state.items= action.payload.cart;
            state.total=action.payload.total;

            console.log(state.items);
        })

           builder.addCase(
            loadCart.rejected,(state)=>{
            state.loading=false;
          state.error=action.error.message

        })
    }

})

export default cartSlice.reducer;