import { createContext, useState } from "react";

export const AuthContext = createContext();

function AuthProvider({ children }) {


  //what ever stored on the browser 
  const [token, setToken] = useState(
    localStorage.getItem("token") || ""
  );

  //send from the backend and called

  function login(jwt, userID) {
    localStorage.setItem("token", jwt);
    localStorage.setItem("userID", userID);

    setToken(jwt);
  }
//requires a fresh login
  function logout() {
    localStorage.removeItem("token");
    setToken("");
  }

  return (
    <AuthContext.Provider
      value={{
        //any children can actually access this token
        token,
        login,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export default AuthProvider;