import { createBrowserRouter } from "react-router";
import Test from "./pages/Test";
import LandingPage from "./pages/Home";

export const AppRoutes = {
  Home: "/",
  Login: "/login",
} as const;

export const router = createBrowserRouter([
  {
    path: AppRoutes.Home,
    Component: LandingPage,
  },
  {
    path: AppRoutes.Login,
    Component: Test,
  },
]);
