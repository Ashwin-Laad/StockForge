import { createContext, useState } from "react";

export const AuthContext = createContext();

function AuthProvider({ children }) {

  const [token, setToken] = useState(
    localStorage.getItem("token") || ""
  );

  function login(jwt, userID) {
    localStorage.setItem("token", jwt);
    localStorage.setItem("userID", userID);

    setToken(jwt);
  }

  function logout() {
    localStorage.removeItem("token");
    setToken("");
  }

  return (
    <AuthContext.Provider
      value={{
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