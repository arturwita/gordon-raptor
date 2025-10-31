import { createBrowserRouter } from "react-router";
import Home from "../pages/Home";
import Recipes from "../pages/Recipes";
import { LoginCallback } from "../components/callbacks/LoginCallback";

export const AppRoutes = {
  Home: "/",
  Recipes: "/recipes",
  LoginCallback: "/login/google/callback",
} as const;

export const router = createBrowserRouter([
  {
    path: AppRoutes.Home,
    Component: Home,
  },
  {
    path: AppRoutes.LoginCallback,
    Component: LoginCallback,
  },
  {
    path: AppRoutes.Recipes,
    Component: Recipes, // todo: secure with AuthGuard
  },
]);
