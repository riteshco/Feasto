import './index.css'
import { Navbar } from './components/Navbar'
import { CustomerHome } from './pages/home/CustomerHome'
import { Landing } from './pages/Landing'
import ReactDOM from "react-dom/client";
import { createBrowserRouter, Routes, Route, RouterProvider } from "react-router-dom";
import { AdminRoute , ProtectedRoute } from './components/ProtectedRoute';
import { RoleBasedRoute } from '@/components/RoleBasedRoute'

function App() {

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Landing />,
    },
    {
      path: "/home",
      element: <RoleBasedRoute />,
    },
    {
    path: "/admin",
    element: (
      <AdminRoute>
        {/* <AdminDashboard /> */}
      </AdminRoute>
    )
  }
  ])

  return <RouterProvider router={router} />;
}

export default App
