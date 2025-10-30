import { createBrowserRouter } from "react-router";
import App from "./App";
import Test from "./pages/Test";

export const AppRoutes = {
  home: "/",
  test: "/abc",
} as const;

export const router = createBrowserRouter([
  {
    path: AppRoutes.home,
    Component: App,
  },
  {
    path: AppRoutes.test,
    Component: Test,
  },
]);
