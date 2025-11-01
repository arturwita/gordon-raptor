import {
  createContext,
  useContext,
  useState,
  type ReactNode,
  type FC,
} from "react";
import { jwtDecode } from "jwt-decode";
import { UserRole, type UserJwtPayload } from "../types/user";
import { AppRoutes } from "../routes/AppRoutes";

interface AuthContextType {
  user: UserJwtPayload | null;
  token: string | null;
  isAdmin: boolean;
  login: (token: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [user, setUser] = useState<UserJwtPayload | null>(() => {
    const storedUser = localStorage.getItem("user");
    return storedUser ? JSON.parse(storedUser) : null;
  });

  const [token, setToken] = useState<string | null>(() => {
    return localStorage.getItem("token");
  });

  const login = (jwtToken: string) => {
    const decoded = jwtDecode<UserJwtPayload>(jwtToken);
    setUser(decoded);
    setToken(jwtToken);
    localStorage.setItem("user", JSON.stringify(decoded));
    localStorage.setItem("token", jwtToken);
  };

  const logout = () => {
    setUser(null);
    setToken(null);
    localStorage.removeItem("user");
    localStorage.removeItem("token");

    window.location.href = AppRoutes.Home;
  };

  const isAdmin = user?.role === UserRole.Admin;

  return (
    <AuthContext.Provider value={{ user, token, login, logout, isAdmin }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = (): AuthContextType => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within AuthProvider");
  return context;
};
