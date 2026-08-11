import {createSlice} from "@reduxjs/toolkit"
import { useNavigate } from "react-router-dom";
function navigateToLogin() {
  const navigate = useNavigate();
  navigate("/login");
}
const authSlice = createSlice({
  name:"auth",
  initialState:{
    token:localStorage.getItem("token") || "",
    userId: localStorage.getItem("userId") || "",
    
  },

  //now define function which can actually change sthe state --> reducers
  reducers:{
    login: (state, action) =>{
      const {jwt, userId}= action.payload;
      localStorage.setItem("token", jwt);
      localStorage.setItem("userId", userId);

      state.token= jwt;
      state.userId=userId;
    

    },

      logout: (state) => {
          
            localStorage.removeItem("token");
            localStorage.removeItem("userId");

            state.token = "";
            state.userId = "";
         

          

        }
  }
})

export const {login, logout}= authSlice.actions; //i have to export these actions
//export reducer auto injects it 
export default authSlice.reducer;

// export default {
//     reducer: authSlice.reducer,
//     actions: authSlice.actions
// };

// Then:

// import auth from "./authSlice";

// dispatch(
//     auth.actions.login({
//         jwt: response.data.token,
//         userID: response.data.userID
//     })
