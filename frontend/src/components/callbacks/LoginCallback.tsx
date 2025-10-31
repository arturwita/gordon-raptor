import { useEffect, type FC } from "react";
import { useNavigate, useSearchParams } from "react-router";
import { AppRoutes } from "@/routes/AppRoutes";
import { useAuth } from "../../context/AuthContext";

export const LoginCallback: FC = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const { login } = useAuth();

  useEffect(() => {
    const token = searchParams.get("token");
    if (!token) {
      navigate(AppRoutes.Home);
      return;
    }

    try {
      login(token);
      navigate(AppRoutes.Recipes);
    } catch (err) {
      console.error("Invalid token:", err);
      navigate(AppRoutes.Home);
    }
  }, [navigate, searchParams]);

  return (
    <div className="flex items-center justify-center h-screen">
      <p className="text-gray-600 dark:text-gray-300">
        Logging you in... please wait
      </p>
    </div>
  );
};
