import React from "react";
import { Navigate } from "react-router-dom";

import { CustomerHome } from "@/pages/home/CustomerHome";
import { getUserFromToken } from "@/utils/auth";
import { ProtectedRoute } from "./ProtectedRoute";

export function RoleBasedRoute() {
  const user = getUserFromToken()

  if (!user) {
    return <Navigate to="/" replace />;
  }

  try {
    const role = user.user_role

    switch (role) {
      case "admin":
        return <ProtectedRoute/>;
      case "chef":
        return <ChefHome />;
      case "customer":
        return (
          <ProtectedRoute >
            <CustomerHome />
          </ProtectedRoute>
      )
      default:
        return <h1>Unauthorized</h1>;
    }
  } catch (err) {
    return <Navigate to="/" replace />;
  }
}
