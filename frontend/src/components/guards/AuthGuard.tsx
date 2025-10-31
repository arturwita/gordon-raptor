import { Navigate } from "react-router";
import { useAuth } from "@/context/AuthContext";
import { AppRoutes } from "../../routes/AppRoutes";
import type { ReactNode } from "react";

export const AuthGuard = ({ children }: { children: ReactNode }) => {
  const { user } = useAuth();

  if (!user) {
    return <Navigate to={AppRoutes.Home} replace />;
  }

  return <>{children}</>;
};
