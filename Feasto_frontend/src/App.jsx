import './index.css'
import { Navbar } from './components/Navbar'
import { CustomerHome } from './pages/home/CustomerHome'
import { CategoriesPage } from './pages/Categories'
import { Landing } from './pages/Landing'
import ReactDOM from "react-dom/client";
import { createBrowserRouter, Routes, Route, RouterProvider } from "react-router-dom";
import { AdminRoute , ProtectedRoute } from './components/ProtectedRoute';
import { RoleBasedRoute } from '@/components/RoleBasedRoute'
import { CartPage } from './pages/Cart'
import { UserOrderPage } from './pages/UserOrders'
import { OrderDetailPage } from './pages/OrderDetail'
import { OrderPaymentPage } from './pages/OrderPayment'
import { PastOrdersPage } from './pages/PastOrders'
import { AddFoodPage } from './pages/AddFood'

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
  },
  {
    path: "/categories",
    element: (
    <ProtectedRoute>
    <CategoriesPage/>,
    </ProtectedRoute>
    )
  },
  {
    path: "/cart",
    element: (
    <ProtectedRoute>
    <CartPage/>
    </ProtectedRoute>
    )
  },
  {
    path: "/orders",
    element: (
      <ProtectedRoute>
        <UserOrderPage/>
      </ProtectedRoute>
    )
  },
  {
    path: "/order/items/:id",
    element: (
      <ProtectedRoute>
        <OrderDetailPage/>
      </ProtectedRoute>
    )
  },
  {
    path: "/order/payment/:id",
    element: (
      <ProtectedRoute>
        <OrderPaymentPage/>
      </ProtectedRoute>
    )
  },
  {
    path: "/past-orders",
    element: (
      <ProtectedRoute>
        <PastOrdersPage/>
      </ProtectedRoute>
    )
  },
  {
    path: "/add-food",
    element: (
      <ProtectedRoute role1="admin" role2="chef">
        <AddFoodPage/>
      </ProtectedRoute>
    )
  }
  ])

  return <RouterProvider router={router} />;
}

export default App
