import { createBrowserRouter } from "react-router";
import Home from "../pages/Home";
import Recipes from "../pages/RecipesPage";

export const AppRoutes = {
  Home: "/",
  Recipes: "/recipes",
} as const;

export const router = createBrowserRouter([
  {
    path: AppRoutes.Home,
    Component: Home,
  },

  {
    path: AppRoutes.Recipes,
    Component: Recipes,
  },
]);
