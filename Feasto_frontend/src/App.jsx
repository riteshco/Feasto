import './index.css'
import { CategoriesPage } from './pages/Categories'
import { Landing } from './pages/Landing'
import { createBrowserRouter, Outlet, RouterProvider } from "react-router-dom";
import { AdminRoute, ProtectedRoute } from './components/ProtectedRoute';
import { RoleBasedRoute } from '@/components/RoleBasedRoute'
import { CartPage } from './pages/Cart'
import { UserOrderPage } from './pages/UserOrders'
import { OrderDetailPage } from './pages/OrderDetail'
import { OrderPaymentPage } from './pages/OrderPayment'
import { PastOrdersPage } from './pages/PastOrders'
import { AddFoodPage } from './pages/AddFood'
import { AdminDashboard } from './pages/admin/AdminDashboard'
import { DataTableDemo } from './pages/admin/Allusers'
import { AllOrdersPage } from './pages/admin/AllOrders'
import { AllPaymentsPage } from './pages/admin/AllPayments'
import Layout from './components/Layout';
import { ChangeRolePage } from './pages/ChangeRole';
import { DeliveredOrdersPage } from './pages/DeliveredOrders';

function App() {

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Layout />,
      children:
        [
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
                <AdminDashboard />
              </AdminRoute>
            )
          },
          {
            path: "/users",
            element: (
              <AdminRoute>
                <DataTableDemo />
              </AdminRoute>
            )
          },
          {
            path: "/all-orders",
            element: (
              <AdminRoute>
                <AllOrdersPage />
              </AdminRoute>
            )
          },
          {
            path: "/all-payments",
            element: (
              <AdminRoute>
                <AllPaymentsPage />
              </AdminRoute>
            )
          },
          {
            path: "/categories",
            element: (
              <ProtectedRoute>
                <CategoriesPage />,
              </ProtectedRoute>
            )
          },
          {
            path: "/cart",
            element: (
              <ProtectedRoute>
                <CartPage />
              </ProtectedRoute>
            )
          },
          {
            path: "/orders",
            element: (
              <ProtectedRoute>
                <UserOrderPage />
              </ProtectedRoute>
            )
          },
          {
            path: "/order",
            element: <ProtectedRoute><Outlet /></ProtectedRoute>,
            children: [
              { path: "items/:id", element: <OrderDetailPage /> },
              { path: "payment/:id", element: <OrderPaymentPage /> },
            ]
          },
          {
            path: "/past-orders",
            element: (
              <ProtectedRoute>
                <PastOrdersPage />
              </ProtectedRoute>
            )
          },
          {
            path: "/add-food",
            element: (
              <ProtectedRoute role1="admin" role2="chef">
                <AddFoodPage />
              </ProtectedRoute>
            )
          },
          {
            path: "/delivered-orders",
            element: (
              <ProtectedRoute role1="chef">
                <DeliveredOrdersPage />
              </ProtectedRoute>
            )
          },
          {
            path: "/change-role",
            element: (
              <ProtectedRoute>
                <ChangeRolePage />
              </ProtectedRoute>
            )
          },
        ]
    }]
  )

  return <RouterProvider router={router} />;
}

export default App
